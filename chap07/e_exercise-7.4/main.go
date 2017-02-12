package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type MyReader string

func (r *MyReader) Read(p []byte) (n int, err error) {
	n = len(*r)
	copy(p, []byte(*r))
	err = io.EOF
	return
}

func NewReader(s string) io.Reader {
	var r MyReader
	r = MyReader(s)
	return &r
}

func main() {
	doc, _ := html.Parse(NewReader("<html><body><h1>hello</h1></body></html>"))
	fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
}
