// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: tran/v1/api.proto

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
	TranV1_ChainList_FullMethodName               = "/tran.v1.TranV1/ChainList"
	TranV1_IsMultiSigAddress_FullMethodName       = "/tran.v1.TranV1/IsMultiSigAddress"
	TranV1_Balance_FullMethodName                 = "/tran.v1.TranV1/Balance"
	TranV1_SendTran_FullMethodName                = "/tran.v1.TranV1/SendTran"
	TranV1_Height_FullMethodName                  = "/tran.v1.TranV1/Height"
	TranV1_GetBlockHashByHeight_FullMethodName    = "/tran.v1.TranV1/GetBlockHashByHeight"
	TranV1_EthGetBlockHashByHeight_FullMethodName = "/tran.v1.TranV1/EthGetBlockHashByHeight"
	TranV1_BtcGetBlockHashByHeight_FullMethodName = "/tran.v1.TranV1/BtcGetBlockHashByHeight"
)

// TranV1Client is the client API for TranV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranV1Client interface {
	ChainList(ctx context.Context, in *ChainListRequest, opts ...grpc.CallOption) (*ChainListReply, error)
	IsMultiSigAddress(ctx context.Context, in *IsMultiSigAddressRequest, opts ...grpc.CallOption) (*IsMultiSigAddressReply, error)
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error)
	SendTran(ctx context.Context, in *SendTranRequest, opts ...grpc.CallOption) (*SendTranReply, error)
	Height(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*HeightReply, error)
	GetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...grpc.CallOption) (*GetBlockHashByHeightReply, error)
	EthGetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...grpc.CallOption) (*TxResult, error)
	BtcGetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...grpc.CallOption) (*BtcGetBlockHashByHeightReply, error)
}

type tranV1Client struct {
	cc grpc.ClientConnInterface
}

func NewTranV1Client(cc grpc.ClientConnInterface) TranV1Client {
	return &tranV1Client{cc}
}

func (c *tranV1Client) ChainList(ctx context.Context, in *ChainListRequest, opts ...grpc.CallOption) (*ChainListReply, error) {
	out := new(ChainListReply)
	err := c.cc.Invoke(ctx, TranV1_ChainList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tranV1Client) IsMultiSigAddress(ctx context.Context, in *IsMultiSigAddressRequest, opts ...grpc.CallOption) (*IsMultiSigAddressReply, error) {
	out := new(IsMultiSigAddressReply)
	err := c.cc.Invoke(ctx, TranV1_IsMultiSigAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tranV1Client) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error) {
	out := new(BalanceReply)
	err := c.cc.Invoke(ctx, TranV1_Balance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tranV1Client) SendTran(ctx context.Context, in *SendTranRequest, opts ...grpc.CallOption) (*SendTranReply, error) {
	out := new(SendTranReply)
	err := c.cc.Invoke(ctx, TranV1_SendTran_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tranV1Client) Height(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*HeightReply, error) {
	out := new(HeightReply)
	err := c.cc.Invoke(ctx, TranV1_Height_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tranV1Client) GetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...grpc.CallOption) (*GetBlockHashByHeightReply, error) {
	out := new(GetBlockHashByHeightReply)
	err := c.cc.Invoke(ctx, TranV1_GetBlockHashByHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tranV1Client) EthGetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, TranV1_EthGetBlockHashByHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tranV1Client) BtcGetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...grpc.CallOption) (*BtcGetBlockHashByHeightReply, error) {
	out := new(BtcGetBlockHashByHeightReply)
	err := c.cc.Invoke(ctx, TranV1_BtcGetBlockHashByHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranV1Server is the server API for TranV1 service.
// All implementations must embed UnimplementedTranV1Server
// for forward compatibility
type TranV1Server interface {
	ChainList(context.Context, *ChainListRequest) (*ChainListReply, error)
	IsMultiSigAddress(context.Context, *IsMultiSigAddressRequest) (*IsMultiSigAddressReply, error)
	Balance(context.Context, *BalanceRequest) (*BalanceReply, error)
	SendTran(context.Context, *SendTranRequest) (*SendTranReply, error)
	Height(context.Context, *HeightRequest) (*HeightReply, error)
	GetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*GetBlockHashByHeightReply, error)
	EthGetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*TxResult, error)
	BtcGetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*BtcGetBlockHashByHeightReply, error)
	mustEmbedUnimplementedTranV1Server()
}

// UnimplementedTranV1Server must be embedded to have forward compatible implementations.
type UnimplementedTranV1Server struct {
}

func (UnimplementedTranV1Server) ChainList(context.Context, *ChainListRequest) (*ChainListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainList not implemented")
}
func (UnimplementedTranV1Server) IsMultiSigAddress(context.Context, *IsMultiSigAddressRequest) (*IsMultiSigAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsMultiSigAddress not implemented")
}
func (UnimplementedTranV1Server) Balance(context.Context, *BalanceRequest) (*BalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedTranV1Server) SendTran(context.Context, *SendTranRequest) (*SendTranReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTran not implemented")
}
func (UnimplementedTranV1Server) Height(context.Context, *HeightRequest) (*HeightReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Height not implemented")
}
func (UnimplementedTranV1Server) GetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*GetBlockHashByHeightReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHashByHeight not implemented")
}
func (UnimplementedTranV1Server) EthGetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthGetBlockHashByHeight not implemented")
}
func (UnimplementedTranV1Server) BtcGetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*BtcGetBlockHashByHeightReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BtcGetBlockHashByHeight not implemented")
}
func (UnimplementedTranV1Server) mustEmbedUnimplementedTranV1Server() {}

// UnsafeTranV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranV1Server will
// result in compilation errors.
type UnsafeTranV1Server interface {
	mustEmbedUnimplementedTranV1Server()
}

func RegisterTranV1Server(s grpc.ServiceRegistrar, srv TranV1Server) {
	s.RegisterService(&TranV1_ServiceDesc, srv)
}

func _TranV1_ChainList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranV1Server).ChainList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranV1_ChainList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranV1Server).ChainList(ctx, req.(*ChainListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranV1_IsMultiSigAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsMultiSigAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranV1Server).IsMultiSigAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranV1_IsMultiSigAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranV1Server).IsMultiSigAddress(ctx, req.(*IsMultiSigAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranV1_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranV1Server).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranV1_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranV1Server).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranV1_SendTran_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTranRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranV1Server).SendTran(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranV1_SendTran_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranV1Server).SendTran(ctx, req.(*SendTranRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranV1_Height_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranV1Server).Height(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranV1_Height_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranV1Server).Height(ctx, req.(*HeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranV1_GetBlockHashByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHashByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranV1Server).GetBlockHashByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranV1_GetBlockHashByHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranV1Server).GetBlockHashByHeight(ctx, req.(*GetBlockHashByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranV1_EthGetBlockHashByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHashByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranV1Server).EthGetBlockHashByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranV1_EthGetBlockHashByHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranV1Server).EthGetBlockHashByHeight(ctx, req.(*GetBlockHashByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranV1_BtcGetBlockHashByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHashByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranV1Server).BtcGetBlockHashByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranV1_BtcGetBlockHashByHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranV1Server).BtcGetBlockHashByHeight(ctx, req.(*GetBlockHashByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TranV1_ServiceDesc is the grpc.ServiceDesc for TranV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tran.v1.TranV1",
	HandlerType: (*TranV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChainList",
			Handler:    _TranV1_ChainList_Handler,
		},
		{
			MethodName: "IsMultiSigAddress",
			Handler:    _TranV1_IsMultiSigAddress_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _TranV1_Balance_Handler,
		},
		{
			MethodName: "SendTran",
			Handler:    _TranV1_SendTran_Handler,
		},
		{
			MethodName: "Height",
			Handler:    _TranV1_Height_Handler,
		},
		{
			MethodName: "GetBlockHashByHeight",
			Handler:    _TranV1_GetBlockHashByHeight_Handler,
		},
		{
			MethodName: "EthGetBlockHashByHeight",
			Handler:    _TranV1_EthGetBlockHashByHeight_Handler,
		},
		{
			MethodName: "BtcGetBlockHashByHeight",
			Handler:    _TranV1_BtcGetBlockHashByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tran/v1/api.proto",
}
