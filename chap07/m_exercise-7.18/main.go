package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{} // *Element or CharData

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Node
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			//fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := newElement(tok)
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				if parent, ok := parent.(*Element); ok {
					parent.Children = append(parent.Children, elem)
					//fmt.Println("child appended to ", parent)
				}
			}
			stack = append(stack, elem)
			//fmt.Println("stack appended: ", elem.Type.Local)
		case xml.EndElement:
			if len(stack) == 1 {
				root = stack[0]
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			text := strings.TrimSpace(string(tok))
			if len(text) > 0 {
				if len(stack) > 0 {
					parent := stack[len(stack)-1]
					if parent, ok := parent.(*Element); ok {
						parent.Children = append(parent.Children, CharData(text))
					}
				}
			}
		}
	}

	visit(root)
}

func newElement(src xml.StartElement) *Element {
	var retval Element
	retval.Type = src.Name
	for _, attr := range src.Attr {
		retval.Attr = append(retval.Attr, attr)
	}
	return &retval
}

func visit(n Node) {
	switch n := n.(type) {
	case CharData:
		fmt.Println("visiting: ", n)
	case *Element:
		fmt.Println("visiting: " + n.Type.Local)
		for _, child := range n.Children {
			visit(child)
		}
	}
}
