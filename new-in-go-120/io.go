package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	err = write(f)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(f *os.File) (string, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// START OMIT
func write(f *os.File) error {
	w := io.NewOffsetWriter(f, 4)
	mw := io.MultiWriter(f, w)
	_, err := mw.Write([]byte("test"))
	if err != nil {
		return err
	}
	s, err := readFile(f)
	fmt.Println(s)
	return nil
}
// END OMIT
