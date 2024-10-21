package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	// blocking until a connection is accepted
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')

		fmt.Printf(">> read: %v \n", str)

		if str != "PING\r\n" {
			continue
		}

		if err != nil {
			if err == io.EOF {
				fmt.Println(">> eof")
				break
			}
			fmt.Println(">> other err")
			break
		}

		res := []byte("+PONG\r\n")
		_, err = conn.Write(res)
		if err != nil {
			fmt.Println("Failed to respond PONG")
			break
		}
		fmt.Println(">> sent pong")
	}

}
