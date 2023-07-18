module github.com/zjllib/go-micro/plugins/transport/grpc

go 1.17

require (
    github.com/zjllib/go-micro v0.0.0-20230717035846-7d40414eaab8
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.56.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace (
	github.com/zjllib/go-micro  => ../../../../go-micro
)
