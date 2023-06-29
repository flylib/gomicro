package main

import (
	"github.com/zjllib/go-micro"
	"github.com/zjllib/go-micro/plugins/transport/grpc"
)

func main() {
	transport := grpc.NewTransport(
		grpc.Address(":8099"))

	service := micro.NewService(
		micro.Transport(transport),
	)
	service.Start()
}
