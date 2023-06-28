package hello

import (
	"github.com/zjllib/go-micro"
	"github.com/zjllib/go-micro/plugins/transport/grpc"
)

func main() {
	service := micro.NewService(
		micro.Address(":8099"),
		micro.Transport(grpc.NewTransport()),
	)
	service.Start()
}
