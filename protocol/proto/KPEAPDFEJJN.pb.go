// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: KPEAPDFEJJN.proto

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

type KPEAPDFEJJN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AGEPMGAOBEP *DANACCHECHA `protobuf:"bytes,12,opt,name=AGEPMGAOBEP,proto3" json:"AGEPMGAOBEP,omitempty"`
	GGGJACPNJNP *EBMPINJCHFB `protobuf:"bytes,2,opt,name=GGGJACPNJNP,proto3" json:"GGGJACPNJNP,omitempty"`
	IPKNIJNEFIJ *NDCKDMOMLFK `protobuf:"bytes,13,opt,name=IPKNIJNEFIJ,proto3" json:"IPKNIJNEFIJ,omitempty"`
	LCKCBBKMDNI *ABCMHGNGFGL `protobuf:"bytes,3,opt,name=LCKCBBKMDNI,proto3" json:"LCKCBBKMDNI,omitempty"`
}

func (x *KPEAPDFEJJN) Reset() {
	*x = KPEAPDFEJJN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KPEAPDFEJJN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KPEAPDFEJJN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KPEAPDFEJJN) ProtoMessage() {}

func (x *KPEAPDFEJJN) ProtoReflect() protoreflect.Message {
	mi := &file_KPEAPDFEJJN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KPEAPDFEJJN.ProtoReflect.Descriptor instead.
func (*KPEAPDFEJJN) Descriptor() ([]byte, []int) {
	return file_KPEAPDFEJJN_proto_rawDescGZIP(), []int{0}
}

func (x *KPEAPDFEJJN) GetAGEPMGAOBEP() *DANACCHECHA {
	if x != nil {
		return x.AGEPMGAOBEP
	}
	return nil
}

func (x *KPEAPDFEJJN) GetGGGJACPNJNP() *EBMPINJCHFB {
	if x != nil {
		return x.GGGJACPNJNP
	}
	return nil
}

func (x *KPEAPDFEJJN) GetIPKNIJNEFIJ() *NDCKDMOMLFK {
	if x != nil {
		return x.IPKNIJNEFIJ
	}
	return nil
}

func (x *KPEAPDFEJJN) GetLCKCBBKMDNI() *ABCMHGNGFGL {
	if x != nil {
		return x.LCKCBBKMDNI
	}
	return nil
}

var File_KPEAPDFEJJN_proto protoreflect.FileDescriptor

var file_KPEAPDFEJJN_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4b, 0x50, 0x45, 0x41, 0x50, 0x44, 0x46, 0x45, 0x4a, 0x4a, 0x4e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x44, 0x43, 0x4b, 0x44, 0x4d, 0x4f, 0x4d, 0x4c, 0x46, 0x4b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x42, 0x4d, 0x50, 0x49, 0x4e, 0x4a, 0x43,
	0x48, 0x46, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x42, 0x43, 0x4d, 0x48,
	0x47, 0x4e, 0x47, 0x46, 0x47, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x41,
	0x4e, 0x41, 0x43, 0x43, 0x48, 0x45, 0x43, 0x48, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcd, 0x01, 0x0a, 0x0b, 0x4b, 0x50, 0x45, 0x41, 0x50, 0x44, 0x46, 0x45, 0x4a, 0x4a, 0x4e, 0x12,
	0x2e, 0x0a, 0x0b, 0x41, 0x47, 0x45, 0x50, 0x4d, 0x47, 0x41, 0x4f, 0x42, 0x45, 0x50, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x41, 0x4e, 0x41, 0x43, 0x43, 0x48, 0x45, 0x43,
	0x48, 0x41, 0x52, 0x0b, 0x41, 0x47, 0x45, 0x50, 0x4d, 0x47, 0x41, 0x4f, 0x42, 0x45, 0x50, 0x12,
	0x2e, 0x0a, 0x0b, 0x47, 0x47, 0x47, 0x4a, 0x41, 0x43, 0x50, 0x4e, 0x4a, 0x4e, 0x50, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x42, 0x4d, 0x50, 0x49, 0x4e, 0x4a, 0x43, 0x48,
	0x46, 0x42, 0x52, 0x0b, 0x47, 0x47, 0x47, 0x4a, 0x41, 0x43, 0x50, 0x4e, 0x4a, 0x4e, 0x50, 0x12,
	0x2e, 0x0a, 0x0b, 0x49, 0x50, 0x4b, 0x4e, 0x49, 0x4a, 0x4e, 0x45, 0x46, 0x49, 0x4a, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x44, 0x43, 0x4b, 0x44, 0x4d, 0x4f, 0x4d, 0x4c,
	0x46, 0x4b, 0x52, 0x0b, 0x49, 0x50, 0x4b, 0x4e, 0x49, 0x4a, 0x4e, 0x45, 0x46, 0x49, 0x4a, 0x12,
	0x2e, 0x0a, 0x0b, 0x4c, 0x43, 0x4b, 0x43, 0x42, 0x42, 0x4b, 0x4d, 0x44, 0x4e, 0x49, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x42, 0x43, 0x4d, 0x48, 0x47, 0x4e, 0x47, 0x46,
	0x47, 0x4c, 0x52, 0x0b, 0x4c, 0x43, 0x4b, 0x43, 0x42, 0x42, 0x4b, 0x4d, 0x44, 0x4e, 0x49, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_KPEAPDFEJJN_proto_rawDescOnce sync.Once
	file_KPEAPDFEJJN_proto_rawDescData = file_KPEAPDFEJJN_proto_rawDesc
)

func file_KPEAPDFEJJN_proto_rawDescGZIP() []byte {
	file_KPEAPDFEJJN_proto_rawDescOnce.Do(func() {
		file_KPEAPDFEJJN_proto_rawDescData = protoimpl.X.CompressGZIP(file_KPEAPDFEJJN_proto_rawDescData)
	})
	return file_KPEAPDFEJJN_proto_rawDescData
}

var file_KPEAPDFEJJN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KPEAPDFEJJN_proto_goTypes = []interface{}{
	(*KPEAPDFEJJN)(nil), // 0: KPEAPDFEJJN
	(*DANACCHECHA)(nil), // 1: DANACCHECHA
	(*EBMPINJCHFB)(nil), // 2: EBMPINJCHFB
	(*NDCKDMOMLFK)(nil), // 3: NDCKDMOMLFK
	(*ABCMHGNGFGL)(nil), // 4: ABCMHGNGFGL
}
var file_KPEAPDFEJJN_proto_depIdxs = []int32{
	1, // 0: KPEAPDFEJJN.AGEPMGAOBEP:type_name -> DANACCHECHA
	2, // 1: KPEAPDFEJJN.GGGJACPNJNP:type_name -> EBMPINJCHFB
	3, // 2: KPEAPDFEJJN.IPKNIJNEFIJ:type_name -> NDCKDMOMLFK
	4, // 3: KPEAPDFEJJN.LCKCBBKMDNI:type_name -> ABCMHGNGFGL
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_KPEAPDFEJJN_proto_init() }
func file_KPEAPDFEJJN_proto_init() {
	if File_KPEAPDFEJJN_proto != nil {
		return
	}
	file_NDCKDMOMLFK_proto_init()
	file_EBMPINJCHFB_proto_init()
	file_ABCMHGNGFGL_proto_init()
	file_DANACCHECHA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_KPEAPDFEJJN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KPEAPDFEJJN); i {
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
			RawDescriptor: file_KPEAPDFEJJN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KPEAPDFEJJN_proto_goTypes,
		DependencyIndexes: file_KPEAPDFEJJN_proto_depIdxs,
		MessageInfos:      file_KPEAPDFEJJN_proto_msgTypes,
	}.Build()
	File_KPEAPDFEJJN_proto = out.File
	file_KPEAPDFEJJN_proto_rawDesc = nil
	file_KPEAPDFEJJN_proto_goTypes = nil
	file_KPEAPDFEJJN_proto_depIdxs = nil
}
