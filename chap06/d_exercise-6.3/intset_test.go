package intset

import (
	"fmt"
)

func ExampleIntersect() {
	var x, y IntSet
	x.AddAll(1, 144)
	y.AddAll(5, 144, 20, 1024)
	x.IntersectWith(&y)
	fmt.Println(x.String()) // "{144}"
	// Output:
	// {144}
}

func ExampleDifference() {
	var x, y IntSet
	x.AddAll(1, 144)
	y.AddAll(5, 144, 1024)
	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{1 5 1024}"
	// Output:
	// {1 5 1024}
}
