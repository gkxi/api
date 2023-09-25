// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.12
// source: tran/v1/api.proto

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

const OperationTranV1Balance = "/tran.v1.TranV1/Balance"
const OperationTranV1ChainList = "/tran.v1.TranV1/ChainList"
const OperationTranV1GetBlockHashByHeight = "/tran.v1.TranV1/GetBlockHashByHeight"
const OperationTranV1GetBlockHashByHeightBtc = "/tran.v1.TranV1/GetBlockHashByHeightBtc"
const OperationTranV1GetBlockHashByHeightEth = "/tran.v1.TranV1/GetBlockHashByHeightEth"
const OperationTranV1Height = "/tran.v1.TranV1/Height"
const OperationTranV1IsMultiSigAddress = "/tran.v1.TranV1/IsMultiSigAddress"
const OperationTranV1SendTran = "/tran.v1.TranV1/SendTran"

type TranV1HTTPServer interface {
	Balance(context.Context, *BalanceRequest) (*BalanceReply, error)
	ChainList(context.Context, *ChainListRequest) (*ChainListReply, error)
	GetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*GetBlockHashByHeightReply, error)
	GetBlockHashByHeightBtc(context.Context, *GetBlockHashByHeightRequest) (*GetBlockHashByHeightBtcReply, error)
	GetBlockHashByHeightEth(context.Context, *GetBlockHashByHeightRequest) (*TxResult, error)
	Height(context.Context, *HeightRequest) (*HeightReply, error)
	IsMultiSigAddress(context.Context, *IsMultiSigAddressRequest) (*IsMultiSigAddressReply, error)
	SendTran(context.Context, *SendTranRequest) (*SendTranReply, error)
}

func RegisterTranV1HTTPServer(s *http.Server, srv TranV1HTTPServer) {
	r := s.Route("/")
	r.POST("/chain/list/{chainCode}", _TranV1_ChainList0_HTTP_Handler(srv))
	r.POST("/chain/multisig", _TranV1_IsMultiSigAddress0_HTTP_Handler(srv))
	r.POST("/chain/getBalance", _TranV1_Balance0_HTTP_Handler(srv))
	r.POST("/chain/transferTo", _TranV1_SendTran0_HTTP_Handler(srv))
	r.POST("/chain/getLastBlockHeight", _TranV1_Height0_HTTP_Handler(srv))
	r.POST("/chain/getBlockHashByHeight/{height}", _TranV1_GetBlockHashByHeight0_HTTP_Handler(srv))
	r.POST("/chain/getBlockHashByHeightEth/{height}", _TranV1_GetBlockHashByHeightEth0_HTTP_Handler(srv))
	r.POST("/chain/getBlockHashByHeightBtc/{height}", _TranV1_GetBlockHashByHeightBtc0_HTTP_Handler(srv))
}

func _TranV1_ChainList0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChainListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1ChainList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChainList(ctx, req.(*ChainListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ChainListReply)
		return ctx.Result(200, reply)
	}
}

func _TranV1_IsMultiSigAddress0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IsMultiSigAddressRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1IsMultiSigAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.IsMultiSigAddress(ctx, req.(*IsMultiSigAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IsMultiSigAddressReply)
		return ctx.Result(200, reply)
	}
}

func _TranV1_Balance0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BalanceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1Balance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Balance(ctx, req.(*BalanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BalanceReply)
		return ctx.Result(200, reply)
	}
}

func _TranV1_SendTran0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendTranRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1SendTran)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendTran(ctx, req.(*SendTranRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendTranReply)
		return ctx.Result(200, reply)
	}
}

func _TranV1_Height0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HeightRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1Height)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Height(ctx, req.(*HeightRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HeightReply)
		return ctx.Result(200, reply)
	}
}

func _TranV1_GetBlockHashByHeight0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBlockHashByHeightRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1GetBlockHashByHeight)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBlockHashByHeight(ctx, req.(*GetBlockHashByHeightRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBlockHashByHeightReply)
		return ctx.Result(200, reply)
	}
}

func _TranV1_GetBlockHashByHeightEth0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBlockHashByHeightRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1GetBlockHashByHeightEth)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBlockHashByHeightEth(ctx, req.(*GetBlockHashByHeightRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TxResult)
		return ctx.Result(200, reply)
	}
}

func _TranV1_GetBlockHashByHeightBtc0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBlockHashByHeightRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1GetBlockHashByHeightBtc)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBlockHashByHeightBtc(ctx, req.(*GetBlockHashByHeightRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBlockHashByHeightBtcReply)
		return ctx.Result(200, reply)
	}
}

type TranV1HTTPClient interface {
	Balance(ctx context.Context, req *BalanceRequest, opts ...http.CallOption) (rsp *BalanceReply, err error)
	ChainList(ctx context.Context, req *ChainListRequest, opts ...http.CallOption) (rsp *ChainListReply, err error)
	GetBlockHashByHeight(ctx context.Context, req *GetBlockHashByHeightRequest, opts ...http.CallOption) (rsp *GetBlockHashByHeightReply, err error)
	GetBlockHashByHeightBtc(ctx context.Context, req *GetBlockHashByHeightRequest, opts ...http.CallOption) (rsp *GetBlockHashByHeightBtcReply, err error)
	GetBlockHashByHeightEth(ctx context.Context, req *GetBlockHashByHeightRequest, opts ...http.CallOption) (rsp *TxResult, err error)
	Height(ctx context.Context, req *HeightRequest, opts ...http.CallOption) (rsp *HeightReply, err error)
	IsMultiSigAddress(ctx context.Context, req *IsMultiSigAddressRequest, opts ...http.CallOption) (rsp *IsMultiSigAddressReply, err error)
	SendTran(ctx context.Context, req *SendTranRequest, opts ...http.CallOption) (rsp *SendTranReply, err error)
}

type TranV1HTTPClientImpl struct {
	cc *http.Client
}

func NewTranV1HTTPClient(client *http.Client) TranV1HTTPClient {
	return &TranV1HTTPClientImpl{client}
}

func (c *TranV1HTTPClientImpl) Balance(ctx context.Context, in *BalanceRequest, opts ...http.CallOption) (*BalanceReply, error) {
	var out BalanceReply
	pattern := "/chain/getBalance"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1Balance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TranV1HTTPClientImpl) ChainList(ctx context.Context, in *ChainListRequest, opts ...http.CallOption) (*ChainListReply, error) {
	var out ChainListReply
	pattern := "/chain/list/{chainCode}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1ChainList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TranV1HTTPClientImpl) GetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...http.CallOption) (*GetBlockHashByHeightReply, error) {
	var out GetBlockHashByHeightReply
	pattern := "/chain/getBlockHashByHeight/{height}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1GetBlockHashByHeight))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TranV1HTTPClientImpl) GetBlockHashByHeightBtc(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...http.CallOption) (*GetBlockHashByHeightBtcReply, error) {
	var out GetBlockHashByHeightBtcReply
	pattern := "/chain/getBlockHashByHeightBtc/{height}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1GetBlockHashByHeightBtc))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TranV1HTTPClientImpl) GetBlockHashByHeightEth(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...http.CallOption) (*TxResult, error) {
	var out TxResult
	pattern := "/chain/getBlockHashByHeightEth/{height}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1GetBlockHashByHeightEth))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TranV1HTTPClientImpl) Height(ctx context.Context, in *HeightRequest, opts ...http.CallOption) (*HeightReply, error) {
	var out HeightReply
	pattern := "/chain/getLastBlockHeight"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1Height))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TranV1HTTPClientImpl) IsMultiSigAddress(ctx context.Context, in *IsMultiSigAddressRequest, opts ...http.CallOption) (*IsMultiSigAddressReply, error) {
	var out IsMultiSigAddressReply
	pattern := "/chain/multisig"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1IsMultiSigAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TranV1HTTPClientImpl) SendTran(ctx context.Context, in *SendTranRequest, opts ...http.CallOption) (*SendTranReply, error) {
	var out SendTranReply
	pattern := "/chain/transferTo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1SendTran))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
