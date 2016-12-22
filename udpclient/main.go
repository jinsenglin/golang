package main 

import "fmt"
import "net"
import "strconv"
import "time"

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}

func main() {
	fmt.Println("starting udp client to server 0.0.0.0:8080 ...")

	server_addr,err := net.ResolveUDPAddr("udp", ":8080")
	check(err, "prepared udp server addr :8080")

	local_addr, err := net.ResolveUDPAddr("udp", ":0")
	check(err, "prepared udp client addr :0")

	conn, err := net.DialUDP("udp", local_addr, server_addr)
	check(err, "")

	defer conn.Close()

	i := 0
	for {
		msg := strconv.Itoa(i)
		buf := []byte(msg)
		_, err := conn.Write(buf)
		check(err, "")

		n, addr, err := conn.ReadFromUDP(buf)
		check(err, "")
		fmt.Println("message :", string(buf[0:n]), " from", addr) 

		time.Sleep(time.Second * 1)
		i++
	}

	fmt.Println("exit 0")
}
