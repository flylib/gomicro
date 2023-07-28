package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zjllib/go-micro"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"log"
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

	if options.dialTimeout.Seconds() == 0 {
		options.dialTimeout = 5 * time.Second
	}

	conf := clientv3.Config{
		Endpoints:   options.endpoints,
		DialTimeout: options.dialTimeout,
	}

	c, err := clientv3.New(conf)
	if err != nil {
		panic(err)
	}
	registry := etcd{
		client: c,
	}
	if options.registerTTL.Seconds() != 0 {
		err = registry.setLease(int64(options.registerTTL.Seconds()))
		if err != nil {
			panic(err)
		}
		go registry.listenLeaseRespChan()
	}

	return &registry
}

func (self *etcd) RegisterNode(node micro.Node) error {
	marshal, err := json.Marshal(node)
	if err != nil {
		return err
	}
	log.Println(node.Name, " registry node info:\n", string(marshal))
	_, err = self.client.Put(context.TODO(), node.Name, string(marshal), clientv3.WithLease(self.leaseResp.ID))
	return err
}

func (self *etcd) DeregisterNode(node micro.Node) error {
	_, err := self.client.Delete(context.Background(), node.Name)
	return err
}

func (self *etcd) GetNodes(prefix string) ([]micro.Node, error) {
	result, err := self.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var nodes []micro.Node
	for _, item := range result.Kvs {
		var node micro.Node
		err = json.Unmarshal(item.Value, &node)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

//watcher 监听前缀
func (s *etcd) WatchNodes(path string, callback func(eventType micro.EventType, node micro.Node)) error {
	watchChan := s.client.Watch(context.Background(), path, clientv3.WithPrefix())
	for resp := range watchChan {
		if resp.Canceled {
			return resp.Err()
		}
		for _, e := range resp.Events {
			var node micro.Node
			switch e.Type {
			case mvccpb.PUT: //修改或者新增
				json.Unmarshal(e.Kv.Value, &node)
				callback(micro.Modify, node)
			case mvccpb.DELETE: //删除
				node.Name = string(e.Kv.Key)
				callback(micro.Delete, node)
			}
		}
	}
	return errors.New(fmt.Sprintf("WatchChan-%s  close", path))
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
func (self *etcd) listenLeaseRespChan() {
	for rsp := range self.keepAliveChan {
		log.Println("续租成功", rsp.ID)
	}
	log.Println("已经关闭续租功能")
	return
}
