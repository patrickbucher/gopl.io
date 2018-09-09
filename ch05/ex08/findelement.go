package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [id]", os.Args[0])
		os.Exit(1)
	}
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		os.Exit(1)
	}
	n := ElementByID(doc, os.Args[1])
	if n != nil {
		var attrList []string
		for _, a := range n.Attr {
			attrList = append(attrList, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
		}
		attributes := strings.Join(attrList, " ")
		fmt.Printf("<%s %s>\n", n.Data, attributes)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var element *html.Node
	find := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					element = n
					return false
				}
			}
		}
		return true
	}
	forEachNode(doc, find, nil)
	return element
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && !pre(n) {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil && !post(n) {
		return
	}
}
