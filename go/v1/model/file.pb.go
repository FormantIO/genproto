// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.2
// source: protos/model/v1/file.proto

package model

import (
	proto "github.com/golang/protobuf/proto"
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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*File_Url
	//	*File_Raw
	Data     isFile_Data `protobuf_oneof:"data"`
	Filename string      `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_file_proto_rawDescGZIP(), []int{0}
}

func (m *File) GetData() isFile_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *File) GetUrl() string {
	if x, ok := x.GetData().(*File_Url); ok {
		return x.Url
	}
	return ""
}

func (x *File) GetRaw() []byte {
	if x, ok := x.GetData().(*File_Raw); ok {
		return x.Raw
	}
	return nil
}

func (x *File) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type isFile_Data interface {
	isFile_Data()
}

type File_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}

type File_Raw struct {
	Raw []byte `protobuf:"bytes,2,opt,name=raw,proto3,oneof"`
}

func (*File_Url) isFile_Data() {}

func (*File_Raw) isFile_Data() {}

var File_protos_model_v1_file_proto protoreflect.FileDescriptor

var file_protos_model_v1_file_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x58, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x12, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x49, 0x4f, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_file_proto_rawDescOnce sync.Once
	file_protos_model_v1_file_proto_rawDescData = file_protos_model_v1_file_proto_rawDesc
)

func file_protos_model_v1_file_proto_rawDescGZIP() []byte {
	file_protos_model_v1_file_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_file_proto_rawDescData)
	})
	return file_protos_model_v1_file_proto_rawDescData
}

var file_protos_model_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_model_v1_file_proto_goTypes = []interface{}{
	(*File)(nil), // 0: v1.model.File
}
var file_protos_model_v1_file_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_model_v1_file_proto_init() }
func file_protos_model_v1_file_proto_init() {
	if File_protos_model_v1_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
	file_protos_model_v1_file_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*File_Url)(nil),
		(*File_Raw)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_file_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_file_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_file_proto_msgTypes,
	}.Build()
	File_protos_model_v1_file_proto = out.File
	file_protos_model_v1_file_proto_rawDesc = nil
	file_protos_model_v1_file_proto_goTypes = nil
	file_protos_model_v1_file_proto_depIdxs = nil
}
