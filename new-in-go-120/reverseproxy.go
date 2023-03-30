package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

// START OMIT
func main() {
	proxyHandler := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(&url.URL{
				Scheme: "https",
				Host:   "example.com",
			}) // Forward request to https://example.com/.
			r.SetXForwarded() // Set X-Forwarded-* headers.
			r.Out.Header.Set("X-Additional-Header", "header set by the proxy")
		},
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	proxyHandler.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())
	fmt.Println(rec.HeaderMap)
}
// END OMIT
