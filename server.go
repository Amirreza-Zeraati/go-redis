package main

import (
	"bufio"
	"fmt"
	"go-redis/RESP"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection")
		}
	}(conn)
	reader := bufio.NewReader(conn)

	for {
		cmd := RESP.DeserializeCommand(reader)
		fmt.Printf("Received: %q\n", cmd)
		_, err := conn.Write([]byte(cmd + "\r\n"))
		if err != nil {
			fmt.Printf("Write error: %v\n", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started and listening on port 6379")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New connection accepted")
		go handleConnection(conn)
	}
}
