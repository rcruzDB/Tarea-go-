package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)


type LimitedByteCounter struct {
	reader io.Reader
	
	rest int64
}

func (c *LimitedByteCounter) Read(p []byte) (int, error) {
	length := int64(len(p))
	if c.rest < length {
		length = c.rest
	}
	c.rest -= length

	n, err := c.reader.Read(p[:length])
	if c.rest == 0 {
		return n, io.EOF
	}
	return n, err
}


func LimitReader(r io.Reader, n int64) io.Reader {
	lbc := LimitedByteCounter{
		reader: r,
		rest:   n,
	}
	return &lbc
}

func main() {
	r := strings.NewReader("Hello, world!")
	s, err := ioutil.ReadAll(LimitReader(r, 5))
	if err != nil {
		log.Fatalf("ch07/ex05: %v", err)
	}
	fmt.Println(string(s)) 
}