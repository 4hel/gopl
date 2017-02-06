package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	node := getElementById(doc, "id1")
	if node != nil {
		fmt.Println(*node)
	}
}

func getElementById(n *html.Node, id string) *html.Node {
	var retval *html.Node = nil

	for _, a := range n.Attr {
		if a.Key == "id" {
			if a.Val == id {
				retval = n
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if retval == nil {
			retval = getElementById(c, id)
		}
	}

	return retval
}
