package transport

//服务端
type IServer interface {
	// 启动监听
	Listen() error
	// 停止服务
	Stop() error
	// 地址
	Addr() string
}
