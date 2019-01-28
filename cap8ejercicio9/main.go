package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"time"

	

type Root struct {
	path  string
	files int64
	bytes int64
}

func main() {
	
	var roots []Root
	for _, rootPath := range os.Args[1:] {
		roots = append(roots, Root{rootPath, 0, 0})
	}
	if len(roots) == 0 {
		roots = []Root{Root{".", 0, 0}}
	}

	go func() {
		for {
			var n sync.WaitGroup
			for i := range roots {
				n.Add(1)
				i := i
				go func() {
					
					fileSizes := make(chan int64)
					var m sync.WaitGroup
					m.Add(1)
					go func() {
						walkDir(roots[i].path, &m, fileSizes)
					}()
					go func() {
						m.Wait()
						close(fileSizes)
						n.Done()
					}()
					var nfiles, nbytes int64
					for {
						size, ok := <-fileSizes
						if !ok {
							break
						}
						nfiles++
						nbytes += size
					}
					roots[i].bytes = nbytes
					roots[i].files = nfiles
				}()
			}
			n.Wait()
			<-time.After(500 * time.Millisecond)
		}
	}()

	for {
		<-time.After(500 * time.Millisecond)
		printDiskUsage(roots) 
	}
}

func printDiskUsage(roots []Root) {
	var rows [][]string

	for _, root := range roots {
		rows = append(rows, []string{
			root.path,
			fmt.Sprintf("%d files", root.files),
			fmt.Sprintf("%.1f GB", float64(root.bytes)/1e9),
		})
	}

	console.Clear()
	fmt.Println(console.SprintTable(rows))
}



func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}


var sema = make(chan struct{}, 20)


func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        
	defer func() { <-sema }() 

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}