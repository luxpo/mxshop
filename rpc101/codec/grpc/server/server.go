package main

import (
	"net"
	"net/rpc"

	"github.com/luxpo/mxshop/rpc101/codec/grpc/stub"
)

func main() {
	listener, _ := net.Listen("tcp", ":1234")
	stub.RegisterHelloService(&stub.NewHelloService{})
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
