package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopl.io/ch04/ex14"
)

const (
	issuesURL             = "https://api.github.com/search/issues"
	htmlDir               = "html/"
	issueListFile         = "index.html"
	issueFileTemplate     = "issue-%d.html"
	userFileTemplate      = "user-%d.html"
	milestoneFileTemplate = "milestone-%d.html"
	pathSep               = string(os.PathSeparator)
)

func main() {
	if len(os.Args) < 2 {
		fail("usage: %s [search params]\n", os.Args[0])
	}
	issues, err := searchIssues(os.Args[1:])
	if err != nil {
		fail("retrieve GitHub issues: %v\n", err)
	}
	if err := os.RemoveAll(htmlDir); err != nil {
		fail("error removing dir %s: %v\n", htmlDir, err)
	}
	if err := os.Mkdir(htmlDir, 0751); err != nil {
		fail("error creating dir %s: %v\n", htmlDir, err)
	}
	if err = createIssueList(issues); err != nil {
		fail("error creating issue list: %v\n", err)
	}
	if err = createIssuePages(issues); err != nil {
		fail("error creating issue pages: %v\n", err)
	}
	// createUserPages(issues)
	// createMilestonePages(issues)
	http.Handle("/", http.FileServer(http.Dir(htmlDir)))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func createIssuePages(result *ex14.SearchResult) error {
	for _, issue := range result.Issues {
		issuePageFileName := fmt.Sprintf(issueFileTemplate, issue.Id)
		issuePageFilePath := htmlDir + pathSep + issuePageFileName
		issuePageHtml, err := os.Create(issuePageFilePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %v",
				issuePageFilePath, err)
		}
		ex14.IssuePageTemplate.Execute(issuePageHtml, issue)
	}
	return nil
}

func createIssueList(result *ex14.SearchResult) error {
	issueListFileName := htmlDir + pathSep + issueListFile
	issueListHtml, err := os.Create(issueListFileName)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", issueListFileName, err)
	}
	ex14.IssueListTemplate.Execute(issueListHtml, result)
	return nil
}

func searchIssues(searchTerms []string) (*ex14.SearchResult, error) {
	var searchResult ex14.SearchResult
	query := strings.Join(searchTerms, " ")
	escapedQuery := url.QueryEscape(query)
	searchURL := fmt.Sprintf("%s?q=%s", issuesURL, escapedQuery)
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
	searchResult.SearchTerms = query
	return &searchResult, nil
}

func fail(format string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, format, params)
	os.Exit(1)
}
