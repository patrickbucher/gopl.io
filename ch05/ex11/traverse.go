package main

import "fmt"

func main() {
	graph := map[string][]string{
		"A": {"B", "D"},
		"B": {"C", "D"},
		"C": {"A"},
		"D": {},
	}
	output := func(node string, list []string) {
		fmt.Printf("%q: %q\n", node, list)
	}
	traverse(graph, output)
}

func traverse(graph map[string][]string, forEach func(string, []string)) {
	for node := range graph {
		visited := map[string]bool{}
		visitAll(graph, node, visited, forEach)
	}
}

func visitAll(g map[string][]string, n string, v map[string]bool, f func(string, []string)) {
	if !v[n] {
		f(n, g[n])
		v[n] = true
	}
	for _, e := range g[n] {
		if !v[e] {
			v[e] = true
			visitAll(g, e, v, f)
		}
	}
}
