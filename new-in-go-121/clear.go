package main

import (
	"fmt"
	"math"
)

func main() {
	m := map[float64]string{
		math.NaN(): "1",
	}
	for k := range m {
		delete(m, k)
	}
	fmt.Println(m)
}
