package main

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
		issue := readIssue()
		json, err := json.Marshal(issue)
		if err != nil {
			log.Fatalf("marshaling %v: %v\n", issue, err)
		}
		client := &http.Client{}
		reader := bytes.NewReader(json)
		req, err := http.NewRequest("POST", issuesURL, reader)
		if err != nil {
			log.Fatalf("init POST to %s: %v\n", issuesURL, err)
		}
		req.Header.Add("Authorization", "token "+token())
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("POST failed: %v\n", err)
		}
		if resp.StatusCode != http.StatusCreated {
			log.Fatalf("POST: %s\n", resp.Status)
		}
	case "read":
		issueNumber, err := readIssueNumber()
		if err != nil {
			log.Fatal("error reading issue #: %v\n", err)
		}
		// TODO: generalize issue retrievement
		url := issuesURL + "/" + fmt.Sprintf("%d", issueNumber)
		fmt.Println(url)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("init GET to %s: %v\n", url, err)
		}
		req.Header.Add("Authorization", "token "+token())
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("GET failed: %v\n", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("GET: %s\n", resp.Status)
		}
		dec := json.NewDecoder(resp.Body)
		var issue IssueView
		err = dec.Decode(&issue)
		if err != nil {
			log.Fatalf("error decoding response: %v\n", err)
		}
		// TODO: end of generalize issue retrievement
		printIssue(issue)
	case "update":
		issueNumber, err := readIssueNumber()
		if err != nil {
			log.Fatal("error reading issue #: %v\n", err)
		}
		// TODO: retrieve issue, read in changed information (using editors), POST it
	case "lock":
		log.Fatal("not implemented yet")
		// TODO: read in issue number; lock it, if found
	case "unlock":
		log.Fatal("not implemented yet")
		// TODO: read in issue number; unlock it, if found
	}
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

func printIssue(issue IssueView) {
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

func readIssue() Issue {
	// TODO expect existing issue, use and modify if given
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("Title: ")
	title, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("reading title: %v\n", err)
	}
	fmt.Printf("Labels (%s): ", strings.Join(labels, ", "))
	labelInput, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("reading labels: %v\n", err)
	}
	labelList := strings.Fields(labelInput)
	if ok, wrongLabel := allContained(labelList, labels); !ok {
		log.Fatalf("illegal label %s\n", wrongLabel)
	}
	// TODO if existing issue given, write text to file and pass in
	body, err := textInput()
	if err != nil {
		log.Fatalf("text input failed: %v\n", err)
	}
	return Issue{Title: title, Body: body, Assignees: []string{assignee},
		Labels: labelList}
}

func token() string {
	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Fatalf("error reading token from %s: %v\n", tokenFile, err)
	}
	return strings.TrimSpace(string(data))
}

func textInput() (string, error) {
	// TODO expect file param, us it if != nil, otherwise create temp file
	file, err := ioutil.TempFile("", "githubissue")
	if err != nil {
		return "", fmt.Errorf("create temp file: %v", err)
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
