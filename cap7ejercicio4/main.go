package main

import (
	"io"
	"log"
	"os"

	"github.com/kdama/gopl/ch05/ex07/prettyhtml"
	"golang.org/x/net/html"
)


type StringReader string

func (s *StringReader) Read(p []byte) (int, error) {
	copy(p, *s)
	return len(*s), io.EOF
}


func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}

func main() {
	doc, err := html.Parse(NewReader("<p>Hello, world!</p>"))
	if err != nil {
		log.Fatalf("ch07/ex04: %v", err)
	}
	prettyhtml.WriteHTML(os.Stdout, doc)
}