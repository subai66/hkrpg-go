// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueAction.proto

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

type RogueAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*RogueAction_BuffSelectInfo
	//	*RogueAction_IJEMJHPDFBF
	//	*RogueAction_PEHHNCDHIGN
	//	*RogueAction_BPMFHEFCINF
	//	*RogueAction_MiracleSelectInfo
	//	*RogueAction_INABPAILEPD
	//	*RogueAction_EALLGGDLGHP
	//	*RogueAction_HCDDAKNPDIB
	//	*RogueAction_EIJJLAIBPOJ
	//	*RogueAction_PJFEMEFHHJG
	//	*RogueAction_LCKAGOEAHHA
	//	*RogueAction_LOJIOMDFFFB
	//	*RogueAction_BonusSelectInfo
	//	*RogueAction_RogueFormulaSelectInfo
	//	*RogueAction_DIODEFICIAG
	//	*RogueAction_FGMNAINGFOL
	Action isRogueAction_Action `protobuf_oneof:"action"`
}

func (x *RogueAction) Reset() {
	*x = RogueAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueAction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueAction) ProtoMessage() {}

func (x *RogueAction) ProtoReflect() protoreflect.Message {
	mi := &file_RogueAction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueAction.ProtoReflect.Descriptor instead.
func (*RogueAction) Descriptor() ([]byte, []int) {
	return file_RogueAction_proto_rawDescGZIP(), []int{0}
}

func (m *RogueAction) GetAction() isRogueAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *RogueAction) GetBuffSelectInfo() *RogueCommonBuffSelectInfo {
	if x, ok := x.GetAction().(*RogueAction_BuffSelectInfo); ok {
		return x.BuffSelectInfo
	}
	return nil
}

func (x *RogueAction) GetIJEMJHPDFBF() *DOCGLJMHMIF {
	if x, ok := x.GetAction().(*RogueAction_IJEMJHPDFBF); ok {
		return x.IJEMJHPDFBF
	}
	return nil
}

func (x *RogueAction) GetPEHHNCDHIGN() *PJFEGMHPDPO {
	if x, ok := x.GetAction().(*RogueAction_PEHHNCDHIGN); ok {
		return x.PEHHNCDHIGN
	}
	return nil
}

func (x *RogueAction) GetBPMFHEFCINF() *IDCEIJBPDAL {
	if x, ok := x.GetAction().(*RogueAction_BPMFHEFCINF); ok {
		return x.BPMFHEFCINF
	}
	return nil
}

func (x *RogueAction) GetMiracleSelectInfo() *RogueMiracleSelectInfo {
	if x, ok := x.GetAction().(*RogueAction_MiracleSelectInfo); ok {
		return x.MiracleSelectInfo
	}
	return nil
}

func (x *RogueAction) GetINABPAILEPD() *GNFLODHEKJJ {
	if x, ok := x.GetAction().(*RogueAction_INABPAILEPD); ok {
		return x.INABPAILEPD
	}
	return nil
}

func (x *RogueAction) GetEALLGGDLGHP() *MEFJJENNJDH {
	if x, ok := x.GetAction().(*RogueAction_EALLGGDLGHP); ok {
		return x.EALLGGDLGHP
	}
	return nil
}

func (x *RogueAction) GetHCDDAKNPDIB() *PBBDLCGHLBP {
	if x, ok := x.GetAction().(*RogueAction_HCDDAKNPDIB); ok {
		return x.HCDDAKNPDIB
	}
	return nil
}

func (x *RogueAction) GetEIJJLAIBPOJ() *BHEIOIMNCJJ {
	if x, ok := x.GetAction().(*RogueAction_EIJJLAIBPOJ); ok {
		return x.EIJJLAIBPOJ
	}
	return nil
}

func (x *RogueAction) GetPJFEMEFHHJG() *HAKKLGAILBN {
	if x, ok := x.GetAction().(*RogueAction_PJFEMEFHHJG); ok {
		return x.PJFEMEFHHJG
	}
	return nil
}

func (x *RogueAction) GetLCKAGOEAHHA() *PCAFNPOEDCF {
	if x, ok := x.GetAction().(*RogueAction_LCKAGOEAHHA); ok {
		return x.LCKAGOEAHHA
	}
	return nil
}

func (x *RogueAction) GetLOJIOMDFFFB() *NGFPPDBIOKJ {
	if x, ok := x.GetAction().(*RogueAction_LOJIOMDFFFB); ok {
		return x.LOJIOMDFFFB
	}
	return nil
}

func (x *RogueAction) GetBonusSelectInfo() *RogueBonusSelectInfo {
	if x, ok := x.GetAction().(*RogueAction_BonusSelectInfo); ok {
		return x.BonusSelectInfo
	}
	return nil
}

func (x *RogueAction) GetRogueFormulaSelectInfo() *RogueFormulaSelectInfo {
	if x, ok := x.GetAction().(*RogueAction_RogueFormulaSelectInfo); ok {
		return x.RogueFormulaSelectInfo
	}
	return nil
}

func (x *RogueAction) GetDIODEFICIAG() *CEGHNKFLKEE {
	if x, ok := x.GetAction().(*RogueAction_DIODEFICIAG); ok {
		return x.DIODEFICIAG
	}
	return nil
}

func (x *RogueAction) GetFGMNAINGFOL() *HCPIEGICMEG {
	if x, ok := x.GetAction().(*RogueAction_FGMNAINGFOL); ok {
		return x.FGMNAINGFOL
	}
	return nil
}

type isRogueAction_Action interface {
	isRogueAction_Action()
}

type RogueAction_BuffSelectInfo struct {
	BuffSelectInfo *RogueCommonBuffSelectInfo `protobuf:"bytes,216,opt,name=buff_select_info,json=buffSelectInfo,proto3,oneof"`
}

type RogueAction_IJEMJHPDFBF struct {
	IJEMJHPDFBF *DOCGLJMHMIF `protobuf:"bytes,550,opt,name=IJEMJHPDFBF,proto3,oneof"`
}

type RogueAction_PEHHNCDHIGN struct {
	PEHHNCDHIGN *PJFEGMHPDPO `protobuf:"bytes,1977,opt,name=PEHHNCDHIGN,proto3,oneof"`
}

type RogueAction_BPMFHEFCINF struct {
	BPMFHEFCINF *IDCEIJBPDAL `protobuf:"bytes,402,opt,name=BPMFHEFCINF,proto3,oneof"`
}

type RogueAction_MiracleSelectInfo struct {
	MiracleSelectInfo *RogueMiracleSelectInfo `protobuf:"bytes,1920,opt,name=miracle_select_info,json=miracleSelectInfo,proto3,oneof"`
}

type RogueAction_INABPAILEPD struct {
	INABPAILEPD *GNFLODHEKJJ `protobuf:"bytes,1781,opt,name=INABPAILEPD,proto3,oneof"`
}

type RogueAction_EALLGGDLGHP struct {
	EALLGGDLGHP *MEFJJENNJDH `protobuf:"bytes,1485,opt,name=EALLGGDLGHP,proto3,oneof"`
}

type RogueAction_HCDDAKNPDIB struct {
	HCDDAKNPDIB *PBBDLCGHLBP `protobuf:"bytes,433,opt,name=HCDDAKNPDIB,proto3,oneof"`
}

type RogueAction_EIJJLAIBPOJ struct {
	EIJJLAIBPOJ *BHEIOIMNCJJ `protobuf:"bytes,747,opt,name=EIJJLAIBPOJ,proto3,oneof"`
}

type RogueAction_PJFEMEFHHJG struct {
	PJFEMEFHHJG *HAKKLGAILBN `protobuf:"bytes,597,opt,name=PJFEMEFHHJG,proto3,oneof"`
}

type RogueAction_LCKAGOEAHHA struct {
	LCKAGOEAHHA *PCAFNPOEDCF `protobuf:"bytes,418,opt,name=LCKAGOEAHHA,proto3,oneof"`
}

type RogueAction_LOJIOMDFFFB struct {
	LOJIOMDFFFB *NGFPPDBIOKJ `protobuf:"bytes,1925,opt,name=LOJIOMDFFFB,proto3,oneof"`
}

type RogueAction_BonusSelectInfo struct {
	BonusSelectInfo *RogueBonusSelectInfo `protobuf:"bytes,388,opt,name=bonus_select_info,json=bonusSelectInfo,proto3,oneof"`
}

type RogueAction_RogueFormulaSelectInfo struct {
	RogueFormulaSelectInfo *RogueFormulaSelectInfo `protobuf:"bytes,485,opt,name=rogue_formula_select_info,json=rogueFormulaSelectInfo,proto3,oneof"`
}

type RogueAction_DIODEFICIAG struct {
	DIODEFICIAG *CEGHNKFLKEE `protobuf:"bytes,1636,opt,name=DIODEFICIAG,proto3,oneof"`
}

type RogueAction_FGMNAINGFOL struct {
	FGMNAINGFOL *HCPIEGICMEG `protobuf:"bytes,575,opt,name=FGMNAINGFOL,proto3,oneof"`
}

func (*RogueAction_BuffSelectInfo) isRogueAction_Action() {}

func (*RogueAction_IJEMJHPDFBF) isRogueAction_Action() {}

func (*RogueAction_PEHHNCDHIGN) isRogueAction_Action() {}

func (*RogueAction_BPMFHEFCINF) isRogueAction_Action() {}

func (*RogueAction_MiracleSelectInfo) isRogueAction_Action() {}

func (*RogueAction_INABPAILEPD) isRogueAction_Action() {}

func (*RogueAction_EALLGGDLGHP) isRogueAction_Action() {}

func (*RogueAction_HCDDAKNPDIB) isRogueAction_Action() {}

func (*RogueAction_EIJJLAIBPOJ) isRogueAction_Action() {}

func (*RogueAction_PJFEMEFHHJG) isRogueAction_Action() {}

func (*RogueAction_LCKAGOEAHHA) isRogueAction_Action() {}

func (*RogueAction_LOJIOMDFFFB) isRogueAction_Action() {}

func (*RogueAction_BonusSelectInfo) isRogueAction_Action() {}

func (*RogueAction_RogueFormulaSelectInfo) isRogueAction_Action() {}

func (*RogueAction_DIODEFICIAG) isRogueAction_Action() {}

func (*RogueAction_FGMNAINGFOL) isRogueAction_Action() {}

var File_RogueAction_proto protoreflect.FileDescriptor

var file_RogueAction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x44, 0x4f, 0x43, 0x47, 0x4c, 0x4a, 0x4d, 0x48, 0x4d, 0x49, 0x46, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x42, 0x48, 0x45, 0x49, 0x4f, 0x49, 0x4d, 0x4e, 0x43, 0x4a, 0x4a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x47, 0x46, 0x50, 0x50, 0x44, 0x42, 0x49, 0x4f,
	0x4b, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4e, 0x46, 0x4c, 0x4f, 0x44,
	0x48, 0x45, 0x4b, 0x4a, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x45, 0x46,
	0x4a, 0x4a, 0x45, 0x4e, 0x4e, 0x4a, 0x44, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x48, 0x41, 0x4b, 0x4b, 0x4c, 0x47, 0x41, 0x49, 0x4c, 0x42, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x50, 0x43, 0x41, 0x46, 0x4e, 0x50, 0x4f, 0x45, 0x44, 0x43, 0x46, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x42, 0x42, 0x44, 0x4c, 0x43, 0x47, 0x48, 0x4c, 0x42,
	0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x45, 0x47, 0x48, 0x4e, 0x4b, 0x46, 0x4c,
	0x4b, 0x45, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x44, 0x43, 0x45, 0x49, 0x4a, 0x42,
	0x50, 0x44, 0x41, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x4a, 0x46, 0x45,
	0x47, 0x4d, 0x48, 0x50, 0x44, 0x50, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48,
	0x43, 0x50, 0x49, 0x45, 0x47, 0x49, 0x43, 0x4d, 0x45, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xad, 0x07, 0x0a, 0x0b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x47, 0x0a, 0x10, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0xd8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0b, 0x49, 0x4a, 0x45,
	0x4d, 0x4a, 0x48, 0x50, 0x44, 0x46, 0x42, 0x46, 0x18, 0xa6, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x44, 0x4f, 0x43, 0x47, 0x4c, 0x4a, 0x4d, 0x48, 0x4d, 0x49, 0x46, 0x48, 0x00, 0x52,
	0x0b, 0x49, 0x4a, 0x45, 0x4d, 0x4a, 0x48, 0x50, 0x44, 0x46, 0x42, 0x46, 0x12, 0x31, 0x0a, 0x0b,
	0x50, 0x45, 0x48, 0x48, 0x4e, 0x43, 0x44, 0x48, 0x49, 0x47, 0x4e, 0x18, 0xb9, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x4a, 0x46, 0x45, 0x47, 0x4d, 0x48, 0x50, 0x44, 0x50, 0x4f,
	0x48, 0x00, 0x52, 0x0b, 0x50, 0x45, 0x48, 0x48, 0x4e, 0x43, 0x44, 0x48, 0x49, 0x47, 0x4e, 0x12,
	0x31, 0x0a, 0x0b, 0x42, 0x50, 0x4d, 0x46, 0x48, 0x45, 0x46, 0x43, 0x49, 0x4e, 0x46, 0x18, 0x92,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x44, 0x43, 0x45, 0x49, 0x4a, 0x42, 0x50,
	0x44, 0x41, 0x4c, 0x48, 0x00, 0x52, 0x0b, 0x42, 0x50, 0x4d, 0x46, 0x48, 0x45, 0x46, 0x43, 0x49,
	0x4e, 0x46, 0x12, 0x4a, 0x0a, 0x13, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x80, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31,
	0x0a, 0x0b, 0x49, 0x4e, 0x41, 0x42, 0x50, 0x41, 0x49, 0x4c, 0x45, 0x50, 0x44, 0x18, 0xf5, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x4e, 0x46, 0x4c, 0x4f, 0x44, 0x48, 0x45, 0x4b,
	0x4a, 0x4a, 0x48, 0x00, 0x52, 0x0b, 0x49, 0x4e, 0x41, 0x42, 0x50, 0x41, 0x49, 0x4c, 0x45, 0x50,
	0x44, 0x12, 0x31, 0x0a, 0x0b, 0x45, 0x41, 0x4c, 0x4c, 0x47, 0x47, 0x44, 0x4c, 0x47, 0x48, 0x50,
	0x18, 0xcd, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x45, 0x46, 0x4a, 0x4a, 0x45,
	0x4e, 0x4e, 0x4a, 0x44, 0x48, 0x48, 0x00, 0x52, 0x0b, 0x45, 0x41, 0x4c, 0x4c, 0x47, 0x47, 0x44,
	0x4c, 0x47, 0x48, 0x50, 0x12, 0x31, 0x0a, 0x0b, 0x48, 0x43, 0x44, 0x44, 0x41, 0x4b, 0x4e, 0x50,
	0x44, 0x49, 0x42, 0x18, 0xb1, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x42, 0x42,
	0x44, 0x4c, 0x43, 0x47, 0x48, 0x4c, 0x42, 0x50, 0x48, 0x00, 0x52, 0x0b, 0x48, 0x43, 0x44, 0x44,
	0x41, 0x4b, 0x4e, 0x50, 0x44, 0x49, 0x42, 0x12, 0x31, 0x0a, 0x0b, 0x45, 0x49, 0x4a, 0x4a, 0x4c,
	0x41, 0x49, 0x42, 0x50, 0x4f, 0x4a, 0x18, 0xeb, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x42, 0x48, 0x45, 0x49, 0x4f, 0x49, 0x4d, 0x4e, 0x43, 0x4a, 0x4a, 0x48, 0x00, 0x52, 0x0b, 0x45,
	0x49, 0x4a, 0x4a, 0x4c, 0x41, 0x49, 0x42, 0x50, 0x4f, 0x4a, 0x12, 0x31, 0x0a, 0x0b, 0x50, 0x4a,
	0x46, 0x45, 0x4d, 0x45, 0x46, 0x48, 0x48, 0x4a, 0x47, 0x18, 0xd5, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x48, 0x41, 0x4b, 0x4b, 0x4c, 0x47, 0x41, 0x49, 0x4c, 0x42, 0x4e, 0x48, 0x00,
	0x52, 0x0b, 0x50, 0x4a, 0x46, 0x45, 0x4d, 0x45, 0x46, 0x48, 0x48, 0x4a, 0x47, 0x12, 0x31, 0x0a,
	0x0b, 0x4c, 0x43, 0x4b, 0x41, 0x47, 0x4f, 0x45, 0x41, 0x48, 0x48, 0x41, 0x18, 0xa2, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x43, 0x41, 0x46, 0x4e, 0x50, 0x4f, 0x45, 0x44, 0x43,
	0x46, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x43, 0x4b, 0x41, 0x47, 0x4f, 0x45, 0x41, 0x48, 0x48, 0x41,
	0x12, 0x31, 0x0a, 0x0b, 0x4c, 0x4f, 0x4a, 0x49, 0x4f, 0x4d, 0x44, 0x46, 0x46, 0x46, 0x42, 0x18,
	0x85, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x47, 0x46, 0x50, 0x50, 0x44, 0x42,
	0x49, 0x4f, 0x4b, 0x4a, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x4f, 0x4a, 0x49, 0x4f, 0x4d, 0x44, 0x46,
	0x46, 0x46, 0x42, 0x12, 0x44, 0x0a, 0x11, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x84, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x55, 0x0a, 0x19, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xe5, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x16, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x31, 0x0a, 0x0b, 0x44, 0x49, 0x4f, 0x44, 0x45, 0x46, 0x49, 0x43, 0x49, 0x41, 0x47, 0x18,
	0xe4, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x45, 0x47, 0x48, 0x4e, 0x4b, 0x46,
	0x4c, 0x4b, 0x45, 0x45, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x49, 0x4f, 0x44, 0x45, 0x46, 0x49, 0x43,
	0x49, 0x41, 0x47, 0x12, 0x31, 0x0a, 0x0b, 0x46, 0x47, 0x4d, 0x4e, 0x41, 0x49, 0x4e, 0x47, 0x46,
	0x4f, 0x4c, 0x18, 0xbf, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x43, 0x50, 0x49,
	0x45, 0x47, 0x49, 0x43, 0x4d, 0x45, 0x47, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x47, 0x4d, 0x4e, 0x41,
	0x49, 0x4e, 0x47, 0x46, 0x4f, 0x4c, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_RogueAction_proto_rawDescOnce sync.Once
	file_RogueAction_proto_rawDescData = file_RogueAction_proto_rawDesc
)

func file_RogueAction_proto_rawDescGZIP() []byte {
	file_RogueAction_proto_rawDescOnce.Do(func() {
		file_RogueAction_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueAction_proto_rawDescData)
	})
	return file_RogueAction_proto_rawDescData
}

var file_RogueAction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueAction_proto_goTypes = []interface{}{
	(*RogueAction)(nil),               // 0: RogueAction
	(*RogueCommonBuffSelectInfo)(nil), // 1: RogueCommonBuffSelectInfo
	(*DOCGLJMHMIF)(nil),               // 2: DOCGLJMHMIF
	(*PJFEGMHPDPO)(nil),               // 3: PJFEGMHPDPO
	(*IDCEIJBPDAL)(nil),               // 4: IDCEIJBPDAL
	(*RogueMiracleSelectInfo)(nil),    // 5: RogueMiracleSelectInfo
	(*GNFLODHEKJJ)(nil),               // 6: GNFLODHEKJJ
	(*MEFJJENNJDH)(nil),               // 7: MEFJJENNJDH
	(*PBBDLCGHLBP)(nil),               // 8: PBBDLCGHLBP
	(*BHEIOIMNCJJ)(nil),               // 9: BHEIOIMNCJJ
	(*HAKKLGAILBN)(nil),               // 10: HAKKLGAILBN
	(*PCAFNPOEDCF)(nil),               // 11: PCAFNPOEDCF
	(*NGFPPDBIOKJ)(nil),               // 12: NGFPPDBIOKJ
	(*RogueBonusSelectInfo)(nil),      // 13: RogueBonusSelectInfo
	(*RogueFormulaSelectInfo)(nil),    // 14: RogueFormulaSelectInfo
	(*CEGHNKFLKEE)(nil),               // 15: CEGHNKFLKEE
	(*HCPIEGICMEG)(nil),               // 16: HCPIEGICMEG
}
var file_RogueAction_proto_depIdxs = []int32{
	1,  // 0: RogueAction.buff_select_info:type_name -> RogueCommonBuffSelectInfo
	2,  // 1: RogueAction.IJEMJHPDFBF:type_name -> DOCGLJMHMIF
	3,  // 2: RogueAction.PEHHNCDHIGN:type_name -> PJFEGMHPDPO
	4,  // 3: RogueAction.BPMFHEFCINF:type_name -> IDCEIJBPDAL
	5,  // 4: RogueAction.miracle_select_info:type_name -> RogueMiracleSelectInfo
	6,  // 5: RogueAction.INABPAILEPD:type_name -> GNFLODHEKJJ
	7,  // 6: RogueAction.EALLGGDLGHP:type_name -> MEFJJENNJDH
	8,  // 7: RogueAction.HCDDAKNPDIB:type_name -> PBBDLCGHLBP
	9,  // 8: RogueAction.EIJJLAIBPOJ:type_name -> BHEIOIMNCJJ
	10, // 9: RogueAction.PJFEMEFHHJG:type_name -> HAKKLGAILBN
	11, // 10: RogueAction.LCKAGOEAHHA:type_name -> PCAFNPOEDCF
	12, // 11: RogueAction.LOJIOMDFFFB:type_name -> NGFPPDBIOKJ
	13, // 12: RogueAction.bonus_select_info:type_name -> RogueBonusSelectInfo
	14, // 13: RogueAction.rogue_formula_select_info:type_name -> RogueFormulaSelectInfo
	15, // 14: RogueAction.DIODEFICIAG:type_name -> CEGHNKFLKEE
	16, // 15: RogueAction.FGMNAINGFOL:type_name -> HCPIEGICMEG
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_RogueAction_proto_init() }
func file_RogueAction_proto_init() {
	if File_RogueAction_proto != nil {
		return
	}
	file_RogueBonusSelectInfo_proto_init()
	file_DOCGLJMHMIF_proto_init()
	file_BHEIOIMNCJJ_proto_init()
	file_NGFPPDBIOKJ_proto_init()
	file_GNFLODHEKJJ_proto_init()
	file_MEFJJENNJDH_proto_init()
	file_HAKKLGAILBN_proto_init()
	file_PCAFNPOEDCF_proto_init()
	file_PBBDLCGHLBP_proto_init()
	file_RogueCommonBuffSelectInfo_proto_init()
	file_RogueFormulaSelectInfo_proto_init()
	file_CEGHNKFLKEE_proto_init()
	file_RogueMiracleSelectInfo_proto_init()
	file_IDCEIJBPDAL_proto_init()
	file_PJFEGMHPDPO_proto_init()
	file_HCPIEGICMEG_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueAction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueAction); i {
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
	file_RogueAction_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RogueAction_BuffSelectInfo)(nil),
		(*RogueAction_IJEMJHPDFBF)(nil),
		(*RogueAction_PEHHNCDHIGN)(nil),
		(*RogueAction_BPMFHEFCINF)(nil),
		(*RogueAction_MiracleSelectInfo)(nil),
		(*RogueAction_INABPAILEPD)(nil),
		(*RogueAction_EALLGGDLGHP)(nil),
		(*RogueAction_HCDDAKNPDIB)(nil),
		(*RogueAction_EIJJLAIBPOJ)(nil),
		(*RogueAction_PJFEMEFHHJG)(nil),
		(*RogueAction_LCKAGOEAHHA)(nil),
		(*RogueAction_LOJIOMDFFFB)(nil),
		(*RogueAction_BonusSelectInfo)(nil),
		(*RogueAction_RogueFormulaSelectInfo)(nil),
		(*RogueAction_DIODEFICIAG)(nil),
		(*RogueAction_FGMNAINGFOL)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueAction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueAction_proto_goTypes,
		DependencyIndexes: file_RogueAction_proto_depIdxs,
		MessageInfos:      file_RogueAction_proto_msgTypes,
	}.Build()
	File_RogueAction_proto = out.File
	file_RogueAction_proto_rawDesc = nil
	file_RogueAction_proto_goTypes = nil
	file_RogueAction_proto_depIdxs = nil
}
