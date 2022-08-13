package stub

import "net/rpc"

type IHelloService interface {
	Hello(req string, resp *string) error
}

func RegisterHelloService(srv IHelloService) error {
	return rpc.RegisterName(HelloServiceName, srv)
}
