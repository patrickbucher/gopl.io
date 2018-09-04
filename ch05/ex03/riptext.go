// Riptext reads HTML and outputs the content of all its text nodes.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse HTML from stdin: %v\n", err)
		os.Exit(1)
	}
	ripText(doc)
}

func ripText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}
	if n.Type != html.ElementNode || n.Data != "script" {
		if n.FirstChild != nil {
			ripText(n.FirstChild)
		}
		if n.NextSibling != nil {
			ripText(n.NextSibling)
		}
	}
}
