package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		printElementNodeStart(n)
		depth++
	} else if n.Type == html.CommentNode {
		printCommentNodeStart(n)
		depth++
	} else if n.Type == html.TextNode {
		printTextNode(n)
		//depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		printElementNodeEnd(n)
	} else if n.Type == html.CommentNode {
		depth--
		printCommentNodeEnd(n)
	} else if n.Type == html.TextNode {
		//depth--
	}
}

func printElementNodeStart(n *html.Node) {
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

func printElementNodeEnd(n *html.Node) {
	if n.FirstChild != nil {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func printCommentNodeStart(n *html.Node) {
	fmt.Printf("<!--%s", n.Data)
}

func printCommentNodeEnd(n *html.Node) {
	fmt.Printf("-->\n")
}

func printTextNode(n *html.Node) {
	trimmed := strings.TrimSpace(n.Data)
	if trimmed != "" {
		fmt.Printf("%*s%s\n", depth*2, "", trimmed)
	}
}
