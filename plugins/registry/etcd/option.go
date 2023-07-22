package etcd

import "time"

type Option func(o *option)

type option struct {
	endpoints []string
	// The register expiry time
	registerTTL time.Duration
	// The interval on which to register
	registerInterval time.Duration
	// dialTimeout is the timeout for failing to establish a connection.
	dialTimeout time.Duration
}

func Endpoints(endpoints ...string) Option {
	return func(o *option) {
		o.endpoints = endpoints
	}
}

func RegisterTTL(t time.Duration) Option {
	return func(o *option) {
		o.registerTTL = t
	}
}

func RegisterInterval(t time.Duration) Option {
	return func(o *option) {
		o.registerInterval = t
	}
}

func DialTimeout(t time.Duration) Option {
	return func(o *option) {
		o.dialTimeout = t
	}
}
