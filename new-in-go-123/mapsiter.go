package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	m := map[int]string{0: "zero", 1: "one", 2: "two", 3: "three"}
	// Note: iteration order is not defined!
	fmt.Println(slices.Collect(maps.Keys(m)))
	fmt.Println(slices.Collect(maps.Values(m)))
	fmt.Println(maps.Collect(slices.All(slices.Collect(maps.Values(m)))))
}
