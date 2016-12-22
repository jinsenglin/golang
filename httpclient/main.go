package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("starting http client to get http://localhost:8080")

	resp, err := http.Get("http://localhost:8080")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("exit 0")
}
