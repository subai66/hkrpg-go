// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueCurrentInfo.proto

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

type ChessRogueCurrentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameMiracleInfo      *ChessRogueMiracleInfo    `protobuf:"bytes,8,opt,name=game_miracle_info,json=gameMiracleInfo,proto3" json:"game_miracle_info,omitempty"`
	RogueVirtualItem     *RogueVirtualItem         `protobuf:"bytes,4,opt,name=rogue_virtual_item,json=rogueVirtualItem,proto3" json:"rogue_virtual_item,omitempty"`
	RogueBuffInfo        *RogueTournBuffInfo       `protobuf:"bytes,6,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	JANOIDAPFMG          *IHKKKMAJKBK              `protobuf:"bytes,9,opt,name=JANOIDAPFMG,proto3" json:"JANOIDAPFMG,omitempty"`
	RogueSubMode         uint32                    `protobuf:"varint,14,opt,name=rogue_sub_mode,json=rogueSubMode,proto3" json:"rogue_sub_mode,omitempty"`
	RogueAeonInfo        *ILACICKKNOF              `protobuf:"bytes,11,opt,name=rogue_aeon_info,json=rogueAeonInfo,proto3" json:"rogue_aeon_info,omitempty"`
	StoryBuffInfo        *NNECGMAFOKB              `protobuf:"bytes,1,opt,name=story_buff_info,json=storyBuffInfo,proto3" json:"story_buff_info,omitempty"`
	NGMILHFCFHI          *IBJDAOLBIGO              `protobuf:"bytes,3,opt,name=NGMILHFCFHI,proto3" json:"NGMILHFCFHI,omitempty"`
	PAFILFHGKPE          *PPCDAKHJNOI              `protobuf:"bytes,10,opt,name=PAFILFHGKPE,proto3" json:"PAFILFHGKPE,omitempty"`
	RogueLineupInfo      *KDEDLOGEMEM              `protobuf:"bytes,13,opt,name=rogue_lineup_info,json=rogueLineupInfo,proto3" json:"rogue_lineup_info,omitempty"`
	IPJMAOIFIAE          []*MKNNGPDEFLB            `protobuf:"bytes,5,rep,name=IPJMAOIFIAE,proto3" json:"IPJMAOIFIAE,omitempty"`
	EBABEFMHIIJ          *FDBLDFEBBBI              `protobuf:"bytes,747,opt,name=EBABEFMHIIJ,proto3" json:"EBABEFMHIIJ,omitempty"`
	PendingAction        *RogueCommonPendingAction `protobuf:"bytes,15,opt,name=pending_action,json=pendingAction,proto3" json:"pending_action,omitempty"`
	LevelInfo            *ChessRogueLevelInfo      `protobuf:"bytes,7,opt,name=level_info,json=levelInfo,proto3" json:"level_info,omitempty"`
	RogueVirtualItemInfo *FJOCDCBHFHI              `protobuf:"bytes,12,opt,name=rogue_virtual_item_info,json=rogueVirtualItemInfo,proto3" json:"rogue_virtual_item_info,omitempty"`
}

func (x *ChessRogueCurrentInfo) Reset() {
	*x = ChessRogueCurrentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueCurrentInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueCurrentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueCurrentInfo) ProtoMessage() {}

func (x *ChessRogueCurrentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueCurrentInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueCurrentInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueCurrentInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueCurrentInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueCurrentInfo) GetGameMiracleInfo() *ChessRogueMiracleInfo {
	if x != nil {
		return x.GameMiracleInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetRogueVirtualItem() *RogueVirtualItem {
	if x != nil {
		return x.RogueVirtualItem
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetRogueBuffInfo() *RogueTournBuffInfo {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetJANOIDAPFMG() *IHKKKMAJKBK {
	if x != nil {
		return x.JANOIDAPFMG
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetRogueSubMode() uint32 {
	if x != nil {
		return x.RogueSubMode
	}
	return 0
}

func (x *ChessRogueCurrentInfo) GetRogueAeonInfo() *ILACICKKNOF {
	if x != nil {
		return x.RogueAeonInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetStoryBuffInfo() *NNECGMAFOKB {
	if x != nil {
		return x.StoryBuffInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetNGMILHFCFHI() *IBJDAOLBIGO {
	if x != nil {
		return x.NGMILHFCFHI
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetPAFILFHGKPE() *PPCDAKHJNOI {
	if x != nil {
		return x.PAFILFHGKPE
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetRogueLineupInfo() *KDEDLOGEMEM {
	if x != nil {
		return x.RogueLineupInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetIPJMAOIFIAE() []*MKNNGPDEFLB {
	if x != nil {
		return x.IPJMAOIFIAE
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetEBABEFMHIIJ() *FDBLDFEBBBI {
	if x != nil {
		return x.EBABEFMHIIJ
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetPendingAction() *RogueCommonPendingAction {
	if x != nil {
		return x.PendingAction
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetLevelInfo() *ChessRogueLevelInfo {
	if x != nil {
		return x.LevelInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetRogueVirtualItemInfo() *FJOCDCBHFHI {
	if x != nil {
		return x.RogueVirtualItemInfo
	}
	return nil
}

var File_ChessRogueCurrentInfo_proto protoreflect.FileDescriptor

var file_ChessRogueCurrentInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49,
	0x48, 0x4b, 0x4b, 0x4b, 0x4d, 0x41, 0x4a, 0x4b, 0x42, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x49, 0x42, 0x4a, 0x44, 0x41, 0x4f, 0x4c, 0x42, 0x49, 0x47, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x44, 0x42, 0x4c, 0x44, 0x46, 0x45, 0x42, 0x42, 0x42, 0x49,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x44, 0x45, 0x44, 0x4c, 0x4f, 0x47, 0x45,
	0x4d, 0x45, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x50, 0x50, 0x43, 0x44, 0x41, 0x4b, 0x48, 0x4a, 0x4e, 0x4f, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x4c,
	0x41, 0x43, 0x49, 0x43, 0x4b, 0x4b, 0x4e, 0x4f, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4e, 0x4e, 0x45, 0x43, 0x47, 0x4d, 0x41, 0x46, 0x4f, 0x4b, 0x42, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x4b, 0x4e, 0x4e, 0x47, 0x50, 0x44, 0x45, 0x46, 0x4c, 0x42, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x4a, 0x4f, 0x43, 0x44, 0x43, 0x42, 0x48, 0x46,
	0x48, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd2, 0x06, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x11,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0f, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x3f, 0x0a, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x3b, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e,
	0x0a, 0x0b, 0x4a, 0x41, 0x4e, 0x4f, 0x49, 0x44, 0x41, 0x50, 0x46, 0x4d, 0x47, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x48, 0x4b, 0x4b, 0x4b, 0x4d, 0x41, 0x4a, 0x4b, 0x42,
	0x4b, 0x52, 0x0b, 0x4a, 0x41, 0x4e, 0x4f, 0x49, 0x44, 0x41, 0x50, 0x46, 0x4d, 0x47, 0x12, 0x24,
	0x0a, 0x0e, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x75, 0x62,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x61, 0x65,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x49, 0x4c, 0x41, 0x43, 0x49, 0x43, 0x4b, 0x4b, 0x4e, 0x4f, 0x46, 0x52, 0x0d, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0f, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x4e, 0x45, 0x43, 0x47, 0x4d, 0x41, 0x46, 0x4f, 0x4b,
	0x42, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x47, 0x4d, 0x49, 0x4c, 0x48, 0x46, 0x43, 0x46, 0x48, 0x49, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x42, 0x4a, 0x44, 0x41, 0x4f, 0x4c, 0x42,
	0x49, 0x47, 0x4f, 0x52, 0x0b, 0x4e, 0x47, 0x4d, 0x49, 0x4c, 0x48, 0x46, 0x43, 0x46, 0x48, 0x49,
	0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x41, 0x46, 0x49, 0x4c, 0x46, 0x48, 0x47, 0x4b, 0x50, 0x45, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x50, 0x43, 0x44, 0x41, 0x4b, 0x48, 0x4a,
	0x4e, 0x4f, 0x49, 0x52, 0x0b, 0x50, 0x41, 0x46, 0x49, 0x4c, 0x46, 0x48, 0x47, 0x4b, 0x50, 0x45,
	0x12, 0x38, 0x0a, 0x11, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x44,
	0x45, 0x44, 0x4c, 0x4f, 0x47, 0x45, 0x4d, 0x45, 0x4d, 0x52, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x50,
	0x4a, 0x4d, 0x41, 0x4f, 0x49, 0x46, 0x49, 0x41, 0x45, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4d, 0x4b, 0x4e, 0x4e, 0x47, 0x50, 0x44, 0x45, 0x46, 0x4c, 0x42, 0x52, 0x0b, 0x49,
	0x50, 0x4a, 0x4d, 0x41, 0x4f, 0x49, 0x46, 0x49, 0x41, 0x45, 0x12, 0x2f, 0x0a, 0x0b, 0x45, 0x42,
	0x41, 0x42, 0x45, 0x46, 0x4d, 0x48, 0x49, 0x49, 0x4a, 0x18, 0xeb, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x46, 0x44, 0x42, 0x4c, 0x44, 0x46, 0x45, 0x42, 0x42, 0x42, 0x49, 0x52, 0x0b,
	0x45, 0x42, 0x41, 0x42, 0x45, 0x46, 0x4d, 0x48, 0x49, 0x49, 0x4a, 0x12, 0x40, 0x0a, 0x0e, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x43, 0x0a, 0x17, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x4a, 0x4f, 0x43, 0x44, 0x43, 0x42, 0x48, 0x46, 0x48,
	0x49, 0x52, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueCurrentInfo_proto_rawDescOnce sync.Once
	file_ChessRogueCurrentInfo_proto_rawDescData = file_ChessRogueCurrentInfo_proto_rawDesc
)

func file_ChessRogueCurrentInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueCurrentInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueCurrentInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueCurrentInfo_proto_rawDescData)
	})
	return file_ChessRogueCurrentInfo_proto_rawDescData
}

var file_ChessRogueCurrentInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueCurrentInfo_proto_goTypes = []interface{}{
	(*ChessRogueCurrentInfo)(nil),    // 0: ChessRogueCurrentInfo
	(*ChessRogueMiracleInfo)(nil),    // 1: ChessRogueMiracleInfo
	(*RogueVirtualItem)(nil),         // 2: RogueVirtualItem
	(*RogueTournBuffInfo)(nil),       // 3: RogueTournBuffInfo
	(*IHKKKMAJKBK)(nil),              // 4: IHKKKMAJKBK
	(*ILACICKKNOF)(nil),              // 5: ILACICKKNOF
	(*NNECGMAFOKB)(nil),              // 6: NNECGMAFOKB
	(*IBJDAOLBIGO)(nil),              // 7: IBJDAOLBIGO
	(*PPCDAKHJNOI)(nil),              // 8: PPCDAKHJNOI
	(*KDEDLOGEMEM)(nil),              // 9: KDEDLOGEMEM
	(*MKNNGPDEFLB)(nil),              // 10: MKNNGPDEFLB
	(*FDBLDFEBBBI)(nil),              // 11: FDBLDFEBBBI
	(*RogueCommonPendingAction)(nil), // 12: RogueCommonPendingAction
	(*ChessRogueLevelInfo)(nil),      // 13: ChessRogueLevelInfo
	(*FJOCDCBHFHI)(nil),              // 14: FJOCDCBHFHI
}
var file_ChessRogueCurrentInfo_proto_depIdxs = []int32{
	1,  // 0: ChessRogueCurrentInfo.game_miracle_info:type_name -> ChessRogueMiracleInfo
	2,  // 1: ChessRogueCurrentInfo.rogue_virtual_item:type_name -> RogueVirtualItem
	3,  // 2: ChessRogueCurrentInfo.rogue_buff_info:type_name -> RogueTournBuffInfo
	4,  // 3: ChessRogueCurrentInfo.JANOIDAPFMG:type_name -> IHKKKMAJKBK
	5,  // 4: ChessRogueCurrentInfo.rogue_aeon_info:type_name -> ILACICKKNOF
	6,  // 5: ChessRogueCurrentInfo.story_buff_info:type_name -> NNECGMAFOKB
	7,  // 6: ChessRogueCurrentInfo.NGMILHFCFHI:type_name -> IBJDAOLBIGO
	8,  // 7: ChessRogueCurrentInfo.PAFILFHGKPE:type_name -> PPCDAKHJNOI
	9,  // 8: ChessRogueCurrentInfo.rogue_lineup_info:type_name -> KDEDLOGEMEM
	10, // 9: ChessRogueCurrentInfo.IPJMAOIFIAE:type_name -> MKNNGPDEFLB
	11, // 10: ChessRogueCurrentInfo.EBABEFMHIIJ:type_name -> FDBLDFEBBBI
	12, // 11: ChessRogueCurrentInfo.pending_action:type_name -> RogueCommonPendingAction
	13, // 12: ChessRogueCurrentInfo.level_info:type_name -> ChessRogueLevelInfo
	14, // 13: ChessRogueCurrentInfo.rogue_virtual_item_info:type_name -> FJOCDCBHFHI
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_ChessRogueCurrentInfo_proto_init() }
func file_ChessRogueCurrentInfo_proto_init() {
	if File_ChessRogueCurrentInfo_proto != nil {
		return
	}
	file_IHKKKMAJKBK_proto_init()
	file_IBJDAOLBIGO_proto_init()
	file_FDBLDFEBBBI_proto_init()
	file_KDEDLOGEMEM_proto_init()
	file_RogueVirtualItem_proto_init()
	file_ChessRogueMiracleInfo_proto_init()
	file_PPCDAKHJNOI_proto_init()
	file_ChessRogueLevelInfo_proto_init()
	file_ILACICKKNOF_proto_init()
	file_NNECGMAFOKB_proto_init()
	file_MKNNGPDEFLB_proto_init()
	file_RogueCommonPendingAction_proto_init()
	file_FJOCDCBHFHI_proto_init()
	file_RogueTournBuffInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueCurrentInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueCurrentInfo); i {
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
			RawDescriptor: file_ChessRogueCurrentInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueCurrentInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueCurrentInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueCurrentInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueCurrentInfo_proto = out.File
	file_ChessRogueCurrentInfo_proto_rawDesc = nil
	file_ChessRogueCurrentInfo_proto_goTypes = nil
	file_ChessRogueCurrentInfo_proto_depIdxs = nil
}
