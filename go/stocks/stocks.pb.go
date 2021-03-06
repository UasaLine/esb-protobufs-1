// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/internal/stocks.proto

package stocks

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

type GetParams struct {
	EnabledBufferSite    bool     `protobuf:"varint,1,opt,name=enabled_buffer_site,json=enabledBufferSite,proto3" json:"enabled_buffer_site,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParams) Reset()         { *m = GetParams{} }
func (m *GetParams) String() string { return proto.CompactTextString(m) }
func (*GetParams) ProtoMessage()    {}
func (*GetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_56473f84fafb4361, []int{0}
}

func (m *GetParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParams.Unmarshal(m, b)
}
func (m *GetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParams.Marshal(b, m, deterministic)
}
func (m *GetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParams.Merge(m, src)
}
func (m *GetParams) XXX_Size() int {
	return xxx_messageInfo_GetParams.Size(m)
}
func (m *GetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetParams proto.InternalMessageInfo

func (m *GetParams) GetEnabledBufferSite() bool {
	if m != nil {
		return m.EnabledBufferSite
	}
	return false
}

// limit default = 1000
// offset default = 0;
type Request struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	EnabledBufferSite    bool     `protobuf:"varint,3,opt,name=enabled_buffer_site,json=enabledBufferSite,proto3" json:"enabled_buffer_site,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_56473f84fafb4361, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Request) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Request) GetEnabledBufferSite() bool {
	if m != nil {
		return m.EnabledBufferSite
	}
	return false
}

type Response struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Stocks               []*Stock `protobuf:"bytes,2,rep,name=stocks,proto3" json:"stocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_56473f84fafb4361, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetStocks() []*Stock {
	if m != nil {
		return m.Stocks
	}
	return nil
}

type SuccessResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessResponse) Reset()         { *m = SuccessResponse{} }
func (m *SuccessResponse) String() string { return proto.CompactTextString(m) }
func (*SuccessResponse) ProtoMessage()    {}
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56473f84fafb4361, []int{3}
}

func (m *SuccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessResponse.Unmarshal(m, b)
}
func (m *SuccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessResponse.Marshal(b, m, deterministic)
}
func (m *SuccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessResponse.Merge(m, src)
}
func (m *SuccessResponse) XXX_Size() int {
	return xxx_messageInfo_SuccessResponse.Size(m)
}
func (m *SuccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessResponse proto.InternalMessageInfo

func (m *SuccessResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListStocks struct {
	Stocks               []*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListStocks) Reset()         { *m = ListStocks{} }
func (m *ListStocks) String() string { return proto.CompactTextString(m) }
func (*ListStocks) ProtoMessage()    {}
func (*ListStocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_56473f84fafb4361, []int{4}
}

func (m *ListStocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStocks.Unmarshal(m, b)
}
func (m *ListStocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStocks.Marshal(b, m, deterministic)
}
func (m *ListStocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStocks.Merge(m, src)
}
func (m *ListStocks) XXX_Size() int {
	return xxx_messageInfo_ListStocks.Size(m)
}
func (m *ListStocks) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStocks.DiscardUnknown(m)
}

var xxx_messageInfo_ListStocks proto.InternalMessageInfo

func (m *ListStocks) GetStocks() []*Stock {
	if m != nil {
		return m.Stocks
	}
	return nil
}

type Stock struct {
	Barcode              int64    `protobuf:"varint,1,opt,name=barcode,proto3" json:"barcode,omitempty"`
	StoreCode            int32    `protobuf:"varint,2,opt,name=store_code,json=storeCode,proto3" json:"store_code,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Available            int32    `protobuf:"varint,4,opt,name=available,proto3" json:"available,omitempty"`
	Reserved             int32    `protobuf:"varint,5,opt,name=reserved,proto3" json:"reserved,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stock) Reset()         { *m = Stock{} }
func (m *Stock) String() string { return proto.CompactTextString(m) }
func (*Stock) ProtoMessage()    {}
func (*Stock) Descriptor() ([]byte, []int) {
	return fileDescriptor_56473f84fafb4361, []int{5}
}

func (m *Stock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stock.Unmarshal(m, b)
}
func (m *Stock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stock.Marshal(b, m, deterministic)
}
func (m *Stock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stock.Merge(m, src)
}
func (m *Stock) XXX_Size() int {
	return xxx_messageInfo_Stock.Size(m)
}
func (m *Stock) XXX_DiscardUnknown() {
	xxx_messageInfo_Stock.DiscardUnknown(m)
}

var xxx_messageInfo_Stock proto.InternalMessageInfo

func (m *Stock) GetBarcode() int64 {
	if m != nil {
		return m.Barcode
	}
	return 0
}

func (m *Stock) GetStoreCode() int32 {
	if m != nil {
		return m.StoreCode
	}
	return 0
}

func (m *Stock) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Stock) GetAvailable() int32 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *Stock) GetReserved() int32 {
	if m != nil {
		return m.Reserved
	}
	return 0
}

func init() {
	proto.RegisterType((*GetParams)(nil), "stocks.GetParams")
	proto.RegisterType((*Request)(nil), "stocks.Request")
	proto.RegisterType((*Response)(nil), "stocks.Response")
	proto.RegisterType((*SuccessResponse)(nil), "stocks.SuccessResponse")
	proto.RegisterType((*ListStocks)(nil), "stocks.ListStocks")
	proto.RegisterType((*Stock)(nil), "stocks.Stock")
}

func init() { proto.RegisterFile("proto/internal/stocks.proto", fileDescriptor_56473f84fafb4361) }

var fileDescriptor_56473f84fafb4361 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0xdd, 0x6c, 0x4c, 0x76, 0x52, 0xcb, 0xb2, 0x6e, 0x2b, 0x1a, 0xa2, 0xc2, 0x12, 0x10, 0x16,
	0x84, 0x2c, 0xec, 0x0a, 0x1e, 0xbc, 0x8d, 0x87, 0xb9, 0x78, 0x90, 0x0c, 0x5e, 0xbc, 0x0c, 0x9d,
	0x4c, 0x65, 0x68, 0xcc, 0xa4, 0x67, 0xba, 0x2a, 0x03, 0xfe, 0x88, 0xbf, 0xe0, 0x6f, 0x4a, 0x3a,
	0xdd, 0x19, 0x05, 0xc7, 0x53, 0xf3, 0x5e, 0x55, 0xd7, 0xe3, 0x3d, 0x1e, 0xbc, 0xda, 0x19, 0xcd,
	0xfa, 0x5e, 0x75, 0x8c, 0xa6, 0x93, 0xed, 0x3d, 0xb1, 0xae, 0xbf, 0x53, 0x61, 0x59, 0x11, 0x8f,
	0x28, 0xff, 0x08, 0xc9, 0x02, 0xf9, 0x8b, 0x34, 0x72, 0x4b, 0xa2, 0x80, 0x67, 0xd8, 0xc9, 0xaa,
	0xc5, 0xf5, 0xaa, 0xea, 0x9b, 0x06, 0xcd, 0x8a, 0x14, 0x63, 0x1a, 0xdc, 0x06, 0x77, 0xb3, 0xf2,
	0xc6, 0x8d, 0xe6, 0x76, 0xb2, 0x54, 0x8c, 0xf9, 0x06, 0x2e, 0x4a, 0xdc, 0xf7, 0x48, 0x2c, 0x9e,
	0x43, 0xd4, 0xaa, 0xad, 0x62, 0xbb, 0x1c, 0x95, 0x23, 0x10, 0x2f, 0x20, 0xd6, 0x4d, 0x43, 0xc8,
	0xe9, 0xb9, 0xa5, 0x1d, 0x3a, 0x25, 0x14, 0x9e, 0x12, 0x5a, 0xc0, 0xac, 0x44, 0xda, 0xe9, 0x8e,
	0x70, 0x50, 0x62, 0xcd, 0xb2, 0xf5, 0x4a, 0x16, 0x88, 0xb7, 0xe0, 0x1c, 0xa5, 0xe7, 0xb7, 0xe1,
	0xdd, 0xe5, 0xc3, 0x55, 0xe1, 0xec, 0x2e, 0x87, 0xa7, 0xf4, 0x76, 0xdf, 0xc1, 0xf5, 0xb2, 0xaf,
	0x6b, 0x24, 0x9a, 0xee, 0xa5, 0x70, 0x41, 0x23, 0xe5, 0x8c, 0x7a, 0x98, 0x3f, 0x02, 0x7c, 0x56,
	0xc4, 0xf6, 0x02, 0xfd, 0xa1, 0x10, 0xfc, 0x4f, 0xe1, 0x67, 0x00, 0x91, 0x65, 0x86, 0xc3, 0x95,
	0x34, 0xb5, 0x5e, 0x8f, 0x09, 0x86, 0xa5, 0x87, 0xe2, 0x0d, 0x00, 0xb1, 0x36, 0xb8, 0xb2, 0xc3,
	0x31, 0x9a, 0xc4, 0x32, 0x9f, 0x86, 0x71, 0x06, 0xb3, 0x7d, 0x2f, 0x3b, 0x56, 0xfc, 0xc3, 0x46,
	0x12, 0x95, 0x13, 0x16, 0xaf, 0x21, 0x91, 0x07, 0xa9, 0xda, 0x21, 0xa1, 0xf4, 0xc9, 0xf8, 0x73,
	0x22, 0x86, 0x9f, 0x06, 0x09, 0xcd, 0x01, 0xd7, 0x69, 0x34, 0xfe, 0xf4, 0xf8, 0xe1, 0x57, 0x00,
	0xb1, 0xb3, 0xf2, 0x01, 0xe2, 0xaf, 0x3b, 0x42, 0xc3, 0x42, 0x78, 0x13, 0x47, 0xa3, 0xd9, 0xcb,
	0xc9, 0xd8, 0xdf, 0x49, 0xe5, 0x67, 0xa2, 0x80, 0x70, 0x81, 0x2c, 0x6e, 0xfc, 0xc6, 0x54, 0x9d,
	0xec, 0x1f, 0x87, 0xf2, 0x33, 0xf1, 0x1e, 0xae, 0xec, 0xca, 0x46, 0x75, 0x92, 0x95, 0xee, 0xc4,
	0xb5, 0x5f, 0x73, 0xbd, 0xc9, 0x9e, 0x1e, 0x09, 0xaf, 0x32, 0xbf, 0xfc, 0x96, 0x6c, 0xb4, 0xab,
	0x6b, 0x15, 0xdb, 0xbe, 0x3e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x3c, 0x20, 0x28, 0xce,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StocksClient is the client API for Stocks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StocksClient interface {
	Upsert(ctx context.Context, in *ListStocks, opts ...grpc.CallOption) (*SuccessResponse, error)
	Get(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*ListStocks, error)
	GetPagination(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type stocksClient struct {
	cc *grpc.ClientConn
}

func NewStocksClient(cc *grpc.ClientConn) StocksClient {
	return &stocksClient{cc}
}

func (c *stocksClient) Upsert(ctx context.Context, in *ListStocks, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/stocks.Stocks/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stocksClient) Get(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*ListStocks, error) {
	out := new(ListStocks)
	err := c.cc.Invoke(ctx, "/stocks.Stocks/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stocksClient) GetPagination(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/stocks.Stocks/GetPagination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StocksServer is the server API for Stocks service.
type StocksServer interface {
	Upsert(context.Context, *ListStocks) (*SuccessResponse, error)
	Get(context.Context, *GetParams) (*ListStocks, error)
	GetPagination(context.Context, *Request) (*Response, error)
}

// UnimplementedStocksServer can be embedded to have forward compatible implementations.
type UnimplementedStocksServer struct {
}

func (*UnimplementedStocksServer) Upsert(ctx context.Context, req *ListStocks) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (*UnimplementedStocksServer) Get(ctx context.Context, req *GetParams) (*ListStocks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedStocksServer) GetPagination(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPagination not implemented")
}

func RegisterStocksServer(s *grpc.Server, srv StocksServer) {
	s.RegisterService(&_Stocks_serviceDesc, srv)
}

func _Stocks_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStocks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StocksServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stocks.Stocks/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StocksServer).Upsert(ctx, req.(*ListStocks))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stocks_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StocksServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stocks.Stocks/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StocksServer).Get(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stocks_GetPagination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StocksServer).GetPagination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stocks.Stocks/GetPagination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StocksServer).GetPagination(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stocks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stocks.Stocks",
	HandlerType: (*StocksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upsert",
			Handler:    _Stocks_Upsert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Stocks_Get_Handler,
		},
		{
			MethodName: "GetPagination",
			Handler:    _Stocks_GetPagination_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/internal/stocks.proto",
}
