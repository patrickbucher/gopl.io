package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byAlbum []*Track

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Multisort struct {
	criteria []sort.Interface
}

func (m Multisort) apply(criterion sort.Interface) {
	criteria := make([]sort.Interface, len(m.criteria))
	// recycle the criteria used before
	for _, s := range m.criteria {
		if criterion != s {
			// except the latest
			criteria = append(criteria, s)
		}
	}
	// use the fresh criterion at the end of the sort chain
	criteria = append(criteria, criterion)
	for _, s := range criteria {
		sort.Sort(s)
	}
	m.criteria = criteria
}

var tableHTML = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Exercise 7.9: Table Multisort</title>
		<style type="text/css">
		td { padding: 0.5em; }
		</style>
	</head>
	<body>
		<h1>Exercise 7.9: Table Multisort</h1>
		<table>
			<tr>
				<th><a href="/tracks?sort=title">Title</a></th>
				<th><a href="/tracks?sort=artist">Artist</a></th>
				<th><a href="/tracks?sort=album">Album</a></th>
				<th><a href="/tracks?sort=year">Year</a></th>
				<th><a href="/tracks?sort=length">Length</a></th>
			</tr>
			{{ range . }}
			<tr>
				<td>{{ .Title }}</td>
				<td>{{ .Artist}}</td>
				<td>{{ .Album}}</td>
				<td>{{ .Year }}</td>
				<td>{{ .Length }}</td>
			</tr>
			{{ end }}
		</table>
	</body>
</html>`

var tableTemplate = template.Must(template.New("tracks").Parse(tableHTML))

func main() {
	var multisort Multisort
	http.HandleFunc("/tracks", func(w http.ResponseWriter, r *http.Request) {
		sortBy := r.FormValue("sort")
		switch sortBy {
		case "title":
			multisort.apply(byTitle(tracks))
		case "artist":
			multisort.apply(byArtist(tracks))
		case "album":
			multisort.apply(byAlbum(tracks))
		case "year":
			multisort.apply(byYear(tracks))
		case "length":
			multisort.apply(byLength(tracks))
		default:
			fmt.Fprintf(w, "sorting criterion '%s' is not supported", sortBy)
		}
		if err := tableTemplate.Execute(w, tracks); err != nil {
			fmt.Fprintln(w, err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
