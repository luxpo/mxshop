package main

import (
	"context"
	"log"

	"github.com/luxpo/mxshop/rpc101/stream/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	resp, err := c.GetStream(context.Background(), &proto.StreamReqData{
		Data: "luxcgo",
	})
	for {
		data, err := resp.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(data)
	}
}
