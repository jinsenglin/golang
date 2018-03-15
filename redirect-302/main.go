package main

import (
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.google.com", http.StatusFound)
	return
}

func main() {
	http.HandleFunc("/", redirect)
	http.ListenAndServe(":8000", nil)
}
