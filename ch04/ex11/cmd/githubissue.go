package main

// TODO: make sure Title is stored without '\n' at the end (trim it)
// TODO: update without input should keep the original content
// TODO: multi word labels must not be split as fields, but by commas

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch04/ex11"
)

const (
	assignee  = "patrickbucher"
	issuesURL = "https://api.github.com/repos/patrickbucher/gopl.io/issues"
	tokenFile = ".github_token"
)

var labels = []string{
	"bug",
	"duplicate",
	"enhancement",
	"good first issue",
	"help wanted",
	"invalid",
	"question",
	"wontfix",
}
var commands = []string{
	"create",
	"read",
	"update",
	"lock", // instead of delete
	"unlock",
}

type Issue struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignees []string `json:"assignees"`
	Labels    []string `json:"labels"`
}

type User struct {
	Name string `json:"login"`
}

type Label struct {
	Name string `json:"name"`
}

type IssueView struct {
	Title    string  `json:"title"`
	Body     string  `json:"body"`
	Assignee User    `json:"assignee"`
	Labels   []Label `json:"labels"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s [action]\n", os.Args[0])
	}
	if !contains(commands, os.Args[1]) {
		log.Fatalf("command %q not supported, must be one of: %v\n",
			os.Args[1], commands)
	}
	command := os.Args[1]
	switch command {
	case "create":
		issue, err := readIssue()
		if err != nil {
			log.Fatalf("read issue: %v\n", err)
		}
		err = postIssue(issue, issuesURL)
		if err != nil {
			log.Fatalf("creating issue %#v: %v\n", issue, err)
		}
	case "read":
		issueNumber, err := readIssueNumber()
		if err != nil {
			log.Fatalf("error reading issue #: %v\n", err)
		}
		if issue, err := fetchIssue(issueNumber); err != nil {
			log.Fatalf("error fetching issue: %v\n", err)
		} else {
			printIssue(issue)
		}
	case "update":
		issueNumber, err := readIssueNumber()
		if err != nil {
			log.Fatalf("error reading issue #: %v\n", err)
		}
		issueView, err := fetchIssue(issueNumber)
		if err != nil {
			log.Fatalf("error fetching issue for editing: %v\n", err)
		}
		issue, err := inputIssueUpdate(issueView)
		if err != nil {
			log.Fatalf("error on input issue update: %v\n", err)
		}
		err = postIssue(issue, issuesURL+"/"+fmt.Sprintf("%d", issueNumber))
		if err != nil {
			log.Fatalf("updating issue #%d as %#v: %v\n",
				issueNumber, issue, err)
		}
	case "lock":
		issueNumber, err := readIssueNumber()
		if err != nil {
			log.Fatalf("error reading issue #: %v\n", err)
		}
		url := issuesURL + "/" + fmt.Sprintf("%d", issueNumber) + "/lock"
		req, err := http.NewRequest("PUT", url, nil)
		req.Header.Add("Authorization", "token "+token())
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("PUT issue #%d: %v\n", issueNumber, err)
		}
		if resp.StatusCode != http.StatusNoContent {
			log.Fatalf("PUT issue #%d: %s\n", resp.Status)
		}
	case "unlock":
		log.Fatal("not implemented yet")
		// TODO: read in issue number; unlock it, if found
	}
}

func postIssue(issue *Issue, url string) error {
	json, err := json.Marshal(issue)
	if err != nil {
		return fmt.Errorf("marshaling %v: %v", issue, err)
	}
	client := &http.Client{}
	reader := bytes.NewReader(json)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return fmt.Errorf("init POST to %s: %v", issuesURL, err)
	}
	req.Header.Add("Authorization", "token "+token())
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("POST failed: %v", err)
	}
	if resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusOK {
		return fmt.Errorf("POST: %s", resp.Status)
	}
	return nil
}

func fetchIssue(issueNumber int) (*IssueView, error) {
	url := issuesURL + "/" + fmt.Sprintf("%d", issueNumber)
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("init GET to %s: %v\n", url, err)
	}
	req.Header.Add("Authorization", "token "+token())
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET failed: %v\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET: %s\n", resp.Status)
	}
	dec := json.NewDecoder(resp.Body)
	var issue IssueView
	err = dec.Decode(&issue)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v\n", err)
	}
	return &issue, nil
}

func readIssueNumber() (int, error) {
	fmt.Printf("Issue #: ")
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("reading issue #: %v\n", err)
	}
	issueNumber, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf("parsing %q to int: %v\n", input, err)
	}
	return issueNumber, nil
}

func printIssue(issue *IssueView) {
	fmt.Printf("Title:\t\t%s\n", issue.Title)
	fmt.Printf("Assignee:\t%s\n", issue.Assignee.Name)
	fmt.Printf("Labels:\t\t%s\n", fmtLabels(issue.Labels))
	fmt.Printf("Body:\n\n%s\n", issue.Body)
}

func fmtLabels(labels []Label) string {
	names := make([]string, len(labels))
	for i := range labels {
		names[i] = labels[i].Name
	}
	return strings.Join(names, ", ")
}

func extractLabelNames(labels []Label) []string {
	labelNames := make([]string, len(labels))
	for i := range labels {
		labelNames[i] = labels[i].Name
	}
	return labelNames
}

func inputIssueUpdate(existing *IssueView) (*Issue, error) {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("New Title [old: %q]: ", existing.Title)
	title, err := r.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("reading title: %v", err)
	}
	fmt.Printf("Labels (%s) [old: %s]: ", strings.Join(labels, ", "),
		strings.Join(extractLabelNames(existing.Labels), ", "))
	labelInput, err := r.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("reading labels: %v\n", err)
	}
	labelList := strings.Fields(labelInput)
	if ok, wrongLabel := allContained(labelList, labels); !ok {
		return nil, fmt.Errorf("illegal label %q\n", wrongLabel)
	}
	body, err := textInput(existing.Body)
	if err != nil {
		return nil, fmt.Errorf("updating body: %v", err)
	}
	return &Issue{
		Title:     title,
		Body:      body,
		Labels:    labelList,
		Assignees: []string{existing.Assignee.Name},
	}, nil
}

func readIssue() (*Issue, error) {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("Title: ")
	title, err := r.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("reading title: %v", err)
	}
	fmt.Printf("Labels (%s): ", strings.Join(labels, ", "))
	labelInput, err := r.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("reading labels: %v", err)
	}
	labelList := strings.Fields(labelInput)
	if ok, wrongLabel := allContained(labelList, labels); !ok {
		return nil, fmt.Errorf("illegal label %s", wrongLabel)
	}
	body, err := textInput("")
	if err != nil {
		return nil, fmt.Errorf("text input failed: %v", err)
	}
	return &Issue{Title: title, Body: body, Assignees: []string{assignee},
		Labels: labelList}, nil
}

func token() string {
	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Fatalf("error reading token from %s: %v\n", tokenFile, err)
	}
	return strings.TrimSpace(string(data))
}

func textInput(content string) (string, error) {
	file, err := ioutil.TempFile("", "githubissue")
	if err != nil {
		return "", fmt.Errorf("create temp file: %v", err)
	}
	if len(content) > 0 {
		w := bufio.NewWriter(file)
		w.WriteString(content)
		if err := w.Flush(); err != nil {
			return "", fmt.Errorf("unable to flush %s: %v", file.Name(), err)
		}
	}
	if err = ex11.Edit(file.Name()); err != nil {
		return "", fmt.Errorf("edit file %q: %v\n", file.Name(), err)
	}
	if data, err := ioutil.ReadFile(file.Name()); err != nil {
		return "", fmt.Errorf("read file %q: %v\n", file.Name(), err)
	} else {
		return string(data), nil
	}
}

func allContained(part, whole []string) (bool, string) {
	for _, s := range part {
		if !contains(whole, s) {
			return false, s
		}
	}
	return true, ""
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
