// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func (s *IntSet) Elems() []uint64 {
	var retval []uint64
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				retval = append(retval, uint64(i*64+j))
			}
		}
	}
	return retval
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(nums ...int) {
	for _, x := range nums {
		s.Add(x)
	}
}

func (s *IntSet) Len() (length int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	// fmt.Println("word:", word)
	// fmt.Println("bit:", bit)
	mask := (^uint64(0)) ^ (uint64(1) << bit)
	// fmt.Printf("%b\n", s.words[0])
	s.words[word] &= mask
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	retval := new(IntSet)
	retval.words = make([]uint64, len(s.words))
	copy(retval.words, s.words)
	return retval
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	var max int
	if len(s.words) > len(t.words) {
		max = len(s.words)
	} else {
		max = len(t.words)
	}
	for i := 0; i < max; i++ {
		if i < len(t.words) && i < len(s.words) {
			s.words[i] ^= t.words[i]
		} else if i >= len(s.words) {
			s.words = append(s.words, t.words[i])
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
