package main

import (
	"fmt"

	"github.com/luxpo/mxshop/rpc101/proto/helloworld"
	"google.golang.org/protobuf/proto"
)

func main() {
	req := &helloworld.HelloRequest{
		Name: "luxcgo",
	}
	resp, _ := proto.Marshal(req)
	fmt.Println(string(resp))
}
