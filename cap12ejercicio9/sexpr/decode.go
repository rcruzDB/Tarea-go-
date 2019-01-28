package sexpr

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)


type lexer struct {
	scan  scanner.Scanner
	token rune 
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { 
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

























func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		
		
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) 
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) 
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() 
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}




func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: 
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: 
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: 
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: 
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}



type Decoder struct {
	lex   *lexer
	err   error
	depth int
}

func NewDecoder(r io.Reader) *Decoder {
	lex := new(lexer)
	lex.scan.Init(r)
	return &Decoder{lex: lex}
}

func (d *Decoder) Token() (Token, error) {
	d.lex.next()
	if d.err != nil {
		return nil, d.err
	}
	text := d.lex.text()
	switch d.lex.token {
	case scanner.Ident:
		return Symbol(text), nil
	case scanner.Int:
		n, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return nil, err
		}
		return Int(n), nil
	case scanner.String:
		return text[1 : len(text)-1], nil
	case '(':
		d.depth++
		return StartList{}, nil
	case ')':
		d.depth--
		return EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected token %q", text)
	}
}

type Token interface{}

type Symbol string 
type String string 
type Int int
type StartList struct{}
type EndList struct{}

func (i Int) String() string {
	return i.String()
}

func (s StartList) String() string {
	return "StartList"
}

func (s EndList) String() string {
	return "EndList"
}