package stub

// namespace
const HelloServiceName = "handler/HelloService"

type NewHelloService struct{}

func (s *NewHelloService) Hello(req string, resp *string) error {
	*resp = "hello " + req
	return nil
}
