package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

//  what do we know about tcp
// - bi directional
// - tcp is reliant and in order
// - there is a handshake ack to establish the connection

// golang net package
// listen: tcp listener
// accept: accepts incoming connections [blocking]
// write: write to connection
// read

type Server struct {
	addr     string
	listener net.Listener
}

func NewServer(addr string) Server {
	return Server{
		addr: addr,
	}
}

func (s *Server) start() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		return err
	}

	defer l.Close()
	s.listener = l

	s.accept()

	return nil
}

func (s *Server) accept() error {
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go s.read(conn)

	}

}

func (s *Server) read(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error reading:", err)
			if err == io.EOF {
				break
			}

			continue
		}

		// read until last byte
		log.Printf(">> buffer: %v", buffer)

		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing:", err)
			continue
		}
	}
}

func main() {
	srv := NewServer(":6379")
	log.Fatal(srv.start())
}
