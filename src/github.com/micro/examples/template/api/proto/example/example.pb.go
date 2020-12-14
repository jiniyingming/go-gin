// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/examples/template/api/proto/example/example.proto

/*
Package go_micro_api_template is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/template/api/proto/example/example.proto

It has these top-level messages:
*/
package go_micro_api_template

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_api "github.com/micro/go-api/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Example service

type ExampleClient interface {
	Call(ctx context.Context, in *go_api.Request, opts ...grpc.CallOption) (*go_api.Response, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Call(ctx context.Context, in *go_api.Request, opts ...grpc.CallOption) (*go_api.Response, error) {
	out := new(go_api.Response)
	err := grpc.Invoke(ctx, "/go.micro.api.template.Example/Call", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleServer interface {
	Call(context.Context, *go_api.Request) (*go_api.Response, error)
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(go_api.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.api.template.Example/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Call(ctx, req.(*go_api.Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.api.template.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Example_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/examples/template/api/proto/example/example.proto",
}

func init() {
	proto.RegisterFile("github.com/micro/examples/template/api/proto/example/example.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0xad, 0x48,
	0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x4f,
	0x2c, 0xc8, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0x81, 0xcb, 0xc1, 0x68, 0x3d, 0xb0, 0xa8, 0x90, 0x68,
	0x7a, 0xbe, 0x1e, 0x58, 0xaf, 0x5e, 0x62, 0x41, 0xa6, 0x1e, 0x4c, 0x9b, 0x94, 0x3a, 0x86, 0xd1,
	0xe9, 0xf9, 0xba, 0x08, 0xb3, 0x40, 0x6a, 0xc1, 0x2c, 0x23, 0x33, 0x2e, 0x76, 0x57, 0x88, 0x81,
	0x42, 0xda, 0x5c, 0x2c, 0xce, 0x89, 0x39, 0x39, 0x42, 0xfc, 0x7a, 0xe9, 0x10, 0xd3, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0x04, 0x10, 0x02, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a,
	0x0c, 0x49, 0x6c, 0x60, 0xed, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x63, 0xb0, 0x5b,
	0xc4, 0x00, 0x00, 0x00,
}
