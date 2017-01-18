// using right shift
package main

import (
	"fmt"
)

func main() {
	var v uint32 = 0x05000500 // b 00000101 00000000 00000101 00000000
	fmt.Println("before:             ", word2string(v))
	fmt.Println("then right shift 8: ", word2string(v>>8))
	fmt.Printf("then cast to byte:   %24s%08b\n", " ", byte(v>>8))

}

func word2string(b uint32) string {
	s := fmt.Sprintf("%b", b)
	for i := len(s); i < 32; i++ {
		s = "0" + s
	}
	return s
}
