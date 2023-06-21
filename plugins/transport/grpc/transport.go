package grpc

//type ITransport interface {
//	Server() IServer
//	Client() IClient
//}

type transport struct {
	s *server
	c *client
}

func (self transport) Server() *server {
	return self.s
}
func (self transport) Client() *client {
	return self.c
}
