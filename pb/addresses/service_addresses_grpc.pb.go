// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: service_addresses.proto

package addresses

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Addresses_CreateAddress_FullMethodName       = "/pb.Addresses/CreateAddress"
	Addresses_DeleteAddress_FullMethodName       = "/pb.Addresses/DeleteAddress"
	Addresses_GetAddress_FullMethodName          = "/pb.Addresses/GetAddress"
	Addresses_ResetDefaultAddress_FullMethodName = "/pb.Addresses/ResetDefaultAddress"
)

// AddressesClient is the client API for Addresses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressesClient interface {
	CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*CreateAddressResponse, error)
	DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
	ResetDefaultAddress(ctx context.Context, in *ResetDefaultAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type addressesClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressesClient(cc grpc.ClientConnInterface) AddressesClient {
	return &addressesClient{cc}
}

func (c *addressesClient) CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*CreateAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAddressResponse)
	err := c.cc.Invoke(ctx, Addresses_CreateAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Addresses_DeleteAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, Addresses_GetAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) ResetDefaultAddress(ctx context.Context, in *ResetDefaultAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Addresses_ResetDefaultAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressesServer is the server API for Addresses service.
// All implementations must embed UnimplementedAddressesServer
// for forward compatibility.
type AddressesServer interface {
	CreateAddress(context.Context, *CreateAddressRequest) (*CreateAddressResponse, error)
	DeleteAddress(context.Context, *DeleteAddressRequest) (*emptypb.Empty, error)
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
	ResetDefaultAddress(context.Context, *ResetDefaultAddressRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAddressesServer()
}

// UnimplementedAddressesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAddressesServer struct{}

func (UnimplementedAddressesServer) CreateAddress(context.Context, *CreateAddressRequest) (*CreateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (UnimplementedAddressesServer) DeleteAddress(context.Context, *DeleteAddressRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (UnimplementedAddressesServer) GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedAddressesServer) ResetDefaultAddress(context.Context, *ResetDefaultAddressRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetDefaultAddress not implemented")
}
func (UnimplementedAddressesServer) mustEmbedUnimplementedAddressesServer() {}
func (UnimplementedAddressesServer) testEmbeddedByValue()                   {}

// UnsafeAddressesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressesServer will
// result in compilation errors.
type UnsafeAddressesServer interface {
	mustEmbedUnimplementedAddressesServer()
}

func RegisterAddressesServer(s grpc.ServiceRegistrar, srv AddressesServer) {
	// If the following call pancis, it indicates UnimplementedAddressesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Addresses_ServiceDesc, srv)
}

func _Addresses_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Addresses_CreateAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).CreateAddress(ctx, req.(*CreateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Addresses_DeleteAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).DeleteAddress(ctx, req.(*DeleteAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Addresses_GetAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_ResetDefaultAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetDefaultAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).ResetDefaultAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Addresses_ResetDefaultAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).ResetDefaultAddress(ctx, req.(*ResetDefaultAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Addresses_ServiceDesc is the grpc.ServiceDesc for Addresses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Addresses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Addresses",
	HandlerType: (*AddressesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAddress",
			Handler:    _Addresses_CreateAddress_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _Addresses_DeleteAddress_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _Addresses_GetAddress_Handler,
		},
		{
			MethodName: "ResetDefaultAddress",
			Handler:    _Addresses_ResetDefaultAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_addresses.proto",
}
