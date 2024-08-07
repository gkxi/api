// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: address/v1/address_v1.proto

package v1

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
	AddressV1_NewBip44_FullMethodName  = "/address.v1.AddressV1/NewBip44"
	AddressV1_NewBip441_FullMethodName = "/address.v1.AddressV1/NewBip441"
)

// AddressV1Client is the client API for AddressV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressV1Client interface {
	NewBip44(ctx context.Context, in *NewBip44Request, opts ...grpc.CallOption) (*NewBip44Reply, error)
	NewBip441(ctx context.Context, in *NewBip44Request, opts ...grpc.CallOption) (*NewBip441Reply, error)
}

type addressV1Client struct {
	cc grpc.ClientConnInterface
}

func NewAddressV1Client(cc grpc.ClientConnInterface) AddressV1Client {
	return &addressV1Client{cc}
}

func (c *addressV1Client) NewBip44(ctx context.Context, in *NewBip44Request, opts ...grpc.CallOption) (*NewBip44Reply, error) {
	out := new(NewBip44Reply)
	err := c.cc.Invoke(ctx, AddressV1_NewBip44_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressV1Client) NewBip441(ctx context.Context, in *NewBip44Request, opts ...grpc.CallOption) (*NewBip441Reply, error) {
	out := new(NewBip441Reply)
	err := c.cc.Invoke(ctx, AddressV1_NewBip441_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressV1Server is the server API for AddressV1 service.
// All implementations must embed UnimplementedAddressV1Server
// for forward compatibility
type AddressV1Server interface {
	NewBip44(context.Context, *NewBip44Request) (*NewBip44Reply, error)
	NewBip441(context.Context, *NewBip44Request) (*NewBip441Reply, error)
	mustEmbedUnimplementedAddressV1Server()
}

// UnimplementedAddressV1Server must be embedded to have forward compatible implementations.
type UnimplementedAddressV1Server struct {
}

func (UnimplementedAddressV1Server) NewBip44(context.Context, *NewBip44Request) (*NewBip44Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBip44 not implemented")
}
func (UnimplementedAddressV1Server) NewBip441(context.Context, *NewBip44Request) (*NewBip441Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBip441 not implemented")
}
func (UnimplementedAddressV1Server) mustEmbedUnimplementedAddressV1Server() {}

// UnsafeAddressV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressV1Server will
// result in compilation errors.
type UnsafeAddressV1Server interface {
	mustEmbedUnimplementedAddressV1Server()
}

func RegisterAddressV1Server(s grpc.ServiceRegistrar, srv AddressV1Server) {
	s.RegisterService(&AddressV1_ServiceDesc, srv)
}

func _AddressV1_NewBip44_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBip44Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressV1Server).NewBip44(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressV1_NewBip44_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressV1Server).NewBip44(ctx, req.(*NewBip44Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressV1_NewBip441_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBip44Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressV1Server).NewBip441(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressV1_NewBip441_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressV1Server).NewBip441(ctx, req.(*NewBip44Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressV1_ServiceDesc is the grpc.ServiceDesc for AddressV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "address.v1.AddressV1",
	HandlerType: (*AddressV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBip44",
			Handler:    _AddressV1_NewBip44_Handler,
		},
		{
			MethodName: "NewBip441",
			Handler:    _AddressV1_NewBip441_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address/v1/address_v1.proto",
}
