// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: service_session.proto

package session

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Session_CreateSessionId_FullMethodName = "/pb.Session/CreateSessionId"
)

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionClient interface {
	CreateSessionId(ctx context.Context, in *CreateSessionIdRequest, opts ...grpc.CallOption) (*CreateSessionIdResponse, error)
}

type sessionClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionClient(cc grpc.ClientConnInterface) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) CreateSessionId(ctx context.Context, in *CreateSessionIdRequest, opts ...grpc.CallOption) (*CreateSessionIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSessionIdResponse)
	err := c.cc.Invoke(ctx, Session_CreateSessionId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServer is the server API for Session service.
// All implementations must embed UnimplementedSessionServer
// for forward compatibility.
type SessionServer interface {
	CreateSessionId(context.Context, *CreateSessionIdRequest) (*CreateSessionIdResponse, error)
	mustEmbedUnimplementedSessionServer()
}

// UnimplementedSessionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSessionServer struct{}

func (UnimplementedSessionServer) CreateSessionId(context.Context, *CreateSessionIdRequest) (*CreateSessionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSessionId not implemented")
}
func (UnimplementedSessionServer) mustEmbedUnimplementedSessionServer() {}
func (UnimplementedSessionServer) testEmbeddedByValue()                 {}

// UnsafeSessionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServer will
// result in compilation errors.
type UnsafeSessionServer interface {
	mustEmbedUnimplementedSessionServer()
}

func RegisterSessionServer(s grpc.ServiceRegistrar, srv SessionServer) {
	// If the following call pancis, it indicates UnimplementedSessionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Session_ServiceDesc, srv)
}

func _Session_CreateSessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).CreateSessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_CreateSessionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).CreateSessionId(ctx, req.(*CreateSessionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Session_ServiceDesc is the grpc.ServiceDesc for Session service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Session_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSessionId",
			Handler:    _Session_CreateSessionId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_session.proto",
}
