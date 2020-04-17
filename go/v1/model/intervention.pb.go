// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/model/v1/intervention.proto

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

type Severity int32

const (
	Severity_INFO     Severity = 0
	Severity_WARNING  Severity = 1
	Severity_ERROR    Severity = 2
	Severity_CRITICAL Severity = 3
)

var Severity_name = map[int32]string{
	0: "INFO",
	1: "WARNING",
	2: "ERROR",
	3: "CRITICAL",
}

var Severity_value = map[string]int32{
	"INFO":     0,
	"WARNING":  1,
	"ERROR":    2,
	"CRITICAL": 3,
}

func (x Severity) String() string {
	return proto.EnumName(Severity_name, int32(x))
}

func (Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{0}
}

type InterventionRequest struct {
	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Severity  Severity `protobuf:"varint,3,opt,name=severity,proto3,enum=v1.model.Severity" json:"severity,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*InterventionRequest_SelectionRequest
	//	*InterventionRequest_LabelingRequest
	Data                 isInterventionRequest_Data `protobuf_oneof:"data"`
	Tags                 map[string]string          `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Responses            []*InterventionResponse    `protobuf:"bytes,7,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *InterventionRequest) Reset()         { *m = InterventionRequest{} }
func (m *InterventionRequest) String() string { return proto.CompactTextString(m) }
func (*InterventionRequest) ProtoMessage()    {}
func (*InterventionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{0}
}

func (m *InterventionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterventionRequest.Unmarshal(m, b)
}
func (m *InterventionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterventionRequest.Marshal(b, m, deterministic)
}
func (m *InterventionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterventionRequest.Merge(m, src)
}
func (m *InterventionRequest) XXX_Size() int {
	return xxx_messageInfo_InterventionRequest.Size(m)
}
func (m *InterventionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InterventionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InterventionRequest proto.InternalMessageInfo

func (m *InterventionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InterventionRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *InterventionRequest) GetSeverity() Severity {
	if m != nil {
		return m.Severity
	}
	return Severity_INFO
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

func (m *InterventionRequest) GetData() isInterventionRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InterventionRequest) GetSelectionRequest() *SelectionRequest {
	if x, ok := m.GetData().(*InterventionRequest_SelectionRequest); ok {
		return x.SelectionRequest
	}
	return nil
}

func (m *InterventionRequest) GetLabelingRequest() *LabelingRequest {
	if x, ok := m.GetData().(*InterventionRequest_LabelingRequest); ok {
		return x.LabelingRequest
	}
	return nil
}

func (m *InterventionRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *InterventionRequest) GetResponses() []*InterventionResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InterventionRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InterventionRequest_SelectionRequest)(nil),
		(*InterventionRequest_LabelingRequest)(nil),
	}
}

type InterventionResponse struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*InterventionResponse_SelectionResponse
	//	*InterventionResponse_LabelingResponse
	Data                 isInterventionResponse_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *InterventionResponse) Reset()         { *m = InterventionResponse{} }
func (m *InterventionResponse) String() string { return proto.CompactTextString(m) }
func (*InterventionResponse) ProtoMessage()    {}
func (*InterventionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{1}
}

func (m *InterventionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterventionResponse.Unmarshal(m, b)
}
func (m *InterventionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterventionResponse.Marshal(b, m, deterministic)
}
func (m *InterventionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterventionResponse.Merge(m, src)
}
func (m *InterventionResponse) XXX_Size() int {
	return xxx_messageInfo_InterventionResponse.Size(m)
}
func (m *InterventionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InterventionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InterventionResponse proto.InternalMessageInfo

func (m *InterventionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InterventionResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *InterventionResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
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

func (m *InterventionResponse) GetData() isInterventionResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InterventionResponse) GetSelectionResponse() *SelectionResponse {
	if x, ok := m.GetData().(*InterventionResponse_SelectionResponse); ok {
		return x.SelectionResponse
	}
	return nil
}

func (m *InterventionResponse) GetLabelingResponse() *LabelingResponse {
	if x, ok := m.GetData().(*InterventionResponse_LabelingResponse); ok {
		return x.LabelingResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InterventionResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InterventionResponse_SelectionResponse)(nil),
		(*InterventionResponse_LabelingResponse)(nil),
	}
}

type Label struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{2}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Label) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type LabeledPolygon struct {
	Vertices             []*Vertex `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	Labels               []string  `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LabeledPolygon) Reset()         { *m = LabeledPolygon{} }
func (m *LabeledPolygon) String() string { return proto.CompactTextString(m) }
func (*LabeledPolygon) ProtoMessage()    {}
func (*LabeledPolygon) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{3}
}

func (m *LabeledPolygon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabeledPolygon.Unmarshal(m, b)
}
func (m *LabeledPolygon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabeledPolygon.Marshal(b, m, deterministic)
}
func (m *LabeledPolygon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabeledPolygon.Merge(m, src)
}
func (m *LabeledPolygon) XXX_Size() int {
	return xxx_messageInfo_LabeledPolygon.Size(m)
}
func (m *LabeledPolygon) XXX_DiscardUnknown() {
	xxx_messageInfo_LabeledPolygon.DiscardUnknown(m)
}

var xxx_messageInfo_LabeledPolygon proto.InternalMessageInfo

func (m *LabeledPolygon) GetVertices() []*Vertex {
	if m != nil {
		return m.Vertices
	}
	return nil
}

func (m *LabeledPolygon) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type Vertex struct {
	X                    float64  `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vertex) Reset()         { *m = Vertex{} }
func (m *Vertex) String() string { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()    {}
func (*Vertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{4}
}

func (m *Vertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vertex.Unmarshal(m, b)
}
func (m *Vertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vertex.Marshal(b, m, deterministic)
}
func (m *Vertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vertex.Merge(m, src)
}
func (m *Vertex) XXX_Size() int {
	return xxx_messageInfo_Vertex.Size(m)
}
func (m *Vertex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vertex.DiscardUnknown(m)
}

var xxx_messageInfo_Vertex proto.InternalMessageInfo

func (m *Vertex) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vertex) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type LabelingRequest struct {
	Title                string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Instruction          string            `protobuf:"bytes,2,opt,name=instruction,proto3" json:"instruction,omitempty"`
	Image                *Image            `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Labels               []*Label          `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Hint                 []*LabeledPolygon `protobuf:"bytes,5,rep,name=hint,proto3" json:"hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LabelingRequest) Reset()         { *m = LabelingRequest{} }
func (m *LabelingRequest) String() string { return proto.CompactTextString(m) }
func (*LabelingRequest) ProtoMessage()    {}
func (*LabelingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{5}
}

func (m *LabelingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelingRequest.Unmarshal(m, b)
}
func (m *LabelingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelingRequest.Marshal(b, m, deterministic)
}
func (m *LabelingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelingRequest.Merge(m, src)
}
func (m *LabelingRequest) XXX_Size() int {
	return xxx_messageInfo_LabelingRequest.Size(m)
}
func (m *LabelingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LabelingRequest proto.InternalMessageInfo

func (m *LabelingRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *LabelingRequest) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

func (m *LabelingRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *LabelingRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LabelingRequest) GetHint() []*LabeledPolygon {
	if m != nil {
		return m.Hint
	}
	return nil
}

type LabelingResponse struct {
	Value                []*LabeledPolygon `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LabelingResponse) Reset()         { *m = LabelingResponse{} }
func (m *LabelingResponse) String() string { return proto.CompactTextString(m) }
func (*LabelingResponse) ProtoMessage()    {}
func (*LabelingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{6}
}

func (m *LabelingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelingResponse.Unmarshal(m, b)
}
func (m *LabelingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelingResponse.Marshal(b, m, deterministic)
}
func (m *LabelingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelingResponse.Merge(m, src)
}
func (m *LabelingResponse) XXX_Size() int {
	return xxx_messageInfo_LabelingResponse.Size(m)
}
func (m *LabelingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LabelingResponse proto.InternalMessageInfo

func (m *LabelingResponse) GetValue() []*LabeledPolygon {
	if m != nil {
		return m.Value
	}
	return nil
}

type SelectionRequest struct {
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*SelectionRequest_Image
	Data                 isSelectionRequest_Data `protobuf_oneof:"data"`
	Instruction          string                  `protobuf:"bytes,3,opt,name=instruction,proto3" json:"instruction,omitempty"`
	Options              []string                `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
	Hint                 uint32                  `protobuf:"varint,5,opt,name=hint,proto3" json:"hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SelectionRequest) Reset()         { *m = SelectionRequest{} }
func (m *SelectionRequest) String() string { return proto.CompactTextString(m) }
func (*SelectionRequest) ProtoMessage()    {}
func (*SelectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{7}
}

func (m *SelectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectionRequest.Unmarshal(m, b)
}
func (m *SelectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectionRequest.Marshal(b, m, deterministic)
}
func (m *SelectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectionRequest.Merge(m, src)
}
func (m *SelectionRequest) XXX_Size() int {
	return xxx_messageInfo_SelectionRequest.Size(m)
}
func (m *SelectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SelectionRequest proto.InternalMessageInfo

func (m *SelectionRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type isSelectionRequest_Data interface {
	isSelectionRequest_Data()
}

type SelectionRequest_Image struct {
	Image *Image `protobuf:"bytes,2,opt,name=image,proto3,oneof"`
}

func (*SelectionRequest_Image) isSelectionRequest_Data() {}

func (m *SelectionRequest) GetData() isSelectionRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SelectionRequest) GetImage() *Image {
	if x, ok := m.GetData().(*SelectionRequest_Image); ok {
		return x.Image
	}
	return nil
}

func (m *SelectionRequest) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

func (m *SelectionRequest) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *SelectionRequest) GetHint() uint32 {
	if m != nil {
		return m.Hint
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SelectionRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SelectionRequest_Image)(nil),
	}
}

type SelectionResponse struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectionResponse) Reset()         { *m = SelectionResponse{} }
func (m *SelectionResponse) String() string { return proto.CompactTextString(m) }
func (*SelectionResponse) ProtoMessage()    {}
func (*SelectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de983aa1dcb2256a, []int{8}
}

func (m *SelectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectionResponse.Unmarshal(m, b)
}
func (m *SelectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectionResponse.Marshal(b, m, deterministic)
}
func (m *SelectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectionResponse.Merge(m, src)
}
func (m *SelectionResponse) XXX_Size() int {
	return xxx_messageInfo_SelectionResponse.Size(m)
}
func (m *SelectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SelectionResponse proto.InternalMessageInfo

func (m *SelectionResponse) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("v1.model.Severity", Severity_name, Severity_value)
	proto.RegisterType((*InterventionRequest)(nil), "v1.model.InterventionRequest")
	proto.RegisterMapType((map[string]string)(nil), "v1.model.InterventionRequest.TagsEntry")
	proto.RegisterType((*InterventionResponse)(nil), "v1.model.InterventionResponse")
	proto.RegisterType((*Label)(nil), "v1.model.Label")
	proto.RegisterType((*LabeledPolygon)(nil), "v1.model.LabeledPolygon")
	proto.RegisterType((*Vertex)(nil), "v1.model.Vertex")
	proto.RegisterType((*LabelingRequest)(nil), "v1.model.LabelingRequest")
	proto.RegisterType((*LabelingResponse)(nil), "v1.model.LabelingResponse")
	proto.RegisterType((*SelectionRequest)(nil), "v1.model.SelectionRequest")
	proto.RegisterType((*SelectionResponse)(nil), "v1.model.SelectionResponse")
}

func init() { proto.RegisterFile("protos/model/v1/intervention.proto", fileDescriptor_de983aa1dcb2256a) }

var fileDescriptor_de983aa1dcb2256a = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5d, 0x4f, 0xd4, 0x40,
	0x14, 0xdd, 0x69, 0x77, 0x97, 0xed, 0x5d, 0x60, 0xcb, 0x48, 0x4c, 0x05, 0x35, 0x6b, 0xa3, 0x61,
	0x51, 0xd2, 0x0d, 0xeb, 0x83, 0x06, 0x7d, 0x10, 0x08, 0x48, 0x13, 0xb2, 0x98, 0x91, 0x60, 0xe2,
	0x0b, 0x29, 0x74, 0x52, 0x26, 0xf6, 0x63, 0xed, 0xcc, 0x6e, 0xd8, 0x77, 0xff, 0x8c, 0x7f, 0xc3,
	0x17, 0xff, 0x96, 0xe9, 0xf4, 0x93, 0xb2, 0xf8, 0xd6, 0xb9, 0x73, 0xe6, 0xcc, 0x3d, 0xe7, 0xdc,
	0x0e, 0x98, 0x93, 0x38, 0x12, 0x11, 0x1f, 0x06, 0x91, 0x4b, 0xfd, 0xe1, 0x6c, 0x77, 0xc8, 0x42,
	0x41, 0xe3, 0x19, 0x0d, 0x05, 0x8b, 0x42, 0x4b, 0x6e, 0xe2, 0xce, 0x6c, 0xd7, 0x92, 0xfb, 0x1b,
	0x9b, 0x75, 0x74, 0x40, 0x5d, 0xe6, 0xa4, 0x30, 0xf3, 0x8f, 0x0a, 0x8f, 0xec, 0xca, 0x69, 0x42,
	0x7f, 0x4e, 0x29, 0x17, 0x78, 0x15, 0x14, 0xe6, 0x1a, 0xa8, 0x8f, 0x06, 0x1a, 0x51, 0x98, 0x8b,
	0x9f, 0x82, 0x26, 0x58, 0x40, 0xb9, 0x70, 0x82, 0x89, 0xa1, 0xf4, 0xd1, 0x40, 0x25, 0x65, 0x01,
	0x5b, 0xd0, 0xe1, 0x74, 0x46, 0x63, 0x26, 0xe6, 0x86, 0xda, 0x47, 0x83, 0xd5, 0x11, 0xb6, 0xf2,
	0xfb, 0xad, 0xaf, 0xd9, 0x0e, 0x29, 0x30, 0xd8, 0x86, 0x35, 0x4e, 0x7d, 0x7a, 0x9d, 0xdc, 0x78,
	0x19, 0xa7, 0x57, 0x1a, 0xcd, 0x3e, 0x1a, 0x74, 0x47, 0x1b, 0xd5, 0x83, 0x19, 0x24, 0x6b, 0xea,
	0xa4, 0x41, 0x74, 0x5e, 0xab, 0xe1, 0x63, 0xd0, 0x7d, 0xe7, 0x8a, 0xfa, 0x2c, 0xf4, 0x0a, 0xa6,
	0x96, 0x64, 0x7a, 0x52, 0x32, 0x9d, 0x66, 0x88, 0x92, 0xa8, 0xe7, 0xdf, 0x2d, 0xe1, 0x0f, 0xd0,
	0x14, 0x8e, 0xc7, 0x8d, 0x76, 0x5f, 0x1d, 0x74, 0x47, 0x5b, 0xe5, 0xd9, 0x05, 0xee, 0x58, 0xe7,
	0x8e, 0xc7, 0x8f, 0x42, 0x11, 0xcf, 0x89, 0x3c, 0x84, 0x3f, 0x82, 0x16, 0x53, 0x3e, 0x89, 0x42,
	0x4e, 0xb9, 0xb1, 0x24, 0x19, 0x9e, 0x3f, 0xc4, 0x90, 0xc2, 0x48, 0x79, 0x60, 0xe3, 0x1d, 0x68,
	0x05, 0x21, 0xd6, 0x41, 0xfd, 0x41, 0xe7, 0x99, 0xf3, 0xc9, 0x27, 0x5e, 0x87, 0xd6, 0xcc, 0xf1,
	0xa7, 0x54, 0xda, 0xae, 0x91, 0x74, 0xb1, 0xa7, 0xbc, 0x47, 0x07, 0x6d, 0x68, 0xba, 0x8e, 0x70,
	0xcc, 0x5f, 0x0a, 0xac, 0x2f, 0xba, 0xe4, 0x5e, 0x8a, 0xcf, 0x00, 0x32, 0x8f, 0x2e, 0x99, 0x9b,
	0xf1, 0x69, 0x59, 0xc5, 0xae, 0x85, 0xac, 0xd6, 0x43, 0x3e, 0x05, 0x5c, 0x0d, 0x2d, 0xbd, 0x22,
	0x4b, 0x6d, 0x73, 0x61, 0x6a, 0x29, 0xe4, 0xa4, 0x41, 0xd6, 0x78, 0xbd, 0x98, 0x8c, 0x40, 0x25,
	0xb7, 0x8c, 0xac, 0x55, 0x1f, 0x81, 0x32, 0xb8, 0x82, 0x4b, 0xf7, 0x6b, 0xb5, 0xc2, 0x86, 0x4f,
	0xd0, 0x92, 0xf8, 0xd2, 0x31, 0x54, 0x71, 0x0c, 0xbf, 0x80, 0x65, 0x97, 0xf1, 0x89, 0xef, 0xcc,
	0x2f, 0x43, 0x27, 0xc8, 0xed, 0xec, 0x66, 0xb5, 0xb1, 0x13, 0x50, 0xf3, 0x02, 0x56, 0x25, 0x03,
	0x75, 0xbf, 0x44, 0xfe, 0xdc, 0x8b, 0x42, 0xbc, 0x03, 0x9d, 0x19, 0x8d, 0x05, 0xbb, 0xa6, 0xdc,
	0x40, 0x32, 0x58, 0xbd, 0xec, 0xee, 0x82, 0xc6, 0x82, 0xde, 0x92, 0x02, 0x81, 0x1f, 0x43, 0x5b,
	0x76, 0xc7, 0x0d, 0xa5, 0xaf, 0x0e, 0x34, 0x92, 0xad, 0xcc, 0x97, 0xd0, 0x4e, 0xb1, 0x78, 0x19,
	0xd0, 0xad, 0x6c, 0x0b, 0x11, 0x24, 0x57, 0x73, 0xd9, 0x07, 0x22, 0x68, 0x6e, 0xfe, 0x45, 0xd0,
	0xab, 0x4d, 0x6a, 0x22, 0x45, 0x30, 0xe1, 0x17, 0x52, 0xe4, 0x02, 0xf7, 0xa1, 0xcb, 0x42, 0x2e,
	0xe2, 0xa9, 0xf4, 0x34, 0x57, 0x52, 0x29, 0xe1, 0x57, 0xd0, 0x62, 0x81, 0xe3, 0x51, 0x19, 0x63,
	0x77, 0xd4, 0xab, 0x4c, 0x63, 0x52, 0x26, 0xe9, 0x2e, 0xde, 0x2a, 0x1a, 0x6e, 0x4a, 0x71, 0xbd,
	0x9a, 0xf5, 0xb9, 0x02, 0xbc, 0x03, 0xcd, 0x1b, 0x16, 0x26, 0xbf, 0x56, 0x02, 0x33, 0x6a, 0xb0,
	0xc2, 0x2f, 0x22, 0x51, 0xe6, 0x01, 0xe8, 0xf5, 0xe4, 0xb0, 0x55, 0x86, 0xf2, 0x7f, 0x8a, 0x14,
	0x66, 0xfe, 0x46, 0xa0, 0xd7, 0x5f, 0x80, 0x07, 0xec, 0xd8, 0xca, 0xc5, 0x2a, 0x0b, 0xc5, 0x9e,
	0x34, 0x72, 0xb9, 0x35, 0xdf, 0xd4, 0xfb, 0xbe, 0x19, 0xb0, 0x14, 0x4d, 0x92, 0xaf, 0xd4, 0x11,
	0x8d, 0xe4, 0x4b, 0x8c, 0x0b, 0x07, 0xd0, 0x60, 0x25, 0xd5, 0x59, 0x4c, 0xde, 0x36, 0xac, 0xdd,
	0x1b, 0xfb, 0xbb, 0x53, 0xb8, 0x92, 0xc9, 0x7a, 0xbd, 0x07, 0x9d, 0xfc, 0x41, 0xc4, 0x1d, 0x68,
	0xda, 0xe3, 0xe3, 0x33, 0xbd, 0x81, 0xbb, 0xb0, 0xf4, 0x6d, 0x9f, 0x8c, 0xed, 0xf1, 0x67, 0x1d,
	0x61, 0x0d, 0x5a, 0x47, 0x84, 0x9c, 0x11, 0x5d, 0xc1, 0xcb, 0xd0, 0x39, 0x24, 0xf6, 0xb9, 0x7d,
	0xb8, 0x7f, 0xaa, 0xab, 0x07, 0x6f, 0xbe, 0x6f, 0x7b, 0x4c, 0xdc, 0x4c, 0xaf, 0xac, 0xeb, 0x28,
	0x18, 0x1e, 0x47, 0x71, 0xe0, 0x84, 0xc2, 0x3e, 0x1b, 0x7a, 0x34, 0x94, 0xaf, 0xf9, 0xd0, 0x8b,
	0xe4, 0xfb, 0x9e, 0x48, 0xbf, 0x6a, 0xcb, 0xd2, 0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e,
	0xe5, 0x02, 0xf0, 0x2d, 0x06, 0x00, 0x00,
}
