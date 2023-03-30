package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	s := "somestring"
	s3, ok := strings.CutPrefix(s, "some")
	fmt.Println(s3, ok)
	s4, ok := strings.CutSuffix(s, "string")
	fmt.Println(s4, ok)
}
// END OMIT
