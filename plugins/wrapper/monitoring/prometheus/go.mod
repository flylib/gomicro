module github.com/zjllib/go-micro/plugins/wrapper/monitoring/prometheus/v3

go 1.16

require (
	github.com/zjllib/go-micro/plugins/broker/memory/v3 v3.0.0-20210630062103-c13bb07171bc
	github.com/zjllib/go-micro/plugins/registry/memory/v3 v3.0.0-20210630062103-c13bb07171bc
	github.com/zjllib/go-micro v3.5.2-0.20210630062103-c13bb07171bc
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/client_model v0.2.0
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/zjllib/go-micro/plugins/broker/memory/v3 => ../../../../plugins/broker/memory
	github.com/zjllib/go-micro/plugins/registry/memory/v3 => ../../../../plugins/registry/memory
	github.com/zjllib/go-micro => ../../../../../go-micro
)
