package main

import "fmt"
import "net/http"
import "io/ioutil"

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}

func main() {
	fmt.Println("starting http client to get http://localhost:8080/message")

	resp, err := http.Get("http://localhost:8080/message")
	check(err, "")

	defer resp.Body.Close()

	message, err := ioutil.ReadAll(resp.Body)
	check(err, "")

	fmt.Println("message :", string(message))

	fmt.Println("exit 0")
}
