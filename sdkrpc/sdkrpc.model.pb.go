// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: sdkrpc.model.proto

package sdkrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 状态码
type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 消息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 国际化key
	I18N string `protobuf:"bytes,3,opt,name=i18n,proto3" json:"i18n,omitempty"`
	// 国际化扩展值
	Meta map[string]string `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdkrpc_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_sdkrpc_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_sdkrpc_model_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Code) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Code) GetI18N() string {
	if x != nil {
		return x.I18N
	}
	return ""
}

func (x *Code) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

// 返回结果
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码
	Code *Code `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// 数据
	Data *anypb.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdkrpc_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_sdkrpc_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_sdkrpc_model_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetCode() *Code {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *Result) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_sdkrpc_model_proto protoreflect.FileDescriptor

var file_sdkrpc_model_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x64, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x64, 0x6b, 0x72, 0x70, 0x63, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x31, 0x38, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x31, 0x38, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x64, 0x6b, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x54, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x64, 0x6b, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x3d, 0x0a, 0x0b, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x1a, 0x0e, 0x2e, 0x73, 0x64, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x72, 0x66, 0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x66, 0x69, 0x73, 0x68, 0x2d, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x73, 0x64, 0x6b, 0x72, 0x70, 0x63, 0x3b, 0x73, 0x64, 0x6b, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdkrpc_model_proto_rawDescOnce sync.Once
	file_sdkrpc_model_proto_rawDescData = file_sdkrpc_model_proto_rawDesc
)

func file_sdkrpc_model_proto_rawDescGZIP() []byte {
	file_sdkrpc_model_proto_rawDescOnce.Do(func() {
		file_sdkrpc_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdkrpc_model_proto_rawDescData)
	})
	return file_sdkrpc_model_proto_rawDescData
}

var file_sdkrpc_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sdkrpc_model_proto_goTypes = []interface{}{
	(*Code)(nil),      // 0: sdkrpc.Code
	(*Result)(nil),    // 1: sdkrpc.Result
	nil,               // 2: sdkrpc.Code.MetaEntry
	(*anypb.Any)(nil), // 3: google.protobuf.Any
}
var file_sdkrpc_model_proto_depIdxs = []int32{
	2, // 0: sdkrpc.Code.meta:type_name -> sdkrpc.Code.MetaEntry
	0, // 1: sdkrpc.Result.code:type_name -> sdkrpc.Code
	3, // 2: sdkrpc.Result.data:type_name -> google.protobuf.Any
	3, // 3: sdkrpc.GRPCService.Call:input_type -> google.protobuf.Any
	1, // 4: sdkrpc.GRPCService.Call:output_type -> sdkrpc.Result
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sdkrpc_model_proto_init() }
func file_sdkrpc_model_proto_init() {
	if File_sdkrpc_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdkrpc_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sdkrpc_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sdkrpc_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sdkrpc_model_proto_goTypes,
		DependencyIndexes: file_sdkrpc_model_proto_depIdxs,
		MessageInfos:      file_sdkrpc_model_proto_msgTypes,
	}.Build()
	File_sdkrpc_model_proto = out.File
	file_sdkrpc_model_proto_rawDesc = nil
	file_sdkrpc_model_proto_goTypes = nil
	file_sdkrpc_model_proto_depIdxs = nil
}
