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

	//if opt.Transport.ITransport != nil {
	//	opt.Transport.Init(opts...)
	//}
	return Service{
		opt,
	}
}

type Service struct {
	Option
}

func (self *Service) Name() string {
	return self.Option.Name
}

func (self *Service) Start() error {
	err := self.IRegistry.Register(self)
	if err != nil {
		return err
	}
	return self.ITransport.Server().Start()
}

func (self *Service) Stop() error {

	return self.ITransport.Server().Stop()
}
