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
	//	*Datapoint_MetricSet
	//	*Datapoint_Bitset
	//	*Datapoint_File
	//	*Datapoint_Image
	//	*Datapoint_PointCloud
	//	*Datapoint_Location
	//	*Datapoint_Localization
	//	*Datapoint_Health
	//	*Datapoint_Json
	//	*Datapoint_Battery
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

type Datapoint_MetricSet struct {
	MetricSet *MetricSet `protobuf:"bytes,6,opt,name=metric_set,json=metricSet,proto3,oneof"`
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

func (*Datapoint_Text) isDatapoint_Data() {}

func (*Datapoint_Numeric) isDatapoint_Data() {}

func (*Datapoint_MetricSet) isDatapoint_Data() {}

func (*Datapoint_Bitset) isDatapoint_Data() {}

func (*Datapoint_File) isDatapoint_Data() {}

func (*Datapoint_Image) isDatapoint_Data() {}

func (*Datapoint_PointCloud) isDatapoint_Data() {}

func (*Datapoint_Location) isDatapoint_Data() {}

func (*Datapoint_Localization) isDatapoint_Data() {}

func (*Datapoint_Health) isDatapoint_Data() {}

func (*Datapoint_Json) isDatapoint_Data() {}

func (*Datapoint_Battery) isDatapoint_Data() {}

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

func (m *Datapoint) GetMetricSet() *MetricSet {
	if x, ok := m.GetData().(*Datapoint_MetricSet); ok {
		return x.MetricSet
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

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Datapoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Datapoint_Text)(nil),
		(*Datapoint_Numeric)(nil),
		(*Datapoint_MetricSet)(nil),
		(*Datapoint_Bitset)(nil),
		(*Datapoint_File)(nil),
		(*Datapoint_Image)(nil),
		(*Datapoint_PointCloud)(nil),
		(*Datapoint_Location)(nil),
		(*Datapoint_Localization)(nil),
		(*Datapoint_Health)(nil),
		(*Datapoint_Json)(nil),
		(*Datapoint_Battery)(nil),
	}
}

func init() {
	proto.RegisterType((*Datapoint)(nil), "v1.model.Datapoint")
	proto.RegisterMapType((map[string]string)(nil), "v1.model.Datapoint.TagsEntry")
}

func init() { proto.RegisterFile("protos/model/v1/datapoint.proto", fileDescriptor_d4daa958809a1527) }

var fileDescriptor_d4daa958809a1527 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5b, 0x6f, 0xd3, 0x4c,
	0x10, 0x86, 0x9d, 0xe6, 0xd0, 0x78, 0xd2, 0xaf, 0xed, 0xb7, 0x54, 0xd5, 0x2a, 0x14, 0x61, 0x21,
	0x24, 0x02, 0x08, 0x9b, 0x14, 0xa4, 0x22, 0xc4, 0x55, 0x80, 0xca, 0x45, 0x9c, 0xb4, 0xf4, 0x8a,
	0x9b, 0x6a, 0xe3, 0x2c, 0xce, 0x82, 0xd7, 0x1b, 0xd9, 0x9b, 0xa8, 0xe1, 0xff, 0xf2, 0x3f, 0xd0,
	0x8e, 0x4f, 0xaa, 0x2b, 0xee, 0x3c, 0xf3, 0xbc, 0xef, 0x68, 0x3c, 0x33, 0x0b, 0xf7, 0x57, 0x99,
	0x36, 0x3a, 0x0f, 0x94, 0x5e, 0x88, 0x24, 0xd8, 0x4c, 0x83, 0x05, 0x37, 0x7c, 0xa5, 0x65, 0x6a,
	0x7c, 0x24, 0x64, 0xb8, 0x99, 0xfa, 0x08, 0xc7, 0xe3, 0xb6, 0xf4, 0x87, 0x4c, 0x44, 0xa1, 0x1a,
	0x9f, 0xb4, 0xd9, 0x52, 0xf0, 0xc4, 0x2c, 0x4b, 0x7a, 0xcb, 0xa9, 0x78, 0xcd, 0xbc, 0x36, 0x4b,
	0xf9, 0x46, 0xc6, 0xdc, 0x48, 0x9d, 0xfe, 0xcb, 0x6d, 0xc4, 0x75, 0xd9, 0xdd, 0xf8, 0xee, 0xad,
	0xca, 0x62, 0x21, 0x79, 0x01, 0x1f, 0xfc, 0xe9, 0x83, 0xfb, 0xae, 0xfa, 0x1d, 0x72, 0x0c, 0x83,
	0xdc, 0x64, 0x82, 0x2b, 0xda, 0xf1, 0x3a, 0x13, 0x97, 0x95, 0x11, 0x39, 0x01, 0xd7, 0x48, 0x25,
	0x72, 0xc3, 0xd5, 0x8a, 0xee, 0x78, 0x9d, 0x49, 0x97, 0x35, 0x09, 0x32, 0x85, 0x9e, 0xe1, 0x71,
	0x4e, 0xbb, 0x5e, 0x77, 0x32, 0x3a, 0xbd, 0xe7, 0x57, 0xd3, 0xf0, 0xeb, 0xc2, 0xfe, 0x25, 0x8f,
	0xf3, 0xf7, 0xa9, 0xc9, 0xb6, 0x0c, 0xa5, 0xe4, 0x21, 0xf4, 0x6c, 0x87, 0xb4, 0xe7, 0x75, 0x26,
	0xa3, 0xd3, 0xfd, 0xc6, 0x72, 0x29, 0xae, 0x4d, 0xe8, 0x30, 0xa4, 0xe4, 0x19, 0xec, 0xa6, 0x6b,
	0x25, 0x32, 0x19, 0xd1, 0x3e, 0x0a, 0xff, 0x6f, 0x84, 0x9f, 0x0b, 0x10, 0x3a, 0xac, 0xd2, 0x90,
	0x97, 0x00, 0x4a, 0x98, 0x4c, 0x46, 0x57, 0xb9, 0x30, 0x74, 0x80, 0x8e, 0x3b, 0x8d, 0xe3, 0x13,
	0xb2, 0x6f, 0xc2, 0xd6, 0x77, 0x55, 0x15, 0x90, 0x27, 0x30, 0x98, 0x4b, 0x63, 0x1d, 0xbb, 0xe8,
	0x38, 0x6c, 0x1c, 0x33, 0xcc, 0x87, 0x0e, 0x2b, 0x15, 0xb6, 0x6d, 0xbb, 0x50, 0x3a, 0x6c, 0xb7,
	0x7d, 0x2e, 0x13, 0x61, 0xdb, 0xb6, 0x94, 0x3c, 0x82, 0xbe, 0x54, 0x3c, 0x16, 0xd4, 0x45, 0xd9,
	0x41, 0x23, 0xbb, 0xb0, 0xe9, 0xd0, 0x61, 0x05, 0x27, 0x67, 0x30, 0xc2, 0xf1, 0x5c, 0x45, 0x89,
	0x5e, 0x2f, 0x28, 0xa0, 0xfc, 0xa8, 0x91, 0x7f, 0xb5, 0xf0, 0xad, 0x65, 0xa1, 0xc3, 0x60, 0x55,
	0x47, 0xe4, 0x39, 0x0c, 0x13, 0x1d, 0xe1, 0x01, 0xd0, 0x11, 0xba, 0x48, 0xe3, 0xfa, 0x58, 0x92,
	0xd0, 0x61, 0xb5, 0x8a, 0xbc, 0x81, 0x3d, 0xfb, 0x9d, 0xc8, 0xdf, 0x85, 0x6b, 0x0f, 0x5d, 0xc7,
	0x37, 0x5d, 0x15, 0x0d, 0x1d, 0x76, 0x43, 0x6d, 0x67, 0x54, 0x1c, 0x2b, 0xfd, 0xaf, 0x3d, 0xa3,
	0x10, 0xf3, 0x76, 0x46, 0x85, 0xc2, 0xce, 0xe8, 0x67, 0xae, 0x53, 0xba, 0xdf, 0x9e, 0xd1, 0x87,
	0x1c, 0x2b, 0x23, 0xb5, 0xab, 0x9d, 0x73, 0x63, 0x44, 0xb6, 0xa5, 0x07, 0xed, 0xd5, 0xce, 0x0a,
	0x60, 0x57, 0x5b, 0x6a, 0xc6, 0x67, 0xe0, 0xd6, 0x27, 0x44, 0x0e, 0xa1, 0xfb, 0x4b, 0x6c, 0xcb,
	0x13, 0xb5, 0x9f, 0xe4, 0x08, 0xfa, 0x1b, 0x9e, 0xac, 0x05, 0xde, 0xa6, 0xcb, 0x8a, 0xe0, 0xf5,
	0xce, 0xab, 0xce, 0x6c, 0x00, 0x3d, 0xfb, 0x5a, 0x67, 0x4f, 0xbf, 0x3f, 0x8e, 0xa5, 0x59, 0xae,
	0xe7, 0x7e, 0xa4, 0x55, 0x70, 0xae, 0x33, 0xc5, 0x53, 0x73, 0xf1, 0x25, 0x88, 0x45, 0x8a, 0x0f,
	0x21, 0x88, 0x35, 0x3e, 0x0d, 0xdb, 0xc0, 0x7c, 0x80, 0xa9, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0xa4, 0xe3, 0x3a, 0xf9, 0x03, 0x00, 0x00,
}
