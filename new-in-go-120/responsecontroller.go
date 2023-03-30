package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func main() {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	RequestHandler(rec, req)
	fmt.Println(rec.Body.String())
}

// START OMIT
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// No more w.(http.Flusher), w.(http.Hijacker), etc
	rc := http.NewResponseController(w)
	rc.SetWriteDeadline(time.Time{}) // disable Server.WriteTimeout when sending a large response
	// Also:
	// rc.Flush()
	// rc.Hijack()
	// rc.SetWriteDeadline(time.Time{})
	io.Copy(w, strings.NewReader("large amounts of data"))
}
// END OMIT
