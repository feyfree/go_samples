package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("Listen error: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			break
		}
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	packet := make([]byte, 1024)
	for {
		n, err := conn.Read(packet)
		if err != nil {
			log.Println("Read socket error:", err)
			return
		}
		log.Println(packet[:n])
		_, _ = conn.Write(packet[:n])
	}
}
