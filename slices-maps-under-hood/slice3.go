package main

import "fmt"

func main() {
	// Copying slices as values is dangerous
	v := make([]int, 0, 6)
	v = append(v, 1, 2, 3)
	v2 := v
	v2 = append(v2, 4, 5, 6)
	fmt.Println(v, v2)
	v = append(v, 7)
	fmt.Println(v, v2) // ??

	// Use the copy builtin
	v3 := make([]int, len(v))
	copy(v3, v)
	v = append(v, 8)
	fmt.Println(v, v3) // ??

	// Always know what the underlying array is doing!
}
