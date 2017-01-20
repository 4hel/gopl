// Comma inserts commas into positive integers
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		var prefix string
		var result string = arg
		var suffix string

		if arg[0] == '-' {
			prefix = "-"
			result = arg[1:]
		}

		if i := strings.LastIndex(result, "."); i >= 0 {
			fmt.Println("lastindex = " + strconv.Itoa(i))
			suffix = result[i:]
			result = result[:i]
		}

		fmt.Println(prefix + comma(result) + suffix)
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	} else {
		return comma(s[:n-3]) + "," + s[n-3:]
	}
}
