package micro

import "log"

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

func init() {
	log.SetPrefix("[micro]")
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
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
