// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.12
// source: address/v1/api.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAddressV1NewBip44 = "/address.v1.AddressV1/NewBip44"
const OperationAddressV1NewBip441 = "/address.v1.AddressV1/NewBip441"

type AddressV1HTTPServer interface {
	NewBip44(context.Context, *NewBip44Request) (*NewBip44Reply, error)
	NewBip441(context.Context, *NewBip44Request) (*NewBip441Reply, error)
}

func RegisterAddressV1HTTPServer(s *http.Server, srv AddressV1HTTPServer) {
	r := s.Route("/")
	r.POST("/chain/createAddressByMnemonic", _AddressV1_NewBip440_HTTP_Handler(srv))
	r.POST("/chain/createAddressByMnemonic1", _AddressV1_NewBip4410_HTTP_Handler(srv))
}

func _AddressV1_NewBip440_HTTP_Handler(srv AddressV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NewBip44Request
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAddressV1NewBip44)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NewBip44(ctx, req.(*NewBip44Request))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NewBip44Reply)
		return ctx.Result(200, reply)
	}
}

func _AddressV1_NewBip4410_HTTP_Handler(srv AddressV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NewBip44Request
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAddressV1NewBip441)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NewBip441(ctx, req.(*NewBip44Request))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NewBip441Reply)
		return ctx.Result(200, reply)
	}
}

type AddressV1HTTPClient interface {
	NewBip44(ctx context.Context, req *NewBip44Request, opts ...http.CallOption) (rsp *NewBip44Reply, err error)
	NewBip441(ctx context.Context, req *NewBip44Request, opts ...http.CallOption) (rsp *NewBip441Reply, err error)
}

type AddressV1HTTPClientImpl struct {
	cc *http.Client
}

func NewAddressV1HTTPClient(client *http.Client) AddressV1HTTPClient {
	return &AddressV1HTTPClientImpl{client}
}

func (c *AddressV1HTTPClientImpl) NewBip44(ctx context.Context, in *NewBip44Request, opts ...http.CallOption) (*NewBip44Reply, error) {
	var out NewBip44Reply
	pattern := "/chain/createAddressByMnemonic"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAddressV1NewBip44))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AddressV1HTTPClientImpl) NewBip441(ctx context.Context, in *NewBip44Request, opts ...http.CallOption) (*NewBip441Reply, error) {
	var out NewBip441Reply
	pattern := "/chain/createAddressByMnemonic1"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAddressV1NewBip441))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
