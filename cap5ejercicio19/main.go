package main

import (
	"fmt"
)

func get(in interface{}) (out interface{}) {
	out = in
	defer func() { recover() }()
	panic(in)
}

func main() {
	fmt.Println(get("Hello, world!"))
}