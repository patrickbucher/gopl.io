package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
func ls(dir string) []string {
	var items []string
	info, err := os.Stat(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "stat %s: %v\n", dir, err)
		return items
	}
	if !info.IsDir() {
		return items
	}
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "list %s: %v\n", dir, err)
		return items
	}
	for _, info := range infos {
		path := filepath.Join(dir, info.Name())
		fmt.Println(path)
		if info.IsDir() {
			items = append(items, path)
		}
	}
	return items
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [folder]\n", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "stat %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "%s is not a folder\n", os.Args[1])
		os.Exit(1)
	}
	breadthFirst(ls, []string{os.Args[1]})
}
