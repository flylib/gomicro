package grpc

type Option func(o *option)

type option struct {
	addres string
}

func Addres(addres string) Option {
	return func(o *option) {
		o.addres = addres
	}
}
