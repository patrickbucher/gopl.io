package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// NOTE: this test is fleaky, for the timing is only an educated guess.
func TestCancel(t *testing.T) {
	url := "https://news.ycombinator.com" // approx. 600ms
	m := New(httpGetBody)
	start := time.Now()
	cancel := make(chan struct{})
	go func() {
		_, err := m.Get(url, cancel)
		if err != nil {
			t.Log(err)
		}
	}()
	time.Sleep(600 * time.Millisecond)
	cancel <- struct{}{}
	t.Log(time.Since(start))
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
