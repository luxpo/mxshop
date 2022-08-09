package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	resp := new(string)
	client.Call("HelloService.Hello", "luxcgo", resp)
	log.Println(*resp)
}
