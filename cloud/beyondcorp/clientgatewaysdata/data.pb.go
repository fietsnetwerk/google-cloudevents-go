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
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: cloud/beyondcorp/clientgateways/v1/data.proto

package clientgatewaysdata

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

// Represents the different states of a gateway.
type ClientGateway_State int32

const (
	// Default value. This value is unused.
	ClientGateway_STATE_UNSPECIFIED ClientGateway_State = 0
	// Gateway is being created.
	ClientGateway_CREATING ClientGateway_State = 1
	// Gateway is being updated.
	ClientGateway_UPDATING ClientGateway_State = 2
	// Gateway is being deleted.
	ClientGateway_DELETING ClientGateway_State = 3
	// Gateway is running.
	ClientGateway_RUNNING ClientGateway_State = 4
	// Gateway is down and may be restored in the future.
	// This happens when CCFE sends ProjectState = OFF.
	ClientGateway_DOWN ClientGateway_State = 5
	// ClientGateway encountered an error and is in indeterministic state.
	ClientGateway_ERROR ClientGateway_State = 6
)

// Enum value maps for ClientGateway_State.
var (
	ClientGateway_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "UPDATING",
		3: "DELETING",
		4: "RUNNING",
		5: "DOWN",
		6: "ERROR",
	}
	ClientGateway_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"UPDATING":          2,
		"DELETING":          3,
		"RUNNING":           4,
		"DOWN":              5,
		"ERROR":             6,
	}
)

func (x ClientGateway_State) Enum() *ClientGateway_State {
	p := new(ClientGateway_State)
	*p = x
	return p
}

func (x ClientGateway_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientGateway_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_clientgateways_v1_data_proto_enumTypes[0].Descriptor()
}

func (ClientGateway_State) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_clientgateways_v1_data_proto_enumTypes[0]
}

func (x ClientGateway_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientGateway_State.Descriptor instead.
func (ClientGateway_State) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

// Message describing ClientGateway object.
type ClientGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. name of resource. The name is ignored during creation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. [Output only] Create time stamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. [Output only] Update time stamp.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. The operational state of the gateway.
	State ClientGateway_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.events.cloud.beyondcorp.clientgateways.v1.ClientGateway_State" json:"state,omitempty"`
	// Output only. A unique identifier for the instance generated by the system.
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The client connector service name that the client gateway is
	// associated to. Client Connector Services, named as follows:
	//
	//	`projects/{project_id}/locations/{location_id}/client_connector_services/{client_connector_service_id}`.
	ClientConnectorService string `protobuf:"bytes,6,opt,name=client_connector_service,json=clientConnectorService,proto3" json:"client_connector_service,omitempty"`
}

func (x *ClientGateway) Reset() {
	*x = ClientGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientGateway) ProtoMessage() {}

func (x *ClientGateway) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientGateway.ProtoReflect.Descriptor instead.
func (*ClientGateway) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *ClientGateway) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientGateway) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ClientGateway) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ClientGateway) GetState() ClientGateway_State {
	if x != nil {
		return x.State
	}
	return ClientGateway_STATE_UNSPECIFIED
}

func (x *ClientGateway) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientGateway) GetClientConnectorService() string {
	if x != nil {
		return x.ClientConnectorService
	}
	return ""
}

// The data within all ClientGateway events.
type ClientGatewayEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The ClientGateway event payload. Unset for deletion events.
	Payload *ClientGateway `protobuf:"bytes,1,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
}

func (x *ClientGatewayEventData) Reset() {
	*x = ClientGatewayEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientGatewayEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientGatewayEventData) ProtoMessage() {}

func (x *ClientGatewayEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientGatewayEventData.ProtoReflect.Descriptor instead.
func (*ClientGatewayEventData) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *ClientGatewayEventData) GetPayload() *ClientGateway {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_cloud_beyondcorp_clientgateways_v1_data_proto protoreflect.FileDescriptor

var file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x03, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x38, 0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x6a, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x06, 0x22, 0x84, 0x01, 0x0a, 0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x5e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x3c, 0xaa, 0x02,
	0x39, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x65,
	0x79, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescOnce sync.Once
	file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescData = file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDesc
)

func file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescData)
	})
	return file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDescData
}

var file_cloud_beyondcorp_clientgateways_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_beyondcorp_clientgateways_v1_data_proto_goTypes = []interface{}{
	(ClientGateway_State)(0),       // 0: google.events.cloud.beyondcorp.clientgateways.v1.ClientGateway.State
	(*ClientGateway)(nil),          // 1: google.events.cloud.beyondcorp.clientgateways.v1.ClientGateway
	(*ClientGatewayEventData)(nil), // 2: google.events.cloud.beyondcorp.clientgateways.v1.ClientGatewayEventData
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_cloud_beyondcorp_clientgateways_v1_data_proto_depIdxs = []int32{
	3, // 0: google.events.cloud.beyondcorp.clientgateways.v1.ClientGateway.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.events.cloud.beyondcorp.clientgateways.v1.ClientGateway.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.events.cloud.beyondcorp.clientgateways.v1.ClientGateway.state:type_name -> google.events.cloud.beyondcorp.clientgateways.v1.ClientGateway.State
	1, // 3: google.events.cloud.beyondcorp.clientgateways.v1.ClientGatewayEventData.payload:type_name -> google.events.cloud.beyondcorp.clientgateways.v1.ClientGateway
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cloud_beyondcorp_clientgateways_v1_data_proto_init() }
func file_cloud_beyondcorp_clientgateways_v1_data_proto_init() {
	if File_cloud_beyondcorp_clientgateways_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientGateway); i {
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
		file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientGatewayEventData); i {
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
	file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_beyondcorp_clientgateways_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_beyondcorp_clientgateways_v1_data_proto_depIdxs,
		EnumInfos:         file_cloud_beyondcorp_clientgateways_v1_data_proto_enumTypes,
		MessageInfos:      file_cloud_beyondcorp_clientgateways_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_beyondcorp_clientgateways_v1_data_proto = out.File
	file_cloud_beyondcorp_clientgateways_v1_data_proto_rawDesc = nil
	file_cloud_beyondcorp_clientgateways_v1_data_proto_goTypes = nil
	file_cloud_beyondcorp_clientgateways_v1_data_proto_depIdxs = nil
}
