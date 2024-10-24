package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"a", "b", "c"}
	it := slices.Values(s)
	fmt.Println(slices.Collect(it))
}
