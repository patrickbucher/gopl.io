package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	issuesURL             = "https://api.github.com/search/issues"
	htmlDir               = "html"
	issueListFile         = "issues.html"
	issueFileTemplate     = "issue-%d.html"
	userFileTemplate      = "user-%d.html"
	milestoneFileTemplate = "milestone-%d.html"
)

type SearchResult struct {
	Count  int     `json:"total_count"`
	Issues []Issue `json:"items"`
}

type Issue struct {
	Id            int       `json:"id"`
	Number        int       `json:"number"`
	User          User      `json:"user"`
	RepositoryURL string    `json:"repository_url"`
	Labels        []Label   `json:"labels"`
	Milestone     Milestone `json:"milestone"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
	ClosedAt      string    `json:"closed_at"`
	Body          string    `json:"body"`
}

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"login"`
	ProfileURL string `json:"url"`
}

type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Milestone struct {
	Id          int    `json:"id"`
	Number      int    `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TODO issue list template
// TODO single issue template
// TODO user template
// TODO milestone template

func main() {
	if len(os.Args) < 2 {
		fail("usage: %s [search params]\n", os.Args[0])
	}
	issues, err := searchIssues(os.Args[1:])
	if err != nil {
		fail("retrieve GitHub issues: %v\n", err)
	}
	fmt.Println(issues)
	// TODO process issue list, issues, users, milestones to html files
	// TODO offer html files
}

func searchIssues(searchTerms []string) (*SearchResult, error) {
	var searchResult SearchResult
	q := url.QueryEscape(strings.Join(searchTerms, " "))
	searchURL := fmt.Sprintf("%s?q=%s", issuesURL, q)
	resp, err := http.Get(searchURL)
	if err != nil {
		fail("GET %s: %v\n", searchURL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fail("GET %s: %s\n", searchURL, resp.Status)
	}
	decoder := json.NewDecoder(bufio.NewReader(resp.Body))
	if err := decoder.Decode(&searchResult); err != nil {
		return &searchResult, fmt.Errorf("unmarshal JSON: %v", err)
	}
	return &searchResult, nil
}

func fail(format string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, format, params)
	os.Exit(1)
}
