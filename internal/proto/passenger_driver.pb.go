// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passenger_driver.proto

package protobuf

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

type RequestDriver struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestDriver) Reset()         { *m = RequestDriver{} }
func (m *RequestDriver) String() string { return proto.CompactTextString(m) }
func (*RequestDriver) ProtoMessage()    {}
func (*RequestDriver) Descriptor() ([]byte, []int) {
	return fileDescriptor_7196ee08884a7dc7, []int{0}
}

func (m *RequestDriver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestDriver.Unmarshal(m, b)
}
func (m *RequestDriver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestDriver.Marshal(b, m, deterministic)
}
func (m *RequestDriver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestDriver.Merge(m, src)
}
func (m *RequestDriver) XXX_Size() int {
	return xxx_messageInfo_RequestDriver.Size(m)
}
func (m *RequestDriver) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestDriver.DiscardUnknown(m)
}

var xxx_messageInfo_RequestDriver proto.InternalMessageInfo

type RequestOrder struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestOrder) Reset()         { *m = RequestOrder{} }
func (m *RequestOrder) String() string { return proto.CompactTextString(m) }
func (*RequestOrder) ProtoMessage()    {}
func (*RequestOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_7196ee08884a7dc7, []int{1}
}

func (m *RequestOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestOrder.Unmarshal(m, b)
}
func (m *RequestOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestOrder.Marshal(b, m, deterministic)
}
func (m *RequestOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestOrder.Merge(m, src)
}
func (m *RequestOrder) XXX_Size() int {
	return xxx_messageInfo_RequestOrder.Size(m)
}
func (m *RequestOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestOrder.DiscardUnknown(m)
}

var xxx_messageInfo_RequestOrder proto.InternalMessageInfo

type ResponseDriver struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CarModel             string   `protobuf:"bytes,2,opt,name=carModel,proto3" json:"carModel,omitempty"`
	Latitude             float64  `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseDriver) Reset()         { *m = ResponseDriver{} }
func (m *ResponseDriver) String() string { return proto.CompactTextString(m) }
func (*ResponseDriver) ProtoMessage()    {}
func (*ResponseDriver) Descriptor() ([]byte, []int) {
	return fileDescriptor_7196ee08884a7dc7, []int{2}
}

func (m *ResponseDriver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseDriver.Unmarshal(m, b)
}
func (m *ResponseDriver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseDriver.Marshal(b, m, deterministic)
}
func (m *ResponseDriver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseDriver.Merge(m, src)
}
func (m *ResponseDriver) XXX_Size() int {
	return xxx_messageInfo_ResponseDriver.Size(m)
}
func (m *ResponseDriver) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseDriver.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseDriver proto.InternalMessageInfo

func (m *ResponseDriver) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResponseDriver) GetCarModel() string {
	if m != nil {
		return m.CarModel
	}
	return ""
}

func (m *ResponseDriver) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *ResponseDriver) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type ResponseOrder struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	LatitudeTo           float64  `protobuf:"fixed64,3,opt,name=latitudeTo,proto3" json:"latitudeTo,omitempty"`
	LongitudeTo          float64  `protobuf:"fixed64,4,opt,name=longitudeTo,proto3" json:"longitudeTo,omitempty"`
	LatitudeFrom         float64  `protobuf:"fixed64,5,opt,name=latitudeFrom,proto3" json:"latitudeFrom,omitempty"`
	LongitudeFrom        float64  `protobuf:"fixed64,6,opt,name=longitudeFrom,proto3" json:"longitudeFrom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseOrder) Reset()         { *m = ResponseOrder{} }
func (m *ResponseOrder) String() string { return proto.CompactTextString(m) }
func (*ResponseOrder) ProtoMessage()    {}
func (*ResponseOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_7196ee08884a7dc7, []int{3}
}

func (m *ResponseOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseOrder.Unmarshal(m, b)
}
func (m *ResponseOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseOrder.Marshal(b, m, deterministic)
}
func (m *ResponseOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseOrder.Merge(m, src)
}
func (m *ResponseOrder) XXX_Size() int {
	return xxx_messageInfo_ResponseOrder.Size(m)
}
func (m *ResponseOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseOrder.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseOrder proto.InternalMessageInfo

func (m *ResponseOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResponseOrder) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ResponseOrder) GetLatitudeTo() float64 {
	if m != nil {
		return m.LatitudeTo
	}
	return 0
}

func (m *ResponseOrder) GetLongitudeTo() float64 {
	if m != nil {
		return m.LongitudeTo
	}
	return 0
}

func (m *ResponseOrder) GetLatitudeFrom() float64 {
	if m != nil {
		return m.LatitudeFrom
	}
	return 0
}

func (m *ResponseOrder) GetLongitudeFrom() float64 {
	if m != nil {
		return m.LongitudeFrom
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestDriver)(nil), "protobuf.RequestDriver")
	proto.RegisterType((*RequestOrder)(nil), "protobuf.RequestOrder")
	proto.RegisterType((*ResponseDriver)(nil), "protobuf.ResponseDriver")
	proto.RegisterType((*ResponseOrder)(nil), "protobuf.ResponseOrder")
}

func init() { proto.RegisterFile("passenger_driver.proto", fileDescriptor_7196ee08884a7dc7) }

var fileDescriptor_7196ee08884a7dc7 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x4d, 0x75, 0x65, 0xbb, 0x6e, 0x1d, 0xe4, 0x61, 0x2b, 0x45, 0xa4, 0x04, 0x1f, 0xfa,
	0x54, 0x44, 0xdf, 0x05, 0x41, 0x7c, 0x13, 0xa5, 0xec, 0x5d, 0x3a, 0x73, 0x1d, 0x85, 0xae, 0xa9,
	0x49, 0xbb, 0x5f, 0xe2, 0x7f, 0xf2, 0x6f, 0xc9, 0x6e, 0x93, 0xda, 0x32, 0x9f, 0xda, 0xfb, 0x9d,
	0x93, 0x9c, 0x70, 0x0f, 0xac, 0xea, 0xdc, 0x18, 0xac, 0x76, 0xa8, 0xdf, 0xa5, 0x2e, 0x0e, 0xa8,
	0xd3, 0x5a, 0xab, 0x46, 0xf1, 0x29, 0x7d, 0xb6, 0xed, 0xa7, 0x58, 0xc2, 0x22, 0xc3, 0xaf, 0x16,
	0x4d, 0xf3, 0x44, 0x06, 0x11, 0xc0, 0xdc, 0x82, 0x57, 0x2d, 0x51, 0x8b, 0x03, 0x04, 0x19, 0x9a,
	0x5a, 0x55, 0x06, 0x3b, 0x07, 0x0f, 0xc0, 0x2b, 0x64, 0xc8, 0x62, 0x96, 0xcc, 0x32, 0xaf, 0x90,
	0x3c, 0x82, 0xe9, 0x47, 0xae, 0x5f, 0x94, 0xc4, 0x32, 0xf4, 0x88, 0xf6, 0xf3, 0x51, 0x2b, 0xf3,
	0xa6, 0x68, 0x5a, 0x89, 0xe1, 0x79, 0xcc, 0x12, 0x96, 0xf5, 0x33, 0xbf, 0x82, 0x59, 0xa9, 0xaa,
	0x5d, 0x27, 0x5e, 0x90, 0xf8, 0x07, 0xc4, 0x0f, 0x3b, 0xbe, 0xac, 0x0b, 0xa6, 0x97, 0xfc, 0x97,
	0xdb, 0x1a, 0xd4, 0x55, 0xbe, 0x47, 0x97, 0xeb, 0x66, 0x7e, 0x0d, 0xe0, 0x72, 0x36, 0xca, 0x26,
	0x0f, 0x08, 0x8f, 0xe1, 0xb2, 0x8f, 0xda, 0x28, 0x9b, 0x3e, 0x44, 0x5c, 0xc0, 0xdc, 0xf9, 0x9f,
	0xb5, 0xda, 0x87, 0x13, 0xb2, 0x8c, 0x18, 0xbf, 0x81, 0x45, 0x7f, 0x84, 0x4c, 0x3e, 0x99, 0xc6,
	0xf0, 0xee, 0x9b, 0xc1, 0xf2, 0xcd, 0xf5, 0x60, 0x77, 0xf8, 0x08, 0xbe, 0xfd, 0x5b, 0xa7, 0xae,
	0x8b, 0x74, 0x54, 0x44, 0x14, 0x0e, 0x85, 0x61, 0x01, 0xe2, 0x2c, 0x61, 0xb7, 0x8c, 0x3f, 0xc0,
	0xa4, 0xdb, 0xcb, 0xea, 0xe4, 0x06, 0xe2, 0xd1, 0xfa, 0xf4, 0x82, 0xae, 0x52, 0x3a, 0xbf, 0xf5,
	0x49, 0xbd, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x3c, 0xf9, 0xcc, 0x24, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PassengerDriverClient is the client API for PassengerDriver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PassengerDriverClient interface {
	Driver(ctx context.Context, opts ...grpc.CallOption) (PassengerDriver_DriverClient, error)
	Order(ctx context.Context, opts ...grpc.CallOption) (PassengerDriver_OrderClient, error)
}

type passengerDriverClient struct {
	cc *grpc.ClientConn
}

func NewPassengerDriverClient(cc *grpc.ClientConn) PassengerDriverClient {
	return &passengerDriverClient{cc}
}

func (c *passengerDriverClient) Driver(ctx context.Context, opts ...grpc.CallOption) (PassengerDriver_DriverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PassengerDriver_serviceDesc.Streams[0], "/protobuf.PassengerDriver/Driver", opts...)
	if err != nil {
		return nil, err
	}
	x := &passengerDriverDriverClient{stream}
	return x, nil
}

type PassengerDriver_DriverClient interface {
	Send(*RequestDriver) error
	Recv() (*ResponseDriver, error)
	grpc.ClientStream
}

type passengerDriverDriverClient struct {
	grpc.ClientStream
}

func (x *passengerDriverDriverClient) Send(m *RequestDriver) error {
	return x.ClientStream.SendMsg(m)
}

func (x *passengerDriverDriverClient) Recv() (*ResponseDriver, error) {
	m := new(ResponseDriver)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *passengerDriverClient) Order(ctx context.Context, opts ...grpc.CallOption) (PassengerDriver_OrderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PassengerDriver_serviceDesc.Streams[1], "/protobuf.PassengerDriver/Order", opts...)
	if err != nil {
		return nil, err
	}
	x := &passengerDriverOrderClient{stream}
	return x, nil
}

type PassengerDriver_OrderClient interface {
	Send(*RequestOrder) error
	Recv() (*ResponseOrder, error)
	grpc.ClientStream
}

type passengerDriverOrderClient struct {
	grpc.ClientStream
}

func (x *passengerDriverOrderClient) Send(m *RequestOrder) error {
	return x.ClientStream.SendMsg(m)
}

func (x *passengerDriverOrderClient) Recv() (*ResponseOrder, error) {
	m := new(ResponseOrder)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PassengerDriverServer is the server API for PassengerDriver service.
type PassengerDriverServer interface {
	Driver(PassengerDriver_DriverServer) error
	Order(PassengerDriver_OrderServer) error
}

// UnimplementedPassengerDriverServer can be embedded to have forward compatible implementations.
type UnimplementedPassengerDriverServer struct {
}

func (*UnimplementedPassengerDriverServer) Driver(srv PassengerDriver_DriverServer) error {
	return status.Errorf(codes.Unimplemented, "method Driver not implemented")
}
func (*UnimplementedPassengerDriverServer) Order(srv PassengerDriver_OrderServer) error {
	return status.Errorf(codes.Unimplemented, "method Order not implemented")
}

func RegisterPassengerDriverServer(s *grpc.Server, srv PassengerDriverServer) {
	s.RegisterService(&_PassengerDriver_serviceDesc, srv)
}

func _PassengerDriver_Driver_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PassengerDriverServer).Driver(&passengerDriverDriverServer{stream})
}

type PassengerDriver_DriverServer interface {
	Send(*ResponseDriver) error
	Recv() (*RequestDriver, error)
	grpc.ServerStream
}

type passengerDriverDriverServer struct {
	grpc.ServerStream
}

func (x *passengerDriverDriverServer) Send(m *ResponseDriver) error {
	return x.ServerStream.SendMsg(m)
}

func (x *passengerDriverDriverServer) Recv() (*RequestDriver, error) {
	m := new(RequestDriver)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PassengerDriver_Order_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PassengerDriverServer).Order(&passengerDriverOrderServer{stream})
}

type PassengerDriver_OrderServer interface {
	Send(*ResponseOrder) error
	Recv() (*RequestOrder, error)
	grpc.ServerStream
}

type passengerDriverOrderServer struct {
	grpc.ServerStream
}

func (x *passengerDriverOrderServer) Send(m *ResponseOrder) error {
	return x.ServerStream.SendMsg(m)
}

func (x *passengerDriverOrderServer) Recv() (*RequestOrder, error) {
	m := new(RequestOrder)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PassengerDriver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.PassengerDriver",
	HandlerType: (*PassengerDriverServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Driver",
			Handler:       _PassengerDriver_Driver_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Order",
			Handler:       _PassengerDriver_Order_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "passenger_driver.proto",
}
