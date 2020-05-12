// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arith.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	arith.proto

It has these top-level messages:
	ArithRequest
	ArithResponse
*/
package pb

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

// 算术运算请求结构
type ArithRequest struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *ArithRequest) Reset()                    { *m = ArithRequest{} }
func (m *ArithRequest) String() string            { return proto.CompactTextString(m) }
func (*ArithRequest) ProtoMessage()               {}
func (*ArithRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ArithRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *ArithRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

// 算术运算响应结构
type ArithResponse struct {
	Pro int32 `protobuf:"varint,1,opt,name=pro" json:"pro,omitempty"`
	Quo int32 `protobuf:"varint,2,opt,name=quo" json:"quo,omitempty"`
	Rem int32 `protobuf:"varint,3,opt,name=rem" json:"rem,omitempty"`
}

func (m *ArithResponse) Reset()                    { *m = ArithResponse{} }
func (m *ArithResponse) String() string            { return proto.CompactTextString(m) }
func (*ArithResponse) ProtoMessage()               {}
func (*ArithResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ArithResponse) GetPro() int32 {
	if m != nil {
		return m.Pro
	}
	return 0
}

func (m *ArithResponse) GetQuo() int32 {
	if m != nil {
		return m.Quo
	}
	return 0
}

func (m *ArithResponse) GetRem() int32 {
	if m != nil {
		return m.Rem
	}
	return 0
}

func init() {
	proto.RegisterType((*ArithRequest)(nil), "pb.ArithRequest")
	proto.RegisterType((*ArithResponse)(nil), "pb.ArithResponse")
}

func init() { proto.RegisterFile("arith.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0xca, 0x2c,
	0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe2, 0xe2, 0x71,
	0x04, 0x09, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1, 0x70, 0x31, 0x26, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x26, 0x82, 0x78, 0x49, 0x12, 0x4c, 0x10, 0x5e, 0x92, 0x92,
	0x2b, 0x17, 0x2f, 0x54, 0x6d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x41,
	0x51, 0x3e, 0x54, 0x39, 0x88, 0x09, 0x12, 0x29, 0x2c, 0xcd, 0x87, 0x6a, 0x01, 0x31, 0x41, 0x22,
	0x45, 0xa9, 0xb9, 0x12, 0xcc, 0x10, 0x91, 0xa2, 0xd4, 0x5c, 0xa3, 0x3c, 0xa8, 0x95, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xfa, 0x5c, 0x1c, 0xb9, 0xa5, 0x39, 0x25, 0x99, 0x05, 0x39,
	0x95, 0x42, 0x02, 0x7a, 0x05, 0x49, 0x7a, 0xc8, 0x0e, 0x92, 0x12, 0x44, 0x12, 0x81, 0x5a, 0xab,
	0xcb, 0xc5, 0x96, 0x92, 0x59, 0x96, 0x99, 0x92, 0x4a, 0x94, 0xf2, 0x24, 0x36, 0xb0, 0x6f, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x03, 0xde, 0xe9, 0xfc, 0x00, 0x00, 0x00,
}