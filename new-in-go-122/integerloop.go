package main

import "fmt"

// START OMIT
func main() {
	for _, v := range []int{0, 1, 2} {
		fmt.Println("v is", v)
	}
	for v := range 3 {
		fmt.Println("v is", v)
	}
}

// END OMIT
