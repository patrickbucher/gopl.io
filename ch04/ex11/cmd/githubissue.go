package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopl.io/ch04/ex11"
)

const (
	assignee = "patrickbucher"
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
		r := bufio.NewReader(os.Stdin)
		fmt.Printf("Title: ")
		title, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("reading title: %v\n", err)
		}
		fmt.Printf("Labels (%v): ", labels)
		labelInput, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("reading labels: %v\n", err)
		}
		labelList := strings.Fields(labelInput)
		if ok, wrongLabel := allContained(labelList, labels); !ok {
			log.Fatalf("illegal label %s\n", wrongLabel)
		}
		body, err := textInput()
		if err != nil {
			log.Fatalf("text input failed: %v\n", err)
		}
		issue := Issue{Title: title, Body: body, Assignees: []string{assignee},
			Labels: labelList}
		fmt.Println(issue)
	case "read":
		log.Fatal("not implemented yet")
	case "update":
		log.Fatal("not implemented yet")
	case "lock":
		log.Fatal("not implemented yet")
	case "unlock":
		log.Fatal("not implemented yet")
	}
}

func textInput() (string, error) {
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
