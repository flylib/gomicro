package micro

type OptionFun func(o *Option)

//options
type (
	Option struct {
		//service name
		Name string
		//Version
		Version string

		IRegistry
		ITransport
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
