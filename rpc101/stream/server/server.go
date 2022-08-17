package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/luxpo/mxshop/rpc101/stream/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) GetStream(req *proto.StreamReqData, resp proto.Greeter_GetStreamServer) error {
	for i := 0; i < 3; i++ {
		resp.Send(&proto.StreamRespData{
			Data: fmt.Sprintf("Hi %s, date is %s", req.Data, time.Now().Format(time.RFC3339)),
		})
		time.Sleep(time.Second)
	}

	return nil
}

func (s *server) PutStream(req proto.Greeter_PutStreamServer) error {
	for {
		if reqData, err := req.Recv(); err != nil {
			log.Println(err)
			break
		} else {
			log.Println(reqData.Data)
		}
	}
	return nil
}

func (s *server) AllStream(greeter proto.Greeter_AllStreamServer) error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := greeter.Recv()
			log.Printf("recv from client: %s\n", data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			greeter.Send(&proto.StreamRespData{
				Data: "I'm server.",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
