package micro

import "time"

type OptionFun func(o *Option)

//options
type (
	Option struct {
		//service name
		Name string
		//Version
		Version string

		TransportOption
		RegistryOption
	}
	// RegistryOption Registry
	RegistryOption struct {
		IRegistry
		Address []string
		// The register expiry time
		RegisterTTL time.Duration
		// The interval on which to register
		RegisterInterval time.Duration
	}

	// TransportOption Transport
	TransportOption struct {
		ITransport
		Address string
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

func Address(addr string) OptionFun {
	return func(o *Option) {
		o.TransportOption.Address = addr
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

// Register the service with a TTL
func RegisterTTL(t time.Duration) OptionFun {
	return func(o *Option) {
		o.RegisterTTL = t
	}
}

// Register the service with at interval
func RegisterInterval(t time.Duration) OptionFun {
	return func(o *Option) {
		o.RegisterInterval = t
	}
}
