// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/infra_proxy/request/nodes.proto

package request

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *Nodes) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Nodes) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Node) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AffectedNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Chef object type (e.g. 'cookbooks', 'roles', 'chef_environment').
	ChefType string `protobuf:"bytes,3,opt,name=chef_type,json=chefType,proto3" json:"chef_type,omitempty"`
	// Chef object name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Chef object version.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *AffectedNodes) Reset() {
	*x = AffectedNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AffectedNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AffectedNodes) ProtoMessage() {}

func (x *AffectedNodes) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AffectedNodes.ProtoReflect.Descriptor instead.
func (*AffectedNodes) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{2}
}

func (x *AffectedNodes) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *AffectedNodes) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *AffectedNodes) GetChefType() string {
	if x != nil {
		return x.ChefType
	}
	return ""
}

func (x *AffectedNodes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AffectedNodes) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type DeleteNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteNode) Reset() {
	*x = DeleteNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNode) ProtoMessage() {}

func (x *DeleteNode) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNode.ProtoReflect.Descriptor instead.
func (*DeleteNode) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNode) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *DeleteNode) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *DeleteNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Node environment.
	Environment string `protobuf:"bytes,4,opt,name=environment,proto3" json:"environment,omitempty"`
	// Node policy name.
	PolicyName string `protobuf:"bytes,5,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// Node policy group.
	PolicyGroup string `protobuf:"bytes,6,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	// Node run-list.
	RunList []string `protobuf:"bytes,7,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	// Node automatic attributes JSON.
	AutomaticAttributes *_struct.Struct `protobuf:"bytes,8,opt,name=automatic_attributes,json=automaticAttributes,proto3" json:"automatic_attributes,omitempty"`
	// Node default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,9,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Node normal attributes JSON.
	NormalAttributes *_struct.Struct `protobuf:"bytes,10,opt,name=normal_attributes,json=normalAttributes,proto3" json:"normal_attributes,omitempty"`
	// Node override attributes JSON.
	OverrideAttributes *_struct.Struct `protobuf:"bytes,11,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
}

func (x *UpdateNode) Reset() {
	*x = UpdateNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNode) ProtoMessage() {}

func (x *UpdateNode) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNode.ProtoReflect.Descriptor instead.
func (*UpdateNode) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateNode) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *UpdateNode) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *UpdateNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNode) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *UpdateNode) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *UpdateNode) GetPolicyGroup() string {
	if x != nil {
		return x.PolicyGroup
	}
	return ""
}

func (x *UpdateNode) GetRunList() []string {
	if x != nil {
		return x.RunList
	}
	return nil
}

func (x *UpdateNode) GetAutomaticAttributes() *_struct.Struct {
	if x != nil {
		return x.AutomaticAttributes
	}
	return nil
}

func (x *UpdateNode) GetDefaultAttributes() *_struct.Struct {
	if x != nil {
		return x.DefaultAttributes
	}
	return nil
}

func (x *UpdateNode) GetNormalAttributes() *_struct.Struct {
	if x != nil {
		return x.NormalAttributes
	}
	return nil
}

func (x *UpdateNode) GetOverrideAttributes() *_struct.Struct {
	if x != nil {
		return x.OverrideAttributes
	}
	return nil
}

type UpdateNodeTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Node tags action (e.g. 'add', 'delete', 'set').
	Action string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	// Node tags.
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateNodeTags) Reset() {
	*x = UpdateNodeTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeTags) ProtoMessage() {}

func (x *UpdateNodeTags) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_nodes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeTags.ProtoReflect.Descriptor instead.
func (*UpdateNodeTags) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNodeTags) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *UpdateNodeTags) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *UpdateNodeTags) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNodeTags) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *UpdateNodeTags) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_external_infra_proxy_request_nodes_proto protoreflect.FileDescriptor

var file_external_infra_proxy_request_nodes_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3b, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8e, 0x01, 0x0a,
	0x0d, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f,
	0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xf9, 0x03, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x14, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x11, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x44, 0x0a, 0x11, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x10, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x13, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x12, 0x6f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22,
	0x84, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x61,
	0x67, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_infra_proxy_request_nodes_proto_rawDescOnce sync.Once
	file_external_infra_proxy_request_nodes_proto_rawDescData = file_external_infra_proxy_request_nodes_proto_rawDesc
)

func file_external_infra_proxy_request_nodes_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_request_nodes_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_request_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_request_nodes_proto_rawDescData)
	})
	return file_external_infra_proxy_request_nodes_proto_rawDescData
}

var file_external_infra_proxy_request_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_external_infra_proxy_request_nodes_proto_goTypes = []interface{}{
	(*Nodes)(nil),          // 0: chef.automate.api.infra_proxy.request.Nodes
	(*Node)(nil),           // 1: chef.automate.api.infra_proxy.request.Node
	(*AffectedNodes)(nil),  // 2: chef.automate.api.infra_proxy.request.AffectedNodes
	(*DeleteNode)(nil),     // 3: chef.automate.api.infra_proxy.request.DeleteNode
	(*UpdateNode)(nil),     // 4: chef.automate.api.infra_proxy.request.UpdateNode
	(*UpdateNodeTags)(nil), // 5: chef.automate.api.infra_proxy.request.UpdateNodeTags
	(*_struct.Struct)(nil), // 6: google.protobuf.Struct
}
var file_external_infra_proxy_request_nodes_proto_depIdxs = []int32{
	6, // 0: chef.automate.api.infra_proxy.request.UpdateNode.automatic_attributes:type_name -> google.protobuf.Struct
	6, // 1: chef.automate.api.infra_proxy.request.UpdateNode.default_attributes:type_name -> google.protobuf.Struct
	6, // 2: chef.automate.api.infra_proxy.request.UpdateNode.normal_attributes:type_name -> google.protobuf.Struct
	6, // 3: chef.automate.api.infra_proxy.request.UpdateNode.override_attributes:type_name -> google.protobuf.Struct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_request_nodes_proto_init() }
func file_external_infra_proxy_request_nodes_proto_init() {
	if File_external_infra_proxy_request_nodes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_request_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
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
		file_external_infra_proxy_request_nodes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_external_infra_proxy_request_nodes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AffectedNodes); i {
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
		file_external_infra_proxy_request_nodes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNode); i {
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
		file_external_infra_proxy_request_nodes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNode); i {
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
		file_external_infra_proxy_request_nodes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeTags); i {
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
			RawDescriptor: file_external_infra_proxy_request_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_request_nodes_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_request_nodes_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_request_nodes_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_request_nodes_proto = out.File
	file_external_infra_proxy_request_nodes_proto_rawDesc = nil
	file_external_infra_proxy_request_nodes_proto_goTypes = nil
	file_external_infra_proxy_request_nodes_proto_depIdxs = nil
}
