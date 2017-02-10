package intset

import (
	"fmt"
)

func ExampleElems() {
	var x IntSet
	x.AddAll(1, 5, 144)
	for _, e := range x.Elems() {
		fmt.Print(e)
	}
	// Output:
	// 15144
}
