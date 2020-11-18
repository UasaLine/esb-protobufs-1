// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/internal/errors.proto

package errors

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

type Error struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e6b5dca6754bf1c, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Error) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e6b5dca6754bf1c, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*Error)(nil), "products.Error")
	proto.RegisterType((*CreateResponse)(nil), "products.CreateResponse")
}

func init() { proto.RegisterFile("proto/internal/errors.proto", fileDescriptor_4e6b5dca6754bf1c) }

var fileDescriptor_4e6b5dca6754bf1c = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x4f, 0x84, 0x50,
	0x0c, 0xc6, 0x03, 0xea, 0x13, 0x6a, 0x82, 0x49, 0x07, 0xf3, 0xa2, 0x0b, 0x61, 0x72, 0x82, 0x44,
	0x66, 0x17, 0x8d, 0xa3, 0x0b, 0xa3, 0x1b, 0x62, 0x43, 0x0c, 0x77, 0xf4, 0xa5, 0x7d, 0x0c, 0xf7,
	0xdf, 0x5f, 0xae, 0x07, 0xb9, 0xdc, 0xd6, 0xef, 0xf7, 0x35, 0x5f, 0xfb, 0xc1, 0x4b, 0x10, 0x8e,
	0xdc, 0xfc, 0xcf, 0x91, 0x64, 0xee, 0x77, 0x0d, 0x89, 0xb0, 0x68, 0x6d, 0x14, 0xb3, 0x20, 0xfc,
	0xb7, 0x0c, 0x51, 0xab, 0x6f, 0xb8, 0xfb, 0x3a, 0x39, 0xf8, 0x04, 0x4e, 0x79, 0x91, 0x81, 0x7c,
	0x52, 0x26, 0xaf, 0x79, 0xb7, 0x2a, 0x44, 0xb8, 0x8d, 0x87, 0x40, 0x3e, 0x35, 0x6a, 0x33, 0x7a,
	0xb8, 0xdf, 0x93, 0x6a, 0x3f, 0x92, 0xbf, 0x31, 0xbc, 0xc9, 0xaa, 0x84, 0xe2, 0x53, 0xa8, 0x8f,
	0xd4, 0x91, 0x06, 0x9e, 0x95, 0xb0, 0x80, 0x94, 0x27, 0xcb, 0xcc, 0xba, 0x94, 0xa7, 0xb7, 0x77,
	0x70, 0x76, 0x50, 0xb1, 0x05, 0x77, 0xde, 0xc5, 0xc7, 0x7a, 0xfb, 0xa7, 0x36, 0xef, 0xd9, 0x5f,
	0xc0, 0x75, 0xdc, 0xc7, 0xc3, 0x4f, 0x3e, 0xf2, 0x5a, 0xe6, 0xd7, 0x59, 0x9b, 0xf6, 0x18, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0x67, 0xc1, 0xec, 0xec, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ErrorsClient is the client API for Errors service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ErrorsClient interface {
	Create(ctx context.Context, in *Error, opts ...grpc.CallOption) (*CreateResponse, error)
}

type errorsClient struct {
	cc *grpc.ClientConn
}

func NewErrorsClient(cc *grpc.ClientConn) ErrorsClient {
	return &errorsClient{cc}
}

func (c *errorsClient) Create(ctx context.Context, in *Error, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/products.Errors/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErrorsServer is the server API for Errors service.
type ErrorsServer interface {
	Create(context.Context, *Error) (*CreateResponse, error)
}

// UnimplementedErrorsServer can be embedded to have forward compatible implementations.
type UnimplementedErrorsServer struct {
}

func (*UnimplementedErrorsServer) Create(ctx context.Context, req *Error) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterErrorsServer(s *grpc.Server, srv ErrorsServer) {
	s.RegisterService(&_Errors_serviceDesc, srv)
}

func _Errors_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Error)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products.Errors/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorsServer).Create(ctx, req.(*Error))
	}
	return interceptor(ctx, in, info, handler)
}

var _Errors_serviceDesc = grpc.ServiceDesc{
	ServiceName: "products.Errors",
	HandlerType: (*ErrorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Errors_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/internal/errors.proto",
}