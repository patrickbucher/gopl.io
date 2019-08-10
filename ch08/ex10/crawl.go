package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gopl.io/ch08/ex10/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	time.Sleep(100 * time.Millisecond) // give user a chance to cancel
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

	// Cancellation: user enters single character to stop
	var wg sync.WaitGroup
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		fmt.Fprintln(os.Stderr, "user cancelled")
		close(done)
		links.Cancel()
		wg.Wait()
		close(worklist)
	}()

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- workItem{urls: flag.Args(), level: 0}
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		select {
		case <-done:
			// drain worklist
			for range worklist {
			}
		case item := <-worklist:
			if item.level > *depth {
				continue
			}
			for _, link := range item.urls {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						defer wg.Done()
						wg.Add(1)
						if !cancelled() {
							worklist <- workItem{urls: crawl(link), level: item.level + 1}
						}
					}(link)
				}
			}
		}
	}
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
