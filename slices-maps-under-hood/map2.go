package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Alice"] = 1

	m2 := m
	m2["Bob"] = 2

	fmt.Println(m["Bob"])
}
