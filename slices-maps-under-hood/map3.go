package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[int]string{1: "one", 10: "Ten", 1000: "THOUSAND"}
	m2 := map[int]string{1: "one", 10: "Ten", 1000: "THOUSAND"}
	m3 := map[int]string{1: "one", 10: "ten", 1000: "thousand"}

	fmt.Println(maps.Equal(m1, m2))
	fmt.Println(maps.Equal(m1, m3))
}
