package main

import (
	"fmt"
	"log"

	
)

func main() {
	expr, err := eval.Parse("x ? pow(x, 3) : pow(y, 3)")
	if err != nil {
		log.Fatalf("ch07/ex14: %v", err)
	}
	fmt.Printf("%s\n", expr)
}