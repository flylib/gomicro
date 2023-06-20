package transport

import "net"

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
