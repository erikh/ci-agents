// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: github.com/tinyci/ci-agents/ci-gen/grpc/types/run.proto

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

// Run is a single CI run, intended to be sent to a runner.
type Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                // ID is the internal ID of the run.
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`             // Name is the name of the run. Typically this is in `dir:run_name` format.
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`   // When was this run created
	StartedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=startedAt,proto3" json:"startedAt,omitempty"`   // When did this run start
	FinishedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"` // When did this run finish
	Status     bool                   `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`        // What is the status of this run
	StatusSet  bool                   `protobuf:"varint,7,opt,name=statusSet,proto3" json:"statusSet,omitempty"`  // Is the status valid? (nil internally for invalid settings, but proto doesn't like nil)
	Settings   *RunSettings           `protobuf:"bytes,8,opt,name=settings,proto3" json:"settings,omitempty"`     // The settings for the Run (image, command etc)
	Task       *Task                  `protobuf:"bytes,9,opt,name=task,proto3" json:"task,omitempty"`             // Task for the Run.
	RanOn      string                 `protobuf:"bytes,10,opt,name=ranOn,proto3" json:"ranOn,omitempty"`          // what host the run happened on
	RanOnSet   bool                   `protobuf:"varint,11,opt,name=ranOnSet,proto3" json:"ranOnSet,omitempty"`   // if the ranOn host was set.
}

func (x *Run) Reset() {
	*x = Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescGZIP(), []int{0}
}

func (x *Run) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Run) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Run) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Run) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Run) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *Run) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Run) GetStatusSet() bool {
	if x != nil {
		return x.StatusSet
	}
	return false
}

func (x *Run) GetSettings() *RunSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Run) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *Run) GetRanOn() string {
	if x != nil {
		return x.RanOn
	}
	return ""
}

func (x *Run) GetRanOnSet() bool {
	if x != nil {
		return x.RanOnSet
	}
	return false
}

// RunList is just an array of runs
type RunList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Run `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"` // The list of runs!
}

func (x *RunList) Reset() {
	*x = RunList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunList) ProtoMessage() {}

func (x *RunList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunList.ProtoReflect.Descriptor instead.
func (*RunList) Descriptor() ([]byte, []int) {
	return file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescGZIP(), []int{1}
}

func (x *RunList) GetList() []*Run {
	if x != nil {
		return x.List
	}
	return nil
}

var File_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto protoreflect.FileDescriptor

var file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e,
	0x79, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x69,
	0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69,
	0x6e, 0x79, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63,
	0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x79, 0x63, 0x69, 0x2f, 0x63,
	0x69, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03,
	0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a,
	0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x74, 0x12,
	0x2e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x1f, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x4f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x61, 0x6e, 0x4f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61, 0x6e, 0x4f, 0x6e, 0x53,
	0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x4f, 0x6e, 0x53,
	0x65, 0x74, 0x22, 0x29, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x79,
	0x63, 0x69, 0x2f, 0x63, 0x69, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x69, 0x2d,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescOnce sync.Once
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescData = file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDesc
)

func file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescGZIP() []byte {
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescOnce.Do(func() {
		file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescData)
	})
	return file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDescData
}

var file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_goTypes = []interface{}{
	(*Run)(nil),                   // 0: types.Run
	(*RunList)(nil),               // 1: types.RunList
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*RunSettings)(nil),           // 3: types.RunSettings
	(*Task)(nil),                  // 4: types.Task
}
var file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_depIdxs = []int32{
	2, // 0: types.Run.createdAt:type_name -> google.protobuf.Timestamp
	2, // 1: types.Run.startedAt:type_name -> google.protobuf.Timestamp
	2, // 2: types.Run.finishedAt:type_name -> google.protobuf.Timestamp
	3, // 3: types.Run.settings:type_name -> types.RunSettings
	4, // 4: types.Run.task:type_name -> types.Task
	0, // 5: types.RunList.list:type_name -> types.Run
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_init() }
func file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_init() {
	if File_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto != nil {
		return
	}
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_task_proto_init()
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Run); i {
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
		file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunList); i {
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
			RawDescriptor: file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_goTypes,
		DependencyIndexes: file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_depIdxs,
		MessageInfos:      file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_msgTypes,
	}.Build()
	File_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto = out.File
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_rawDesc = nil
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_goTypes = nil
	file_github_com_tinyci_ci_agents_ci_gen_grpc_types_run_proto_depIdxs = nil
}
