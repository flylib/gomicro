// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package proto is a generated protocol buffer package.

https://blog.csdn.net/u014308482/article/details/52958148 Protobuf语言指南——.proto文件语法详解
然后将proto文件编译为go文件 protoc --go_out=plugins=grpc:./proto/ ./test.proto
定义包名

It is generated from these files:
	test.proto

It has these top-level messages:
	Req
	Res
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// 定义 Req 消息结构
type Req struct {
	// 类型 字段 = 标识号
	JsonStr string `protobuf:"bytes,1,opt,name=jsonStr" json:"jsonStr,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto1.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetJsonStr() string {
	if m != nil {
		return m.JsonStr
	}
	return ""
}

// 定义 Res 消息结构
type Res struct {
	BackJson string `protobuf:"bytes,1,opt,name=backJson" json:"backJson,omitempty"`
}

func (m *Res) Reset()                    { *m = Res{} }
func (m *Res) String() string            { return proto1.CompactTextString(m) }
func (*Res) ProtoMessage()               {}
func (*Res) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Res) GetBackJson() string {
	if m != nil {
		return m.BackJson
	}
	return ""
}

func init() {
	proto1.RegisterType((*Req)(nil), "proto.Req")
	proto1.RegisterType((*Res)(nil), "proto.Res")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Waiter service

type WaiterClient interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	DoMD5(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type waiterClient struct {
	cc *grpc.ClientConn
}

func NewWaiterClient(cc *grpc.ClientConn) WaiterClient {
	return &waiterClient{cc}
}

func (c *waiterClient) DoMD5(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := grpc.Invoke(ctx, "/proto.Waiter/DoMD5", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Waiter service

type WaiterServer interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	DoMD5(context.Context, *Req) (*Res, error)
}

func RegisterWaiterServer(s *grpc.Server, srv WaiterServer) {
	s.RegisterService(&_Waiter_serviceDesc, srv)
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
		FullMethod: "/proto.Waiter/DoMD5",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaiterServer).DoMD5(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Waiter_serviceDesc = grpc.ServiceDesc{
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

func init() { proto1.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xf2, 0x5c, 0xcc, 0x41, 0xa9,
	0x85, 0x42, 0x12, 0x5c, 0xec, 0x59, 0xc5, 0xf9, 0x79, 0xc1, 0x25, 0x45, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92, 0x22, 0x48, 0x41, 0xb1, 0x90, 0x14, 0x17, 0x47, 0x52, 0x62,
	0x72, 0xb6, 0x57, 0x71, 0x7e, 0x1e, 0x54, 0x05, 0x9c, 0x6f, 0xa4, 0xcd, 0xc5, 0x16, 0x9e, 0x98,
	0x59, 0x92, 0x5a, 0x24, 0xa4, 0xc8, 0xc5, 0xea, 0x92, 0xef, 0xeb, 0x62, 0x2a, 0xc4, 0x05, 0xb1,
	0x45, 0x2f, 0x28, 0xb5, 0x50, 0x0a, 0xc1, 0x2e, 0x56, 0x62, 0x48, 0x62, 0x03, 0x73, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x09, 0x2f, 0x60, 0x8c, 0x00, 0x00, 0x00,
}