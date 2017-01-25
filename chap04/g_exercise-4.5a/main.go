// remove adjacent duplicates from a string slice

package main

import (
	"fmt"
)

func main() {
	s := []string{"one", "two", "two", "three", "four"}
	fmt.Println(removedup(s))
}

func removedup(strings []string) []string {

	last := ""
	offset := 0

	for i, s := range strings {
		if s == "" {
			continue
		}
		if last == s {
			offset++
		}
		fmt.Println(i - offset)
		strings[i-offset] = s
		last = s
	}

	return strings[:len(strings)-offset]
}
