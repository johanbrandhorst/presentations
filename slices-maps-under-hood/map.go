package main

func main() {
	m := make(map[string]int)
	// Equivalent to
	m2 := map[string]int{}

	// Panic waiting to happen
	var m3 map[string]int

	// Set the size ahead of time
	m4 := make(map[string]int, 10)
}
