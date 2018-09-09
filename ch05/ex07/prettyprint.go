package ex07

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

const indent = 4

func PrettyPrint(r io.Reader, w io.Writer) error {
	doc, err := html.Parse(r)
	if err != nil {
		return fmt.Errorf("parse HTML from %v: %v", r, err)
	}
	prettyPrint(doc, w)
	return nil
}

var depth int

func startElement(n *html.Node, w io.Writer) (open bool) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(w, "%*s<%s", depth*indent, "", n.Data)
		// TODO process attributes
		if n.FirstChild != nil {
			open = true
			fmt.Fprint(w, ">\n")
			depth++
		} else {
			open = false
			fmt.Fprint(w, "/>\n")
		}
	}
	// TODO process TextNodes
	return
}

func endElement(n *html.Node, w io.Writer) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(w, "%*s</%s>\n", depth*indent, "", n.Data)
	}
}

func prettyPrint(n *html.Node, w io.Writer) {
	remainedOpen := startElement(n, w)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		prettyPrint(c, w)
	}
	if remainedOpen {
		endElement(n, w)
	}
}
