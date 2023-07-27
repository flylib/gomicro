package main

import (
	"github.com/zjllib/go-micro"
	"github.com/zjllib/go-micro/plugins/registry/etcd"
	"github.com/zjllib/go-micro/plugins/transport/grpc"
	"hello/handler"
	"hello/proto"
	"log"
	"time"
)

func main() {
	registry := etcd.NewRegistry(
		etcd.Endpoints("127.0.0.1:2379"),
		etcd.RegisterTTL(time.Second*15),
	)

	transport := grpc.NewTransport(
		grpc.Address(":8028"),
		grpc.M(proto.RegisterWaiterServer, &handler.MD5Handler{}),
	)

	service := micro.NewService(
		micro.Name("test"),
		micro.Transport(transport),
		micro.Registry(registry),
		micro.RegistryAddress("127.0.0.1:8028"),
	)

	err := service.Start()
	if err != nil {
		log.Fatal(err)
	}
}
