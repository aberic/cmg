// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: rwset.proto

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

// key read version
type KeyVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the transaction identifier that last modified the key
	RefTxId string `protobuf:"bytes,3,opt,name=ref_tx_id,json=refTxId,proto3" json:"ref_tx_id,omitempty"`
	// the offset of the key in the write set of the transaction, starts from 0
	RefOffset int32 `protobuf:"varint,4,opt,name=ref_offset,json=refOffset,proto3" json:"ref_offset,omitempty"`
}

func (x *KeyVersion) Reset() {
	*x = KeyVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rwset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyVersion) ProtoMessage() {}

func (x *KeyVersion) ProtoReflect() protoreflect.Message {
	mi := &file_rwset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyVersion.ProtoReflect.Descriptor instead.
func (*KeyVersion) Descriptor() ([]byte, []int) {
	return file_rwset_proto_rawDescGZIP(), []int{0}
}

func (x *KeyVersion) GetRefTxId() string {
	if x != nil {
		return x.RefTxId
	}
	return ""
}

func (x *KeyVersion) GetRefOffset() int32 {
	if x != nil {
		return x.RefOffset
	}
	return 0
}

// TxRead describes a read operation on a key
type TxRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// read key
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// the value of the key
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// contract name, used in cross-contract calls
	// set to null if only the contract in transaction request is called
	ContractName string `protobuf:"bytes,3,opt,name=contract_name,json=contractName,proto3" json:"contract_name,omitempty"`
	// read key version
	Version *KeyVersion `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *TxRead) Reset() {
	*x = TxRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rwset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxRead) ProtoMessage() {}

func (x *TxRead) ProtoReflect() protoreflect.Message {
	mi := &file_rwset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxRead.ProtoReflect.Descriptor instead.
func (*TxRead) Descriptor() ([]byte, []int) {
	return file_rwset_proto_rawDescGZIP(), []int{1}
}

func (x *TxRead) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *TxRead) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *TxRead) GetContractName() string {
	if x != nil {
		return x.ContractName
	}
	return ""
}

func (x *TxRead) GetVersion() *KeyVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

// TxRead describes a write/delete operation on a key
type TxWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// write key
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// write value
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// contract name, used in cross-contract calls
	// set to null if only the contract in transaction request is called
	ContractName string `protobuf:"bytes,3,opt,name=contract_name,json=contractName,proto3" json:"contract_name,omitempty"`
}

func (x *TxWrite) Reset() {
	*x = TxWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rwset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxWrite) ProtoMessage() {}

func (x *TxWrite) ProtoReflect() protoreflect.Message {
	mi := &file_rwset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxWrite.ProtoReflect.Descriptor instead.
func (*TxWrite) Descriptor() ([]byte, []int) {
	return file_rwset_proto_rawDescGZIP(), []int{2}
}

func (x *TxWrite) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *TxWrite) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *TxWrite) GetContractName() string {
	if x != nil {
		return x.ContractName
	}
	return ""
}

// TxRWSet describes all the operations of a transaction on ledger
type TxRWSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// transaction identifier
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// read set
	TxReads []*TxRead `protobuf:"bytes,2,rep,name=tx_reads,json=txReads,proto3" json:"tx_reads,omitempty"`
	// write set
	TxWrites []*TxWrite `protobuf:"bytes,3,rep,name=tx_writes,json=txWrites,proto3" json:"tx_writes,omitempty"`
}

func (x *TxRWSet) Reset() {
	*x = TxRWSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rwset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxRWSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxRWSet) ProtoMessage() {}

func (x *TxRWSet) ProtoReflect() protoreflect.Message {
	mi := &file_rwset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxRWSet.ProtoReflect.Descriptor instead.
func (*TxRWSet) Descriptor() ([]byte, []int) {
	return file_rwset_proto_rawDescGZIP(), []int{3}
}

func (x *TxRWSet) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *TxRWSet) GetTxReads() []*TxRead {
	if x != nil {
		return x.TxReads
	}
	return nil
}

func (x *TxRWSet) GetTxWrites() []*TxWrite {
	if x != nil {
		return x.TxWrites
	}
	return nil
}

var File_rwset_proto protoreflect.FileDescriptor

var file_rwset_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x47, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x5f, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x54, 0x78, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x66, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x65, 0x66, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x7f, 0x0a, 0x06, 0x54, 0x78,
	0x52, 0x65, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x07, 0x54,
	0x78, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x07, 0x54, 0x78, 0x52, 0x57, 0x53, 0x65, 0x74, 0x12, 0x13,
	0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x78, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x07, 0x74, 0x78, 0x52, 0x65, 0x61, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x09, 0x74, 0x78,
	0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x54, 0x78, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x08, 0x74, 0x78, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x42, 0x3b, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5a, 0x1f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rwset_proto_rawDescOnce sync.Once
	file_rwset_proto_rawDescData = file_rwset_proto_rawDesc
)

func file_rwset_proto_rawDescGZIP() []byte {
	file_rwset_proto_rawDescOnce.Do(func() {
		file_rwset_proto_rawDescData = protoimpl.X.CompressGZIP(file_rwset_proto_rawDescData)
	})
	return file_rwset_proto_rawDescData
}

var file_rwset_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rwset_proto_goTypes = []interface{}{
	(*KeyVersion)(nil), // 0: pb.KeyVersion
	(*TxRead)(nil),     // 1: pb.TxRead
	(*TxWrite)(nil),    // 2: pb.TxWrite
	(*TxRWSet)(nil),    // 3: pb.TxRWSet
}
var file_rwset_proto_depIdxs = []int32{
	0, // 0: pb.TxRead.version:type_name -> pb.KeyVersion
	1, // 1: pb.TxRWSet.tx_reads:type_name -> pb.TxRead
	2, // 2: pb.TxRWSet.tx_writes:type_name -> pb.TxWrite
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rwset_proto_init() }
func file_rwset_proto_init() {
	if File_rwset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rwset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyVersion); i {
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
		file_rwset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxRead); i {
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
		file_rwset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxWrite); i {
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
		file_rwset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxRWSet); i {
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
			RawDescriptor: file_rwset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rwset_proto_goTypes,
		DependencyIndexes: file_rwset_proto_depIdxs,
		MessageInfos:      file_rwset_proto_msgTypes,
	}.Build()
	File_rwset_proto = out.File
	file_rwset_proto_rawDesc = nil
	file_rwset_proto_goTypes = nil
	file_rwset_proto_depIdxs = nil
}
