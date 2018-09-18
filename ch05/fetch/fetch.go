package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [urls]\n", os.Args[0])
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		filename, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetching %s: %v\n", url, err)
			continue
		}
		fmt.Printf("stored %s to %s (%d bytes)\n", url, filename, n)
	}
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
