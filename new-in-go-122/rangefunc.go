package main

import (
	"fmt"
	"iter"
)

type myDataStructure struct {
	nums    []int
	strings []string
}

// START OMIT
func main() {
	m := &myDataStructure{
		nums:    []int{1, 2, 3},
		strings: []string{"one", "two", "three"},
	}
	for num, str := range m.Walk() {
		fmt.Println(num, str)
	}
}

func (m *myDataStructure) Walk() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, num := range m.nums {
			if !yield(num, m.strings[i]) {
				break
			}
		}
	}
}

// END OMIT
