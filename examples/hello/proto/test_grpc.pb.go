// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.1
// source: test.proto

// 定义包名

package proto

import (
	context "context"
	go_micro "github.com/zjllib/go-micro"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Waiter_DoMD5_FullMethodName = "/proto.Waiter/DoMD5"
)

// WaiterClient is the client API for Waiter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WaiterClient interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	DoMD5(ctx context.Context, in *Req) (*Res, error)
}

type waiterClient struct {
	cc *go_micro.Client
}

func NewWaiterClient(cc *go_micro.Client) WaiterClient {
	return &waiterClient{cc}
}

func (c *waiterClient) DoMD5(ctx context.Context, in *Req) (*Res, error) {
	out := new(Res)
	err := c.cc.Call(ctx, Waiter_DoMD5_FullMethodName, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaiterServer is the server API for Waiter service.
// All implementations should embed UnimplementedWaiterServer
// for forward compatibility
type WaiterServer interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	DoMD5(context.Context, *Req) (*Res, error)
}

// UnimplementedWaiterServer should be embedded to have forward compatible implementations.
type UnimplementedWaiterServer struct {
}

func (UnimplementedWaiterServer) DoMD5(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoMD5 not implemented")
}

// UnsafeWaiterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WaiterServer will
// result in compilation errors.
type UnsafeWaiterServer interface {
	mustEmbedUnimplementedWaiterServer()
}

func RegisterWaiterServer(s grpc.ServiceRegistrar, srv WaiterServer) {
	s.RegisterService(&Waiter_ServiceDesc, srv)
}

func _Waiter_DoMD5_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaiterServer).DoMD5(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Waiter_DoMD5_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaiterServer).DoMD5(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Waiter_ServiceDesc is the grpc.ServiceDesc for Waiter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Waiter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Waiter",
	HandlerType: (*WaiterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoMD5",
			Handler:    _Waiter_DoMD5_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
