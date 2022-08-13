package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = "hello " + req
	return nil
}

func main() {
	// 1. 实例化server
	lis, err := net.Listen("tcp", ":1234")
	checkError(err)
	// 2. 注册handler
	err = rpc.RegisterName("HelloService", &HelloService{})
	checkError(err)
	// 3. 启动服务
	for {
		conn, err := lis.Accept()
		checkError(err)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
