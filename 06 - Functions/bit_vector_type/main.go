package main

import (
	"bytes"
	"fmt"
)

/*
---- IntSet ----
- Bit vector implementation of set
- using `uint64` so we have 64 bits in each bucket (index in arr)
- 64 bits = 64 integers(positive) to represent

*/

type IntSet struct {
	words  []uint64
	length int
}

// Add non-negative val x to set
func (s *IntSet) Add(x int) {
	// word - bucket in arr of words
	// bit - which bit to modify
	word, bit := x/64, uint(x%64)

	// Curr bucket doesnt exist
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// Modify bit
	s.words[word] |= 1 << bit
	s.length++
}

func (s *IntSet) AddAll(x ...int) {
	for _, i := range x {
		s.Add(i)
	}
}

// return number of elements
func (s *IntSet) Len() int {
	return s.length
}

// remove x from set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if !s.Has(x) {
		fmt.Println("You high or something? the elem doesnt exist...")
	} else {
		s.words[word] &^= 1 << bit
		s.length--
	}
}

// Remove all elem from set
func (s *IntSet) Clear() {
	s.words = nil
	s.length = 0
}

// Check whether non-negative val x exists
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	// bucket exists AND bit is not zero
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// set `s` to union of `s` and `t`
func (s *IntSet) UnionsWith(t *IntSet) {
	for i, tword := range t.words {
		if i > len(s.words) {
			s.words = append(s.words, tword)
		} else {
			s.words[i] |= tword
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range s.words {
		// None of the bits are set
		if word == 0 {
			continue
		}

		// Iterate over each bit
		for j := 0; j < 64; j++ {
			// Check if bit is set
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

func (s *IntSet) Copy() *IntSet {
	cp := new(IntSet)
	cp.words = make([]uint64, len(s.words))
	copy(cp.words, s.words)
	cp.length = s.length
	return cp
}

func main() {
	var x, y IntSet
	x.AddAll(1, 3)
	fmt.Println("x: ", x.String(), " len= ", x.length)
	fmt.Printf("x has 1? %v\n", x.Has(1))
	x.Remove(1)
	fmt.Println("x after removing 1: ", x.String())
	fmt.Printf("x has 1? %v\n", x.Has(1))
	fmt.Printf("x has 2? %v\n", x.Has(1))
	x.Add(5)
	fmt.Println("x: ", x.String())

	y.Add(2)
	y.Add(3)
	fmt.Println("y: ", y.String())

	x.UnionsWith(&y)
	fmt.Println("x after union with y: ", x.String())

	z := x.Copy()
	fmt.Printf("z: %v", z)
}
