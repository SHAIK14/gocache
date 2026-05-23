package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/SHAIK14/gocache/store"
)

func handleConn(conn net.Conn, s *store.Store) {
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
			s.Set(parts[1], parts[2])
			conn.Write([]byte("OK\n"))
		} else if parts[0] == "GET" && len(parts) == 2 {
			val, ok := s.Get(parts[1])
			if !ok {
				conn.Write([]byte("-ERR not found \n"))

			} else {

				conn.Write([]byte(val + "\n"))
			}

		} else if parts[0] == "DEL" && len(parts) == 2 {
			s.Del(parts[1])
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
	s := store.NewStore()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn, s)

		fmt.Println("client connected - remote:", conn.RemoteAddr(), "local:",
			conn.LocalAddr())
	}

}
