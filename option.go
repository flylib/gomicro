package micro

type OptionFun func(o *Option)

//options
type Option struct {
	//service name
	ServiceName string
	//Version
	Version string
	//
	Address string

	transport ITransport
}

func Name(name string) OptionFun {
	return func(o *Option) {
		o.ServiceName = name
	}
}

func Version(version string) OptionFun {
	return func(o *Option) {
		o.Version = version
	}
}

func Address(addr string) OptionFun {
	return func(o *Option) {
		o.Address = addr
	}
}

func Transport(t ITransport) OptionFun {
	return func(o *Option) {
		o.transport = t
	}
}
