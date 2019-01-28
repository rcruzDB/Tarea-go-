package main

import (
	"fmt"
	"os"

	
)

func main() {
	err := omdb.GetPoster(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/ex12: %v\n", err)
		os.Exit(1)
	}
}