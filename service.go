package micro

import (
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

	cli := &Client{
		servicePath: RegistryPrefix + opt.serviceName,
		Option:      self.Option,
		CallOption:  opt,
	}
	return cli, cli.connectToServices()
}
