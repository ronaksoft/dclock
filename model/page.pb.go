// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: page.proto

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

// Page
type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ReplicaSet uint64 `protobuf:"varint,2,opt,name=ReplicaSet,proto3" json:"ReplicaSet,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Page) GetReplicaSet() uint64 {
	if x != nil {
		return x.ReplicaSet
	}
	return 0
}

var File_page_proto protoreflect.FileDescriptor

var file_page_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x6f, 0x6e, 0x61, 0x6b, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x72, 0x6f, 0x6e, 0x79, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x04,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x53, 0x65, 0x74, 0x3a, 0x28, 0x88, 0xb5, 0x18, 0x01, 0xd2, 0xb5, 0x18, 0x04, 0x63, 0x72,
	0x75, 0x64, 0xf2, 0xb5, 0x18, 0x04, 0x28, 0x49, 0x44, 0x29, 0xfa, 0xb5, 0x18, 0x10, 0x28, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x65, 0x74, 0x2c, 0x20, 0x49, 0x44, 0x29, 0x42, 0x2c,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6e,
	0x61, 0x6b, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x64, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x8a, 0xb5, 0x18, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_page_proto_rawDescOnce sync.Once
	file_page_proto_rawDescData = file_page_proto_rawDesc
)

func file_page_proto_rawDescGZIP() []byte {
	file_page_proto_rawDescOnce.Do(func() {
		file_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_page_proto_rawDescData)
	})
	return file_page_proto_rawDescData
}

var file_page_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_page_proto_goTypes = []interface{}{
	(*Page)(nil), // 0: model.Page
}
var file_page_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_page_proto_init() }
func file_page_proto_init() {
	if File_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_page_proto_goTypes,
		DependencyIndexes: file_page_proto_depIdxs,
		MessageInfos:      file_page_proto_msgTypes,
	}.Build()
	File_page_proto = out.File
	file_page_proto_rawDesc = nil
	file_page_proto_goTypes = nil
	file_page_proto_depIdxs = nil
}