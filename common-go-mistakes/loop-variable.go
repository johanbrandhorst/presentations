package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i, val := range []string{"a", "b", "c", "d"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, val)
		}()
	}
	wg.Wait()
}
