// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payments_errors.proto

package business_errors

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PaymentsError int32

const (
	// 初始值, 无意义
	PaymentsError_INIT PaymentsError = 0
	// 支付方式查询失败
	PaymentsError_PAYMENT_METHODS_NOT_FOUND PaymentsError = 60300
	// 修改支付方式失败
	PaymentsError_PAYMENT_METHODS_UPDATE PaymentsError = 60302
)

var PaymentsError_name = map[int32]string{
	0:     "INIT",
	60300: "PAYMENT_METHODS_NOT_FOUND",
	60302: "PAYMENT_METHODS_UPDATE",
}
var PaymentsError_value = map[string]int32{
	"INIT":                      0,
	"PAYMENT_METHODS_NOT_FOUND": 60300,
	"PAYMENT_METHODS_UPDATE":    60302,
}

func (x PaymentsError) String() string {
	return proto.EnumName(PaymentsError_name, int32(x))
}
func (PaymentsError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payments_errors_5df25e1b88e0d1e6, []int{0}
}

func init() {
	proto.RegisterEnum("payments.PaymentsError", PaymentsError_name, PaymentsError_value)
}

func init() {
	proto.RegisterFile("payments_errors.proto", fileDescriptor_payments_errors_5df25e1b88e0d1e6)
}

var fileDescriptor_payments_errors_5df25e1b88e0d1e6 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x48, 0xac, 0xcc,
	0x4d, 0xcd, 0x2b, 0x29, 0x8e, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x80, 0x09, 0x6b, 0x45, 0x70, 0xf1, 0x06, 0x40, 0xd9, 0xae, 0x20, 0x15, 0x42,
	0x1c, 0x5c, 0x2c, 0x9e, 0x7e, 0x9e, 0x21, 0x02, 0x0c, 0x42, 0xf2, 0x5c, 0x92, 0x01, 0x8e, 0x91,
	0xbe, 0xae, 0x7e, 0x21, 0xf1, 0xbe, 0xae, 0x21, 0x1e, 0xfe, 0x2e, 0xc1, 0xf1, 0x7e, 0xfe, 0x21,
	0xf1, 0x6e, 0xfe, 0xa1, 0x7e, 0x2e, 0x02, 0x3d, 0xd7, 0x99, 0x85, 0x64, 0xb8, 0xc4, 0xd0, 0x15,
	0x84, 0x06, 0xb8, 0x38, 0x86, 0xb8, 0x0a, 0xf4, 0x5d, 0x67, 0x76, 0xd2, 0x8a, 0xe2, 0x4f, 0x2a,
	0x2d, 0xce, 0xcc, 0x4b, 0x2d, 0x86, 0x59, 0x7e, 0x8a, 0x49, 0x2c, 0xb4, 0x24, 0xdf, 0xb5, 0x22,
	0x39, 0xb5, 0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x38, 0x06, 0x66, 0x73, 0x12, 0x1b, 0xd8, 0x59, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xcb, 0xf5, 0x23, 0xaf, 0x00, 0x00, 0x00,
}
