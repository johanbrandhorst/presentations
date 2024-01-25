package main

import "fmt"

// START OMIT
func main() {
	nums := []int{1, 2, 4, 6}
	var fs []func()
	for _, v := range nums {
		fs = append(fs, func() {
			fmt.Println("num is", v)
		})
	}
	for _, f := range fs {
		f()
	}
}

// END OMIT
