package etcd

import (
	"context"
	"fmt"
	"github.com/zjllib/go-micro"
	"go.etcd.io/etcd/client/v3"
	"time"
)

type etcd struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	canclefunc    func()
}

func NewRegistry(opts ...Option) micro.IRegistry {
	var options option
	for _, o := range opts {
		o(&options)
	}

	conf := clientv3.Config{
		Endpoints:   options.address,
		DialTimeout: 5 * time.Second,
	}

	c, err := clientv3.New(conf)
	if err != nil {
		panic(err)
	}
	registry := etcd{
		client: c,
	}
	registry.setLease(int64(options.registerttl.Seconds()))
	go registry.ListenLeaseRespChan()
	return &registry
}

//设置租约
func (self *etcd) setLease(ttl int64) error {
	//create a ne lease
	lease := clientv3.NewLease(self.client)

	//set
	leaseResp, err := lease.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}

	//设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)

	if err != nil {
		return err
	}

	self.lease = lease
	self.leaseResp = leaseResp
	self.canclefunc = cancelFunc
	self.keepAliveChan = leaseRespChan
	return nil
}

//监听 续租情况
func (self *etcd) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-self.keepAliveChan:
			if leaseKeepResp == nil {
				fmt.Printf("已经关闭续租功能\n")
				return
			} else {
				fmt.Printf("续租成功\n")
			}
		}
	}
}

func (self *etcd) Register(service *micro.Service) error {
	kv := clientv3.NewKV(self.client)
	_, err := kv.Put(context.TODO(), service.Option.Name, service.ITransport.Server().Addr(), clientv3.WithLease(self.leaseResp.ID))
	return err
}

func (self *etcd) Deregister(service *micro.Service) error {
	kv := clientv3.NewKV(self.client)
	response, err := kv.Delete(context.Background(), service.Option.Name)
	fmt.Println(*response)
	return err
}

func (self *etcd) GetService(s string) ([]*micro.Service, error) {
	panic("implement me")
}

func (self *etcd) ListServices() ([]*micro.Service, error) {
	panic("implement me")
}
