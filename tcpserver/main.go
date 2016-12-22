package main

import "fmt"
import "net"
import "bufio"

func check(err error, message string) {
	if err != nil {
		panic(err)
	}

	if message != "" {
		fmt.Println(message)
	}
}

func main() {
	fmt.Println("starting tcp server 0.0.0.0:8080 ...")

	ln, err := net.Listen("tcp", ":8080")
	check(err, "started tcp server 0.0.0.0:8080")

	for {
		conn, err := ln.Accept()
		check(err, "accepted connection")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				message, err := buf.ReadString('\n')

				if err != nil {
					fmt.Println("disconnected client")
					break
				}

				fmt.Println("message :", message)
				conn.Write([]byte(message))
			}
		}()
	}

	fmt.Println("exit 0")
}
