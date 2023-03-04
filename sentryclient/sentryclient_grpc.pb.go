// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: sentryclient.proto

package sentryclient

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

// SentryClient is the client API for Sentry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SentryClient interface {
	// SendEvent
	SendEvent(ctx context.Context, in *SendEventReq, opts ...grpc.CallOption) (*EmptyResp, error)
}

type sentryClient struct {
	cc grpc.ClientConnInterface
}

func NewSentryClient(cc grpc.ClientConnInterface) SentryClient {
	return &sentryClient{cc}
}

func (c *sentryClient) SendEvent(ctx context.Context, in *SendEventReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, "/sentryclient.Sentry/SendEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SentryServer is the server API for Sentry service.
// All implementations must embed UnimplementedSentryServer
// for forward compatibility
type SentryServer interface {
	// SendEvent
	SendEvent(context.Context, *SendEventReq) (*EmptyResp, error)
	mustEmbedUnimplementedSentryServer()
}

// UnimplementedSentryServer must be embedded to have forward compatible implementations.
type UnimplementedSentryServer struct {
}

func (UnimplementedSentryServer) SendEvent(context.Context, *SendEventReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEvent not implemented")
}
func (UnimplementedSentryServer) mustEmbedUnimplementedSentryServer() {}

// UnsafeSentryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SentryServer will
// result in compilation errors.
type UnsafeSentryServer interface {
	mustEmbedUnimplementedSentryServer()
}

func RegisterSentryServer(s grpc.ServiceRegistrar, srv SentryServer) {
	s.RegisterService(&Sentry_ServiceDesc, srv)
}

func _Sentry_SendEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentryclient.Sentry/SendEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendEvent(ctx, req.(*SendEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sentry_ServiceDesc is the grpc.ServiceDesc for Sentry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sentry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sentryclient.Sentry",
	HandlerType: (*SentryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEvent",
			Handler:    _Sentry_SendEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentryclient.proto",
}
