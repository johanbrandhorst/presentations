package main

import (
	"errors"
	"fmt"
)

func someFunc() (int, error) {
	i := 5
	var err error
	defer func() {
		if err != nil {
			i = 1
		}
	}()
	// do some operation
	err = errors.New("there was an error!")
	return i, err
}

func main() {
	val, err := someFunc()
	fmt.Println(val, err)
}
