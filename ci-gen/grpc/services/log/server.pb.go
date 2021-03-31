// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.14.0
// source: grpc/services/log/server.proto

package log

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// LogMessage is a log message in struct form.
type LogMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	At      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=at,proto3" json:"at,omitempty"`           // Time of log
	Level   string                 `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`     // Level of log (debug, info, error are supported)
	Fields  *structpb.Struct       `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`   // Fields in map[string]interface{} format
	Service string                 `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"` // Service name
	Message string                 `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"` // Message
}

func (x *LogMessage) Reset() {
	*x = LogMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_services_log_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogMessage) ProtoMessage() {}

func (x *LogMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_services_log_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogMessage.ProtoReflect.Descriptor instead.
func (*LogMessage) Descriptor() ([]byte, []int) {
	return file_grpc_services_log_server_proto_rawDescGZIP(), []int{0}
}

func (x *LogMessage) GetAt() *timestamppb.Timestamp {
	if x != nil {
		return x.At
	}
	return nil
}

func (x *LogMessage) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LogMessage) GetFields() *structpb.Struct {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *LogMessage) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *LogMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_services_log_server_proto protoreflect.FileDescriptor

var file_grpc_services_log_server_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x6c, 0x6f, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2a, 0x0a, 0x02, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x35, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x2e,
	0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e,
	0x79, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x69,
	0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_services_log_server_proto_rawDescOnce sync.Once
	file_grpc_services_log_server_proto_rawDescData = file_grpc_services_log_server_proto_rawDesc
)

func file_grpc_services_log_server_proto_rawDescGZIP() []byte {
	file_grpc_services_log_server_proto_rawDescOnce.Do(func() {
		file_grpc_services_log_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_services_log_server_proto_rawDescData)
	})
	return file_grpc_services_log_server_proto_rawDescData
}

var file_grpc_services_log_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_services_log_server_proto_goTypes = []interface{}{
	(*LogMessage)(nil),            // 0: log.LogMessage
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 2: google.protobuf.Struct
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_grpc_services_log_server_proto_depIdxs = []int32{
	1, // 0: log.LogMessage.at:type_name -> google.protobuf.Timestamp
	2, // 1: log.LogMessage.fields:type_name -> google.protobuf.Struct
	0, // 2: log.Log.Put:input_type -> log.LogMessage
	3, // 3: log.Log.Put:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_services_log_server_proto_init() }
func file_grpc_services_log_server_proto_init() {
	if File_grpc_services_log_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_services_log_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogMessage); i {
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
			RawDescriptor: file_grpc_services_log_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_services_log_server_proto_goTypes,
		DependencyIndexes: file_grpc_services_log_server_proto_depIdxs,
		MessageInfos:      file_grpc_services_log_server_proto_msgTypes,
	}.Build()
	File_grpc_services_log_server_proto = out.File
	file_grpc_services_log_server_proto_rawDesc = nil
	file_grpc_services_log_server_proto_goTypes = nil
	file_grpc_services_log_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogClient interface {
	Put(ctx context.Context, in *LogMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) Put(ctx context.Context, in *LogMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/log.Log/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
type LogServer interface {
	Put(context.Context, *LogMessage) (*emptypb.Empty, error)
}

// UnimplementedLogServer can be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (*UnimplementedLogServer) Put(context.Context, *LogMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}

func RegisterLogServer(s *grpc.Server, srv LogServer) {
	s.RegisterService(&_Log_serviceDesc, srv)
}

func _Log_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.Log/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Put(ctx, req.(*LogMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Log_serviceDesc = grpc.ServiceDesc{
	ServiceName: "log.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Log_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/services/log/server.proto",
}
