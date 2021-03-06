// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/ddgadmin.proto

package ddgadmin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ValidateRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Application          string   `protobuf:"bytes,3,opt,name=application,proto3" json:"application,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3cd3dcdd80ebcc8, []int{0}
}
func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (dst *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(dst, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *ValidateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ValidateRequest) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

type ValidateResponse struct {
	Success              int32    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Uid                  int64    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateResponse) Reset()         { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()    {}
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3cd3dcdd80ebcc8, []int{1}
}
func (m *ValidateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateResponse.Unmarshal(m, b)
}
func (m *ValidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateResponse.Marshal(b, m, deterministic)
}
func (dst *ValidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateResponse.Merge(dst, src)
}
func (m *ValidateResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateResponse.Size(m)
}
func (m *ValidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateResponse proto.InternalMessageInfo

func (m *ValidateResponse) GetSuccess() int32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *ValidateResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ValidateResponse) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidateRequest)(nil), "ddgadmin.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "ddgadmin.ValidateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DdgAdminFrontClient is the client API for DdgAdminFront service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DdgAdminFrontClient interface {
	ValidateUserAccountAndPwd(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type ddgAdminFrontClient struct {
	cc *grpc.ClientConn
}

func NewDdgAdminFrontClient(cc *grpc.ClientConn) DdgAdminFrontClient {
	return &ddgAdminFrontClient{cc}
}

func (c *ddgAdminFrontClient) ValidateUserAccountAndPwd(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/ddgadmin.ddgAdminFront/ValidateUserAccountAndPwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DdgAdminFrontServer is the server API for DdgAdminFront service.
type DdgAdminFrontServer interface {
	ValidateUserAccountAndPwd(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

func RegisterDdgAdminFrontServer(s *grpc.Server, srv DdgAdminFrontServer) {
	s.RegisterService(&_DdgAdminFront_serviceDesc, srv)
}

func _DdgAdminFront_ValidateUserAccountAndPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DdgAdminFrontServer).ValidateUserAccountAndPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddgadmin.ddgAdminFront/ValidateUserAccountAndPwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DdgAdminFrontServer).ValidateUserAccountAndPwd(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DdgAdminFront_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddgadmin.ddgAdminFront",
	HandlerType: (*DdgAdminFrontServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateUserAccountAndPwd",
			Handler:    _DdgAdminFront_ValidateUserAccountAndPwd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/ddgadmin.proto",
}

func init() { proto.RegisterFile("protobuf/ddgadmin.proto", fileDescriptor_f3cd3dcdd80ebcc8) }

var fileDescriptor_f3cd3dcdd80ebcc8 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xad, 0x8b, 0x5a, 0x47, 0xc4, 0x92, 0x8b, 0xe9, 0x9e, 0xca, 0x9e, 0x3c, 0x55, 0xd0,
	0x5f, 0xb0, 0x17, 0xcf, 0x25, 0xa8, 0xf7, 0x34, 0x13, 0x97, 0x40, 0x37, 0x89, 0x99, 0x84, 0xfd,
	0xfb, 0x92, 0xac, 0x59, 0x45, 0xbc, 0xcd, 0x7b, 0x0f, 0xe6, 0x9b, 0x79, 0x70, 0xef, 0x83, 0x8b,
	0xee, 0x98, 0x3e, 0x1e, 0x11, 0x07, 0x89, 0xa3, 0xb1, 0xfb, 0xe2, 0xb0, 0x75, 0xd5, 0x9d, 0x81,
	0xbb, 0x77, 0x79, 0x32, 0x28, 0xa3, 0x16, 0xfa, 0x33, 0x69, 0x8a, 0x8c, 0xc3, 0x95, 0x54, 0xca,
	0x25, 0x1b, 0xf9, 0x6a, 0xb7, 0x7a, 0xb8, 0x16, 0x55, 0xb2, 0x16, 0xd6, 0x5e, 0x12, 0x4d, 0x2e,
	0x20, 0x3f, 0x2f, 0xd1, 0xa2, 0xd9, 0x0e, 0x6e, 0xa4, 0xf7, 0x27, 0xa3, 0x64, 0x34, 0xce, 0xf2,
	0xa6, 0xc4, 0xbf, 0xad, 0xee, 0x00, 0x9b, 0x1f, 0x14, 0x79, 0x67, 0x49, 0x67, 0x16, 0x25, 0xa5,
	0x34, 0x51, 0x61, 0x5d, 0x88, 0x2a, 0xd9, 0x06, 0x9a, 0x91, 0x86, 0x6f, 0x4c, 0x1e, 0xb3, 0x93,
	0x0c, 0x96, 0xcd, 0x8d, 0xc8, 0xe3, 0x93, 0x86, 0x5b, 0xc4, 0xa1, 0xcf, 0x8f, 0xbc, 0x04, 0x67,
	0x23, 0x7b, 0x85, 0x6d, 0x45, 0xbc, 0x91, 0x0e, 0xfd, 0x7c, 0x77, 0x6f, 0xf1, 0x30, 0x21, 0xdb,
	0xee, 0x97, 0x16, 0xfe, 0xbc, 0xdc, 0xb6, 0xff, 0x45, 0xf3, 0x89, 0xdd, 0xd9, 0xf1, 0xb2, 0x94,
	0xf6, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x32, 0xb3, 0x4f, 0x81, 0x4f, 0x01, 0x00, 0x00,
}
