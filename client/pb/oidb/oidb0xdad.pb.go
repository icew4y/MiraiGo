// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: oidb0xdad.proto

package oidb

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

type DADReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client    int64        `protobuf:"varint,1,opt,name=client,proto3" json:"client,omitempty"`
	ProductId uint64       `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Amount    int64        `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	ToUin     uint64       `protobuf:"varint,4,opt,name=to_uin,json=toUin,proto3" json:"to_uin,omitempty"`
	Gc        uint64       `protobuf:"varint,5,opt,name=gc,proto3" json:"gc,omitempty"`
	Ip        string       `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
	Version   string       `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	Sig       *DADLoginSig `protobuf:"bytes,8,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *DADReqBody) Reset() {
	*x = DADReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xdad_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DADReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DADReqBody) ProtoMessage() {}

func (x *DADReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xdad_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DADReqBody.ProtoReflect.Descriptor instead.
func (*DADReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0xdad_proto_rawDescGZIP(), []int{0}
}

func (x *DADReqBody) GetClient() int64 {
	if x != nil {
		return x.Client
	}
	return 0
}

func (x *DADReqBody) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *DADReqBody) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DADReqBody) GetToUin() uint64 {
	if x != nil {
		return x.ToUin
	}
	return 0
}

func (x *DADReqBody) GetGc() uint64 {
	if x != nil {
		return x.Gc
	}
	return 0
}

func (x *DADReqBody) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *DADReqBody) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DADReqBody) GetSig() *DADLoginSig {
	if x != nil {
		return x.Sig
	}
	return nil
}

type DADLoginSig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Sig   []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	Appid uint32 `protobuf:"varint,3,opt,name=appid,proto3" json:"appid,omitempty"`
}

func (x *DADLoginSig) Reset() {
	*x = DADLoginSig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xdad_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DADLoginSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DADLoginSig) ProtoMessage() {}

func (x *DADLoginSig) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xdad_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DADLoginSig.ProtoReflect.Descriptor instead.
func (*DADLoginSig) Descriptor() ([]byte, []int) {
	return file_oidb0xdad_proto_rawDescGZIP(), []int{1}
}

func (x *DADLoginSig) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *DADLoginSig) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *DADLoginSig) GetAppid() uint32 {
	if x != nil {
		return x.Appid
	}
	return 0
}

var File_oidb0xdad_proto protoreflect.FileDescriptor

var file_oidb0xdad_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x64, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x0a, 0x44, 0x41, 0x44, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x75, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x74, 0x6f, 0x55, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x67, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x67, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x44, 0x41, 0x44, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x52, 0x03, 0x73, 0x69, 0x67,
	0x22, 0x49, 0x0a, 0x0b, 0x44, 0x41, 0x44, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x3b, 0x6f, 0x69, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oidb0xdad_proto_rawDescOnce sync.Once
	file_oidb0xdad_proto_rawDescData = file_oidb0xdad_proto_rawDesc
)

func file_oidb0xdad_proto_rawDescGZIP() []byte {
	file_oidb0xdad_proto_rawDescOnce.Do(func() {
		file_oidb0xdad_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0xdad_proto_rawDescData)
	})
	return file_oidb0xdad_proto_rawDescData
}

var file_oidb0xdad_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oidb0xdad_proto_goTypes = []interface{}{
	(*DADReqBody)(nil),  // 0: DADReqBody
	(*DADLoginSig)(nil), // 1: DADLoginSig
}
var file_oidb0xdad_proto_depIdxs = []int32{
	1, // 0: DADReqBody.sig:type_name -> DADLoginSig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_oidb0xdad_proto_init() }
func file_oidb0xdad_proto_init() {
	if File_oidb0xdad_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0xdad_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DADReqBody); i {
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
		file_oidb0xdad_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DADLoginSig); i {
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
			RawDescriptor: file_oidb0xdad_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0xdad_proto_goTypes,
		DependencyIndexes: file_oidb0xdad_proto_depIdxs,
		MessageInfos:      file_oidb0xdad_proto_msgTypes,
	}.Build()
	File_oidb0xdad_proto = out.File
	file_oidb0xdad_proto_rawDesc = nil
	file_oidb0xdad_proto_goTypes = nil
	file_oidb0xdad_proto_depIdxs = nil
}
