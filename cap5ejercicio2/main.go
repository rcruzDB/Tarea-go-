package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch05/ex02: %v\n", err)
		os.Exit(1)
	}
	for key, value := range visit(make(map[string]int), doc) {
		fmt.Printf("%s: %d\n", key, value)
	}
}


func visit(counts map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return counts
	}
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	visit(counts, n.FirstChild)
	visit(counts, n.NextSibling)
	return counts
}