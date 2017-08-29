// Code generated by protoc-gen-go. DO NOT EDIT.
// source: els.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	els.proto

It has these top-level messages:
	RoutingKeyRequest
	ServiceInstanceResponse
	ServiceInstanceListResponse
	AddRoutingKeyRequest
	DeleteRoutingKeyRequest
	HealthCheckRequest
	HealthCheckResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type RoutingKeyRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RoutingKeyRequest) Reset()                    { *m = RoutingKeyRequest{} }
func (m *RoutingKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*RoutingKeyRequest) ProtoMessage()               {}
func (*RoutingKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RoutingKeyRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ServiceInstanceResponse struct {
	ServiceUri string `protobuf:"bytes,1,opt,name=serviceUri" json:"serviceUri,omitempty"`
	Tags       string `protobuf:"bytes,2,opt,name=tags" json:"tags,omitempty"`
}

func (m *ServiceInstanceResponse) Reset()                    { *m = ServiceInstanceResponse{} }
func (m *ServiceInstanceResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceInstanceResponse) ProtoMessage()               {}
func (*ServiceInstanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceInstanceResponse) GetServiceUri() string {
	if m != nil {
		return m.ServiceUri
	}
	return ""
}

func (m *ServiceInstanceResponse) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

type ServiceInstanceListResponse struct {
	ServiceInstances []*ServiceInstanceResponse `protobuf:"bytes,1,rep,name=serviceInstances" json:"serviceInstances,omitempty"`
}

func (m *ServiceInstanceListResponse) Reset()                    { *m = ServiceInstanceListResponse{} }
func (m *ServiceInstanceListResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceInstanceListResponse) ProtoMessage()               {}
func (*ServiceInstanceListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceInstanceListResponse) GetServiceInstances() []*ServiceInstanceResponse {
	if m != nil {
		return m.ServiceInstances
	}
	return nil
}

type AddRoutingKeyRequest struct {
	ServiceUri string `protobuf:"bytes,1,opt,name=serviceUri" json:"serviceUri,omitempty"`
	Tags       string `protobuf:"bytes,2,opt,name=tags" json:"tags,omitempty"`
	RoutingKey string `protobuf:"bytes,3,opt,name=routingKey" json:"routingKey,omitempty"`
}

func (m *AddRoutingKeyRequest) Reset()                    { *m = AddRoutingKeyRequest{} }
func (m *AddRoutingKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRoutingKeyRequest) ProtoMessage()               {}
func (*AddRoutingKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddRoutingKeyRequest) GetServiceUri() string {
	if m != nil {
		return m.ServiceUri
	}
	return ""
}

func (m *AddRoutingKeyRequest) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *AddRoutingKeyRequest) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

type DeleteRoutingKeyRequest struct {
	ServiceUri string `protobuf:"bytes,1,opt,name=serviceUri" json:"serviceUri,omitempty"`
	RoutingKey string `protobuf:"bytes,3,opt,name=routingKey" json:"routingKey,omitempty"`
}

func (m *DeleteRoutingKeyRequest) Reset()                    { *m = DeleteRoutingKeyRequest{} }
func (m *DeleteRoutingKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRoutingKeyRequest) ProtoMessage()               {}
func (*DeleteRoutingKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteRoutingKeyRequest) GetServiceUri() string {
	if m != nil {
		return m.ServiceUri
	}
	return ""
}

func (m *DeleteRoutingKeyRequest) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=api.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*RoutingKeyRequest)(nil), "api.RoutingKeyRequest")
	proto.RegisterType((*ServiceInstanceResponse)(nil), "api.ServiceInstanceResponse")
	proto.RegisterType((*ServiceInstanceListResponse)(nil), "api.ServiceInstanceListResponse")
	proto.RegisterType((*AddRoutingKeyRequest)(nil), "api.AddRoutingKeyRequest")
	proto.RegisterType((*DeleteRoutingKeyRequest)(nil), "api.DeleteRoutingKeyRequest")
	proto.RegisterType((*HealthCheckRequest)(nil), "api.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "api.HealthCheckResponse")
	proto.RegisterEnum("api.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Els service

type ElsClient interface {
	// Get a service by routingKey
	GetServiceInstanceByKey(ctx context.Context, in *RoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceResponse, error)
	// Get a service by routingKey
	ListServiceInstances(ctx context.Context, in *RoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceListResponse, error)
	// Add a routingKey to a service
	AddRoutingKey(ctx context.Context, in *AddRoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceResponse, error)
	// Add a routingKey to a service
	RemoveRoutingKey(ctx context.Context, in *DeleteRoutingKeyRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type elsClient struct {
	cc *grpc.ClientConn
}

func NewElsClient(cc *grpc.ClientConn) ElsClient {
	return &elsClient{cc}
}

func (c *elsClient) GetServiceInstanceByKey(ctx context.Context, in *RoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceResponse, error) {
	out := new(ServiceInstanceResponse)
	err := grpc.Invoke(ctx, "/api.Els/GetServiceInstanceByKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elsClient) ListServiceInstances(ctx context.Context, in *RoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceListResponse, error) {
	out := new(ServiceInstanceListResponse)
	err := grpc.Invoke(ctx, "/api.Els/ListServiceInstances", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elsClient) AddRoutingKey(ctx context.Context, in *AddRoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceResponse, error) {
	out := new(ServiceInstanceResponse)
	err := grpc.Invoke(ctx, "/api.Els/AddRoutingKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elsClient) RemoveRoutingKey(ctx context.Context, in *DeleteRoutingKeyRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/api.Els/RemoveRoutingKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Els service

type ElsServer interface {
	// Get a service by routingKey
	GetServiceInstanceByKey(context.Context, *RoutingKeyRequest) (*ServiceInstanceResponse, error)
	// Get a service by routingKey
	ListServiceInstances(context.Context, *RoutingKeyRequest) (*ServiceInstanceListResponse, error)
	// Add a routingKey to a service
	AddRoutingKey(context.Context, *AddRoutingKeyRequest) (*ServiceInstanceResponse, error)
	// Add a routingKey to a service
	RemoveRoutingKey(context.Context, *DeleteRoutingKeyRequest) (*google_protobuf.Empty, error)
}

func RegisterElsServer(s *grpc.Server, srv ElsServer) {
	s.RegisterService(&_Els_serviceDesc, srv)
}

func _Els_GetServiceInstanceByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutingKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElsServer).GetServiceInstanceByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Els/GetServiceInstanceByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElsServer).GetServiceInstanceByKey(ctx, req.(*RoutingKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Els_ListServiceInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutingKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElsServer).ListServiceInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Els/ListServiceInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElsServer).ListServiceInstances(ctx, req.(*RoutingKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Els_AddRoutingKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoutingKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElsServer).AddRoutingKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Els/AddRoutingKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElsServer).AddRoutingKey(ctx, req.(*AddRoutingKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Els_RemoveRoutingKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoutingKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElsServer).RemoveRoutingKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Els/RemoveRoutingKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElsServer).RemoveRoutingKey(ctx, req.(*DeleteRoutingKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Els_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Els",
	HandlerType: (*ElsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceInstanceByKey",
			Handler:    _Els_GetServiceInstanceByKey_Handler,
		},
		{
			MethodName: "ListServiceInstances",
			Handler:    _Els_ListServiceInstances_Handler,
		},
		{
			MethodName: "AddRoutingKey",
			Handler:    _Els_AddRoutingKey_Handler,
		},
		{
			MethodName: "RemoveRoutingKey",
			Handler:    _Els_RemoveRoutingKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "els.proto",
}

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/api.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "els.proto",
}

func init() { proto.RegisterFile("els.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x63, 0x07, 0x52, 0x75, 0xa2, 0x96, 0x30, 0x54, 0x8d, 0x49, 0x11, 0x8a, 0x16, 0x09,
	0xf5, 0xe4, 0x4a, 0xe1, 0xd6, 0x03, 0x12, 0xd0, 0xa8, 0x7f, 0x02, 0xae, 0xd8, 0x50, 0x10, 0x27,
	0xe4, 0x26, 0x83, 0xbb, 0xe0, 0xda, 0x26, 0xbb, 0xae, 0xe4, 0xc7, 0xe0, 0x9d, 0x78, 0x30, 0xe4,
	0x5d, 0xbb, 0xc4, 0x71, 0x4c, 0x05, 0xb7, 0xec, 0xcc, 0x7c, 0xbf, 0x8c, 0x67, 0xbe, 0x81, 0x4d,
	0x0a, 0xa5, 0x9b, 0x2c, 0x62, 0x15, 0x63, 0xdb, 0x4f, 0xc4, 0x60, 0x2f, 0x88, 0xe3, 0x20, 0xa4,
	0x03, 0x1d, 0xba, 0x4c, 0xbf, 0x1e, 0xd0, 0x75, 0xa2, 0x32, 0x53, 0xc1, 0x9e, 0xc1, 0x43, 0x1e,
	0xa7, 0x4a, 0x44, 0xc1, 0x84, 0x32, 0x4e, 0x3f, 0x52, 0x92, 0x0a, 0xb7, 0xc1, 0x16, 0x73, 0xc7,
	0x1a, 0x5a, 0xfb, 0x9b, 0xdc, 0x16, 0x73, 0xf6, 0x0e, 0xfa, 0x53, 0x5a, 0xdc, 0x88, 0x19, 0x9d,
	0x46, 0x52, 0xf9, 0xd1, 0x8c, 0x38, 0xc9, 0x24, 0x8e, 0x24, 0xe1, 0x53, 0x00, 0x69, 0x52, 0x17,
	0x0b, 0x51, 0x48, 0x96, 0x22, 0x88, 0x70, 0x4f, 0xf9, 0x81, 0x74, 0x6c, 0x9d, 0xd1, 0xbf, 0x59,
	0x00, 0x7b, 0x2b, 0xb8, 0xb7, 0x42, 0xaa, 0x5b, 0xe4, 0x09, 0xf4, 0x64, 0x35, 0x2d, 0x1d, 0x6b,
	0xd8, 0xde, 0xef, 0x8e, 0x9e, 0xb8, 0x7e, 0x22, 0xdc, 0x86, 0x56, 0x78, 0x4d, 0xc5, 0xbe, 0xc1,
	0xce, 0xab, 0xf9, 0xbc, 0xfe, 0x7d, 0xff, 0xd1, 0x74, 0xae, 0x59, 0xdc, 0x82, 0x9c, 0xb6, 0xd1,
	0xfc, 0x89, 0xb0, 0xcf, 0xd0, 0x3f, 0xa2, 0x90, 0x14, 0xfd, 0xfb, 0xdf, 0xdd, 0x85, 0x76, 0x01,
	0x4f, 0xc8, 0x0f, 0xd5, 0xd5, 0x9b, 0x2b, 0x9a, 0x7d, 0x2f, 0xa9, 0x0e, 0x6c, 0x14, 0x8c, 0x02,
	0x59, 0x3e, 0xd9, 0x4f, 0x0b, 0x1e, 0x55, 0x04, 0xc5, 0x60, 0x5f, 0x42, 0x47, 0x2a, 0x5f, 0xa5,
	0x52, 0x0b, 0xb6, 0x47, 0xcf, 0xf5, 0x38, 0xd7, 0x54, 0x9a, 0x11, 0x47, 0xc1, 0x54, 0x57, 0xf3,
	0x42, 0xc5, 0x0e, 0x61, 0xab, 0x92, 0xc0, 0x2e, 0x6c, 0x5c, 0x78, 0x13, 0xef, 0xfc, 0x93, 0xd7,
	0x6b, 0xe5, 0x8f, 0xe9, 0x98, 0x7f, 0x3c, 0xf5, 0x8e, 0x7b, 0x16, 0x3e, 0x80, 0xae, 0x77, 0xfe,
	0xe1, 0x4b, 0x19, 0xb0, 0x47, 0xbf, 0x6c, 0x68, 0x8f, 0x43, 0x89, 0xef, 0xa1, 0x7f, 0x4c, 0x6a,
	0x65, 0x85, 0xaf, 0xb3, 0x09, 0x65, 0xb8, 0xab, 0xdb, 0xa9, 0x8d, 0x6f, 0xf0, 0xd7, 0xad, 0xb3,
	0x16, 0x72, 0xd8, 0xc9, 0xfd, 0xb3, 0x52, 0x20, 0x1b, 0x79, 0xc3, 0x75, 0xbc, 0x65, 0x07, 0xb2,
	0x16, 0x9e, 0xc1, 0x56, 0xc5, 0x39, 0xf8, 0x58, 0x8b, 0xd6, 0xb9, 0xe9, 0xce, 0xfe, 0xce, 0xa0,
	0xc7, 0xe9, 0x3a, 0xbe, 0x59, 0x72, 0x06, 0x1a, 0x4d, 0x83, 0x61, 0x06, 0xbb, 0xae, 0x39, 0x59,
	0xb7, 0x3c, 0x59, 0x77, 0x9c, 0x9f, 0x2c, 0x6b, 0x8d, 0x8e, 0xa0, 0x63, 0xf6, 0x85, 0x87, 0x70,
	0x5f, 0xef, 0x0c, 0xfb, 0xf5, 0x2d, 0x1a, 0x8a, 0xd3, 0xb4, 0xde, 0xcb, 0x8e, 0xe6, 0xbe, 0xf8,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x72, 0x22, 0x9d, 0x8c, 0x2a, 0x04, 0x00, 0x00,
}
