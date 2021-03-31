// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.14.0
// source: github.com/tinyci/ci-agents/ci-gen/grpc/types/submission.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                 // ID of the submission
	HeadRef    *Ref                   `protobuf:"bytes,2,opt,name=headRef,proto3" json:"headRef,omitempty"`        // Head git ref of the submission
	BaseRef    *Ref                   `protobuf:"bytes,3,opt,name=baseRef,proto3" json:"baseRef,omitempty"`        // Base git ref of the submission
	User       *User                  `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`              // User who submitted it
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`    // When it was submitted
	TasksCount int64                  `protobuf:"varint,6,opt,name=tasksCount,proto3" json:"tasksCount,omitempty"` // The number of tasks in this submission
	FinishedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`  // When it completed
	Status     bool                   `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`         // What is the status of this submission
	StatusSet  bool                   `protobuf:"varint,9,opt,name=statusSet,proto3" json:"statusSet,omitempty"`   // Is the status valid? (nil internally for invalid settings, but proto doesn't like nil)
	StartedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=startedAt,proto3" json:"startedAt,omitempty"`   // When it started
	Canceled   bool                   `protobuf:"varint,11,opt,name=canceled,proto3" json:"canceled,omitempty"`    // If the whole submission was canceled
	TicketID   int64                  `protobuf:"varint,12,opt,name=ticketID,proto3" json:"ticketID,omitempty"`    // ID of the corresponding ticket in source control
	RunsCount  int64                  `protobuf:"varint,13,opt,name=runsCount,proto3" json:"runsCount,omitempty"`  // The number of runs in this submission
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescGZIP(), []int{0}
}

func (x *Submission) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Submission) GetHeadRef() *Ref {
	if x != nil {
		return x.HeadRef
	}
	return nil
}

func (x *Submission) GetBaseRef() *Ref {
	if x != nil {
		return x.BaseRef
	}
	return nil
}

func (x *Submission) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Submission) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Submission) GetTasksCount() int64 {
	if x != nil {
		return x.TasksCount
	}
	return 0
}

func (x *Submission) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *Submission) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Submission) GetStatusSet() bool {
	if x != nil {
		return x.StatusSet
	}
	return false
}

func (x *Submission) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Submission) GetCanceled() bool {
	if x != nil {
		return x.Canceled
	}
	return false
}

func (x *Submission) GetTicketID() int64 {
	if x != nil {
		return x.TicketID
	}
	return 0
}

func (x *Submission) GetRunsCount() int64 {
	if x != nil {
		return x.RunsCount
	}
	return 0
}

type SubmissionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submissions []*Submission `protobuf:"bytes,1,rep,name=submissions,proto3" json:"submissions,omitempty"`
}

func (x *SubmissionList) Reset() {
	*x = SubmissionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmissionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionList) ProtoMessage() {}

func (x *SubmissionList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionList.ProtoReflect.Descriptor instead.
func (*SubmissionList) Descriptor() ([]byte, []int) {
	return file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescGZIP(), []int{1}
}

func (x *SubmissionList) GetSubmissions() []*Submission {
	if x != nil {
		return x.Submissions
	}
	return nil
}

var File_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto protoreflect.FileDescriptor

var file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e,
	0x79, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x69,
	0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x79, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x2d, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69,
	0x6e, 0x79, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63,
	0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x03, 0x0a, 0x0a,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x52, 0x65, 0x66,
	0x12, 0x24, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x07, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x66, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53,
	0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x53, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x79, 0x63, 0x69, 0x2f,
	0x63, 0x69, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x69, 0x2d, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescOnce sync.Once
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescData = file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDesc
)

func file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescGZIP() []byte {
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescOnce.Do(func() {
		file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescData)
	})
	return file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDescData
}

var file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_goTypes = []interface{}{
	(*Submission)(nil),            // 0: types.Submission
	(*SubmissionList)(nil),        // 1: types.SubmissionList
	(*Ref)(nil),                   // 2: types.Ref
	(*User)(nil),                  // 3: types.User
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_depIdxs = []int32{
	2, // 0: types.Submission.headRef:type_name -> types.Ref
	2, // 1: types.Submission.baseRef:type_name -> types.Ref
	3, // 2: types.Submission.user:type_name -> types.User
	4, // 3: types.Submission.createdAt:type_name -> google.protobuf.Timestamp
	4, // 4: types.Submission.finishedAt:type_name -> google.protobuf.Timestamp
	4, // 5: types.Submission.startedAt:type_name -> google.protobuf.Timestamp
	0, // 6: types.SubmissionList.submissions:type_name -> types.Submission
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_init() }
func file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_init() {
	if File_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto != nil {
		return
	}
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_ref_proto_init()
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmissionList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_goTypes,
		DependencyIndexes: file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_depIdxs,
		MessageInfos:      file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_msgTypes,
	}.Build()
	File_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto = out.File
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_rawDesc = nil
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_goTypes = nil
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_submission_proto_depIdxs = nil
}
