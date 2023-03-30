package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	t := time.Now()
	fmt.Println(t.Format(time.DateTime))
	fmt.Println(t.Format(time.DateOnly))
	fmt.Println(t.Format(time.TimeOnly))
	t2 := time.Now()
	// -1 for smaller, 0 for equal, 1 for greater
	fmt.Println(t.Compare(t2))
}
// END OMIT
