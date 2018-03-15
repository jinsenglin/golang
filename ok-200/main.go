package main

import (
	"io"
	"net/http"
)

func ok(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", ok)
	http.ListenAndServe(":8000", nil)
}
