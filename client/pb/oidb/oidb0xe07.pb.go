// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: oidb0xe07.proto

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

type DE07ReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    int32       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Client     int32       `protobuf:"varint,2,opt,name=client,proto3" json:"client,omitempty"`
	Entrance   int32       `protobuf:"varint,3,opt,name=entrance,proto3" json:"entrance,omitempty"`
	OcrReqBody *OCRReqBody `protobuf:"bytes,10,opt,name=ocrReqBody,proto3" json:"ocrReqBody,omitempty"`
}

func (x *DE07ReqBody) Reset() {
	*x = DE07ReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xe07_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DE07ReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DE07ReqBody) ProtoMessage() {}

func (x *DE07ReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xe07_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DE07ReqBody.ProtoReflect.Descriptor instead.
func (*DE07ReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0xe07_proto_rawDescGZIP(), []int{0}
}

func (x *DE07ReqBody) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DE07ReqBody) GetClient() int32 {
	if x != nil {
		return x.Client
	}
	return 0
}

func (x *DE07ReqBody) GetEntrance() int32 {
	if x != nil {
		return x.Entrance
	}
	return 0
}

func (x *DE07ReqBody) GetOcrReqBody() *OCRReqBody {
	if x != nil {
		return x.OcrReqBody
	}
	return nil
}

type OCRReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl              string `protobuf:"bytes,1,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	LanguageType          string `protobuf:"bytes,2,opt,name=languageType,proto3" json:"languageType,omitempty"`
	Scene                 string `protobuf:"bytes,3,opt,name=scene,proto3" json:"scene,omitempty"`
	OriginMd5             string `protobuf:"bytes,10,opt,name=originMd5,proto3" json:"originMd5,omitempty"`
	AfterCompressMd5      string `protobuf:"bytes,11,opt,name=afterCompressMd5,proto3" json:"afterCompressMd5,omitempty"`
	AfterCompressFileSize int32  `protobuf:"varint,12,opt,name=afterCompressFileSize,proto3" json:"afterCompressFileSize,omitempty"`
	AfterCompressWeight   int32  `protobuf:"varint,13,opt,name=afterCompressWeight,proto3" json:"afterCompressWeight,omitempty"`
	AfterCompressHeight   int32  `protobuf:"varint,14,opt,name=afterCompressHeight,proto3" json:"afterCompressHeight,omitempty"`
	IsCut                 bool   `protobuf:"varint,15,opt,name=isCut,proto3" json:"isCut,omitempty"`
}

func (x *OCRReqBody) Reset() {
	*x = OCRReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xe07_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCRReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCRReqBody) ProtoMessage() {}

func (x *OCRReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xe07_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCRReqBody.ProtoReflect.Descriptor instead.
func (*OCRReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0xe07_proto_rawDescGZIP(), []int{1}
}

func (x *OCRReqBody) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *OCRReqBody) GetLanguageType() string {
	if x != nil {
		return x.LanguageType
	}
	return ""
}

func (x *OCRReqBody) GetScene() string {
	if x != nil {
		return x.Scene
	}
	return ""
}

func (x *OCRReqBody) GetOriginMd5() string {
	if x != nil {
		return x.OriginMd5
	}
	return ""
}

func (x *OCRReqBody) GetAfterCompressMd5() string {
	if x != nil {
		return x.AfterCompressMd5
	}
	return ""
}

func (x *OCRReqBody) GetAfterCompressFileSize() int32 {
	if x != nil {
		return x.AfterCompressFileSize
	}
	return 0
}

func (x *OCRReqBody) GetAfterCompressWeight() int32 {
	if x != nil {
		return x.AfterCompressWeight
	}
	return 0
}

func (x *OCRReqBody) GetAfterCompressHeight() int32 {
	if x != nil {
		return x.AfterCompressHeight
	}
	return 0
}

func (x *OCRReqBody) GetIsCut() bool {
	if x != nil {
		return x.IsCut
	}
	return false
}

type DE07RspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode    int32       `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	ErrMsg     string      `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	Wording    string      `protobuf:"bytes,3,opt,name=wording,proto3" json:"wording,omitempty"`
	OcrRspBody *OCRRspBody `protobuf:"bytes,10,opt,name=ocrRspBody,proto3" json:"ocrRspBody,omitempty"`
}

func (x *DE07RspBody) Reset() {
	*x = DE07RspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xe07_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DE07RspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DE07RspBody) ProtoMessage() {}

func (x *DE07RspBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xe07_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DE07RspBody.ProtoReflect.Descriptor instead.
func (*DE07RspBody) Descriptor() ([]byte, []int) {
	return file_oidb0xe07_proto_rawDescGZIP(), []int{2}
}

func (x *DE07RspBody) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *DE07RspBody) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *DE07RspBody) GetWording() string {
	if x != nil {
		return x.Wording
	}
	return ""
}

func (x *DE07RspBody) GetOcrRspBody() *OCRRspBody {
	if x != nil {
		return x.OcrRspBody
	}
	return nil
}

type TextDetection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetectedText string   `protobuf:"bytes,1,opt,name=detectedText,proto3" json:"detectedText,omitempty"`
	Confidence   int32    `protobuf:"varint,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Polygon      *Polygon `protobuf:"bytes,3,opt,name=polygon,proto3" json:"polygon,omitempty"`
	AdvancedInfo string   `protobuf:"bytes,4,opt,name=advancedInfo,proto3" json:"advancedInfo,omitempty"`
}

func (x *TextDetection) Reset() {
	*x = TextDetection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xe07_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextDetection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextDetection) ProtoMessage() {}

func (x *TextDetection) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xe07_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextDetection.ProtoReflect.Descriptor instead.
func (*TextDetection) Descriptor() ([]byte, []int) {
	return file_oidb0xe07_proto_rawDescGZIP(), []int{3}
}

func (x *TextDetection) GetDetectedText() string {
	if x != nil {
		return x.DetectedText
	}
	return ""
}

func (x *TextDetection) GetConfidence() int32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *TextDetection) GetPolygon() *Polygon {
	if x != nil {
		return x.Polygon
	}
	return nil
}

func (x *TextDetection) GetAdvancedInfo() string {
	if x != nil {
		return x.AdvancedInfo
	}
	return ""
}

type Polygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates []*Coordinate `protobuf:"bytes,1,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *Polygon) Reset() {
	*x = Polygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xe07_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Polygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Polygon) ProtoMessage() {}

func (x *Polygon) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xe07_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Polygon.ProtoReflect.Descriptor instead.
func (*Polygon) Descriptor() ([]byte, []int) {
	return file_oidb0xe07_proto_rawDescGZIP(), []int{4}
}

func (x *Polygon) GetCoordinates() []*Coordinate {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xe07_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xe07_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_oidb0xe07_proto_rawDescGZIP(), []int{5}
}

func (x *Coordinate) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinate) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language     string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	LanguageDesc string `protobuf:"bytes,2,opt,name=languageDesc,proto3" json:"languageDesc,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xe07_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xe07_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_oidb0xe07_proto_rawDescGZIP(), []int{6}
}

func (x *Language) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Language) GetLanguageDesc() string {
	if x != nil {
		return x.LanguageDesc
	}
	return ""
}

type OCRRspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TextDetections           []*TextDetection `protobuf:"bytes,1,rep,name=textDetections,proto3" json:"textDetections,omitempty"`
	Language                 string           `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	RequestId                string           `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	OcrLanguageList          []string         `protobuf:"bytes,101,rep,name=ocrLanguageList,proto3" json:"ocrLanguageList,omitempty"`
	DstTranslateLanguageList []string         `protobuf:"bytes,102,rep,name=dstTranslateLanguageList,proto3" json:"dstTranslateLanguageList,omitempty"`
	LanguageList             []*Language      `protobuf:"bytes,103,rep,name=languageList,proto3" json:"languageList,omitempty"`
	AfterCompressWeight      int32            `protobuf:"varint,111,opt,name=afterCompressWeight,proto3" json:"afterCompressWeight,omitempty"`
	AfterCompressHeight      int32            `protobuf:"varint,112,opt,name=afterCompressHeight,proto3" json:"afterCompressHeight,omitempty"`
}

func (x *OCRRspBody) Reset() {
	*x = OCRRspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xe07_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCRRspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCRRspBody) ProtoMessage() {}

func (x *OCRRspBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xe07_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCRRspBody.ProtoReflect.Descriptor instead.
func (*OCRRspBody) Descriptor() ([]byte, []int) {
	return file_oidb0xe07_proto_rawDescGZIP(), []int{7}
}

func (x *OCRRspBody) GetTextDetections() []*TextDetection {
	if x != nil {
		return x.TextDetections
	}
	return nil
}

func (x *OCRRspBody) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *OCRRspBody) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *OCRRspBody) GetOcrLanguageList() []string {
	if x != nil {
		return x.OcrLanguageList
	}
	return nil
}

func (x *OCRRspBody) GetDstTranslateLanguageList() []string {
	if x != nil {
		return x.DstTranslateLanguageList
	}
	return nil
}

func (x *OCRRspBody) GetLanguageList() []*Language {
	if x != nil {
		return x.LanguageList
	}
	return nil
}

func (x *OCRRspBody) GetAfterCompressWeight() int32 {
	if x != nil {
		return x.AfterCompressWeight
	}
	return 0
}

func (x *OCRRspBody) GetAfterCompressHeight() int32 {
	if x != nil {
		return x.AfterCompressHeight
	}
	return 0
}

var File_oidb0xe07_proto protoreflect.FileDescriptor

var file_oidb0xe07_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x65, 0x30, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0b, 0x44, 0x45, 0x30, 0x37, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x0a, 0x6f, 0x63, 0x72, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4f, 0x43, 0x52, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79,
	0x52, 0x0a, 0x6f, 0x63, 0x72, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xdc, 0x02, 0x0a,
	0x0a, 0x4f, 0x43, 0x52, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d, 0x64, 0x35, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d, 0x64, 0x35, 0x12,
	0x2a, 0x0a, 0x10, 0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x4d, 0x64, 0x35, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x64, 0x35, 0x12, 0x34, 0x0a, 0x15, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x13, 0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x43, 0x75, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x43, 0x75, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x0b,
	0x44, 0x45, 0x30, 0x37, 0x52, 0x73, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2b, 0x0a, 0x0a, 0x6f, 0x63, 0x72, 0x52, 0x73,
	0x70, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4f, 0x43,
	0x52, 0x52, 0x73, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0a, 0x6f, 0x63, 0x72, 0x52, 0x73, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x78, 0x74, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x6f,
	0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x6f,
	0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x38, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52,
	0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x0a,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x59, 0x22, 0x4a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x22, 0xf7, 0x02, 0x0a, 0x0a, 0x4f, 0x43, 0x52, 0x52, 0x73, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x36, 0x0a, 0x0e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x54, 0x65, 0x78, 0x74,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x74, 0x65, 0x78, 0x74, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x63, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x63,
	0x72, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a,
	0x18, 0x64, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x66, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x18, 0x64, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0c, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x67, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x6f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x3b, 0x6f, 0x69, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oidb0xe07_proto_rawDescOnce sync.Once
	file_oidb0xe07_proto_rawDescData = file_oidb0xe07_proto_rawDesc
)

func file_oidb0xe07_proto_rawDescGZIP() []byte {
	file_oidb0xe07_proto_rawDescOnce.Do(func() {
		file_oidb0xe07_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0xe07_proto_rawDescData)
	})
	return file_oidb0xe07_proto_rawDescData
}

var file_oidb0xe07_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_oidb0xe07_proto_goTypes = []interface{}{
	(*DE07ReqBody)(nil),   // 0: DE07ReqBody
	(*OCRReqBody)(nil),    // 1: OCRReqBody
	(*DE07RspBody)(nil),   // 2: DE07RspBody
	(*TextDetection)(nil), // 3: TextDetection
	(*Polygon)(nil),       // 4: Polygon
	(*Coordinate)(nil),    // 5: Coordinate
	(*Language)(nil),      // 6: Language
	(*OCRRspBody)(nil),    // 7: OCRRspBody
}
var file_oidb0xe07_proto_depIdxs = []int32{
	1, // 0: DE07ReqBody.ocrReqBody:type_name -> OCRReqBody
	7, // 1: DE07RspBody.ocrRspBody:type_name -> OCRRspBody
	4, // 2: TextDetection.polygon:type_name -> Polygon
	5, // 3: Polygon.coordinates:type_name -> Coordinate
	3, // 4: OCRRspBody.textDetections:type_name -> TextDetection
	6, // 5: OCRRspBody.languageList:type_name -> Language
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_oidb0xe07_proto_init() }
func file_oidb0xe07_proto_init() {
	if File_oidb0xe07_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0xe07_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DE07ReqBody); i {
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
		file_oidb0xe07_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCRReqBody); i {
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
		file_oidb0xe07_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DE07RspBody); i {
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
		file_oidb0xe07_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextDetection); i {
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
		file_oidb0xe07_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Polygon); i {
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
		file_oidb0xe07_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinate); i {
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
		file_oidb0xe07_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
		file_oidb0xe07_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCRRspBody); i {
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
			RawDescriptor: file_oidb0xe07_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0xe07_proto_goTypes,
		DependencyIndexes: file_oidb0xe07_proto_depIdxs,
		MessageInfos:      file_oidb0xe07_proto_msgTypes,
	}.Build()
	File_oidb0xe07_proto = out.File
	file_oidb0xe07_proto_rawDesc = nil
	file_oidb0xe07_proto_goTypes = nil
	file_oidb0xe07_proto_depIdxs = nil
}
