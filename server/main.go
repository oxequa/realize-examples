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
	fmt.Println("Watch on 0.0.0.0:5000")
	http.HandleFunc("/", hello)
	http.ListenAndServe("0.0.0.0:5000", nil)
}