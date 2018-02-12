// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pez_desk.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_req_kickout from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_bc_leaveTimeout from common_client.proto

// Ignoring public import of common_desk_by_agent from common_client.proto

// Ignoring public import of common_req_list_coin_desk from common_client.proto

// Ignoring public import of common_ack_list_coin_desk from common_client.proto

// Ignoring public import of CommonCoinDeskInfo from common_client.proto

// Ignoring public import of common_req_list_coin_info from common_client.proto

// Ignoring public import of common_ack_list_coin_info from common_client.proto

// Ignoring public import of CommonCoinLevelInfo from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMCARD_BILL_TYPE from common_enum.proto

// Ignoring public import of pez_base_PaiInfo from pez_base.proto

// Ignoring public import of pez_base_PlayConf from pez_base.proto

// Ignoring public import of pez_base_RoomTypeInfo from pez_base.proto

// Ignoring public import of pez_base_timerInfo from pez_base.proto

// Ignoring public import of pez_base_PaiValue from pez_base.proto

// Ignoring public import of pez_base_PlayerCard from pez_base.proto

// Ignoring public import of pez_lastScore from pez_base.proto

// Ignoring public import of pez_base_PlayerInfo from pez_base.proto

// Ignoring public import of pez_base_DeskGameInfo from pez_base.proto

// Ignoring public import of pez_tuozi from pez_base.proto

// Ignoring public import of pez_enum_protoId from pez_base.proto

// Ignoring public import of pez_enum_PEZTYPE from pez_base.proto

// Ignoring public import of pez_enum_ErrorCode from pez_base.proto

// Ignoring public import of pez_enum_Option from pez_base.proto

// Ignoring public import of pez_RoomType from pez_base.proto

// Ignoring public import of pez_enum_mjColor from pez_base.proto

// Ignoring public import of pez_enum_PaiType from pez_base.proto

// Ignoring public import of pez_enum_Base from pez_base.proto

// Ignoring public import of pez_enum_Bet from pez_base.proto

// Ignoring public import of pez_enum_UserGameStatus from pez_base.proto

// Ignoring public import of pez_enum_DeskGameStatus from pez_base.proto

// Ignoring public import of pez_enum_TuoziType from pez_base.proto

// 房主解散房间(未开局)
type Pez_DissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_DissolveDesk) Reset()                    { *m = Pez_DissolveDesk{} }
func (m *Pez_DissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Pez_DissolveDesk) ProtoMessage()               {}
func (*Pez_DissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{0} }

func (m *Pez_DissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_DissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 解散房间回复
type Pez_AckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=passWord" json:"passWord,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckDissolveDesk) Reset()                    { *m = Pez_AckDissolveDesk{} }
func (m *Pez_AckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckDissolveDesk) ProtoMessage()               {}
func (*Pez_AckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{1} }

func (m *Pez_AckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_AckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *Pez_AckDissolveDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

// 准备游戏
type Pez_Ready struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_Ready) Reset()                    { *m = Pez_Ready{} }
func (m *Pez_Ready) String() string            { return proto.CompactTextString(m) }
func (*Pez_Ready) ProtoMessage()               {}
func (*Pez_Ready) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{2} }

func (m *Pez_Ready) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_Ready) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 准备游戏的结果
type Pez_AckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckReady) Reset()                    { *m = Pez_AckReady{} }
func (m *Pez_AckReady) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckReady) ProtoMessage()               {}
func (*Pez_AckReady) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{3} }

func (m *Pez_AckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *Pez_AckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type Pez_EndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	IsbigWin         *bool   `protobuf:"varint,3,opt,name=isbigWin" json:"isbigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	BaseValue        *int32  `protobuf:"varint,5,opt,name=baseValue" json:"baseValue,omitempty"`
	WinScore         *int64  `protobuf:"varint,6,opt,name=winScore" json:"winScore,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_EndLotteryInfo) Reset()                    { *m = Pez_EndLotteryInfo{} }
func (m *Pez_EndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*Pez_EndLotteryInfo) ProtoMessage()               {}
func (*Pez_EndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{4} }

func (m *Pez_EndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_EndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *Pez_EndLotteryInfo) GetIsbigWin() bool {
	if m != nil && m.IsbigWin != nil {
		return *m.IsbigWin
	}
	return false
}

func (m *Pez_EndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *Pez_EndLotteryInfo) GetBaseValue() int32 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *Pez_EndLotteryInfo) GetWinScore() int64 {
	if m != nil && m.WinScore != nil {
		return *m.WinScore
	}
	return 0
}

// 本局结果(广播)
type Pez_SendCurrentResult struct {
	CurrentRound     *int32          `protobuf:"varint,1,opt,name=currentRound" json:"currentRound,omitempty"`
	CurrBill         []*Pez_UserBean `protobuf:"bytes,2,rep,name=CurrBill" json:"CurrBill,omitempty"`
	BoardsCout       *int32          `protobuf:"varint,3,opt,name=boardsCout" json:"boardsCout,omitempty"`
	CardsID          []int32         `protobuf:"varint,4,rep,name=cardsID" json:"cardsID,omitempty"`
	LanDiScore       *int64          `protobuf:"varint,5,opt,name=lanDiScore" json:"lanDiScore,omitempty"`
	OpenResult       *int32          `protobuf:"varint,6,opt,name=openResult" json:"openResult,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Pez_SendCurrentResult) Reset()                    { *m = Pez_SendCurrentResult{} }
func (m *Pez_SendCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendCurrentResult) ProtoMessage()               {}
func (*Pez_SendCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{5} }

func (m *Pez_SendCurrentResult) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *Pez_SendCurrentResult) GetCurrBill() []*Pez_UserBean {
	if m != nil {
		return m.CurrBill
	}
	return nil
}

func (m *Pez_SendCurrentResult) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *Pez_SendCurrentResult) GetCardsID() []int32 {
	if m != nil {
		return m.CardsID
	}
	return nil
}

func (m *Pez_SendCurrentResult) GetLanDiScore() int64 {
	if m != nil && m.LanDiScore != nil {
		return *m.LanDiScore
	}
	return 0
}

func (m *Pez_SendCurrentResult) GetOpenResult() int32 {
	if m != nil && m.OpenResult != nil {
		return *m.OpenResult
	}
	return 0
}

type EndLottery struct {
	UserId           *uint32     `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	IsbigWin         *bool       `protobuf:"varint,2,opt,name=isbigWin" json:"isbigWin,omitempty"`
	IsOwner          *bool       `protobuf:"varint,3,opt,name=isOwner" json:"isOwner,omitempty"`
	WinScore         *int64      `protobuf:"varint,4,opt,name=winScore" json:"winScore,omitempty"`
	WxInfo           *WeixinInfo `protobuf:"bytes,5,opt,name=wxInfo" json:"wxInfo,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *EndLottery) Reset()                    { *m = EndLottery{} }
func (m *EndLottery) String() string            { return proto.CompactTextString(m) }
func (*EndLottery) ProtoMessage()               {}
func (*EndLottery) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{6} }

func (m *EndLottery) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *EndLottery) GetIsbigWin() bool {
	if m != nil && m.IsbigWin != nil {
		return *m.IsbigWin
	}
	return false
}

func (m *EndLottery) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *EndLottery) GetWinScore() int64 {
	if m != nil && m.WinScore != nil {
		return *m.WinScore
	}
	return 0
}

func (m *EndLottery) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

// 牌局结束(广播)
type Pez_SendEndLottery struct {
	EndLotteryInfo   []*EndLottery `protobuf:"bytes,1,rep,name=endLotteryInfo" json:"endLotteryInfo,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Pez_SendEndLottery) Reset()                    { *m = Pez_SendEndLottery{} }
func (m *Pez_SendEndLottery) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendEndLottery) ProtoMessage()               {}
func (*Pez_SendEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{7} }

func (m *Pez_SendEndLottery) GetEndLotteryInfo() []*EndLottery {
	if m != nil {
		return m.EndLotteryInfo
	}
	return nil
}

type Pez_UserBean struct {
	UserId           *uint32   `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Score            *int64    `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	Coin             *int64    `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Tuozi            *PezTuozi `protobuf:"bytes,4,opt,name=tuozi" json:"tuozi,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Pez_UserBean) Reset()                    { *m = Pez_UserBean{} }
func (m *Pez_UserBean) String() string            { return proto.CompactTextString(m) }
func (*Pez_UserBean) ProtoMessage()               {}
func (*Pez_UserBean) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{8} }

func (m *Pez_UserBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_UserBean) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *Pez_UserBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *Pez_UserBean) GetTuozi() *PezTuozi {
	if m != nil {
		return m.Tuozi
	}
	return nil
}

type Pez_Bill struct {
	BeanList         []*Pez_UserBean `protobuf:"bytes,1,rep,name=BeanList" json:"BeanList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Pez_Bill) Reset()                    { *m = Pez_Bill{} }
func (m *Pez_Bill) String() string            { return proto.CompactTextString(m) }
func (*Pez_Bill) ProtoMessage()               {}
func (*Pez_Bill) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{9} }

func (m *Pez_Bill) GetBeanList() []*Pez_UserBean {
	if m != nil {
		return m.BeanList
	}
	return nil
}

func init() {
	proto.RegisterType((*Pez_DissolveDesk)(nil), "ddproto.pez_DissolveDesk")
	proto.RegisterType((*Pez_AckDissolveDesk)(nil), "ddproto.pez_AckDissolveDesk")
	proto.RegisterType((*Pez_Ready)(nil), "ddproto.pez_Ready")
	proto.RegisterType((*Pez_AckReady)(nil), "ddproto.pez_AckReady")
	proto.RegisterType((*Pez_EndLotteryInfo)(nil), "ddproto.pez_EndLotteryInfo")
	proto.RegisterType((*Pez_SendCurrentResult)(nil), "ddproto.pez_SendCurrentResult")
	proto.RegisterType((*EndLottery)(nil), "ddproto.EndLottery")
	proto.RegisterType((*Pez_SendEndLottery)(nil), "ddproto.pez_SendEndLottery")
	proto.RegisterType((*Pez_UserBean)(nil), "ddproto.pez_UserBean")
	proto.RegisterType((*Pez_Bill)(nil), "ddproto.pez_Bill")
}

func init() { proto.RegisterFile("pez_desk.proto", fileDescriptor45) }

var fileDescriptor45 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0xfd, 0x56, 0xb2, 0x1c, 0x79, 0x9c, 0x2f, 0xa4, 0x9b, 0xa4, 0x08, 0x53, 0x8a, 0xd0, 0x95,
	0xa0, 0x25, 0x10, 0xdf, 0x96, 0x5e, 0x34, 0x71, 0x21, 0x86, 0xd0, 0x3a, 0x1b, 0xda, 0xf4, 0x2e,
	0x28, 0xd2, 0x24, 0x5d, 0x2c, 0xef, 0x1a, 0xad, 0xd4, 0xfc, 0xbc, 0x44, 0x1f, 0xa2, 0x2f, 0xd1,
	0x47, 0xea, 0x63, 0x94, 0x5d, 0xad, 0x65, 0xa9, 0xe0, 0x5e, 0x94, 0xdc, 0x98, 0x3d, 0x67, 0x46,
	0x33, 0x73, 0xe6, 0x8c, 0x61, 0x67, 0x89, 0x8f, 0x57, 0x19, 0xaa, 0xf9, 0xe1, 0xb2, 0x90, 0xa5,
	0xa4, 0x5b, 0x59, 0x66, 0x1e, 0xa3, 0xbd, 0x54, 0x2e, 0x16, 0x52, 0x5c, 0xa5, 0x39, 0x47, 0x51,
	0xd6, 0xd1, 0xd1, 0x33, 0x4b, 0xa2, 0xa8, 0x16, 0x96, 0x32, 0x05, 0xae, 0x13, 0x85, 0x35, 0x8e,
	0xbe, 0xc0, 0xae, 0x66, 0x26, 0x5c, 0x29, 0x99, 0x7f, 0xc3, 0x09, 0xaa, 0x39, 0x7d, 0x0d, 0xfd,
	0xaf, 0x98, 0x64, 0x58, 0x04, 0x24, 0x24, 0xf1, 0x70, 0xbc, 0x7f, 0x68, 0xbb, 0x1c, 0xce, 0xf4,
	0xef, 0xa9, 0x89, 0x31, 0x9b, 0x43, 0x9f, 0x43, 0xbf, 0x52, 0x58, 0x4c, 0xb3, 0xc0, 0x09, 0x49,
	0xfc, 0x3f, 0xb3, 0x28, 0xfa, 0x4e, 0x60, 0x4f, 0x97, 0x7e, 0x97, 0xce, 0x9f, 0xbe, 0xba, 0xe6,
	0xf5, 0x1a, 0xa6, 0x59, 0xe0, 0x86, 0x24, 0xf6, 0x98, 0x45, 0x74, 0x04, 0xfe, 0x32, 0x51, 0xea,
	0x52, 0x16, 0x59, 0xd0, 0x0b, 0x49, 0x3c, 0x60, 0x0d, 0x8e, 0xce, 0x61, 0xa0, 0x07, 0x62, 0x98,
	0x64, 0x0f, 0x4f, 0x24, 0xf2, 0x06, 0xb6, 0xad, 0xc6, 0x7f, 0xa9, 0xba, 0x0b, 0xee, 0x42, 0xdd,
	0x9a, 0x92, 0x03, 0xa6, 0x9f, 0xad, 0x3e, 0x6e, 0xa7, 0xcf, 0x4f, 0x02, 0x54, 0x37, 0x7a, 0x2f,
	0xb2, 0x33, 0x59, 0x96, 0x58, 0x3c, 0x4c, 0xc5, 0x8d, 0x6c, 0xa5, 0x93, 0xce, 0x76, 0x46, 0xe0,
	0x0b, 0x9e, 0xce, 0x3f, 0x24, 0x0b, 0xb4, 0xd5, 0x1b, 0xac, 0x63, 0x5c, 0x5d, 0xf3, 0xdb, 0x4b,
	0x2e, 0x4c, 0x13, 0x9f, 0x35, 0x98, 0x06, 0xb0, 0xc5, 0xd5, 0xc7, 0x3b, 0x81, 0x85, 0x59, 0x9e,
	0xcf, 0x56, 0x90, 0xbe, 0x80, 0x81, 0xbe, 0x9a, 0xcf, 0x49, 0x5e, 0x61, 0xe0, 0x99, 0x95, 0xaf,
	0x09, 0x5d, 0xf3, 0x8e, 0x8b, 0x8b, 0x54, 0x16, 0x18, 0xf4, 0x43, 0x12, 0xbb, 0xac, 0xc1, 0xd1,
	0x2f, 0x02, 0x07, 0x7a, 0xf4, 0x0b, 0x14, 0xd9, 0x49, 0x55, 0x14, 0x28, 0x4a, 0x86, 0xaa, 0xca,
	0x4b, 0x1a, 0xc1, 0x76, 0x6a, 0x09, 0x59, 0x89, 0x5a, 0x83, 0xc7, 0x3a, 0x1c, 0x3d, 0x02, 0x5f,
	0x7f, 0x74, 0xcc, 0xf3, 0x3c, 0x70, 0x42, 0x37, 0x1e, 0x8e, 0x0f, 0x9a, 0x95, 0xea, 0xaa, 0x9f,
	0x14, 0x16, 0xc7, 0x98, 0x08, 0xd6, 0xa4, 0xd1, 0x97, 0x00, 0xd7, 0x32, 0x29, 0x32, 0x75, 0x22,
	0xab, 0xd2, 0x9e, 0x47, 0x8b, 0xd1, 0x22, 0x53, 0x0d, 0xa6, 0x93, 0xa0, 0x17, 0xba, 0xb1, 0xc7,
	0x56, 0x50, 0x7f, 0x99, 0x27, 0x62, 0xc2, 0x6b, 0x21, 0x9e, 0x11, 0xd2, 0x62, 0x74, 0x5c, 0x2e,
	0x51, 0xd4, 0xe3, 0x1b, 0xa1, 0x1e, 0x6b, 0x31, 0xd1, 0x0f, 0x02, 0xb0, 0x76, 0xe8, 0x6f, 0xee,
	0x34, 0x0e, 0x38, 0x9b, 0x1d, 0x70, 0xbb, 0x0e, 0xb4, 0x77, 0xdc, 0xeb, 0xee, 0x98, 0xbe, 0x82,
	0xfe, 0xdd, 0xbd, 0xbe, 0x08, 0x33, 0xf4, 0x70, 0xbc, 0xd7, 0xec, 0xe8, 0x12, 0xf9, 0x3d, 0x17,
	0x3a, 0xc4, 0x6c, 0x4a, 0x74, 0x5e, 0x9f, 0x92, 0xf6, 0xa3, 0x35, 0xec, 0x1b, 0xd8, 0xc1, 0xce,
	0x71, 0x05, 0xc4, 0xac, 0x7b, 0x5d, 0x6a, 0x9d, 0xcc, 0xfe, 0x48, 0x8d, 0x1e, 0xeb, 0xbf, 0xc1,
	0xca, 0x8c, 0x8d, 0xca, 0xf7, 0xc1, 0x53, 0x46, 0x80, 0x63, 0x04, 0xd4, 0x80, 0x52, 0xe8, 0xa5,
	0xd2, 0x5e, 0xa3, 0xcb, 0xcc, 0x9b, 0xc6, 0xe0, 0x95, 0x95, 0x7c, 0xe4, 0x46, 0xea, 0x70, 0x4c,
	0x3b, 0xa6, 0x9b, 0x08, 0xab, 0x13, 0xa2, 0xb7, 0xe0, 0x6b, 0xce, 0x58, 0x7f, 0x04, 0xbe, 0xee,
	0x7f, 0xc6, 0x55, 0x69, 0xc7, 0xdf, 0x74, 0x2d, 0xab, 0xb4, 0x63, 0xe7, 0xd4, 0x9d, 0xfd, 0x37,
	0x23, 0x33, 0xe7, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x76, 0x66, 0x40, 0x5c, 0x05, 0x00,
	0x00,
}
