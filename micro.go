package micro

func NewService(opts ...OptionFun) IService {
	opt := Option{}
	for _, f := range opts {
		f(&opt)
	}

	return &Service{
		opt: opt,
	}
}
