package grpc

import (
	"github.com/zjllib/go-micro"
)

type grpcTransport struct {
	s *server
	c *client
}

func NewTransport(opts ...Option) micro.ITransport {
	var options option
	for _, o := range opts {
		o(&options)
	}
	s := &server{opt: options}
	c := &client{opt: options}
	return &grpcTransport{s, c}
}

func (self *grpcTransport) Server() micro.IServer {
	return self.s
}

func (self *grpcTransport) Client() micro.IClient {
	return self.c
}
