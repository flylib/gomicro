package grpc

import (
	. "github.com/zjllib/go-micro"
)

type grpcTransport struct {
	s *server
	c *client
}

func NewTransport(opts ...OptionFun) ITransport {
	var options Option
	for _, o := range opts {
		o(&options)
	}
	s:=&server{opt: options}
	c:=&client{opt: options}
	return &grpcTransport{s,c}
}

func (self *grpcTransport)Init(opts ...OptionFun)  error {
	var options Option
	for _, o := range opts {
		o(&options)
	}
	s:=&server{opt: options}
	c:=&client{opt: options}
	self.s = s
	self.c = c
	return nil
}


func (self *grpcTransport) Server() IServer {
	return self.s
}

func (self *grpcTransport) Client() IClient {
	return self.c
}
