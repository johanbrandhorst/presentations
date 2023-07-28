package main

import "fmt"

func main() {
	var s []string
	s = append(s, "a")
	fmt.Println(s)

	var m map[string]string
	m["a"] = "b"
	fmt.Println(m)
}
