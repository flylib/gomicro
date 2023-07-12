package micro

type Registry interface {
	Init(...Option) error
	Register(*Service) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
}
