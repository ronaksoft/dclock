// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: model.proto

package model

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

type Hook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID    []byte `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	ID          []byte `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Timestamp   int64  `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"` // UTC unix epoch time
	CallbackUrl []byte `protobuf:"bytes,4,opt,name=CallbackUrl,proto3" json:"CallbackUrl,omitempty"`
	JsonData    []byte `protobuf:"bytes,5,opt,name=JsonData,proto3" json:"JsonData,omitempty"`
	Fired       bool   `protobuf:"varint,6,opt,name=Fired,proto3" json:"Fired,omitempty"`
	Success     bool   `protobuf:"varint,7,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Hook) Reset() {
	*x = Hook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hook) ProtoMessage() {}

func (x *Hook) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hook.ProtoReflect.Descriptor instead.
func (*Hook) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *Hook) GetClientID() []byte {
	if x != nil {
		return x.ClientID
	}
	return nil
}

func (x *Hook) GetID() []byte {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Hook) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Hook) GetCallbackUrl() []byte {
	if x != nil {
		return x.CallbackUrl
	}
	return nil
}

func (x *Hook) GetJsonData() []byte {
	if x != nil {
		return x.JsonData
	}
	return nil
}

func (x *Hook) GetFired() bool {
	if x != nil {
		return x.Fired
	}
	return false
}

func (x *Hook) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type HookHolder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID []byte `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	ID       []byte `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Hook     *Hook  `protobuf:"bytes,3,opt,name=Hook,proto3" json:"Hook,omitempty"`
}

func (x *HookHolder) Reset() {
	*x = HookHolder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookHolder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookHolder) ProtoMessage() {}

func (x *HookHolder) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookHolder.ProtoReflect.Descriptor instead.
func (*HookHolder) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *HookHolder) GetClientID() []byte {
	if x != nil {
		return x.ClientID
	}
	return nil
}

func (x *HookHolder) GetID() []byte {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *HookHolder) GetHook() *Hook {
	if x != nil {
		return x.Hook
	}
	return nil
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x6f, 0x6e, 0x61, 0x6b, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x72, 0x6f, 0x6e, 0x79, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01,
	0x0a, 0x04, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x46, 0x69, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x46,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a, 0x2b,
	0x88, 0xb5, 0x18, 0x01, 0x92, 0xb5, 0x18, 0x0e, 0x28, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x2c, 0x20, 0x49, 0x44, 0x29, 0x9a, 0xb5, 0x18, 0x11, 0x28, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x2c, 0x20, 0x49, 0x44, 0x29, 0x22, 0x71, 0x0a, 0x0a, 0x48,
	0x6f, 0x6f, 0x6b, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x04, 0x48, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48, 0x6f, 0x6f, 0x6b,
	0x52, 0x04, 0x48, 0x6f, 0x6f, 0x6b, 0x3a, 0x16, 0x88, 0xb5, 0x18, 0x01, 0x92, 0xb5, 0x18, 0x0e,
	0x28, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x2c, 0x20, 0x49, 0x44, 0x29, 0x42, 0x2c,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6e,
	0x61, 0x6b, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x64, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x8a, 0xb5, 0x18, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_proto_goTypes = []interface{}{
	(*Hook)(nil),       // 0: model.Hook
	(*HookHolder)(nil), // 1: model.HookHolder
}
var file_model_proto_depIdxs = []int32{
	0, // 0: model.HookHolder.Hook:type_name -> model.Hook
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hook); i {
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
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookHolder); i {
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
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
