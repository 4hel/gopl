package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	t.Log("testing PopCount function...")

	// 0x01 -> 1 bit set
	result := PopCount(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	// 0x05 -> 2 bits set
	result = PopCount(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0x0500 but got %d", result)
	} else {
		t.Log("2 bits set for 0x0500")
	}

}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
