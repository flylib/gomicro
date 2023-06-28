module hello

go 1.17


require (
    github.com/zjllib/go-micro latest
)

replace (
	github.com/zjllib/go-micro  => ../../../go-micro
)

