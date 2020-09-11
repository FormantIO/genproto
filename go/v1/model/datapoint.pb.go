// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/model/v1/datapoint.proto

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

type Datapoint struct {
	Stream    string            `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	Timestamp int64             `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Tags      map[string]string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to Data:
	//	*Datapoint_Text
	//	*Datapoint_Numeric
	//	*Datapoint_NumericSet
	//	*Datapoint_Bitset
	//	*Datapoint_File
	//	*Datapoint_Image
	//	*Datapoint_PointCloud
	//	*Datapoint_Location
	//	*Datapoint_Localization
	//	*Datapoint_Health
	//	*Datapoint_Json
	//	*Datapoint_Battery
	//	*Datapoint_Video
	Data                 isDatapoint_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Datapoint) Reset()         { *m = Datapoint{} }
func (m *Datapoint) String() string { return proto.CompactTextString(m) }
func (*Datapoint) ProtoMessage()    {}
func (*Datapoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4daa958809a1527, []int{0}
}

func (m *Datapoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Datapoint.Unmarshal(m, b)
}
func (m *Datapoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Datapoint.Marshal(b, m, deterministic)
}
func (m *Datapoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Datapoint.Merge(m, src)
}
func (m *Datapoint) XXX_Size() int {
	return xxx_messageInfo_Datapoint.Size(m)
}
func (m *Datapoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Datapoint.DiscardUnknown(m)
}

var xxx_messageInfo_Datapoint proto.InternalMessageInfo

func (m *Datapoint) GetStream() string {
	if m != nil {
		return m.Stream
	}
	return ""
}

func (m *Datapoint) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Datapoint) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type isDatapoint_Data interface {
	isDatapoint_Data()
}

type Datapoint_Text struct {
	Text *Text `protobuf:"bytes,4,opt,name=text,proto3,oneof"`
}

type Datapoint_Numeric struct {
	Numeric *Numeric `protobuf:"bytes,5,opt,name=numeric,proto3,oneof"`
}

type Datapoint_NumericSet struct {
	NumericSet *NumericSet `protobuf:"bytes,17,opt,name=numeric_set,json=numericSet,proto3,oneof"`
}

type Datapoint_Bitset struct {
	Bitset *Bitset `protobuf:"bytes,7,opt,name=bitset,proto3,oneof"`
}

type Datapoint_File struct {
	File *File `protobuf:"bytes,8,opt,name=file,proto3,oneof"`
}

type Datapoint_Image struct {
	Image *Image `protobuf:"bytes,9,opt,name=image,proto3,oneof"`
}

type Datapoint_PointCloud struct {
	PointCloud *PointCloud `protobuf:"bytes,10,opt,name=point_cloud,json=pointCloud,proto3,oneof"`
}

type Datapoint_Location struct {
	Location *Location `protobuf:"bytes,11,opt,name=location,proto3,oneof"`
}

type Datapoint_Localization struct {
	Localization *Localization `protobuf:"bytes,12,opt,name=localization,proto3,oneof"`
}

type Datapoint_Health struct {
	Health *Health `protobuf:"bytes,13,opt,name=health,proto3,oneof"`
}

type Datapoint_Json struct {
	Json *Json `protobuf:"bytes,14,opt,name=json,proto3,oneof"`
}

type Datapoint_Battery struct {
	Battery *Battery `protobuf:"bytes,15,opt,name=battery,proto3,oneof"`
}

type Datapoint_Video struct {
	Video *Video `protobuf:"bytes,16,opt,name=video,proto3,oneof"`
}

func (*Datapoint_Text) isDatapoint_Data() {}

func (*Datapoint_Numeric) isDatapoint_Data() {}

func (*Datapoint_NumericSet) isDatapoint_Data() {}

func (*Datapoint_Bitset) isDatapoint_Data() {}

func (*Datapoint_File) isDatapoint_Data() {}

func (*Datapoint_Image) isDatapoint_Data() {}

func (*Datapoint_PointCloud) isDatapoint_Data() {}

func (*Datapoint_Location) isDatapoint_Data() {}

func (*Datapoint_Localization) isDatapoint_Data() {}

func (*Datapoint_Health) isDatapoint_Data() {}

func (*Datapoint_Json) isDatapoint_Data() {}

func (*Datapoint_Battery) isDatapoint_Data() {}

func (*Datapoint_Video) isDatapoint_Data() {}

func (m *Datapoint) GetData() isDatapoint_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Datapoint) GetText() *Text {
	if x, ok := m.GetData().(*Datapoint_Text); ok {
		return x.Text
	}
	return nil
}

func (m *Datapoint) GetNumeric() *Numeric {
	if x, ok := m.GetData().(*Datapoint_Numeric); ok {
		return x.Numeric
	}
	return nil
}

func (m *Datapoint) GetNumericSet() *NumericSet {
	if x, ok := m.GetData().(*Datapoint_NumericSet); ok {
		return x.NumericSet
	}
	return nil
}

func (m *Datapoint) GetBitset() *Bitset {
	if x, ok := m.GetData().(*Datapoint_Bitset); ok {
		return x.Bitset
	}
	return nil
}

func (m *Datapoint) GetFile() *File {
	if x, ok := m.GetData().(*Datapoint_File); ok {
		return x.File
	}
	return nil
}

func (m *Datapoint) GetImage() *Image {
	if x, ok := m.GetData().(*Datapoint_Image); ok {
		return x.Image
	}
	return nil
}

func (m *Datapoint) GetPointCloud() *PointCloud {
	if x, ok := m.GetData().(*Datapoint_PointCloud); ok {
		return x.PointCloud
	}
	return nil
}

func (m *Datapoint) GetLocation() *Location {
	if x, ok := m.GetData().(*Datapoint_Location); ok {
		return x.Location
	}
	return nil
}

func (m *Datapoint) GetLocalization() *Localization {
	if x, ok := m.GetData().(*Datapoint_Localization); ok {
		return x.Localization
	}
	return nil
}

func (m *Datapoint) GetHealth() *Health {
	if x, ok := m.GetData().(*Datapoint_Health); ok {
		return x.Health
	}
	return nil
}

func (m *Datapoint) GetJson() *Json {
	if x, ok := m.GetData().(*Datapoint_Json); ok {
		return x.Json
	}
	return nil
}

func (m *Datapoint) GetBattery() *Battery {
	if x, ok := m.GetData().(*Datapoint_Battery); ok {
		return x.Battery
	}
	return nil
}

func (m *Datapoint) GetVideo() *Video {
	if x, ok := m.GetData().(*Datapoint_Video); ok {
		return x.Video
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Datapoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Datapoint_Text)(nil),
		(*Datapoint_Numeric)(nil),
		(*Datapoint_NumericSet)(nil),
		(*Datapoint_Bitset)(nil),
		(*Datapoint_File)(nil),
		(*Datapoint_Image)(nil),
		(*Datapoint_PointCloud)(nil),
		(*Datapoint_Location)(nil),
		(*Datapoint_Localization)(nil),
		(*Datapoint_Health)(nil),
		(*Datapoint_Json)(nil),
		(*Datapoint_Battery)(nil),
		(*Datapoint_Video)(nil),
	}
}

func init() {
	proto.RegisterType((*Datapoint)(nil), "v1.model.Datapoint")
	proto.RegisterMapType((map[string]string)(nil), "v1.model.Datapoint.TagsEntry")
}

func init() { proto.RegisterFile("protos/model/v1/datapoint.proto", fileDescriptor_d4daa958809a1527) }

var fileDescriptor_d4daa958809a1527 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x51, 0x6f, 0xd3, 0x30,
	0x14, 0x85, 0xd3, 0xb5, 0x4d, 0x1b, 0x77, 0x6c, 0x9d, 0x35, 0x4d, 0x56, 0x19, 0xa2, 0x42, 0x48,
	0x14, 0x10, 0x09, 0x1d, 0x0f, 0x43, 0x88, 0xa7, 0x02, 0x53, 0x36, 0x21, 0x40, 0x66, 0xe2, 0x81,
	0x97, 0xc9, 0x6d, 0x4c, 0x6a, 0x48, 0xe2, 0x2a, 0x71, 0xa3, 0x95, 0x3f, 0xc2, 0xdf, 0x45, 0xf7,
	0x26, 0x4d, 0xb4, 0x54, 0x7b, 0xb3, 0xef, 0x77, 0x8e, 0x73, 0x73, 0x7c, 0x4d, 0x1e, 0xaf, 0x52,
	0x6d, 0x74, 0xe6, 0xc5, 0x3a, 0x90, 0x91, 0x97, 0x4f, 0xbd, 0x40, 0x18, 0xb1, 0xd2, 0x2a, 0x31,
	0x2e, 0x12, 0xda, 0xcf, 0xa7, 0x2e, 0xc2, 0xd1, 0xa8, 0x29, 0xfd, 0xa5, 0x22, 0x59, 0xa8, 0x46,
	0xa7, 0x4d, 0xb6, 0x94, 0x22, 0x32, 0xcb, 0x92, 0xee, 0x38, 0x63, 0x51, 0xb1, 0x71, 0x93, 0x25,
	0x22, 0x57, 0xa1, 0x30, 0x4a, 0x27, 0xf7, 0xb9, 0x8d, 0xbc, 0x2d, 0xbb, 0x1b, 0x3d, 0xdc, 0x39,
	0x59, 0x06, 0x4a, 0x14, 0xf0, 0xc9, 0x3f, 0x9b, 0x38, 0x1f, 0xb7, 0xbf, 0x43, 0x4f, 0x88, 0x9d,
	0x99, 0x54, 0x8a, 0x98, 0xb5, 0xc6, 0xad, 0x89, 0xc3, 0xcb, 0x1d, 0x3d, 0x25, 0x8e, 0x51, 0xb1,
	0xcc, 0x8c, 0x88, 0x57, 0x6c, 0x6f, 0xdc, 0x9a, 0xb4, 0x79, 0x5d, 0xa0, 0x53, 0xd2, 0x31, 0x22,
	0xcc, 0x58, 0x7b, 0xdc, 0x9e, 0x0c, 0xce, 0x1e, 0xb9, 0xdb, 0x34, 0xdc, 0xea, 0x60, 0xf7, 0x5a,
	0x84, 0xd9, 0xa7, 0xc4, 0xa4, 0x1b, 0x8e, 0x52, 0xfa, 0x94, 0x74, 0xa0, 0x43, 0xd6, 0x19, 0xb7,
	0x26, 0x83, 0xb3, 0x83, 0xda, 0x72, 0x2d, 0x6f, 0x8d, 0x6f, 0x71, 0xa4, 0xf4, 0x15, 0xe9, 0x25,
	0xeb, 0x58, 0xa6, 0x6a, 0xc1, 0xba, 0x28, 0x3c, 0xaa, 0x85, 0x5f, 0x0a, 0xe0, 0x5b, 0x7c, 0xab,
	0xa1, 0xe7, 0x64, 0x50, 0x2e, 0x6f, 0x32, 0x69, 0xd8, 0x11, 0x5a, 0x8e, 0x77, 0x2c, 0xdf, 0x25,
	0x7c, 0x81, 0x24, 0xd5, 0x8e, 0xbe, 0x20, 0xf6, 0x5c, 0x19, 0xf0, 0xf4, 0xd0, 0x33, 0xac, 0x3d,
	0x33, 0xac, 0xfb, 0x16, 0x2f, 0x15, 0xd0, 0x39, 0xdc, 0x29, 0xeb, 0x37, 0x3b, 0xbf, 0x50, 0x91,
	0x84, 0xce, 0x81, 0xd2, 0x67, 0xa4, 0xab, 0x62, 0x11, 0x4a, 0xe6, 0xa0, 0xec, 0xb0, 0x96, 0x5d,
	0x42, 0xd9, 0xb7, 0x78, 0xc1, 0xa1, 0x67, 0x4c, 0xe8, 0x66, 0x11, 0xe9, 0x75, 0xc0, 0x48, 0xb3,
	0xe7, 0x6f, 0x00, 0x3f, 0x00, 0x83, 0x9e, 0x57, 0xd5, 0x8e, 0xbe, 0x26, 0xfd, 0x48, 0x2f, 0x70,
	0x06, 0xd8, 0x00, 0x5d, 0xb4, 0x76, 0x7d, 0x2e, 0x89, 0x6f, 0xf1, 0x4a, 0x45, 0xdf, 0x93, 0x7d,
	0x58, 0x47, 0xea, 0x6f, 0xe1, 0xda, 0x47, 0xd7, 0xc9, 0x5d, 0xd7, 0x96, 0xfa, 0x16, 0xbf, 0xa3,
	0x86, 0x8c, 0x8a, 0x79, 0x65, 0x0f, 0x9a, 0x19, 0xf9, 0x58, 0x87, 0x8c, 0x0a, 0x05, 0x64, 0xf4,
	0x3b, 0xd3, 0x09, 0x3b, 0x68, 0x66, 0x74, 0x95, 0xe1, 0xc9, 0x48, 0xe1, 0x76, 0xe7, 0xc2, 0x18,
	0x99, 0x6e, 0xd8, 0x61, 0xf3, 0x76, 0x67, 0x05, 0x80, 0xdb, 0x2d, 0x35, 0x10, 0x69, 0xae, 0x02,
	0xa9, 0xd9, 0xb0, 0x19, 0xe9, 0x0f, 0x28, 0x43, 0xa4, 0xc8, 0x47, 0xe7, 0xc4, 0xa9, 0xc6, 0x8d,
	0x0e, 0x49, 0xfb, 0x8f, 0xdc, 0x94, 0xe3, 0x0c, 0x4b, 0x7a, 0x4c, 0xba, 0xb9, 0x88, 0xd6, 0x12,
	0xe7, 0xd8, 0xe1, 0xc5, 0xe6, 0xdd, 0xde, 0xdb, 0xd6, 0xcc, 0x26, 0x1d, 0x78, 0xd9, 0x57, 0x9d,
	0xbe, 0x3d, 0xec, 0xcd, 0x5e, 0xfe, 0x7c, 0x1e, 0x2a, 0xb3, 0x5c, 0xcf, 0xdd, 0x85, 0x8e, 0xbd,
	0x0b, 0x9d, 0xc6, 0x22, 0x31, 0x97, 0x5f, 0xbd, 0x50, 0x26, 0xf8, 0x74, 0xbc, 0x50, 0xe3, 0x63,
	0x82, 0x16, 0xe6, 0x36, 0x96, 0xde, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x92, 0xfb, 0x67,
	0x2b, 0x04, 0x00, 0x00,
}
