package micro

type IRegistry interface {
	Register(*Service) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
}
