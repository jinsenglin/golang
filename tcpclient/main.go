package main

import "fmt"
import "net"

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

	fmt.Println("starting tcp client to server 0.0.0.0:8080 ...")

	conn, err := net.Dial("tcp", ":8080")
	check(err, "")

	defer conn.Close()

	fmt.Fprintf(conn, data)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	check(err, "")

	message := string(buf[0:n])
	fmt.Println("message :", message)

	fmt.Println("exit 0")
}
