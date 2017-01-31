// parse html TextNodes and print to stdout
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"unicode"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode && n.Parent.Data != "script" {
		if containsNonWhiteSpace([]rune(n.Data)) {
			fmt.Print(n.Data, " ")
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}

func containsNonWhiteSpace(s []rune) bool {
	retval := false
	for _, c := range s {
		if !unicode.IsSpace(c) {
			retval = true
		}
	}
	return retval
}
