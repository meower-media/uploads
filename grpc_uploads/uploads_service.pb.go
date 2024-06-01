// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: uploads_service.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClaimFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
}

func (x *ClaimFileReq) Reset() {
	*x = ClaimFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimFileReq) ProtoMessage() {}

func (x *ClaimFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimFileReq.ProtoReflect.Descriptor instead.
func (*ClaimFileReq) Descriptor() ([]byte, []int) {
	return file_uploads_service_proto_rawDescGZIP(), []int{0}
}

func (x *ClaimFileReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClaimFileReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

type ClaimFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mime     string `protobuf:"bytes,2,opt,name=mime,proto3" json:"mime,omitempty"`
	Filename string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Size     int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Width    int32  `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height   int32  `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ClaimFileResp) Reset() {
	*x = ClaimFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimFileResp) ProtoMessage() {}

func (x *ClaimFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimFileResp.ProtoReflect.Descriptor instead.
func (*ClaimFileResp) Descriptor() ([]byte, []int) {
	return file_uploads_service_proto_rawDescGZIP(), []int{1}
}

func (x *ClaimFileResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClaimFileResp) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *ClaimFileResp) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ClaimFileResp) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ClaimFileResp) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ClaimFileResp) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type DeleteFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFileReq) Reset() {
	*x = DeleteFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileReq) ProtoMessage() {}

func (x *DeleteFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileReq.ProtoReflect.Descriptor instead.
func (*DeleteFileReq) Descriptor() ([]byte, []int) {
	return file_uploads_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteFileReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ClearFilesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ClearFilesReq) Reset() {
	*x = ClearFilesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearFilesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearFilesReq) ProtoMessage() {}

func (x *ClearFilesReq) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearFilesReq.ProtoReflect.Descriptor instead.
func (*ClearFilesReq) Descriptor() ([]byte, []int) {
	return file_uploads_service_proto_rawDescGZIP(), []int{3}
}

func (x *ClearFilesReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_uploads_service_proto protoreflect.FileDescriptor

var file_uploads_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x0c, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0d, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x32, 0xc1, 0x01, 0x0a, 0x07, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x12, 0x3a, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uploads_service_proto_rawDescOnce sync.Once
	file_uploads_service_proto_rawDescData = file_uploads_service_proto_rawDesc
)

func file_uploads_service_proto_rawDescGZIP() []byte {
	file_uploads_service_proto_rawDescOnce.Do(func() {
		file_uploads_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_uploads_service_proto_rawDescData)
	})
	return file_uploads_service_proto_rawDescData
}

var file_uploads_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_uploads_service_proto_goTypes = []interface{}{
	(*ClaimFileReq)(nil),  // 0: uploads.ClaimFileReq
	(*ClaimFileResp)(nil), // 1: uploads.ClaimFileResp
	(*DeleteFileReq)(nil), // 2: uploads.DeleteFileReq
	(*ClearFilesReq)(nil), // 3: uploads.ClearFilesReq
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_uploads_service_proto_depIdxs = []int32{
	0, // 0: uploads.Uploads.ClaimFile:input_type -> uploads.ClaimFileReq
	2, // 1: uploads.Uploads.DeleteFile:input_type -> uploads.DeleteFileReq
	3, // 2: uploads.Uploads.ClearFiles:input_type -> uploads.ClearFilesReq
	1, // 3: uploads.Uploads.ClaimFile:output_type -> uploads.ClaimFileResp
	4, // 4: uploads.Uploads.DeleteFile:output_type -> google.protobuf.Empty
	4, // 5: uploads.Uploads.ClearFiles:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_uploads_service_proto_init() }
func file_uploads_service_proto_init() {
	if File_uploads_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uploads_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimFileReq); i {
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
		file_uploads_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimFileResp); i {
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
		file_uploads_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileReq); i {
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
		file_uploads_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearFilesReq); i {
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
			RawDescriptor: file_uploads_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uploads_service_proto_goTypes,
		DependencyIndexes: file_uploads_service_proto_depIdxs,
		MessageInfos:      file_uploads_service_proto_msgTypes,
	}.Build()
	File_uploads_service_proto = out.File
	file_uploads_service_proto_rawDesc = nil
	file_uploads_service_proto_goTypes = nil
	file_uploads_service_proto_depIdxs = nil
}