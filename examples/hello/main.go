package main

import (
	"github.com/zjllib/go-micro"
	"github.com/zjllib/go-micro/plugins/registry/etcd"
	"github.com/zjllib/go-micro/plugins/transport/grpc"
	"time"
)

func main() {

	registry := etcd.NewRegistry(
		etcd.Address("127.0.0.1:2379"),
		etcd.RegisterTTL(time.Second*5),
	)

	transport := grpc.NewTransport(
		grpc.Addres(":8090"),
	)

	service := micro.NewService(
		micro.Transport(transport),
		micro.Registry(registry),
	)
	service.Start()
}
