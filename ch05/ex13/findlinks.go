package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopl.io/ch05/links"
)

const downloadDir = "output"

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(link string) []string {
	originURL, err := url.Parse(link)
	if err != nil {
		log.Print(err)
		return nil
	}
	fmt.Println(link)
	list, err := links.Extract(link)
	if err != nil {
		log.Print(err)
	}
	for _, l := range list {
		fmt.Print(l)
		linkURL, err := url.Parse(l)
		if err != nil {
			log.Print(err)
			continue
		}
		if linkURL.Host == originURL.Host {
			download(linkURL)
		}
	}
	return list
}

func download(url *url.URL) {
	var filename string
	var directory string
	path := url.EscapedPath()
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 {
		filename = path
		directory = ""
	} else {
		filename = path[lastSlash+1:]
		directory = strings.TrimLeft(path[:lastSlash], "/")
		d := downloadDir + string(os.PathSeparator) + directory
		if err := os.MkdirAll(d, 0755); err != nil {
			log.Print(err)
			return
		}
	}
	if filename == "" {
		filename = "index.html"
	}
	target := strings.Join([]string{downloadDir, directory, filename},
		string(os.PathSeparator))
	fmt.Println("\n save to: ", target)
	f, err := os.Create(target)
	defer f.Close()
	if err != nil {
		log.Print(err)
		return
	}
	resp, err := http.Get(url.String())
	if err != nil {
		log.Print(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Print(resp.Status)
		return
	}
	defer resp.Body.Close()
	io.Copy(f, resp.Body)
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	initOutputDirectory()
	breadthFirst(crawl, os.Args[1:])
}

func initOutputDirectory() {
	if err := os.RemoveAll(downloadDir); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	if err := os.Mkdir(downloadDir, 0755); err != nil {
		log.Fatal(err)
	}
}
