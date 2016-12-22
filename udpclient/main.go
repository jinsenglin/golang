package main 

import "fmt"
import "net"
import "strconv"
import "time"

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func main() {
	fmt.Println("starting udp client to server 0.0.0.0:8080 ...")

	server_addr,err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	check(err, "")

	local_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	check(err, "")

	conn, err := net.DialUDP("udp", local_addr, server_addr)
	check(err, "")

	defer conn.Close()

	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := conn.Write(buf)
		check(err, "")
		time.Sleep(time.Second * 1)
	}

	fmt.Println("exit 0")
}
