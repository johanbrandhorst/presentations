package main

// START OMIT
import (
	"log"
	"net/http"

	"github.com/stealthrocket/net/wasip1"
)

func main() {
	listener, err := wasip1.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!\n"))
		}),
	}
	_ = server.Serve(listener)
}

// END OMIT
