// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: photo.proto

package photo

import (
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

type PhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Databytes []byte `protobuf:"bytes,2,opt,name=databytes,proto3" json:"databytes,omitempty"`
}

func (x *PhotoRequest) Reset() {
	*x = PhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoRequest) ProtoMessage() {}

func (x *PhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_photo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoRequest.ProtoReflect.Descriptor instead.
func (*PhotoRequest) Descriptor() ([]byte, []int) {
	return file_photo_proto_rawDescGZIP(), []int{0}
}

func (x *PhotoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PhotoRequest) GetDatabytes() []byte {
	if x != nil {
		return x.Databytes
	}
	return nil
}

type PhotoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *PhotoReply) Reset() {
	*x = PhotoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoReply) ProtoMessage() {}

func (x *PhotoReply) ProtoReflect() protoreflect.Message {
	mi := &file_photo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoReply.ProtoReflect.Descriptor instead.
func (*PhotoReply) Descriptor() ([]byte, []int) {
	return file_photo_proto_rawDescGZIP(), []int{1}
}

func (x *PhotoReply) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

var File_photo_proto protoreflect.FileDescriptor

var file_photo_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0c, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x0a, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0x40, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x65,
	0x72, 0x12, 0x35, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_photo_proto_rawDescOnce sync.Once
	file_photo_proto_rawDescData = file_photo_proto_rawDesc
)

func file_photo_proto_rawDescGZIP() []byte {
	file_photo_proto_rawDescOnce.Do(func() {
		file_photo_proto_rawDescData = protoimpl.X.CompressGZIP(file_photo_proto_rawDescData)
	})
	return file_photo_proto_rawDescData
}

var file_photo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_photo_proto_goTypes = []interface{}{
	(*PhotoRequest)(nil), // 0: photo.PhotoRequest
	(*PhotoReply)(nil),   // 1: photo.PhotoReply
}
var file_photo_proto_depIdxs = []int32{
	0, // 0: photo.Photoer.SendPhoto:input_type -> photo.PhotoRequest
	1, // 1: photo.Photoer.SendPhoto:output_type -> photo.PhotoReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_photo_proto_init() }
func file_photo_proto_init() {
	if File_photo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_photo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoRequest); i {
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
		file_photo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoReply); i {
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
			RawDescriptor: file_photo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_photo_proto_goTypes,
		DependencyIndexes: file_photo_proto_depIdxs,
		MessageInfos:      file_photo_proto_msgTypes,
	}.Build()
	File_photo_proto = out.File
	file_photo_proto_rawDesc = nil
	file_photo_proto_goTypes = nil
	file_photo_proto_depIdxs = nil
}
