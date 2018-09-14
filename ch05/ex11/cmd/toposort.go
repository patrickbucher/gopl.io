package main

import (
	"fmt"
	"os"
	"sort"

	"gopl.io/ch05/ex11"
)

// prereqs maps computer science courses to their prerequisites.
var prereps = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},

	// this creates a cycle
	"linear algebra": {"calculus"},
}

func main() {
	if ex11.IsCyclic(prereps) {
		fmt.Fprintln(os.Stderr, "topological sorting impossible for cyclic graph")
		os.Exit(1)
	}
	for i, course := range topoSort(prereps) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// TODO: incorporate logic of ex11.IsCyclic() here
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
