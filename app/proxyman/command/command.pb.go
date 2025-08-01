// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: app/proxyman/command/command.proto

package command

import (
	protocol "github.com/NamiraNet/xray-core/common/protocol"
	serial "github.com/NamiraNet/xray-core/common/serial"
	core "github.com/NamiraNet/xray-core/core"
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

type AddUserOperation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *protocol.User         `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddUserOperation) Reset() {
	*x = AddUserOperation{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddUserOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserOperation) ProtoMessage() {}

func (x *AddUserOperation) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserOperation.ProtoReflect.Descriptor instead.
func (*AddUserOperation) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserOperation) GetUser() *protocol.User {
	if x != nil {
		return x.User
	}
	return nil
}

type RemoveUserOperation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveUserOperation) Reset() {
	*x = RemoveUserOperation{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveUserOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserOperation) ProtoMessage() {}

func (x *RemoveUserOperation) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserOperation.ProtoReflect.Descriptor instead.
func (*RemoveUserOperation) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveUserOperation) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AddInboundRequest struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Inbound       *core.InboundHandlerConfig `protobuf:"bytes,1,opt,name=inbound,proto3" json:"inbound,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddInboundRequest) Reset() {
	*x = AddInboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInboundRequest) ProtoMessage() {}

func (x *AddInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInboundRequest.ProtoReflect.Descriptor instead.
func (*AddInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{2}
}

func (x *AddInboundRequest) GetInbound() *core.InboundHandlerConfig {
	if x != nil {
		return x.Inbound
	}
	return nil
}

type AddInboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddInboundResponse) Reset() {
	*x = AddInboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInboundResponse) ProtoMessage() {}

func (x *AddInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInboundResponse.ProtoReflect.Descriptor instead.
func (*AddInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{3}
}

type RemoveInboundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveInboundRequest) Reset() {
	*x = RemoveInboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInboundRequest) ProtoMessage() {}

func (x *RemoveInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInboundRequest.ProtoReflect.Descriptor instead.
func (*RemoveInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveInboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type RemoveInboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveInboundResponse) Reset() {
	*x = RemoveInboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInboundResponse) ProtoMessage() {}

func (x *RemoveInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInboundResponse.ProtoReflect.Descriptor instead.
func (*RemoveInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{5}
}

type AlterInboundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Operation     *serial.TypedMessage   `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterInboundRequest) Reset() {
	*x = AlterInboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterInboundRequest) ProtoMessage() {}

func (x *AlterInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterInboundRequest.ProtoReflect.Descriptor instead.
func (*AlterInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{6}
}

func (x *AlterInboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AlterInboundRequest) GetOperation() *serial.TypedMessage {
	if x != nil {
		return x.Operation
	}
	return nil
}

type AlterInboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterInboundResponse) Reset() {
	*x = AlterInboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterInboundResponse) ProtoMessage() {}

func (x *AlterInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterInboundResponse.ProtoReflect.Descriptor instead.
func (*AlterInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{7}
}

type ListInboundsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsOnlyTags    bool                   `protobuf:"varint,1,opt,name=isOnlyTags,proto3" json:"isOnlyTags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInboundsRequest) Reset() {
	*x = ListInboundsRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInboundsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInboundsRequest) ProtoMessage() {}

func (x *ListInboundsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInboundsRequest.ProtoReflect.Descriptor instead.
func (*ListInboundsRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{8}
}

func (x *ListInboundsRequest) GetIsOnlyTags() bool {
	if x != nil {
		return x.IsOnlyTags
	}
	return false
}

type ListInboundsResponse struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Inbounds      []*core.InboundHandlerConfig `protobuf:"bytes,1,rep,name=inbounds,proto3" json:"inbounds,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInboundsResponse) Reset() {
	*x = ListInboundsResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInboundsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInboundsResponse) ProtoMessage() {}

func (x *ListInboundsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInboundsResponse.ProtoReflect.Descriptor instead.
func (*ListInboundsResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{9}
}

func (x *ListInboundsResponse) GetInbounds() []*core.InboundHandlerConfig {
	if x != nil {
		return x.Inbounds
	}
	return nil
}

type GetInboundUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInboundUserRequest) Reset() {
	*x = GetInboundUserRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInboundUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboundUserRequest) ProtoMessage() {}

func (x *GetInboundUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboundUserRequest.ProtoReflect.Descriptor instead.
func (*GetInboundUserRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{10}
}

func (x *GetInboundUserRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *GetInboundUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetInboundUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*protocol.User       `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInboundUserResponse) Reset() {
	*x = GetInboundUserResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInboundUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboundUserResponse) ProtoMessage() {}

func (x *GetInboundUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboundUserResponse.ProtoReflect.Descriptor instead.
func (*GetInboundUserResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{11}
}

func (x *GetInboundUserResponse) GetUsers() []*protocol.User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetInboundUsersCountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInboundUsersCountResponse) Reset() {
	*x = GetInboundUsersCountResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInboundUsersCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboundUsersCountResponse) ProtoMessage() {}

func (x *GetInboundUsersCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboundUsersCountResponse.ProtoReflect.Descriptor instead.
func (*GetInboundUsersCountResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{12}
}

func (x *GetInboundUsersCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddOutboundRequest struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Outbound      *core.OutboundHandlerConfig `protobuf:"bytes,1,opt,name=outbound,proto3" json:"outbound,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddOutboundRequest) Reset() {
	*x = AddOutboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOutboundRequest) ProtoMessage() {}

func (x *AddOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOutboundRequest.ProtoReflect.Descriptor instead.
func (*AddOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{13}
}

func (x *AddOutboundRequest) GetOutbound() *core.OutboundHandlerConfig {
	if x != nil {
		return x.Outbound
	}
	return nil
}

type AddOutboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddOutboundResponse) Reset() {
	*x = AddOutboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOutboundResponse) ProtoMessage() {}

func (x *AddOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOutboundResponse.ProtoReflect.Descriptor instead.
func (*AddOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{14}
}

type RemoveOutboundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveOutboundRequest) Reset() {
	*x = RemoveOutboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOutboundRequest) ProtoMessage() {}

func (x *RemoveOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOutboundRequest.ProtoReflect.Descriptor instead.
func (*RemoveOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{15}
}

func (x *RemoveOutboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type RemoveOutboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveOutboundResponse) Reset() {
	*x = RemoveOutboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOutboundResponse) ProtoMessage() {}

func (x *RemoveOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOutboundResponse.ProtoReflect.Descriptor instead.
func (*RemoveOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{16}
}

type AlterOutboundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Operation     *serial.TypedMessage   `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterOutboundRequest) Reset() {
	*x = AlterOutboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterOutboundRequest) ProtoMessage() {}

func (x *AlterOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterOutboundRequest.ProtoReflect.Descriptor instead.
func (*AlterOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{17}
}

func (x *AlterOutboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AlterOutboundRequest) GetOperation() *serial.TypedMessage {
	if x != nil {
		return x.Operation
	}
	return nil
}

type AlterOutboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterOutboundResponse) Reset() {
	*x = AlterOutboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterOutboundResponse) ProtoMessage() {}

func (x *AlterOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterOutboundResponse.ProtoReflect.Descriptor instead.
func (*AlterOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{18}
}

type ListOutboundsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOutboundsRequest) Reset() {
	*x = ListOutboundsRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[19]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOutboundsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOutboundsRequest) ProtoMessage() {}

func (x *ListOutboundsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[19]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOutboundsRequest.ProtoReflect.Descriptor instead.
func (*ListOutboundsRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{19}
}

type ListOutboundsResponse struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Outbounds     []*core.OutboundHandlerConfig `protobuf:"bytes,1,rep,name=outbounds,proto3" json:"outbounds,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOutboundsResponse) Reset() {
	*x = ListOutboundsResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[20]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOutboundsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOutboundsResponse) ProtoMessage() {}

func (x *ListOutboundsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[20]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOutboundsResponse.ProtoReflect.Descriptor instead.
func (*ListOutboundsResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{20}
}

func (x *ListOutboundsResponse) GetOutbounds() []*core.OutboundHandlerConfig {
	if x != nil {
		return x.Outbounds
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[21]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[21]
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
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{21}
}

var File_app_proxyman_command_command_proto protoreflect.FileDescriptor

const file_app_proxyman_command_command_proto_rawDesc = "" +
	"\n" +
	"\"app/proxyman/command/command.proto\x12\x19xray.app.proxyman.command\x1a\x1acommon/protocol/user.proto\x1a!common/serial/typed_message.proto\x1a\x11core/config.proto\"B\n" +
	"\x10AddUserOperation\x12.\n" +
	"\x04user\x18\x01 \x01(\v2\x1a.xray.common.protocol.UserR\x04user\"+\n" +
	"\x13RemoveUserOperation\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\"N\n" +
	"\x11AddInboundRequest\x129\n" +
	"\ainbound\x18\x01 \x01(\v2\x1f.xray.core.InboundHandlerConfigR\ainbound\"\x14\n" +
	"\x12AddInboundResponse\"(\n" +
	"\x14RemoveInboundRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\"\x17\n" +
	"\x15RemoveInboundResponse\"g\n" +
	"\x13AlterInboundRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\x12>\n" +
	"\toperation\x18\x02 \x01(\v2 .xray.common.serial.TypedMessageR\toperation\"\x16\n" +
	"\x14AlterInboundResponse\"5\n" +
	"\x13ListInboundsRequest\x12\x1e\n" +
	"\n" +
	"isOnlyTags\x18\x01 \x01(\bR\n" +
	"isOnlyTags\"S\n" +
	"\x14ListInboundsResponse\x12;\n" +
	"\binbounds\x18\x01 \x03(\v2\x1f.xray.core.InboundHandlerConfigR\binbounds\"?\n" +
	"\x15GetInboundUserRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\"J\n" +
	"\x16GetInboundUserResponse\x120\n" +
	"\x05users\x18\x01 \x03(\v2\x1a.xray.common.protocol.UserR\x05users\"4\n" +
	"\x1cGetInboundUsersCountResponse\x12\x14\n" +
	"\x05count\x18\x01 \x01(\x03R\x05count\"R\n" +
	"\x12AddOutboundRequest\x12<\n" +
	"\boutbound\x18\x01 \x01(\v2 .xray.core.OutboundHandlerConfigR\boutbound\"\x15\n" +
	"\x13AddOutboundResponse\")\n" +
	"\x15RemoveOutboundRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\"\x18\n" +
	"\x16RemoveOutboundResponse\"h\n" +
	"\x14AlterOutboundRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\x12>\n" +
	"\toperation\x18\x02 \x01(\v2 .xray.common.serial.TypedMessageR\toperation\"\x17\n" +
	"\x15AlterOutboundResponse\"\x16\n" +
	"\x14ListOutboundsRequest\"W\n" +
	"\x15ListOutboundsResponse\x12>\n" +
	"\toutbounds\x18\x01 \x03(\v2 .xray.core.OutboundHandlerConfigR\toutbounds\"\b\n" +
	"\x06Config2\xae\t\n" +
	"\x0eHandlerService\x12k\n" +
	"\n" +
	"AddInbound\x12,.xray.app.proxyman.command.AddInboundRequest\x1a-.xray.app.proxyman.command.AddInboundResponse\"\x00\x12t\n" +
	"\rRemoveInbound\x12/.xray.app.proxyman.command.RemoveInboundRequest\x1a0.xray.app.proxyman.command.RemoveInboundResponse\"\x00\x12q\n" +
	"\fAlterInbound\x12..xray.app.proxyman.command.AlterInboundRequest\x1a/.xray.app.proxyman.command.AlterInboundResponse\"\x00\x12q\n" +
	"\fListInbounds\x12..xray.app.proxyman.command.ListInboundsRequest\x1a/.xray.app.proxyman.command.ListInboundsResponse\"\x00\x12x\n" +
	"\x0fGetInboundUsers\x120.xray.app.proxyman.command.GetInboundUserRequest\x1a1.xray.app.proxyman.command.GetInboundUserResponse\"\x00\x12\x83\x01\n" +
	"\x14GetInboundUsersCount\x120.xray.app.proxyman.command.GetInboundUserRequest\x1a7.xray.app.proxyman.command.GetInboundUsersCountResponse\"\x00\x12n\n" +
	"\vAddOutbound\x12-.xray.app.proxyman.command.AddOutboundRequest\x1a..xray.app.proxyman.command.AddOutboundResponse\"\x00\x12w\n" +
	"\x0eRemoveOutbound\x120.xray.app.proxyman.command.RemoveOutboundRequest\x1a1.xray.app.proxyman.command.RemoveOutboundResponse\"\x00\x12t\n" +
	"\rAlterOutbound\x12/.xray.app.proxyman.command.AlterOutboundRequest\x1a0.xray.app.proxyman.command.AlterOutboundResponse\"\x00\x12t\n" +
	"\rListOutbounds\x12/.xray.app.proxyman.command.ListOutboundsRequest\x1a0.xray.app.proxyman.command.ListOutboundsResponse\"\x00Br\n" +
	"\x1dcom.xray.app.proxyman.commandP\x01Z3github.com/NamiraNet/xray-core/app/proxyman/command\xaa\x02\x19Xray.App.Proxyman.Commandb\x06proto3"

var (
	file_app_proxyman_command_command_proto_rawDescOnce sync.Once
	file_app_proxyman_command_command_proto_rawDescData []byte
)

func file_app_proxyman_command_command_proto_rawDescGZIP() []byte {
	file_app_proxyman_command_command_proto_rawDescOnce.Do(func() {
		file_app_proxyman_command_command_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_proxyman_command_command_proto_rawDesc), len(file_app_proxyman_command_command_proto_rawDesc)))
	})
	return file_app_proxyman_command_command_proto_rawDescData
}

var file_app_proxyman_command_command_proto_msgTypes = make([]protoimpl.MessageInfo, 22)
var file_app_proxyman_command_command_proto_goTypes = []any{
	(*AddUserOperation)(nil),             // 0: xray.app.proxyman.command.AddUserOperation
	(*RemoveUserOperation)(nil),          // 1: xray.app.proxyman.command.RemoveUserOperation
	(*AddInboundRequest)(nil),            // 2: xray.app.proxyman.command.AddInboundRequest
	(*AddInboundResponse)(nil),           // 3: xray.app.proxyman.command.AddInboundResponse
	(*RemoveInboundRequest)(nil),         // 4: xray.app.proxyman.command.RemoveInboundRequest
	(*RemoveInboundResponse)(nil),        // 5: xray.app.proxyman.command.RemoveInboundResponse
	(*AlterInboundRequest)(nil),          // 6: xray.app.proxyman.command.AlterInboundRequest
	(*AlterInboundResponse)(nil),         // 7: xray.app.proxyman.command.AlterInboundResponse
	(*ListInboundsRequest)(nil),          // 8: xray.app.proxyman.command.ListInboundsRequest
	(*ListInboundsResponse)(nil),         // 9: xray.app.proxyman.command.ListInboundsResponse
	(*GetInboundUserRequest)(nil),        // 10: xray.app.proxyman.command.GetInboundUserRequest
	(*GetInboundUserResponse)(nil),       // 11: xray.app.proxyman.command.GetInboundUserResponse
	(*GetInboundUsersCountResponse)(nil), // 12: xray.app.proxyman.command.GetInboundUsersCountResponse
	(*AddOutboundRequest)(nil),           // 13: xray.app.proxyman.command.AddOutboundRequest
	(*AddOutboundResponse)(nil),          // 14: xray.app.proxyman.command.AddOutboundResponse
	(*RemoveOutboundRequest)(nil),        // 15: xray.app.proxyman.command.RemoveOutboundRequest
	(*RemoveOutboundResponse)(nil),       // 16: xray.app.proxyman.command.RemoveOutboundResponse
	(*AlterOutboundRequest)(nil),         // 17: xray.app.proxyman.command.AlterOutboundRequest
	(*AlterOutboundResponse)(nil),        // 18: xray.app.proxyman.command.AlterOutboundResponse
	(*ListOutboundsRequest)(nil),         // 19: xray.app.proxyman.command.ListOutboundsRequest
	(*ListOutboundsResponse)(nil),        // 20: xray.app.proxyman.command.ListOutboundsResponse
	(*Config)(nil),                       // 21: xray.app.proxyman.command.Config
	(*protocol.User)(nil),                // 22: xray.common.protocol.User
	(*core.InboundHandlerConfig)(nil),    // 23: xray.core.InboundHandlerConfig
	(*serial.TypedMessage)(nil),          // 24: xray.common.serial.TypedMessage
	(*core.OutboundHandlerConfig)(nil),   // 25: xray.core.OutboundHandlerConfig
}
var file_app_proxyman_command_command_proto_depIdxs = []int32{
	22, // 0: xray.app.proxyman.command.AddUserOperation.user:type_name -> xray.common.protocol.User
	23, // 1: xray.app.proxyman.command.AddInboundRequest.inbound:type_name -> xray.core.InboundHandlerConfig
	24, // 2: xray.app.proxyman.command.AlterInboundRequest.operation:type_name -> xray.common.serial.TypedMessage
	23, // 3: xray.app.proxyman.command.ListInboundsResponse.inbounds:type_name -> xray.core.InboundHandlerConfig
	22, // 4: xray.app.proxyman.command.GetInboundUserResponse.users:type_name -> xray.common.protocol.User
	25, // 5: xray.app.proxyman.command.AddOutboundRequest.outbound:type_name -> xray.core.OutboundHandlerConfig
	24, // 6: xray.app.proxyman.command.AlterOutboundRequest.operation:type_name -> xray.common.serial.TypedMessage
	25, // 7: xray.app.proxyman.command.ListOutboundsResponse.outbounds:type_name -> xray.core.OutboundHandlerConfig
	2,  // 8: xray.app.proxyman.command.HandlerService.AddInbound:input_type -> xray.app.proxyman.command.AddInboundRequest
	4,  // 9: xray.app.proxyman.command.HandlerService.RemoveInbound:input_type -> xray.app.proxyman.command.RemoveInboundRequest
	6,  // 10: xray.app.proxyman.command.HandlerService.AlterInbound:input_type -> xray.app.proxyman.command.AlterInboundRequest
	8,  // 11: xray.app.proxyman.command.HandlerService.ListInbounds:input_type -> xray.app.proxyman.command.ListInboundsRequest
	10, // 12: xray.app.proxyman.command.HandlerService.GetInboundUsers:input_type -> xray.app.proxyman.command.GetInboundUserRequest
	10, // 13: xray.app.proxyman.command.HandlerService.GetInboundUsersCount:input_type -> xray.app.proxyman.command.GetInboundUserRequest
	13, // 14: xray.app.proxyman.command.HandlerService.AddOutbound:input_type -> xray.app.proxyman.command.AddOutboundRequest
	15, // 15: xray.app.proxyman.command.HandlerService.RemoveOutbound:input_type -> xray.app.proxyman.command.RemoveOutboundRequest
	17, // 16: xray.app.proxyman.command.HandlerService.AlterOutbound:input_type -> xray.app.proxyman.command.AlterOutboundRequest
	19, // 17: xray.app.proxyman.command.HandlerService.ListOutbounds:input_type -> xray.app.proxyman.command.ListOutboundsRequest
	3,  // 18: xray.app.proxyman.command.HandlerService.AddInbound:output_type -> xray.app.proxyman.command.AddInboundResponse
	5,  // 19: xray.app.proxyman.command.HandlerService.RemoveInbound:output_type -> xray.app.proxyman.command.RemoveInboundResponse
	7,  // 20: xray.app.proxyman.command.HandlerService.AlterInbound:output_type -> xray.app.proxyman.command.AlterInboundResponse
	9,  // 21: xray.app.proxyman.command.HandlerService.ListInbounds:output_type -> xray.app.proxyman.command.ListInboundsResponse
	11, // 22: xray.app.proxyman.command.HandlerService.GetInboundUsers:output_type -> xray.app.proxyman.command.GetInboundUserResponse
	12, // 23: xray.app.proxyman.command.HandlerService.GetInboundUsersCount:output_type -> xray.app.proxyman.command.GetInboundUsersCountResponse
	14, // 24: xray.app.proxyman.command.HandlerService.AddOutbound:output_type -> xray.app.proxyman.command.AddOutboundResponse
	16, // 25: xray.app.proxyman.command.HandlerService.RemoveOutbound:output_type -> xray.app.proxyman.command.RemoveOutboundResponse
	18, // 26: xray.app.proxyman.command.HandlerService.AlterOutbound:output_type -> xray.app.proxyman.command.AlterOutboundResponse
	20, // 27: xray.app.proxyman.command.HandlerService.ListOutbounds:output_type -> xray.app.proxyman.command.ListOutboundsResponse
	18, // [18:28] is the sub-list for method output_type
	8,  // [8:18] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_app_proxyman_command_command_proto_init() }
func file_app_proxyman_command_command_proto_init() {
	if File_app_proxyman_command_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_proxyman_command_command_proto_rawDesc), len(file_app_proxyman_command_command_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   22,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_proxyman_command_command_proto_goTypes,
		DependencyIndexes: file_app_proxyman_command_command_proto_depIdxs,
		MessageInfos:      file_app_proxyman_command_command_proto_msgTypes,
	}.Build()
	File_app_proxyman_command_command_proto = out.File
	file_app_proxyman_command_command_proto_goTypes = nil
	file_app_proxyman_command_command_proto_depIdxs = nil
}
