package main

import (
	"fmt"
	"iter"
)

type Set[E comparable] struct {
	m map[E]struct{}
}

func (s *Set[E]) Add(v E) {
	s.m[v] = struct{}{}
}

// START OMIT
// All is an iterator over the elements of s.
func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	t := &Set[int]{m: make(map[int]struct{})}
	t.Add(1)
	t.Add(2)
	t.Add(3)
	for v := range t.All() {
		fmt.Println(v)
	}
}

// END OMIT
