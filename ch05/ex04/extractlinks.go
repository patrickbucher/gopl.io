package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing HTML from stdiné %v\n", err)
		os.Exit(1)
	}
	var links []string
	extractLinks(doc, &links)
	for _, v := range links {
		fmt.Println(v)
	}
}

func extractLinks(n *html.Node, links *[]string) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "link":
			fallthrough
		case "a":
			if val, ok := extractAttrVal(n, "href"); ok {
				*links = append(*links, val)
			}
		case "script":
			fallthrough
		case "img":
			if val, ok := extractAttrVal(n, "src"); ok {
				*links = append(*links, val)
			}
		}
	}
	if n.FirstChild != nil {
		extractLinks(n.FirstChild, links)
	}
	if n.NextSibling != nil {
		extractLinks(n.NextSibling, links)
	}
}

func extractAttrVal(n *html.Node, attrKey string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == attrKey {
			return attr.Val, true
		}
	}
	return "", false
}
