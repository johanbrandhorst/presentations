package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

// START OMIT
func main() {
	s := []int{0, 3, 2, 1}
	slices.Sort(s)
	fmt.Println(s) // Sorted
	slices.SortFunc(s, func(i, j int) int {
		if j < i {
			return -1
		}
		if j > i {
			return 1
		}
		return 0
	})
	fmt.Println(s) // Reversed
	m := map[string]struct{}{"a": {}, "b": {}}
	fmt.Println(maps.Clone(m))
	slices.SortFunc(s, cmp.Compare) // Use cmp.Compare as comparer, handling NaNs
	fmt.Println(s)
}

// END OMIT
