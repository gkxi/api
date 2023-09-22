// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: address/ecode.proto

package address

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
	ErrorReason_USER_NOT_FOUND  ErrorReason = 0
	ErrorReason_CONTENT_MISSING ErrorReason = 1
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "USER_NOT_FOUND",
		1: "CONTENT_MISSING",
	}
	ErrorReason_value = map[string]int32{
		"USER_NOT_FOUND":  0,
		"CONTENT_MISSING": 1,
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
	return file_address_ecode_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_address_ecode_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_address_ecode_proto_rawDescGZIP(), []int{0}
}

var File_address_ecode_proto protoreflect.FileDescriptor

var file_address_ecode_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x48, 0x0a, 0x0b, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x1a, 0x04,
	0xa8, 0x45, 0x94, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x1a,
	0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x2e, 0x67, 0x6b, 0x78,
	0x69, 0x2e, 0x74, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x6b, 0x2f, 0x61, 0x70, 0x6b, 0x2d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x3b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_ecode_proto_rawDescOnce sync.Once
	file_address_ecode_proto_rawDescData = file_address_ecode_proto_rawDesc
)

func file_address_ecode_proto_rawDescGZIP() []byte {
	file_address_ecode_proto_rawDescOnce.Do(func() {
		file_address_ecode_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_ecode_proto_rawDescData)
	})
	return file_address_ecode_proto_rawDescData
}

var file_address_ecode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_address_ecode_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: address.ErrorReason
}
var file_address_ecode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_address_ecode_proto_init() }
func file_address_ecode_proto_init() {
	if File_address_ecode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_address_ecode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_address_ecode_proto_goTypes,
		DependencyIndexes: file_address_ecode_proto_depIdxs,
		EnumInfos:         file_address_ecode_proto_enumTypes,
	}.Build()
	File_address_ecode_proto = out.File
	file_address_ecode_proto_rawDesc = nil
	file_address_ecode_proto_goTypes = nil
	file_address_ecode_proto_depIdxs = nil
}
