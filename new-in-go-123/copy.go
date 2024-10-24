package main

import (
	"fmt"
	"os"
)

func main() {
	rootFS := os.DirFS(".")

	err := os.CopyFS("/tmp", rootFS)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Copied current directory to /tmp")
}
