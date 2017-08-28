// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pdk_desk.proto

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 房主解散房间(未开局)
type PdkReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsDaikai         *bool        `protobuf:"varint,3,opt,name=isDaikai" json:"isDaikai,omitempty"`
	DeskId           *int32       `protobuf:"varint,4,opt,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkReqDissolveDesk) Reset()                    { *m = PdkReqDissolveDesk{} }
func (m *PdkReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkReqDissolveDesk) ProtoMessage()               {}
func (*PdkReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{0} }

func (m *PdkReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkReqDissolveDesk) GetIsDaikai() bool {
	if m != nil && m.IsDaikai != nil {
		return *m.IsDaikai
	}
	return false
}

func (m *PdkReqDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

// 解散房间回复
type PdkAckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=passWord" json:"passWord,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkAckDissolveDesk) Reset()                    { *m = PdkAckDissolveDesk{} }
func (m *PdkAckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkAckDissolveDesk) ProtoMessage()               {}
func (*PdkAckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{1} }

func (m *PdkAckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkAckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkAckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PdkAckDissolveDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

// 准备游戏
type PdkReqReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsShowHandPokers *bool        `protobuf:"varint,3,opt,name=isShowHandPokers" json:"isShowHandPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkReqReady) Reset()                    { *m = PdkReqReady{} }
func (m *PdkReqReady) String() string            { return proto.CompactTextString(m) }
func (*PdkReqReady) ProtoMessage()               {}
func (*PdkReqReady) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{2} }

func (m *PdkReqReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkReqReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkReqReady) GetIsShowHandPokers() bool {
	if m != nil && m.IsShowHandPokers != nil {
		return *m.IsShowHandPokers
	}
	return false
}

// 准备游戏的结果
type PdkAckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkAckReady) Reset()                    { *m = PdkAckReady{} }
func (m *PdkAckReady) String() string            { return proto.CompactTextString(m) }
func (*PdkAckReady) ProtoMessage()               {}
func (*PdkAckReady) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{3} }

func (m *PdkAckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkAckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *PdkAckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 赢牌信息：谁赢了多少
type PdkBaseWinCoinInfo struct {
	NickName         *string            `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	BaseValue        *int32             `protobuf:"varint,3,opt,name=baseValue" json:"baseValue,omitempty"`
	WinCoin          *int64             `protobuf:"varint,4,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64             `protobuf:"varint,5,opt,name=coin" json:"coin,omitempty"`
	IsDiZhu          *bool              `protobuf:"varint,6,opt,name=isDiZhu" json:"isDiZhu,omitempty"`
	Rate             *int32             `protobuf:"varint,7,opt,name=rate" json:"rate,omitempty"`
	Description      *string            `protobuf:"bytes,8,opt,name=description" json:"description,omitempty"`
	Bomb             *int32             `protobuf:"varint,9,opt,name=bomb" json:"bomb,omitempty"`
	IsSpring         *bool              `protobuf:"varint,10,opt,name=isSpring" json:"isSpring,omitempty"`
	HandPokers       []*ClientBasePoker `protobuf:"bytes,11,rep,name=handPokers" json:"handPokers,omitempty"`
	HaveNiaoPai      *bool              `protobuf:"varint,12,opt,name=haveNiaoPai" json:"haveNiaoPai,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *PdkBaseWinCoinInfo) Reset()                    { *m = PdkBaseWinCoinInfo{} }
func (m *PdkBaseWinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkBaseWinCoinInfo) ProtoMessage()               {}
func (*PdkBaseWinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{4} }

func (m *PdkBaseWinCoinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PdkBaseWinCoinInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetBaseValue() int32 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetIsDiZhu() bool {
	if m != nil && m.IsDiZhu != nil {
		return *m.IsDiZhu
	}
	return false
}

func (m *PdkBaseWinCoinInfo) GetRate() int32 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *PdkBaseWinCoinInfo) GetBomb() int32 {
	if m != nil && m.Bomb != nil {
		return *m.Bomb
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

func (m *PdkBaseWinCoinInfo) GetHandPokers() []*ClientBasePoker {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *PdkBaseWinCoinInfo) GetHaveNiaoPai() bool {
	if m != nil && m.HaveNiaoPai != nil {
		return *m.HaveNiaoPai
	}
	return false
}

// 本局结果(广播)
type PdkBaCurrentResult struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WinCoinInfo      []*PdkBaseWinCoinInfo `protobuf:"bytes,2,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	RemainPais       []int32               `protobuf:"varint,3,rep,name=RemainPais" json:"RemainPais,omitempty"`
	CurrPlayCount    *int32                `protobuf:"varint,4,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PdkBaCurrentResult) Reset()                    { *m = PdkBaCurrentResult{} }
func (m *PdkBaCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*PdkBaCurrentResult) ProtoMessage()               {}
func (*PdkBaCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{5} }

func (m *PdkBaCurrentResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkBaCurrentResult) GetWinCoinInfo() []*PdkBaseWinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

func (m *PdkBaCurrentResult) GetRemainPais() []int32 {
	if m != nil {
		return m.RemainPais
	}
	return nil
}

func (m *PdkBaCurrentResult) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

type PdkBaseEndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	BigWin           *bool   `protobuf:"varint,3,opt,name=bigWin" json:"bigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	MaxWinCoin       *int32  `protobuf:"varint,6,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	CountBomb        *int32  `protobuf:"varint,7,opt,name=countBomb" json:"countBomb,omitempty"`
	CountWin         *int32  `protobuf:"varint,8,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32  `protobuf:"varint,9,opt,name=countLose" json:"countLose,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PdkBaseEndLotteryInfo) Reset()                    { *m = PdkBaseEndLotteryInfo{} }
func (m *PdkBaseEndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkBaseEndLotteryInfo) ProtoMessage()               {}
func (*PdkBaseEndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{6} }

func (m *PdkBaseEndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PdkBaseEndLotteryInfo) GetBigWin() bool {
	if m != nil && m.BigWin != nil {
		return *m.BigWin
	}
	return false
}

func (m *PdkBaseEndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PdkBaseEndLotteryInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetMaxWinCoin() int32 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

// 牌局结束(广播)
type PdkBcEndLottery struct {
	Header           *ProtoHeader             `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CoinInfo         []*PdkBaseEndLotteryInfo `protobuf:"bytes,2,rep,name=coinInfo" json:"coinInfo,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *PdkBcEndLottery) Reset()                    { *m = PdkBcEndLottery{} }
func (m *PdkBcEndLottery) String() string            { return proto.CompactTextString(m) }
func (*PdkBcEndLottery) ProtoMessage()               {}
func (*PdkBcEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{7} }

func (m *PdkBcEndLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkBcEndLottery) GetCoinInfo() []*PdkBaseEndLotteryInfo {
	if m != nil {
		return m.CoinInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PdkReqDissolveDesk)(nil), "ddproto.pdk_req_dissolveDesk")
	proto.RegisterType((*PdkAckDissolveDesk)(nil), "ddproto.pdk_ack_dissolveDesk")
	proto.RegisterType((*PdkReqReady)(nil), "ddproto.pdk_req_ready")
	proto.RegisterType((*PdkAckReady)(nil), "ddproto.pdk_ack_ready")
	proto.RegisterType((*PdkBaseWinCoinInfo)(nil), "ddproto.pdk_base_winCoinInfo")
	proto.RegisterType((*PdkBaCurrentResult)(nil), "ddproto.pdk_ba_currentResult")
	proto.RegisterType((*PdkBaseEndLotteryInfo)(nil), "ddproto.pdk_base_endLotteryInfo")
	proto.RegisterType((*PdkBcEndLottery)(nil), "ddproto.pdk_bc_endLottery")
}

func init() { proto.RegisterFile("pdk_desk.proto", fileDescriptor39) }

var fileDescriptor39 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x25, 0xcd, 0xda, 0xb5, 0x2e, 0x43, 0x23, 0x4c, 0x60, 0x4d, 0x30, 0x45, 0x11, 0x87, 0x0a,
	0xa1, 0x1d, 0x76, 0x44, 0x48, 0x48, 0xdb, 0x0e, 0x9b, 0x34, 0x8d, 0xc8, 0x93, 0x98, 0xc4, 0xa5,
	0x72, 0x63, 0xd3, 0x5a, 0x69, 0xec, 0x60, 0x27, 0x1b, 0x3b, 0x21, 0xfe, 0x01, 0xfc, 0x2e, 0xfe,
	0x0e, 0x3f, 0x00, 0x7d, 0x8e, 0x9b, 0xb8, 0x0c, 0x0e, 0x43, 0x70, 0xa9, 0xfc, 0x9e, 0xfd, 0xf9,
	0x7b, 0x79, 0xdf, 0x73, 0xd1, 0x83, 0x92, 0xe5, 0x53, 0xc6, 0x4d, 0xbe, 0x5f, 0x6a, 0x55, 0xa9,
	0x68, 0x93, 0x31, 0xbb, 0xd8, 0x7d, 0x94, 0xa9, 0xa2, 0x50, 0x72, 0x9a, 0x2d, 0x05, 0x97, 0x55,
	0xb3, 0x9b, 0x7c, 0x0d, 0xd0, 0x0e, 0x14, 0x68, 0xfe, 0x71, 0xca, 0x84, 0x31, 0x6a, 0x79, 0xc5,
	0x8f, 0xb9, 0xc9, 0xa3, 0x97, 0x68, 0xb0, 0xe0, 0x94, 0x71, 0x8d, 0x83, 0x38, 0x98, 0x8c, 0x0f,
	0x76, 0xf6, 0xdd, 0x3d, 0xfb, 0x29, 0xfc, 0x9e, 0xd8, 0x3d, 0xe2, 0xce, 0x44, 0x8f, 0xd1, 0xa0,
	0x36, 0x5c, 0x9f, 0x32, 0xdc, 0x8b, 0x83, 0xc9, 0x16, 0x71, 0x28, 0xda, 0x45, 0x43, 0x61, 0x8e,
	0xa9, 0xc8, 0xa9, 0xc0, 0x61, 0x1c, 0x4c, 0x86, 0xa4, 0xc5, 0x50, 0x03, 0x32, 0x4f, 0x19, 0xde,
	0x88, 0x83, 0x49, 0x9f, 0x38, 0xd4, 0x4a, 0xa2, 0x59, 0xfe, 0x3f, 0x24, 0x75, 0x6d, 0x43, 0xbf,
	0x2d, 0x48, 0x2d, 0xa9, 0x31, 0x97, 0x4a, 0x37, 0x82, 0x46, 0xa4, 0xc5, 0xc9, 0x97, 0x00, 0x6d,
	0xad, 0x5c, 0xd2, 0x9c, 0xb2, 0x9b, 0x7f, 0xa4, 0xe5, 0x05, 0xda, 0x16, 0xe6, 0x62, 0xa1, 0xae,
	0x4f, 0xa8, 0x64, 0xa9, 0xca, 0xb9, 0x36, 0xce, 0xa6, 0x5b, 0x7c, 0x32, 0x6f, 0x24, 0x80, 0x2b,
	0x7f, 0x23, 0x61, 0x1b, 0x85, 0x85, 0x99, 0xdb, 0xfe, 0x23, 0x02, 0x4b, 0x4f, 0x54, 0xe8, 0x8b,
	0x4a, 0x7e, 0xf4, 0x1a, 0xff, 0x67, 0xd4, 0xf0, 0xe9, 0xb5, 0x90, 0x47, 0x4a, 0xc8, 0x53, 0xf9,
	0x41, 0x81, 0x43, 0x52, 0x64, 0xf9, 0x39, 0x2d, 0xb8, 0x6d, 0x39, 0x22, 0x2d, 0xfe, 0xe3, 0x17,
	0x3e, 0x45, 0x23, 0xb8, 0xe7, 0x1d, 0x5d, 0xd6, 0xdc, 0x19, 0xde, 0x11, 0x11, 0x46, 0x9b, 0xae,
	0x81, 0xb5, 0x3c, 0x24, 0x2b, 0x18, 0x45, 0x68, 0x23, 0x03, 0xba, 0x6f, 0x69, 0xbb, 0x86, 0xd3,
	0xc2, 0x1c, 0x8b, 0xf7, 0x8b, 0x1a, 0x0f, 0xac, 0x49, 0x2b, 0x08, 0xa7, 0x35, 0xad, 0x38, 0xde,
	0xb4, 0x0d, 0xec, 0x3a, 0x8a, 0xd1, 0x98, 0x71, 0x93, 0x69, 0x51, 0x56, 0x42, 0x49, 0x3c, 0xb4,
	0x82, 0x7d, 0x0a, 0xaa, 0x66, 0xaa, 0x98, 0xe1, 0x51, 0x53, 0x05, 0xeb, 0x26, 0xb0, 0x17, 0xa5,
	0x16, 0x72, 0x8e, 0xd1, 0x2a, 0xb0, 0x0d, 0x8e, 0x5e, 0x21, 0xb4, 0xe8, 0xe6, 0x34, 0x8e, 0xc3,
	0xc9, 0xf8, 0x60, 0xb7, 0x35, 0xbd, 0x79, 0x56, 0x8d, 0x6b, 0x25, 0x1c, 0x21, 0xde, 0x69, 0x50,
	0xb3, 0xa0, 0x57, 0xfc, 0x5c, 0x50, 0x95, 0x52, 0x81, 0xef, 0xdb, 0xab, 0x7d, 0x2a, 0xf9, 0x1e,
	0xac, 0x6c, 0x9f, 0x66, 0xb5, 0xd6, 0x5c, 0x56, 0x84, 0x9b, 0x7a, 0x59, 0xdd, 0x71, 0xce, 0x6f,
	0xd0, 0xd8, 0x9b, 0x19, 0xee, 0x59, 0x95, 0xcf, 0xda, 0x92, 0xdf, 0x0d, 0x96, 0xf8, 0x15, 0xd1,
	0x1e, 0x42, 0x84, 0x17, 0x54, 0xc8, 0x94, 0x0a, 0x48, 0x63, 0x38, 0xe9, 0x13, 0x8f, 0x89, 0x9e,
	0xa3, 0x2d, 0xd0, 0x97, 0x2e, 0xe9, 0xcd, 0x91, 0xaa, 0x65, 0xe5, 0x5e, 0xef, 0x3a, 0x99, 0x7c,
	0xeb, 0xa1, 0x27, 0x6d, 0x2f, 0x2e, 0xd9, 0x99, 0xaa, 0x2a, 0xae, 0x6f, 0x6c, 0x87, 0x2e, 0x2b,
	0xc1, 0xaf, 0x7f, 0x16, 0x6d, 0xbe, 0x7a, 0xb7, 0xf3, 0x35, 0x13, 0xf3, 0x4b, 0x21, 0xdd, 0xfb,
	0x70, 0xa8, 0xc9, 0xc4, 0xdb, 0x6b, 0xc9, 0xb5, 0xd5, 0x61, 0x33, 0x61, 0xa1, 0x9f, 0xad, 0xfe,
	0x7a, 0xb6, 0xf6, 0x10, 0x2a, 0xe8, 0xa7, 0x4b, 0xb7, 0x39, 0xb0, 0xf2, 0x3d, 0x06, 0x32, 0x9b,
	0xc1, 0x47, 0x1c, 0x42, 0x38, 0x9a, 0x48, 0x75, 0x04, 0xa8, 0xb4, 0x00, 0xb4, 0x0c, 0xed, 0x66,
	0x8b, 0xdb, 0xca, 0x33, 0x65, 0xb8, 0x8b, 0x55, 0x47, 0x24, 0x9f, 0xd1, 0x43, 0x6b, 0x49, 0xe6,
	0x19, 0x72, 0xc7, 0xe9, 0xbe, 0x86, 0xe6, 0x6b, 0xa3, 0x8d, 0x6f, 0x8f, 0x76, 0xdd, 0x6e, 0xd2,
	0x56, 0x1c, 0xf6, 0x4e, 0xc2, 0xf4, 0xde, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xb2, 0xd8,
	0xe7, 0x21, 0x06, 0x00, 0x00,
}
