package main

import (
	"fmt"
	"log"

	"github.com/luxpo/mxshop/rpc101/codec/grpc/stub"
)

func main() {
	client, _ := stub.NewHelloServiceClient("tcp", "localhost:1234")
	var resp string
	err := client.Hello("luxcgo", &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
