package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
)

var cancel = make(chan struct{})

func cancelled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

func main() {
	flag.Parse()
	urls := flag.Args()
	responses := make(chan string, len(urls))
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "build request to %s: %v\n", url, err)
				return
			}
			req.Cancel = cancel
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			defer resp.Body.Close()
			responses <- url
			fmt.Println("finished lookup of", url)
		}(url)
	}
	for {
		select {
		case url := <-responses:
			fmt.Println(url, "is the fastest")
			close(cancel)
			wg.Wait()
			return
		}
	}
}
