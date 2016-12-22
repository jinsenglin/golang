package main 

import "fmt"
import "net"
import "bufio"

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}

func main() {
	data := "message"

	fmt.Println("starting udp client to server 0.0.0.0:8080 ...")

	conn, err := net.Dial("udp", ":8080")
	check(err, "")

	defer conn.Close()

	fmt.Fprintf(conn, data + "\n")

	message, err := bufio.NewReader(conn).ReadString('\n')
	check(err, "")

	fmt.Println("message :", message)

	fmt.Println("exit 0")
}
