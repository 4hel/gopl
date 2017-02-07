package intset

import (
	"fmt"
)

func ExampleUnion() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println("x:", x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println("y:", y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println("union:", x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	// Output:
	// x: {1 9 144}
	// y: {9 42}
	// union: {1 9 42 144}
	// true false
}

func ExampleLen() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.Len()) // "3"
	// Output:
	// 3
}

func ExampleRemove() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Remove(1)
	fmt.Println(x.String()) // "{9 144}"
	// Output:
	// {9 144}
}

func ExampleClear() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Clear()
	fmt.Println(x.String()) // "{}"
	// Output:
	// {}
}

func ExampleCopy() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y := x.Copy()
	fmt.Println(y.String()) // "{1 9 144}"
	// Output:
	// {1 9 144}
}
