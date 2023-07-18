package main

import (
	"github.com/zjllib/go-micro"
	"github.com/zjllib/go-micro/plugins/registry/etcd"
	"github.com/zjllib/go-micro/plugins/transport/grpc"
	"hello/handler"
	"hello/proto"
	"time"
)

func main() {
	registry := etcd.NewRegistry(
		etcd.Address("127.0.0.1:2379"),
		etcd.RegisterTTL(time.Second*5),
	)

	transport := grpc.NewTransport(
		grpc.Address(":8090"),
		grpc.M(proto.RegisterWaiterServer, &handler.MD5Handler{}),
	)

	service := micro.NewService(
		micro.Name("test"),
		micro.Transport(transport),
		micro.Registry(registry),
	)

	service.Start()
}
