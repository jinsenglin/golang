package main

import "fmt"
import "net"
import "bufio"

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
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
				data, err := buf.ReadString('\n')

				if err != nil {
					fmt.Println("disconnected client")
					break
				}

				conn.Write([]byte(data))
			}
		}()
	}

	fmt.Println("exit 0")
}
