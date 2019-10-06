// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/ci-gen/grpc/types/submission.proto

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

type Submission struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	HeadRef              *Ref                 `protobuf:"bytes,2,opt,name=headRef,proto3" json:"headRef,omitempty"`
	BaseRef              *Ref                 `protobuf:"bytes,3,opt,name=baseRef,proto3" json:"baseRef,omitempty"`
	User                 *User                `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	TasksCount           int64                `protobuf:"varint,6,opt,name=tasksCount,proto3" json:"tasksCount,omitempty"`
	FinishedAt           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`
	Status               bool                 `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	StatusSet            bool                 `protobuf:"varint,9,opt,name=statusSet,proto3" json:"statusSet,omitempty"`
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	Canceled             bool                 `protobuf:"varint,11,opt,name=canceled,proto3" json:"canceled,omitempty"`
	TicketID             int64                `protobuf:"varint,12,opt,name=ticketID,proto3" json:"ticketID,omitempty"`
	RunsCount            int64                `protobuf:"varint,13,opt,name=runsCount,proto3" json:"runsCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Submission) Reset()         { *m = Submission{} }
func (m *Submission) String() string { return proto.CompactTextString(m) }
func (*Submission) ProtoMessage()    {}
func (*Submission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f00c03eb97056e8, []int{0}
}

func (m *Submission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Submission.Unmarshal(m, b)
}
func (m *Submission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Submission.Marshal(b, m, deterministic)
}
func (m *Submission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Submission.Merge(m, src)
}
func (m *Submission) XXX_Size() int {
	return xxx_messageInfo_Submission.Size(m)
}
func (m *Submission) XXX_DiscardUnknown() {
	xxx_messageInfo_Submission.DiscardUnknown(m)
}

var xxx_messageInfo_Submission proto.InternalMessageInfo

func (m *Submission) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Submission) GetHeadRef() *Ref {
	if m != nil {
		return m.HeadRef
	}
	return nil
}

func (m *Submission) GetBaseRef() *Ref {
	if m != nil {
		return m.BaseRef
	}
	return nil
}

func (m *Submission) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Submission) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Submission) GetTasksCount() int64 {
	if m != nil {
		return m.TasksCount
	}
	return 0
}

func (m *Submission) GetFinishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

func (m *Submission) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Submission) GetStatusSet() bool {
	if m != nil {
		return m.StatusSet
	}
	return false
}

func (m *Submission) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *Submission) GetCanceled() bool {
	if m != nil {
		return m.Canceled
	}
	return false
}

func (m *Submission) GetTicketID() int64 {
	if m != nil {
		return m.TicketID
	}
	return 0
}

func (m *Submission) GetRunsCount() int64 {
	if m != nil {
		return m.RunsCount
	}
	return 0
}

type SubmissionList struct {
	Submissions          []*Submission `protobuf:"bytes,1,rep,name=submissions,proto3" json:"submissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SubmissionList) Reset()         { *m = SubmissionList{} }
func (m *SubmissionList) String() string { return proto.CompactTextString(m) }
func (*SubmissionList) ProtoMessage()    {}
func (*SubmissionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f00c03eb97056e8, []int{1}
}

func (m *SubmissionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmissionList.Unmarshal(m, b)
}
func (m *SubmissionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmissionList.Marshal(b, m, deterministic)
}
func (m *SubmissionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmissionList.Merge(m, src)
}
func (m *SubmissionList) XXX_Size() int {
	return xxx_messageInfo_SubmissionList.Size(m)
}
func (m *SubmissionList) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmissionList.DiscardUnknown(m)
}

var xxx_messageInfo_SubmissionList proto.InternalMessageInfo

func (m *SubmissionList) GetSubmissions() []*Submission {
	if m != nil {
		return m.Submissions
	}
	return nil
}

func init() {
	proto.RegisterType((*Submission)(nil), "types.Submission")
	proto.RegisterType((*SubmissionList)(nil), "types.SubmissionList")
}

func init() {
	proto.RegisterFile("github.com/tinyci/ci-agents/ci-gen/grpc/types/submission.proto", fileDescriptor_9f00c03eb97056e8)
}

var fileDescriptor_9f00c03eb97056e8 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0x91, 0x9d, 0x38, 0xf6, 0xa8, 0x0d, 0x74, 0x0f, 0x65, 0x31, 0xa5, 0x35, 0xa1, 0x07,
	0x5f, 0x22, 0x41, 0x72, 0x68, 0x68, 0xa1, 0xd0, 0x7f, 0x87, 0x42, 0x4f, 0x4a, 0x7b, 0xe9, 0x6d,
	0xb5, 0x1a, 0xad, 0x97, 0xd8, 0xbb, 0x62, 0x67, 0x74, 0xc8, 0x07, 0xee, 0xf7, 0x28, 0x5a, 0xc9,
	0x96, 0x21, 0x87, 0xe0, 0xdb, 0xee, 0xbc, 0xdf, 0x3c, 0xcd, 0x1b, 0x2d, 0x7c, 0x36, 0x96, 0x37,
	0x6d, 0x99, 0x69, 0xbf, 0xcb, 0xd9, 0xba, 0x47, 0x6d, 0x73, 0x6d, 0xaf, 0x95, 0x41, 0xc7, 0xd4,
	0x9d, 0x0c, 0xba, 0xdc, 0x84, 0x46, 0xe7, 0xfc, 0xd8, 0x20, 0xe5, 0xd4, 0x96, 0x3b, 0x4b, 0x64,
	0xbd, 0xcb, 0x9a, 0xe0, 0xd9, 0x8b, 0xf3, 0x58, 0x5f, 0x7e, 0x3a, 0xb2, 0x31, 0x7e, 0xab, 0x9c,
	0xc9, 0xa3, 0x5e, 0xb6, 0x75, 0xde, 0xf4, 0xad, 0x6c, 0x77, 0x48, 0xac, 0x76, 0xcd, 0x78, 0xea,
	0x3d, 0x96, 0x1f, 0x4e, 0x9b, 0x21, 0x60, 0x3d, 0x34, 0xde, 0x9d, 0xd6, 0xd8, 0x12, 0x86, 0xbe,
	0xf3, 0xea, 0xdf, 0x14, 0xe0, 0xfe, 0x90, 0x45, 0x5c, 0xc2, 0xc4, 0x56, 0x32, 0x59, 0x25, 0xeb,
	0x69, 0x31, 0xb1, 0x95, 0x78, 0x0f, 0x17, 0x1b, 0x54, 0x55, 0x81, 0xb5, 0x9c, 0xac, 0x92, 0x75,
	0x7a, 0x03, 0x59, 0xb4, 0xc8, 0x0a, 0xac, 0x8b, 0xbd, 0xd4, 0x51, 0xa5, 0x22, 0xec, 0xa8, 0xe9,
	0x53, 0x6a, 0x90, 0xc4, 0x3b, 0x38, 0xeb, 0x3e, 0x2c, 0xcf, 0x22, 0x92, 0x0e, 0xc8, 0x1f, 0xc2,
	0x50, 0x44, 0x41, 0xdc, 0xc1, 0x42, 0x07, 0x54, 0x8c, 0xd5, 0x17, 0x96, 0xe7, 0x91, 0x5a, 0x66,
	0xc6, 0x7b, 0xb3, 0xc5, 0x6c, 0xbf, 0xc4, 0xec, 0xf7, 0x7e, 0x67, 0xc5, 0x08, 0x8b, 0xb7, 0x00,
	0xac, 0xe8, 0x81, 0xbe, 0xf9, 0xd6, 0xb1, 0x9c, 0xc5, 0xf1, 0x8f, 0x2a, 0xe2, 0x23, 0x40, 0x6d,
	0x9d, 0xa5, 0x4d, 0xb4, 0xbe, 0x78, 0xd6, 0xfa, 0x88, 0x16, 0xaf, 0x61, 0x46, 0xac, 0xb8, 0x25,
	0x39, 0x5f, 0x25, 0xeb, 0x79, 0x31, 0xdc, 0xc4, 0x1b, 0x58, 0xf4, 0xa7, 0x7b, 0x64, 0xb9, 0x88,
	0xd2, 0x58, 0xe8, 0xb2, 0x10, 0xab, 0xd0, 0x67, 0x81, 0xe7, 0xb3, 0x1c, 0x60, 0xb1, 0x84, 0xb9,
	0x56, 0x4e, 0xe3, 0x16, 0x2b, 0x99, 0x46, 0xdb, 0xc3, 0xbd, 0xd3, 0xd8, 0xea, 0x07, 0xe4, 0x9f,
	0xdf, 0xe5, 0x8b, 0x98, 0xf2, 0x70, 0xef, 0xe6, 0x09, 0xad, 0x1b, 0x56, 0xf0, 0x32, 0x8a, 0x63,
	0xe1, 0xea, 0x07, 0x5c, 0x8e, 0xbf, 0xf9, 0x97, 0x25, 0x16, 0xb7, 0x90, 0x8e, 0x8f, 0x98, 0x64,
	0xb2, 0x9a, 0xae, 0xd3, 0x9b, 0x57, 0xc3, 0x5f, 0x19, 0xd9, 0xe2, 0x98, 0xfa, 0x9a, 0xff, 0xbd,
	0x3e, 0xe9, 0xa9, 0x95, 0xb3, 0x18, 0xf6, 0xf6, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xc5,
	0xa1, 0xbc, 0x5f, 0x03, 0x00, 0x00,
}
