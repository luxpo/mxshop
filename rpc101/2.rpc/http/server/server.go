package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = "hello " + req
	return nil
}

func main() {
	rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}
