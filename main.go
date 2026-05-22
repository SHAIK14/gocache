package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConn(conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		result := string(buf[:n])
		parts := strings.Fields(result)
		if len(parts) == 0 {
			return
		}
		if parts[0] == "SET" && len(parts) == 3 {
			fmt.Println(parts)
			conn.Write([]byte("OK\n"))
		} else if parts[0] == "GET" && len(parts) == 2 {
			fmt.Println(parts)
			conn.Write([]byte("OK\n"))
		} else if parts[0] == "DEL" && len(parts) == 2 {
			fmt.Println(parts)
			conn.Write([]byte("OK\n"))
		} else {
			fmt.Println("data sahi se de bsdk")
			conn.Write([]byte("-ERR unknown command\n"))

		}

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
