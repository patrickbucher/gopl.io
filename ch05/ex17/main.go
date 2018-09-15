package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [html element names]\n", os.Args[0])
		os.Exit(1)
	}
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing HTML from Stdin: %v\n", err)
		os.Exit(1)
	}
	elements := ElementsByTagName(doc, os.Args[1:]...)
	for _, e := range elements {
		var attributes []string
		for _, a := range e.Attr {
			attr := fmt.Sprintf(`%s="%s"`, a.Key, a.Val)
			attributes = append(attributes, attr)
		}
		if len(attributes) > 0 {
			args := strings.Join(attributes, " ")
			fmt.Printf("<%s %s>\n", e.Data, args)
		} else {
			fmt.Printf("<%s>\n", e.Data)
		}
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var elements []*html.Node
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		elements = append(elements, visitAll(c, name...)...)
	}
	return elements
}

func visitAll(node *html.Node, name ...string) []*html.Node {
	var found []*html.Node
	if node.Type == html.ElementNode {
		for _, n := range name {
			if node.Data == n {
				found = append(found, node)
				break
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			found = append(found, visitAll(c, name...)...)
		}
	}
	return found
}
