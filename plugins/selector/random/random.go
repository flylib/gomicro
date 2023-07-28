package random

import (
	"github.com/zjllib/go-micro"
	"sync"
)

type randomSelector struct {
	sync.Map
}

func (self *randomSelector) Remove(node micro.Node) {
	self.Delete(node.Name)
}

func (self *randomSelector) Init(nodes []*micro.Node) {
	for _, node := range nodes {
		self.Store(node.Name, node)
	}
}

func (self *randomSelector) Next() *micro.Node {
	var node *micro.Node
	self.Range(func(key, value interface{}) bool {
		node, _ = value.(*micro.Node)
		return false
	})
	return node
}
