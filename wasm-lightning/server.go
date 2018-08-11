package main

import (
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./html"))
	http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}

		fs.ServeHTTP(resp, req)
	}))
}
