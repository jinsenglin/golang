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
	fmt.Println("starting udp server 0.0.0.0:8080 ...")

	addr, err := net.ResolveUDPAddr("udp",":8080")
	check(err, "")

	conn, err := net.ListenUDP("udp", addr)
	check(err, "started udp server 0.0.0.0:8080 ...")

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		check(err, "")

		go func(conn *net.UDPConn, addr *net.UDPAddr, message string) {
			fmt.Println("message :", message)
			_, err := conn.WriteToUDP([]byte(message), addr)
			check(err, "")
		}(conn, addr, string(buf[0:n]))
	}

	fmt.Println("exit 0")
}
