package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	reader := bufio.NewReader(conn)

	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
			} else {
				fmt.Printf("Read error: %v\n", err)
			}
			return
		}

		fmt.Printf("Received: %q\n", buf[:n])
		response := "+OK\r\n"
		_, err = conn.Write([]byte(response))
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
