package grpc

import (
	"github.com/zjllib/go-micro"
)

type grpcTransport struct {
	opt option
	s   *server
	//c *client
}

func NewTransport(opts ...Option) micro.ITransport {
	var options option
	for _, o := range opts {
		o(&options)
	}
	s := &server{opt: options}
	//c := &client{opt: options}
	return &grpcTransport{
		opt: options,
		s:   s}
}

func (self *grpcTransport) Server() micro.IServer {
	return self.s
}

func (self *grpcTransport) Client() micro.IClient {
	return &client{
		opt: self.opt,
	}
}

func (self *grpcTransport) Type() string {
	return "grpc"
}
