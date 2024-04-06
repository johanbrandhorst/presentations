package main

import "fmt"

func main() {
	v := [3]int{1, 2, 3}

	// Arrays are value types, copies are deep copies
	v2 := v
	v2[0] = 100
	fmt.Println(v)  // [1 2 3]
	fmt.Println(v2) // [100 2 3]

	// Cannot use in append directly
	// v3 = append(v, 4) // Compiler error

	// But they can be sliced!
	v4 := append(v[:], 4) // OK
	fmt.Println(v4)       // [1 2 3 4]
	fmt.Println(v[:2])    // [1 2]
}
