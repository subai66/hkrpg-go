// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: RogueScoreRewardInfo.proto

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

type RogueScoreRewardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginTime            int64    `protobuf:"varint,2,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	HasTakenInitialScore bool     `protobuf:"varint,4,opt,name=has_taken_initial_score,json=hasTakenInitialScore,proto3" json:"has_taken_initial_score,omitempty"`
	PoolRefreshed        bool     `protobuf:"varint,9,opt,name=pool_refreshed,json=poolRefreshed,proto3" json:"pool_refreshed,omitempty"`
	PoolId               uint32   `protobuf:"varint,11,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	EndTime              int64    `protobuf:"varint,15,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	TakenScoreRewardList []uint32 `protobuf:"varint,1,rep,packed,name=taken_score_reward_list,json=takenScoreRewardList,proto3" json:"taken_score_reward_list,omitempty"`
	Score                uint32   `protobuf:"varint,8,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *RogueScoreRewardInfo) Reset() {
	*x = RogueScoreRewardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueScoreRewardInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueScoreRewardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueScoreRewardInfo) ProtoMessage() {}

func (x *RogueScoreRewardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueScoreRewardInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueScoreRewardInfo.ProtoReflect.Descriptor instead.
func (*RogueScoreRewardInfo) Descriptor() ([]byte, []int) {
	return file_RogueScoreRewardInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueScoreRewardInfo) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *RogueScoreRewardInfo) GetHasTakenInitialScore() bool {
	if x != nil {
		return x.HasTakenInitialScore
	}
	return false
}

func (x *RogueScoreRewardInfo) GetPoolRefreshed() bool {
	if x != nil {
		return x.PoolRefreshed
	}
	return false
}

func (x *RogueScoreRewardInfo) GetPoolId() uint32 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *RogueScoreRewardInfo) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *RogueScoreRewardInfo) GetTakenScoreRewardList() []uint32 {
	if x != nil {
		return x.TakenScoreRewardList
	}
	return nil
}

func (x *RogueScoreRewardInfo) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_RogueScoreRewardInfo_proto protoreflect.FileDescriptor

var file_RogueScoreRewardInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a,
	0x14, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x68, 0x61, 0x73, 0x5f, 0x74, 0x61, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x68, 0x61, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x6f, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x14, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueScoreRewardInfo_proto_rawDescOnce sync.Once
	file_RogueScoreRewardInfo_proto_rawDescData = file_RogueScoreRewardInfo_proto_rawDesc
)

func file_RogueScoreRewardInfo_proto_rawDescGZIP() []byte {
	file_RogueScoreRewardInfo_proto_rawDescOnce.Do(func() {
		file_RogueScoreRewardInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueScoreRewardInfo_proto_rawDescData)
	})
	return file_RogueScoreRewardInfo_proto_rawDescData
}

var file_RogueScoreRewardInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueScoreRewardInfo_proto_goTypes = []interface{}{
	(*RogueScoreRewardInfo)(nil), // 0: RogueScoreRewardInfo
}
var file_RogueScoreRewardInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueScoreRewardInfo_proto_init() }
func file_RogueScoreRewardInfo_proto_init() {
	if File_RogueScoreRewardInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueScoreRewardInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueScoreRewardInfo); i {
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
			RawDescriptor: file_RogueScoreRewardInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueScoreRewardInfo_proto_goTypes,
		DependencyIndexes: file_RogueScoreRewardInfo_proto_depIdxs,
		MessageInfos:      file_RogueScoreRewardInfo_proto_msgTypes,
	}.Build()
	File_RogueScoreRewardInfo_proto = out.File
	file_RogueScoreRewardInfo_proto_rawDesc = nil
	file_RogueScoreRewardInfo_proto_goTypes = nil
	file_RogueScoreRewardInfo_proto_depIdxs = nil
}
