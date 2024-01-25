package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// START OMIT
func main() {
	mux := http.NewServeMux()
	userHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User ID is %q", r.PathValue("user_id"))
	})
	mux.Handle("/v1/users/{user_id}", userHandler)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/users/johan", nil)
	mux.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())
}

// END OMIT
