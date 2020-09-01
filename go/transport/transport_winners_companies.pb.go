// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/transport_winners_companies.proto

package transport

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

type GeoType int32

const (
	GeoType_dadata   GeoType = 0
	GeoType_geonames GeoType = 1
)

var GeoType_name = map[int32]string{
	0: "dadata",
	1: "geonames",
}

var GeoType_value = map[string]int32{
	"dadata":   0,
	"geonames": 1,
}

func (x GeoType) String() string {
	return proto.EnumName(GeoType_name, int32(x))
}

func (GeoType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2ffd909b5960fa39, []int{0}
}

type WinnersParams struct {
	GeoId      string  `protobuf:"bytes,1,opt,name=geo_id,json=geoId,proto3" json:"geo_id,omitempty"`
	GeoTypeId  GeoType `protobuf:"varint,2,opt,name=geo_type_id,json=geoTypeId,proto3,enum=transport.GeoType" json:"geo_type_id,omitempty"`
	OrderPrice int32   `protobuf:"varint,3,opt,name=order_price,json=orderPrice,proto3" json:"order_price,omitempty"`
	// в рублях
	Locale               string   `protobuf:"bytes,4,opt,name=locale,proto3" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WinnersParams) Reset()         { *m = WinnersParams{} }
func (m *WinnersParams) String() string { return proto.CompactTextString(m) }
func (*WinnersParams) ProtoMessage()    {}
func (*WinnersParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffd909b5960fa39, []int{0}
}

func (m *WinnersParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WinnersParams.Unmarshal(m, b)
}
func (m *WinnersParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WinnersParams.Marshal(b, m, deterministic)
}
func (m *WinnersParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WinnersParams.Merge(m, src)
}
func (m *WinnersParams) XXX_Size() int {
	return xxx_messageInfo_WinnersParams.Size(m)
}
func (m *WinnersParams) XXX_DiscardUnknown() {
	xxx_messageInfo_WinnersParams.DiscardUnknown(m)
}

var xxx_messageInfo_WinnersParams proto.InternalMessageInfo

func (m *WinnersParams) GetGeoId() string {
	if m != nil {
		return m.GeoId
	}
	return ""
}

func (m *WinnersParams) GetGeoTypeId() GeoType {
	if m != nil {
		return m.GeoTypeId
	}
	return GeoType_dadata
}

func (m *WinnersParams) GetOrderPrice() int32 {
	if m != nil {
		return m.OrderPrice
	}
	return 0
}

func (m *WinnersParams) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type WinnersResponse struct {
	Result               []*Winner `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WinnersResponse) Reset()         { *m = WinnersResponse{} }
func (m *WinnersResponse) String() string { return proto.CompactTextString(m) }
func (*WinnersResponse) ProtoMessage()    {}
func (*WinnersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffd909b5960fa39, []int{1}
}

func (m *WinnersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WinnersResponse.Unmarshal(m, b)
}
func (m *WinnersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WinnersResponse.Marshal(b, m, deterministic)
}
func (m *WinnersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WinnersResponse.Merge(m, src)
}
func (m *WinnersResponse) XXX_Size() int {
	return xxx_messageInfo_WinnersResponse.Size(m)
}
func (m *WinnersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WinnersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WinnersResponse proto.InternalMessageInfo

func (m *WinnersResponse) GetResult() []*Winner {
	if m != nil {
		return m.Result
	}
	return nil
}

type Winner struct {
	DeliveryTypeId    DeliveryType     `protobuf:"varint,1,opt,name=delivery_type_id,json=deliveryTypeId,proto3,enum=transport.DeliveryType" json:"delivery_type_id,omitempty"`
	DeliveryTypeTitle string           `protobuf:"bytes,2,opt,name=delivery_type_title,json=deliveryTypeTitle,proto3" json:"delivery_type_title,omitempty"`
	WinnerCompany     *DeliveryCompany `protobuf:"bytes,3,opt,name=winner_company,json=winnerCompany,proto3" json:"winner_company,omitempty"`
	Payments          []*PaymentType   `protobuf:"bytes,4,rep,name=payments,proto3" json:"payments,omitempty"`
	// 'YYYY-MM-DD'
	DeliveryDate         []*DeliveryTime `protobuf:"bytes,6,rep,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Winner) Reset()         { *m = Winner{} }
func (m *Winner) String() string { return proto.CompactTextString(m) }
func (*Winner) ProtoMessage()    {}
func (*Winner) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffd909b5960fa39, []int{2}
}

func (m *Winner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Winner.Unmarshal(m, b)
}
func (m *Winner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Winner.Marshal(b, m, deterministic)
}
func (m *Winner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Winner.Merge(m, src)
}
func (m *Winner) XXX_Size() int {
	return xxx_messageInfo_Winner.Size(m)
}
func (m *Winner) XXX_DiscardUnknown() {
	xxx_messageInfo_Winner.DiscardUnknown(m)
}

var xxx_messageInfo_Winner proto.InternalMessageInfo

func (m *Winner) GetDeliveryTypeId() DeliveryType {
	if m != nil {
		return m.DeliveryTypeId
	}
	return DeliveryType_courier
}

func (m *Winner) GetDeliveryTypeTitle() string {
	if m != nil {
		return m.DeliveryTypeTitle
	}
	return ""
}

func (m *Winner) GetWinnerCompany() *DeliveryCompany {
	if m != nil {
		return m.WinnerCompany
	}
	return nil
}

func (m *Winner) GetPayments() []*PaymentType {
	if m != nil {
		return m.Payments
	}
	return nil
}

func (m *Winner) GetDeliveryDate() []*DeliveryTime {
	if m != nil {
		return m.DeliveryDate
	}
	return nil
}

type DeliveryCompany struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliveryCompany) Reset()         { *m = DeliveryCompany{} }
func (m *DeliveryCompany) String() string { return proto.CompactTextString(m) }
func (*DeliveryCompany) ProtoMessage()    {}
func (*DeliveryCompany) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffd909b5960fa39, []int{3}
}

func (m *DeliveryCompany) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryCompany.Unmarshal(m, b)
}
func (m *DeliveryCompany) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryCompany.Marshal(b, m, deterministic)
}
func (m *DeliveryCompany) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryCompany.Merge(m, src)
}
func (m *DeliveryCompany) XXX_Size() int {
	return xxx_messageInfo_DeliveryCompany.Size(m)
}
func (m *DeliveryCompany) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryCompany.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryCompany proto.InternalMessageInfo

func (m *DeliveryCompany) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DeliveryCompany) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type PaymentType struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	DeliveryPrice        int32    `protobuf:"varint,3,opt,name=delivery_price,json=deliveryPrice,proto3" json:"delivery_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentType) Reset()         { *m = PaymentType{} }
func (m *PaymentType) String() string { return proto.CompactTextString(m) }
func (*PaymentType) ProtoMessage()    {}
func (*PaymentType) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffd909b5960fa39, []int{4}
}

func (m *PaymentType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentType.Unmarshal(m, b)
}
func (m *PaymentType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentType.Marshal(b, m, deterministic)
}
func (m *PaymentType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentType.Merge(m, src)
}
func (m *PaymentType) XXX_Size() int {
	return xxx_messageInfo_PaymentType.Size(m)
}
func (m *PaymentType) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentType.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentType proto.InternalMessageInfo

func (m *PaymentType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PaymentType) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PaymentType) GetDeliveryPrice() int32 {
	if m != nil {
		return m.DeliveryPrice
	}
	return 0
}

type DeliveryTime struct {
	Date                 string      `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Intervals            []*Interval `protobuf:"bytes,2,rep,name=intervals,proto3" json:"intervals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeliveryTime) Reset()         { *m = DeliveryTime{} }
func (m *DeliveryTime) String() string { return proto.CompactTextString(m) }
func (*DeliveryTime) ProtoMessage()    {}
func (*DeliveryTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffd909b5960fa39, []int{5}
}

func (m *DeliveryTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryTime.Unmarshal(m, b)
}
func (m *DeliveryTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryTime.Marshal(b, m, deterministic)
}
func (m *DeliveryTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryTime.Merge(m, src)
}
func (m *DeliveryTime) XXX_Size() int {
	return xxx_messageInfo_DeliveryTime.Size(m)
}
func (m *DeliveryTime) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryTime.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryTime proto.InternalMessageInfo

func (m *DeliveryTime) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *DeliveryTime) GetIntervals() []*Interval {
	if m != nil {
		return m.Intervals
	}
	return nil
}

type Interval struct {
	// 10:30 for example
	TimeFrom             string   `protobuf:"bytes,1,opt,name=time_from,json=timeFrom,proto3" json:"time_from,omitempty"`
	TimeTo               string   `protobuf:"bytes,2,opt,name=time_to,json=timeTo,proto3" json:"time_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Interval) Reset()         { *m = Interval{} }
func (m *Interval) String() string { return proto.CompactTextString(m) }
func (*Interval) ProtoMessage()    {}
func (*Interval) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffd909b5960fa39, []int{6}
}

func (m *Interval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Interval.Unmarshal(m, b)
}
func (m *Interval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Interval.Marshal(b, m, deterministic)
}
func (m *Interval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interval.Merge(m, src)
}
func (m *Interval) XXX_Size() int {
	return xxx_messageInfo_Interval.Size(m)
}
func (m *Interval) XXX_DiscardUnknown() {
	xxx_messageInfo_Interval.DiscardUnknown(m)
}

var xxx_messageInfo_Interval proto.InternalMessageInfo

func (m *Interval) GetTimeFrom() string {
	if m != nil {
		return m.TimeFrom
	}
	return ""
}

func (m *Interval) GetTimeTo() string {
	if m != nil {
		return m.TimeTo
	}
	return ""
}

func init() {
	proto.RegisterEnum("transport.GeoType", GeoType_name, GeoType_value)
	proto.RegisterType((*WinnersParams)(nil), "transport.WinnersParams")
	proto.RegisterType((*WinnersResponse)(nil), "transport.WinnersResponse")
	proto.RegisterType((*Winner)(nil), "transport.winner")
	proto.RegisterType((*DeliveryCompany)(nil), "transport.deliveryCompany")
	proto.RegisterType((*PaymentType)(nil), "transport.paymentType")
	proto.RegisterType((*DeliveryTime)(nil), "transport.deliveryTime")
	proto.RegisterType((*Interval)(nil), "transport.interval")
}

func init() {
	proto.RegisterFile("proto/transport_winners_companies.proto", fileDescriptor_2ffd909b5960fa39)
}

var fileDescriptor_2ffd909b5960fa39 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0x39, 0x4d, 0x9c, 0x78, 0xf2, 0xa7, 0xe9, 0xf4, 0x47, 0x63, 0x85, 0x03, 0x91, 0x51,
	0x45, 0xe0, 0x10, 0x84, 0x39, 0x52, 0x09, 0x4a, 0x2b, 0x50, 0x6e, 0x95, 0x09, 0x42, 0xe2, 0x80,
	0xb5, 0x64, 0x87, 0x60, 0xc9, 0xf6, 0x5a, 0xeb, 0xa5, 0x28, 0x9f, 0x83, 0x2f, 0xc2, 0x47, 0x44,
	0x5e, 0xaf, 0xcd, 0xa6, 0x70, 0xe0, 0xe6, 0x99, 0x79, 0x6f, 0xe6, 0xbd, 0xf5, 0x0c, 0x3c, 0x2a,
	0xa4, 0x50, 0xe2, 0xa9, 0x92, 0x2c, 0x2f, 0x0b, 0x21, 0x55, 0xfc, 0x3d, 0xc9, 0x73, 0x92, 0x65,
	0xbc, 0x15, 0x59, 0xc1, 0xf2, 0x84, 0xca, 0x95, 0x46, 0xa0, 0xd7, 0x42, 0xe6, 0xe7, 0x77, 0x39,
	0x9c, 0xd2, 0xe4, 0x96, 0xe4, 0x3e, 0xce, 0x48, 0x7d, 0x15, 0xbc, 0x66, 0x04, 0x3f, 0x1c, 0x18,
	0x7f, 0xa8, 0xbb, 0xdd, 0x30, 0xc9, 0xb2, 0x12, 0xef, 0x81, 0xbb, 0x23, 0x11, 0x27, 0xdc, 0x77,
	0x16, 0xce, 0xd2, 0x8b, 0x7a, 0x3b, 0x12, 0x6b, 0x8e, 0x21, 0x0c, 0xab, 0xb4, 0xda, 0x17, 0x54,
	0xd5, 0x3a, 0x0b, 0x67, 0x39, 0x09, 0x71, 0xd5, 0xf6, 0x5f, 0xbd, 0x25, 0xb1, 0xd9, 0x17, 0x14,
	0x79, 0xbb, 0xfa, 0x63, 0xcd, 0xf1, 0x01, 0x0c, 0x85, 0xe4, 0x24, 0xe3, 0x42, 0x26, 0x5b, 0xf2,
	0x8f, 0x16, 0xce, 0xb2, 0x17, 0x81, 0x4e, 0xdd, 0x54, 0x19, 0x3c, 0x03, 0x37, 0x15, 0x5b, 0x96,
	0x92, 0xdf, 0xd5, 0xb3, 0x4c, 0x14, 0x5c, 0xc0, 0xb1, 0x11, 0x15, 0x51, 0x59, 0x88, 0xbc, 0x24,
	0x7c, 0x0c, 0xae, 0xa4, 0xf2, 0x5b, 0xaa, 0x7c, 0x67, 0x71, 0xb4, 0x1c, 0x86, 0x27, 0xd6, 0xe8,
	0xfa, 0x39, 0x22, 0x03, 0x08, 0x7e, 0x76, 0xc0, 0xad, 0x53, 0x78, 0x09, 0xd3, 0xd6, 0x77, 0x23,
	0xdd, 0xd1, 0xd2, 0x67, 0x16, 0xff, 0xda, 0x40, 0xb4, 0xfe, 0x09, 0xb7, 0xa2, 0x35, 0xc7, 0x15,
	0x9c, 0x1e, 0xb6, 0x50, 0x89, 0x4a, 0x49, 0x3f, 0x80, 0x17, 0x9d, 0xd8, 0xe0, 0x4d, 0x55, 0xc0,
	0x4b, 0x98, 0xd4, 0xc3, 0xcd, 0xdf, 0xd9, 0x6b, 0xdf, 0xc3, 0x70, 0x6e, 0x0d, 0x6c, 0x58, 0x57,
	0x35, 0x22, 0x1a, 0xd7, 0x0c, 0x13, 0x62, 0x08, 0x83, 0x82, 0xed, 0x33, 0xca, 0x55, 0xe9, 0x77,
	0xb5, 0xdb, 0x33, 0x8b, 0x6c, 0x4a, 0x5a, 0x6c, 0x8b, 0xc3, 0x0b, 0x18, 0xb7, 0x32, 0x39, 0x53,
	0xe4, 0xbb, 0x9a, 0x38, 0xfb, 0xcb, 0xd4, 0x4d, 0x92, 0x51, 0x34, 0x6a, 0xa2, 0x6b, 0xa6, 0x28,
	0x78, 0x01, 0xc7, 0x77, 0x34, 0x21, 0x42, 0x77, 0x2b, 0x38, 0x99, 0x2d, 0xd0, 0xdf, 0xf8, 0x3f,
	0xf4, 0x6c, 0xf7, 0x75, 0x10, 0x7c, 0x82, 0xa1, 0xa5, 0xe9, 0xdf, 0x89, 0x78, 0x0e, 0xed, 0x63,
	0x1f, 0xac, 0x48, 0xeb, 0x44, 0x6f, 0x49, 0xf0, 0x1e, 0x46, 0xb6, 0xf4, 0x6a, 0x80, 0x76, 0x68,
	0x06, 0x54, 0xdf, 0xf8, 0x0c, 0xbc, 0x24, 0x57, 0x24, 0x6f, 0x59, 0x5a, 0xfa, 0x1d, 0x6d, 0xfd,
	0xd4, 0xb2, 0xde, 0xd4, 0xa2, 0xdf, 0xa8, 0xe0, 0x15, 0x0c, 0x9a, 0x00, 0xef, 0x83, 0xa7, 0x92,
	0x8c, 0xe2, 0x2f, 0x52, 0x64, 0xa6, 0xef, 0xa0, 0x4a, 0xbc, 0x91, 0x22, 0xc3, 0x19, 0xf4, 0x75,
	0x51, 0x09, 0x23, 0xdf, 0xad, 0xc2, 0x8d, 0x78, 0xf2, 0x10, 0xfa, 0x66, 0xeb, 0x11, 0xc0, 0xe5,
	0x8c, 0x33, 0xc5, 0xa6, 0xff, 0xe1, 0x08, 0x06, 0x3b, 0x12, 0x39, 0xcb, 0xa8, 0x9c, 0x3a, 0xe1,
	0x3b, 0x98, 0x9a, 0x5d, 0xbe, 0x6a, 0xae, 0x15, 0x5f, 0x42, 0xdf, 0xe4, 0xd0, 0xb7, 0x54, 0x1e,
	0x1c, 0xe2, 0x7c, 0xfe, 0x67, 0xa5, 0xb9, 0x86, 0xd7, 0x93, 0x8f, 0xa3, 0x9d, 0x75, 0xdc, 0x9f,
	0x5d, 0x7d, 0xcd, 0xcf, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xef, 0x32, 0xd9, 0xaa, 0x2a, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WinnersCompaniesClient is the client API for WinnersCompanies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WinnersCompaniesClient interface {
	Winners(ctx context.Context, in *WinnersParams, opts ...grpc.CallOption) (*WinnersResponse, error)
}

type winnersCompaniesClient struct {
	cc *grpc.ClientConn
}

func NewWinnersCompaniesClient(cc *grpc.ClientConn) WinnersCompaniesClient {
	return &winnersCompaniesClient{cc}
}

func (c *winnersCompaniesClient) Winners(ctx context.Context, in *WinnersParams, opts ...grpc.CallOption) (*WinnersResponse, error) {
	out := new(WinnersResponse)
	err := c.cc.Invoke(ctx, "/transport.WinnersCompanies/Winners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WinnersCompaniesServer is the server API for WinnersCompanies service.
type WinnersCompaniesServer interface {
	Winners(context.Context, *WinnersParams) (*WinnersResponse, error)
}

// UnimplementedWinnersCompaniesServer can be embedded to have forward compatible implementations.
type UnimplementedWinnersCompaniesServer struct {
}

func (*UnimplementedWinnersCompaniesServer) Winners(ctx context.Context, req *WinnersParams) (*WinnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Winners not implemented")
}

func RegisterWinnersCompaniesServer(s *grpc.Server, srv WinnersCompaniesServer) {
	s.RegisterService(&_WinnersCompanies_serviceDesc, srv)
}

func _WinnersCompanies_Winners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WinnersParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WinnersCompaniesServer).Winners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.WinnersCompanies/Winners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WinnersCompaniesServer).Winners(ctx, req.(*WinnersParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _WinnersCompanies_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transport.WinnersCompanies",
	HandlerType: (*WinnersCompaniesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Winners",
			Handler:    _WinnersCompanies_Winners_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transport_winners_companies.proto",
}