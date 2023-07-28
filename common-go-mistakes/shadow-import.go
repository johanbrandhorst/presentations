package main

import "time"

func main() {
	now := time.Now()
	time := now.Add(time.Hour)
	nowAgain := time.Now()
}
