// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: clock.proto

package service

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ronaksoft/rony"
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

type HookSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueID     []byte `protobuf:"bytes,1,opt,name=UniqueID,proto3" json:"UniqueID,omitempty"`
	Timestamp    int64  `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"` // UTC unix epoch time
	HookUrl      []byte `protobuf:"bytes,3,opt,name=HookUrl,proto3" json:"HookUrl,omitempty"`
	HookJsonData []byte `protobuf:"bytes,4,opt,name=HookJsonData,proto3" json:"HookJsonData,omitempty"`
}

func (x *HookSetRequest) Reset() {
	*x = HookSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookSetRequest) ProtoMessage() {}

func (x *HookSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookSetRequest.ProtoReflect.Descriptor instead.
func (*HookSetRequest) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{0}
}

func (x *HookSetRequest) GetUniqueID() []byte {
	if x != nil {
		return x.UniqueID
	}
	return nil
}

func (x *HookSetRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HookSetRequest) GetHookUrl() []byte {
	if x != nil {
		return x.HookUrl
	}
	return nil
}

func (x *HookSetRequest) GetHookJsonData() []byte {
	if x != nil {
		return x.HookJsonData
	}
	return nil
}

type HookSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
}

func (x *HookSetResponse) Reset() {
	*x = HookSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookSetResponse) ProtoMessage() {}

func (x *HookSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookSetResponse.ProtoReflect.Descriptor instead.
func (*HookSetResponse) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{1}
}

func (x *HookSetResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

type HookDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueID []byte `protobuf:"bytes,1,opt,name=UniqueID,proto3" json:"UniqueID,omitempty"`
}

func (x *HookDeleteRequest) Reset() {
	*x = HookDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookDeleteRequest) ProtoMessage() {}

func (x *HookDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookDeleteRequest.ProtoReflect.Descriptor instead.
func (*HookDeleteRequest) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{2}
}

func (x *HookDeleteRequest) GetUniqueID() []byte {
	if x != nil {
		return x.UniqueID
	}
	return nil
}

type HookDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
}

func (x *HookDeleteResponse) Reset() {
	*x = HookDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookDeleteResponse) ProtoMessage() {}

func (x *HookDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookDeleteResponse.ProtoReflect.Descriptor instead.
func (*HookDeleteResponse) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{3}
}

func (x *HookDeleteResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

var File_clock_proto protoreflect.FileDescriptor

var file_clock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6e, 0x61, 0x6b, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x72, 0x6f, 0x6e,
	0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x88, 0x01, 0x0a, 0x0e, 0x48, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x48, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x48,
	0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x48, 0x6f, 0x6f, 0x6b, 0x4a, 0x73,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x48, 0x6f,
	0x6f, 0x6b, 0x4a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x0f, 0x48, 0x6f,
	0x6f, 0x6b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x22, 0x2f, 0x0a,
	0x11, 0x48, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x22, 0x34,
	0x0a, 0x12, 0x48, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66,
	0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x66, 0x75, 0x6c, 0x32, 0x92, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x3c,
	0x0a, 0x07, 0x48, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x6f,
	0x6b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a,
	0x48, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6e, 0x61, 0x6b, 0x73, 0x6f, 0x66,
	0x74, 0x2f, 0x64, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clock_proto_rawDescOnce sync.Once
	file_clock_proto_rawDescData = file_clock_proto_rawDesc
)

func file_clock_proto_rawDescGZIP() []byte {
	file_clock_proto_rawDescOnce.Do(func() {
		file_clock_proto_rawDescData = protoimpl.X.CompressGZIP(file_clock_proto_rawDescData)
	})
	return file_clock_proto_rawDescData
}

var file_clock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_clock_proto_goTypes = []interface{}{
	(*HookSetRequest)(nil),     // 0: service.HookSetRequest
	(*HookSetResponse)(nil),    // 1: service.HookSetResponse
	(*HookDeleteRequest)(nil),  // 2: service.HookDeleteRequest
	(*HookDeleteResponse)(nil), // 3: service.HookDeleteResponse
}
var file_clock_proto_depIdxs = []int32{
	0, // 0: service.Clock.HookSet:input_type -> service.HookSetRequest
	2, // 1: service.Clock.HookDelete:input_type -> service.HookDeleteRequest
	1, // 2: service.Clock.HookSet:output_type -> service.HookSetResponse
	3, // 3: service.Clock.HookDelete:output_type -> service.HookDeleteResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clock_proto_init() }
func file_clock_proto_init() {
	if File_clock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookSetRequest); i {
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
		file_clock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookSetResponse); i {
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
		file_clock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookDeleteRequest); i {
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
		file_clock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookDeleteResponse); i {
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
			RawDescriptor: file_clock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clock_proto_goTypes,
		DependencyIndexes: file_clock_proto_depIdxs,
		MessageInfos:      file_clock_proto_msgTypes,
	}.Build()
	File_clock_proto = out.File
	file_clock_proto_rawDesc = nil
	file_clock_proto_goTypes = nil
	file_clock_proto_depIdxs = nil
}
