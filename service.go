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
+		registry,transport(udp、tcp、ws、quic、rpc)、codec+
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

func (self *Service) NewClient(opts ...CallOptionFun) (*Client, error) {
	var opt CallOption
	for i := 0; i < len(opts); i++ {
		opts[i](&opt)
	}

	//must have a node selector
	if opt.ISelector == nil {
		opt.ISelector = &defaultSelector{}
	}

	cli := &Client{
		servicePath: RegistryPrefix + opt.serviceName,
		Option:      self.Option,
		CallOption:  opt,
	}
	return cli, cli.connectToServices()
}

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
	node := self.Next()
	if node == nil {
		return errors.New("not found the node")
	}
	return node.cli.Call(ctx, method, in, out)
}
