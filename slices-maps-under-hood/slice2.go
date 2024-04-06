package main

import "fmt"

func main() {
	// Copies are not deep
	v := []int{1, 2, 3}
	v2 := v
	v2[0] = 100
	fmt.Println(v)              // [100 2 3]
	fmt.Println(len(v), cap(v)) // 3 3

	// Appending creates a new array
	v3 := append(v, 4)
	fmt.Println(len(v3), cap(v3)) // 4 6

	// Slicing creates a new slice over the same array
	v4 := v[:2]
	fmt.Println(len(v4), cap(v4)) // 2 3
}
