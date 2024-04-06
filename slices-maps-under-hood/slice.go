package main

import (
	"unsafe"
)

// START OMIT
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {
	v := [3]int{1, 2, 3}

	v2 := []int{1, 2, 3}
	// Essentially equivalent to
	v3 := SliceHeader{
		Data: uintptr(unsafe.Pointer(&v[0])),
		Len:  len(v),
		Cap:  cap(v),
	}
	_ = v2
	_ = v3
}

// END OMIT
