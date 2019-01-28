package main

import (
	"fmt"

	
)

func main() {
	is := &intset.IntSet{}

	is.Add(1)

	fmt.Println(is) 

	is.AddAll(2, 3, 4)

	fmt.Println(is) 
}