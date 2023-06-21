package transport

//客户端
type IClient interface {
	Dial() error
}
