package main

import "fmt"

func main() {
	x, y := 1, 2
	fmt.Println(min(x))                // 1
	fmt.Println(min(x, y))             // 1
	fmt.Println(max(x, y, 10))         // 10
	fmt.Println(max(1, 2.0, 10))       // 10.0 (floating-point kind)
	fmt.Println(max(0, float32(x)))    // type of result is float32
	// var s []string
	// _ = min(s...)                   // invalid: slice arguments are not permitted
	fmt.Println(max("", "foo", "bar")) // "foo" (string kind)
}
