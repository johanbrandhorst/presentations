package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan struct{})
	for i, val := range []string{"a", "b", "c", "d"} {
		go func() {
			fmt.Println(i, val)
			<-ch
		}()
	}
	fmt.Println("Eeep, we still have", runtime.NumGoroutine(), "goroutines running!")
}
