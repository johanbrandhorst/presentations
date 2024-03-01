package main

import "fmt"

// START OMIT
func main() {
	fmt.Println(maxInt(4, 7))
	fmt.Println(maxFloat(4.4, 7.1))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// END OMIT
