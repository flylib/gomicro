package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
	opt        option
	grpcServer *grpc.Server
}

func (self *server) Start() error {
	ln, err := net.Listen("tcp", self.opt.addres)
	if err != nil {
		return err
	}
	self.grpcServer = grpc.NewServer(grpc.StatsHandler(&StatsHandler{})) //创建gRPC服务
	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */

	//proto.RegisterWaiterServer(grpcServer, &server{})
	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(self.grpcServer)
	// 将监听交给gRPC服务处理
	err = self.grpcServer.Serve(ln)
	return err
}

func (self *server) Stop() error {
	self.grpcServer.Stop()
	return nil
}

func (self *server) Addr() string {
	return self.opt.addres
}
