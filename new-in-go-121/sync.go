package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	var i int
	// Also sync.OnceValue, sync.OnceValues for
	// 1, 2 return values respectively
	onceFn := sync.OnceFunc(func() {
		i++
	})
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start // Wait for start to be closed
			onceFn()
		}()
	}
	close(start) // Unblock all goroutines
	wg.Wait()
	fmt.Println(i)
}

// END OMIT
