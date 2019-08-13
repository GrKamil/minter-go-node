// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manager/manager.proto

package manager

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

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80040e3adc73a4a5, []int{0}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusReply struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusReply) Reset()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()    {}
func (*StatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_80040e3adc73a4a5, []int{1}
}

func (m *StatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReply.Unmarshal(m, b)
}
func (m *StatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReply.Marshal(b, m, deterministic)
}
func (m *StatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReply.Merge(m, src)
}
func (m *StatusReply) XXX_Size() int {
	return xxx_messageInfo_StatusReply.Size(m)
}
func (m *StatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReply proto.InternalMessageInfo

func (m *StatusReply) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type PruneStatesRequest struct {
	FromHeight           int64    `protobuf:"varint,1,opt,name=from_height,json=fromHeight,proto3" json:"from_height,omitempty"`
	ToHeight             int64    `protobuf:"varint,2,opt,name=to_height,json=toHeight,proto3" json:"to_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PruneStatesRequest) Reset()         { *m = PruneStatesRequest{} }
func (m *PruneStatesRequest) String() string { return proto.CompactTextString(m) }
func (*PruneStatesRequest) ProtoMessage()    {}
func (*PruneStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80040e3adc73a4a5, []int{2}
}

func (m *PruneStatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PruneStatesRequest.Unmarshal(m, b)
}
func (m *PruneStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PruneStatesRequest.Marshal(b, m, deterministic)
}
func (m *PruneStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PruneStatesRequest.Merge(m, src)
}
func (m *PruneStatesRequest) XXX_Size() int {
	return xxx_messageInfo_PruneStatesRequest.Size(m)
}
func (m *PruneStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PruneStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PruneStatesRequest proto.InternalMessageInfo

func (m *PruneStatesRequest) GetFromHeight() int64 {
	if m != nil {
		return m.FromHeight
	}
	return 0
}

func (m *PruneStatesRequest) GetToHeight() int64 {
	if m != nil {
		return m.ToHeight
	}
	return 0
}

type PruneStatesReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PruneStatesReply) Reset()         { *m = PruneStatesReply{} }
func (m *PruneStatesReply) String() string { return proto.CompactTextString(m) }
func (*PruneStatesReply) ProtoMessage()    {}
func (*PruneStatesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_80040e3adc73a4a5, []int{3}
}

func (m *PruneStatesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PruneStatesReply.Unmarshal(m, b)
}
func (m *PruneStatesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PruneStatesReply.Marshal(b, m, deterministic)
}
func (m *PruneStatesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PruneStatesReply.Merge(m, src)
}
func (m *PruneStatesReply) XXX_Size() int {
	return xxx_messageInfo_PruneStatesReply.Size(m)
}
func (m *PruneStatesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PruneStatesReply.DiscardUnknown(m)
}

var xxx_messageInfo_PruneStatesReply proto.InternalMessageInfo

func (m *PruneStatesReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ExportStatesRequest struct {
	FromHeight           int64    `protobuf:"varint,1,opt,name=from_height,json=fromHeight,proto3" json:"from_height,omitempty"`
	ToHeight             int64    `protobuf:"varint,2,opt,name=to_height,json=toHeight,proto3" json:"to_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportStatesRequest) Reset()         { *m = ExportStatesRequest{} }
func (m *ExportStatesRequest) String() string { return proto.CompactTextString(m) }
func (*ExportStatesRequest) ProtoMessage()    {}
func (*ExportStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80040e3adc73a4a5, []int{4}
}

func (m *ExportStatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportStatesRequest.Unmarshal(m, b)
}
func (m *ExportStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportStatesRequest.Marshal(b, m, deterministic)
}
func (m *ExportStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportStatesRequest.Merge(m, src)
}
func (m *ExportStatesRequest) XXX_Size() int {
	return xxx_messageInfo_ExportStatesRequest.Size(m)
}
func (m *ExportStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportStatesRequest proto.InternalMessageInfo

func (m *ExportStatesRequest) GetFromHeight() int64 {
	if m != nil {
		return m.FromHeight
	}
	return 0
}

func (m *ExportStatesRequest) GetToHeight() int64 {
	if m != nil {
		return m.ToHeight
	}
	return 0
}

type ExportStatesReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportStatesReply) Reset()         { *m = ExportStatesReply{} }
func (m *ExportStatesReply) String() string { return proto.CompactTextString(m) }
func (*ExportStatesReply) ProtoMessage()    {}
func (*ExportStatesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_80040e3adc73a4a5, []int{5}
}

func (m *ExportStatesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportStatesReply.Unmarshal(m, b)
}
func (m *ExportStatesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportStatesReply.Marshal(b, m, deterministic)
}
func (m *ExportStatesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportStatesReply.Merge(m, src)
}
func (m *ExportStatesReply) XXX_Size() int {
	return xxx_messageInfo_ExportStatesReply.Size(m)
}
func (m *ExportStatesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportStatesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ExportStatesReply proto.InternalMessageInfo

func (m *ExportStatesReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "manager.StatusRequest")
	proto.RegisterType((*StatusReply)(nil), "manager.StatusReply")
	proto.RegisterType((*PruneStatesRequest)(nil), "manager.PruneStatesRequest")
	proto.RegisterType((*PruneStatesReply)(nil), "manager.PruneStatesReply")
	proto.RegisterType((*ExportStatesRequest)(nil), "manager.ExportStatesRequest")
	proto.RegisterType((*ExportStatesReply)(nil), "manager.ExportStatesReply")
}

func init() { proto.RegisterFile("manager/manager.proto", fileDescriptor_80040e3adc73a4a5) }

var fileDescriptor_80040e3adc73a4a5 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0x2d, 0xd2, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x12, 0x3f, 0x17, 0x6f, 0x70, 0x49, 0x62, 0x49, 0x69, 0x71, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x92, 0x2a, 0x17, 0x37, 0x4c, 0xa0, 0x20, 0xa7, 0x52, 0x48, 0x8c, 0x8b, 0x2d, 0x23, 0x35,
	0x33, 0x3d, 0xa3, 0x44, 0x82, 0x51, 0x81, 0x51, 0x83, 0x39, 0x08, 0xca, 0x53, 0x0a, 0xe2, 0x12,
	0x0a, 0x28, 0x2a, 0xcd, 0x4b, 0x05, 0xa9, 0x4d, 0x85, 0x69, 0x16, 0x92, 0xe7, 0xe2, 0x4e, 0x2b,
	0xca, 0xcf, 0x8d, 0x47, 0xd1, 0xc2, 0x05, 0x12, 0xf2, 0x00, 0x8b, 0x08, 0x49, 0x73, 0x71, 0x96,
	0xe4, 0xc3, 0xa4, 0x99, 0xc0, 0xd2, 0x1c, 0x25, 0xf9, 0x10, 0x49, 0x25, 0x1d, 0x2e, 0x01, 0x14,
	0x33, 0x41, 0xf6, 0x4b, 0x70, 0xb1, 0x17, 0x97, 0x26, 0x27, 0xa7, 0x16, 0x17, 0x83, 0x4d, 0xe3,
	0x08, 0x82, 0x71, 0x95, 0x82, 0xb9, 0x84, 0x5d, 0x2b, 0x0a, 0xf2, 0x8b, 0x4a, 0xa8, 0xe9, 0x04,
	0x5d, 0x2e, 0x41, 0x54, 0x43, 0xf1, 0xba, 0xc1, 0xe8, 0x36, 0x23, 0x17, 0xbb, 0x2f, 0x24, 0x24,
	0x85, 0xac, 0xb9, 0x38, 0xdd, 0x53, 0x4b, 0x20, 0x61, 0x27, 0x24, 0xa6, 0x07, 0x0b, 0x6f, 0x94,
	0xd0, 0x95, 0x12, 0xc1, 0x10, 0x2f, 0xc8, 0xa9, 0x54, 0x62, 0x10, 0x72, 0xe7, 0xe2, 0x46, 0xf2,
	0xba, 0x90, 0x34, 0x5c, 0x19, 0x66, 0x20, 0x4b, 0x49, 0x62, 0x97, 0x84, 0x18, 0xe4, 0xc5, 0xc5,
	0x83, 0xec, 0x01, 0x21, 0x19, 0xb8, 0x62, 0x2c, 0x81, 0x25, 0x25, 0x85, 0x43, 0x16, 0x6c, 0x56,
	0x12, 0x1b, 0x38, 0xad, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0x8d, 0x48, 0x7c, 0x44,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagerClient interface {
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	PruneStates(ctx context.Context, in *PruneStatesRequest, opts ...grpc.CallOption) (*PruneStatesReply, error)
	ExportStates(ctx context.Context, in *ExportStatesRequest, opts ...grpc.CallOption) (*ExportStatesReply, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/manager.Manager/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) PruneStates(ctx context.Context, in *PruneStatesRequest, opts ...grpc.CallOption) (*PruneStatesReply, error) {
	out := new(PruneStatesReply)
	err := c.cc.Invoke(ctx, "/manager.Manager/PruneStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExportStates(ctx context.Context, in *ExportStatesRequest, opts ...grpc.CallOption) (*ExportStatesReply, error) {
	out := new(ExportStatesReply)
	err := c.cc.Invoke(ctx, "/manager.Manager/ExportStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
type ManagerServer interface {
	GetStatus(context.Context, *StatusRequest) (*StatusReply, error)
	PruneStates(context.Context, *PruneStatesRequest) (*PruneStatesReply, error)
	ExportStates(context.Context, *ExportStatesRequest) (*ExportStatesReply, error)
}

// UnimplementedManagerServer can be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (*UnimplementedManagerServer) GetStatus(ctx context.Context, req *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (*UnimplementedManagerServer) PruneStates(ctx context.Context, req *PruneStatesRequest) (*PruneStatesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PruneStates not implemented")
}
func (*UnimplementedManagerServer) ExportStates(ctx context.Context, req *ExportStatesRequest) (*ExportStatesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportStates not implemented")
}

func RegisterManagerServer(s *grpc.Server, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_PruneStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PruneStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).PruneStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/PruneStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).PruneStates(ctx, req.(*PruneStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExportStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExportStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/ExportStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExportStates(ctx, req.(*ExportStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manager.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Manager_GetStatus_Handler,
		},
		{
			MethodName: "PruneStates",
			Handler:    _Manager_PruneStates_Handler,
		},
		{
			MethodName: "ExportStates",
			Handler:    _Manager_ExportStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager/manager.proto",
}