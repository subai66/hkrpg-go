// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ADJGFJECBEI.proto

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

type ADJGFJECBEI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DNABPNFKDKJ uint32         `protobuf:"varint,3,opt,name=DNABPNFKDKJ,proto3" json:"DNABPNFKDKJ,omitempty"`
	FCLLFHGGAPF *KPEAPDFEJJN   `protobuf:"bytes,1,opt,name=FCLLFHGGAPF,proto3" json:"FCLLFHGGAPF,omitempty"`
	AKEBAFKANGF []*HMINCBJOEOD `protobuf:"bytes,9,rep,name=AKEBAFKANGF,proto3" json:"AKEBAFKANGF,omitempty"`
}

func (x *ADJGFJECBEI) Reset() {
	*x = ADJGFJECBEI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ADJGFJECBEI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADJGFJECBEI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADJGFJECBEI) ProtoMessage() {}

func (x *ADJGFJECBEI) ProtoReflect() protoreflect.Message {
	mi := &file_ADJGFJECBEI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADJGFJECBEI.ProtoReflect.Descriptor instead.
func (*ADJGFJECBEI) Descriptor() ([]byte, []int) {
	return file_ADJGFJECBEI_proto_rawDescGZIP(), []int{0}
}

func (x *ADJGFJECBEI) GetDNABPNFKDKJ() uint32 {
	if x != nil {
		return x.DNABPNFKDKJ
	}
	return 0
}

func (x *ADJGFJECBEI) GetFCLLFHGGAPF() *KPEAPDFEJJN {
	if x != nil {
		return x.FCLLFHGGAPF
	}
	return nil
}

func (x *ADJGFJECBEI) GetAKEBAFKANGF() []*HMINCBJOEOD {
	if x != nil {
		return x.AKEBAFKANGF
	}
	return nil
}

var File_ADJGFJECBEI_proto protoreflect.FileDescriptor

var file_ADJGFJECBEI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x44, 0x4a, 0x47, 0x46, 0x4a, 0x45, 0x43, 0x42, 0x45, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x4d, 0x49, 0x4e, 0x43, 0x42, 0x4a, 0x4f, 0x45, 0x4f, 0x44,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x50, 0x45, 0x41, 0x50, 0x44, 0x46, 0x45,
	0x4a, 0x4a, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x41, 0x44,
	0x4a, 0x47, 0x46, 0x4a, 0x45, 0x43, 0x42, 0x45, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4e, 0x41,
	0x42, 0x50, 0x4e, 0x46, 0x4b, 0x44, 0x4b, 0x4a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x44, 0x4e, 0x41, 0x42, 0x50, 0x4e, 0x46, 0x4b, 0x44, 0x4b, 0x4a, 0x12, 0x2e, 0x0a, 0x0b, 0x46,
	0x43, 0x4c, 0x4c, 0x46, 0x48, 0x47, 0x47, 0x41, 0x50, 0x46, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4b, 0x50, 0x45, 0x41, 0x50, 0x44, 0x46, 0x45, 0x4a, 0x4a, 0x4e, 0x52, 0x0b,
	0x46, 0x43, 0x4c, 0x4c, 0x46, 0x48, 0x47, 0x47, 0x41, 0x50, 0x46, 0x12, 0x2e, 0x0a, 0x0b, 0x41,
	0x4b, 0x45, 0x42, 0x41, 0x46, 0x4b, 0x41, 0x4e, 0x47, 0x46, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x48, 0x4d, 0x49, 0x4e, 0x43, 0x42, 0x4a, 0x4f, 0x45, 0x4f, 0x44, 0x52, 0x0b,
	0x41, 0x4b, 0x45, 0x42, 0x41, 0x46, 0x4b, 0x41, 0x4e, 0x47, 0x46, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ADJGFJECBEI_proto_rawDescOnce sync.Once
	file_ADJGFJECBEI_proto_rawDescData = file_ADJGFJECBEI_proto_rawDesc
)

func file_ADJGFJECBEI_proto_rawDescGZIP() []byte {
	file_ADJGFJECBEI_proto_rawDescOnce.Do(func() {
		file_ADJGFJECBEI_proto_rawDescData = protoimpl.X.CompressGZIP(file_ADJGFJECBEI_proto_rawDescData)
	})
	return file_ADJGFJECBEI_proto_rawDescData
}

var file_ADJGFJECBEI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ADJGFJECBEI_proto_goTypes = []interface{}{
	(*ADJGFJECBEI)(nil), // 0: ADJGFJECBEI
	(*KPEAPDFEJJN)(nil), // 1: KPEAPDFEJJN
	(*HMINCBJOEOD)(nil), // 2: HMINCBJOEOD
}
var file_ADJGFJECBEI_proto_depIdxs = []int32{
	1, // 0: ADJGFJECBEI.FCLLFHGGAPF:type_name -> KPEAPDFEJJN
	2, // 1: ADJGFJECBEI.AKEBAFKANGF:type_name -> HMINCBJOEOD
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ADJGFJECBEI_proto_init() }
func file_ADJGFJECBEI_proto_init() {
	if File_ADJGFJECBEI_proto != nil {
		return
	}
	file_HMINCBJOEOD_proto_init()
	file_KPEAPDFEJJN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ADJGFJECBEI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADJGFJECBEI); i {
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
			RawDescriptor: file_ADJGFJECBEI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ADJGFJECBEI_proto_goTypes,
		DependencyIndexes: file_ADJGFJECBEI_proto_depIdxs,
		MessageInfos:      file_ADJGFJECBEI_proto_msgTypes,
	}.Build()
	File_ADJGFJECBEI_proto = out.File
	file_ADJGFJECBEI_proto_rawDesc = nil
	file_ADJGFJECBEI_proto_goTypes = nil
	file_ADJGFJECBEI_proto_depIdxs = nil
}
