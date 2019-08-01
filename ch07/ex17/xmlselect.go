// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type selector struct {
	name string
	attr []xml.Attr
}

type selectorStack []selector

func (s selectorStack) flatten() []string {
	var stack []string
	for _, sel := range s {
		stack = append(stack, sel.name)
		for _, attr := range sel.attr {
			stack = append(stack, fmt.Sprintf(`%s="%s"`, attr.Name.Local, attr.Value))
		}
	}
	return stack
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack selectorStack
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			sel := selector{name: tok.Name.Local, attr: tok.Attr}
			stack = append(stack, sel) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			criteria := stack.flatten()
			if containsAll(criteria, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(criteria, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
