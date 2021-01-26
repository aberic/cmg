// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: local_config.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DebugConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*KeyValuePair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *DebugConfigRequest) Reset() {
	*x = DebugConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugConfigRequest) ProtoMessage() {}

func (x *DebugConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_local_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugConfigRequest.ProtoReflect.Descriptor instead.
func (*DebugConfigRequest) Descriptor() ([]byte, []int) {
	return file_local_config_proto_rawDescGZIP(), []int{0}
}

func (x *DebugConfigRequest) GetPairs() []*KeyValuePair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type DebugConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DebugConfigResponse) Reset() {
	*x = DebugConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugConfigResponse) ProtoMessage() {}

func (x *DebugConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_local_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugConfigResponse.ProtoReflect.Descriptor instead.
func (*DebugConfigResponse) Descriptor() ([]byte, []int) {
	return file_local_config_proto_rawDescGZIP(), []int{1}
}

func (x *DebugConfigResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DebugConfigResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CheckNewBlockChainConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckNewBlockChainConfigRequest) Reset() {
	*x = CheckNewBlockChainConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckNewBlockChainConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckNewBlockChainConfigRequest) ProtoMessage() {}

func (x *CheckNewBlockChainConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_local_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckNewBlockChainConfigRequest.ProtoReflect.Descriptor instead.
func (*CheckNewBlockChainConfigRequest) Descriptor() ([]byte, []int) {
	return file_local_config_proto_rawDescGZIP(), []int{2}
}

type CheckNewBlockChainConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CheckNewBlockChainConfigResponse) Reset() {
	*x = CheckNewBlockChainConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckNewBlockChainConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckNewBlockChainConfigResponse) ProtoMessage() {}

func (x *CheckNewBlockChainConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_local_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckNewBlockChainConfigResponse.ProtoReflect.Descriptor instead.
func (*CheckNewBlockChainConfigResponse) Descriptor() ([]byte, []int) {
	return file_local_config_proto_rawDescGZIP(), []int{3}
}

func (x *CheckNewBlockChainConfigResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CheckNewBlockChainConfigResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_local_config_proto protoreflect.FileDescriptor

var file_local_config_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x12, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x69, 0x72, 0x73, 0x22, 0x43, 0x0a, 0x13, 0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x1f, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a,
	0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x21, 0x5a, 0x1f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_local_config_proto_rawDescOnce sync.Once
	file_local_config_proto_rawDescData = file_local_config_proto_rawDesc
)

func file_local_config_proto_rawDescGZIP() []byte {
	file_local_config_proto_rawDescOnce.Do(func() {
		file_local_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_local_config_proto_rawDescData)
	})
	return file_local_config_proto_rawDescData
}

var file_local_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_local_config_proto_goTypes = []interface{}{
	(*DebugConfigRequest)(nil),               // 0: pb.DebugConfigRequest
	(*DebugConfigResponse)(nil),              // 1: pb.DebugConfigResponse
	(*CheckNewBlockChainConfigRequest)(nil),  // 2: pb.CheckNewBlockChainConfigRequest
	(*CheckNewBlockChainConfigResponse)(nil), // 3: pb.CheckNewBlockChainConfigResponse
	(*KeyValuePair)(nil),                     // 4: pb.KeyValuePair
}
var file_local_config_proto_depIdxs = []int32{
	4, // 0: pb.DebugConfigRequest.pairs:type_name -> pb.KeyValuePair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_local_config_proto_init() }
func file_local_config_proto_init() {
	if File_local_config_proto != nil {
		return
	}
	file_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_local_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugConfigRequest); i {
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
		file_local_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugConfigResponse); i {
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
		file_local_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckNewBlockChainConfigRequest); i {
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
		file_local_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckNewBlockChainConfigResponse); i {
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
			RawDescriptor: file_local_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_local_config_proto_goTypes,
		DependencyIndexes: file_local_config_proto_depIdxs,
		MessageInfos:      file_local_config_proto_msgTypes,
	}.Build()
	File_local_config_proto = out.File
	file_local_config_proto_rawDesc = nil
	file_local_config_proto_goTypes = nil
	file_local_config_proto_depIdxs = nil
}
