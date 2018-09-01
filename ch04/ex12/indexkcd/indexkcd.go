package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopl.io/ch04/ex12"
)

const (
	urlTemplate = "https://xkcd.com/%d/info.0.json"
	indexFile   = "indexkcd.json"
)

func main() {
	finished := false
	var entries []ex12.XKCD
	for number := 2000; !finished; number++ {
		url := fmt.Sprintf(urlTemplate, number)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "GET %s: %v\n", url, err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		switch resp.StatusCode {
		case http.StatusOK:
			decoder := json.NewDecoder(resp.Body)
			xkcd := ex12.XKCD{}
			if err := decoder.Decode(&xkcd); err != nil {
				fmt.Fprintf(os.Stderr, "Decode %d: %v\n", number, err)
				os.Exit(1)
			}
			entries = append(entries, xkcd)
			fmt.Println(number, len(entries), xkcd)
		case http.StatusNotFound:
			finished = true
		default:
			fmt.Fprintf(os.Stderr, "GET %s: %s", url, resp.Status)
			os.Exit(1)
		}
	}
	index := ex12.Index{Entries: entries}
	persist(index, indexFile)
}

func persist(index ex12.Index, file string) error {
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
	encoder.Encode(index)
	return nil
}
