package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// START OMIT
var cookie = &http.Cookie{Name: "session", Value: "abc123"}

func handler(w http.ResponseWriter, r *http.Request) {
	// Easily get the cookie we care about from the request
	cookies := r.CookiesNamed("session")
	if len(cookies) > 0 {
		fmt.Fprintf(w, "session cookie = %s", cookies[0].Value)
	} else {
		fmt.Fprint(w, "session cookie not found")
	}
}

func main() {
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(cookie)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// END OMIT
