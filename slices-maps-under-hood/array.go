package main

import "fmt"

func main() {
	v := [3]int{1, 2, 3}
	v2 := [...]int{1, 2, 3}
	fmt.Println(v == v2) // true

	// Arrays are constant size
	v3 := [3]int{1, 2}
	fmt.Println(v3) // [1 2 0]
}
