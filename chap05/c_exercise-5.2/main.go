// count types of element nodes in map
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var m map[string]int = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
	for k, v := range m {
		fmt.Printf("%8s -> %d\n", k, v)
	}
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
