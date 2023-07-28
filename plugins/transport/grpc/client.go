package grpc

import (
	"context"
	"github.com/zjllib/go-micro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	opt  option
	conn *grpc.ClientConn
	node micro.Node
}

func (self *client) Dial() error {
	conn, err := grpc.Dial(self.opt.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	self.conn = conn
	return nil
}

func (self *client) DialNode(node micro.Node) error {
	conn, err := grpc.Dial(node.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	self.conn = conn
	return nil
}

func (self *client) Close() error {
	return self.conn.Close()
}

func (self *client) Call(ctx context.Context, method string, in, out interface{}) error {
	return self.conn.Invoke(ctx, method, in, out)
}
