package main

import (
	"fmt"


)

func main() {
	s := &intset.IntSet{}

	s.Add(1)
	s.Add(2)
	s.Add(42)

	for i, e := range s.Elems() {
		fmt.Printf("%d\t%d\n", i, e)
	}
}