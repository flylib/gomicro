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

func NewService(opts ...OptionFun) Service {
	opt := Option{}
	for _, f := range opts {
		f(&opt)
	}

	if opt.transport != nil {
		opt.transport.Init(opts...)
	}
	return Service{
		opt,
	}
}

type Service struct {
	Option
}

func (self *Service) Name() string {
	return self.ServiceName
}

func (self *Service) Start() error {
	return self.transport.Server().Start()
}

func (self *Service) Stop() error {
	return self.transport.Server().Stop()
}
