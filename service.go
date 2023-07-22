package micro

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zjllib/goutils/net"
	"log"
	"strings"
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
		opt,
	}
}

type Service struct {
	Option
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
	return self.ITransport.Server().Start()
}

func (self *Service) Stop() error {
	return self.ITransport.Server().Stop()
}
