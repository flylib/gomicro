package micro

import (
	"context"
	"net"
)

/*
+---------------------------------------------------+
+		server		|	transport(tcp、udp、grpc、http)	|   client	    +
+---------------------------------------------------+
*/

//会话
type ISession interface {
	//ID
	ID() uint64
	//断开
	Close() error
	//发送消息
	Send(msg interface{}) error
	//设置键值对，存储关联数据
	Store(key, value interface{})
	//获取键值对
	Load(key interface{}) (value interface{}, ok bool)
	//地址
	RemoteAddr() net.Addr
}

//客户端
type IClient interface {
	Dial() error
	Call(ctx context.Context, method string, in, out interface{}) error
}

//服务端
type IServer interface {
	// 启动监听
	Start() error
	// 停止服务
	Stop() error
	// 地址
	Addr() string
}

type ITransport interface {
	//Init(...OptionFun) error
	Server() IServer
	Client() IClient
	Type() string
}
