// counting the number of bits set to 1 in a byte
package main

import (
	"fmt"
)

var counts [16]byte

func main() {
	for i := range counts {
		//
		// in i&1 the bitwise AND is used to mask the least significant bit
		// so the byte toggles between 0 and 1 for even / odd i
		//
		counts[i] = counts[i/2] + byte(i&1)
		str_i := byte2string(byte(i))
		fmt.Printf("i = %02d : %s -> %d bits set to 1 (%d + %d) \n", i, str_i, counts[i], counts[i/2], i&1)
	}
}

func byte2string(b byte) string {
	s := fmt.Sprintf("%b", b)
	for i := len(s); i < 8; i++ {
		s = "0" + s
	}
	return s
}
