package main

import (
	"context"
	"log"
	"net"

	"github.com/luxpo/mxshop/rpc101/proto/helloworld"
	"google.golang.org/grpc"
)

type Server struct {
	helloworld.UnimplementedHelloServer
}

func (s *Server) Hello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.Response, error) {
	return &helloworld.Response{
		Reply: "hello " + req.Name,
	}, nil
}

func main() {
	s := grpc.NewServer()
	helloworld.RegisterHelloServer(s, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
