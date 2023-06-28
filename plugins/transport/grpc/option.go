package grpc

type OptionFun func(o *Option)

//options
type Option struct {
	//service name
	serviceName string
	//
	address string
}

func Name(name string) OptionFun {
	return func(o *Option) {
		o.serviceName = name
	}
}

func Address(addr string) OptionFun {
	return func(o *Option) {
		o.address = addr
	}
}
