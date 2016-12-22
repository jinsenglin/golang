package main

import "fmt"
import "net/http"
import "io/ioutil"

func check(err error, message string) {
	if err != nil {
		panic(err)
	}

	if message != "" {
		fmt.Println(message)
	}
}

func main() {
	data := "message"

	fmt.Println("starting http client to get http://localhost:8080/" + data)

	resp, err := http.Get("http://localhost:8080/" + data)
	check(err, "")

	defer resp.Body.Close()

	message, err := ioutil.ReadAll(resp.Body)
	check(err, "")

	fmt.Println("message :", string(message))

	fmt.Println("exit 0")
}
