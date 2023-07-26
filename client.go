package micro

import "context"

type Client struct {
	clients []IClient
}

func (c Client) Call(ctx context.Context, method string, in, out interface{}) error {
	return c.clients[0].Call(ctx, method, in, out)
}

func (c Client) Init() {

}

func (c Client) watchNodes() {

}
