// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v5.27.1
// source: tran/v1/tran_v1.proto

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
const OperationTranV1BtcGetBlockHashByHeight = "/tran.v1.TranV1/BtcGetBlockHashByHeight"
const OperationTranV1ChainList = "/tran.v1.TranV1/ChainList"
const OperationTranV1EthGetBlockHashByHeight = "/tran.v1.TranV1/EthGetBlockHashByHeight"
const OperationTranV1GetBlockHashByHeight = "/tran.v1.TranV1/GetBlockHashByHeight"
const OperationTranV1GetTxByHash = "/tran.v1.TranV1/GetTxByHash"
const OperationTranV1Height = "/tran.v1.TranV1/Height"
const OperationTranV1IsMultiSigAddress = "/tran.v1.TranV1/IsMultiSigAddress"
const OperationTranV1MinerFee = "/tran.v1.TranV1/MinerFee"
const OperationTranV1MinerFee1 = "/tran.v1.TranV1/MinerFee1"
const OperationTranV1SendTran = "/tran.v1.TranV1/SendTran"

type TranV1HTTPServer interface {
	Balance(context.Context, *BalanceRequest) (*BalanceReply, error)
	BtcGetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*BtcGetBlockHashByHeightReply, error)
	ChainList(context.Context, *ChainListRequest) (*ChainListReply, error)
	EthGetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*TxResult, error)
	GetBlockHashByHeight(context.Context, *GetBlockHashByHeightRequest) (*GetBlockHashByHeightReply, error)
	GetTxByHash(context.Context, *GetTxByHashRequest) (*GetTxByHashReply, error)
	Height(context.Context, *HeightRequest) (*HeightReply, error)
	IsMultiSigAddress(context.Context, *IsMultiSigAddressRequest) (*IsMultiSigAddressReply, error)
	MinerFee(context.Context, *MinerFeeRequest) (*MinerFeeReply, error)
	MinerFee1(context.Context, *MinerFee1Request) (*MinerFee1Reply, error)
	SendTran(context.Context, *SendTranRequest) (*SendTranReply, error)
}

func RegisterTranV1HTTPServer(s *http.Server, srv TranV1HTTPServer) {
	r := s.Route("/")
	r.POST("/chain/list/{chainCode}", _TranV1_ChainList0_HTTP_Handler(srv))
	r.POST("/chain/multisig", _TranV1_IsMultiSigAddress0_HTTP_Handler(srv))
	r.POST("/chain/getBalance", _TranV1_Balance0_HTTP_Handler(srv))
	r.POST("/chain/free/{master}/{gasAdd}", _TranV1_MinerFee10_HTTP_Handler(srv))
	r.POST("/chain/minerfee", _TranV1_MinerFee0_HTTP_Handler(srv))
	r.POST("/chain/transferTo", _TranV1_SendTran0_HTTP_Handler(srv))
	r.POST("/chain/getLastBlockHeight", _TranV1_Height0_HTTP_Handler(srv))
	r.POST("/chain/getBlockHashByHeight/{height}", _TranV1_GetBlockHashByHeight0_HTTP_Handler(srv))
	r.POST("/chain/tx/{hash}", _TranV1_GetTxByHash0_HTTP_Handler(srv))
	r.POST("/chain/getBlockHashByHeightEth/{height}", _TranV1_EthGetBlockHashByHeight0_HTTP_Handler(srv))
	r.POST("/chain/getBlockHashByHeightBtc/{height}", _TranV1_BtcGetBlockHashByHeight0_HTTP_Handler(srv))
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

func _TranV1_MinerFee10_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MinerFee1Request
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1MinerFee1)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MinerFee1(ctx, req.(*MinerFee1Request))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MinerFee1Reply)
		return ctx.Result(200, reply)
	}
}

func _TranV1_MinerFee0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MinerFeeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1MinerFee)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MinerFee(ctx, req.(*MinerFeeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MinerFeeReply)
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

func _TranV1_GetTxByHash0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTxByHashRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTranV1GetTxByHash)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTxByHash(ctx, req.(*GetTxByHashRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTxByHashReply)
		return ctx.Result(200, reply)
	}
}

func _TranV1_EthGetBlockHashByHeight0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
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
		http.SetOperation(ctx, OperationTranV1EthGetBlockHashByHeight)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EthGetBlockHashByHeight(ctx, req.(*GetBlockHashByHeightRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TxResult)
		return ctx.Result(200, reply)
	}
}

func _TranV1_BtcGetBlockHashByHeight0_HTTP_Handler(srv TranV1HTTPServer) func(ctx http.Context) error {
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
		http.SetOperation(ctx, OperationTranV1BtcGetBlockHashByHeight)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BtcGetBlockHashByHeight(ctx, req.(*GetBlockHashByHeightRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BtcGetBlockHashByHeightReply)
		return ctx.Result(200, reply)
	}
}

type TranV1HTTPClient interface {
	Balance(ctx context.Context, req *BalanceRequest, opts ...http.CallOption) (rsp *BalanceReply, err error)
	BtcGetBlockHashByHeight(ctx context.Context, req *GetBlockHashByHeightRequest, opts ...http.CallOption) (rsp *BtcGetBlockHashByHeightReply, err error)
	ChainList(ctx context.Context, req *ChainListRequest, opts ...http.CallOption) (rsp *ChainListReply, err error)
	EthGetBlockHashByHeight(ctx context.Context, req *GetBlockHashByHeightRequest, opts ...http.CallOption) (rsp *TxResult, err error)
	GetBlockHashByHeight(ctx context.Context, req *GetBlockHashByHeightRequest, opts ...http.CallOption) (rsp *GetBlockHashByHeightReply, err error)
	GetTxByHash(ctx context.Context, req *GetTxByHashRequest, opts ...http.CallOption) (rsp *GetTxByHashReply, err error)
	Height(ctx context.Context, req *HeightRequest, opts ...http.CallOption) (rsp *HeightReply, err error)
	IsMultiSigAddress(ctx context.Context, req *IsMultiSigAddressRequest, opts ...http.CallOption) (rsp *IsMultiSigAddressReply, err error)
	MinerFee(ctx context.Context, req *MinerFeeRequest, opts ...http.CallOption) (rsp *MinerFeeReply, err error)
	MinerFee1(ctx context.Context, req *MinerFee1Request, opts ...http.CallOption) (rsp *MinerFee1Reply, err error)
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

func (c *TranV1HTTPClientImpl) BtcGetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...http.CallOption) (*BtcGetBlockHashByHeightReply, error) {
	var out BtcGetBlockHashByHeightReply
	pattern := "/chain/getBlockHashByHeightBtc/{height}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1BtcGetBlockHashByHeight))
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

func (c *TranV1HTTPClientImpl) EthGetBlockHashByHeight(ctx context.Context, in *GetBlockHashByHeightRequest, opts ...http.CallOption) (*TxResult, error) {
	var out TxResult
	pattern := "/chain/getBlockHashByHeightEth/{height}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1EthGetBlockHashByHeight))
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

func (c *TranV1HTTPClientImpl) GetTxByHash(ctx context.Context, in *GetTxByHashRequest, opts ...http.CallOption) (*GetTxByHashReply, error) {
	var out GetTxByHashReply
	pattern := "/chain/tx/{hash}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1GetTxByHash))
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

func (c *TranV1HTTPClientImpl) MinerFee(ctx context.Context, in *MinerFeeRequest, opts ...http.CallOption) (*MinerFeeReply, error) {
	var out MinerFeeReply
	pattern := "/chain/minerfee"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1MinerFee))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TranV1HTTPClientImpl) MinerFee1(ctx context.Context, in *MinerFee1Request, opts ...http.CallOption) (*MinerFee1Reply, error) {
	var out MinerFee1Reply
	pattern := "/chain/free/{master}/{gasAdd}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTranV1MinerFee1))
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