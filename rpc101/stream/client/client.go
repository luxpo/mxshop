package main

import (
	"context"
	"fmt"
	"log"
	"sync"
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

	// 双向流模式
	allStream, _ := c.AllStream(context.Background())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStream.Recv()
			log.Printf("recv from server: %s", data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allStream.Send(&proto.StreamReqData{
				Data: "I'm client.",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
