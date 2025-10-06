package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		parts := strings.Split(string(buf), "\r\n")
		parts = parts[:len(parts)-1]
		fmt.Println(parts)
		if len(parts) == 7 {
			if strings.ToLower(parts[2]) == "set" {
				res := ReadSet(parts)
				fmt.Println(res)
			}
		}
		if len(parts) == 5 {
			if strings.ToLower(parts[2]) == "get" {
				res := ReadGet(parts[4])
				fmt.Println(res)
			}
		}
		conn.Write([]byte("+OK\r\n"))
	}
}
