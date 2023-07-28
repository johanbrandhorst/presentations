package main

import "fmt"

type myStruct struct {
	someVal string
}

func main() {
	s := []myStruct{{"a"}, {"b"}, {"c"}, {"d"}}
	for _, val := range s {
		val.someVal = "e"
	}
	for _, val := range s {
		fmt.Println(val)
	}
}
