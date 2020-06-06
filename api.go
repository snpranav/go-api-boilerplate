package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":3000", Handler: http.HandlerFunc(apiHandler)}
	log.Printf("Serving on https://0.0.0.0:3000")
	log.Fatal(srv.ListenAndServe())
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	log.Printf("Got connection: %s", r.Proto)
	// Send a message back to the client
	w.Write([]byte("Hello World! The Go API is running inside a container using the go-api-boilerplate docker image.\nPlease visit - https://hub.docker.com/r/snpranav/go-api-boilerplate"))
}
