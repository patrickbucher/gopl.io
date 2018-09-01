package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const (
	omdbURLFile   = ".omdb_url"
	posterURLFile = ".poster_url"
)

type Result struct {
	Entries []Movie `json:"Search"`
}
type Movie struct {
	Id     string `json:"imdbID"`
	Title  string
	Poster string
}

func main() {
	if len(os.Args) < 2 {
		fail("usage: %s [search term]\n", os.Args[0])
	}
	searchTerm := url.QueryEscape(os.Args[1])
	b, err := ioutil.ReadFile(omdbURLFile)
	if err != nil {
		fail("unable to read file %s: %v\n", omdbURLFile, err)
	}
	baseURL := strings.TrimSpace(string(b))
	omdbURL := fmt.Sprintf("%s&type=movie&r=json&s=%s", baseURL, searchTerm)
	resp, err := get(omdbURL)
	if err != nil {
		fail("%v\n", err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(bufio.NewReader(resp.Body))
	var result Result
	if err := decoder.Decode(&result); err != nil {
		fail("decoding JSON: %v\n", err)
	}
	removeNonAscii := regexp.MustCompile("[[:^ascii:]]")
	removeWhitespace := regexp.MustCompile("[[:space:]]")
	removeNonAlnum := regexp.MustCompile("[[:^alnum:]]")
	for _, e := range result.Entries {
		if strings.TrimSpace(e.Poster) == "" {
			continue
		}
		title := removeNonAscii.ReplaceAllLiteralString(e.Title, "")
		title = removeWhitespace.ReplaceAllLiteralString(title, "-")
		title = removeNonAlnum.ReplaceAllLiteralString(title, "")
		fileName := fmt.Sprintf("%s_%s.jpg", title, e.Id)
		resp, err := get(e.Poster)
		if err != nil {
			fmt.Fprintf(os.Stderr, "retrieving poster %s: %v\n", e.Poster, err)
			continue
		}
		defer resp.Body.Close()
		f, err := os.Create(fileName)
		defer f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "create file %q: %v\n", fileName, err)
			continue
		}
		io.Copy(f, resp.Body)
	}
}

func get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %v", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s", url, resp.Status)
	}
	return resp, nil
}

func fail(format string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, format, params)
	os.Exit(1)
}
