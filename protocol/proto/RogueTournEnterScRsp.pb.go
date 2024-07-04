// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournEnterScRsp.proto

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

type RogueTournEnterScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueTournCurInfo      *RogueTournCurInfo      `protobuf:"bytes,15,opt,name=rogue_tourn_cur_info,json=rogueTournCurInfo,proto3" json:"rogue_tourn_cur_info,omitempty"`
	Retcode                uint32                  `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RogueTournCurSceneInfo *RogueTournCurSceneInfo `protobuf:"bytes,2,opt,name=rogue_tourn_cur_scene_info,json=rogueTournCurSceneInfo,proto3" json:"rogue_tourn_cur_scene_info,omitempty"`
}

func (x *RogueTournEnterScRsp) Reset() {
	*x = RogueTournEnterScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournEnterScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournEnterScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournEnterScRsp) ProtoMessage() {}

func (x *RogueTournEnterScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournEnterScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournEnterScRsp.ProtoReflect.Descriptor instead.
func (*RogueTournEnterScRsp) Descriptor() ([]byte, []int) {
	return file_RogueTournEnterScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournEnterScRsp) GetRogueTournCurInfo() *RogueTournCurInfo {
	if x != nil {
		return x.RogueTournCurInfo
	}
	return nil
}

func (x *RogueTournEnterScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *RogueTournEnterScRsp) GetRogueTournCurSceneInfo() *RogueTournCurSceneInfo {
	if x != nil {
		return x.RogueTournCurSceneInfo
	}
	return nil
}

var File_RogueTournEnterScRsp_proto protoreflect.FileDescriptor

var file_RogueTournEnterScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x14, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x14,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x53, 0x0a, 0x1a, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x5f, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x16, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_RogueTournEnterScRsp_proto_rawDescOnce sync.Once
	file_RogueTournEnterScRsp_proto_rawDescData = file_RogueTournEnterScRsp_proto_rawDesc
)

func file_RogueTournEnterScRsp_proto_rawDescGZIP() []byte {
	file_RogueTournEnterScRsp_proto_rawDescOnce.Do(func() {
		file_RogueTournEnterScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournEnterScRsp_proto_rawDescData)
	})
	return file_RogueTournEnterScRsp_proto_rawDescData
}

var file_RogueTournEnterScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournEnterScRsp_proto_goTypes = []interface{}{
	(*RogueTournEnterScRsp)(nil),   // 0: RogueTournEnterScRsp
	(*RogueTournCurInfo)(nil),      // 1: RogueTournCurInfo
	(*RogueTournCurSceneInfo)(nil), // 2: RogueTournCurSceneInfo
}
var file_RogueTournEnterScRsp_proto_depIdxs = []int32{
	1, // 0: RogueTournEnterScRsp.rogue_tourn_cur_info:type_name -> RogueTournCurInfo
	2, // 1: RogueTournEnterScRsp.rogue_tourn_cur_scene_info:type_name -> RogueTournCurSceneInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueTournEnterScRsp_proto_init() }
func file_RogueTournEnterScRsp_proto_init() {
	if File_RogueTournEnterScRsp_proto != nil {
		return
	}
	file_RogueTournCurSceneInfo_proto_init()
	file_RogueTournCurInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournEnterScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournEnterScRsp); i {
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
			RawDescriptor: file_RogueTournEnterScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournEnterScRsp_proto_goTypes,
		DependencyIndexes: file_RogueTournEnterScRsp_proto_depIdxs,
		MessageInfos:      file_RogueTournEnterScRsp_proto_msgTypes,
	}.Build()
	File_RogueTournEnterScRsp_proto = out.File
	file_RogueTournEnterScRsp_proto_rawDesc = nil
	file_RogueTournEnterScRsp_proto_goTypes = nil
	file_RogueTournEnterScRsp_proto_depIdxs = nil
}
