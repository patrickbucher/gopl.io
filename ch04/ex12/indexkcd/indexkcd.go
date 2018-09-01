package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"gopl.io/ch04/ex12"
)

const (
	urlTemplate   = "https://xkcd.com/%d/info.0.json"
	indexFile     = "indexkcd.json"
	throttleAfter = 100
	throttlePause = 100 * time.Millisecond
)

var errNotFound error = errors.New("Not Found")

type Result struct {
	xkcd *ex12.XKCD
	err  error
}

var wg sync.WaitGroup

func main() {
	var entries []ex12.XKCD
	results := make(chan Result)
	done := make(chan bool)
	quitted := false
	var started int
	go func() {
		finished := false
		for number := 1; !finished; number++ {
			// sending out requests happens way faster than getting responsees
			if number%throttleAfter == 0 {
				time.Sleep(throttlePause)
			}
			select {
			case <-done:
				finished = true
			default:
				wg.Add(1)
				started++
				go fetchXKCD(number, results)
			}
		}
		wg.Wait()
		close(results)
	}()
	for result := range results {
		if result.err == errNotFound {
			if !quitted {
				done <- true
			}
			quitted = true
		} else if result.err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", result.err)
			os.Exit(1)
		} else if result.xkcd != nil {
			entries = append(entries, *result.xkcd)
		}
	}
	index := ex12.Index{Entries: entries}
	persist(&index, indexFile)
}

func fetchXKCD(number int, results chan<- Result) {
	defer wg.Done()
	url := fmt.Sprintf(urlTemplate, number)
	resp, err := http.Get(url)
	if err != nil {
		results <- Result{
			xkcd: nil,
			err:  fmt.Errorf("GET %s: %v\n", url, err),
		}
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		xkcd := ex12.XKCD{}
		if err := decoder.Decode(&xkcd); err != nil {
			results <- Result{
				xkcd: nil,
				err:  fmt.Errorf("Decode %d: %v\n", number, err),
			}
		} else {
			results <- Result{xkcd: &xkcd, err: nil}
		}
	case http.StatusNotFound:
		if number != 404 {
			results <- Result{xkcd: nil, err: errNotFound}
		}
	default:
		results <- Result{
			xkcd: nil,
			err:  fmt.Errorf("GET %s: %s", url, resp.Status),
		}
	}
}

func persist(index *ex12.Index, file string) error {
	if _, err := os.Stat(file); err == nil {
		if err = os.Remove(file); err != nil {
			return fmt.Errorf("delete %s: %v", file, err)
		}
	}
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("unable to create file %s: %v", file, err)
	}
	encoder := json.NewEncoder(bufio.NewWriter(f))
	encoder.SetIndent("", "  ")
	return encoder.Encode(index)
}
