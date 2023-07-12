package micro

/*
+---------------------------------------------------+
+				     service						+
+---------------------------------------------------+
+		server			|		client				+
+---------------------------------------------------+
+		registry,bee worker、conn pool、codec		+
+---------------------------------------------------+
+		transport(udp、tcp、ws、quic、rpc)			+
+---------------------------------------------------+
*/

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
