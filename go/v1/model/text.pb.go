// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/model/v1/text.proto

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

type Text struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_e75f6da50fa3739c, []int{0}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Text)(nil), "v1.model.Text")
}

func init() { proto.RegisterFile("protos/model/v1/text.proto", fileDescriptor_e75f6da50fa3739c) }

var fileDescriptor_e75f6da50fa3739c = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0x2f, 0x33, 0xd4, 0x2f, 0x49, 0xad, 0x28,
	0xd1, 0x03, 0x0b, 0x0a, 0x71, 0x94, 0x19, 0xea, 0x81, 0xc5, 0x95, 0x64, 0xb8, 0x58, 0x42, 0x52,
	0x2b, 0x4a, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x20, 0x1c, 0x27, 0xed, 0x28, 0xcd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0x7d, 0xb7, 0xfc, 0xa2, 0xdc, 0xc4, 0xbc, 0x12, 0x4f, 0x7f, 0xfd, 0xf4, 0xd4, 0x3c,
	0xb0, 0x41, 0xfa, 0xe9, 0xf9, 0x20, 0x93, 0xc1, 0x46, 0x25, 0xb1, 0x81, 0x85, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x09, 0x2e, 0x85, 0xba, 0x79, 0x00, 0x00, 0x00,
}
