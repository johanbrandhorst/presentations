package main

import "fmt"

func main() {
	// Iterate over integers
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Iterate over a slice
	s := []string{"a", "b", "c"}
	for _, v := range s {
		fmt.Println(v)
	}
}
