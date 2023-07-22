package grpc

import (
	"context"
	"google.golang.org/grpc"
)

type client struct {
	opt  option
	conn *grpc.ClientConn
}

func (self *client) Dial() error {
	dial, err := grpc.Dial(self.opt.address)
	if err != nil {
		return err
	}
	self.conn = dial
	return nil
}

func (self *client) Call(ctx context.Context, method string, in, out interface{}) error {
	return self.conn.Invoke(ctx, method, in, out)
}
