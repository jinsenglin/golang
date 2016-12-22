package main

import "fmt"
import "net"
import "bufio"

func main() {
	fmt.Println("starting tcp client to server 0.0.0.0:8080 ...")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("sending message ...")
	fmt.Fprintf(conn, "message\n")

	fmt.Println("receiving message ...")
	message, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		panic(err)
	}

	fmt.Println("received message :", message)

	fmt.Println("exit 0")
}
