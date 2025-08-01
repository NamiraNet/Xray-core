// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: app/router/command/command.proto

package command

import (
	net "github.com/NamiraNet/xray-core/common/net"
	serial "github.com/NamiraNet/xray-core/common/serial"
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

// RoutingContext is the context with information relative to routing process.
// It conforms to the structure of xray.features.routing.Context and
// xray.features.routing.Route.
type RoutingContext struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	InboundTag        string                 `protobuf:"bytes,1,opt,name=InboundTag,proto3" json:"InboundTag,omitempty"`
	Network           net.Network            `protobuf:"varint,2,opt,name=Network,proto3,enum=xray.common.net.Network" json:"Network,omitempty"`
	SourceIPs         [][]byte               `protobuf:"bytes,3,rep,name=SourceIPs,proto3" json:"SourceIPs,omitempty"`
	TargetIPs         [][]byte               `protobuf:"bytes,4,rep,name=TargetIPs,proto3" json:"TargetIPs,omitempty"`
	SourcePort        uint32                 `protobuf:"varint,5,opt,name=SourcePort,proto3" json:"SourcePort,omitempty"`
	TargetPort        uint32                 `protobuf:"varint,6,opt,name=TargetPort,proto3" json:"TargetPort,omitempty"`
	TargetDomain      string                 `protobuf:"bytes,7,opt,name=TargetDomain,proto3" json:"TargetDomain,omitempty"`
	Protocol          string                 `protobuf:"bytes,8,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	User              string                 `protobuf:"bytes,9,opt,name=User,proto3" json:"User,omitempty"`
	Attributes        map[string]string      `protobuf:"bytes,10,rep,name=Attributes,proto3" json:"Attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	OutboundGroupTags []string               `protobuf:"bytes,11,rep,name=OutboundGroupTags,proto3" json:"OutboundGroupTags,omitempty"`
	OutboundTag       string                 `protobuf:"bytes,12,opt,name=OutboundTag,proto3" json:"OutboundTag,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *RoutingContext) Reset() {
	*x = RoutingContext{}
	mi := &file_app_router_command_command_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoutingContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingContext) ProtoMessage() {}

func (x *RoutingContext) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingContext.ProtoReflect.Descriptor instead.
func (*RoutingContext) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{0}
}

func (x *RoutingContext) GetInboundTag() string {
	if x != nil {
		return x.InboundTag
	}
	return ""
}

func (x *RoutingContext) GetNetwork() net.Network {
	if x != nil {
		return x.Network
	}
	return net.Network(0)
}

func (x *RoutingContext) GetSourceIPs() [][]byte {
	if x != nil {
		return x.SourceIPs
	}
	return nil
}

func (x *RoutingContext) GetTargetIPs() [][]byte {
	if x != nil {
		return x.TargetIPs
	}
	return nil
}

func (x *RoutingContext) GetSourcePort() uint32 {
	if x != nil {
		return x.SourcePort
	}
	return 0
}

func (x *RoutingContext) GetTargetPort() uint32 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

func (x *RoutingContext) GetTargetDomain() string {
	if x != nil {
		return x.TargetDomain
	}
	return ""
}

func (x *RoutingContext) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *RoutingContext) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RoutingContext) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *RoutingContext) GetOutboundGroupTags() []string {
	if x != nil {
		return x.OutboundGroupTags
	}
	return nil
}

func (x *RoutingContext) GetOutboundTag() string {
	if x != nil {
		return x.OutboundTag
	}
	return ""
}

// SubscribeRoutingStatsRequest subscribes to routing statistics channel if
// opened by xray-core.
// * FieldSelectors selects a subset of fields in routing statistics to return.
// Valid selectors:
//   - inbound: Selects connection's inbound tag.
//   - network: Selects connection's network.
//   - ip: Equivalent as "ip_source" and "ip_target", selects both source and
//     target IP.
//   - port: Equivalent as "port_source" and "port_target", selects both source
//     and target port.
//   - domain: Selects target domain.
//   - protocol: Select connection's protocol.
//   - user: Select connection's inbound user email.
//   - attributes: Select connection's additional attributes.
//   - outbound: Equivalent as "outbound" and "outbound_group", select both
//     outbound tag and outbound group tags.
//
// * If FieldSelectors is left empty, all fields will be returned.
type SubscribeRoutingStatsRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	FieldSelectors []string               `protobuf:"bytes,1,rep,name=FieldSelectors,proto3" json:"FieldSelectors,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SubscribeRoutingStatsRequest) Reset() {
	*x = SubscribeRoutingStatsRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRoutingStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRoutingStatsRequest) ProtoMessage() {}

func (x *SubscribeRoutingStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRoutingStatsRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRoutingStatsRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeRoutingStatsRequest) GetFieldSelectors() []string {
	if x != nil {
		return x.FieldSelectors
	}
	return nil
}

// TestRouteRequest manually tests a routing result according to the routing
// context message.
// * RoutingContext is the routing message without outbound information.
// * FieldSelectors selects the fields to return in the routing result. All
// fields are returned if left empty.
// * PublishResult broadcasts the routing result to routing statistics channel
// if set true.
type TestRouteRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	RoutingContext *RoutingContext        `protobuf:"bytes,1,opt,name=RoutingContext,proto3" json:"RoutingContext,omitempty"`
	FieldSelectors []string               `protobuf:"bytes,2,rep,name=FieldSelectors,proto3" json:"FieldSelectors,omitempty"`
	PublishResult  bool                   `protobuf:"varint,3,opt,name=PublishResult,proto3" json:"PublishResult,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TestRouteRequest) Reset() {
	*x = TestRouteRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRouteRequest) ProtoMessage() {}

func (x *TestRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRouteRequest.ProtoReflect.Descriptor instead.
func (*TestRouteRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{2}
}

func (x *TestRouteRequest) GetRoutingContext() *RoutingContext {
	if x != nil {
		return x.RoutingContext
	}
	return nil
}

func (x *TestRouteRequest) GetFieldSelectors() []string {
	if x != nil {
		return x.FieldSelectors
	}
	return nil
}

func (x *TestRouteRequest) GetPublishResult() bool {
	if x != nil {
		return x.PublishResult
	}
	return false
}

type PrincipleTargetInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           []string               `protobuf:"bytes,1,rep,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrincipleTargetInfo) Reset() {
	*x = PrincipleTargetInfo{}
	mi := &file_app_router_command_command_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrincipleTargetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipleTargetInfo) ProtoMessage() {}

func (x *PrincipleTargetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipleTargetInfo.ProtoReflect.Descriptor instead.
func (*PrincipleTargetInfo) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{3}
}

func (x *PrincipleTargetInfo) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

type OverrideInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        string                 `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverrideInfo) Reset() {
	*x = OverrideInfo{}
	mi := &file_app_router_command_command_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverrideInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideInfo) ProtoMessage() {}

func (x *OverrideInfo) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideInfo.ProtoReflect.Descriptor instead.
func (*OverrideInfo) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{4}
}

func (x *OverrideInfo) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type BalancerMsg struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Override        *OverrideInfo          `protobuf:"bytes,5,opt,name=override,proto3" json:"override,omitempty"`
	PrincipleTarget *PrincipleTargetInfo   `protobuf:"bytes,6,opt,name=principle_target,json=principleTarget,proto3" json:"principle_target,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *BalancerMsg) Reset() {
	*x = BalancerMsg{}
	mi := &file_app_router_command_command_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalancerMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerMsg) ProtoMessage() {}

func (x *BalancerMsg) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerMsg.ProtoReflect.Descriptor instead.
func (*BalancerMsg) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{5}
}

func (x *BalancerMsg) GetOverride() *OverrideInfo {
	if x != nil {
		return x.Override
	}
	return nil
}

func (x *BalancerMsg) GetPrincipleTarget() *PrincipleTargetInfo {
	if x != nil {
		return x.PrincipleTarget
	}
	return nil
}

type GetBalancerInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBalancerInfoRequest) Reset() {
	*x = GetBalancerInfoRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBalancerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalancerInfoRequest) ProtoMessage() {}

func (x *GetBalancerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalancerInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBalancerInfoRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{6}
}

func (x *GetBalancerInfoRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type GetBalancerInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Balancer      *BalancerMsg           `protobuf:"bytes,1,opt,name=balancer,proto3" json:"balancer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBalancerInfoResponse) Reset() {
	*x = GetBalancerInfoResponse{}
	mi := &file_app_router_command_command_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBalancerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalancerInfoResponse) ProtoMessage() {}

func (x *GetBalancerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalancerInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBalancerInfoResponse) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{7}
}

func (x *GetBalancerInfoResponse) GetBalancer() *BalancerMsg {
	if x != nil {
		return x.Balancer
	}
	return nil
}

type OverrideBalancerTargetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BalancerTag   string                 `protobuf:"bytes,1,opt,name=balancerTag,proto3" json:"balancerTag,omitempty"`
	Target        string                 `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverrideBalancerTargetRequest) Reset() {
	*x = OverrideBalancerTargetRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverrideBalancerTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideBalancerTargetRequest) ProtoMessage() {}

func (x *OverrideBalancerTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideBalancerTargetRequest.ProtoReflect.Descriptor instead.
func (*OverrideBalancerTargetRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{8}
}

func (x *OverrideBalancerTargetRequest) GetBalancerTag() string {
	if x != nil {
		return x.BalancerTag
	}
	return ""
}

func (x *OverrideBalancerTargetRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type OverrideBalancerTargetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverrideBalancerTargetResponse) Reset() {
	*x = OverrideBalancerTargetResponse{}
	mi := &file_app_router_command_command_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverrideBalancerTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideBalancerTargetResponse) ProtoMessage() {}

func (x *OverrideBalancerTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideBalancerTargetResponse.ProtoReflect.Descriptor instead.
func (*OverrideBalancerTargetResponse) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{9}
}

type AddRuleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Config        *serial.TypedMessage   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	ShouldAppend  bool                   `protobuf:"varint,2,opt,name=shouldAppend,proto3" json:"shouldAppend,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRuleRequest) Reset() {
	*x = AddRuleRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRuleRequest) ProtoMessage() {}

func (x *AddRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRuleRequest.ProtoReflect.Descriptor instead.
func (*AddRuleRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{10}
}

func (x *AddRuleRequest) GetConfig() *serial.TypedMessage {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *AddRuleRequest) GetShouldAppend() bool {
	if x != nil {
		return x.ShouldAppend
	}
	return false
}

type AddRuleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRuleResponse) Reset() {
	*x = AddRuleResponse{}
	mi := &file_app_router_command_command_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRuleResponse) ProtoMessage() {}

func (x *AddRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRuleResponse.ProtoReflect.Descriptor instead.
func (*AddRuleResponse) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{11}
}

type RemoveRuleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RuleTag       string                 `protobuf:"bytes,1,opt,name=ruleTag,proto3" json:"ruleTag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRuleRequest) Reset() {
	*x = RemoveRuleRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRuleRequest) ProtoMessage() {}

func (x *RemoveRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRuleRequest.ProtoReflect.Descriptor instead.
func (*RemoveRuleRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{12}
}

func (x *RemoveRuleRequest) GetRuleTag() string {
	if x != nil {
		return x.RuleTag
	}
	return ""
}

type RemoveRuleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRuleResponse) Reset() {
	*x = RemoveRuleResponse{}
	mi := &file_app_router_command_command_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRuleResponse) ProtoMessage() {}

func (x *RemoveRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRuleResponse.ProtoReflect.Descriptor instead.
func (*RemoveRuleResponse) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{13}
}

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_router_command_command_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[14]
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
	return file_app_router_command_command_proto_rawDescGZIP(), []int{14}
}

var File_app_router_command_command_proto protoreflect.FileDescriptor

const file_app_router_command_command_proto_rawDesc = "" +
	"\n" +
	" app/router/command/command.proto\x12\x17xray.app.router.command\x1a\x18common/net/network.proto\x1a!common/serial/typed_message.proto\"\x9c\x04\n" +
	"\x0eRoutingContext\x12\x1e\n" +
	"\n" +
	"InboundTag\x18\x01 \x01(\tR\n" +
	"InboundTag\x122\n" +
	"\aNetwork\x18\x02 \x01(\x0e2\x18.xray.common.net.NetworkR\aNetwork\x12\x1c\n" +
	"\tSourceIPs\x18\x03 \x03(\fR\tSourceIPs\x12\x1c\n" +
	"\tTargetIPs\x18\x04 \x03(\fR\tTargetIPs\x12\x1e\n" +
	"\n" +
	"SourcePort\x18\x05 \x01(\rR\n" +
	"SourcePort\x12\x1e\n" +
	"\n" +
	"TargetPort\x18\x06 \x01(\rR\n" +
	"TargetPort\x12\"\n" +
	"\fTargetDomain\x18\a \x01(\tR\fTargetDomain\x12\x1a\n" +
	"\bProtocol\x18\b \x01(\tR\bProtocol\x12\x12\n" +
	"\x04User\x18\t \x01(\tR\x04User\x12W\n" +
	"\n" +
	"Attributes\x18\n" +
	" \x03(\v27.xray.app.router.command.RoutingContext.AttributesEntryR\n" +
	"Attributes\x12,\n" +
	"\x11OutboundGroupTags\x18\v \x03(\tR\x11OutboundGroupTags\x12 \n" +
	"\vOutboundTag\x18\f \x01(\tR\vOutboundTag\x1a=\n" +
	"\x0fAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"F\n" +
	"\x1cSubscribeRoutingStatsRequest\x12&\n" +
	"\x0eFieldSelectors\x18\x01 \x03(\tR\x0eFieldSelectors\"\xb1\x01\n" +
	"\x10TestRouteRequest\x12O\n" +
	"\x0eRoutingContext\x18\x01 \x01(\v2'.xray.app.router.command.RoutingContextR\x0eRoutingContext\x12&\n" +
	"\x0eFieldSelectors\x18\x02 \x03(\tR\x0eFieldSelectors\x12$\n" +
	"\rPublishResult\x18\x03 \x01(\bR\rPublishResult\"'\n" +
	"\x13PrincipleTargetInfo\x12\x10\n" +
	"\x03tag\x18\x01 \x03(\tR\x03tag\"&\n" +
	"\fOverrideInfo\x12\x16\n" +
	"\x06target\x18\x02 \x01(\tR\x06target\"\xa9\x01\n" +
	"\vBalancerMsg\x12A\n" +
	"\boverride\x18\x05 \x01(\v2%.xray.app.router.command.OverrideInfoR\boverride\x12W\n" +
	"\x10principle_target\x18\x06 \x01(\v2,.xray.app.router.command.PrincipleTargetInfoR\x0fprincipleTarget\"*\n" +
	"\x16GetBalancerInfoRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\"[\n" +
	"\x17GetBalancerInfoResponse\x12@\n" +
	"\bbalancer\x18\x01 \x01(\v2$.xray.app.router.command.BalancerMsgR\bbalancer\"Y\n" +
	"\x1dOverrideBalancerTargetRequest\x12 \n" +
	"\vbalancerTag\x18\x01 \x01(\tR\vbalancerTag\x12\x16\n" +
	"\x06target\x18\x02 \x01(\tR\x06target\" \n" +
	"\x1eOverrideBalancerTargetResponse\"n\n" +
	"\x0eAddRuleRequest\x128\n" +
	"\x06config\x18\x01 \x01(\v2 .xray.common.serial.TypedMessageR\x06config\x12\"\n" +
	"\fshouldAppend\x18\x02 \x01(\bR\fshouldAppend\"\x11\n" +
	"\x0fAddRuleResponse\"-\n" +
	"\x11RemoveRuleRequest\x12\x18\n" +
	"\aruleTag\x18\x01 \x01(\tR\aruleTag\"\x14\n" +
	"\x12RemoveRuleResponse\"\b\n" +
	"\x06Config2\xbf\x05\n" +
	"\x0eRoutingService\x12{\n" +
	"\x15SubscribeRoutingStats\x125.xray.app.router.command.SubscribeRoutingStatsRequest\x1a'.xray.app.router.command.RoutingContext\"\x000\x01\x12a\n" +
	"\tTestRoute\x12).xray.app.router.command.TestRouteRequest\x1a'.xray.app.router.command.RoutingContext\"\x00\x12v\n" +
	"\x0fGetBalancerInfo\x12/.xray.app.router.command.GetBalancerInfoRequest\x1a0.xray.app.router.command.GetBalancerInfoResponse\"\x00\x12\x8b\x01\n" +
	"\x16OverrideBalancerTarget\x126.xray.app.router.command.OverrideBalancerTargetRequest\x1a7.xray.app.router.command.OverrideBalancerTargetResponse\"\x00\x12^\n" +
	"\aAddRule\x12'.xray.app.router.command.AddRuleRequest\x1a(.xray.app.router.command.AddRuleResponse\"\x00\x12g\n" +
	"\n" +
	"RemoveRule\x12*.xray.app.router.command.RemoveRuleRequest\x1a+.xray.app.router.command.RemoveRuleResponse\"\x00Bl\n" +
	"\x1bcom.xray.app.router.commandP\x01Z1github.com/NamiraNet/xray-core/app/router/command\xaa\x02\x17Xray.App.Router.Commandb\x06proto3"

var (
	file_app_router_command_command_proto_rawDescOnce sync.Once
	file_app_router_command_command_proto_rawDescData []byte
)

func file_app_router_command_command_proto_rawDescGZIP() []byte {
	file_app_router_command_command_proto_rawDescOnce.Do(func() {
		file_app_router_command_command_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_router_command_command_proto_rawDesc), len(file_app_router_command_command_proto_rawDesc)))
	})
	return file_app_router_command_command_proto_rawDescData
}

var file_app_router_command_command_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_app_router_command_command_proto_goTypes = []any{
	(*RoutingContext)(nil),                 // 0: xray.app.router.command.RoutingContext
	(*SubscribeRoutingStatsRequest)(nil),   // 1: xray.app.router.command.SubscribeRoutingStatsRequest
	(*TestRouteRequest)(nil),               // 2: xray.app.router.command.TestRouteRequest
	(*PrincipleTargetInfo)(nil),            // 3: xray.app.router.command.PrincipleTargetInfo
	(*OverrideInfo)(nil),                   // 4: xray.app.router.command.OverrideInfo
	(*BalancerMsg)(nil),                    // 5: xray.app.router.command.BalancerMsg
	(*GetBalancerInfoRequest)(nil),         // 6: xray.app.router.command.GetBalancerInfoRequest
	(*GetBalancerInfoResponse)(nil),        // 7: xray.app.router.command.GetBalancerInfoResponse
	(*OverrideBalancerTargetRequest)(nil),  // 8: xray.app.router.command.OverrideBalancerTargetRequest
	(*OverrideBalancerTargetResponse)(nil), // 9: xray.app.router.command.OverrideBalancerTargetResponse
	(*AddRuleRequest)(nil),                 // 10: xray.app.router.command.AddRuleRequest
	(*AddRuleResponse)(nil),                // 11: xray.app.router.command.AddRuleResponse
	(*RemoveRuleRequest)(nil),              // 12: xray.app.router.command.RemoveRuleRequest
	(*RemoveRuleResponse)(nil),             // 13: xray.app.router.command.RemoveRuleResponse
	(*Config)(nil),                         // 14: xray.app.router.command.Config
	nil,                                    // 15: xray.app.router.command.RoutingContext.AttributesEntry
	(net.Network)(0),                       // 16: xray.common.net.Network
	(*serial.TypedMessage)(nil),            // 17: xray.common.serial.TypedMessage
}
var file_app_router_command_command_proto_depIdxs = []int32{
	16, // 0: xray.app.router.command.RoutingContext.Network:type_name -> xray.common.net.Network
	15, // 1: xray.app.router.command.RoutingContext.Attributes:type_name -> xray.app.router.command.RoutingContext.AttributesEntry
	0,  // 2: xray.app.router.command.TestRouteRequest.RoutingContext:type_name -> xray.app.router.command.RoutingContext
	4,  // 3: xray.app.router.command.BalancerMsg.override:type_name -> xray.app.router.command.OverrideInfo
	3,  // 4: xray.app.router.command.BalancerMsg.principle_target:type_name -> xray.app.router.command.PrincipleTargetInfo
	5,  // 5: xray.app.router.command.GetBalancerInfoResponse.balancer:type_name -> xray.app.router.command.BalancerMsg
	17, // 6: xray.app.router.command.AddRuleRequest.config:type_name -> xray.common.serial.TypedMessage
	1,  // 7: xray.app.router.command.RoutingService.SubscribeRoutingStats:input_type -> xray.app.router.command.SubscribeRoutingStatsRequest
	2,  // 8: xray.app.router.command.RoutingService.TestRoute:input_type -> xray.app.router.command.TestRouteRequest
	6,  // 9: xray.app.router.command.RoutingService.GetBalancerInfo:input_type -> xray.app.router.command.GetBalancerInfoRequest
	8,  // 10: xray.app.router.command.RoutingService.OverrideBalancerTarget:input_type -> xray.app.router.command.OverrideBalancerTargetRequest
	10, // 11: xray.app.router.command.RoutingService.AddRule:input_type -> xray.app.router.command.AddRuleRequest
	12, // 12: xray.app.router.command.RoutingService.RemoveRule:input_type -> xray.app.router.command.RemoveRuleRequest
	0,  // 13: xray.app.router.command.RoutingService.SubscribeRoutingStats:output_type -> xray.app.router.command.RoutingContext
	0,  // 14: xray.app.router.command.RoutingService.TestRoute:output_type -> xray.app.router.command.RoutingContext
	7,  // 15: xray.app.router.command.RoutingService.GetBalancerInfo:output_type -> xray.app.router.command.GetBalancerInfoResponse
	9,  // 16: xray.app.router.command.RoutingService.OverrideBalancerTarget:output_type -> xray.app.router.command.OverrideBalancerTargetResponse
	11, // 17: xray.app.router.command.RoutingService.AddRule:output_type -> xray.app.router.command.AddRuleResponse
	13, // 18: xray.app.router.command.RoutingService.RemoveRule:output_type -> xray.app.router.command.RemoveRuleResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_app_router_command_command_proto_init() }
func file_app_router_command_command_proto_init() {
	if File_app_router_command_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_router_command_command_proto_rawDesc), len(file_app_router_command_command_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_router_command_command_proto_goTypes,
		DependencyIndexes: file_app_router_command_command_proto_depIdxs,
		MessageInfos:      file_app_router_command_command_proto_msgTypes,
	}.Build()
	File_app_router_command_command_proto = out.File
	file_app_router_command_command_proto_goTypes = nil
	file_app_router_command_command_proto_depIdxs = nil
}
