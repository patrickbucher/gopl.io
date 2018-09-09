package ex07

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

var testURLs = []string{
	"http://paedubucher.ch",
	"http://marmaro.de",
	"http://blog.fefe.de",
	"http://gaegs.ch",
}

func TestPrettyPrint(t *testing.T) {
	for _, url := range testURLs {
		body, err := fetch(url)
		if err != nil {
			t.Logf("error fetching %s: %v", url, err)
			continue
		}
		defer body.Close()
		var buf bytes.Buffer
		w := bufio.NewWriter(&buf)
		PrettyPrint(body, w)
		r := bufio.NewReader(&buf)
		if _, err := html.Parse(r); err != nil {
			t.Fatalf("PrettyPrint %s: %v", url, err)
		}
	}
}

func fetch(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %v", url, resp.Status)
	}
	return resp.Body, nil
}
