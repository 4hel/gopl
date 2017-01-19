package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	t.Log("testing PopCount function...")

	result := PopCount(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	result = PopCount(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0xffef but got %d", result)
	} else {
		t.Log("15 bits set for 0xffef")
	}

}

func TestPopCountLoop(t *testing.T) {
	t.Log("testing PopCountLoop function...")

	result := PopCountLoop(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	result = PopCountLoop(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0xffef but got %d", result)
	} else {
		t.Log("15 bits set for 0xffef")
	}
}

func TestPopCountShift1(t *testing.T) {
	t.Log("testing PopCountShift1 function...")

	result := PopCountShift1(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	result = PopCountShift1(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0xffef but got %d", result)
	} else {
		t.Log("15 bits set for 0xffef")
	}
}

func TestPopCountClear1(t *testing.T) {
	t.Log("testing PopCountShift1 function...")

	result := PopCountClear1(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	result = PopCountClear1(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0xffef but got %d", result)
	} else {
		t.Log("15 bits set for 0xffef")
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(uint64(i))
	}
}

func BenchmarkPopCountShift1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift1(uint64(i))
	}
}

func BenchmarkPopCountClear1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear1(uint64(i))
	}
}
