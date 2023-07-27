package micro

import "time"

type OptionFun func(o *Option)
type CallOptionFun func(o *CallCallOption)

//options
type (
	Option struct {
		//service name
		Name string
		//Version
		Version string
		//Registered address, used to distinguish container ip and LAN ip
		RegistryAddress string

		IRegistry
		ITransport
	}

	CallCallOption struct {
		ServiceName string
		// Number of Call attempts
		Retries int
		// Request/Response timeout
		RequestTimeout time.Duration
	}
)

func Name(name string) OptionFun {
	return func(o *Option) {
		o.Name = name
	}
}

func Version(version string) OptionFun {
	return func(o *Option) {
		o.Version = version
	}
}

func RegistryAddress(address string) OptionFun {
	return func(o *Option) {
		o.RegistryAddress = address
	}
}

func Transport(transport ITransport) OptionFun {
	return func(o *Option) {
		o.ITransport = transport
	}
}

func Registry(registry IRegistry) OptionFun {
	return func(o *Option) {
		o.IRegistry = registry
	}
}
