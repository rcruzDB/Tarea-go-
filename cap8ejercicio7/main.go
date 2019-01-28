package main

import (
	"os"

	"github.com/kdama/gopl/ch08/ex07/crawl"
)


const out = "./out"

func main() {
	worklist := make(chan []string)
	var n int 

	
	n++
	go func() { worklist <- os.Args[1:] }()

	
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl.Crawl(link, out)
				}(link)
			}
		}
	}
}