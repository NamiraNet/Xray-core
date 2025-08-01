// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: transport/internet/grpc/config.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Authority           string                 `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	ServiceName         string                 `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	MultiMode           bool                   `protobuf:"varint,3,opt,name=multi_mode,json=multiMode,proto3" json:"multi_mode,omitempty"`
	IdleTimeout         int32                  `protobuf:"varint,4,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	HealthCheckTimeout  int32                  `protobuf:"varint,5,opt,name=health_check_timeout,json=healthCheckTimeout,proto3" json:"health_check_timeout,omitempty"`
	PermitWithoutStream bool                   `protobuf:"varint,6,opt,name=permit_without_stream,json=permitWithoutStream,proto3" json:"permit_without_stream,omitempty"`
	InitialWindowsSize  int32                  `protobuf:"varint,7,opt,name=initial_windows_size,json=initialWindowsSize,proto3" json:"initial_windows_size,omitempty"`
	UserAgent           string                 `protobuf:"bytes,8,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_internet_grpc_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_grpc_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_transport_internet_grpc_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *Config) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Config) GetMultiMode() bool {
	if x != nil {
		return x.MultiMode
	}
	return false
}

func (x *Config) GetIdleTimeout() int32 {
	if x != nil {
		return x.IdleTimeout
	}
	return 0
}

func (x *Config) GetHealthCheckTimeout() int32 {
	if x != nil {
		return x.HealthCheckTimeout
	}
	return 0
}

func (x *Config) GetPermitWithoutStream() bool {
	if x != nil {
		return x.PermitWithoutStream
	}
	return false
}

func (x *Config) GetInitialWindowsSize() int32 {
	if x != nil {
		return x.InitialWindowsSize
	}
	return 0
}

func (x *Config) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

var File_transport_internet_grpc_config_proto protoreflect.FileDescriptor

const file_transport_internet_grpc_config_proto_rawDesc = "" +
	"\n" +
	"$transport/internet/grpc/config.proto\x12%xray.transport.internet.grpc.encoding\"\xc2\x02\n" +
	"\x06Config\x12\x1c\n" +
	"\tauthority\x18\x01 \x01(\tR\tauthority\x12!\n" +
	"\fservice_name\x18\x02 \x01(\tR\vserviceName\x12\x1d\n" +
	"\n" +
	"multi_mode\x18\x03 \x01(\bR\tmultiMode\x12!\n" +
	"\fidle_timeout\x18\x04 \x01(\x05R\vidleTimeout\x120\n" +
	"\x14health_check_timeout\x18\x05 \x01(\x05R\x12healthCheckTimeout\x122\n" +
	"\x15permit_without_stream\x18\x06 \x01(\bR\x13permitWithoutStream\x120\n" +
	"\x14initial_windows_size\x18\a \x01(\x05R\x12initialWindowsSize\x12\x1d\n" +
	"\n" +
	"user_agent\x18\b \x01(\tR\tuserAgentB8Z6github.com/NamiraNet/xray-core/transport/internet/grpcb\x06proto3"

var (
	file_transport_internet_grpc_config_proto_rawDescOnce sync.Once
	file_transport_internet_grpc_config_proto_rawDescData []byte
)

func file_transport_internet_grpc_config_proto_rawDescGZIP() []byte {
	file_transport_internet_grpc_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_grpc_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_internet_grpc_config_proto_rawDesc), len(file_transport_internet_grpc_config_proto_rawDesc)))
	})
	return file_transport_internet_grpc_config_proto_rawDescData
}

var file_transport_internet_grpc_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_internet_grpc_config_proto_goTypes = []any{
	(*Config)(nil), // 0: xray.transport.internet.grpc.encoding.Config
}
var file_transport_internet_grpc_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_internet_grpc_config_proto_init() }
func file_transport_internet_grpc_config_proto_init() {
	if File_transport_internet_grpc_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_internet_grpc_config_proto_rawDesc), len(file_transport_internet_grpc_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_grpc_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_grpc_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_grpc_config_proto_msgTypes,
	}.Build()
	File_transport_internet_grpc_config_proto = out.File
	file_transport_internet_grpc_config_proto_goTypes = nil
	file_transport_internet_grpc_config_proto_depIdxs = nil
}
