package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"gopl.io/ch04/ex12"
)

const (
	indexFile            = "indexkcd.json"
	scoreTitleMatch      = 25
	scoreAltMatch        = 5
	scoreTranscriptMatch = 1
)

type SearchResult struct {
	xkcd  ex12.XKCD
	score int
}

func (s SearchResult) String() string {
	return fmt.Sprintf("Score: %d\nURL: %s\nTranscript: %s\n",
		s.score, s.xkcd.URL, s.xkcd.Transcript)
}

func main() {
	var index ex12.Index
	var found []SearchResult
	if len(os.Args) < 2 {
		fail("usage: %s [search term]\n", os.Args[0])
	}
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		fail("index file %q does not exist\n", indexFile)
	}
	indexFile, err := os.Open(indexFile)
	if err != nil {
		fail("unable to open %s: %v\n", indexFile, err)
	}
	decoder := json.NewDecoder(bufio.NewReader(indexFile))
	if err := decoder.Decode(&index); err != nil {
		fail("error decoding JSON index: %v\n", err)
	}
	for _, xkcd := range index.Entries {
		s := score(xkcd, os.Args[1])
		if s > 0 {
			res := SearchResult{xkcd: xkcd, score: s}
			found = append(found, res)
		}
	}
	sort.Slice(found, func(i, j int) bool {
		return found[i].score > found[j].score
	})
	for _, res := range found {
		fmt.Println(res)
	}
}

func score(entry ex12.XKCD, searchTerm string) int {
	var score int
	if contains(entry.Title, searchTerm) {
		score += scoreTitleMatch
	}
	if contains(entry.Alt, searchTerm) {
		score += scoreAltMatch
	}
	if contains(entry.Transcript, searchTerm) {
		score += scoreTranscriptMatch
	}
	return score
}

func contains(haystack, needle string) bool {
	haystack, needle = strings.ToLower(haystack), strings.ToLower(needle)
	words := strings.Fields(haystack)
	for _, word := range words {
		if word == needle {
			return true
		}
	}
	return false
}

func fail(format string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, format, params)
	os.Exit(1)
}
