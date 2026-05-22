package main

import (
	"fmt"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[:n]))

	}

}

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)

		fmt.Println("client connected - remote:", conn.RemoteAddr(), "local:",
			conn.LocalAddr())
	}

}
