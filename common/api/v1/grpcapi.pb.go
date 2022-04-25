// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcapi.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PolarisGRPCClient is the client API for PolarisGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisGRPCClient interface {
	// 客户端上报
	ReportClient(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Response, error)
	// 被调方注册服务实例
	RegisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
	// 被调方反注册服务实例
	DeregisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
	// 统一发现接口
	Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisGRPC_DiscoverClient, error)
	// 被调方上报心跳
	Heartbeat(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
}

type polarisGRPCClient struct {
	cc *grpc.ClientConn
}

func NewPolarisGRPCClient(cc *grpc.ClientConn) PolarisGRPCClient {
	return &polarisGRPCClient{cc}
}

func (c *polarisGRPCClient) ReportClient(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/ReportClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) RegisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/RegisterInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) DeregisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/DeregisterInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisGRPC_DiscoverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PolarisGRPC_serviceDesc.Streams[0], "/v1.PolarisGRPC/Discover", opts...)
	if err != nil {
		return nil, err
	}
	x := &polarisGRPCDiscoverClient{stream}
	return x, nil
}

type PolarisGRPC_DiscoverClient interface {
	Send(*DiscoverRequest) error
	Recv() (*DiscoverResponse, error)
	grpc.ClientStream
}

type polarisGRPCDiscoverClient struct {
	grpc.ClientStream
}

func (x *polarisGRPCDiscoverClient) Send(m *DiscoverRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *polarisGRPCDiscoverClient) Recv() (*DiscoverResponse, error) {
	m := new(DiscoverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *polarisGRPCClient) Heartbeat(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolarisGRPCServer is the server API for PolarisGRPC service.
type PolarisGRPCServer interface {
	// 客户端上报
	ReportClient(context.Context, *Client) (*Response, error)
	// 被调方注册服务实例
	RegisterInstance(context.Context, *Instance) (*Response, error)
	// 被调方反注册服务实例
	DeregisterInstance(context.Context, *Instance) (*Response, error)
	// 统一发现接口
	Discover(PolarisGRPC_DiscoverServer) error
	// 被调方上报心跳
	Heartbeat(context.Context, *Instance) (*Response, error)
}

func RegisterPolarisGRPCServer(s *grpc.Server, srv PolarisGRPCServer) {
	s.RegisterService(&_PolarisGRPC_serviceDesc, srv)
}

func _PolarisGRPC_ReportClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).ReportClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/ReportClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).ReportClient(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_RegisterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).RegisterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/RegisterInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).RegisterInstance(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_DeregisterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).DeregisterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/DeregisterInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).DeregisterInstance(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PolarisGRPCServer).Discover(&polarisGRPCDiscoverServer{stream})
}

type PolarisGRPC_DiscoverServer interface {
	Send(*DiscoverResponse) error
	Recv() (*DiscoverRequest, error)
	grpc.ServerStream
}

type polarisGRPCDiscoverServer struct {
	grpc.ServerStream
}

func (x *polarisGRPCDiscoverServer) Send(m *DiscoverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *polarisGRPCDiscoverServer) Recv() (*DiscoverRequest, error) {
	m := new(DiscoverRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PolarisGRPC_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).Heartbeat(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolarisGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PolarisGRPC",
	HandlerType: (*PolarisGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportClient",
			Handler:    _PolarisGRPC_ReportClient_Handler,
		},
		{
			MethodName: "RegisterInstance",
			Handler:    _PolarisGRPC_RegisterInstance_Handler,
		},
		{
			MethodName: "DeregisterInstance",
			Handler:    _PolarisGRPC_DeregisterInstance_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _PolarisGRPC_Heartbeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _PolarisGRPC_Discover_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcapi.proto",
}

func init() { proto.RegisterFile("grpcapi.proto", fileDescriptor_grpcapi_5af82093fbf54d81) }

var fileDescriptor_grpcapi_5af82093fbf54d81 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbd, 0x4a, 0x04, 0x31,
	0x10, 0xc7, 0xbd, 0x2b, 0x44, 0xc7, 0x3d, 0x91, 0xd1, 0x2a, 0xa5, 0x95, 0x5a, 0x2c, 0x77, 0x6b,
	0x69, 0xb9, 0x0b, 0x6a, 0xb7, 0xe4, 0x0d, 0xb2, 0x61, 0x58, 0x02, 0x4b, 0x12, 0x67, 0x62, 0x5e,
	0xc1, 0xd7, 0x96, 0xfd, 0x48, 0x21, 0x82, 0x60, 0xf7, 0xff, 0xc8, 0x2f, 0xc3, 0x0c, 0x1c, 0x46,
	0x8e, 0xd6, 0x44, 0x57, 0x47, 0x0e, 0x29, 0xe0, 0x3e, 0x9f, 0x54, 0x65, 0x27, 0x47, 0x3e, 0xad,
	0x89, 0x3a, 0x08, 0x71, 0x76, 0x96, 0x8a, 0x65, 0xfa, 0xf8, 0x24, 0x29, 0xed, 0x35, 0x93, 0xc4,
	0xe0, 0x65, 0xab, 0x9b, 0xaf, 0x3d, 0x5c, 0xf5, 0x61, 0x32, 0xec, 0xe4, 0x55, 0xf7, 0x2d, 0x3e,
	0x41, 0xa5, 0x29, 0x06, 0x4e, 0xed, 0xf2, 0x27, 0x42, 0x9d, 0x4f, 0xf5, 0xaa, 0x55, 0x35, 0x6b,
	0xbd, 0xf1, 0xf7, 0x67, 0x78, 0x84, 0x1b, 0x4d, 0xa3, 0x93, 0x44, 0xfc, 0xee, 0x25, 0x19, 0x6f,
	0x09, 0x97, 0x37, 0xc5, 0xfd, 0x22, 0x1a, 0xc0, 0x8e, 0xf8, 0x7f, 0xcc, 0x0b, 0x5c, 0x74, 0x4e,
	0x6c, 0xc8, 0xc4, 0x78, 0x3b, 0x77, 0xc5, 0xe9, 0x75, 0x31, 0x75, 0xf7, 0x33, 0x2c, 0xe0, 0xc3,
	0xee, 0xb8, 0xc3, 0x47, 0xb8, 0x7c, 0x23, 0xc3, 0x69, 0x20, 0x93, 0xfe, 0x9e, 0x33, 0x9c, 0x2f,
	0x07, 0x79, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x5a, 0x87, 0xc9, 0x61, 0x01, 0x00, 0x00,
}
