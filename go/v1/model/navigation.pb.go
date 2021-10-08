// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.2
// source: protos/model/v1/navigation.proto

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Localization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Odometry    *Odometry     `protobuf:"bytes,1,opt,name=odometry,proto3" json:"odometry,omitempty"`
	Map         *Map          `protobuf:"bytes,2,opt,name=map,proto3" json:"map,omitempty"`
	PointClouds []*PointCloud `protobuf:"bytes,3,rep,name=point_clouds,json=pointClouds,proto3" json:"point_clouds,omitempty"`
	Path        *Path         `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Goal        *Goal         `protobuf:"bytes,5,opt,name=goal,proto3" json:"goal,omitempty"`
}

func (x *Localization) Reset() {
	*x = Localization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Localization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Localization) ProtoMessage() {}

func (x *Localization) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Localization.ProtoReflect.Descriptor instead.
func (*Localization) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{1}
}

func (x *Localization) GetOdometry() *Odometry {
	if x != nil {
		return x.Odometry
	}
	return nil
}

func (x *Localization) GetMap() *Map {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *Localization) GetPointClouds() []*PointCloud {
	if x != nil {
		return x.PointClouds
	}
	return nil
}

func (x *Localization) GetPath() *Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Localization) GetGoal() *Goal {
	if x != nil {
		return x.Goal
	}
	return nil
}

type Odometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pose         *Transform `protobuf:"bytes,1,opt,name=pose,proto3" json:"pose,omitempty"`
	Twist        *Twist     `protobuf:"bytes,2,opt,name=twist,proto3" json:"twist,omitempty"`
	WorldToLocal *Transform `protobuf:"bytes,3,opt,name=world_to_local,json=worldToLocal,proto3" json:"world_to_local,omitempty"`
}

func (x *Odometry) Reset() {
	*x = Odometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Odometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Odometry) ProtoMessage() {}

func (x *Odometry) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Odometry.ProtoReflect.Descriptor instead.
func (*Odometry) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{2}
}

func (x *Odometry) GetPose() *Transform {
	if x != nil {
		return x.Pose
	}
	return nil
}

func (x *Odometry) GetTwist() *Twist {
	if x != nil {
		return x.Twist
	}
	return nil
}

func (x *Odometry) GetWorldToLocal() *Transform {
	if x != nil {
		return x.WorldToLocal
	}
	return nil
}

type Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resolution   float64    `protobuf:"fixed64,1,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Width        uint32     `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height       uint32     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Origin       *Transform `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	WorldToLocal *Transform `protobuf:"bytes,5,opt,name=world_to_local,json=worldToLocal,proto3" json:"world_to_local,omitempty"`
	// Types that are assignable to Data:
	//	*Map_OccupancyGrid
	//	*Map_Url
	//	*Map_Raw
	Data isMap_Data `protobuf_oneof:"data"`
}

func (x *Map) Reset() {
	*x = Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Map) ProtoMessage() {}

func (x *Map) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Map.ProtoReflect.Descriptor instead.
func (*Map) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{3}
}

func (x *Map) GetResolution() float64 {
	if x != nil {
		return x.Resolution
	}
	return 0
}

func (x *Map) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Map) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Map) GetOrigin() *Transform {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Map) GetWorldToLocal() *Transform {
	if x != nil {
		return x.WorldToLocal
	}
	return nil
}

func (m *Map) GetData() isMap_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Map) GetOccupancyGrid() *OccupancyGrid {
	if x, ok := x.GetData().(*Map_OccupancyGrid); ok {
		return x.OccupancyGrid
	}
	return nil
}

func (x *Map) GetUrl() string {
	if x, ok := x.GetData().(*Map_Url); ok {
		return x.Url
	}
	return ""
}

func (x *Map) GetRaw() []byte {
	if x, ok := x.GetData().(*Map_Raw); ok {
		return x.Raw
	}
	return nil
}

type isMap_Data interface {
	isMap_Data()
}

type Map_OccupancyGrid struct {
	OccupancyGrid *OccupancyGrid `protobuf:"bytes,6,opt,name=occupancy_grid,json=occupancyGrid,proto3,oneof"`
}

type Map_Url struct {
	Url string `protobuf:"bytes,7,opt,name=url,proto3,oneof"`
}

type Map_Raw struct {
	Raw []byte `protobuf:"bytes,8,opt,name=raw,proto3,oneof"`
}

func (*Map_OccupancyGrid) isMap_Data() {}

func (*Map_Url) isMap_Data() {}

func (*Map_Raw) isMap_Data() {}

type OccupancyGrid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []int32 `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
}

func (x *OccupancyGrid) Reset() {
	*x = OccupancyGrid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupancyGrid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupancyGrid) ProtoMessage() {}

func (x *OccupancyGrid) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupancyGrid.ProtoReflect.Descriptor instead.
func (*OccupancyGrid) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{4}
}

func (x *OccupancyGrid) GetData() []int32 {
	if x != nil {
		return x.Data
	}
	return nil
}

type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldToLocal *Transform   `protobuf:"bytes,1,opt,name=world_to_local,json=worldToLocal,proto3" json:"world_to_local,omitempty"`
	Poses        []*Transform `protobuf:"bytes,2,rep,name=poses,proto3" json:"poses,omitempty"`
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{5}
}

func (x *Path) GetWorldToLocal() *Transform {
	if x != nil {
		return x.WorldToLocal
	}
	return nil
}

func (x *Path) GetPoses() []*Transform {
	if x != nil {
		return x.Poses
	}
	return nil
}

type JointState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldToLocal *Transform `protobuf:"bytes,5,opt,name=world_to_local,json=worldToLocal,proto3" json:"world_to_local,omitempty"`
	Name         []string   `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
	Position     []float64  `protobuf:"fixed64,2,rep,packed,name=position,proto3" json:"position,omitempty"`
	Velocity     []float64  `protobuf:"fixed64,3,rep,packed,name=velocity,proto3" json:"velocity,omitempty"`
	Effort       []float64  `protobuf:"fixed64,4,rep,packed,name=effort,proto3" json:"effort,omitempty"`
}

func (x *JointState) Reset() {
	*x = JointState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JointState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JointState) ProtoMessage() {}

func (x *JointState) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JointState.ProtoReflect.Descriptor instead.
func (*JointState) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{6}
}

func (x *JointState) GetWorldToLocal() *Transform {
	if x != nil {
		return x.WorldToLocal
	}
	return nil
}

func (x *JointState) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *JointState) GetPosition() []float64 {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *JointState) GetVelocity() []float64 {
	if x != nil {
		return x.Velocity
	}
	return nil
}

func (x *JointState) GetEffort() []float64 {
	if x != nil {
		return x.Effort
	}
	return nil
}

type Goal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldToLocal *Transform `protobuf:"bytes,1,opt,name=world_to_local,json=worldToLocal,proto3" json:"world_to_local,omitempty"`
	Pose         *Transform `protobuf:"bytes,2,opt,name=pose,proto3" json:"pose,omitempty"`
}

func (x *Goal) Reset() {
	*x = Goal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goal) ProtoMessage() {}

func (x *Goal) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goal.ProtoReflect.Descriptor instead.
func (*Goal) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{7}
}

func (x *Goal) GetWorldToLocal() *Transform {
	if x != nil {
		return x.WorldToLocal
	}
	return nil
}

func (x *Goal) GetPose() *Transform {
	if x != nil {
		return x.Pose
	}
	return nil
}

type GoalID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GoalID) Reset() {
	*x = GoalID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalID) ProtoMessage() {}

func (x *GoalID) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalID.ProtoReflect.Descriptor instead.
func (*GoalID) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{8}
}

func (x *GoalID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PoseWithCovariance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pose       *Transform `protobuf:"bytes,1,opt,name=pose,proto3" json:"pose,omitempty"`
	Covariance []float64  `protobuf:"fixed64,2,rep,packed,name=covariance,proto3" json:"covariance,omitempty"`
}

func (x *PoseWithCovariance) Reset() {
	*x = PoseWithCovariance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_navigation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoseWithCovariance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoseWithCovariance) ProtoMessage() {}

func (x *PoseWithCovariance) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_navigation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoseWithCovariance.ProtoReflect.Descriptor instead.
func (*PoseWithCovariance) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_navigation_proto_rawDescGZIP(), []int{9}
}

func (x *PoseWithCovariance) GetPose() *Transform {
	if x != nil {
		return x.Pose
	}
	return nil
}

func (x *PoseWithCovariance) GetCovariance() []float64 {
	if x != nil {
		return x.Covariance
	}
	return nil
}

var File_protos_model_v1_navigation_proto protoreflect.FileDescriptor

var file_protos_model_v1_navigation_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xe0, 0x01, 0x0a, 0x0c,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x08,
	0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x64, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x03,
	0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x37, 0x0a,
	0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x04, 0x67, 0x6f,
	0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x04, 0x67, 0x6f, 0x61, 0x6c, 0x22, 0x95,
	0x01, 0x0a, 0x08, 0x4f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x70,
	0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x04,
	0x70, 0x6f, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x74, 0x77, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54,
	0x77, 0x69, 0x73, 0x74, 0x52, 0x05, 0x74, 0x77, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0e, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x54,
	0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0xad, 0x02, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x39, 0x0a, 0x0e, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x54, 0x6f, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x12, 0x40, 0x0a, 0x0e, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63,
	0x79, 0x5f, 0x67, 0x72, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63,
	0x79, 0x47, 0x72, 0x69, 0x64, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e,
	0x63, 0x79, 0x47, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x03, 0x72, 0x61,
	0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x0d, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61,
	0x6e, 0x63, 0x79, 0x47, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6c, 0x0a, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x74, 0x6f, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x29,
	0x0a, 0x05, 0x70, 0x6f, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x65, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x0a, 0x4a, 0x6f,
	0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x54, 0x6f, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x06, 0x65, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x22, 0x6a, 0x0a, 0x04, 0x47, 0x6f, 0x61, 0x6c, 0x12,
	0x39, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0c, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x6f,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x47, 0x6f, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a,
	0x12, 0x50, 0x6f, 0x73, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01,
	0x52, 0x0a, 0x63, 0x6f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x74, 0x49, 0x4f, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_model_v1_navigation_proto_rawDescOnce sync.Once
	file_protos_model_v1_navigation_proto_rawDescData = file_protos_model_v1_navigation_proto_rawDesc
)

func file_protos_model_v1_navigation_proto_rawDescGZIP() []byte {
	file_protos_model_v1_navigation_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_navigation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_navigation_proto_rawDescData)
	})
	return file_protos_model_v1_navigation_proto_rawDescData
}

var file_protos_model_v1_navigation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_model_v1_navigation_proto_goTypes = []interface{}{
	(*Location)(nil),           // 0: v1.model.Location
	(*Localization)(nil),       // 1: v1.model.Localization
	(*Odometry)(nil),           // 2: v1.model.Odometry
	(*Map)(nil),                // 3: v1.model.Map
	(*OccupancyGrid)(nil),      // 4: v1.model.OccupancyGrid
	(*Path)(nil),               // 5: v1.model.Path
	(*JointState)(nil),         // 6: v1.model.JointState
	(*Goal)(nil),               // 7: v1.model.Goal
	(*GoalID)(nil),             // 8: v1.model.GoalID
	(*PoseWithCovariance)(nil), // 9: v1.model.PoseWithCovariance
	(*PointCloud)(nil),         // 10: v1.model.PointCloud
	(*Transform)(nil),          // 11: v1.model.Transform
	(*Twist)(nil),              // 12: v1.model.Twist
}
var file_protos_model_v1_navigation_proto_depIdxs = []int32{
	2,  // 0: v1.model.Localization.odometry:type_name -> v1.model.Odometry
	3,  // 1: v1.model.Localization.map:type_name -> v1.model.Map
	10, // 2: v1.model.Localization.point_clouds:type_name -> v1.model.PointCloud
	5,  // 3: v1.model.Localization.path:type_name -> v1.model.Path
	7,  // 4: v1.model.Localization.goal:type_name -> v1.model.Goal
	11, // 5: v1.model.Odometry.pose:type_name -> v1.model.Transform
	12, // 6: v1.model.Odometry.twist:type_name -> v1.model.Twist
	11, // 7: v1.model.Odometry.world_to_local:type_name -> v1.model.Transform
	11, // 8: v1.model.Map.origin:type_name -> v1.model.Transform
	11, // 9: v1.model.Map.world_to_local:type_name -> v1.model.Transform
	4,  // 10: v1.model.Map.occupancy_grid:type_name -> v1.model.OccupancyGrid
	11, // 11: v1.model.Path.world_to_local:type_name -> v1.model.Transform
	11, // 12: v1.model.Path.poses:type_name -> v1.model.Transform
	11, // 13: v1.model.JointState.world_to_local:type_name -> v1.model.Transform
	11, // 14: v1.model.Goal.world_to_local:type_name -> v1.model.Transform
	11, // 15: v1.model.Goal.pose:type_name -> v1.model.Transform
	11, // 16: v1.model.PoseWithCovariance.pose:type_name -> v1.model.Transform
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_protos_model_v1_navigation_proto_init() }
func file_protos_model_v1_navigation_proto_init() {
	if File_protos_model_v1_navigation_proto != nil {
		return
	}
	file_protos_model_v1_math_proto_init()
	file_protos_model_v1_media_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_navigation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Localization); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Odometry); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Map); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupancyGrid); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JointState); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Goal); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalID); i {
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
		file_protos_model_v1_navigation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoseWithCovariance); i {
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
	file_protos_model_v1_navigation_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Map_OccupancyGrid)(nil),
		(*Map_Url)(nil),
		(*Map_Raw)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_navigation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_navigation_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_navigation_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_navigation_proto_msgTypes,
	}.Build()
	File_protos_model_v1_navigation_proto = out.File
	file_protos_model_v1_navigation_proto_rawDesc = nil
	file_protos_model_v1_navigation_proto_goTypes = nil
	file_protos_model_v1_navigation_proto_depIdxs = nil
}
