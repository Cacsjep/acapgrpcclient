// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: cctv.proto

package cctv

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

type FrameMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frame       []byte `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	Height      int32  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Width       int32  `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	SequenceNbr int32  `protobuf:"varint,4,opt,name=sequence_nbr,json=sequenceNbr,proto3" json:"sequence_nbr,omitempty"`
	Framerate   int32  `protobuf:"varint,5,opt,name=framerate,proto3" json:"framerate,omitempty"`
}

func (x *FrameMessage) Reset() {
	*x = FrameMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameMessage) ProtoMessage() {}

func (x *FrameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameMessage.ProtoReflect.Descriptor instead.
func (*FrameMessage) Descriptor() ([]byte, []int) {
	return file_cctv_proto_rawDescGZIP(), []int{0}
}

func (x *FrameMessage) GetFrame() []byte {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *FrameMessage) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *FrameMessage) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *FrameMessage) GetSequenceNbr() int32 {
	if x != nil {
		return x.SequenceNbr
	}
	return 0
}

func (x *FrameMessage) GetFramerate() int32 {
	if x != nil {
		return x.Framerate
	}
	return 0
}

type Acknowledgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Acknowledgement) Reset() {
	*x = Acknowledgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledgement) ProtoMessage() {}

func (x *Acknowledgement) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acknowledgement.ProtoReflect.Descriptor instead.
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return file_cctv_proto_rawDescGZIP(), []int{1}
}

func (x *Acknowledgement) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cctv_proto protoreflect.FileDescriptor

var file_cctv_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x63, 0x74, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a,
	0x0c, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x62,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x4e, 0x62, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x41, 0x0a, 0x0b, 0x43, 0x43, 0x54, 0x56, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32,
	0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0d, 0x2e,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x41,
	0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x61, 0x63, 0x73, 0x6a, 0x65, 0x70, 0x2f, 0x61, 0x63, 0x61, 0x70, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x63, 0x74, 0x76, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cctv_proto_rawDescOnce sync.Once
	file_cctv_proto_rawDescData = file_cctv_proto_rawDesc
)

func file_cctv_proto_rawDescGZIP() []byte {
	file_cctv_proto_rawDescOnce.Do(func() {
		file_cctv_proto_rawDescData = protoimpl.X.CompressGZIP(file_cctv_proto_rawDescData)
	})
	return file_cctv_proto_rawDescData
}

var file_cctv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cctv_proto_goTypes = []any{
	(*FrameMessage)(nil),    // 0: FrameMessage
	(*Acknowledgement)(nil), // 1: Acknowledgement
}
var file_cctv_proto_depIdxs = []int32{
	0, // 0: CCTVService.StreamVideo:input_type -> FrameMessage
	1, // 1: CCTVService.StreamVideo:output_type -> Acknowledgement
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cctv_proto_init() }
func file_cctv_proto_init() {
	if File_cctv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cctv_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FrameMessage); i {
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
		file_cctv_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Acknowledgement); i {
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
			RawDescriptor: file_cctv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cctv_proto_goTypes,
		DependencyIndexes: file_cctv_proto_depIdxs,
		MessageInfos:      file_cctv_proto_msgTypes,
	}.Build()
	File_cctv_proto = out.File
	file_cctv_proto_rawDesc = nil
	file_cctv_proto_goTypes = nil
	file_cctv_proto_depIdxs = nil
}
