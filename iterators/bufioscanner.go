package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scan := bufio.NewScanner(strings.NewReader("foo\nbar\nbaz\n"))
	for scan.Scan() {
		fmt.Println(scan.Text())
	}
}
