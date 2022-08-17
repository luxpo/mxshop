package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/luxpo/mxshop/rpc101/stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *server) AllStream(proto.Greeter_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
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
