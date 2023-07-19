package grpc

import (
	"reflect"
	"sync"
)

type Option func(o *option)

type option struct {
	address          string
	registerHandlers sync.Map
}

func Address(address string) Option {
	return func(o *option) {
		o.address = address
	}
}

func M(register, handler interface{}) Option {

	if reflect.TypeOf(register).Kind() != reflect.Func {
		panic("register mut be a fun")
	}

	if reflect.TypeOf(handler).Kind() != reflect.Struct && reflect.TypeOf(handler).Kind() != reflect.Ptr {
		panic("handler mut be a struct")
	}

	return func(o *option) {
		o.registerHandlers.Store(handler, register)
	}
}
