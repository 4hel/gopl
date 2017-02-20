package main

import (
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
	{"Ready 2 Go", "Martin Solveig", "Smash", 2007, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var html = template.Must(template.ParseFiles("tracks"))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	var sorting sort.Interface

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if len(r.Form["sort"]) > 0 {
		switch r.Form["sort"][0] {
		case "year":
			sorting = customSort{tracks, func(x, y *Track) bool {
				return x.Year < y.Year
			}}

		case "artist":
			sorting = customSort{tracks, func(x, y *Track) bool {
				return x.Artist < y.Artist
			}}
		case "title":
			sorting = customSort{tracks, func(x, y *Track) bool {
				return x.Title < y.Title
			}}
		case "album":
			sorting = customSort{tracks, func(x, y *Track) bool {
				return x.Album < y.Album
			}}
		case "length":
			sorting = customSort{tracks, func(x, y *Track) bool {
				return x.Length < y.Length
			}}
		}
	}

	if sorting != nil {
		sort.Sort(sorting)
	}

	if err := html.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}
