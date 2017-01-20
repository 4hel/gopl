// Comma inserts commas into positive integers
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n < 4 {
		return s
	}
	offset := n % 3
	if offset > 0 {
		buf.WriteString(s[:offset])
		buf.WriteByte(',')
	}
	s = s[offset:]
	steps := len(s) / 3
	for i := 0; i < steps; i++ {
		pos := 3 * i
		buf.WriteString(s[pos : pos+3])
		if i < steps-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
