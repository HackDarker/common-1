// Code generated by protoc-gen-go. DO NOT EDIT.
// source: majiang_desk.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of CardInfo from common_mj.proto

// Ignoring public import of RoomTypeInfo from common_mj.proto

// Ignoring public import of PlayOptions from common_mj.proto

// Ignoring public import of ChangShaPlayOptions from common_mj.proto

// Ignoring public import of BaiShanPlayOptions from common_mj.proto

// Ignoring public import of ZhuanZhuanPlayOptions from common_mj.proto

// Ignoring public import of HaiNanPlayOptions from common_mj.proto

// Ignoring public import of ComposeCard from common_mj.proto

// Ignoring public import of PlayerCard from common_mj.proto

// Ignoring public import of GangBean from common_mj.proto

// Ignoring public import of ErrorCode from common_mj.proto

// Ignoring public import of MahjongColor from common_mj.proto

// Ignoring public import of GangType from common_mj.proto

// Ignoring public import of HuType from common_mj.proto

// Ignoring public import of ComposeCardType from common_mj.proto

// Ignoring public import of PaiType from common_mj.proto

// Ignoring public import of MJUserGameStatus from common_mj.proto

// Ignoring public import of DeskGameStatus from common_mj.proto

// Ignoring public import of MJRoomType from common_mj.proto

// 房主解散房间(未开局)
type Game_DissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_DissolveDesk) Reset()                    { *m = Game_DissolveDesk{} }
func (m *Game_DissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Game_DissolveDesk) ProtoMessage()               {}
func (*Game_DissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{0} }

func (m *Game_DissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_DissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_DissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

// 解散房间回复
type Game_AckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=passWord" json:"passWord,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckDissolveDesk) Reset()                    { *m = Game_AckDissolveDesk{} }
func (m *Game_AckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Game_AckDissolveDesk) ProtoMessage()               {}
func (*Game_AckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{1} }

func (m *Game_AckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_AckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *Game_AckDissolveDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

// 申请解散房间(游戏中)
type Game_ReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_ReqDissolveDesk) Reset()                    { *m = Game_ReqDissolveDesk{} }
func (m *Game_ReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Game_ReqDissolveDesk) ProtoMessage()               {}
func (*Game_ReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{2} }

func (m *Game_ReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_ReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type Game_AckReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserIdAgree      []uint32     `protobuf:"varint,2,rep,name=userIdAgree" json:"userIdAgree,omitempty"`
	UserIdWait       []uint32     `protobuf:"varint,3,rep,name=userIdWait" json:"userIdWait,omitempty"`
	UserIdDisagree   []uint32     `protobuf:"varint,4,rep,name=userIdDisagree" json:"userIdDisagree,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckReqDissolveDesk) Reset()                    { *m = Game_AckReqDissolveDesk{} }
func (m *Game_AckReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Game_AckReqDissolveDesk) ProtoMessage()               {}
func (*Game_AckReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{3} }

func (m *Game_AckReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckReqDissolveDesk) GetUserIdAgree() []uint32 {
	if m != nil {
		return m.UserIdAgree
	}
	return nil
}

func (m *Game_AckReqDissolveDesk) GetUserIdWait() []uint32 {
	if m != nil {
		return m.UserIdWait
	}
	return nil
}

func (m *Game_AckReqDissolveDesk) GetUserIdDisagree() []uint32 {
	if m != nil {
		return m.UserIdDisagree
	}
	return nil
}

// 准备游戏
type Game_Ready struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_Ready) Reset()                    { *m = Game_Ready{} }
func (m *Game_Ready) String() string            { return proto.CompactTextString(m) }
func (*Game_Ready) ProtoMessage()               {}
func (*Game_Ready) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{4} }

func (m *Game_Ready) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_Ready) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 准备游戏的结果
type Game_AckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckReady) Reset()                    { *m = Game_AckReady{} }
func (m *Game_AckReady) String() string            { return proto.CompactTextString(m) }
func (*Game_AckReady) ProtoMessage()               {}
func (*Game_AckReady) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{5} }

func (m *Game_AckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *Game_AckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type BirdInfo struct {
	BirdPai          *CardInfo `protobuf:"bytes,1,opt,name=birdPai" json:"birdPai,omitempty"`
	ZhuaUser         *uint32   `protobuf:"varint,2,opt,name=zhuaUser" json:"zhuaUser,omitempty"`
	BirdUser         *uint32   `protobuf:"varint,3,opt,name=birdUser" json:"birdUser,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BirdInfo) Reset()                    { *m = BirdInfo{} }
func (m *BirdInfo) String() string            { return proto.CompactTextString(m) }
func (*BirdInfo) ProtoMessage()               {}
func (*BirdInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{6} }

func (m *BirdInfo) GetBirdPai() *CardInfo {
	if m != nil {
		return m.BirdPai
	}
	return nil
}

func (m *BirdInfo) GetZhuaUser() uint32 {
	if m != nil && m.ZhuaUser != nil {
		return *m.ZhuaUser
	}
	return 0
}

func (m *BirdInfo) GetBirdUser() uint32 {
	if m != nil && m.BirdUser != nil {
		return *m.BirdUser
	}
	return 0
}

// 赢牌信息：谁赢了多少
type WinCoinInfo struct {
	NickName         *string     `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32     `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	WinCoin          *int64      `protobuf:"varint,3,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64      `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	CardTitle        *string     `protobuf:"bytes,5,opt,name=cardTitle" json:"cardTitle,omitempty"`
	Cards            *PlayerCard `protobuf:"bytes,6,opt,name=cards" json:"cards,omitempty"`
	IsDealer         *bool       `protobuf:"varint,7,opt,name=isDealer" json:"isDealer,omitempty"`
	HuCount          *int32      `protobuf:"varint,8,opt,name=huCount" json:"huCount,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *WinCoinInfo) Reset()                    { *m = WinCoinInfo{} }
func (m *WinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*WinCoinInfo) ProtoMessage()               {}
func (*WinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{7} }

func (m *WinCoinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *WinCoinInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *WinCoinInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *WinCoinInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *WinCoinInfo) GetCardTitle() string {
	if m != nil && m.CardTitle != nil {
		return *m.CardTitle
	}
	return ""
}

func (m *WinCoinInfo) GetCards() *PlayerCard {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *WinCoinInfo) GetIsDealer() bool {
	if m != nil && m.IsDealer != nil {
		return *m.IsDealer
	}
	return false
}

func (m *WinCoinInfo) GetHuCount() int32 {
	if m != nil && m.HuCount != nil {
		return *m.HuCount
	}
	return 0
}

type EndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	BigWin           *bool   `protobuf:"varint,3,opt,name=bigWin" json:"bigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	CountHu          *int32  `protobuf:"varint,6,opt,name=countHu" json:"countHu,omitempty"`
	CountZiMo        *int32  `protobuf:"varint,7,opt,name=countZiMo" json:"countZiMo,omitempty"`
	CountDianPao     *int32  `protobuf:"varint,8,opt,name=countDianPao" json:"countDianPao,omitempty"`
	CountAnGang      *int32  `protobuf:"varint,9,opt,name=countAnGang" json:"countAnGang,omitempty"`
	CountMingGang    *int32  `protobuf:"varint,10,opt,name=countMingGang" json:"countMingGang,omitempty"`
	CountDianGang    *int32  `protobuf:"varint,11,opt,name=countDianGang" json:"countDianGang,omitempty"`
	CountChaJiao     *int32  `protobuf:"varint,12,opt,name=countChaJiao" json:"countChaJiao,omitempty"`
	BestGunner       *bool   `protobuf:"varint,13,opt,name=bestGunner" json:"bestGunner,omitempty"`
	CountLianZhuang  *int32  `protobuf:"varint,14,opt,name=countLianZhuang" json:"countLianZhuang,omitempty"`
	CountZhaMa       *int32  `protobuf:"varint,15,opt,name=countZhaMa" json:"countZhaMa,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EndLotteryInfo) Reset()                    { *m = EndLotteryInfo{} }
func (m *EndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*EndLotteryInfo) ProtoMessage()               {}
func (*EndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{8} }

func (m *EndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *EndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *EndLotteryInfo) GetBigWin() bool {
	if m != nil && m.BigWin != nil {
		return *m.BigWin
	}
	return false
}

func (m *EndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *EndLotteryInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *EndLotteryInfo) GetCountHu() int32 {
	if m != nil && m.CountHu != nil {
		return *m.CountHu
	}
	return 0
}

func (m *EndLotteryInfo) GetCountZiMo() int32 {
	if m != nil && m.CountZiMo != nil {
		return *m.CountZiMo
	}
	return 0
}

func (m *EndLotteryInfo) GetCountDianPao() int32 {
	if m != nil && m.CountDianPao != nil {
		return *m.CountDianPao
	}
	return 0
}

func (m *EndLotteryInfo) GetCountAnGang() int32 {
	if m != nil && m.CountAnGang != nil {
		return *m.CountAnGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountMingGang() int32 {
	if m != nil && m.CountMingGang != nil {
		return *m.CountMingGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountDianGang() int32 {
	if m != nil && m.CountDianGang != nil {
		return *m.CountDianGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountChaJiao() int32 {
	if m != nil && m.CountChaJiao != nil {
		return *m.CountChaJiao
	}
	return 0
}

func (m *EndLotteryInfo) GetBestGunner() bool {
	if m != nil && m.BestGunner != nil {
		return *m.BestGunner
	}
	return false
}

func (m *EndLotteryInfo) GetCountLianZhuang() int32 {
	if m != nil && m.CountLianZhuang != nil {
		return *m.CountLianZhuang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountZhaMa() int32 {
	if m != nil && m.CountZhaMa != nil {
		return *m.CountZhaMa
	}
	return 0
}

// 本局结果(广播)
type Game_SendCurrentResult struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WinCoinInfo      []*WinCoinInfo `protobuf:"bytes,2,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	BridInfo         []*BirdInfo    `protobuf:"bytes,3,rep,name=bridInfo" json:"bridInfo,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Game_SendCurrentResult) Reset()                    { *m = Game_SendCurrentResult{} }
func (m *Game_SendCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*Game_SendCurrentResult) ProtoMessage()               {}
func (*Game_SendCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{9} }

func (m *Game_SendCurrentResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_SendCurrentResult) GetWinCoinInfo() []*WinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

func (m *Game_SendCurrentResult) GetBridInfo() []*BirdInfo {
	if m != nil {
		return m.BridInfo
	}
	return nil
}

// 牌局结束(广播)
type Game_SendEndLottery struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CoinInfo         []*EndLotteryInfo `protobuf:"bytes,2,rep,name=coinInfo" json:"coinInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Game_SendEndLottery) Reset()                    { *m = Game_SendEndLottery{} }
func (m *Game_SendEndLottery) String() string            { return proto.CompactTextString(m) }
func (*Game_SendEndLottery) ProtoMessage()               {}
func (*Game_SendEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{10} }

func (m *Game_SendEndLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_SendEndLottery) GetCoinInfo() []*EndLotteryInfo {
	if m != nil {
		return m.CoinInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Game_DissolveDesk)(nil), "ddproto.game_DissolveDesk")
	proto.RegisterType((*Game_AckDissolveDesk)(nil), "ddproto.game_AckDissolveDesk")
	proto.RegisterType((*Game_ReqDissolveDesk)(nil), "ddproto.game_ReqDissolveDesk")
	proto.RegisterType((*Game_AckReqDissolveDesk)(nil), "ddproto.game_AckReqDissolveDesk")
	proto.RegisterType((*Game_Ready)(nil), "ddproto.game_Ready")
	proto.RegisterType((*Game_AckReady)(nil), "ddproto.game_AckReady")
	proto.RegisterType((*BirdInfo)(nil), "ddproto.BirdInfo")
	proto.RegisterType((*WinCoinInfo)(nil), "ddproto.WinCoinInfo")
	proto.RegisterType((*EndLotteryInfo)(nil), "ddproto.EndLotteryInfo")
	proto.RegisterType((*Game_SendCurrentResult)(nil), "ddproto.game_SendCurrentResult")
	proto.RegisterType((*Game_SendEndLottery)(nil), "ddproto.game_SendEndLottery")
}

func init() { proto.RegisterFile("majiang_desk.proto", fileDescriptor31) }

var fileDescriptor31 = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x7d, 0x8e, 0x93, 0xc6, 0xb9, 0x69, 0xda, 0xd7, 0x69, 0xd5, 0x8e, 0xaa, 0xa7, 0xa7, 0xc8,
	0x7a, 0x7a, 0x0a, 0x02, 0xba, 0x00, 0x89, 0x7d, 0x9b, 0xa0, 0xb6, 0xa8, 0x85, 0x68, 0x00, 0x45,
	0xaa, 0x90, 0xaa, 0x49, 0x3c, 0x38, 0xd3, 0x24, 0xe3, 0x76, 0xc6, 0xa6, 0x94, 0x5f, 0xc1, 0x1f,
	0x61, 0xc1, 0xbf, 0x62, 0xc5, 0x6f, 0x40, 0x73, 0xfd, 0x11, 0xbb, 0x88, 0x45, 0x45, 0xc5, 0x26,
	0x9a, 0x73, 0xe6, 0xe4, 0x9e, 0x7b, 0x8f, 0xe6, 0x1a, 0xc8, 0x82, 0x5f, 0x48, 0xae, 0xc2, 0xf3,
	0x40, 0x98, 0xd9, 0xde, 0xa5, 0x8e, 0xe2, 0x88, 0x34, 0x83, 0x00, 0x0f, 0xbb, 0xeb, 0x93, 0x68,
	0xb1, 0x88, 0xd4, 0xf9, 0xe2, 0x22, 0xbd, 0xf1, 0xaf, 0x60, 0x23, 0xe4, 0x0b, 0x71, 0x3e, 0x90,
	0xc6, 0x44, 0xf3, 0x0f, 0x62, 0x20, 0xcc, 0x8c, 0x3c, 0x82, 0x95, 0xa9, 0xe0, 0x81, 0xd0, 0xd4,
	0xe9, 0x3a, 0xbd, 0xf6, 0x93, 0xad, 0xbd, 0xec, 0xff, 0x7b, 0x43, 0xfb, 0x7b, 0x84, 0x77, 0x2c,
	0xd3, 0x90, 0x6d, 0x58, 0x49, 0x8c, 0xd0, 0xc7, 0x01, 0xad, 0x75, 0x9d, 0x5e, 0x87, 0x65, 0xc8,
	0xf2, 0xb6, 0x85, 0xe3, 0x80, 0xba, 0x5d, 0xa7, 0xd7, 0x60, 0x19, 0xf2, 0x3f, 0x3b, 0xb0, 0x85,
	0x9e, 0xfb, 0x93, 0xd9, 0x9f, 0xb3, 0x25, 0xbb, 0xe0, 0x5d, 0x72, 0x63, 0x46, 0x91, 0x0e, 0x68,
	0xbd, 0xeb, 0xf4, 0x5a, 0xac, 0xc0, 0xfe, 0xbb, 0xac, 0x23, 0x26, 0xae, 0xee, 0xbf, 0x23, 0xff,
	0xab, 0x03, 0x3b, 0xf9, 0xc0, 0xbf, 0xe7, 0xd0, 0x85, 0x76, 0x5a, 0x73, 0x3f, 0xd4, 0x42, 0xd0,
	0x5a, 0xd7, 0xed, 0x75, 0x58, 0x99, 0x22, 0xff, 0x02, 0xa4, 0x70, 0xc4, 0x65, 0x4c, 0x5d, 0x14,
	0x94, 0x18, 0xf2, 0x3f, 0xac, 0xa5, 0x68, 0x20, 0x0d, 0xc7, 0x22, 0x75, 0xd4, 0xdc, 0x62, 0x7d,
	0x06, 0x90, 0x25, 0xc2, 0x83, 0x9b, 0x7b, 0xca, 0x21, 0x84, 0xce, 0x32, 0x86, 0xbb, 0x97, 0xfd,
	0x1b, 0xdc, 0x85, 0x09, 0xb1, 0x66, 0x8b, 0xd9, 0x63, 0xc9, 0xc8, 0xad, 0x18, 0x45, 0xe0, 0x1d,
	0x48, 0x1d, 0x1c, 0xab, 0xf7, 0x11, 0x79, 0x08, 0xcd, 0xb1, 0xd4, 0xc1, 0x90, 0xcb, 0xcc, 0x64,
	0xa3, 0x30, 0xe9, 0xf3, 0x54, 0xc3, 0x72, 0x85, 0x7d, 0x23, 0x9f, 0xa6, 0x09, 0x7f, 0x6b, 0x84,
	0xce, 0x7a, 0x2f, 0xb0, 0xbd, 0xb3, 0x32, 0xbc, 0x4b, 0xed, 0x0a, 0xec, 0x7f, 0x77, 0xa0, 0x3d,
	0x92, 0xaa, 0x1f, 0x49, 0x85, 0xa6, 0xbb, 0xe0, 0x29, 0x39, 0x99, 0xbd, 0xe4, 0x0b, 0x81, 0xae,
	0x2d, 0x56, 0xe0, 0x5f, 0xbe, 0x5b, 0x0a, 0xcd, 0xeb, 0xb4, 0x04, 0x96, 0x77, 0x59, 0x0e, 0x09,
	0x81, 0xfa, 0xc4, 0xd2, 0x75, 0xa4, 0xf1, 0x4c, 0xfe, 0x81, 0xd6, 0x84, 0xeb, 0xe0, 0x8d, 0x8c,
	0xe7, 0x82, 0x36, 0xd0, 0x62, 0x49, 0x90, 0x07, 0xd0, 0xb0, 0xc0, 0xd0, 0x15, 0x1c, 0x79, 0x73,
	0x99, 0xeb, 0x9c, 0xdf, 0x08, 0x6d, 0x07, 0x67, 0xa9, 0xc2, 0xb6, 0x2a, 0xcd, 0x40, 0xf0, 0xb9,
	0xd0, 0xb4, 0xd9, 0x75, 0x7a, 0x1e, 0x2b, 0xb0, 0x6d, 0x69, 0x9a, 0xf4, 0xa3, 0x44, 0xc5, 0xd4,
	0xc3, 0x5d, 0xca, 0xa1, 0xff, 0xcd, 0x85, 0xb5, 0xe7, 0x2a, 0x38, 0x89, 0xe2, 0x58, 0xe8, 0x1b,
	0x9c, 0x79, 0x39, 0x97, 0x53, 0x99, 0xab, 0x9c, 0x45, 0xed, 0xe7, 0x2c, 0xc6, 0x32, 0x1c, 0x65,
	0x23, 0x7b, 0x2c, 0x43, 0xd6, 0x58, 0x9a, 0x57, 0xd7, 0x4a, 0x68, 0x1c, 0xda, 0x63, 0x39, 0x2c,
	0xa7, 0xd4, 0xa8, 0xa6, 0x44, 0xa1, 0x39, 0xb1, 0xbd, 0x1d, 0x25, 0x38, 0x75, 0x83, 0xe5, 0x10,
	0xb3, 0xb2, 0xc7, 0x33, 0x79, 0x1a, 0xe1, 0x8c, 0x0d, 0xb6, 0x24, 0x88, 0x0f, 0xab, 0x08, 0x06,
	0x92, 0xab, 0x21, 0x8f, 0xb2, 0x49, 0x2b, 0x9c, 0xdd, 0x3b, 0xc4, 0xfb, 0xea, 0x90, 0xab, 0x90,
	0xb6, 0x50, 0x52, 0xa6, 0xc8, 0x7f, 0xd0, 0x41, 0x78, 0x2a, 0x55, 0x88, 0x1a, 0x40, 0x4d, 0x95,
	0x2c, 0x54, 0xb6, 0x2e, 0xaa, 0xda, 0x25, 0x55, 0x4e, 0x16, 0x1d, 0xf5, 0xa7, 0xfc, 0x85, 0xe4,
	0x11, 0x5d, 0x2d, 0x75, 0x94, 0x71, 0x76, 0xcf, 0xc7, 0xc2, 0xc4, 0x87, 0x89, 0xb2, 0x21, 0x75,
	0x30, 0xa4, 0x12, 0x43, 0x7a, 0xb0, 0x8e, 0xfa, 0x13, 0xc9, 0xd5, 0xd9, 0x34, 0xb1, 0x5e, 0x6b,
	0x58, 0xe6, 0x36, 0x6d, 0x2b, 0xa5, 0x61, 0x4c, 0xf9, 0x29, 0xa7, 0xeb, 0x28, 0x2a, 0x31, 0xfe,
	0x17, 0x07, 0xb6, 0x71, 0x6d, 0x5f, 0x0b, 0x15, 0xf4, 0x13, 0xad, 0x85, 0x8a, 0x99, 0x30, 0xc9,
	0x3c, 0xbe, 0xe3, 0xfe, 0x3e, 0x83, 0xf6, 0xf5, 0x72, 0x47, 0xf0, 0xe3, 0x55, 0xfe, 0x4b, 0x69,
	0x7f, 0x58, 0x59, 0x48, 0x1e, 0x83, 0x37, 0xd6, 0x12, 0x37, 0x15, 0x3f, 0x68, 0xe5, 0x15, 0xce,
	0xd7, 0x9c, 0x15, 0x12, 0xff, 0x23, 0x6c, 0x16, 0xed, 0x2e, 0x9f, 0xe8, 0x1d, 0x7b, 0x7d, 0x0a,
	0xde, 0xa4, 0xda, 0xe8, 0x4e, 0xa1, 0xaf, 0xbe, 0x7b, 0x56, 0x08, 0x0f, 0x6a, 0x47, 0xee, 0xf0,
	0xaf, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x92, 0x00, 0x61, 0x80, 0x07, 0x00, 0x00,
}
