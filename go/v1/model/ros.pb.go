// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/model/v1/ros.proto

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

type ROSTopicType int32

const (
	ROSTopicType_UNKNOWN                      ROSTopicType = 0
	ROSTopicType_STD_MSGS_BOOL                ROSTopicType = 1
	ROSTopicType_SENSOR_MSGS_COMPRESSED_IMAGE ROSTopicType = 2
	ROSTopicType_STD_MSGS_STRING              ROSTopicType = 3
	ROSTopicType_GEOMETRY_MSGS_POSE           ROSTopicType = 4
	ROSTopicType_ACTIONLIB_MSGS_GOALID        ROSTopicType = 5
	ROSTopicType_GEOMETRY_MSGS_TWIST          ROSTopicType = 6
	ROSTopicType_H264_VIDEO_FRAME             ROSTopicType = 7
)

var ROSTopicType_name = map[int32]string{
	0: "UNKNOWN",
	1: "STD_MSGS_BOOL",
	2: "SENSOR_MSGS_COMPRESSED_IMAGE",
	3: "STD_MSGS_STRING",
	4: "GEOMETRY_MSGS_POSE",
	5: "ACTIONLIB_MSGS_GOALID",
	6: "GEOMETRY_MSGS_TWIST",
	7: "H264_VIDEO_FRAME",
}

var ROSTopicType_value = map[string]int32{
	"UNKNOWN":                      0,
	"STD_MSGS_BOOL":                1,
	"SENSOR_MSGS_COMPRESSED_IMAGE": 2,
	"STD_MSGS_STRING":              3,
	"GEOMETRY_MSGS_POSE":           4,
	"ACTIONLIB_MSGS_GOALID":        5,
	"GEOMETRY_MSGS_TWIST":          6,
	"H264_VIDEO_FRAME":             7,
}

func (x ROSTopicType) String() string {
	return proto.EnumName(ROSTopicType_name, int32(x))
}

func (ROSTopicType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9010adf44cfe47e0, []int{0}
}

type ROSTopic struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ROSTopic) Reset()         { *m = ROSTopic{} }
func (m *ROSTopic) String() string { return proto.CompactTextString(m) }
func (*ROSTopic) ProtoMessage()    {}
func (*ROSTopic) Descriptor() ([]byte, []int) {
	return fileDescriptor_9010adf44cfe47e0, []int{0}
}

func (m *ROSTopic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ROSTopic.Unmarshal(m, b)
}
func (m *ROSTopic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ROSTopic.Marshal(b, m, deterministic)
}
func (m *ROSTopic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ROSTopic.Merge(m, src)
}
func (m *ROSTopic) XXX_Size() int {
	return xxx_messageInfo_ROSTopic.Size(m)
}
func (m *ROSTopic) XXX_DiscardUnknown() {
	xxx_messageInfo_ROSTopic.DiscardUnknown(m)
}

var xxx_messageInfo_ROSTopic proto.InternalMessageInfo

func (m *ROSTopic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ROSTopic) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ROSLocalization struct {
	MapTopic             string   `protobuf:"bytes,1,opt,name=map_topic,json=mapTopic,proto3" json:"map_topic,omitempty"`
	OdomTopic            string   `protobuf:"bytes,2,opt,name=odom_topic,json=odomTopic,proto3" json:"odom_topic,omitempty"`
	PointCloudTopics     []string `protobuf:"bytes,3,rep,name=point_cloud_topics,json=pointCloudTopics,proto3" json:"point_cloud_topics,omitempty"`
	PathTopic            string   `protobuf:"bytes,4,opt,name=path_topic,json=pathTopic,proto3" json:"path_topic,omitempty"`
	GoalTopic            string   `protobuf:"bytes,5,opt,name=goal_topic,json=goalTopic,proto3" json:"goal_topic,omitempty"`
	BaseReferenceFrame   string   `protobuf:"bytes,6,opt,name=base_reference_frame,json=baseReferenceFrame,proto3" json:"base_reference_frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ROSLocalization) Reset()         { *m = ROSLocalization{} }
func (m *ROSLocalization) String() string { return proto.CompactTextString(m) }
func (*ROSLocalization) ProtoMessage()    {}
func (*ROSLocalization) Descriptor() ([]byte, []int) {
	return fileDescriptor_9010adf44cfe47e0, []int{1}
}

func (m *ROSLocalization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ROSLocalization.Unmarshal(m, b)
}
func (m *ROSLocalization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ROSLocalization.Marshal(b, m, deterministic)
}
func (m *ROSLocalization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ROSLocalization.Merge(m, src)
}
func (m *ROSLocalization) XXX_Size() int {
	return xxx_messageInfo_ROSLocalization.Size(m)
}
func (m *ROSLocalization) XXX_DiscardUnknown() {
	xxx_messageInfo_ROSLocalization.DiscardUnknown(m)
}

var xxx_messageInfo_ROSLocalization proto.InternalMessageInfo

func (m *ROSLocalization) GetMapTopic() string {
	if m != nil {
		return m.MapTopic
	}
	return ""
}

func (m *ROSLocalization) GetOdomTopic() string {
	if m != nil {
		return m.OdomTopic
	}
	return ""
}

func (m *ROSLocalization) GetPointCloudTopics() []string {
	if m != nil {
		return m.PointCloudTopics
	}
	return nil
}

func (m *ROSLocalization) GetPathTopic() string {
	if m != nil {
		return m.PathTopic
	}
	return ""
}

func (m *ROSLocalization) GetGoalTopic() string {
	if m != nil {
		return m.GoalTopic
	}
	return ""
}

func (m *ROSLocalization) GetBaseReferenceFrame() string {
	if m != nil {
		return m.BaseReferenceFrame
	}
	return ""
}

type ROSTransformTree struct {
	BaseReferenceFrame   string   `protobuf:"bytes,1,opt,name=base_reference_frame,json=baseReferenceFrame,proto3" json:"base_reference_frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ROSTransformTree) Reset()         { *m = ROSTransformTree{} }
func (m *ROSTransformTree) String() string { return proto.CompactTextString(m) }
func (*ROSTransformTree) ProtoMessage()    {}
func (*ROSTransformTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_9010adf44cfe47e0, []int{2}
}

func (m *ROSTransformTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ROSTransformTree.Unmarshal(m, b)
}
func (m *ROSTransformTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ROSTransformTree.Marshal(b, m, deterministic)
}
func (m *ROSTransformTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ROSTransformTree.Merge(m, src)
}
func (m *ROSTransformTree) XXX_Size() int {
	return xxx_messageInfo_ROSTransformTree.Size(m)
}
func (m *ROSTransformTree) XXX_DiscardUnknown() {
	xxx_messageInfo_ROSTransformTree.DiscardUnknown(m)
}

var xxx_messageInfo_ROSTransformTree proto.InternalMessageInfo

func (m *ROSTransformTree) GetBaseReferenceFrame() string {
	if m != nil {
		return m.BaseReferenceFrame
	}
	return ""
}

type ROSMessageToPublish struct {
	Stream    string `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	FrameId   string `protobuf:"bytes,7,opt,name=frame_id,json=frameId,proto3" json:"frame_id,omitempty"`
	Timestamp uint64 `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*ROSMessageToPublish_Twist
	//	*ROSMessageToPublish_Bool
	//	*ROSMessageToPublish_CompressedImage
	//	*ROSMessageToPublish_Text
	//	*ROSMessageToPublish_Pose
	//	*ROSMessageToPublish_GoalID
	Data                 isROSMessageToPublish_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ROSMessageToPublish) Reset()         { *m = ROSMessageToPublish{} }
func (m *ROSMessageToPublish) String() string { return proto.CompactTextString(m) }
func (*ROSMessageToPublish) ProtoMessage()    {}
func (*ROSMessageToPublish) Descriptor() ([]byte, []int) {
	return fileDescriptor_9010adf44cfe47e0, []int{3}
}

func (m *ROSMessageToPublish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ROSMessageToPublish.Unmarshal(m, b)
}
func (m *ROSMessageToPublish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ROSMessageToPublish.Marshal(b, m, deterministic)
}
func (m *ROSMessageToPublish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ROSMessageToPublish.Merge(m, src)
}
func (m *ROSMessageToPublish) XXX_Size() int {
	return xxx_messageInfo_ROSMessageToPublish.Size(m)
}
func (m *ROSMessageToPublish) XXX_DiscardUnknown() {
	xxx_messageInfo_ROSMessageToPublish.DiscardUnknown(m)
}

var xxx_messageInfo_ROSMessageToPublish proto.InternalMessageInfo

func (m *ROSMessageToPublish) GetStream() string {
	if m != nil {
		return m.Stream
	}
	return ""
}

func (m *ROSMessageToPublish) GetFrameId() string {
	if m != nil {
		return m.FrameId
	}
	return ""
}

func (m *ROSMessageToPublish) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type isROSMessageToPublish_Data interface {
	isROSMessageToPublish_Data()
}

type ROSMessageToPublish_Twist struct {
	Twist *Twist `protobuf:"bytes,2,opt,name=twist,proto3,oneof"`
}

type ROSMessageToPublish_Bool struct {
	Bool bool `protobuf:"varint,3,opt,name=bool,proto3,oneof"`
}

type ROSMessageToPublish_CompressedImage struct {
	CompressedImage []byte `protobuf:"bytes,4,opt,name=compressed_image,json=compressedImage,proto3,oneof"`
}

type ROSMessageToPublish_Text struct {
	Text string `protobuf:"bytes,5,opt,name=text,proto3,oneof"`
}

type ROSMessageToPublish_Pose struct {
	Pose *Transform `protobuf:"bytes,6,opt,name=pose,proto3,oneof"`
}

type ROSMessageToPublish_GoalID struct {
	GoalID *GoalID `protobuf:"bytes,9,opt,name=goalID,proto3,oneof"`
}

func (*ROSMessageToPublish_Twist) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_Bool) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_CompressedImage) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_Text) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_Pose) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_GoalID) isROSMessageToPublish_Data() {}

func (m *ROSMessageToPublish) GetData() isROSMessageToPublish_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ROSMessageToPublish) GetTwist() *Twist {
	if x, ok := m.GetData().(*ROSMessageToPublish_Twist); ok {
		return x.Twist
	}
	return nil
}

func (m *ROSMessageToPublish) GetBool() bool {
	if x, ok := m.GetData().(*ROSMessageToPublish_Bool); ok {
		return x.Bool
	}
	return false
}

func (m *ROSMessageToPublish) GetCompressedImage() []byte {
	if x, ok := m.GetData().(*ROSMessageToPublish_CompressedImage); ok {
		return x.CompressedImage
	}
	return nil
}

func (m *ROSMessageToPublish) GetText() string {
	if x, ok := m.GetData().(*ROSMessageToPublish_Text); ok {
		return x.Text
	}
	return ""
}

func (m *ROSMessageToPublish) GetPose() *Transform {
	if x, ok := m.GetData().(*ROSMessageToPublish_Pose); ok {
		return x.Pose
	}
	return nil
}

func (m *ROSMessageToPublish) GetGoalID() *GoalID {
	if x, ok := m.GetData().(*ROSMessageToPublish_GoalID); ok {
		return x.GoalID
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ROSMessageToPublish) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ROSMessageToPublish_Twist)(nil),
		(*ROSMessageToPublish_Bool)(nil),
		(*ROSMessageToPublish_CompressedImage)(nil),
		(*ROSMessageToPublish_Text)(nil),
		(*ROSMessageToPublish_Pose)(nil),
		(*ROSMessageToPublish_GoalID)(nil),
	}
}

func init() {
	proto.RegisterEnum("v1.model.ROSTopicType", ROSTopicType_name, ROSTopicType_value)
	proto.RegisterType((*ROSTopic)(nil), "v1.model.ROSTopic")
	proto.RegisterType((*ROSLocalization)(nil), "v1.model.ROSLocalization")
	proto.RegisterType((*ROSTransformTree)(nil), "v1.model.ROSTransformTree")
	proto.RegisterType((*ROSMessageToPublish)(nil), "v1.model.ROSMessageToPublish")
}

func init() { proto.RegisterFile("protos/model/v1/ros.proto", fileDescriptor_9010adf44cfe47e0) }

var fileDescriptor_9010adf44cfe47e0 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5d, 0x6f, 0xda, 0x3c,
	0x1c, 0xc5, 0x49, 0xa1, 0xbc, 0xfc, 0xdb, 0x47, 0xe4, 0x31, 0x5d, 0x47, 0xbb, 0x4e, 0x42, 0xbd,
	0x19, 0x6d, 0x27, 0x58, 0xd9, 0xb4, 0x7b, 0x28, 0x29, 0x44, 0x03, 0x5c, 0x39, 0xd9, 0xaa, 0xed,
	0x26, 0x32, 0xe0, 0x42, 0x24, 0x1c, 0x47, 0xb1, 0xdb, 0xbd, 0x7c, 0xc1, 0x7e, 0x98, 0x7d, 0x89,
	0xc9, 0x36, 0xb4, 0x5b, 0xa5, 0xde, 0xf9, 0x7f, 0x7e, 0xc7, 0x07, 0x73, 0x62, 0xc3, 0x41, 0x9a,
	0x09, 0x25, 0x64, 0x9b, 0x8b, 0x39, 0x5b, 0xb5, 0xef, 0xce, 0xdb, 0x99, 0x90, 0x2d, 0xa3, 0xa1,
	0xf2, 0xdd, 0x79, 0xcb, 0xc8, 0x87, 0x87, 0x4f, 0x4d, 0x9c, 0xaa, 0xa5, 0x75, 0x1d, 0x36, 0x9e,
	0xb2, 0x84, 0xde, 0xc5, 0x0b, 0xaa, 0x62, 0x91, 0x58, 0xc7, 0x71, 0x07, 0xca, 0x04, 0x07, 0xa1,
	0x48, 0xe3, 0x19, 0x42, 0x50, 0x48, 0x28, 0x67, 0x75, 0xa7, 0xe1, 0x34, 0x2b, 0xc4, 0xac, 0xb5,
	0x96, 0x52, 0xb5, 0xac, 0x6f, 0x59, 0x4d, 0xaf, 0x8f, 0x7f, 0x3b, 0x50, 0x25, 0x38, 0x18, 0x89,
	0x19, 0x5d, 0xc5, 0xbf, 0x4c, 0x1a, 0x7a, 0x05, 0x15, 0x4e, 0xd3, 0x48, 0xe9, 0xa0, 0x75, 0x40,
	0x99, 0xd3, 0xd4, 0x06, 0xbf, 0x06, 0x10, 0x73, 0xc1, 0xd7, 0xd4, 0x46, 0x55, 0xb4, 0x62, 0xf1,
	0x5b, 0x40, 0xa9, 0x88, 0x13, 0x15, 0xcd, 0x56, 0xe2, 0x76, 0x6e, 0x5d, 0xb2, 0x9e, 0x6f, 0xe4,
	0x9b, 0x15, 0xe2, 0x1a, 0x72, 0xa1, 0x81, 0x31, 0x4b, 0x1d, 0xa6, 0x4f, 0xb1, 0x0e, 0x2b, 0xd8,
	0x30, 0xad, 0x3c, 0xfc, 0xd6, 0x42, 0xd0, 0xd5, 0x1a, 0x6f, 0x5b, 0xac, 0x15, 0x8b, 0xdf, 0xc1,
	0xde, 0x94, 0x4a, 0x16, 0x65, 0xec, 0x86, 0x65, 0x2c, 0x99, 0xb1, 0xe8, 0x26, 0xd3, 0xff, 0xb9,
	0x68, 0x8c, 0x48, 0x33, 0xb2, 0x41, 0x97, 0x9a, 0x1c, 0xf7, 0xc1, 0xd5, 0x0d, 0x65, 0x34, 0x91,
	0x37, 0x22, 0xe3, 0x61, 0xc6, 0xd8, 0xb3, 0x29, 0xce, 0xb3, 0x29, 0xf7, 0x5b, 0x50, 0x23, 0x38,
	0x18, 0x33, 0x29, 0xe9, 0x82, 0x85, 0xe2, 0xea, 0x76, 0xba, 0x8a, 0xe5, 0x12, 0xed, 0x43, 0x51,
	0xaa, 0x8c, 0x51, 0xbe, 0xde, 0xbb, 0x9e, 0xd0, 0x01, 0x94, 0x4d, 0x64, 0x14, 0xcf, 0xeb, 0x25,
	0x43, 0x4a, 0x66, 0xf6, 0xe7, 0xe8, 0x08, 0x2a, 0x2a, 0xe6, 0x4c, 0x2a, 0xca, 0xd3, 0x7a, 0xb9,
	0xe1, 0x34, 0x0b, 0xe4, 0x51, 0x40, 0x6f, 0x60, 0x5b, 0x7d, 0x8f, 0xa5, 0x32, 0x35, 0xef, 0x74,
	0xaa, 0xad, 0xcd, 0x45, 0x69, 0x85, 0x5a, 0x1e, 0xe6, 0x88, 0xe5, 0x68, 0x0f, 0x0a, 0x53, 0x21,
	0x56, 0xf5, 0x7c, 0xc3, 0x69, 0x96, 0x87, 0x39, 0x62, 0x26, 0x74, 0x06, 0xee, 0x4c, 0xf0, 0x34,
	0x63, 0x52, 0xb2, 0x79, 0x14, 0x73, 0xba, 0x60, 0xa6, 0xe3, 0xdd, 0x61, 0x8e, 0x54, 0x1f, 0x89,
	0xaf, 0x81, 0x8e, 0x50, 0xec, 0x87, 0xb2, 0x2d, 0xeb, 0x08, 0x3d, 0xa1, 0x13, 0x28, 0xa4, 0x42,
	0xda, 0x4a, 0x77, 0x3a, 0xb5, 0xbf, 0x0e, 0xb0, 0xe9, 0x50, 0x5b, 0xb5, 0x05, 0x9d, 0x42, 0x51,
	0x7f, 0x1a, 0xbf, 0x5f, 0xaf, 0x18, 0xb3, 0xfb, 0x68, 0x1e, 0x18, 0x7d, 0x98, 0x23, 0x6b, 0x47,
	0xaf, 0x08, 0x85, 0x39, 0x55, 0xf4, 0xf4, 0xde, 0x81, 0xdd, 0xcd, 0x95, 0x0d, 0x7f, 0xa6, 0x0c,
	0xed, 0x40, 0xe9, 0xf3, 0xe4, 0xd3, 0x04, 0x5f, 0x4f, 0xdc, 0x1c, 0xfa, 0x1f, 0xfe, 0x0b, 0xc2,
	0x7e, 0x34, 0x0e, 0x06, 0x41, 0xd4, 0xc3, 0x78, 0xe4, 0x3a, 0xa8, 0x01, 0x47, 0x81, 0x37, 0x09,
	0x30, 0xb1, 0xea, 0x05, 0x1e, 0x5f, 0x11, 0x2f, 0x08, 0xbc, 0x7e, 0xe4, 0x8f, 0xbb, 0x03, 0xcf,
	0xdd, 0x42, 0x35, 0xa8, 0x3e, 0x6c, 0x0a, 0x42, 0xe2, 0x4f, 0x06, 0x6e, 0x1e, 0xed, 0x03, 0x1a,
	0x78, 0x78, 0xec, 0x85, 0xe4, 0xab, 0x25, 0x57, 0x38, 0xf0, 0xdc, 0x02, 0x3a, 0x80, 0x17, 0xdd,
	0x8b, 0xd0, 0xc7, 0x93, 0x91, 0xdf, 0xb3, 0x60, 0x80, 0xbb, 0x23, 0xbf, 0xef, 0x6e, 0xa3, 0x97,
	0x50, 0xfb, 0x77, 0x4b, 0x78, 0xed, 0x07, 0xa1, 0x5b, 0x44, 0x7b, 0xe0, 0x0e, 0x3b, 0x1f, 0x3f,
	0x44, 0x5f, 0xfc, 0xbe, 0x87, 0xa3, 0x4b, 0xd2, 0x1d, 0x7b, 0x6e, 0xa9, 0x77, 0xf6, 0xed, 0x64,
	0x11, 0xab, 0xe5, 0xed, 0xb4, 0x35, 0x13, 0xbc, 0x7d, 0x29, 0x32, 0x4e, 0x13, 0xe5, 0xe3, 0xf6,
	0x82, 0x25, 0xe6, 0x71, 0xb6, 0x17, 0xc2, 0xbc, 0x67, 0xdd, 0xc7, 0xb4, 0x68, 0xa4, 0xf7, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x11, 0x1e, 0x7f, 0x9b, 0x14, 0x04, 0x00, 0x00,
}
