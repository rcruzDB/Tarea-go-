package main

import (
	"fmt"

	"github.com/kdama/gopl/ch06/ex03/intset"
)

func main() {
	s := &intset.IntSet{}
	t := &intset.IntSet{}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	t.Add(1)
	t.Add(4)

	fmt.Println(s) 

	s.IntersectWith(t)
	fmt.Println(s) 

	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.DifferenceWith(t)
	fmt.Println(s) 

	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.SymmetricDifference(t)
	fmt.Println(s) 
}