package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/command"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
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
	addr       string
	listener   net.Listener
	cmdHandler command.CommandHandler
}

func NewServer(addr string) Server {
	store := store.NewStore()

	return Server{
		addr:       addr,
		cmdHandler: command.NewCommandHandler(&store),
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

		go s.readConnection(conn)

	}

}

func (s *Server) readConnection(conn net.Conn) {
	for {
		err := s.readMsg(conn)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
				return
			}
			fmt.Printf("Error reading message: %v\n", err)
			return
		}
	}
}

func (s *Server) readMsg(conn net.Conn) error {
	reader := bufio.NewReader(conn)

	response, err := s.parse(reader)

	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(response))

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) parse(reader *bufio.Reader) (string, error) {
	m, err := resp.RespRead(reader)

	if err != nil {
		fmt.Println("failed to parse resp message")
		return "", err
	}

	redisResp, err := s.cmdHandler.RunCommand(m)

	if err != nil {
		fmt.Println("failed to run command")
		return "", err
	}

	return redisResp, nil
}

func main() {
	srv := NewServer(":6379")
	log.Fatal(srv.start())
}
