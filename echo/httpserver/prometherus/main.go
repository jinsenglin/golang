package main

import "fmt"
import "net/http"
import "github.com/prometheus/client_golang/prometheus/promhttp"

func handler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path[0:]
	fmt.Println("message :", message)
	fmt.Fprintf(w, message)
}

func main() {
	fmt.Println("starting http server 0.0.0.0:8080 ...")

	http.HandleFunc("/", handler)
    http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)

	fmt.Println("exit 0")
}
