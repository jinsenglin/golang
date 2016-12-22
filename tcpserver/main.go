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
	fmt.Println("starting tcp server 0.0.0.0:8080 ...")

	ln, err := net.Listen("tcp", ":8080")
	check(err, "started tcp server 0.0.0.0:8080")

	for {
		conn, err := ln.Accept()
		check(err, "accepted connection")

		go func() {
			buf := make([]byte, 1024)

			for {
				n, err := conn.Read(buf)

				if err != nil {
					fmt.Println("disconnected connection")
					break
				}

				message := string(buf[0:n])

				fmt.Println("message :", message)
				conn.Write([]byte(message))
			}
		}()
	}

	fmt.Println("exit 0")
}
