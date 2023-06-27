// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.6
// source: firebase/remoteconfig/v1/data.proto

package remoteconfigdata

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

// What type of update was associated with the Remote Config template version.
type RemoteConfigUpdateOrigin int32

const (
	// Catch-all for unrecognized values.
	RemoteConfigUpdateOrigin_REMOTE_CONFIG_UPDATE_ORIGIN_UNSPECIFIED RemoteConfigUpdateOrigin = 0
	// The update came from the Firebase UI.
	RemoteConfigUpdateOrigin_CONSOLE RemoteConfigUpdateOrigin = 1
	// The update came from the Remote Config REST API.
	RemoteConfigUpdateOrigin_REST_API RemoteConfigUpdateOrigin = 2
	// The update came from the Firebase Admin Node SDK.
	RemoteConfigUpdateOrigin_ADMIN_SDK_NODE RemoteConfigUpdateOrigin = 3
)

// Enum value maps for RemoteConfigUpdateOrigin.
var (
	RemoteConfigUpdateOrigin_name = map[int32]string{
		0: "REMOTE_CONFIG_UPDATE_ORIGIN_UNSPECIFIED",
		1: "CONSOLE",
		2: "REST_API",
		3: "ADMIN_SDK_NODE",
	}
	RemoteConfigUpdateOrigin_value = map[string]int32{
		"REMOTE_CONFIG_UPDATE_ORIGIN_UNSPECIFIED": 0,
		"CONSOLE":        1,
		"REST_API":       2,
		"ADMIN_SDK_NODE": 3,
	}
)

func (x RemoteConfigUpdateOrigin) Enum() *RemoteConfigUpdateOrigin {
	p := new(RemoteConfigUpdateOrigin)
	*p = x
	return p
}

func (x RemoteConfigUpdateOrigin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoteConfigUpdateOrigin) Descriptor() protoreflect.EnumDescriptor {
	return file_firebase_remoteconfig_v1_data_proto_enumTypes[0].Descriptor()
}

func (RemoteConfigUpdateOrigin) Type() protoreflect.EnumType {
	return &file_firebase_remoteconfig_v1_data_proto_enumTypes[0]
}

func (x RemoteConfigUpdateOrigin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RemoteConfigUpdateOrigin.Descriptor instead.
func (RemoteConfigUpdateOrigin) EnumDescriptor() ([]byte, []int) {
	return file_firebase_remoteconfig_v1_data_proto_rawDescGZIP(), []int{0}
}

// Where the Remote Config update action originated.
type RemoteConfigUpdateType int32

const (
	// Catch-all for unrecognized enum values.
	RemoteConfigUpdateType_REMOTE_CONFIG_UPDATE_TYPE_UNSPECIFIED RemoteConfigUpdateType = 0
	// A regular incremental update.
	RemoteConfigUpdateType_INCREMENTAL_UPDATE RemoteConfigUpdateType = 1
	// A forced update.
	// The ETag was specified as "*" in an UpdateRemoteConfigRequest
	// request or the "Force Update" button was pressed on the console.
	RemoteConfigUpdateType_FORCED_UPDATE RemoteConfigUpdateType = 2
	// A rollback to a previous Remote Config template.
	RemoteConfigUpdateType_ROLLBACK RemoteConfigUpdateType = 3
)

// Enum value maps for RemoteConfigUpdateType.
var (
	RemoteConfigUpdateType_name = map[int32]string{
		0: "REMOTE_CONFIG_UPDATE_TYPE_UNSPECIFIED",
		1: "INCREMENTAL_UPDATE",
		2: "FORCED_UPDATE",
		3: "ROLLBACK",
	}
	RemoteConfigUpdateType_value = map[string]int32{
		"REMOTE_CONFIG_UPDATE_TYPE_UNSPECIFIED": 0,
		"INCREMENTAL_UPDATE":                    1,
		"FORCED_UPDATE":                         2,
		"ROLLBACK":                              3,
	}
)

func (x RemoteConfigUpdateType) Enum() *RemoteConfigUpdateType {
	p := new(RemoteConfigUpdateType)
	*p = x
	return p
}

func (x RemoteConfigUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoteConfigUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_firebase_remoteconfig_v1_data_proto_enumTypes[1].Descriptor()
}

func (RemoteConfigUpdateType) Type() protoreflect.EnumType {
	return &file_firebase_remoteconfig_v1_data_proto_enumTypes[1]
}

func (x RemoteConfigUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RemoteConfigUpdateType.Descriptor instead.
func (RemoteConfigUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_firebase_remoteconfig_v1_data_proto_rawDescGZIP(), []int{1}
}

// The data within all Firebase Remote Config events.
type RemoteConfigEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version number of the version's corresponding Remote Config template.
	VersionNumber int64 `protobuf:"varint,1,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
	// When the Remote Config template was written to the Remote Config server.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Aggregation of all metadata fields about the account that performed the
	// update.
	UpdateUser *RemoteConfigUser `protobuf:"bytes,3,opt,name=update_user,json=updateUser,proto3" json:"update_user,omitempty"`
	// The user-provided description of the corresponding Remote Config template.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Where the update action originated.
	UpdateOrigin RemoteConfigUpdateOrigin `protobuf:"varint,5,opt,name=update_origin,json=updateOrigin,proto3,enum=google.events.firebase.remoteconfig.v1.RemoteConfigUpdateOrigin" json:"update_origin,omitempty"`
	// What type of update was made.
	UpdateType RemoteConfigUpdateType `protobuf:"varint,6,opt,name=update_type,json=updateType,proto3,enum=google.events.firebase.remoteconfig.v1.RemoteConfigUpdateType" json:"update_type,omitempty"`
	// Only present if this version is the result of a rollback, and will be the
	// version number of the Remote Config template that was rolled-back to.
	RollbackSource int64 `protobuf:"varint,7,opt,name=rollback_source,json=rollbackSource,proto3" json:"rollback_source,omitempty"`
}

func (x *RemoteConfigEventData) Reset() {
	*x = RemoteConfigEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_remoteconfig_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfigEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfigEventData) ProtoMessage() {}

func (x *RemoteConfigEventData) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_remoteconfig_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfigEventData.ProtoReflect.Descriptor instead.
func (*RemoteConfigEventData) Descriptor() ([]byte, []int) {
	return file_firebase_remoteconfig_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteConfigEventData) GetVersionNumber() int64 {
	if x != nil {
		return x.VersionNumber
	}
	return 0
}

func (x *RemoteConfigEventData) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *RemoteConfigEventData) GetUpdateUser() *RemoteConfigUser {
	if x != nil {
		return x.UpdateUser
	}
	return nil
}

func (x *RemoteConfigEventData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RemoteConfigEventData) GetUpdateOrigin() RemoteConfigUpdateOrigin {
	if x != nil {
		return x.UpdateOrigin
	}
	return RemoteConfigUpdateOrigin_REMOTE_CONFIG_UPDATE_ORIGIN_UNSPECIFIED
}

func (x *RemoteConfigEventData) GetUpdateType() RemoteConfigUpdateType {
	if x != nil {
		return x.UpdateType
	}
	return RemoteConfigUpdateType_REMOTE_CONFIG_UPDATE_TYPE_UNSPECIFIED
}

func (x *RemoteConfigEventData) GetRollbackSource() int64 {
	if x != nil {
		return x.RollbackSource
	}
	return 0
}

// All the fields associated with the person/service account
// that wrote a Remote Config template.
type RemoteConfigUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Display name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Email address.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Image URL.
	ImageUrl string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *RemoteConfigUser) Reset() {
	*x = RemoteConfigUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_remoteconfig_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfigUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfigUser) ProtoMessage() {}

func (x *RemoteConfigUser) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_remoteconfig_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfigUser.ProtoReflect.Descriptor instead.
func (*RemoteConfigUser) Descriptor() ([]byte, []int) {
	return file_firebase_remoteconfig_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteConfigUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoteConfigUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RemoteConfigUser) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_firebase_remoteconfig_v1_data_proto protoreflect.FileDescriptor

var file_firebase_remoteconfig_v1_data_proto_rawDesc = []byte{
	0x0a, 0x23, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9,
	0x03, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x0d, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x5f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x59, 0x0a, 0x10, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x2a, 0x76, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x12, 0x2b, 0x0a, 0x27, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x53, 0x4f, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x53, 0x54, 0x5f, 0x41, 0x50, 0x49, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x5f, 0x53, 0x44, 0x4b, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x03, 0x2a, 0x7c, 0x0a,
	0x16, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x25, 0x52, 0x45, 0x4d, 0x4f, 0x54,
	0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x43, 0x52, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x41,
	0x4c, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x4f,
	0x52, 0x43, 0x45, 0x44, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x03, 0x42, 0xbf, 0x01, 0x0a, 0x2a,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x5c, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x2a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_firebase_remoteconfig_v1_data_proto_rawDescOnce sync.Once
	file_firebase_remoteconfig_v1_data_proto_rawDescData = file_firebase_remoteconfig_v1_data_proto_rawDesc
)

func file_firebase_remoteconfig_v1_data_proto_rawDescGZIP() []byte {
	file_firebase_remoteconfig_v1_data_proto_rawDescOnce.Do(func() {
		file_firebase_remoteconfig_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_firebase_remoteconfig_v1_data_proto_rawDescData)
	})
	return file_firebase_remoteconfig_v1_data_proto_rawDescData
}

var file_firebase_remoteconfig_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_firebase_remoteconfig_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_firebase_remoteconfig_v1_data_proto_goTypes = []interface{}{
	(RemoteConfigUpdateOrigin)(0), // 0: google.events.firebase.remoteconfig.v1.RemoteConfigUpdateOrigin
	(RemoteConfigUpdateType)(0),   // 1: google.events.firebase.remoteconfig.v1.RemoteConfigUpdateType
	(*RemoteConfigEventData)(nil), // 2: google.events.firebase.remoteconfig.v1.RemoteConfigEventData
	(*RemoteConfigUser)(nil),      // 3: google.events.firebase.remoteconfig.v1.RemoteConfigUser
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_firebase_remoteconfig_v1_data_proto_depIdxs = []int32{
	4, // 0: google.events.firebase.remoteconfig.v1.RemoteConfigEventData.update_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.events.firebase.remoteconfig.v1.RemoteConfigEventData.update_user:type_name -> google.events.firebase.remoteconfig.v1.RemoteConfigUser
	0, // 2: google.events.firebase.remoteconfig.v1.RemoteConfigEventData.update_origin:type_name -> google.events.firebase.remoteconfig.v1.RemoteConfigUpdateOrigin
	1, // 3: google.events.firebase.remoteconfig.v1.RemoteConfigEventData.update_type:type_name -> google.events.firebase.remoteconfig.v1.RemoteConfigUpdateType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_firebase_remoteconfig_v1_data_proto_init() }
func file_firebase_remoteconfig_v1_data_proto_init() {
	if File_firebase_remoteconfig_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_firebase_remoteconfig_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfigEventData); i {
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
		file_firebase_remoteconfig_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfigUser); i {
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
			RawDescriptor: file_firebase_remoteconfig_v1_data_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_firebase_remoteconfig_v1_data_proto_goTypes,
		DependencyIndexes: file_firebase_remoteconfig_v1_data_proto_depIdxs,
		EnumInfos:         file_firebase_remoteconfig_v1_data_proto_enumTypes,
		MessageInfos:      file_firebase_remoteconfig_v1_data_proto_msgTypes,
	}.Build()
	File_firebase_remoteconfig_v1_data_proto = out.File
	file_firebase_remoteconfig_v1_data_proto_rawDesc = nil
	file_firebase_remoteconfig_v1_data_proto_goTypes = nil
	file_firebase_remoteconfig_v1_data_proto_depIdxs = nil
}
