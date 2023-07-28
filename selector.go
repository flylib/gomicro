package micro

type ISelector interface {
	Init(nodes []*Node)
	Next() *Node
	Remove(Node)
}

type singleSelector struct {
	node *Node
}

func (self *singleSelector) Init(nodes []*Node) {
	if len(nodes) == 0 {
		panic("The length of nodes must be greater than 0")
	}
	self.node = nodes[0]
}

func (self *singleSelector) Next() *Node {
	return self.node
}

func (self *singleSelector) Remove(node Node) {
	if node.Name == self.node.Name {
		self.node = nil
	}
}
