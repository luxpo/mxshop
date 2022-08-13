package main

import (
	"context"
	"log"

	"github.com/luxpo/mxshop/rpc101/proto/helloworld"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := helloworld.NewHelloClient(conn)
	resp, err := c.Hello(context.Background(), &helloworld.HelloRequest{Name: "luxcgo"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Reply)
}
