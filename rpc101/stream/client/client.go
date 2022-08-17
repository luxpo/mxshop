package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/luxpo/mxshop/rpc101/stream/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 服务端流模式
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

	// 客户端流模式
	greeter, _ := c.PutStream(context.Background())
	for i := 0; i < 3; i++ {
		greeter.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("putstream %d", i),
		})
		time.Sleep(time.Second)
	}

}
