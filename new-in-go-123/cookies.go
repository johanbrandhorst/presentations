package main

import (
	"fmt"
	"net/http"
)

// START OMIT
func main() {
	// Parses request cookies
	line := "session_id=abc123; dnt=1; lang=en; lang=de"
	cookies, err := http.ParseCookie(line)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cookie := range cookies {
		fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)
	}
	// Can also parse response header cookies (Set-Cookie)
	line = "session_id=abc123; SameSite=None; Secure; Partitioned; Path=/; Domain=.example.com"
	cookie, err := http.ParseSetCookie(line)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cookie.Name, cookie.Partitioned, cookie.Path, cookie.Domain, cookie.Secure)
}

// END OMIT
