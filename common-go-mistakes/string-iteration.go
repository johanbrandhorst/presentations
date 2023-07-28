package main

import "fmt"

func main() {
	s := "hello"
	for i, val := range s {
		fmt.Println(i, string(val), string(s[i]))
	}
	s = "👍🆗👌💯"
	for i, val := range s {
		fmt.Println(i, string(val), string(s[i]))
	}
}
