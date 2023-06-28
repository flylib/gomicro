package micro

import "github.com/zjllib/go-micro/transport"

/*
+---------------------------------------------------+
+				     service						+
+---------------------------------------------------+
+		server			|		client				+
+---------------------------------------------------+
+		registry,bee worker、conn pool、codec					+
+---------------------------------------------------+
+		transport(udp、tcp、ws、quic、rpc)			    +
+---------------------------------------------------+
*/

//一切皆服务
type IService interface {
	//服务名
	Name() string
	// 开启服务
	Start() error
	// 停止服务
	Stop() error
	// Client is used to call services
	Client() transport.IClient
	// Server is for handling requests and events
	Server() transport.IServer
}

type Service struct {
	Option
}

func (self *Service) Name() string {
	return self.serviceName
}

func (self *Service) Start() error {
	return self.transport.Server().Start()
}

func (self *Service) Stop() error {
	return self.transport.Server().Stop()
}

func (self *Service) Server() transport.IServer {
	return self.transport.Server()
}

func (self *Service) Client() transport.IClient {
	return self.transport.Client()
}
