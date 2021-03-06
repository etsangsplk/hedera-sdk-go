// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ConsensusDeleteTopic.proto

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

// See [ConsensusService.deleteTopic()](#proto.ConsensusService)
type ConsensusDeleteTopicTransactionBody struct {
	TopicID              *TopicID `protobuf:"bytes,1,opt,name=topicID,proto3" json:"topicID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusDeleteTopicTransactionBody) Reset()         { *m = ConsensusDeleteTopicTransactionBody{} }
func (m *ConsensusDeleteTopicTransactionBody) String() string { return proto.CompactTextString(m) }
func (*ConsensusDeleteTopicTransactionBody) ProtoMessage()    {}
func (*ConsensusDeleteTopicTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_321862972fa0dd57, []int{0}
}

func (m *ConsensusDeleteTopicTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusDeleteTopicTransactionBody.Unmarshal(m, b)
}
func (m *ConsensusDeleteTopicTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusDeleteTopicTransactionBody.Marshal(b, m, deterministic)
}
func (m *ConsensusDeleteTopicTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusDeleteTopicTransactionBody.Merge(m, src)
}
func (m *ConsensusDeleteTopicTransactionBody) XXX_Size() int {
	return xxx_messageInfo_ConsensusDeleteTopicTransactionBody.Size(m)
}
func (m *ConsensusDeleteTopicTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusDeleteTopicTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusDeleteTopicTransactionBody proto.InternalMessageInfo

func (m *ConsensusDeleteTopicTransactionBody) GetTopicID() *TopicID {
	if m != nil {
		return m.TopicID
	}
	return nil
}

func init() {
	proto.RegisterType((*ConsensusDeleteTopicTransactionBody)(nil), "proto.ConsensusDeleteTopicTransactionBody")
}

func init() {
	proto.RegisterFile("proto/ConsensusDeleteTopic.proto", fileDescriptor_321862972fa0dd57)
}

var fileDescriptor_321862972fa0dd57 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8b, 0xb1, 0xaa, 0xc2, 0x30,
	0x18, 0x46, 0xe9, 0x70, 0xaf, 0x10, 0xc1, 0xa1, 0x83, 0x48, 0xa7, 0xa2, 0x4b, 0x97, 0x26, 0xa0,
	0x6f, 0x50, 0x3b, 0xe8, 0xa4, 0x48, 0x26, 0xb7, 0x34, 0x0d, 0x4d, 0xd0, 0xe6, 0x0f, 0xf9, 0xd3,
	0xa1, 0x6f, 0x2f, 0x26, 0xe8, 0xe4, 0xf4, 0xc1, 0x77, 0xce, 0x21, 0xa5, 0xf3, 0x10, 0x80, 0x1d,
	0xc1, 0xa2, 0xb2, 0x38, 0x61, 0xab, 0x9e, 0x2a, 0x28, 0x0e, 0xce, 0x48, 0x1a, 0x51, 0xfe, 0x17,
	0xa7, 0x58, 0x27, 0xb1, 0x11, 0x68, 0x24, 0x9f, 0x9d, 0xc2, 0x84, 0xb7, 0x17, 0xb2, 0xfb, 0x15,
	0x73, 0x2f, 0x2c, 0x0a, 0x19, 0x0c, 0xd8, 0x06, 0xfa, 0x39, 0xaf, 0xc8, 0x22, 0xbc, 0xff, 0x73,
	0xbb, 0xc9, 0xca, 0xac, 0x5a, 0xee, 0x57, 0xa9, 0xa7, 0x3c, 0xbd, 0xb7, 0x0f, 0x6e, 0x4e, 0xa4,
	0x90, 0x30, 0x52, 0xad, 0x7a, 0xe5, 0x05, 0xd5, 0x02, 0xf5, 0xe0, 0x85, 0xd3, 0x49, 0xbf, 0x66,
	0xf7, 0x6a, 0x30, 0x41, 0x4f, 0x1d, 0x95, 0x30, 0xb2, 0x2f, 0x65, 0x49, 0xaf, 0xb1, 0x7f, 0xd4,
	0x03, 0xb0, 0xe8, 0x76, 0xff, 0x71, 0x0e, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xb3, 0xca,
	0x8e, 0xe4, 0x00, 0x00, 0x00,
}
