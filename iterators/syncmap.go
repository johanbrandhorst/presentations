package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("foo", "bar")
	m.Store("baz", "qux")
	m.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	})
}
