package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Hello, world!")
	if err != nil {
		panic(err)
	}
	log.Printf("Host: %s", req.Host)
}
