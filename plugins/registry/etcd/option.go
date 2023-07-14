package etcd

import "time"

type Option func(o *option)

type option struct {
	address []string
	// The register expiry time
	registerttl time.Duration
	// The interval on which to register
	registerinterval time.Duration
}

func Address(address ...string) Option {
	return func(o *option) {
		o.address = address
	}
}

// Register the service with a TTL
func RegisterTTL(t time.Duration) Option {
	return func(o *option) {
		o.registerttl = t
	}
}

// Register the service with at interval
func RegisterInterval(t time.Duration) Option {
	return func(o *option) {
		o.registerinterval = t
	}
}
