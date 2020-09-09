// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduling.proto

package resScheduling

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SchedulingCGpuRequest struct {
	ProjectID            int64    `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	CgpuType             int64    `protobuf:"varint,2,opt,name=cgpuType,proto3" json:"cgpuType,omitempty"`
	NodesAfter           []int64  `protobuf:"varint,3,rep,packed,name=nodesAfter,proto3" json:"nodesAfter,omitempty"`
	CtrlID               int64    `protobuf:"varint,4,opt,name=ctrlID,proto3" json:"ctrlID,omitempty"`
	CtrlCN               string   `protobuf:"bytes,5,opt,name=ctrlCN,proto3" json:"ctrlCN,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulingCGpuRequest) Reset()         { *m = SchedulingCGpuRequest{} }
func (m *SchedulingCGpuRequest) String() string { return proto.CompactTextString(m) }
func (*SchedulingCGpuRequest) ProtoMessage()    {}
func (*SchedulingCGpuRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb732522b6cbe36, []int{0}
}

func (m *SchedulingCGpuRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulingCGpuRequest.Unmarshal(m, b)
}
func (m *SchedulingCGpuRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulingCGpuRequest.Marshal(b, m, deterministic)
}
func (m *SchedulingCGpuRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulingCGpuRequest.Merge(m, src)
}
func (m *SchedulingCGpuRequest) XXX_Size() int {
	return xxx_messageInfo_SchedulingCGpuRequest.Size(m)
}
func (m *SchedulingCGpuRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulingCGpuRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulingCGpuRequest proto.InternalMessageInfo

func (m *SchedulingCGpuRequest) GetProjectID() int64 {
	if m != nil {
		return m.ProjectID
	}
	return 0
}

func (m *SchedulingCGpuRequest) GetCgpuType() int64 {
	if m != nil {
		return m.CgpuType
	}
	return 0
}

func (m *SchedulingCGpuRequest) GetNodesAfter() []int64 {
	if m != nil {
		return m.NodesAfter
	}
	return nil
}

func (m *SchedulingCGpuRequest) GetCtrlID() int64 {
	if m != nil {
		return m.CtrlID
	}
	return 0
}

func (m *SchedulingCGpuRequest) GetCtrlCN() string {
	if m != nil {
		return m.CtrlCN
	}
	return ""
}

type SchedulingStorageRequest struct {
	ProjectID             int64    `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	StorageSizeAfter      int64    `protobuf:"varint,2,opt,name=storageSizeAfter,proto3" json:"storageSizeAfter,omitempty"`
	StorageAllocInfoAfter string   `protobuf:"bytes,3,opt,name=storageAllocInfoAfter,proto3" json:"storageAllocInfoAfter,omitempty"`
	CtrlID                int64    `protobuf:"varint,4,opt,name=ctrlID,proto3" json:"ctrlID,omitempty"`
	CtrlCN                string   `protobuf:"bytes,5,opt,name=ctrlCN,proto3" json:"ctrlCN,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SchedulingStorageRequest) Reset()         { *m = SchedulingStorageRequest{} }
func (m *SchedulingStorageRequest) String() string { return proto.CompactTextString(m) }
func (*SchedulingStorageRequest) ProtoMessage()    {}
func (*SchedulingStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb732522b6cbe36, []int{1}
}

func (m *SchedulingStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulingStorageRequest.Unmarshal(m, b)
}
func (m *SchedulingStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulingStorageRequest.Marshal(b, m, deterministic)
}
func (m *SchedulingStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulingStorageRequest.Merge(m, src)
}
func (m *SchedulingStorageRequest) XXX_Size() int {
	return xxx_messageInfo_SchedulingStorageRequest.Size(m)
}
func (m *SchedulingStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulingStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulingStorageRequest proto.InternalMessageInfo

func (m *SchedulingStorageRequest) GetProjectID() int64 {
	if m != nil {
		return m.ProjectID
	}
	return 0
}

func (m *SchedulingStorageRequest) GetStorageSizeAfter() int64 {
	if m != nil {
		return m.StorageSizeAfter
	}
	return 0
}

func (m *SchedulingStorageRequest) GetStorageAllocInfoAfter() string {
	if m != nil {
		return m.StorageAllocInfoAfter
	}
	return ""
}

func (m *SchedulingStorageRequest) GetCtrlID() int64 {
	if m != nil {
		return m.CtrlID
	}
	return 0
}

func (m *SchedulingStorageRequest) GetCtrlCN() string {
	if m != nil {
		return m.CtrlCN
	}
	return ""
}

type SchedulingReply struct {
	ErrorMessage         string   `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulingReply) Reset()         { *m = SchedulingReply{} }
func (m *SchedulingReply) String() string { return proto.CompactTextString(m) }
func (*SchedulingReply) ProtoMessage()    {}
func (*SchedulingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb732522b6cbe36, []int{2}
}

func (m *SchedulingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulingReply.Unmarshal(m, b)
}
func (m *SchedulingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulingReply.Marshal(b, m, deterministic)
}
func (m *SchedulingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulingReply.Merge(m, src)
}
func (m *SchedulingReply) XXX_Size() int {
	return xxx_messageInfo_SchedulingReply.Size(m)
}
func (m *SchedulingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulingReply.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulingReply proto.InternalMessageInfo

func (m *SchedulingReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

///////////////////////////////////////////////////////////
type QueryTreeRequest struct {
	ProjectID            int64    `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	CgpuType             int64    `protobuf:"varint,2,opt,name=cgpuType,proto3" json:"cgpuType,omitempty"`
	QueryType            int64    `protobuf:"varint,3,opt,name=queryType,proto3" json:"queryType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryTreeRequest) Reset()         { *m = QueryTreeRequest{} }
func (m *QueryTreeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTreeRequest) ProtoMessage()    {}
func (*QueryTreeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb732522b6cbe36, []int{3}
}

func (m *QueryTreeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTreeRequest.Unmarshal(m, b)
}
func (m *QueryTreeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTreeRequest.Marshal(b, m, deterministic)
}
func (m *QueryTreeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTreeRequest.Merge(m, src)
}
func (m *QueryTreeRequest) XXX_Size() int {
	return xxx_messageInfo_QueryTreeRequest.Size(m)
}
func (m *QueryTreeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTreeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTreeRequest proto.InternalMessageInfo

func (m *QueryTreeRequest) GetProjectID() int64 {
	if m != nil {
		return m.ProjectID
	}
	return 0
}

func (m *QueryTreeRequest) GetCgpuType() int64 {
	if m != nil {
		return m.CgpuType
	}
	return 0
}

func (m *QueryTreeRequest) GetQueryType() int64 {
	if m != nil {
		return m.QueryType
	}
	return 0
}

type QueryTreeReply struct {
	JsonTree             string   `protobuf:"bytes,1,opt,name=jsonTree,proto3" json:"jsonTree,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryTreeReply) Reset()         { *m = QueryTreeReply{} }
func (m *QueryTreeReply) String() string { return proto.CompactTextString(m) }
func (*QueryTreeReply) ProtoMessage()    {}
func (*QueryTreeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb732522b6cbe36, []int{4}
}

func (m *QueryTreeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTreeReply.Unmarshal(m, b)
}
func (m *QueryTreeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTreeReply.Marshal(b, m, deterministic)
}
func (m *QueryTreeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTreeReply.Merge(m, src)
}
func (m *QueryTreeReply) XXX_Size() int {
	return xxx_messageInfo_QueryTreeReply.Size(m)
}
func (m *QueryTreeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTreeReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTreeReply proto.InternalMessageInfo

func (m *QueryTreeReply) GetJsonTree() string {
	if m != nil {
		return m.JsonTree
	}
	return ""
}

func (m *QueryTreeReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

///////////////////////////////////////////////////////////
type QueryStorageRequest struct {
	ProjectID            int64    `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStorageRequest) Reset()         { *m = QueryStorageRequest{} }
func (m *QueryStorageRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStorageRequest) ProtoMessage()    {}
func (*QueryStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb732522b6cbe36, []int{5}
}

func (m *QueryStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStorageRequest.Unmarshal(m, b)
}
func (m *QueryStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStorageRequest.Marshal(b, m, deterministic)
}
func (m *QueryStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageRequest.Merge(m, src)
}
func (m *QueryStorageRequest) XXX_Size() int {
	return xxx_messageInfo_QueryStorageRequest.Size(m)
}
func (m *QueryStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageRequest proto.InternalMessageInfo

func (m *QueryStorageRequest) GetProjectID() int64 {
	if m != nil {
		return m.ProjectID
	}
	return 0
}

type QueryStorageReply struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	AllocInfo            string   `protobuf:"bytes,2,opt,name=allocInfo,proto3" json:"allocInfo,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStorageReply) Reset()         { *m = QueryStorageReply{} }
func (m *QueryStorageReply) String() string { return proto.CompactTextString(m) }
func (*QueryStorageReply) ProtoMessage()    {}
func (*QueryStorageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb732522b6cbe36, []int{6}
}

func (m *QueryStorageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStorageReply.Unmarshal(m, b)
}
func (m *QueryStorageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStorageReply.Marshal(b, m, deterministic)
}
func (m *QueryStorageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageReply.Merge(m, src)
}
func (m *QueryStorageReply) XXX_Size() int {
	return xxx_messageInfo_QueryStorageReply.Size(m)
}
func (m *QueryStorageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageReply proto.InternalMessageInfo

func (m *QueryStorageReply) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *QueryStorageReply) GetAllocInfo() string {
	if m != nil {
		return m.AllocInfo
	}
	return ""
}

func (m *QueryStorageReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*SchedulingCGpuRequest)(nil), "resScheduling.SchedulingCGpuRequest")
	proto.RegisterType((*SchedulingStorageRequest)(nil), "resScheduling.SchedulingStorageRequest")
	proto.RegisterType((*SchedulingReply)(nil), "resScheduling.SchedulingReply")
	proto.RegisterType((*QueryTreeRequest)(nil), "resScheduling.QueryTreeRequest")
	proto.RegisterType((*QueryTreeReply)(nil), "resScheduling.QueryTreeReply")
	proto.RegisterType((*QueryStorageRequest)(nil), "resScheduling.QueryStorageRequest")
	proto.RegisterType((*QueryStorageReply)(nil), "resScheduling.QueryStorageReply")
}

func init() { proto.RegisterFile("scheduling.proto", fileDescriptor_9cb732522b6cbe36) }

var fileDescriptor_9cb732522b6cbe36 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0xce, 0xd2, 0x40,
	0x10, 0xc7, 0x2d, 0x45, 0x42, 0x27, 0x80, 0xb0, 0x06, 0xd3, 0x34, 0x88, 0xcd, 0xc6, 0x03, 0xf1,
	0xc0, 0x41, 0xf4, 0x01, 0x08, 0x24, 0x86, 0x83, 0x46, 0x0a, 0x31, 0x26, 0x9e, 0xb0, 0x0c, 0xb5,
	0xa4, 0xe9, 0x96, 0xdd, 0xf6, 0x00, 0x6f, 0xe3, 0xeb, 0xf8, 0x30, 0x3e, 0x83, 0xe9, 0xb6, 0xb4,
	0x94, 0x16, 0xf3, 0x91, 0xef, 0xbb, 0x31, 0xff, 0x9d, 0x99, 0xfd, 0x0d, 0xff, 0x9d, 0x42, 0x57,
	0xd8, 0xbf, 0x70, 0x1b, 0x79, 0xae, 0xef, 0x8c, 0x03, 0xce, 0x42, 0x46, 0xda, 0x1c, 0xc5, 0x2a,
	0x13, 0xe9, 0x6f, 0x05, 0xfa, 0x79, 0x38, 0xfb, 0x14, 0x44, 0x16, 0x1e, 0x22, 0x14, 0x21, 0x19,
	0x80, 0x16, 0x70, 0xb6, 0x47, 0x3b, 0x5c, 0xcc, 0x75, 0xc5, 0x54, 0x46, 0xaa, 0x95, 0x0b, 0xc4,
	0x80, 0xa6, 0xed, 0x04, 0xd1, 0xfa, 0x18, 0xa0, 0x5e, 0x93, 0x87, 0x59, 0x4c, 0x86, 0x00, 0x3e,
	0xdb, 0xa2, 0x98, 0xee, 0x42, 0xe4, 0xba, 0x6a, 0xaa, 0x23, 0xd5, 0xba, 0x50, 0xc8, 0x2b, 0x68,
	0xd8, 0x21, 0xf7, 0x16, 0x73, 0xbd, 0x2e, 0x2b, 0xd3, 0xe8, 0xac, 0xcf, 0xbe, 0xe8, 0xcf, 0x4d,
	0x65, 0xa4, 0x59, 0x69, 0x44, 0xff, 0x28, 0xa0, 0xe7, 0x8c, 0xab, 0x90, 0xf1, 0x8d, 0x83, 0x0f,
	0xc3, 0x7c, 0x07, 0x5d, 0x91, 0xe4, 0xaf, 0xdc, 0x13, 0x26, 0x40, 0x09, 0x6e, 0x49, 0x27, 0x1f,
	0xa0, 0x9f, 0x6a, 0x53, 0xcf, 0x63, 0xf6, 0xc2, 0xdf, 0xb1, 0xf3, 0x04, 0x31, 0x4d, 0xf5, 0xe1,
	0xdd, 0xc3, 0x7c, 0x84, 0x17, 0xf9, 0x2c, 0x16, 0x06, 0xde, 0x91, 0x50, 0x68, 0x21, 0xe7, 0x8c,
	0x7f, 0x46, 0x21, 0x36, 0x0e, 0xca, 0x29, 0x34, 0xab, 0xa0, 0xd1, 0x3d, 0x74, 0x97, 0x11, 0xf2,
	0xe3, 0x9a, 0x23, 0x3e, 0xde, 0xa1, 0x01, 0x68, 0x07, 0xd9, 0x2d, 0x3e, 0x54, 0x93, 0xca, 0x4c,
	0xa0, 0x5f, 0xa1, 0x73, 0x71, 0x57, 0x4c, 0x68, 0x40, 0x73, 0x2f, 0x98, 0x1f, 0x0b, 0x29, 0x5d,
	0x16, 0x97, 0xe8, 0x6b, 0x15, 0xf4, 0x13, 0x78, 0x29, 0x3b, 0xde, 0xe3, 0x1d, 0x75, 0xa1, 0x57,
	0x2c, 0x8a, 0x49, 0x08, 0xd4, 0x85, 0x7b, 0xc2, 0x34, 0x5b, 0xfe, 0x8e, 0xdb, 0x6c, 0xce, 0xa6,
	0xa4, 0xd7, 0xe7, 0x42, 0x89, 0x4f, 0x2d, 0xf3, 0xbd, 0xff, 0x5b, 0x03, 0xc8, 0x5d, 0x21, 0xdf,
	0xa1, 0x53, 0xdc, 0x09, 0xf2, 0x76, 0x5c, 0x58, 0x9b, 0x71, 0xe5, 0xca, 0x18, 0xc3, 0x9b, 0x59,
	0x12, 0x9e, 0x3e, 0x23, 0x3f, 0xa0, 0x57, 0x7a, 0xc9, 0x4f, 0xd6, 0x7c, 0x09, 0x6d, 0xf9, 0x87,
	0xc5, 0x55, 0xd2, 0x9a, 0x37, 0x57, 0x25, 0xd7, 0x2f, 0xc8, 0x78, 0x7d, 0x3b, 0x21, 0x69, 0xf9,
	0x0d, 0x5a, 0x97, 0x1e, 0x10, 0x5a, 0x55, 0x50, 0x74, 0xd5, 0x30, 0xff, 0x9b, 0x23, 0xfb, 0xfe,
	0x6c, 0xc8, 0x8f, 0xd1, 0xe4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0xa4, 0x69, 0x57, 0xa0,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchedulingClient is the client API for Scheduling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulingClient interface {
	// Scheduling Resource
	SchedulingCGpu(ctx context.Context, in *SchedulingCGpuRequest, opts ...grpc.CallOption) (*SchedulingReply, error)
	SchedulingStorage(ctx context.Context, in *SchedulingCGpuRequest, opts ...grpc.CallOption) (*SchedulingReply, error)
	// Query Resource Tree
	QueryCGpuTree(ctx context.Context, in *QueryTreeRequest, opts ...grpc.CallOption) (*QueryTreeReply, error)
	QueryStorage(ctx context.Context, in *QueryStorageRequest, opts ...grpc.CallOption) (*QueryStorageReply, error)
}

type schedulingClient struct {
	cc *grpc.ClientConn
}

func NewSchedulingClient(cc *grpc.ClientConn) SchedulingClient {
	return &schedulingClient{cc}
}

func (c *schedulingClient) SchedulingCGpu(ctx context.Context, in *SchedulingCGpuRequest, opts ...grpc.CallOption) (*SchedulingReply, error) {
	out := new(SchedulingReply)
	err := c.cc.Invoke(ctx, "/resScheduling.Scheduling/SchedulingCGpu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulingClient) SchedulingStorage(ctx context.Context, in *SchedulingCGpuRequest, opts ...grpc.CallOption) (*SchedulingReply, error) {
	out := new(SchedulingReply)
	err := c.cc.Invoke(ctx, "/resScheduling.Scheduling/SchedulingStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulingClient) QueryCGpuTree(ctx context.Context, in *QueryTreeRequest, opts ...grpc.CallOption) (*QueryTreeReply, error) {
	out := new(QueryTreeReply)
	err := c.cc.Invoke(ctx, "/resScheduling.Scheduling/QueryCGpuTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulingClient) QueryStorage(ctx context.Context, in *QueryStorageRequest, opts ...grpc.CallOption) (*QueryStorageReply, error) {
	out := new(QueryStorageReply)
	err := c.cc.Invoke(ctx, "/resScheduling.Scheduling/QueryStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulingServer is the server API for Scheduling service.
type SchedulingServer interface {
	// Scheduling Resource
	SchedulingCGpu(context.Context, *SchedulingCGpuRequest) (*SchedulingReply, error)
	SchedulingStorage(context.Context, *SchedulingCGpuRequest) (*SchedulingReply, error)
	// Query Resource Tree
	QueryCGpuTree(context.Context, *QueryTreeRequest) (*QueryTreeReply, error)
	QueryStorage(context.Context, *QueryStorageRequest) (*QueryStorageReply, error)
}

// UnimplementedSchedulingServer can be embedded to have forward compatible implementations.
type UnimplementedSchedulingServer struct {
}

func (*UnimplementedSchedulingServer) SchedulingCGpu(ctx context.Context, req *SchedulingCGpuRequest) (*SchedulingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingCGpu not implemented")
}
func (*UnimplementedSchedulingServer) SchedulingStorage(ctx context.Context, req *SchedulingCGpuRequest) (*SchedulingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingStorage not implemented")
}
func (*UnimplementedSchedulingServer) QueryCGpuTree(ctx context.Context, req *QueryTreeRequest) (*QueryTreeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCGpuTree not implemented")
}
func (*UnimplementedSchedulingServer) QueryStorage(ctx context.Context, req *QueryStorageRequest) (*QueryStorageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStorage not implemented")
}

func RegisterSchedulingServer(s *grpc.Server, srv SchedulingServer) {
	s.RegisterService(&_Scheduling_serviceDesc, srv)
}

func _Scheduling_SchedulingCGpu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingCGpuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulingServer).SchedulingCGpu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resScheduling.Scheduling/SchedulingCGpu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulingServer).SchedulingCGpu(ctx, req.(*SchedulingCGpuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduling_SchedulingStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingCGpuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulingServer).SchedulingStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resScheduling.Scheduling/SchedulingStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulingServer).SchedulingStorage(ctx, req.(*SchedulingCGpuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduling_QueryCGpuTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulingServer).QueryCGpuTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resScheduling.Scheduling/QueryCGpuTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulingServer).QueryCGpuTree(ctx, req.(*QueryTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduling_QueryStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulingServer).QueryStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resScheduling.Scheduling/QueryStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulingServer).QueryStorage(ctx, req.(*QueryStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduling_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resScheduling.Scheduling",
	HandlerType: (*SchedulingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SchedulingCGpu",
			Handler:    _Scheduling_SchedulingCGpu_Handler,
		},
		{
			MethodName: "SchedulingStorage",
			Handler:    _Scheduling_SchedulingStorage_Handler,
		},
		{
			MethodName: "QueryCGpuTree",
			Handler:    _Scheduling_QueryCGpuTree_Handler,
		},
		{
			MethodName: "QueryStorage",
			Handler:    _Scheduling_QueryStorage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduling.proto",
}