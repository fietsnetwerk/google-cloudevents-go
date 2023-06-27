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
// source: firebase/auth/v1/data.proto

package authdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// The data within all Firebase Auth events.
type AuthEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user identifier in the Firebase app.
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// The user's primary email, if set.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Whether or not the user's primary email is verified.
	EmailVerified bool `protobuf:"varint,3,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	// The user's display name.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The user's photo URL.
	Photo_URL string `protobuf:"bytes,5,opt,name=photo_URL,json=photoURL,proto3" json:"photo_URL,omitempty"`
	// Whether the user is disabled.
	Disabled bool `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Additional metadata about the user.
	Metadata *UserMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// User's info at the providers
	ProviderData []*UserInfo `protobuf:"bytes,8,rep,name=provider_data,json=providerData,proto3" json:"provider_data,omitempty"`
	// The user's phone number.
	PhoneNumber string `protobuf:"bytes,9,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// User's custom claims, typically used to define user roles and propagated
	// to an authenticated user's ID token.
	CustomClaims *structpb.Struct `protobuf:"bytes,10,opt,name=custom_claims,json=customClaims,proto3" json:"custom_claims,omitempty"`
}

func (x *AuthEventData) Reset() {
	*x = AuthEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_auth_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthEventData) ProtoMessage() {}

func (x *AuthEventData) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_auth_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthEventData.ProtoReflect.Descriptor instead.
func (*AuthEventData) Descriptor() ([]byte, []int) {
	return file_firebase_auth_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *AuthEventData) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AuthEventData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthEventData) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *AuthEventData) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AuthEventData) GetPhoto_URL() string {
	if x != nil {
		return x.Photo_URL
	}
	return ""
}

func (x *AuthEventData) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *AuthEventData) GetMetadata() *UserMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AuthEventData) GetProviderData() []*UserInfo {
	if x != nil {
		return x.ProviderData
	}
	return nil
}

func (x *AuthEventData) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *AuthEventData) GetCustomClaims() *structpb.Struct {
	if x != nil {
		return x.CustomClaims
	}
	return nil
}

// Additional metadata about the user.
type UserMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The date the user was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The date the user last signed in.
	LastSignInTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_sign_in_time,json=lastSignInTime,proto3" json:"last_sign_in_time,omitempty"`
}

func (x *UserMetadata) Reset() {
	*x = UserMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_auth_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMetadata) ProtoMessage() {}

func (x *UserMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_auth_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMetadata.ProtoReflect.Descriptor instead.
func (*UserMetadata) Descriptor() ([]byte, []int) {
	return file_firebase_auth_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *UserMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *UserMetadata) GetLastSignInTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSignInTime
	}
	return nil
}

// User's info at the identity provider
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user identifier for the linked provider.
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// The email for the linked provider.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// The display name for the linked provider.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The photo URL for the linked provider.
	Photo_URL string `protobuf:"bytes,4,opt,name=photo_URL,json=photoURL,proto3" json:"photo_URL,omitempty"`
	// The linked provider ID (e.g. "google.com" for the Google provider).
	ProviderId string `protobuf:"bytes,5,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_auth_v1_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_auth_v1_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_firebase_auth_v1_data_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UserInfo) GetPhoto_URL() string {
	if x != nil {
		return x.Photo_URL
	}
	return ""
}

func (x *UserInfo) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

var File_firebase_auth_v1_data_proto protoreflect.FileDescriptor

var file_firebase_auth_v1_data_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x03, 0x0a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x55, 0x52, 0x4c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x45, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x55, 0x52, 0x4c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x52, 0x4c, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x70,
	0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x46, 0x69,
	0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_firebase_auth_v1_data_proto_rawDescOnce sync.Once
	file_firebase_auth_v1_data_proto_rawDescData = file_firebase_auth_v1_data_proto_rawDesc
)

func file_firebase_auth_v1_data_proto_rawDescGZIP() []byte {
	file_firebase_auth_v1_data_proto_rawDescOnce.Do(func() {
		file_firebase_auth_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_firebase_auth_v1_data_proto_rawDescData)
	})
	return file_firebase_auth_v1_data_proto_rawDescData
}

var file_firebase_auth_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_firebase_auth_v1_data_proto_goTypes = []interface{}{
	(*AuthEventData)(nil),         // 0: google.events.firebase.auth.v1.AuthEventData
	(*UserMetadata)(nil),          // 1: google.events.firebase.auth.v1.UserMetadata
	(*UserInfo)(nil),              // 2: google.events.firebase.auth.v1.UserInfo
	(*structpb.Struct)(nil),       // 3: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_firebase_auth_v1_data_proto_depIdxs = []int32{
	1, // 0: google.events.firebase.auth.v1.AuthEventData.metadata:type_name -> google.events.firebase.auth.v1.UserMetadata
	2, // 1: google.events.firebase.auth.v1.AuthEventData.provider_data:type_name -> google.events.firebase.auth.v1.UserInfo
	3, // 2: google.events.firebase.auth.v1.AuthEventData.custom_claims:type_name -> google.protobuf.Struct
	4, // 3: google.events.firebase.auth.v1.UserMetadata.create_time:type_name -> google.protobuf.Timestamp
	4, // 4: google.events.firebase.auth.v1.UserMetadata.last_sign_in_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_firebase_auth_v1_data_proto_init() }
func file_firebase_auth_v1_data_proto_init() {
	if File_firebase_auth_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_firebase_auth_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthEventData); i {
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
		file_firebase_auth_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMetadata); i {
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
		file_firebase_auth_v1_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_firebase_auth_v1_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_firebase_auth_v1_data_proto_goTypes,
		DependencyIndexes: file_firebase_auth_v1_data_proto_depIdxs,
		MessageInfos:      file_firebase_auth_v1_data_proto_msgTypes,
	}.Build()
	File_firebase_auth_v1_data_proto = out.File
	file_firebase_auth_v1_data_proto_rawDesc = nil
	file_firebase_auth_v1_data_proto_goTypes = nil
	file_firebase_auth_v1_data_proto_depIdxs = nil
}
