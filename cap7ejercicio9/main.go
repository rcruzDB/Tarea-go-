package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/kdama/gopl/ch07/ex08/sorting"
)


func printTracks(w io.Writer, tracks []*sorting.Track) {
	tracksTemplate.Execute(w, tracks)
}

var tracksTemplate = template.Must(template.New("tracks").Parse(`
<h1>{{len .}} track{{if ne (len .) 1}}s{{end}}</h1>
<table>
<tr style='text-align: left'>
<th><a href='/?sortby=Title'>Title</th>
<th><a href='/?sortby=Artist'>Artist</th>
<th><a href='/?sortby=Album'>Album</th>
<th><a href='/?sortby=Year'>Year</th>
<th><a href='/?sortby=Length'>Length</th>
</tr>
{{range .}}
<tr>
<td>{{.Title}}</td>
<td>{{.Artist}}</td>
<td>{{.Album}}</td>
<td>{{.Year}}</td>
<td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func getTracks() []*sorting.Track {
	return []*sorting.Track{
		{"Go", "Delilah", "From the Roots Up", 2012, sorting.Length("3m38s")},
		{"Go", "Moby", "Moby", 1992, sorting.Length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, sorting.Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, sorting.Length("4m24s")},
	}
}

var columns = []string{}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http:
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	sortby := r.URL.Query().Get("sortby")
	columns = pushAsSet(columns, sortby)

	tracks := getTracks()
	sort.Sort(sorting.MultiSort(tracks, columns))
	printTracks(w, tracks)
}



func pushAsSet(slice []string, s string) []string {
	return append(removeString(slice, s), s)
}


func removeString(slice []string, s string) []string {
	for i := range slice {
		if slice[i] == s {
			return remove(slice, i)
		}
	}
	return slice
}


func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}