package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	resp := new(string)
	err = client.Call("HelloService.Hello", "luxcgo", resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*resp)
}
