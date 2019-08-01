package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var nodeStack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "iterating over tokens: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			var children []Node
			current := &Element{tok.Name, tok.Attr, children}
			if len(nodeStack) > 0 {
				var parentNode *Element
				parentNode = nodeStack[len(nodeStack)-1]
				parentNode.Children = append(parentNode.Children, current)
			}
			nodeStack = append(nodeStack, current)
		case xml.EndElement:
			if len(nodeStack) > 1 { // don't pop root element
				nodeStack = nodeStack[:len(nodeStack)-1]
			}
		case xml.CharData:
			str := strings.TrimSpace(string(tok))
			if str != "" && len(nodeStack) > 0 {
				var parentNode *Element
				parentNode = nodeStack[len(nodeStack)-1]
				parentNode.Children = append(parentNode.Children, CharData(str))
			}
		}
	}
	root := nodeStack[0]
	printTree(root)
}

var level = 0

func printTree(node Node) {
	indent := strings.Repeat("\t", level)
	switch n := node.(type) {
	case *Element:
		fmt.Printf("%s<%s %v>\n", indent, n.Type.Local, attrToStr(n.Attr))
		level += 1
		for _, child := range n.Children {
			printTree(child)
		}
		level -= 1
		fmt.Printf("%s</%s>\n", indent, n.Type.Local)
	case CharData:
		fmt.Printf("%s%s\n", indent, n)
	}
}

func attrToStr(attr []xml.Attr) string {
	var list []string
	for _, a := range attr {
		list = append(list, fmt.Sprintf(`%s="%s"`, a.Name.Local, a.Value))
	}
	return strings.Join(list, " ")
}
