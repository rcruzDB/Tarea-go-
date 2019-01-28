package main

import (
	"fmt"

	
)

func main() {
	is := &intset.IntSet{}

	is.Add(1)
	is.Add(2)
	is.Add(3)
	is.Add(42)

	fmt.Println(is)       
	fmt.Println(is.Len()) 

	fmt.Println(is) 
	is.Remove(2)
	fmt.Println(is) 

	is2 := is.Copy()
	is.Clear()
	fmt.Println(is)  
	fmt.Println(is2) 