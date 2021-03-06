// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ContractCreate.proto

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

// Start a new smart contract instance. After the instance is created, the ContractID for it is in the receipt, or can be retrieved with a GetByKey query, or by asking for a Record of the transaction to be created, and retrieving that. The instance will run the bytecode stored in the given file, referenced either by FileID or by the transaction ID of the transaction that created the file. The constructor will be executed using the given amount of gas, and any unspent gas will be refunded to the paying account. Constructor inputs come from the given constructorParameters.
//
// The instance will exist for autoRenewPeriod seconds. When that is reached, it will renew itself for another autoRenewPeriod seconds by charging its associated cryptocurrency account (which it creates here). If it has insufficient cryptocurrency to extend that long, it will extend as long as it can. If its balance is zero, the instance will be deleted.
//
// A smart contract instance normally enforces rules, so "the code is law". For example, an ERC-20 contract prevents a transfer from being undone without a signature by the recipient of the transfer. This is always enforced if the contract instance was created with the adminKeys being null. But for some uses, it might be desirable to create something like an ERC-20 contract that has a specific group of trusted individuals who can act as a "supreme court" with the ability to override the normal operation, when a sufficient number of them agree to do so. If adminKeys is not null, then they can sign a transaction that can change the state of the smart contract in arbitrary ways, such as to reverse a transaction that violates some standard of behavior that is not covered by the code itself. The admin keys can also be used to change the autoRenewPeriod, and change the adminKeys field itself. The API currently does not implement this ability. But it does allow the adminKeys field to be set and queried, and will in the future implement such admin abilities for any instance that has a non-null adminKeys.
//
// If this constructor stores information, it is charged gas to store it. There is a fee in hbars to maintain that storage until the expiration time, and that fee is added as part of the transaction fee.
//
// An entity (account, file, or smart contract instance) must be created in a particular realm. If the realmID is left null, then a new realm will be created with the given admin key. If a new realm has a null adminKey, then anyone can create/modify/delete entities in that realm. But if an admin key is given, then any transaction to create/modify/delete an entity in that realm must be signed by that key, though anyone can still call functions on smart contract instances that exist in that realm. A realm ceases to exist when everything within it has expired and no longer exists.
//
// The current API ignores shardID, realmID, and newRealmAdminKey, and creates everything in shard 0 and realm 0, with a null key. Future versions of the API will support multiple realms and multiple shards.
//
// The optional memo field can contain a string whose length is up to 100 bytes. That is the size after Unicode NFD then UTF-8 conversion. This field can be used to describe the smart contract. It could also be used for other purposes. One recommended purpose is to hold a hexadecimal string that is the SHA-384 hash of a PDF file containing a human-readable legal contract. Then, if the admin keys are the public keys of human arbitrators, they can use that legal document to guide their decisions during a binding arbitration tribunal, convened to consider any changes to the smart contract in the future. The memo field can only be changed using the admin keys. If there are no admin keys, then it cannot be changed after the smart contract is created.
type ContractCreateTransactionBody struct {
	FileID                *FileID    `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	AdminKey              *Key       `protobuf:"bytes,3,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	Gas                   int64      `protobuf:"varint,4,opt,name=gas,proto3" json:"gas,omitempty"`
	InitialBalance        int64      `protobuf:"varint,5,opt,name=initialBalance,proto3" json:"initialBalance,omitempty"`
	ProxyAccountID        *AccountID `protobuf:"bytes,6,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`
	AutoRenewPeriod       *Duration  `protobuf:"bytes,8,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	ConstructorParameters []byte     `protobuf:"bytes,9,opt,name=constructorParameters,proto3" json:"constructorParameters,omitempty"`
	ShardID               *ShardID   `protobuf:"bytes,10,opt,name=shardID,proto3" json:"shardID,omitempty"`
	RealmID               *RealmID   `protobuf:"bytes,11,opt,name=realmID,proto3" json:"realmID,omitempty"`
	NewRealmAdminKey      *Key       `protobuf:"bytes,12,opt,name=newRealmAdminKey,proto3" json:"newRealmAdminKey,omitempty"`
	Memo                  string     `protobuf:"bytes,13,opt,name=memo,proto3" json:"memo,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}   `json:"-"`
	XXX_unrecognized      []byte     `json:"-"`
	XXX_sizecache         int32      `json:"-"`
}

func (m *ContractCreateTransactionBody) Reset()         { *m = ContractCreateTransactionBody{} }
func (m *ContractCreateTransactionBody) String() string { return proto.CompactTextString(m) }
func (*ContractCreateTransactionBody) ProtoMessage()    {}
func (*ContractCreateTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_009340fb2588b611, []int{0}
}

func (m *ContractCreateTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractCreateTransactionBody.Unmarshal(m, b)
}
func (m *ContractCreateTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractCreateTransactionBody.Marshal(b, m, deterministic)
}
func (m *ContractCreateTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCreateTransactionBody.Merge(m, src)
}
func (m *ContractCreateTransactionBody) XXX_Size() int {
	return xxx_messageInfo_ContractCreateTransactionBody.Size(m)
}
func (m *ContractCreateTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCreateTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCreateTransactionBody proto.InternalMessageInfo

func (m *ContractCreateTransactionBody) GetFileID() *FileID {
	if m != nil {
		return m.FileID
	}
	return nil
}

func (m *ContractCreateTransactionBody) GetAdminKey() *Key {
	if m != nil {
		return m.AdminKey
	}
	return nil
}

func (m *ContractCreateTransactionBody) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *ContractCreateTransactionBody) GetInitialBalance() int64 {
	if m != nil {
		return m.InitialBalance
	}
	return 0
}

func (m *ContractCreateTransactionBody) GetProxyAccountID() *AccountID {
	if m != nil {
		return m.ProxyAccountID
	}
	return nil
}

func (m *ContractCreateTransactionBody) GetAutoRenewPeriod() *Duration {
	if m != nil {
		return m.AutoRenewPeriod
	}
	return nil
}

func (m *ContractCreateTransactionBody) GetConstructorParameters() []byte {
	if m != nil {
		return m.ConstructorParameters
	}
	return nil
}

func (m *ContractCreateTransactionBody) GetShardID() *ShardID {
	if m != nil {
		return m.ShardID
	}
	return nil
}

func (m *ContractCreateTransactionBody) GetRealmID() *RealmID {
	if m != nil {
		return m.RealmID
	}
	return nil
}

func (m *ContractCreateTransactionBody) GetNewRealmAdminKey() *Key {
	if m != nil {
		return m.NewRealmAdminKey
	}
	return nil
}

func (m *ContractCreateTransactionBody) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func init() {
	proto.RegisterType((*ContractCreateTransactionBody)(nil), "proto.ContractCreateTransactionBody")
}

func init() {
	proto.RegisterFile("proto/ContractCreate.proto", fileDescriptor_009340fb2588b611)
}

var fileDescriptor_009340fb2588b611 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x6a, 0xdc, 0x30,
	0x10, 0x86, 0x31, 0x9b, 0x6c, 0x13, 0x25, 0xd9, 0x2c, 0xa2, 0x2d, 0x62, 0xa1, 0x60, 0x0a, 0x0d,
	0xbe, 0xc4, 0x86, 0xb6, 0x94, 0xf6, 0x18, 0xc7, 0x94, 0x9a, 0x5c, 0x16, 0x35, 0xa7, 0xde, 0x26,
	0xf2, 0xd4, 0x16, 0xb5, 0x25, 0x23, 0xc9, 0xa4, 0x7e, 0xee, 0xbe, 0x40, 0x59, 0xc9, 0x5e, 0xc8,
	0x26, 0x27, 0x89, 0xef, 0xff, 0x66, 0x7e, 0x18, 0xb2, 0xe9, 0x8d, 0x76, 0x3a, 0xbb, 0xd5, 0xca,
	0x19, 0x10, 0xee, 0xd6, 0x20, 0x38, 0x4c, 0x3d, 0xa4, 0xc7, 0xfe, 0xd9, 0xbc, 0x0d, 0x4a, 0x0e,
	0x56, 0x8a, 0xfb, 0xb1, 0x47, 0x1b, 0xe2, 0xcd, 0xeb, 0xc0, 0x8b, 0xc1, 0x80, 0x93, 0x5a, 0x05,
	0xfa, 0xfe, 0xdf, 0x82, 0xbc, 0x7b, 0xba, 0xed, 0xde, 0x80, 0xb2, 0x20, 0x76, 0x4e, 0xae, 0xab,
	0x91, 0x7e, 0x20, 0xcb, 0xdf, 0xb2, 0xc5, 0xb2, 0x60, 0x51, 0x1c, 0x25, 0x67, 0x1f, 0x2f, 0xc2,
	0x64, 0xfa, 0xdd, 0x43, 0x3e, 0x85, 0xf4, 0x8a, 0x9c, 0x40, 0xd5, 0x49, 0x75, 0x87, 0x23, 0x5b,
	0x78, 0x91, 0x4c, 0xe2, 0x1d, 0x8e, 0x7c, 0x9f, 0xd1, 0x35, 0x59, 0xd4, 0x60, 0xd9, 0x51, 0x1c,
	0x25, 0x0b, 0xbe, 0xfb, 0xd2, 0x2b, 0xb2, 0x92, 0x4a, 0x3a, 0x09, 0x6d, 0x0e, 0x2d, 0x28, 0x81,
	0xec, 0xd8, 0x87, 0x07, 0x94, 0x7e, 0x25, 0xab, 0xde, 0xe8, 0xbf, 0xe3, 0x8d, 0x10, 0x7a, 0x50,
	0xae, 0x2c, 0xd8, 0xd2, 0xf7, 0xac, 0xa7, 0x9e, 0x3d, 0xe7, 0x07, 0x1e, 0xfd, 0x46, 0x2e, 0x61,
	0x70, 0x9a, 0xa3, 0xc2, 0xc7, 0x2d, 0x1a, 0xa9, 0x2b, 0x76, 0xe2, 0x47, 0x2f, 0xa7, 0xd1, 0xf9,
	0x28, 0xfc, 0xd0, 0xa3, 0x9f, 0xc9, 0x1b, 0xa1, 0x95, 0x75, 0x66, 0x10, 0x4e, 0x9b, 0x2d, 0x18,
	0xe8, 0xd0, 0xa1, 0xb1, 0xec, 0x34, 0x8e, 0x92, 0x73, 0xfe, 0x72, 0x48, 0x13, 0xf2, 0xca, 0x36,
	0x60, 0xaa, 0xb2, 0x60, 0xc4, 0x17, 0xad, 0xa6, 0xa2, 0x9f, 0x81, 0xf2, 0x39, 0xde, 0x99, 0x06,
	0xa1, 0xed, 0xca, 0x82, 0x9d, 0x3d, 0x31, 0x79, 0xa0, 0x7c, 0x8e, 0xe9, 0x17, 0xb2, 0x56, 0xf8,
	0xe8, 0xf1, 0xcd, 0x7c, 0xe8, 0xf3, 0x67, 0x87, 0x7e, 0xe6, 0x50, 0x4a, 0x8e, 0x3a, 0xec, 0x34,
	0xbb, 0x88, 0xa3, 0xe4, 0x94, 0xfb, 0x7f, 0xfe, 0x83, 0x6c, 0x84, 0xee, 0xd2, 0x06, 0x2b, 0x34,
	0x90, 0x36, 0x60, 0x9b, 0xda, 0x40, 0xdf, 0x84, 0x3d, 0xdb, 0xe8, 0x57, 0x52, 0x4b, 0xd7, 0x0c,
	0x0f, 0xa9, 0xd0, 0x5d, 0xb6, 0x4f, 0xb3, 0xa0, 0x5f, 0xdb, 0xea, 0xcf, 0x75, 0xad, 0x33, 0xef,
	0x3e, 0x2c, 0xfd, 0xf3, 0xe9, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x91, 0x8e, 0xb8, 0x99,
	0x02, 0x00, 0x00,
}
