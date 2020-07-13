// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/model/v1/event.proto

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

type Event struct {
	Timestamp            int64             `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	StreamName           string            `protobuf:"bytes,3,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	StreamType           string            `protobuf:"bytes,4,opt,name=stream_type,json=streamType,proto3" json:"stream_type,omitempty"`
	NotificationEnabled  bool              `protobuf:"varint,5,opt,name=notification_enabled,json=notificationEnabled,proto3" json:"notification_enabled,omitempty"`
	Tags                 map[string]string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1f400acba3d6e90, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Event) GetStreamName() string {
	if m != nil {
		return m.StreamName
	}
	return ""
}

func (m *Event) GetStreamType() string {
	if m != nil {
		return m.StreamType
	}
	return ""
}

func (m *Event) GetNotificationEnabled() bool {
	if m != nil {
		return m.NotificationEnabled
	}
	return false
}

func (m *Event) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "v1.model.Event")
	proto.RegisterMapType((map[string]string)(nil), "v1.model.Event.TagsEntry")
}

func init() { proto.RegisterFile("protos/model/v1/event.proto", fileDescriptor_f1f400acba3d6e90) }

var fileDescriptor_f1f400acba3d6e90 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0xba, 0xcd, 0x35, 0xbb, 0x48, 0xdc, 0x21, 0xfe, 0x01, 0x8b, 0xa7, 0x8a, 0x98,
	0x30, 0x3d, 0x28, 0x1e, 0x85, 0x0a, 0x5e, 0x14, 0xca, 0x4e, 0x5e, 0x46, 0xba, 0xbd, 0xc6, 0x62,
	0x93, 0x94, 0xe6, 0x5d, 0xa1, 0xdf, 0xc5, 0x0f, 0x2b, 0x4d, 0xa9, 0xf6, 0x96, 0xf7, 0xf7, 0x7b,
	0x20, 0x3c, 0x0f, 0x39, 0xaf, 0x6a, 0x8b, 0xd6, 0x09, 0x6d, 0xf7, 0x50, 0x8a, 0x66, 0x2d, 0xa0,
	0x01, 0x83, 0xdc, 0x53, 0xba, 0x68, 0xd6, 0xdc, 0x8b, 0xab, 0x9f, 0x09, 0x99, 0xa5, 0x9d, 0xa1,
	0x17, 0x24, 0xc2, 0x42, 0x83, 0x43, 0xa9, 0x2b, 0x16, 0xc4, 0x41, 0x12, 0x66, 0xff, 0x80, 0x32,
	0x72, 0xa4, 0xc1, 0x39, 0xa9, 0x80, 0x4d, 0xe2, 0x20, 0x89, 0xb2, 0xe1, 0xa4, 0x97, 0x64, 0xe9,
	0xb0, 0x06, 0xa9, 0xb7, 0x46, 0x6a, 0x60, 0xa1, 0xb7, 0xa4, 0x47, 0x6f, 0x52, 0x8f, 0x03, 0xd8,
	0x56, 0xc0, 0xa6, 0xe3, 0xc0, 0xa6, 0xad, 0x80, 0xae, 0xc9, 0xca, 0x58, 0x2c, 0x3e, 0x8b, 0x9d,
	0xc4, 0xc2, 0x9a, 0x2d, 0x18, 0x99, 0x97, 0xb0, 0x67, 0xb3, 0x38, 0x48, 0x16, 0xd9, 0xc9, 0xd8,
	0xa5, 0xbd, 0xa2, 0xb7, 0x64, 0x8a, 0x52, 0x39, 0x36, 0x8f, 0xc3, 0x64, 0x79, 0x77, 0xca, 0x87,
	0x3e, 0xdc, 0x77, 0xe1, 0x1b, 0xa9, 0x5c, 0x6a, 0xb0, 0x6e, 0x33, 0x1f, 0x3b, 0x7b, 0x20, 0xd1,
	0x1f, 0xa2, 0xc7, 0x24, 0xfc, 0x86, 0xd6, 0x57, 0x8c, 0xb2, 0xee, 0x49, 0x57, 0x64, 0xd6, 0xc8,
	0xf2, 0x30, 0x54, 0xeb, 0x8f, 0xa7, 0xc9, 0x63, 0xf0, 0x7c, 0xf3, 0x71, 0xad, 0x0a, 0xfc, 0x3a,
	0xe4, 0x7c, 0x67, 0xb5, 0x78, 0xb1, 0xb5, 0x96, 0x06, 0x5f, 0xdf, 0x85, 0x02, 0xe3, 0x97, 0x14,
	0xca, 0x76, 0xdb, 0xfa, 0xbf, 0xf3, 0xb9, 0x47, 0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1,
	0xac, 0x14, 0xaa, 0x7b, 0x01, 0x00, 0x00,
}