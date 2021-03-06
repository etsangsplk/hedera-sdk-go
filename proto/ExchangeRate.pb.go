// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ExchangeRate.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//  Values from these proto denotes habr and cents(USD) conversion
type ExchangeRate struct {
	HbarEquiv            int32             `protobuf:"varint,1,opt,name=hbarEquiv,proto3" json:"hbarEquiv,omitempty"`
	CentEquiv            int32             `protobuf:"varint,2,opt,name=centEquiv,proto3" json:"centEquiv,omitempty"`
	ExpirationTime       *TimestampSeconds `protobuf:"bytes,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExchangeRate) Reset()         { *m = ExchangeRate{} }
func (m *ExchangeRate) String() string { return proto.CompactTextString(m) }
func (*ExchangeRate) ProtoMessage()    {}
func (*ExchangeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6948300117a927ac, []int{0}
}

func (m *ExchangeRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeRate.Unmarshal(m, b)
}
func (m *ExchangeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeRate.Marshal(b, m, deterministic)
}
func (m *ExchangeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeRate.Merge(m, src)
}
func (m *ExchangeRate) XXX_Size() int {
	return xxx_messageInfo_ExchangeRate.Size(m)
}
func (m *ExchangeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeRate.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeRate proto.InternalMessageInfo

func (m *ExchangeRate) GetHbarEquiv() int32 {
	if m != nil {
		return m.HbarEquiv
	}
	return 0
}

func (m *ExchangeRate) GetCentEquiv() int32 {
	if m != nil {
		return m.CentEquiv
	}
	return 0
}

func (m *ExchangeRate) GetExpirationTime() *TimestampSeconds {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

// Two sets of exchange rate
type ExchangeRateSet struct {
	CurrentRate          *ExchangeRate `protobuf:"bytes,1,opt,name=currentRate,proto3" json:"currentRate,omitempty"`
	NextRate             *ExchangeRate `protobuf:"bytes,2,opt,name=nextRate,proto3" json:"nextRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExchangeRateSet) Reset()         { *m = ExchangeRateSet{} }
func (m *ExchangeRateSet) String() string { return proto.CompactTextString(m) }
func (*ExchangeRateSet) ProtoMessage()    {}
func (*ExchangeRateSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_6948300117a927ac, []int{1}
}

func (m *ExchangeRateSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeRateSet.Unmarshal(m, b)
}
func (m *ExchangeRateSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeRateSet.Marshal(b, m, deterministic)
}
func (m *ExchangeRateSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeRateSet.Merge(m, src)
}
func (m *ExchangeRateSet) XXX_Size() int {
	return xxx_messageInfo_ExchangeRateSet.Size(m)
}
func (m *ExchangeRateSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeRateSet.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeRateSet proto.InternalMessageInfo

func (m *ExchangeRateSet) GetCurrentRate() *ExchangeRate {
	if m != nil {
		return m.CurrentRate
	}
	return nil
}

func (m *ExchangeRateSet) GetNextRate() *ExchangeRate {
	if m != nil {
		return m.NextRate
	}
	return nil
}

func init() {
	proto.RegisterType((*ExchangeRate)(nil), "proto.ExchangeRate")
	proto.RegisterType((*ExchangeRateSet)(nil), "proto.ExchangeRateSet")
}

func init() {
	proto.RegisterFile("proto/ExchangeRate.proto", fileDescriptor_6948300117a927ac)
}

var fileDescriptor_6948300117a927ac = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xc9, 0x64, 0xa2, 0xa9, 0x28, 0x44, 0xc4, 0x32, 0x3c, 0x8c, 0x9d, 0x7a, 0x59, 0x02,
	0x13, 0xcf, 0x82, 0x30, 0xf0, 0x28, 0x9d, 0x27, 0x6f, 0x69, 0xfa, 0xd1, 0x04, 0x69, 0x52, 0xd3,
	0x54, 0xea, 0x6f, 0xf0, 0x4f, 0x4b, 0xf3, 0x49, 0x0d, 0xc2, 0x4e, 0x81, 0xf7, 0x79, 0xde, 0xbc,
	0x1f, 0xcd, 0x3b, 0xef, 0x82, 0x13, 0xfb, 0x51, 0x69, 0x69, 0x1b, 0x28, 0x65, 0x00, 0x1e, 0x23,
	0xb6, 0x8c, 0xcf, 0xea, 0x06, 0x85, 0x57, 0xd3, 0x42, 0x1f, 0x64, 0xdb, 0x21, 0xdd, 0x7c, 0x13,
	0x7a, 0x91, 0x96, 0xd8, 0x1d, 0x3d, 0xd7, 0x95, 0xf4, 0xfb, 0x8f, 0xc1, 0x7c, 0xe6, 0x64, 0x4d,
	0x8a, 0x65, 0xf9, 0x17, 0x4c, 0x54, 0x81, 0x0d, 0x48, 0x17, 0x48, 0xe7, 0x80, 0x3d, 0xd2, 0x4b,
	0x18, 0x3b, 0xe3, 0x65, 0x30, 0xce, 0x4e, 0x4b, 0xf9, 0xc9, 0x9a, 0x14, 0xd9, 0xee, 0x16, 0xc7,
	0xf8, 0x3c, 0x7e, 0x00, 0xe5, 0x6c, 0xdd, 0x97, 0xff, 0xf4, 0xcd, 0x17, 0xbd, 0x4a, 0x8f, 0x39,
	0x40, 0x60, 0x0f, 0x34, 0x53, 0x83, 0xf7, 0x60, 0xc3, 0x94, 0xc4, 0x8b, 0xb2, 0xdd, 0xf5, 0xef,
	0x87, 0xa9, 0x5c, 0xa6, 0x1e, 0x13, 0xf4, 0xcc, 0xc2, 0x88, 0x9d, 0xc5, 0xf1, 0xce, 0x2c, 0x3d,
	0x3d, 0xd3, 0x95, 0x72, 0x2d, 0xd7, 0x50, 0x83, 0x97, 0x5c, 0xcb, 0x5e, 0x37, 0x5e, 0x76, 0x1a,
	0x4b, 0x2f, 0xe4, 0xad, 0x68, 0x4c, 0xd0, 0x43, 0xc5, 0x95, 0x6b, 0xc5, 0x4c, 0x05, 0xea, 0xdb,
	0xbe, 0x7e, 0xdf, 0x36, 0x4e, 0x44, 0xb7, 0x3a, 0x8d, 0xcf, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2f, 0xc0, 0x93, 0xe1, 0x93, 0x01, 0x00, 0x00,
}
