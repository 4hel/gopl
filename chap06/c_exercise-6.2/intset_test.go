package intset

import (
	"fmt"
)

func ExampleAddAll() {
	var x IntSet
	x.AddAll(1, 144, 9)
	fmt.Println(x.String()) // "{1 9 144}"
	// Output:
	// {1 9 144}
}
