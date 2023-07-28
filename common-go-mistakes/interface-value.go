package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// START OMIT
type myStruct struct{}

func (m myStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func main() {
	var myHandler http.Handler
	var myStructHandler http.Handler = (*myStruct)(nil)
	if myHandler == nil {
		fmt.Println("myHandler is nil")
	}
	if myStructHandler == nil {
		fmt.Println("myStructHandler is nil")
	} else {
		fmt.Println("Great, myStructHandler is not nil!")
		myStructHandler.ServeHTTP(httptest.NewRecorder(), &http.Request{})
	}
}

// END OMIT
