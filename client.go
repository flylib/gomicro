package micro

import (
	"context"
	"errors"
	"log"
	"sync"
)

type Client struct {
	Option
	CallOption
	servicePath string
	nodes       []*Node
	sync.Once
}

func (self *Client) connectToServices() error {
	var nodes []*Node
	nodes, err := self.IRegistry.GetNodes(self.servicePath)
	if err != nil {
		return err
	}
	//connect all service
	for i := 0; i < len(nodes); i++ {
		cli := self.ITransport.Client()
		err = cli.DialNode(*nodes[i])
		if err != nil {
			return err
		}
		nodes[i].cli = cli
	}
	self.nodes = nodes
	//must have a node selector
	if self.ISelector == nil || len(nodes) <= 1 {
		self.ISelector = &singleSelector{}
	}
	self.ISelector.Init(nodes)
	self.watchServices()
	return nil
}

func (self *Client) watchServices() {
	self.Once.Do(func() {
		go func() {
			err := self.IRegistry.WatchNodes(self.servicePath, func(eventType EventType, node Node) {
				switch eventType {
				case Delete:
					for i := 0; i < len(self.nodes); i++ {
						if self.nodes[i].Name == node.Name {
							self.nodes[i].cli.Close()
							self.nodes = append(self.nodes[:i], self.nodes[i+1:]...)
							self.ISelector.Remove(node)
						}
					}
				case Modify:

				}
			})
			if err != nil {
				log.Println(err)
			}
		}()
	})
}

func (self *Client) Call(ctx context.Context, method string, in, out interface{}) (err error) {
	node := self.ISelector.Next()
	if node == nil {
		return errors.New("not found the node")
	}
	return node.cli.Call(ctx, method, in, out)
}
