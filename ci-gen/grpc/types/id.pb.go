// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/ci-gen/grpc/types/id.proto

package types

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

// IntID is a basic integer ID -- it's used to control a variety of things that
// use basic sequences and so forth.
type IntID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntID) Reset()         { *m = IntID{} }
func (m *IntID) String() string { return proto.CompactTextString(m) }
func (*IntID) ProtoMessage()    {}
func (*IntID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0b61b7a4bc4522, []int{0}
}

func (m *IntID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntID.Unmarshal(m, b)
}
func (m *IntID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntID.Marshal(b, m, deterministic)
}
func (m *IntID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntID.Merge(m, src)
}
func (m *IntID) XXX_Size() int {
	return xxx_messageInfo_IntID.Size(m)
}
func (m *IntID) XXX_DiscardUnknown() {
	xxx_messageInfo_IntID.DiscardUnknown(m)
}

var xxx_messageInfo_IntID proto.InternalMessageInfo

func (m *IntID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

// StringID is just like types.IntID but instead it's a string!
type StringID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringID) Reset()         { *m = StringID{} }
func (m *StringID) String() string { return proto.CompactTextString(m) }
func (*StringID) ProtoMessage()    {}
func (*StringID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0b61b7a4bc4522, []int{1}
}

func (m *StringID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringID.Unmarshal(m, b)
}
func (m *StringID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringID.Marshal(b, m, deterministic)
}
func (m *StringID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringID.Merge(m, src)
}
func (m *StringID) XXX_Size() int {
	return xxx_messageInfo_StringID.Size(m)
}
func (m *StringID) XXX_DiscardUnknown() {
	xxx_messageInfo_StringID.DiscardUnknown(m)
}

var xxx_messageInfo_StringID proto.InternalMessageInfo

func (m *StringID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*IntID)(nil), "types.IntID")
	proto.RegisterType((*StringID)(nil), "types.StringID")
}

func init() {
	proto.RegisterFile("github.com/tinyci/ci-agents/ci-gen/grpc/types/id.proto", fileDescriptor_fe0b61b7a4bc4522)
}

var fileDescriptor_fe0b61b7a4bc4522 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xc9, 0xcc, 0xab, 0x4c, 0xce, 0xd4, 0x4f, 0xce,
	0xd4, 0x4d, 0x4c, 0x4f, 0xcd, 0x2b, 0x29, 0x06, 0xb1, 0xd2, 0x53, 0xf3, 0xf4, 0xd3, 0x8b, 0x0a,
	0x92, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x33, 0x53, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x58, 0xc1, 0x7c, 0x25, 0x71, 0x2e, 0x56, 0xcf, 0xbc, 0x12, 0x4f, 0x17, 0x21, 0x3e, 0x2e,
	0x26, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x26, 0x4f, 0x17, 0x25, 0x29, 0x2e,
	0x8e, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x14, 0x39, 0x4e, 0x90, 0x9c, 0x93, 0x7e, 0x94, 0x2e,
	0x49, 0xb6, 0x26, 0xb1, 0x81, 0xed, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x28, 0x2c, 0x7d,
	0x31, 0xad, 0x00, 0x00, 0x00,
}
