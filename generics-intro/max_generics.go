package main

import (
	"cmp"
	"fmt"
)

// START OMIT
func main() {
	fmt.Println(max(4, 7))
	fmt.Println(max(4.4, 7.1))
}

func max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Or even MORE GENERIC
func megaMax[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// END OMIT
