package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(slices.Repeat(s, 2))
}
