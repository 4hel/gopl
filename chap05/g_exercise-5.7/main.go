package main

import (
	"fmt"

	"os"

	"golang.org/x/net/html"
)

func main() {
	/*
		for _, url := range os.Args[1:] {
			outline(url)
		}
	*/
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

/*
func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, nil)
	//!-call

	return nil
}
*/

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		//fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		printNodeStart(n)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func printNodeStart(n *html.Node) {
	fmt.Printf("%*s<%s", depth*2, "", n.Data)
	for _, a := range n.Attr {
		fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
	}
	if n.FirstChild != nil {
		fmt.Print(">\n")
	} else {
		fmt.Print("/>\n")
	}
}

//!-startend
