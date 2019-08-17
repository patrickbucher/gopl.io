package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string, cancel <-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// NOTE: this test is fleaky, for the timing is only an educated guess.
func TestCancel(t *testing.T) {
	var tests = []struct {
		url   string
		await time.Duration
	}{
		{"https://news.ycombinator.com", time.Millisecond * 100},
		{"https://golang.org", time.Second * 2},
	}
	m := New(httpGetBody)
	for _, test := range tests {
		start := time.Now()
		cancel := make(chan struct{}, 1) // non-blocking
		finished := make(chan struct{})
		go func() {
			_, err := m.Get(test.url, cancel)
			if err != nil {
				t.Log(test.url, err)
			} else {
				t.Log(test.url, "retrieved")
			}
			finished <- struct{}{}
		}()
		time.Sleep(test.await)
		cancel <- struct{}{}
		t.Log(time.Since(start))
		<-finished
	}
	m.Close()
}

func incomingURLs() []string {
	testURLs := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://gopl.io",
	}
	// add every URL twice
	for _, url := range testURLs {
		testURLs = append(testURLs, url)
	}
	return testURLs
}

func TestMemo(t *testing.T) {
	m := New(httpGetBody)
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			t.Error(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	m.Close()
}

func TestMemoAsync(t *testing.T) {
	m := New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
	m.Close()
}
