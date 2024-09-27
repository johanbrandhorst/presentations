package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	}); err != nil {
		fmt.Println(err)
		return
	}
}
