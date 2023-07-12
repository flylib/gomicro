package etcd

import (
	. "github.com/zjllib/go-micro"
)

type etcd struct {

}

func NewRegistry(opts ...OptionFun) Registry {

}

func (e *etcd) Register(service *Service) error {
	panic("implement me")
}

func (e *etcd) Deregister(service *Service) error {
	panic("implement me")
}

func (e *etcd) GetService(s string) ([]*Service, error) {
	panic("implement me")
}

func (e *etcd) ListServices() ([]*Service, error) {
	panic("implement me")
}

