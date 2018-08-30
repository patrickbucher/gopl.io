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
		fmt.Printf("Issue #: ")
		r := bufio.NewReader(os.Stdin)
		input, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("reading issue #: %v\n", err)
		}
		issueNumber, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatalf("parsing %q to int: %v\n", input, err)
		}
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
		var issue Issue // different structure for retrieval needed (see labels)
		err = dec.Decode(&issue)
		if err != nil {
			log.Fatalf("error decoding response: %v\n", err)
		}
		fmt.Println(issue)
	case "update":
		log.Fatal("not implemented yet")
	case "lock":
		log.Fatal("not implemented yet")
	case "unlock":
		log.Fatal("not implemented yet")
	}
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
