package outline

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func NewReader(s string) io.Reader {
	return strings.NewReader(s)
}

func Parse(input string) ([]string, error) {
	stack := make([]string, 0)
	doc, err := html.Parse(NewReader(input))
	if err != nil {
		return stack, fmt.Errorf("error parsing HTML: %v", err)
	}
	stack = outline(stack, doc)
	return stack, nil
}

func outline(stack []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		stack = outline(stack, c)
	}
	return stack
}
