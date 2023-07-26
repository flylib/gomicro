package micro

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zjllib/goutils/net"
	"log"
	"strings"
	"sync"
)

/*
+---------------------------------------------------+
+				     service						+
+---------------------------------------------------+
+		server			|		client				+
+---------------------------------------------------+
+		registry,bee worker、conn pool、codec		+
+---------------------------------------------------+
+		transport(udp、tcp、ws、quic、rpc)			+
+---------------------------------------------------+
*/

func init() {
	log.SetPrefix("[micro]")
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
func NewService(opts ...OptionFun) Service {
	opt := Option{}
	for _, f := range opts {
		f(&opt)
	}
	return Service{
		Option: opt,
	}
}

type Service struct {
	Option
	nodes []Node

	sync.Once
}

func (self *Service) Name() string {
	return self.Option.Name
}

func (self *Service) Start() error {
	addr := self.RegistryAddress
	if addr == "" {
		splits := strings.Split(self.ITransport.Server().Addr(), ":")
		if len(splits) != 2 {
			return errors.New("bad addr:" + addr)
		}
		//Get LAN address
		ip, err := net.GetOutboundIP()
		if err != nil {
			return err
		}
		addr = ip.String() + ":" + splits[1]
	}

	node := Node{
		Name:    RegistryPrefix + self.Name() + "-" + uuid.NewV4().String(),
		Version: self.Version,
		Address: addr,
		Type:    self.ITransport.Type(),
	}
	err := self.IRegistry.RegisterNode(node)
	if err != nil {
		return err
	}
	return self.ITransport.Server().Listen()
}

func (self *Service) Stop() error {
	return self.ITransport.Server().Stop()
}

func (self *Service) Call(ctx context.Context, method string, in, out interface{}) (err error) {
	self.Once.Do(func() {
		var nodes []Node
		nodes, err = self.IRegistry.GetNodes(RegistryPrefix + self.Name())
		if err != nil {
			return
		}
		//connect all service
		for i := 0; i < len(nodes); i++ {
			cli := self.ITransport.Client()
			err = cli.DialNode(nodes[i])
			if err != nil {
				return
			}
			nodes[i].clients = append(nodes[i].clients, cli)
		}
		self.nodes = nodes

		go self.IRegistry.WatchNodes(RegistryPrefix+self.Name(), func(eventType EventType, node Node) {
			switch eventType {
			case Delete:
				for i := 0; i < len(self.nodes); i++ {
					if self.nodes[i].Name == node.Name {
						for _, client := range self.nodes[i].clients {
							client.Close()
						}
						self.nodes = append(self.nodes[:i], self.nodes[i+1:]...)
					}
				}
			case Modify:

			}
		})
	})

	return self.nodes[0].clients[0].Call(ctx, method, in, out)
}
