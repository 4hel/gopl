// Couting the number of bits set to 1 in number (population)
package popcount

// pc[i] is the population for the number i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount return the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Exercise 2.3
func PopCountLoop(x uint64) int {
	var retval byte
	for i := uint64(0); i < 8; i++ {
		retval += pc[byte(x>>(i*8))]
	}
	return int(retval)
}
