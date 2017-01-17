// return pointer to local variable is possible
// go performs pointers escape analysis
// so it knows it has to allocate v on the heap
package main

import (
	"fmt"
)

func main() {
	var p = f()
	fmt.Println(*p)
}

func f() *int {
	v := 1
	return &v
}
