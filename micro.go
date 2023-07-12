package micro

func NewService(opts ...OptionFun) Service {
	opt := Option{}
	for _, f := range opts {
		f(&opt)
	}

	return Service{
		opt,
	}
}
