package grpc

import (
	"github.com/zjllib/go-micro/transport"
)

type grpcTransport struct {
	s *server
	c *client
}

func NewTransport(opts ...OptionFun) transport.ITransport {
	var options Option
	for _, o := range opts {
		o(&options)
	}
	s:=&server{opt: options}
	c:=&client{opt: options}
	return &grpcTransport{s,c}
}



func (self grpcTransport) Server() transport.IServer {
	return self.s
}

func (self grpcTransport) Client() transport.IClient {
	return self.c
}
