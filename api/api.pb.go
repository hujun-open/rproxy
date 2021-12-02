// vrtchef plugin API

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api.proto

package api

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type CreateWorkerCrossReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CreateWorkerCrossReq) Reset() {
	*x = CreateWorkerCrossReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkerCrossReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkerCrossReq) ProtoMessage() {}

func (x *CreateWorkerCrossReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkerCrossReq.ProtoReflect.Descriptor instead.
func (*CreateWorkerCrossReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWorkerCrossReq) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReportWorkerCrossReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *ReportWorkerCrossReq) Reset() {
	*x = ReportWorkerCrossReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportWorkerCrossReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportWorkerCrossReq) ProtoMessage() {}

func (x *ReportWorkerCrossReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportWorkerCrossReq.ProtoReflect.Descriptor instead.
func (*ReportWorkerCrossReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *ReportWorkerCrossReq) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ReportWorkerCrossReq) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x3a, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x32, 0xd0, 0x01,
	0x0a, 0x09, 0x52, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x50, 0x49, 0x12, 0x20, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x6f, 0x6e, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x21, 0x0a,
	0x07, 0x53, 0x69, 0x67, 0x6e, 0x6f, 0x66, 0x66, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x43, 0x72, 0x6f, 0x73, 0x73, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x71, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x3e, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x43, 0x72, 0x6f, 0x73, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x42, 0x0c, 0x5a, 0x0a, 0x72, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_goTypes = []interface{}{
	(*Empty)(nil),                // 0: api.Empty
	(*CreateWorkerCrossReq)(nil), // 1: api.CreateWorkerCrossReq
	(*ReportWorkerCrossReq)(nil), // 2: api.ReportWorkerCrossReq
}
var file_api_proto_depIdxs = []int32{
	0, // 0: api.RProxyAPI.Signon:input_type -> api.Empty
	0, // 1: api.RProxyAPI.Signoff:input_type -> api.Empty
	0, // 2: api.RProxyAPI.CreateWorkerCross:input_type -> api.Empty
	2, // 3: api.RProxyAPI.ReportWorkerCross:input_type -> api.ReportWorkerCrossReq
	0, // 4: api.RProxyAPI.Signon:output_type -> api.Empty
	0, // 5: api.RProxyAPI.Signoff:output_type -> api.Empty
	1, // 6: api.RProxyAPI.CreateWorkerCross:output_type -> api.CreateWorkerCrossReq
	0, // 7: api.RProxyAPI.ReportWorkerCross:output_type -> api.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkerCrossReq); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportWorkerCrossReq); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
