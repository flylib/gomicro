package main

import (
	"context"
	"github.com/zjllib/go-micro"
	"github.com/zjllib/go-micro/plugins/registry/etcd"
	"github.com/zjllib/go-micro/plugins/transport/grpc"
	"hello/proto"
	"log"
)

func main() {
	registry := etcd.NewRegistry(
		etcd.Endpoints("127.0.0.1:2379"),
	)

	service := micro.NewService(
		micro.Name("test"),
		micro.Transport(grpc.NewTransport()),
		micro.Registry(registry),
	)

	client, err := service.NewClient(micro.CallTarget("test"))
	if err != nil {
		panic(err)
	}

	waiterService := proto.NewWaiterClient(client)

	res, err := waiterService.DoMD5(context.Background(), &proto.Req{
		JsonStr: "hello",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("micro server response: %s", res.BackJson)
}
