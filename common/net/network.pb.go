// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: common/net/network.proto

package net

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

type Network int32

const (
	Network_Unknown Network = 0
	Network_TCP     Network = 2
	Network_UDP     Network = 3
	Network_UNIX    Network = 4
)

// Enum value maps for Network.
var (
	Network_name = map[int32]string{
		0: "Unknown",
		2: "TCP",
		3: "UDP",
		4: "UNIX",
	}
	Network_value = map[string]int32{
		"Unknown": 0,
		"TCP":     2,
		"UDP":     3,
		"UNIX":    4,
	}
)

func (x Network) Enum() *Network {
	p := new(Network)
	*p = x
	return p
}

func (x Network) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Network) Descriptor() protoreflect.EnumDescriptor {
	return file_common_net_network_proto_enumTypes[0].Descriptor()
}

func (Network) Type() protoreflect.EnumType {
	return &file_common_net_network_proto_enumTypes[0]
}

func (x Network) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Network.Descriptor instead.
func (Network) EnumDescriptor() ([]byte, []int) {
	return file_common_net_network_proto_rawDescGZIP(), []int{0}
}

// NetworkList is a list of Networks.
type NetworkList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       []Network              `protobuf:"varint,1,rep,packed,name=network,proto3,enum=xray.common.net.Network" json:"network,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkList) Reset() {
	*x = NetworkList{}
	mi := &file_common_net_network_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkList) ProtoMessage() {}

func (x *NetworkList) ProtoReflect() protoreflect.Message {
	mi := &file_common_net_network_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkList.ProtoReflect.Descriptor instead.
func (*NetworkList) Descriptor() ([]byte, []int) {
	return file_common_net_network_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkList) GetNetwork() []Network {
	if x != nil {
		return x.Network
	}
	return nil
}

var File_common_net_network_proto protoreflect.FileDescriptor

const file_common_net_network_proto_rawDesc = "" +
	"\n" +
	"\x18common/net/network.proto\x12\x0fxray.common.net\"A\n" +
	"\vNetworkList\x122\n" +
	"\anetwork\x18\x01 \x03(\x0e2\x18.xray.common.net.NetworkR\anetwork*2\n" +
	"\aNetwork\x12\v\n" +
	"\aUnknown\x10\x00\x12\a\n" +
	"\x03TCP\x10\x02\x12\a\n" +
	"\x03UDP\x10\x03\x12\b\n" +
	"\x04UNIX\x10\x04BT\n" +
	"\x13com.xray.common.netP\x01Z)github.com/NamiraNet/xray-core/common/net\xaa\x02\x0fXray.Common.Netb\x06proto3"

var (
	file_common_net_network_proto_rawDescOnce sync.Once
	file_common_net_network_proto_rawDescData []byte
)

func file_common_net_network_proto_rawDescGZIP() []byte {
	file_common_net_network_proto_rawDescOnce.Do(func() {
		file_common_net_network_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_net_network_proto_rawDesc), len(file_common_net_network_proto_rawDesc)))
	})
	return file_common_net_network_proto_rawDescData
}

var file_common_net_network_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_net_network_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_net_network_proto_goTypes = []any{
	(Network)(0),        // 0: xray.common.net.Network
	(*NetworkList)(nil), // 1: xray.common.net.NetworkList
}
var file_common_net_network_proto_depIdxs = []int32{
	0, // 0: xray.common.net.NetworkList.network:type_name -> xray.common.net.Network
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_net_network_proto_init() }
func file_common_net_network_proto_init() {
	if File_common_net_network_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_net_network_proto_rawDesc), len(file_common_net_network_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_net_network_proto_goTypes,
		DependencyIndexes: file_common_net_network_proto_depIdxs,
		EnumInfos:         file_common_net_network_proto_enumTypes,
		MessageInfos:      file_common_net_network_proto_msgTypes,
	}.Build()
	File_common_net_network_proto = out.File
	file_common_net_network_proto_goTypes = nil
	file_common_net_network_proto_depIdxs = nil
}
