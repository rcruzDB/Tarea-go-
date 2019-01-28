package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fetchToDir(os.Args[1:], "out/1")
	fetchToDir(os.Args[1:], "out/2")
}


func fetchToDir(urls []string, dirName string) {
	start := time.Now()
	ch := make(chan string)
	for idx, url := range urls {
		file, err := os.Create(fmt.Sprintf("%s/%d.txt", dirName, idx))
		if err != nil {
			fmt.Println(err) 
			return
		}
		defer file.Close()
		go fetch(url, file, ch) 
	}
	for range urls {
		fmt.Println(<-ch) 
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}


func fetch(url string, writer io.Writer, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) 
		return
	}

	nbytes, err := io.Copy(writer, resp.Body)
	resp.Body.Close() 
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}