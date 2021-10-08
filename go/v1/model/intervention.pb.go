// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.2
// source: protos/model/v1/intervention.proto

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

type InterventionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Severity  Severity `protobuf:"varint,3,opt,name=severity,proto3,enum=v1.model.Severity" json:"severity,omitempty"`
	// Types that are assignable to Data:
	//	*InterventionRequest_SelectionRequest
	//	*InterventionRequest_LabelingRequest
	Data      isInterventionRequest_Data `protobuf_oneof:"data"`
	Tags      map[string]string          `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Responses []*InterventionResponse    `protobuf:"bytes,7,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *InterventionRequest) Reset() {
	*x = InterventionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterventionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterventionRequest) ProtoMessage() {}

func (x *InterventionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterventionRequest.ProtoReflect.Descriptor instead.
func (*InterventionRequest) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{0}
}

func (x *InterventionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InterventionRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *InterventionRequest) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_INFO
}

func (m *InterventionRequest) GetData() isInterventionRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *InterventionRequest) GetSelectionRequest() *SelectionRequest {
	if x, ok := x.GetData().(*InterventionRequest_SelectionRequest); ok {
		return x.SelectionRequest
	}
	return nil
}

func (x *InterventionRequest) GetLabelingRequest() *LabelingRequest {
	if x, ok := x.GetData().(*InterventionRequest_LabelingRequest); ok {
		return x.LabelingRequest
	}
	return nil
}

func (x *InterventionRequest) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *InterventionRequest) GetResponses() []*InterventionResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

type isInterventionRequest_Data interface {
	isInterventionRequest_Data()
}

type InterventionRequest_SelectionRequest struct {
	SelectionRequest *SelectionRequest `protobuf:"bytes,4,opt,name=selection_request,json=selectionRequest,proto3,oneof"`
}

type InterventionRequest_LabelingRequest struct {
	LabelingRequest *LabelingRequest `protobuf:"bytes,5,opt,name=labeling_request,json=labelingRequest,proto3,oneof"`
}

func (*InterventionRequest_SelectionRequest) isInterventionRequest_Data() {}

func (*InterventionRequest_LabelingRequest) isInterventionRequest_Data() {}

type InterventionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Data:
	//	*InterventionResponse_SelectionResponse
	//	*InterventionResponse_LabelingResponse
	Data isInterventionResponse_Data `protobuf_oneof:"data"`
}

func (x *InterventionResponse) Reset() {
	*x = InterventionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterventionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterventionResponse) ProtoMessage() {}

func (x *InterventionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterventionResponse.ProtoReflect.Descriptor instead.
func (*InterventionResponse) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{1}
}

func (x *InterventionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InterventionResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *InterventionResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (m *InterventionResponse) GetData() isInterventionResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *InterventionResponse) GetSelectionResponse() *SelectionResponse {
	if x, ok := x.GetData().(*InterventionResponse_SelectionResponse); ok {
		return x.SelectionResponse
	}
	return nil
}

func (x *InterventionResponse) GetLabelingResponse() *LabelingResponse {
	if x, ok := x.GetData().(*InterventionResponse_LabelingResponse); ok {
		return x.LabelingResponse
	}
	return nil
}

type isInterventionResponse_Data interface {
	isInterventionResponse_Data()
}

type InterventionResponse_SelectionResponse struct {
	SelectionResponse *SelectionResponse `protobuf:"bytes,4,opt,name=selection_response,json=selectionResponse,proto3,oneof"`
}

type InterventionResponse_LabelingResponse struct {
	LabelingResponse *LabelingResponse `protobuf:"bytes,5,opt,name=labeling_response,json=labelingResponse,proto3,oneof"`
}

func (*InterventionResponse_SelectionResponse) isInterventionResponse_Data() {}

func (*InterventionResponse_LabelingResponse) isInterventionResponse_Data() {}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{2}
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Label) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type LabeledPolygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vertices []*Vertex `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	Labels   []string  `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *LabeledPolygon) Reset() {
	*x = LabeledPolygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabeledPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabeledPolygon) ProtoMessage() {}

func (x *LabeledPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabeledPolygon.ProtoReflect.Descriptor instead.
func (*LabeledPolygon) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{3}
}

func (x *LabeledPolygon) GetVertices() []*Vertex {
	if x != nil {
		return x.Vertices
	}
	return nil
}

func (x *LabeledPolygon) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Vertex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Vertex) Reset() {
	*x = Vertex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vertex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vertex) ProtoMessage() {}

func (x *Vertex) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vertex.ProtoReflect.Descriptor instead.
func (*Vertex) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{4}
}

func (x *Vertex) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vertex) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type LabelingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Instruction string            `protobuf:"bytes,2,opt,name=instruction,proto3" json:"instruction,omitempty"`
	Image       *Image            `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Labels      []*Label          `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Hint        []*LabeledPolygon `protobuf:"bytes,5,rep,name=hint,proto3" json:"hint,omitempty"`
}

func (x *LabelingRequest) Reset() {
	*x = LabelingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelingRequest) ProtoMessage() {}

func (x *LabelingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelingRequest.ProtoReflect.Descriptor instead.
func (*LabelingRequest) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{5}
}

func (x *LabelingRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *LabelingRequest) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

func (x *LabelingRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *LabelingRequest) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *LabelingRequest) GetHint() []*LabeledPolygon {
	if x != nil {
		return x.Hint
	}
	return nil
}

type LabelingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*LabeledPolygon `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *LabelingResponse) Reset() {
	*x = LabelingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelingResponse) ProtoMessage() {}

func (x *LabelingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelingResponse.ProtoReflect.Descriptor instead.
func (*LabelingResponse) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{6}
}

func (x *LabelingResponse) GetValue() []*LabeledPolygon {
	if x != nil {
		return x.Value
	}
	return nil
}

type SelectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Types that are assignable to Data:
	//	*SelectionRequest_Image
	Data        isSelectionRequest_Data `protobuf_oneof:"data"`
	Instruction string                  `protobuf:"bytes,3,opt,name=instruction,proto3" json:"instruction,omitempty"`
	Options     []string                `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
	Hint        uint32                  `protobuf:"varint,5,opt,name=hint,proto3" json:"hint,omitempty"`
}

func (x *SelectionRequest) Reset() {
	*x = SelectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectionRequest) ProtoMessage() {}

func (x *SelectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectionRequest.ProtoReflect.Descriptor instead.
func (*SelectionRequest) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{7}
}

func (x *SelectionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (m *SelectionRequest) GetData() isSelectionRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *SelectionRequest) GetImage() *Image {
	if x, ok := x.GetData().(*SelectionRequest_Image); ok {
		return x.Image
	}
	return nil
}

func (x *SelectionRequest) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

func (x *SelectionRequest) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *SelectionRequest) GetHint() uint32 {
	if x != nil {
		return x.Hint
	}
	return 0
}

type isSelectionRequest_Data interface {
	isSelectionRequest_Data()
}

type SelectionRequest_Image struct {
	Image *Image `protobuf:"bytes,2,opt,name=image,proto3,oneof"`
}

func (*SelectionRequest_Image) isSelectionRequest_Data() {}

type SelectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SelectionResponse) Reset() {
	*x = SelectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_intervention_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectionResponse) ProtoMessage() {}

func (x *SelectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_intervention_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectionResponse.ProtoReflect.Descriptor instead.
func (*SelectionResponse) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_intervention_proto_rawDescGZIP(), []int{8}
}

func (x *SelectionResponse) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_protos_model_v1_intervention_proto protoreflect.FileDescriptor

var file_protos_model_v1_intervention_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x03, 0x0a, 0x13, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e,
	0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x49,
	0x0a, 0x11, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x10, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x0f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3c,
	0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x37, 0x0a, 0x09,
	0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x84, 0x02,
	0x0a, 0x14, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x4c, 0x0a, 0x12, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x11,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x11, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x10, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x0e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65,
	0x64, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52, 0x08, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x24,
	0x0a, 0x06, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x01, 0x79, 0x22, 0xc7, 0x01, 0x0a, 0x0f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x2c, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65,
	0x64, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x22, 0x42,
	0x0a, 0x10, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29,
	0x0a, 0x11, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x49,
	0x4f, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_intervention_proto_rawDescOnce sync.Once
	file_protos_model_v1_intervention_proto_rawDescData = file_protos_model_v1_intervention_proto_rawDesc
)

func file_protos_model_v1_intervention_proto_rawDescGZIP() []byte {
	file_protos_model_v1_intervention_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_intervention_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_intervention_proto_rawDescData)
	})
	return file_protos_model_v1_intervention_proto_rawDescData
}

var file_protos_model_v1_intervention_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_model_v1_intervention_proto_goTypes = []interface{}{
	(*InterventionRequest)(nil),  // 0: v1.model.InterventionRequest
	(*InterventionResponse)(nil), // 1: v1.model.InterventionResponse
	(*Label)(nil),                // 2: v1.model.Label
	(*LabeledPolygon)(nil),       // 3: v1.model.LabeledPolygon
	(*Vertex)(nil),               // 4: v1.model.Vertex
	(*LabelingRequest)(nil),      // 5: v1.model.LabelingRequest
	(*LabelingResponse)(nil),     // 6: v1.model.LabelingResponse
	(*SelectionRequest)(nil),     // 7: v1.model.SelectionRequest
	(*SelectionResponse)(nil),    // 8: v1.model.SelectionResponse
	nil,                          // 9: v1.model.InterventionRequest.TagsEntry
	(Severity)(0),                // 10: v1.model.Severity
	(*Image)(nil),                // 11: v1.model.Image
}
var file_protos_model_v1_intervention_proto_depIdxs = []int32{
	10, // 0: v1.model.InterventionRequest.severity:type_name -> v1.model.Severity
	7,  // 1: v1.model.InterventionRequest.selection_request:type_name -> v1.model.SelectionRequest
	5,  // 2: v1.model.InterventionRequest.labeling_request:type_name -> v1.model.LabelingRequest
	9,  // 3: v1.model.InterventionRequest.tags:type_name -> v1.model.InterventionRequest.TagsEntry
	1,  // 4: v1.model.InterventionRequest.responses:type_name -> v1.model.InterventionResponse
	8,  // 5: v1.model.InterventionResponse.selection_response:type_name -> v1.model.SelectionResponse
	6,  // 6: v1.model.InterventionResponse.labeling_response:type_name -> v1.model.LabelingResponse
	4,  // 7: v1.model.LabeledPolygon.vertices:type_name -> v1.model.Vertex
	11, // 8: v1.model.LabelingRequest.image:type_name -> v1.model.Image
	2,  // 9: v1.model.LabelingRequest.labels:type_name -> v1.model.Label
	3,  // 10: v1.model.LabelingRequest.hint:type_name -> v1.model.LabeledPolygon
	3,  // 11: v1.model.LabelingResponse.value:type_name -> v1.model.LabeledPolygon
	11, // 12: v1.model.SelectionRequest.image:type_name -> v1.model.Image
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_protos_model_v1_intervention_proto_init() }
func file_protos_model_v1_intervention_proto_init() {
	if File_protos_model_v1_intervention_proto != nil {
		return
	}
	file_protos_model_v1_media_proto_init()
	file_protos_model_v1_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_intervention_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterventionRequest); i {
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
		file_protos_model_v1_intervention_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterventionResponse); i {
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
		file_protos_model_v1_intervention_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_protos_model_v1_intervention_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabeledPolygon); i {
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
		file_protos_model_v1_intervention_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vertex); i {
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
		file_protos_model_v1_intervention_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelingRequest); i {
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
		file_protos_model_v1_intervention_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelingResponse); i {
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
		file_protos_model_v1_intervention_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectionRequest); i {
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
		file_protos_model_v1_intervention_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectionResponse); i {
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
	file_protos_model_v1_intervention_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*InterventionRequest_SelectionRequest)(nil),
		(*InterventionRequest_LabelingRequest)(nil),
	}
	file_protos_model_v1_intervention_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*InterventionResponse_SelectionResponse)(nil),
		(*InterventionResponse_LabelingResponse)(nil),
	}
	file_protos_model_v1_intervention_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*SelectionRequest_Image)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_intervention_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_intervention_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_intervention_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_intervention_proto_msgTypes,
	}.Build()
	File_protos_model_v1_intervention_proto = out.File
	file_protos_model_v1_intervention_proto_rawDesc = nil
	file_protos_model_v1_intervention_proto_goTypes = nil
	file_protos_model_v1_intervention_proto_depIdxs = nil
}
