package grpc

import (
	"google.golang.org/grpc"
)

//客户端
//type IClient interface {
//	Dial() error
//}

type client struct {
	opt  option
	conn *grpc.ClientConn
}

func (self client) Dial() error {
	dial, err := grpc.Dial(self.opt.address)
	if err != nil {
		return err
	}
	self.conn = dial
	return nil
}
