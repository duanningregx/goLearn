package main

import (
	"fmt"
	"net"
)

func main() {
	const strTcp = "tcp4"
	tcpAddr, err := net.ResolveTCPAddr(strTcp, "127.0.0.1:8080")
	if err != nil {
		return
	}

	tcp, err := net.ListenTCP(strTcp, tcpAddr)
	if err != nil {
		return
	}

	conn, err := tcp.Accept()
	if err != nil {
		return
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	fmt.Println(string(buf[:n]))

	_, err = conn.Write([]byte("hello! i got it"))
	if err != nil {
		return
	}
}
