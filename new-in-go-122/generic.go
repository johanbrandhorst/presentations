package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fmt.Println(cmp.Or("", "", "myval", ""))
	fmt.Println(slices.Concat([]int{1, 2, 3}, []int{4, 5, 6}))
}
