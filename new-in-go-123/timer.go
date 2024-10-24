package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	const timeout = 10 * time.Millisecond
	t := time.NewTimer(timeout)
	time.Sleep(20 * time.Millisecond)

	start := time.Now()
	t.Reset(timeout)
	<-t.C
	fmt.Printf("Time elapsed: %dms\n", time.Since(start).Milliseconds())
	// expected: Time elapsed: 10ms
	// actual:   Time elapsed:  0ms
}

// END OMIT
