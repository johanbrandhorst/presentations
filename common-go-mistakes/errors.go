package main

import (
	"errors"
	"fmt"
)

// START OMIT
type myError struct{}

func (m myError) Error() string { return "an error happened!" }

func someFunc() error {
	return fmt.Errorf("some error happened: %w", &myError{})
}

func main() {
	err := someFunc()
	if err != nil {
		if errors.Is(err, myError{}) {
			fmt.Println("it was a my error")
		}
		var mErr myError
		if errors.As(err, &mErr) {
			fmt.Println("unwrapped the error:", mErr)
		}
	}
}

// END OMIT
