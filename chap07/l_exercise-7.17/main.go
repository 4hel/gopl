// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack [][]string // stack of slices which contain name and attributes
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
			elem := []string{tok.Name.Local}
			for _, attr := range tok.Attr {
				elem = append(elem, attr.Value)
			}
			stack = append(stack, elem) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(flatten(stack), " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x [][]string, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if matches(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func matches(stackElement []string, param string) bool {
	for _, elem := range stackElement {
		if elem == param {
			return true
		}
	}
	return false
}

func flatten(stack [][]string) []string {
	retval := make([]string, 0)
	for _, elem := range stack {
		var token string
		if len(elem) > 1 {
			token = elem[0] + "(" + strings.Join(elem[1:], " ") + ")"
		} else {
			token = elem[0]
		}
		retval = append(retval, token)
	}
	return retval
}

//!-
