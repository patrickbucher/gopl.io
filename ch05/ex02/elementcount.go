// Elementcount processes HTML input and counts the occuring elements by type.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing HTML from stdin: %v\n", err)
		os.Exit(1)
	}
	count := make(map[string]int)
	countElements(doc, count)
	for k, v := range count {
		fmt.Printf("%4d\t%s\n", v, k)
	}
}

func countElements(n *html.Node, count map[string]int) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	if n.FirstChild != nil {
		countElements(n.FirstChild, count)
	}
	if n.NextSibling != nil {
		countElements(n.NextSibling, count)
	}
}
