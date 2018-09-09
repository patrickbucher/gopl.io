package ex07

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

const indent = 2

func PrettyPrint(r io.Reader, w io.Writer) error {
	doc, err := html.Parse(r)
	if err != nil {
		return fmt.Errorf("parse HTML from %v: %v", r, err)
	}
	prettyPrint(doc, w)
	return nil
}

var depth int
var onNewLine bool = true

func startElement(n *html.Node, w io.Writer) (open bool) {
	switch n.Type {
	case html.ElementNode:
		if isBlockElement(n) && !onNewLine {
			fmt.Fprint(w, "\n")
			onNewLine = true
		}
		if onNewLine {
			fmt.Fprintf(w, "%*s<%s", depth*indent, "", n.Data)
		} else {
			fmt.Fprintf(w, "<%s", n.Data)
		}
		onNewLine = false
		if attributes, has := formatAttributes(n); has {
			fmt.Fprintf(w, " %s", attributes)
		}
		if n.FirstChild != nil {
			fmt.Fprint(w, ">")
			open = true
			onNewLine = false
			depth++
		} else {
			fmt.Fprint(w, "/>\n")
			open = false
			onNewLine = true
		}
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		visibleSpacing := (text == "" && n.Data != "")
		if visibleSpacing {
			if !onNewLine {
				fmt.Fprint(w, " ")
			}
		} else {
			if onNewLine {
				fmt.Fprintf(w, "%*s%s", depth*indent, "", n.Data)
				onNewLine = false
			} else {
				fmt.Fprintf(w, "%s", n.Data)
			}
		}
	case html.CommentNode:
		if onNewLine {
			fmt.Fprintf(w, "%*s<!-- %s -->\n", depth*indent, "", n.Data)
		} else {
			fmt.Fprintf(w, "<!-- %s -->", n.Data)
		}
	case html.DocumentNode:
		// the output will be proper HTML, no matter what the input was
		fmt.Fprintln(w, "<!DOCTYPE html>")
	}
	return
}

func formatAttributes(n *html.Node) (string, bool) {
	if len(n.Attr) == 0 {
		return "", false
	}
	var attributes []string
	for _, a := range n.Attr {
		attr := fmt.Sprintf(`%s="%s"`, a.Key, a.Val)
		attributes = append(attributes, attr)
	}
	return strings.Join(attributes, " "), true
}

func endElement(n *html.Node, w io.Writer) {
	if n.Type == html.ElementNode {
		depth--
		if onNewLine {
			fmt.Fprintf(w, "%*s</%s>", depth*indent, "", n.Data)
			onNewLine = false
		} else {
			fmt.Fprintf(w, "</%s>", n.Data)
		}
		if isBlockElement(n) {
			fmt.Fprintf(w, "\n")
			onNewLine = true
		}
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

var blockElementNames = []string{
	"html", "head", "meta", "body", "link", "title",
	"h1", "h2", "h3", "p", "ol", "ul", "li", "blockquote", "div",
}

func isBlockElement(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	return contains(blockElementNames, n.Data)
}

func contains(haystack []string, needle string) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}
