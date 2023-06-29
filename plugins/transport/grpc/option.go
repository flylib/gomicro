package grpc

type OptionFun func(o *Option)

//options
type Option struct {
	address string
}

func Address(addr string) OptionFun {
	return func(o *Option) {
		o.address = addr
	}
}
