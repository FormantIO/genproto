// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/model/v1/media.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Image struct {
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Image_Url
	//	*Image_Raw
	Data                 isImage_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_a874c151d5cd4bde, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type isImage_Data interface {
	isImage_Data()
}

type Image_Url struct {
	Url string `protobuf:"bytes,2,opt,name=url,proto3,oneof"`
}

type Image_Raw struct {
	Raw []byte `protobuf:"bytes,3,opt,name=raw,proto3,oneof"`
}

func (*Image_Url) isImage_Data() {}

func (*Image_Raw) isImage_Data() {}

func (m *Image) GetData() isImage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Image) GetUrl() string {
	if x, ok := m.GetData().(*Image_Url); ok {
		return x.Url
	}
	return ""
}

func (m *Image) GetRaw() []byte {
	if x, ok := m.GetData().(*Image_Raw); ok {
		return x.Raw
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Image) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Image_Url)(nil),
		(*Image_Raw)(nil),
	}
}

type PointCloud struct {
	// Types that are valid to be assigned to Data:
	//	*PointCloud_Url
	//	*PointCloud_Raw
	Data                 isPointCloud_Data `protobuf_oneof:"data"`
	WorldToLocal         *Transform        `protobuf:"bytes,3,opt,name=world_to_local,json=worldToLocal,proto3" json:"world_to_local,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PointCloud) Reset()         { *m = PointCloud{} }
func (m *PointCloud) String() string { return proto.CompactTextString(m) }
func (*PointCloud) ProtoMessage()    {}
func (*PointCloud) Descriptor() ([]byte, []int) {
	return fileDescriptor_a874c151d5cd4bde, []int{1}
}

func (m *PointCloud) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointCloud.Unmarshal(m, b)
}
func (m *PointCloud) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointCloud.Marshal(b, m, deterministic)
}
func (m *PointCloud) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointCloud.Merge(m, src)
}
func (m *PointCloud) XXX_Size() int {
	return xxx_messageInfo_PointCloud.Size(m)
}
func (m *PointCloud) XXX_DiscardUnknown() {
	xxx_messageInfo_PointCloud.DiscardUnknown(m)
}

var xxx_messageInfo_PointCloud proto.InternalMessageInfo

type isPointCloud_Data interface {
	isPointCloud_Data()
}

type PointCloud_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}

type PointCloud_Raw struct {
	Raw []byte `protobuf:"bytes,2,opt,name=raw,proto3,oneof"`
}

func (*PointCloud_Url) isPointCloud_Data() {}

func (*PointCloud_Raw) isPointCloud_Data() {}

func (m *PointCloud) GetData() isPointCloud_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PointCloud) GetUrl() string {
	if x, ok := m.GetData().(*PointCloud_Url); ok {
		return x.Url
	}
	return ""
}

func (m *PointCloud) GetRaw() []byte {
	if x, ok := m.GetData().(*PointCloud_Raw); ok {
		return x.Raw
	}
	return nil
}

func (m *PointCloud) GetWorldToLocal() *Transform {
	if m != nil {
		return m.WorldToLocal
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PointCloud) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PointCloud_Url)(nil),
		(*PointCloud_Raw)(nil),
	}
}

type H264VideoFrame struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Flags                int32    `protobuf:"varint,2,opt,name=flags,proto3" json:"flags,omitempty"`
	FrameData            []byte   `protobuf:"bytes,3,opt,name=frame_data,json=frameData,proto3" json:"frame_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *H264VideoFrame) Reset()         { *m = H264VideoFrame{} }
func (m *H264VideoFrame) String() string { return proto.CompactTextString(m) }
func (*H264VideoFrame) ProtoMessage()    {}
func (*H264VideoFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a874c151d5cd4bde, []int{2}
}

func (m *H264VideoFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_H264VideoFrame.Unmarshal(m, b)
}
func (m *H264VideoFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_H264VideoFrame.Marshal(b, m, deterministic)
}
func (m *H264VideoFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_H264VideoFrame.Merge(m, src)
}
func (m *H264VideoFrame) XXX_Size() int {
	return xxx_messageInfo_H264VideoFrame.Size(m)
}
func (m *H264VideoFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_H264VideoFrame.DiscardUnknown(m)
}

var xxx_messageInfo_H264VideoFrame proto.InternalMessageInfo

func (m *H264VideoFrame) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *H264VideoFrame) GetFlags() int32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *H264VideoFrame) GetFrameData() []byte {
	if m != nil {
		return m.FrameData
	}
	return nil
}

type AudioChunk struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Format               string   `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	ChunkData            []byte   `protobuf:"bytes,3,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AudioChunk) Reset()         { *m = AudioChunk{} }
func (m *AudioChunk) String() string { return proto.CompactTextString(m) }
func (*AudioChunk) ProtoMessage()    {}
func (*AudioChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_a874c151d5cd4bde, []int{3}
}

func (m *AudioChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioChunk.Unmarshal(m, b)
}
func (m *AudioChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioChunk.Marshal(b, m, deterministic)
}
func (m *AudioChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioChunk.Merge(m, src)
}
func (m *AudioChunk) XXX_Size() int {
	return xxx_messageInfo_AudioChunk.Size(m)
}
func (m *AudioChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioChunk.DiscardUnknown(m)
}

var xxx_messageInfo_AudioChunk proto.InternalMessageInfo

func (m *AudioChunk) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AudioChunk) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *AudioChunk) GetChunkData() []byte {
	if m != nil {
		return m.ChunkData
	}
	return nil
}

type Video struct {
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Duration int64  `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Video_Url
	//	*Video_Raw
	Data                 isVideo_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Video) Reset()         { *m = Video{} }
func (m *Video) String() string { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()    {}
func (*Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_a874c151d5cd4bde, []int{4}
}

func (m *Video) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Video.Unmarshal(m, b)
}
func (m *Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Video.Marshal(b, m, deterministic)
}
func (m *Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Video.Merge(m, src)
}
func (m *Video) XXX_Size() int {
	return xxx_messageInfo_Video.Size(m)
}
func (m *Video) XXX_DiscardUnknown() {
	xxx_messageInfo_Video.DiscardUnknown(m)
}

var xxx_messageInfo_Video proto.InternalMessageInfo

func (m *Video) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Video) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type isVideo_Data interface {
	isVideo_Data()
}

type Video_Url struct {
	Url string `protobuf:"bytes,3,opt,name=url,proto3,oneof"`
}

type Video_Raw struct {
	Raw []byte `protobuf:"bytes,4,opt,name=raw,proto3,oneof"`
}

func (*Video_Url) isVideo_Data() {}

func (*Video_Raw) isVideo_Data() {}

func (m *Video) GetData() isVideo_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Video) GetUrl() string {
	if x, ok := m.GetData().(*Video_Url); ok {
		return x.Url
	}
	return ""
}

func (m *Video) GetRaw() []byte {
	if x, ok := m.GetData().(*Video_Raw); ok {
		return x.Raw
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Video) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Video_Url)(nil),
		(*Video_Raw)(nil),
	}
}

func init() {
	proto.RegisterType((*Image)(nil), "v1.model.Image")
	proto.RegisterType((*PointCloud)(nil), "v1.model.PointCloud")
	proto.RegisterType((*H264VideoFrame)(nil), "v1.model.H264VideoFrame")
	proto.RegisterType((*AudioChunk)(nil), "v1.model.AudioChunk")
	proto.RegisterType((*Video)(nil), "v1.model.Video")
}

func init() { proto.RegisterFile("protos/model/v1/media.proto", fileDescriptor_a874c151d5cd4bde) }

var fileDescriptor_a874c151d5cd4bde = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0xab, 0xda, 0x30,
	0x14, 0xc6, 0xd7, 0xdb, 0x5b, 0xd1, 0x73, 0xe5, 0x3e, 0x64, 0x63, 0x88, 0x97, 0x81, 0xf3, 0xc9,
	0x31, 0x68, 0xd1, 0x8d, 0xc1, 0x1e, 0xa7, 0x43, 0x14, 0x06, 0x1b, 0x41, 0x06, 0x73, 0x0f, 0x25,
	0x36, 0xb1, 0x86, 0x35, 0x39, 0x25, 0xa6, 0x76, 0xfe, 0xf7, 0x23, 0xa9, 0x8a, 0x63, 0xee, 0xf1,
	0xfb, 0xe5, 0x7c, 0x5f, 0x4e, 0xce, 0x09, 0x3c, 0x95, 0x06, 0x2d, 0xee, 0x13, 0x85, 0x5c, 0x14,
	0xc9, 0x61, 0x9c, 0x28, 0xc1, 0x25, 0x8b, 0x3d, 0x25, 0xed, 0xc3, 0x38, 0xf6, 0x07, 0xfd, 0xfe,
	0x3f, 0x65, 0xcc, 0xee, 0x9a, 0xaa, 0xe1, 0x1a, 0xa2, 0xa5, 0x62, 0xb9, 0x20, 0xaf, 0xa1, 0x9b,
	0xa1, 0xb6, 0x42, 0xdb, 0xd4, 0x1e, 0x4b, 0xd1, 0x0b, 0x06, 0xc1, 0xa8, 0x43, 0x1f, 0x4e, 0x6c,
	0x75, 0x2c, 0x05, 0x21, 0x10, 0x56, 0xa6, 0xe8, 0xdd, 0xb9, 0x93, 0xc5, 0x33, 0xea, 0x84, 0x63,
	0x86, 0xd5, 0xbd, 0x70, 0x10, 0x8c, 0xba, 0x8e, 0x19, 0x56, 0x4f, 0x5b, 0x70, 0xcf, 0x99, 0x65,
	0xc3, 0x1a, 0xe0, 0x1b, 0x4a, 0x6d, 0x67, 0x05, 0x56, 0xfc, 0xec, 0x0e, 0x6e, 0xb8, 0xef, 0xae,
	0xdc, 0xe4, 0x23, 0x3c, 0xd6, 0x68, 0x0a, 0x9e, 0x5a, 0x4c, 0x0b, 0xcc, 0x58, 0xe1, 0xc3, 0x1f,
	0x26, 0xcf, 0xe3, 0xf3, 0x83, 0xe2, 0x95, 0x61, 0x7a, 0xbf, 0x45, 0xa3, 0x68, 0xd7, 0x97, 0xae,
	0xf0, 0x8b, 0x2b, 0xbc, 0x5c, 0xfc, 0x13, 0x1e, 0x17, 0x93, 0x0f, 0xef, 0xbf, 0x4b, 0x2e, 0x70,
	0x6e, 0x98, 0x12, 0xe4, 0x05, 0x44, 0x52, 0x73, 0xf1, 0xdb, 0x5f, 0x1f, 0xd1, 0x46, 0x38, 0xba,
	0x2d, 0x58, 0xbe, 0xf7, 0x0d, 0x44, 0xb4, 0x11, 0xe4, 0x15, 0xc0, 0xd6, 0x99, 0x52, 0x97, 0xd5,
	0xbc, 0x8c, 0x76, 0x3c, 0xf9, 0xec, 0xc2, 0x7f, 0x00, 0x7c, 0xaa, 0xb8, 0xc4, 0xd9, 0xae, 0xd2,
	0xbf, 0xfe, 0x13, 0xfc, 0x12, 0x5a, 0xae, 0x3d, 0x66, 0x9b, 0x61, 0xd1, 0x93, 0x72, 0xd1, 0x99,
	0xb3, 0xfd, 0x15, 0xed, 0x89, 0x8f, 0x2e, 0x21, 0xf2, 0x3d, 0x93, 0x27, 0xe8, 0x28, 0xa9, 0xc4,
	0xf5, 0x26, 0xda, 0x0e, 0xf8, 0x35, 0xf4, 0xa1, 0xcd, 0x2b, 0xc3, 0xac, 0x44, 0xed, 0xe3, 0x43,
	0x7a, 0xd1, 0xe7, 0x21, 0x87, 0x37, 0x86, 0x7c, 0x7f, 0x63, 0x45, 0xd3, 0xb7, 0xeb, 0x37, 0xb9,
	0xb4, 0xbb, 0x6a, 0x13, 0x67, 0xa8, 0x92, 0xb9, 0xeb, 0x52, 0xdb, 0xe5, 0xd7, 0x24, 0x17, 0xda,
	0xff, 0x8f, 0x24, 0x47, 0xff, 0x61, 0xdc, 0xd8, 0x37, 0x2d, 0x8f, 0xde, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x6c, 0x2e, 0x47, 0x77, 0x02, 0x00, 0x00,
}
