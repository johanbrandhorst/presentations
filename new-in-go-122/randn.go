package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// START OMIT
func main() {
	{
		var max int = 100
		n := rand.N(max)
		fmt.Println("integer n =", n)
	}

	{
		n := rand.N(time.Second)
		fmt.Println("random duration =", n)
	}
}

// END OMIT
