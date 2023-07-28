package micro

import "time"

type OptionFun func(o *Option)
type CallOptionFun func(o *CallOption)

//options
type Option struct {
	//service name
	Name string
	//Version
	Version string
	//Registered address, used to distinguish container ip and LAN ip
	RegistryAddress string

	IRegistry
	ITransport
}

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

type CallOption struct {
	serviceName string
	// Number of Call attempts
	retries int
	// Request/Response timeout
	requestTimeout time.Duration

	ISelector
}

func CallTarget(serviceName string) CallOptionFun {
	return func(o *CallOption) {
		o.serviceName = serviceName
	}
}

func CallRetries(retries int) CallOptionFun {
	return func(o *CallOption) {
		o.retries = retries
	}
}

func CallTimeout(requestTimeout time.Duration) CallOptionFun {
	return func(o *CallOption) {
		o.requestTimeout = requestTimeout
	}
}
