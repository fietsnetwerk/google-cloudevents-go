// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: cloud/beyondcorp/appconnectors/v1/data.proto

package appconnectorsdata

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

// HealthStatus represents the health status.
type HealthStatus int32

const (
	// Health status is unknown: not initialized or failed to retrieve.
	HealthStatus_HEALTH_STATUS_UNSPECIFIED HealthStatus = 0
	// The resource is healthy.
	HealthStatus_HEALTHY HealthStatus = 1
	// The resource is unhealthy.
	HealthStatus_UNHEALTHY HealthStatus = 2
	// The resource is unresponsive.
	HealthStatus_UNRESPONSIVE HealthStatus = 3
	// Some sub-resources are UNHEALTHY.
	HealthStatus_DEGRADED HealthStatus = 4
)

// Enum value maps for HealthStatus.
var (
	HealthStatus_name = map[int32]string{
		0: "HEALTH_STATUS_UNSPECIFIED",
		1: "HEALTHY",
		2: "UNHEALTHY",
		3: "UNRESPONSIVE",
		4: "DEGRADED",
	}
	HealthStatus_value = map[string]int32{
		"HEALTH_STATUS_UNSPECIFIED": 0,
		"HEALTHY":                   1,
		"UNHEALTHY":                 2,
		"UNRESPONSIVE":              3,
		"DEGRADED":                  4,
	}
)

func (x HealthStatus) Enum() *HealthStatus {
	p := new(HealthStatus)
	*p = x
	return p
}

func (x HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_enumTypes[0].Descriptor()
}

func (HealthStatus) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_appconnectors_v1_data_proto_enumTypes[0]
}

func (x HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthStatus.Descriptor instead.
func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescGZIP(), []int{0}
}

// Represents the different states of a AppConnector.
type AppConnector_State int32

const (
	// Default value. This value is unused.
	AppConnector_STATE_UNSPECIFIED AppConnector_State = 0
	// AppConnector is being created.
	AppConnector_CREATING AppConnector_State = 1
	// AppConnector has been created.
	AppConnector_CREATED AppConnector_State = 2
	// AppConnector's configuration is being updated.
	AppConnector_UPDATING AppConnector_State = 3
	// AppConnector is being deleted.
	AppConnector_DELETING AppConnector_State = 4
	// AppConnector is down and may be restored in the future.
	// This happens when CCFE sends ProjectState = OFF.
	AppConnector_DOWN AppConnector_State = 5
)

// Enum value maps for AppConnector_State.
var (
	AppConnector_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "CREATED",
		3: "UPDATING",
		4: "DELETING",
		5: "DOWN",
	}
	AppConnector_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"CREATED":           2,
		"UPDATING":          3,
		"DELETING":          4,
		"DOWN":              5,
	}
)

func (x AppConnector_State) Enum() *AppConnector_State {
	p := new(AppConnector_State)
	*p = x
	return p
}

func (x AppConnector_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppConnector_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_enumTypes[1].Descriptor()
}

func (AppConnector_State) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_appconnectors_v1_data_proto_enumTypes[1]
}

func (x AppConnector_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppConnector_State.Descriptor instead.
func (AppConnector_State) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescGZIP(), []int{1, 0}
}

// ResourceInfo represents the information/status of an app connector resource.
// Such as:
// - remote_agent
//   - container
//   - runtime
//   - appgateway
//   - appconnector
//   - appconnection
//   - tunnel
//   - logagent
type ResourceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Unique Id for the resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Overall health status. Overall status is derived based on the status of
	// each sub level resources.
	Status HealthStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.events.cloud.beyondcorp.appconnectors.v1.HealthStatus" json:"status,omitempty"`
	// The timestamp to collect the info. It is suggested to be set by
	// the topmost level resource only.
	Time *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	// List of Info for the sub level resources.
	Sub []*ResourceInfo `protobuf:"bytes,5,rep,name=sub,proto3" json:"sub,omitempty"`
}

func (x *ResourceInfo) Reset() {
	*x = ResourceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInfo) ProtoMessage() {}

func (x *ResourceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInfo.ProtoReflect.Descriptor instead.
func (*ResourceInfo) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceInfo) GetStatus() HealthStatus {
	if x != nil {
		return x.Status
	}
	return HealthStatus_HEALTH_STATUS_UNSPECIFIED
}

func (x *ResourceInfo) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ResourceInfo) GetSub() []*ResourceInfo {
	if x != nil {
		return x.Sub
	}
	return nil
}

// A BeyondCorp connector resource that represents an application facing
// component deployed proximal to and with direct access to the application
// instances. It is used to establish connectivity between the remote enterprise
// environment and GCP. It initiates connections to the applications and can
// proxy the data from users over the connection.
type AppConnector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Unique resource name of the AppConnector.
	// The name is ignored when creating a AppConnector.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp when the resource was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when the resource was last modified.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Resource labels to represent user provided metadata.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. An arbitrary user-provided name for the AppConnector. Cannot
	// exceed 64 characters.
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. A unique identifier for the instance generated by the
	// system.
	Uid string `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. The current state of the AppConnector.
	State AppConnector_State `protobuf:"varint,7,opt,name=state,proto3,enum=google.events.cloud.beyondcorp.appconnectors.v1.AppConnector_State" json:"state,omitempty"`
	// Required. Principal information about the Identity of the AppConnector.
	PrincipalInfo *AppConnector_PrincipalInfo `protobuf:"bytes,8,opt,name=principal_info,json=principalInfo,proto3" json:"principal_info,omitempty"`
	// Optional. Resource info of the connector.
	ResourceInfo *ResourceInfo `protobuf:"bytes,11,opt,name=resource_info,json=resourceInfo,proto3" json:"resource_info,omitempty"`
}

func (x *AppConnector) Reset() {
	*x = AppConnector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnector) ProtoMessage() {}

func (x *AppConnector) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnector.ProtoReflect.Descriptor instead.
func (*AppConnector) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *AppConnector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppConnector) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AppConnector) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *AppConnector) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AppConnector) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AppConnector) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AppConnector) GetState() AppConnector_State {
	if x != nil {
		return x.State
	}
	return AppConnector_STATE_UNSPECIFIED
}

func (x *AppConnector) GetPrincipalInfo() *AppConnector_PrincipalInfo {
	if x != nil {
		return x.PrincipalInfo
	}
	return nil
}

func (x *AppConnector) GetResourceInfo() *ResourceInfo {
	if x != nil {
		return x.ResourceInfo
	}
	return nil
}

// The data within all AppConnector events.
type AppConnectorEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The AppConnector event payload. Unset for deletion events.
	Payload *AppConnector `protobuf:"bytes,1,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
}

func (x *AppConnectorEventData) Reset() {
	*x = AppConnectorEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnectorEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnectorEventData) ProtoMessage() {}

func (x *AppConnectorEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnectorEventData.ProtoReflect.Descriptor instead.
func (*AppConnectorEventData) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescGZIP(), []int{2}
}

func (x *AppConnectorEventData) GetPayload() *AppConnector {
	if x != nil {
		return x.Payload
	}
	return nil
}

// PrincipalInfo represents an Identity oneof.
type AppConnector_PrincipalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*AppConnector_PrincipalInfo_ServiceAccount_
	Type isAppConnector_PrincipalInfo_Type `protobuf_oneof:"type"`
}

func (x *AppConnector_PrincipalInfo) Reset() {
	*x = AppConnector_PrincipalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnector_PrincipalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnector_PrincipalInfo) ProtoMessage() {}

func (x *AppConnector_PrincipalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnector_PrincipalInfo.ProtoReflect.Descriptor instead.
func (*AppConnector_PrincipalInfo) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescGZIP(), []int{1, 0}
}

func (m *AppConnector_PrincipalInfo) GetType() isAppConnector_PrincipalInfo_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *AppConnector_PrincipalInfo) GetServiceAccount() *AppConnector_PrincipalInfo_ServiceAccount {
	if x, ok := x.GetType().(*AppConnector_PrincipalInfo_ServiceAccount_); ok {
		return x.ServiceAccount
	}
	return nil
}

type isAppConnector_PrincipalInfo_Type interface {
	isAppConnector_PrincipalInfo_Type()
}

type AppConnector_PrincipalInfo_ServiceAccount_ struct {
	// A GCP service account.
	ServiceAccount *AppConnector_PrincipalInfo_ServiceAccount `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount,proto3,oneof"`
}

func (*AppConnector_PrincipalInfo_ServiceAccount_) isAppConnector_PrincipalInfo_Type() {}

// ServiceAccount represents a GCP service account.
type AppConnector_PrincipalInfo_ServiceAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Email address of the service account.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AppConnector_PrincipalInfo_ServiceAccount) Reset() {
	*x = AppConnector_PrincipalInfo_ServiceAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnector_PrincipalInfo_ServiceAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnector_PrincipalInfo_ServiceAccount) ProtoMessage() {}

func (x *AppConnector_PrincipalInfo_ServiceAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnector_PrincipalInfo_ServiceAccount.ProtoReflect.Descriptor instead.
func (*AppConnector_PrincipalInfo_ServiceAccount) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *AppConnector_PrincipalInfo_ServiceAccount) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_cloud_beyondcorp_appconnectors_v1_data_proto protoreflect.FileDescriptor

var file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61,
	0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf6, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x55, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f,
	0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x73, 0x75, 0x62, 0x22, 0xcd, 0x07, 0x0a, 0x0c, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x61, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x59, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x72, 0x0a, 0x0e, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x62, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79,
	0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0xc7, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x85, 0x01, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x5a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x26, 0x0a,
	0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e,
	0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x69, 0x0a,
	0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x19, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x48,
	0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x52, 0x45,
	0x53, 0x50, 0x4f, 0x4e, 0x53, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45,
	0x47, 0x52, 0x41, 0x44, 0x45, 0x44, 0x10, 0x04, 0x42, 0x3b, 0xaa, 0x02, 0x38, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x65, 0x79, 0x6f, 0x6e, 0x64,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescOnce sync.Once
	file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescData = file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDesc
)

func file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescData)
	})
	return file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDescData
}

var file_cloud_beyondcorp_appconnectors_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloud_beyondcorp_appconnectors_v1_data_proto_goTypes = []interface{}{
	(HealthStatus)(0),                  // 0: google.events.cloud.beyondcorp.appconnectors.v1.HealthStatus
	(AppConnector_State)(0),            // 1: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.State
	(*ResourceInfo)(nil),               // 2: google.events.cloud.beyondcorp.appconnectors.v1.ResourceInfo
	(*AppConnector)(nil),               // 3: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector
	(*AppConnectorEventData)(nil),      // 4: google.events.cloud.beyondcorp.appconnectors.v1.AppConnectorEventData
	(*AppConnector_PrincipalInfo)(nil), // 5: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.PrincipalInfo
	nil,                                // 6: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.LabelsEntry
	(*AppConnector_PrincipalInfo_ServiceAccount)(nil), // 7: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.PrincipalInfo.ServiceAccount
	(*timestamppb.Timestamp)(nil),                     // 8: google.protobuf.Timestamp
}
var file_cloud_beyondcorp_appconnectors_v1_data_proto_depIdxs = []int32{
	0,  // 0: google.events.cloud.beyondcorp.appconnectors.v1.ResourceInfo.status:type_name -> google.events.cloud.beyondcorp.appconnectors.v1.HealthStatus
	8,  // 1: google.events.cloud.beyondcorp.appconnectors.v1.ResourceInfo.time:type_name -> google.protobuf.Timestamp
	2,  // 2: google.events.cloud.beyondcorp.appconnectors.v1.ResourceInfo.sub:type_name -> google.events.cloud.beyondcorp.appconnectors.v1.ResourceInfo
	8,  // 3: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.create_time:type_name -> google.protobuf.Timestamp
	8,  // 4: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.update_time:type_name -> google.protobuf.Timestamp
	6,  // 5: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.labels:type_name -> google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.LabelsEntry
	1,  // 6: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.state:type_name -> google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.State
	5,  // 7: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.principal_info:type_name -> google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.PrincipalInfo
	2,  // 8: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.resource_info:type_name -> google.events.cloud.beyondcorp.appconnectors.v1.ResourceInfo
	3,  // 9: google.events.cloud.beyondcorp.appconnectors.v1.AppConnectorEventData.payload:type_name -> google.events.cloud.beyondcorp.appconnectors.v1.AppConnector
	7,  // 10: google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.PrincipalInfo.service_account:type_name -> google.events.cloud.beyondcorp.appconnectors.v1.AppConnector.PrincipalInfo.ServiceAccount
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_cloud_beyondcorp_appconnectors_v1_data_proto_init() }
func file_cloud_beyondcorp_appconnectors_v1_data_proto_init() {
	if File_cloud_beyondcorp_appconnectors_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInfo); i {
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
		file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConnector); i {
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
		file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConnectorEventData); i {
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
		file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConnector_PrincipalInfo); i {
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
		file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConnector_PrincipalInfo_ServiceAccount); i {
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
	file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*AppConnector_PrincipalInfo_ServiceAccount_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_beyondcorp_appconnectors_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_beyondcorp_appconnectors_v1_data_proto_depIdxs,
		EnumInfos:         file_cloud_beyondcorp_appconnectors_v1_data_proto_enumTypes,
		MessageInfos:      file_cloud_beyondcorp_appconnectors_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_beyondcorp_appconnectors_v1_data_proto = out.File
	file_cloud_beyondcorp_appconnectors_v1_data_proto_rawDesc = nil
	file_cloud_beyondcorp_appconnectors_v1_data_proto_goTypes = nil
	file_cloud_beyondcorp_appconnectors_v1_data_proto_depIdxs = nil
}