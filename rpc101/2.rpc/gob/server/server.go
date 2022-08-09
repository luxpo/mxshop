package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = "hello " + req
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":1234")
	checkError(err)
	err = rpc.RegisterName("HelloService", &HelloService{})
	checkError(err)
	conn, err := lis.Accept()
	checkError(err)
	rpc.ServeConn(conn)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
