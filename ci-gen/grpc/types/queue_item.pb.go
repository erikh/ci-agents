// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/ci-gen/grpc/types/queue_item.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// QueueItems are the subject sent to runners when runners are able to execute
// a job. Runners poll for these endless through the queuesvc.
type QueueItem struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Running              bool                 `protobuf:"varint,2,opt,name=running,proto3" json:"running,omitempty"`
	RunningOn            string               `protobuf:"bytes,3,opt,name=runningOn,proto3" json:"runningOn,omitempty"`
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	QueueName            string               `protobuf:"bytes,5,opt,name=queueName,proto3" json:"queueName,omitempty"`
	Run                  *Run                 `protobuf:"bytes,6,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueueItem) Reset()         { *m = QueueItem{} }
func (m *QueueItem) String() string { return proto.CompactTextString(m) }
func (*QueueItem) ProtoMessage()    {}
func (*QueueItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeff48510487ca01, []int{0}
}

func (m *QueueItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueItem.Unmarshal(m, b)
}
func (m *QueueItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueItem.Marshal(b, m, deterministic)
}
func (m *QueueItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueItem.Merge(m, src)
}
func (m *QueueItem) XXX_Size() int {
	return xxx_messageInfo_QueueItem.Size(m)
}
func (m *QueueItem) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueItem.DiscardUnknown(m)
}

var xxx_messageInfo_QueueItem proto.InternalMessageInfo

func (m *QueueItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueueItem) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *QueueItem) GetRunningOn() string {
	if m != nil {
		return m.RunningOn
	}
	return ""
}

func (m *QueueItem) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *QueueItem) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueItem) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

// QueueRequest is issued by runners to the queuesvc.
type QueueRequest struct {
	QueueName            string   `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
	RunningOn            string   `protobuf:"bytes,2,opt,name=runningOn,proto3" json:"runningOn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueRequest) Reset()         { *m = QueueRequest{} }
func (m *QueueRequest) String() string { return proto.CompactTextString(m) }
func (*QueueRequest) ProtoMessage()    {}
func (*QueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeff48510487ca01, []int{1}
}

func (m *QueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueRequest.Unmarshal(m, b)
}
func (m *QueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueRequest.Marshal(b, m, deterministic)
}
func (m *QueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueRequest.Merge(m, src)
}
func (m *QueueRequest) XXX_Size() int {
	return xxx_messageInfo_QueueRequest.Size(m)
}
func (m *QueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueueRequest proto.InternalMessageInfo

func (m *QueueRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueRequest) GetRunningOn() string {
	if m != nil {
		return m.RunningOn
	}
	return ""
}

// Status is reported to the queuesvc on completion of a run.
type Status struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	AdditionalMessage    string   `protobuf:"bytes,3,opt,name=additionalMessage,proto3" json:"additionalMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeff48510487ca01, []int{2}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Status) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Status) GetAdditionalMessage() string {
	if m != nil {
		return m.AdditionalMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*QueueItem)(nil), "types.QueueItem")
	proto.RegisterType((*QueueRequest)(nil), "types.QueueRequest")
	proto.RegisterType((*Status)(nil), "types.Status")
}

func init() {
	proto.RegisterFile("github.com/tinyci/ci-agents/ci-gen/grpc/types/queue_item.proto", fileDescriptor_aeff48510487ca01)
}

var fileDescriptor_aeff48510487ca01 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x49, 0xfa, 0x6f, 0xfe, 0x66, 0x15, 0xc1, 0x1c, 0x24, 0x94, 0x1e, 0x42, 0x4f, 0x39,
	0xd8, 0x5d, 0xd0, 0x83, 0x82, 0x20, 0xe8, 0x4d, 0x41, 0xc5, 0xd5, 0x93, 0x07, 0x65, 0x9b, 0x8c,
	0xeb, 0x42, 0xb3, 0x9b, 0x66, 0x67, 0x0f, 0xfd, 0x88, 0x7e, 0x2b, 0xc9, 0x26, 0xb5, 0xa5, 0x3d,
	0xf5, 0x36, 0xf3, 0x66, 0xe7, 0xc7, 0xbc, 0xb7, 0xe4, 0x46, 0x2a, 0xfc, 0x76, 0x33, 0x5a, 0x98,
	0x8a, 0xa1, 0xd2, 0xcb, 0x42, 0xb1, 0x42, 0x4d, 0x85, 0x04, 0x8d, 0xb6, 0xad, 0x24, 0x68, 0x26,
	0x9b, 0xba, 0x60, 0xb8, 0xac, 0xc1, 0xb2, 0x85, 0x03, 0x07, 0x9f, 0x0a, 0xa1, 0xa2, 0x75, 0x63,
	0xd0, 0x24, 0x43, 0xaf, 0x8f, 0xae, 0x37, 0x30, 0xd2, 0xcc, 0x85, 0x96, 0xcc, 0xcf, 0x67, 0xee,
	0x8b, 0xd5, 0xdd, 0x2a, 0xaa, 0x0a, 0x2c, 0x8a, 0xaa, 0x5e, 0x57, 0x1d, 0x63, 0x74, 0xb9, 0xdf,
	0x0d, 0x8d, 0xd3, 0xdd, 0xe2, 0xe4, 0x27, 0x20, 0xf1, 0x4b, 0x7b, 0xd1, 0x3d, 0x42, 0x95, 0x1c,
	0x93, 0x50, 0x95, 0x69, 0x90, 0x05, 0xf9, 0x80, 0x87, 0xaa, 0x4c, 0x52, 0xf2, 0xbf, 0x71, 0x5a,
	0x2b, 0x2d, 0xd3, 0x30, 0x0b, 0xf2, 0x03, 0xbe, 0x6a, 0x93, 0x31, 0x89, 0xfb, 0xf2, 0x59, 0xa7,
	0x83, 0x2c, 0xc8, 0x63, 0xbe, 0x16, 0x92, 0x2b, 0x12, 0x5b, 0x14, 0x0d, 0x42, 0x79, 0x8b, 0xe9,
	0xbf, 0x2c, 0xc8, 0x0f, 0xcf, 0x47, 0x54, 0x1a, 0x23, 0xe7, 0x40, 0x57, 0xa6, 0xe8, 0xdb, 0xca,
	0x03, 0x5f, 0x3f, 0x6e, 0xb9, 0x3e, 0xa0, 0x27, 0x51, 0x41, 0x3a, 0xec, 0xb8, 0x7f, 0x42, 0x32,
	0x26, 0x83, 0xc6, 0xe9, 0x34, 0xf2, 0x44, 0x42, 0xbd, 0x19, 0xca, 0x9d, 0xe6, 0xad, 0x3c, 0x79,
	0x20, 0x47, 0xde, 0x0a, 0x87, 0x85, 0x03, 0xbb, 0xc5, 0x0a, 0x76, 0x59, 0x1b, 0x0e, 0xc2, 0x2d,
	0x07, 0x93, 0x0f, 0x12, 0xbd, 0xa2, 0x40, 0x67, 0x77, 0x32, 0x39, 0x25, 0x91, 0xf5, 0x93, 0x3e,
	0x92, 0xbe, 0x4b, 0xce, 0xc8, 0x89, 0x28, 0x4b, 0x85, 0xca, 0x68, 0x31, 0x7f, 0x04, 0x6b, 0x85,
	0x84, 0x3e, 0x99, 0xdd, 0xc1, 0x1d, 0x7b, 0x9f, 0xee, 0xf5, 0x65, 0xb3, 0xc8, 0xe7, 0x76, 0xf1,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x8e, 0x81, 0x6b, 0x6e, 0x02, 0x00, 0x00,
}
