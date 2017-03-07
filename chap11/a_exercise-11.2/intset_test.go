package intset

import "testing"

func TestAddHas(t *testing.T) {
	var iset IntSet
	addList := []int{2, 8, 128, 456, 1099}
	hasTable := []struct {
		n   int
		has bool
	}{
		{1, false},
		{2, true},
		{3, false},
		{8, true},
		{128, true},
		{129, false},
		{456, true},
		{1099, true},
		{2000, false},
	}
	for _, n := range addList {
		iset.Add(n)
	}
	for _, check := range hasTable {
		if iset.Has(check.n) != check.has {
			t.Errorf("iset.Has(%d) != %v\n", check.n, check.has)
		}
	}
}
