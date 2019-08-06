package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch05/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

type workItem struct {
	urls  []string
	level int
}

func main() {
	worklist := make(chan workItem)
	var n int // number of pending sends to worklist

	depth := flag.Int("depth", 3, "maximum number of link levels to follow")
	flag.Parse()
	if *depth < 0 {
		fmt.Fprintf(os.Stderr, "depth=%d: nothing to do\n", *depth)
		os.Exit(1)
	}

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- workItem{urls: flag.Args(), level: 0}
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		item := <-worklist
		if item.level > *depth {
			continue
		}
		for _, link := range item.urls {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- workItem{urls: crawl(link), level: item.level + 1}
				}(link)
			}
		}
	}
}
