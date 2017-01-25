// Rotate a slice

package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 2)
	fmt.Println(a)
	//fmt.Println(rotate(a[:], 2))
}

func rotate(s []int, x int) {
	length := len(s)
	tmp := make([]int, x)
	copy(tmp, s[:x])
	copy(s, s[x:])
	copy(s[length-x:], tmp)
}
