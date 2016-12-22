package main

import "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request in")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("starting http server 0.0.0.0:8080 ...")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("exit 0")
}
