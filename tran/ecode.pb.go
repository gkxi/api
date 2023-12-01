// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: tran/ecode.proto

package tran

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/go-kratos/kratos/v2/errors"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	// 为某个枚举单独设置错误码
	ErrorReason_USER_NOT_FOUND    ErrorReason = 0
	ErrorReason_CONTENT_MISSING   ErrorReason = 1
	ErrorReason_Invalid_Pub_ECDSA ErrorReason = 10000
	ErrorReason_TxOutNil          ErrorReason = 10001
	ErrorReason_InvalidScriptType ErrorReason = 10002
	ErrorReason_ErrMnemonic       ErrorReason = 10003
	ErrorReason_ErrPrivateKey     ErrorReason = 10004
	ErrorReason_InvalidBlock      ErrorReason = 10005
	ErrorReason_BalanceNotEnough  ErrorReason = 10006
	ErrorReason_TxFailed          ErrorReason = 10007
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:     "USER_NOT_FOUND",
		1:     "CONTENT_MISSING",
		10000: "Invalid_Pub_ECDSA",
		10001: "TxOutNil",
		10002: "InvalidScriptType",
		10003: "ErrMnemonic",
		10004: "ErrPrivateKey",
		10005: "InvalidBlock",
		10006: "BalanceNotEnough",
		10007: "TxFailed",
	}
	ErrorReason_value = map[string]int32{
		"USER_NOT_FOUND":    0,
		"CONTENT_MISSING":   1,
		"Invalid_Pub_ECDSA": 10000,
		"TxOutNil":          10001,
		"InvalidScriptType": 10002,
		"ErrMnemonic":       10003,
		"ErrPrivateKey":     10004,
		"InvalidBlock":      10005,
		"BalanceNotEnough":  10006,
		"TxFailed":          10007,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_tran_ecode_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_tran_ecode_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_tran_ecode_proto_rawDescGZIP(), []int{0}
}

var File_tran_ecode_proto protoreflect.FileDescriptor

var file_tran_ecode_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x72, 0x61, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe6, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12,
	0x19, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x50, 0x75, 0x62, 0x5f, 0x45, 0x43, 0x44, 0x53, 0x41, 0x10,
	0x90, 0x4e, 0x12, 0x0d, 0x0a, 0x08, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x4e, 0x69, 0x6c, 0x10, 0x91,
	0x4e, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x92, 0x4e, 0x12, 0x10, 0x0a, 0x0b, 0x45, 0x72, 0x72,
	0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x10, 0x93, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x45,
	0x72, 0x72, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x10, 0x94, 0x4e, 0x12,
	0x11, 0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x10,
	0x95, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x74,
	0x45, 0x6e, 0x6f, 0x75, 0x67, 0x68, 0x10, 0x96, 0x4e, 0x12, 0x0d, 0x0a, 0x08, 0x54, 0x78, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x97, 0x4e, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x1f,
	0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6b, 0x78,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tran_ecode_proto_rawDescOnce sync.Once
	file_tran_ecode_proto_rawDescData = file_tran_ecode_proto_rawDesc
)

func file_tran_ecode_proto_rawDescGZIP() []byte {
	file_tran_ecode_proto_rawDescOnce.Do(func() {
		file_tran_ecode_proto_rawDescData = protoimpl.X.CompressGZIP(file_tran_ecode_proto_rawDescData)
	})
	return file_tran_ecode_proto_rawDescData
}

var file_tran_ecode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tran_ecode_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: tran.ErrorReason
}
var file_tran_ecode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tran_ecode_proto_init() }
func file_tran_ecode_proto_init() {
	if File_tran_ecode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tran_ecode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tran_ecode_proto_goTypes,
		DependencyIndexes: file_tran_ecode_proto_depIdxs,
		EnumInfos:         file_tran_ecode_proto_enumTypes,
	}.Build()
	File_tran_ecode_proto = out.File
	file_tran_ecode_proto_rawDesc = nil
	file_tran_ecode_proto_goTypes = nil
	file_tran_ecode_proto_depIdxs = nil
}
