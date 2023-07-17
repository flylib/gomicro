package micro

type IRegistry interface {
	Register(*Service) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
}

//Node 节点信息
type Node struct {
	Name     string                 `json:"name"`
	Version  string                 `json:"version"`
	Address  string                 `json:"address"`
	Metadata map[string]interface{} `json:"metadata"`
}

const RegistryPrefix = "micro/nodes/"
