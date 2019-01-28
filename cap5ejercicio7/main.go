package main

import (
	"net/http"
	"os"

	"github.com/kdama/gopl/ch05/ex07/prettyhtml"
	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	prettyhtml.WriteHTML(os.Stdout, doc)
	return nil
}