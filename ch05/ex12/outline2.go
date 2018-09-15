package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, 0)
}

func forEachNode(n *html.Node, depth int) {
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
	}
	startElement(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, depth)
	}
	endElement(n)
}
