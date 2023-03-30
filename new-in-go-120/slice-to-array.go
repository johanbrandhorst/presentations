package main

var a0 [0]byte
var a1 [1]byte
var a2 [2]byte
var a4 [4]byte

// START OMIT
func main() {
	s := make([]byte, 2, 4)
	a0 = [0]byte(s)
	a1 = [1]byte(s[1:]) // a1[0] == s[1]
	a2 = [2]byte(s)     // a2[0] == s[0]
	//a4 = [4]byte(s)     // panics: len([4]byte) > len(s)
}
// END OMIT
