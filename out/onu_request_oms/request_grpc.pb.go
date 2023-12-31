// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: request.proto

package onu_request_oms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Request_RequestCondition_FullMethodName = "/onu_request_oms.Request/RequestCondition"
)

// RequestClient is the client API for Request service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestClient interface {
	RequestCondition(ctx context.Context, in *ConditionRequest, opts ...grpc.CallOption) (*ConditionReply, error)
}

type requestClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestClient(cc grpc.ClientConnInterface) RequestClient {
	return &requestClient{cc}
}

func (c *requestClient) RequestCondition(ctx context.Context, in *ConditionRequest, opts ...grpc.CallOption) (*ConditionReply, error) {
	out := new(ConditionReply)
	err := c.cc.Invoke(ctx, Request_RequestCondition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServer is the server API for Request service.
// All implementations must embed UnimplementedRequestServer
// for forward compatibility
type RequestServer interface {
	RequestCondition(context.Context, *ConditionRequest) (*ConditionReply, error)
	mustEmbedUnimplementedRequestServer()
}

// UnimplementedRequestServer must be embedded to have forward compatible implementations.
type UnimplementedRequestServer struct {
}

func (UnimplementedRequestServer) RequestCondition(context.Context, *ConditionRequest) (*ConditionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestCondition not implemented")
}
func (UnimplementedRequestServer) mustEmbedUnimplementedRequestServer() {}

// UnsafeRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServer will
// result in compilation errors.
type UnsafeRequestServer interface {
	mustEmbedUnimplementedRequestServer()
}

func RegisterRequestServer(s grpc.ServiceRegistrar, srv RequestServer) {
	s.RegisterService(&Request_ServiceDesc, srv)
}

func _Request_RequestCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServer).RequestCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Request_RequestCondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServer).RequestCondition(ctx, req.(*ConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Request_ServiceDesc is the grpc.ServiceDesc for Request service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Request_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "onu_request_oms.Request",
	HandlerType: (*RequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestCondition",
			Handler:    _Request_RequestCondition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "request.proto",
}
