package registry

import . "github.com/zjllib/go-micro"

type IRegistry interface {
	Register(*IService) error
	GetService(string) ([]*IService, error)
}
