package main

import (
	"bufio"
	"fmt"
	"go-redis/RESP"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(conn)
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		command, _ := inputReader.ReadString('\n')
		cmd, err := RESP.SerializeCommand(command)
		if err != nil {
			fmt.Printf("error : %s\n", err)
			continue
		}
		_, err = conn.Write(cmd)
		if err != nil {
			fmt.Printf("write error : %s\n", err)
			break
		}

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read error : %s\n", err)
			break
		}
		fmt.Printf("Response: %q\n", response)
	}
}
