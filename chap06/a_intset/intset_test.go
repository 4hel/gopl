package intset

import (
	"fmt"
	"testing"
)

func TestUnion(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println("x: ", x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println("y: ", y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println("union: ", x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
}
