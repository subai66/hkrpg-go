// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: OOLKOIOABDL.proto

package proto

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

type OOLKOIOABDL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ICKOEBGMNNP *OINKLEHKHED `protobuf:"bytes,6,opt,name=ICKOEBGMNNP,proto3" json:"ICKOEBGMNNP,omitempty"`
	LCOBBGLKKLC *PLBNBBDNOOK `protobuf:"bytes,15,opt,name=LCOBBGLKKLC,proto3" json:"LCOBBGLKKLC,omitempty"`
	FMKNLENCOOD *AEGBGMGODEE `protobuf:"bytes,3,opt,name=FMKNLENCOOD,proto3" json:"FMKNLENCOOD,omitempty"`
}

func (x *OOLKOIOABDL) Reset() {
	*x = OOLKOIOABDL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OOLKOIOABDL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OOLKOIOABDL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OOLKOIOABDL) ProtoMessage() {}

func (x *OOLKOIOABDL) ProtoReflect() protoreflect.Message {
	mi := &file_OOLKOIOABDL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OOLKOIOABDL.ProtoReflect.Descriptor instead.
func (*OOLKOIOABDL) Descriptor() ([]byte, []int) {
	return file_OOLKOIOABDL_proto_rawDescGZIP(), []int{0}
}

func (x *OOLKOIOABDL) GetICKOEBGMNNP() *OINKLEHKHED {
	if x != nil {
		return x.ICKOEBGMNNP
	}
	return nil
}

func (x *OOLKOIOABDL) GetLCOBBGLKKLC() *PLBNBBDNOOK {
	if x != nil {
		return x.LCOBBGLKKLC
	}
	return nil
}

func (x *OOLKOIOABDL) GetFMKNLENCOOD() *AEGBGMGODEE {
	if x != nil {
		return x.FMKNLENCOOD
	}
	return nil
}

var File_OOLKOIOABDL_proto protoreflect.FileDescriptor

var file_OOLKOIOABDL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x4f, 0x4c, 0x4b, 0x4f, 0x49, 0x4f, 0x41, 0x42, 0x44, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x45, 0x47, 0x42, 0x47, 0x4d, 0x47, 0x4f, 0x44, 0x45, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x4c, 0x42, 0x4e, 0x42, 0x42, 0x44, 0x4e,
	0x4f, 0x4f, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x49, 0x4e, 0x4b, 0x4c,
	0x45, 0x48, 0x4b, 0x48, 0x45, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a,
	0x0b, 0x4f, 0x4f, 0x4c, 0x4b, 0x4f, 0x49, 0x4f, 0x41, 0x42, 0x44, 0x4c, 0x12, 0x2e, 0x0a, 0x0b,
	0x49, 0x43, 0x4b, 0x4f, 0x45, 0x42, 0x47, 0x4d, 0x4e, 0x4e, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x49, 0x4e, 0x4b, 0x4c, 0x45, 0x48, 0x4b, 0x48, 0x45, 0x44, 0x52,
	0x0b, 0x49, 0x43, 0x4b, 0x4f, 0x45, 0x42, 0x47, 0x4d, 0x4e, 0x4e, 0x50, 0x12, 0x2e, 0x0a, 0x0b,
	0x4c, 0x43, 0x4f, 0x42, 0x42, 0x47, 0x4c, 0x4b, 0x4b, 0x4c, 0x43, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x4c, 0x42, 0x4e, 0x42, 0x42, 0x44, 0x4e, 0x4f, 0x4f, 0x4b, 0x52,
	0x0b, 0x4c, 0x43, 0x4f, 0x42, 0x42, 0x47, 0x4c, 0x4b, 0x4b, 0x4c, 0x43, 0x12, 0x2e, 0x0a, 0x0b,
	0x46, 0x4d, 0x4b, 0x4e, 0x4c, 0x45, 0x4e, 0x43, 0x4f, 0x4f, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x45, 0x47, 0x42, 0x47, 0x4d, 0x47, 0x4f, 0x44, 0x45, 0x45, 0x52,
	0x0b, 0x46, 0x4d, 0x4b, 0x4e, 0x4c, 0x45, 0x4e, 0x43, 0x4f, 0x4f, 0x44, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OOLKOIOABDL_proto_rawDescOnce sync.Once
	file_OOLKOIOABDL_proto_rawDescData = file_OOLKOIOABDL_proto_rawDesc
)

func file_OOLKOIOABDL_proto_rawDescGZIP() []byte {
	file_OOLKOIOABDL_proto_rawDescOnce.Do(func() {
		file_OOLKOIOABDL_proto_rawDescData = protoimpl.X.CompressGZIP(file_OOLKOIOABDL_proto_rawDescData)
	})
	return file_OOLKOIOABDL_proto_rawDescData
}

var file_OOLKOIOABDL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OOLKOIOABDL_proto_goTypes = []interface{}{
	(*OOLKOIOABDL)(nil), // 0: OOLKOIOABDL
	(*OINKLEHKHED)(nil), // 1: OINKLEHKHED
	(*PLBNBBDNOOK)(nil), // 2: PLBNBBDNOOK
	(*AEGBGMGODEE)(nil), // 3: AEGBGMGODEE
}
var file_OOLKOIOABDL_proto_depIdxs = []int32{
	1, // 0: OOLKOIOABDL.ICKOEBGMNNP:type_name -> OINKLEHKHED
	2, // 1: OOLKOIOABDL.LCOBBGLKKLC:type_name -> PLBNBBDNOOK
	3, // 2: OOLKOIOABDL.FMKNLENCOOD:type_name -> AEGBGMGODEE
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_OOLKOIOABDL_proto_init() }
func file_OOLKOIOABDL_proto_init() {
	if File_OOLKOIOABDL_proto != nil {
		return
	}
	file_AEGBGMGODEE_proto_init()
	file_PLBNBBDNOOK_proto_init()
	file_OINKLEHKHED_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OOLKOIOABDL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OOLKOIOABDL); i {
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
			RawDescriptor: file_OOLKOIOABDL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OOLKOIOABDL_proto_goTypes,
		DependencyIndexes: file_OOLKOIOABDL_proto_depIdxs,
		MessageInfos:      file_OOLKOIOABDL_proto_msgTypes,
	}.Build()
	File_OOLKOIOABDL_proto = out.File
	file_OOLKOIOABDL_proto_rawDesc = nil
	file_OOLKOIOABDL_proto_goTypes = nil
	file_OOLKOIOABDL_proto_depIdxs = nil
}
