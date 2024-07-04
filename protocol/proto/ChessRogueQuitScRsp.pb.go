// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueQuitScRsp.proto

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

type ChessRogueQuitScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishInfo       *ChessRogueFinishInfo    `protobuf:"bytes,15,opt,name=finish_info,json=finishInfo,proto3" json:"finish_info,omitempty"`
	RogueCurrentInfo *ChessRogueQueryGameInfo `protobuf:"bytes,7,opt,name=rogue_current_info,json=rogueCurrentInfo,proto3" json:"rogue_current_info,omitempty"`
	Retcode          uint32                   `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Info             *ChessRogueCurrentInfo   `protobuf:"bytes,11,opt,name=info,proto3" json:"info,omitempty"`
	RogueGetInfo     *ChessRogueGetInfo       `protobuf:"bytes,9,opt,name=rogue_get_info,json=rogueGetInfo,proto3" json:"rogue_get_info,omitempty"`
	PlayerInfo       *BLHOEGDLHNA             `protobuf:"bytes,2,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	RogueAeonInfo    *ILACICKKNOF             `protobuf:"bytes,10,opt,name=rogue_aeon_info,json=rogueAeonInfo,proto3" json:"rogue_aeon_info,omitempty"`
	LevelInfo        *ChessRogueLevelInfo     `protobuf:"bytes,3,opt,name=level_info,json=levelInfo,proto3" json:"level_info,omitempty"`
	JDFLOJFCJMO      *AJPAMBEHJKM             `protobuf:"bytes,13,opt,name=JDFLOJFCJMO,proto3" json:"JDFLOJFCJMO,omitempty"`
}

func (x *ChessRogueQuitScRsp) Reset() {
	*x = ChessRogueQuitScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueQuitScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueQuitScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueQuitScRsp) ProtoMessage() {}

func (x *ChessRogueQuitScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueQuitScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueQuitScRsp.ProtoReflect.Descriptor instead.
func (*ChessRogueQuitScRsp) Descriptor() ([]byte, []int) {
	return file_ChessRogueQuitScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueQuitScRsp) GetFinishInfo() *ChessRogueFinishInfo {
	if x != nil {
		return x.FinishInfo
	}
	return nil
}

func (x *ChessRogueQuitScRsp) GetRogueCurrentInfo() *ChessRogueQueryGameInfo {
	if x != nil {
		return x.RogueCurrentInfo
	}
	return nil
}

func (x *ChessRogueQuitScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ChessRogueQuitScRsp) GetInfo() *ChessRogueCurrentInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ChessRogueQuitScRsp) GetRogueGetInfo() *ChessRogueGetInfo {
	if x != nil {
		return x.RogueGetInfo
	}
	return nil
}

func (x *ChessRogueQuitScRsp) GetPlayerInfo() *BLHOEGDLHNA {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *ChessRogueQuitScRsp) GetRogueAeonInfo() *ILACICKKNOF {
	if x != nil {
		return x.RogueAeonInfo
	}
	return nil
}

func (x *ChessRogueQuitScRsp) GetLevelInfo() *ChessRogueLevelInfo {
	if x != nil {
		return x.LevelInfo
	}
	return nil
}

func (x *ChessRogueQuitScRsp) GetJDFLOJFCJMO() *AJPAMBEHJKM {
	if x != nil {
		return x.JDFLOJFCJMO
	}
	return nil
}

var File_ChessRogueQuitScRsp_proto protoreflect.FileDescriptor

var file_ChessRogueQuitScRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x69, 0x74,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x4c, 0x41, 0x43, 0x49, 0x43, 0x4b,
	0x4b, 0x4e, 0x4f, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x4a, 0x50, 0x41,
	0x4d, 0x42, 0x45, 0x48, 0x4a, 0x4b, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42,
	0x4c, 0x48, 0x4f, 0x45, 0x47, 0x44, 0x4c, 0x48, 0x4e, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdf, 0x03, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51,
	0x75, 0x69, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x46, 0x0a, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x38,
	0x0a, 0x0e, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x42, 0x4c, 0x48, 0x4f, 0x45, 0x47, 0x44, 0x4c, 0x48, 0x4e, 0x41, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x5f, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x49, 0x4c, 0x41, 0x43, 0x49, 0x43, 0x4b, 0x4b, 0x4e, 0x4f, 0x46, 0x52, 0x0d,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a,
	0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x44, 0x46, 0x4c, 0x4f, 0x4a, 0x46, 0x43, 0x4a, 0x4d,
	0x4f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x4a, 0x50, 0x41, 0x4d, 0x42,
	0x45, 0x48, 0x4a, 0x4b, 0x4d, 0x52, 0x0b, 0x4a, 0x44, 0x46, 0x4c, 0x4f, 0x4a, 0x46, 0x43, 0x4a,
	0x4d, 0x4f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueQuitScRsp_proto_rawDescOnce sync.Once
	file_ChessRogueQuitScRsp_proto_rawDescData = file_ChessRogueQuitScRsp_proto_rawDesc
)

func file_ChessRogueQuitScRsp_proto_rawDescGZIP() []byte {
	file_ChessRogueQuitScRsp_proto_rawDescOnce.Do(func() {
		file_ChessRogueQuitScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueQuitScRsp_proto_rawDescData)
	})
	return file_ChessRogueQuitScRsp_proto_rawDescData
}

var file_ChessRogueQuitScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueQuitScRsp_proto_goTypes = []interface{}{
	(*ChessRogueQuitScRsp)(nil),     // 0: ChessRogueQuitScRsp
	(*ChessRogueFinishInfo)(nil),    // 1: ChessRogueFinishInfo
	(*ChessRogueQueryGameInfo)(nil), // 2: ChessRogueQueryGameInfo
	(*ChessRogueCurrentInfo)(nil),   // 3: ChessRogueCurrentInfo
	(*ChessRogueGetInfo)(nil),       // 4: ChessRogueGetInfo
	(*BLHOEGDLHNA)(nil),             // 5: BLHOEGDLHNA
	(*ILACICKKNOF)(nil),             // 6: ILACICKKNOF
	(*ChessRogueLevelInfo)(nil),     // 7: ChessRogueLevelInfo
	(*AJPAMBEHJKM)(nil),             // 8: AJPAMBEHJKM
}
var file_ChessRogueQuitScRsp_proto_depIdxs = []int32{
	1, // 0: ChessRogueQuitScRsp.finish_info:type_name -> ChessRogueFinishInfo
	2, // 1: ChessRogueQuitScRsp.rogue_current_info:type_name -> ChessRogueQueryGameInfo
	3, // 2: ChessRogueQuitScRsp.info:type_name -> ChessRogueCurrentInfo
	4, // 3: ChessRogueQuitScRsp.rogue_get_info:type_name -> ChessRogueGetInfo
	5, // 4: ChessRogueQuitScRsp.player_info:type_name -> BLHOEGDLHNA
	6, // 5: ChessRogueQuitScRsp.rogue_aeon_info:type_name -> ILACICKKNOF
	7, // 6: ChessRogueQuitScRsp.level_info:type_name -> ChessRogueLevelInfo
	8, // 7: ChessRogueQuitScRsp.JDFLOJFCJMO:type_name -> AJPAMBEHJKM
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_ChessRogueQuitScRsp_proto_init() }
func file_ChessRogueQuitScRsp_proto_init() {
	if File_ChessRogueQuitScRsp_proto != nil {
		return
	}
	file_ChessRogueFinishInfo_proto_init()
	file_ChessRogueCurrentInfo_proto_init()
	file_ChessRogueGetInfo_proto_init()
	file_ChessRogueLevelInfo_proto_init()
	file_ChessRogueQueryGameInfo_proto_init()
	file_ILACICKKNOF_proto_init()
	file_AJPAMBEHJKM_proto_init()
	file_BLHOEGDLHNA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueQuitScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueQuitScRsp); i {
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
			RawDescriptor: file_ChessRogueQuitScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueQuitScRsp_proto_goTypes,
		DependencyIndexes: file_ChessRogueQuitScRsp_proto_depIdxs,
		MessageInfos:      file_ChessRogueQuitScRsp_proto_msgTypes,
	}.Build()
	File_ChessRogueQuitScRsp_proto = out.File
	file_ChessRogueQuitScRsp_proto_rawDesc = nil
	file_ChessRogueQuitScRsp_proto_goTypes = nil
	file_ChessRogueQuitScRsp_proto_depIdxs = nil
}
