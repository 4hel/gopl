// Counting the number of different bits in two sha256 sums

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x")) // c1 is of type [32]byte
	c2 := sha256.Sum256([]byte("X")) // c2 is of type [32]byte
	fmt.Println(countDiff(c1, c2))
}

func countDiff(c1, c2 [32]byte) int {
	count := 0
	for i, _ := range c1 {
		count += countDiffByte(c1[i], c2[i])
	}
	return count
}

func countDiffByte(b1, b2 byte) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		bit1 := (b1 >> i) & 1
		bit2 := (b2 >> i) & 1
		if bit1 != bit2 {
			count++
		}
	}
	return count
}
