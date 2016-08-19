package main

import (
	"io"
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	fmt.Println("Server listening on port 0.0.0.0:5000")
	http.HandleFunc("/", hello)
	http.ListenAndServe("0.0.0.0:5000", nil)
}