package main

import (
	"fmt"
	"io"
	"io/ioutil"
)


type ByteCounter struct {
	writer io.Writer
	count  int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	c.writer.Write(p)
	c.count += int64(len(p))
	return len(p), nil
}


func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{
		writer: w,
		count:  0,
	}
	return &bc, &bc.count
}

func main() {
	w, count := CountingWriter(ioutil.Discard)
	fmt.Fprint(w, "hello")
	fmt.Println(*count) 
}