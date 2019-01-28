package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string            
	var attrs []map[string]string 
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) 
			attr := make(map[string]string)
			for _, a := range tok.Attr {
				attr[a.Name.Local] = a.Value
			}
			attrs = append(attrs, attr)
		case xml.EndElement:
			stack = stack[:len(stack)-1] 
			attrs = attrs[:len(attrs)-1]
		case xml.CharData:
			if containsAll(toStringSlice(stack, attrs), os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}



func toStringSlice(stack []string, attrs []map[string]string) []string {
	result := []string{}
	for i := range stack {
		result = append(result, stack[i])
		for k, v := range attrs[i] {
			result = append(result, k+"="+v)
		}
	}
	return result
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}