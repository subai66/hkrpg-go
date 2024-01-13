// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SceneEntityMoveScNotify.proto

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

type SceneEntityMoveScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId          uint32      `protobuf:"varint,11,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	EntityId         uint32      `protobuf:"varint,8,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	ClientPosVersion uint32      `protobuf:"varint,5,opt,name=client_pos_version,json=clientPosVersion,proto3" json:"client_pos_version,omitempty"`
	Motion           *MotionInfo `protobuf:"bytes,12,opt,name=motion,proto3" json:"motion,omitempty"`
}

func (x *SceneEntityMoveScNotify) Reset() {
	*x = SceneEntityMoveScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityMoveScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityMoveScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityMoveScNotify) ProtoMessage() {}

func (x *SceneEntityMoveScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityMoveScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityMoveScNotify.ProtoReflect.Descriptor instead.
func (*SceneEntityMoveScNotify) Descriptor() ([]byte, []int) {
	return file_SceneEntityMoveScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityMoveScNotify) GetEntryId() uint32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *SceneEntityMoveScNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *SceneEntityMoveScNotify) GetClientPosVersion() uint32 {
	if x != nil {
		return x.ClientPosVersion
	}
	return 0
}

func (x *SceneEntityMoveScNotify) GetMotion() *MotionInfo {
	if x != nil {
		return x.Motion
	}
	return nil
}

var File_SceneEntityMoveScNotify_proto protoreflect.FileDescriptor

var file_SceneEntityMoveScNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76,
	0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x17, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneEntityMoveScNotify_proto_rawDescOnce sync.Once
	file_SceneEntityMoveScNotify_proto_rawDescData = file_SceneEntityMoveScNotify_proto_rawDesc
)

func file_SceneEntityMoveScNotify_proto_rawDescGZIP() []byte {
	file_SceneEntityMoveScNotify_proto_rawDescOnce.Do(func() {
		file_SceneEntityMoveScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityMoveScNotify_proto_rawDescData)
	})
	return file_SceneEntityMoveScNotify_proto_rawDescData
}

var file_SceneEntityMoveScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityMoveScNotify_proto_goTypes = []interface{}{
	(*SceneEntityMoveScNotify)(nil), // 0: SceneEntityMoveScNotify
	(*MotionInfo)(nil),              // 1: MotionInfo
}
var file_SceneEntityMoveScNotify_proto_depIdxs = []int32{
	1, // 0: SceneEntityMoveScNotify.motion:type_name -> MotionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneEntityMoveScNotify_proto_init() }
func file_SceneEntityMoveScNotify_proto_init() {
	if File_SceneEntityMoveScNotify_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityMoveScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityMoveScNotify); i {
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
			RawDescriptor: file_SceneEntityMoveScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityMoveScNotify_proto_goTypes,
		DependencyIndexes: file_SceneEntityMoveScNotify_proto_depIdxs,
		MessageInfos:      file_SceneEntityMoveScNotify_proto_msgTypes,
	}.Build()
	File_SceneEntityMoveScNotify_proto = out.File
	file_SceneEntityMoveScNotify_proto_rawDesc = nil
	file_SceneEntityMoveScNotify_proto_goTypes = nil
	file_SceneEntityMoveScNotify_proto_depIdxs = nil
}
