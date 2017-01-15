package main

import (
	"os"
	"strings"
)

func main() {
	Concat(os.Args)
	Join(os.Args)
}

func Join(slice [...]src) {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	s := strings.Join(src[1:], " ")
}

func Concat(slice [...]src) {
	var s, sep string
	for i := 1; i < len(src); i++ {
		s += sep + src[i]
		sep = " "
	}
	// fmt.Println(s)
}
