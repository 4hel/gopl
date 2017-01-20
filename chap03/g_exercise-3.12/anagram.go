// Anagram checks if two strings contain the same letters
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("please give two argument strings")
	} else {
		fmt.Println(
			anagram(os.Args[1], os.Args[2]))
	}
}

func anagram(str1 string, str2 string) bool {
	map1 := getMap(str1)
	map2 := getMap(str2)
	return reflect.DeepEqual(map1, map2)
}

func getMap(str string) map[rune]int {
	rmap := make(map[rune]int)
	rs := []rune(str)
	for _, r := range rs {
		rmap[r] += 1
	}
	return rmap
}
