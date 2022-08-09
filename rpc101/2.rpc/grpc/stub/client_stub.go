package stub

import "net/rpc"

type HelloServiceClientStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, addr string) (*HelloServiceClientStub, error) {
	conn, err := rpc.Dial(protocol, addr)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClientStub{
		Client: conn,
	}, nil
}

func (this *HelloServiceClientStub) Hello(req string, resp *string) error {
	return this.Call(HelloServiceName+".Hello", req, resp)
}
