package micro

type IRegistry interface {
	RegisterNode(node Node) error
	DeregisterNode(node Node) error
	GetNodes(string) ([]*Node, error)
	WatchNodes(path string, callback func(eventType EventType, node Node)) error
}

type EventType string

const (
	Delete EventType = "DELETE"
	Modify EventType = "MODIFY"
)

//Node 节点信息
type Node struct {
	Name     string                 `json:"name"`
	Version  string                 `json:"version"`
	Address  string                 `json:"address"`
	Type     string                 `json:"type"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	cli IClient
}

const RegistryPrefix = "micro/nodes/"
