package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./closing.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	contents, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(contents[:10]))
}
