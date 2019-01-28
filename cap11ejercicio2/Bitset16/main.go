package bitset16

import (
	"bytes"
	"fmt"
)


type IntSet struct {
	words []uint16
}


func NewIntSet() *IntSet {
	return &IntSet{words: []uint16{}}
}


func (s *IntSet) Has(x int) bool {
	word, bit := x/16, uint(x%16)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}


func (s *IntSet) Add(x int) {
	word, bit := x/16, uint(x%16)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}


func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}


func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 16; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 16*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}


func (s *IntSet) Len() int {
	count := 0
	for _, w := range s.words {
		for w != 0 {
			w &= w - 1
			count++
		}
	}
	return count
}


func (s *IntSet) Remove(x int) {
	word, bit := x/16, uint(x%16)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}


func (s *IntSet) Clear() {
	s.words = []uint16{}
}


func (s *IntSet) Copy() *IntSet {
	out := &IntSet{}
	for _, w := range s.words {
		out.words = append(out.words, w)
	}
	return out
}


func (s *IntSet) AddAll(vals ...int) {
	for _, x := range vals {
		word, bit := x/16, uint(x%16)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}


func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}


func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}


func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}
}


func (s *IntSet) Elems() []int {
	out := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 16; j++ {
			if word&(1<<uint(j)) != 0 {
				out = append(out, 16*i+j)
			}
		}
	}
	return out
}