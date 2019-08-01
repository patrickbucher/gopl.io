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
	_, err := dec.Token() // _: root
	if err != nil {
		fmt.Fprintf(os.Stderr, "no root node found: %v\n", err)
		os.Exit(1)
	}
	// rootElem := root.(*xml.StartElement)
	// rootNode := Element{Type: rootElem.Name, Attr: rootElem.Attr, Children: []Node{}}
	// var nodeStack []Node
	level := 0
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "iterating over tokens: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			fmt.Printf("%s", strings.Repeat("\t", level))
			fmt.Println("start", tok.Name.Local)
			level += 1
			// add to topmost nodeStack item as child
			// add current element to stack
		case xml.EndElement:
			level -= 1
			fmt.Printf("%s", strings.Repeat("\t", level))
			fmt.Println("end", tok.Name.Local)
			// pop topmost nodeStack item
		case xml.CharData:
			str := strings.TrimSpace(string(tok))
			if str != "" {
				fmt.Printf("%s", strings.Repeat("\t", level))
				fmt.Println(str)
			}
			// add to topmost nodeStack item as child
		}
	}
}
