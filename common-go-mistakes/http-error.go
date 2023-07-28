package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

// START OMIT
type myStruct struct {
	SomeVal string
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	var s myStruct
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
	}
	fmt.Println("Happily handling request:", s)
}

func main() {
	rec := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"someVal": "a"}`))
	myHandler(rec, r)
	fmt.Println("Status", rec.Result().Status)
	rec = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`not json`))
	myHandler(rec, r)
	fmt.Println("Status", rec.Result().Status)
}

// END OMIT
