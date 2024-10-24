package main

import (
	"fmt"
	"sync/atomic"
)

const (
	modeRead  = 0b100
	modeWrite = 0b010
	modeExec  = 0b001
)

func main() {
	var mode atomic.Int32
	mode.Store(modeRead)
	old := mode.Or(modeWrite)

	fmt.Printf("mode: %b -> %b\n", old, mode.Load())
}
