package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	if err != nil {
		return
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		return
	}
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(conn)
	_, err = conn.Write([]byte(" hello server"))
	if err != nil {
		return
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	fmt.Println(string(buf[:n]))
}
