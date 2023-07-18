package grpc

type Option func(o *option)

//type Register struct {
//	registerFun interface{}
//	handler     interface{}
//}

type option struct {
	address          string
	registerHandlers map[interface{}]interface{}
}

func Address(address string) Option {
	return func(o *option) {
		o.address = address
	}
}

func M(registerFun, handler interface{}) Option {
	return func(o *option) {
		o.registerHandlers[handler] = registerFun
	}
}
