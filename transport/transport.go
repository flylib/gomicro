package transport

type ITransport interface {
	Server() IServer
	Client() IClient
}
