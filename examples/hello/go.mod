module hello

go 1.17

require (
	github.com/golang/protobuf v1.5.3
	github.com/zjllib/go-micro v0.0.0-20230717035846-7d40414eaab8
	github.com/zjllib/go-micro/plugins/registry/etcd v0.0.0-20230717035846-7d40414eaab8
	github.com/zjllib/go-micro/plugins/transport/grpc v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.9.0
	google.golang.org/grpc v1.56.2
)

require (
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/zjllib/goutils v1.0.15-0.20230717032249-1bb6c05812a7 // indirect
	go.etcd.io/etcd/api/v3 v3.5.9 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.9 // indirect
	go.etcd.io/etcd/client/v3 v3.5.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace (
	github.com/zjllib/go-micro => ../../../go-micro
	github.com/zjllib/go-micro/plugins/registry/etcd => ../../../go-micro/plugins/registry/etcd
	github.com/zjllib/go-micro/plugins/transport/grpc => ../../../go-micro/plugins/transport/grpc
)
