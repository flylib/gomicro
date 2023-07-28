package micro

type ISelector interface {
	Init(nodes []*Node)
	Next() *Node
	Remove(Node)
}

type defaultSelector struct {
	node *Node
}

func (self *defaultSelector) Init(nodes []*Node) {
	if len(nodes) == 0 {
		panic("The length of nodes must be greater than 0")
	}
	self.node = nodes[0]
}

func (self *defaultSelector) Next() *Node {
	return self.node
}

func (self *defaultSelector) Remove(node Node) {
	if node.Name == self.node.Name {
		self.node = nil
	}
}
