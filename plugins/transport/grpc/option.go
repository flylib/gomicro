package grpc

type OptionFun func(o *Option)

//options
type Option struct {
	//service name
	serviceName string
	//version
	version string
	//
	address string
}

func Name(name string) OptionFun {
	return func(o *Option) {
		o.serviceName = name
	}
}

func Version(version string) OptionFun {
	return func(o *Option) {
		o.version = version
	}
}

func Address(addr string) OptionFun {
	return func(o *Option) {
		o.address = addr
	}
}
