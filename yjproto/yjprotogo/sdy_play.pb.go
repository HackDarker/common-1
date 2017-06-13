// Code generated by protoc-gen-go.
// source: sdy_play.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of cm_offline from common_client.proto

// Ignoring public import of cm_hearbeat from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of sdy_base_userPaiIds from sdy_base.proto

// Ignoring public import of sdy_base_roomTypeInfo from sdy_base.proto

// Ignoring public import of sdy_base_timerInfo from sdy_base.proto

// Ignoring public import of sdy_base_playerInfo from sdy_base.proto

// Ignoring public import of sdy_base_commonRateInfo from sdy_base.proto

// Ignoring public import of sdy_base_deskInfo from sdy_base.proto

// Ignoring public import of sdy_enum_protoId from sdy_base.proto

// Ignoring public import of sdy_enum_errorCode from sdy_base.proto

// Ignoring public import of sdy_enum_actType from sdy_base.proto

// Ignoring public import of sdy_enum_deskStatus from sdy_base.proto

// Ignoring public import of sdy_enum_userStatus from sdy_base.proto

// Ignoring public import of sdy_enum_Option from sdy_base.proto

// Ignoring public import of sdy_enum_roomLevel from sdy_base.proto

// Ignoring public import of sdy_enum_flowers from sdy_base.proto

type SdyEnumJScore int32

const (
	SdyEnumJScore_SDY_SIXTY       SdyEnumJScore = 1
	SdyEnumJScore_SDY_SIXTYFIVE   SdyEnumJScore = 2
	SdyEnumJScore_SDY_SEVENTY     SdyEnumJScore = 3
	SdyEnumJScore_SDY_SEVENTYFIVE SdyEnumJScore = 4
	SdyEnumJScore_SDY_NONE        SdyEnumJScore = 5
)

var SdyEnumJScore_name = map[int32]string{
	1: "SDY_SIXTY",
	2: "SDY_SIXTYFIVE",
	3: "SDY_SEVENTY",
	4: "SDY_SEVENTYFIVE",
	5: "SDY_NONE",
}
var SdyEnumJScore_value = map[string]int32{
	"SDY_SIXTY":       1,
	"SDY_SIXTYFIVE":   2,
	"SDY_SEVENTY":     3,
	"SDY_SEVENTYFIVE": 4,
	"SDY_NONE":        5,
}

func (x SdyEnumJScore) Enum() *SdyEnumJScore {
	p := new(SdyEnumJScore)
	*p = x
	return p
}
func (x SdyEnumJScore) String() string {
	return proto.EnumName(SdyEnumJScore_name, int32(x))
}
func (x *SdyEnumJScore) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SdyEnumJScore_value, data, "SdyEnumJScore")
	if err != nil {
		return err
	}
	*x = SdyEnumJScore(value)
	return nil
}
func (SdyEnumJScore) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

// 开局（接收服务端消息）
type SdyBcOpening struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Ticket           *int64               `protobuf:"varint,3,opt,name=ticket" json:"ticket,omitempty"`
	UserCoin         *int64               `protobuf:"varint,4,opt,name=userCoin" json:"userCoin,omitempty"`
	UserRoomCard     *int64               `protobuf:"varint,5,opt,name=userRoomCard" json:"userRoomCard,omitempty"`
	PlayerInfos      []*SdyBasePlayerInfo `protobuf:"bytes,6,rep,name=playerInfos" json:"playerInfos,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcOpening) Reset()                    { *m = SdyBcOpening{} }
func (m *SdyBcOpening) String() string            { return proto.CompactTextString(m) }
func (*SdyBcOpening) ProtoMessage()               {}
func (*SdyBcOpening) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *SdyBcOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcOpening) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcOpening) GetTicket() int64 {
	if m != nil && m.Ticket != nil {
		return *m.Ticket
	}
	return 0
}

func (m *SdyBcOpening) GetUserCoin() int64 {
	if m != nil && m.UserCoin != nil {
		return *m.UserCoin
	}
	return 0
}

func (m *SdyBcOpening) GetUserRoomCard() int64 {
	if m != nil && m.UserRoomCard != nil {
		return *m.UserRoomCard
	}
	return 0
}

func (m *SdyBcOpening) GetPlayerInfos() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

// 发牌
type SdyDealCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerPokers     []int32      `protobuf:"varint,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyDealCards) Reset()                    { *m = SdyDealCards{} }
func (m *SdyDealCards) String() string            { return proto.CompactTextString(m) }
func (*SdyDealCards) ProtoMessage()               {}
func (*SdyDealCards) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *SdyDealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyDealCards) GetPlayerPokers() []int32 {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

type SdyBcPlayerPokers struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerPokers     []*SdyBaseUserPaiIds `protobuf:"bytes,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcPlayerPokers) Reset()                    { *m = SdyBcPlayerPokers{} }
func (m *SdyBcPlayerPokers) String() string            { return proto.CompactTextString(m) }
func (*SdyBcPlayerPokers) ProtoMessage()               {}
func (*SdyBcPlayerPokers) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *SdyBcPlayerPokers) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcPlayerPokers) GetPlayerPokers() []*SdyBaseUserPaiIds {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

// 广播叫分的消息，包含当前谁叫分，下一个叫分的玩家
type SdyBcJiaoFen struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CurrentUserId    *uint32         `protobuf:"varint,2,opt,name=currentUserId" json:"currentUserId,omitempty"`
	NextUserId       *uint32         `protobuf:"varint,3,opt,name=nextUserId" json:"nextUserId,omitempty"`
	CanJScore        []SdyEnumJScore `protobuf:"varint,4,rep,name=canJScore,enum=yjprotogo.SdyEnumJScore" json:"canJScore,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SdyBcJiaoFen) Reset()                    { *m = SdyBcJiaoFen{} }
func (m *SdyBcJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyBcJiaoFen) ProtoMessage()               {}
func (*SdyBcJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *SdyBcJiaoFen) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcJiaoFen) GetCurrentUserId() uint32 {
	if m != nil && m.CurrentUserId != nil {
		return *m.CurrentUserId
	}
	return 0
}

func (m *SdyBcJiaoFen) GetNextUserId() uint32 {
	if m != nil && m.NextUserId != nil {
		return *m.NextUserId
	}
	return 0
}

func (m *SdyBcJiaoFen) GetCanJScore() []SdyEnumJScore {
	if m != nil {
		return m.CanJScore
	}
	return nil
}

// 叫分,玩家选择分数后由客户端请求叫分的消息
type SdyReqJiaoFen struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32        `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Score            *SdyEnumJScore `protobuf:"varint,3,opt,name=score,enum=yjprotogo.SdyEnumJScore" json:"score,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SdyReqJiaoFen) Reset()                    { *m = SdyReqJiaoFen{} }
func (m *SdyReqJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyReqJiaoFen) ProtoMessage()               {}
func (*SdyReqJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *SdyReqJiaoFen) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqJiaoFen) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReqJiaoFen) GetScore() SdyEnumJScore {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return SdyEnumJScore_SDY_SIXTY
}

// 叫分回复，失败时回复
type SdyAckJiaoFen struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckJiaoFen) Reset()                    { *m = SdyAckJiaoFen{} }
func (m *SdyAckJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyAckJiaoFen) ProtoMessage()               {}
func (*SdyAckJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *SdyAckJiaoFen) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckJiaoFen) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 叫分成功后广播叫分的结果
type SdyBcJiaoFenResult struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32        `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Score            *SdyEnumJScore `protobuf:"varint,3,opt,name=score,enum=yjprotogo.SdyEnumJScore" json:"score,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SdyBcJiaoFenResult) Reset()                    { *m = SdyBcJiaoFenResult{} }
func (m *SdyBcJiaoFenResult) String() string            { return proto.CompactTextString(m) }
func (*SdyBcJiaoFenResult) ProtoMessage()               {}
func (*SdyBcJiaoFenResult) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *SdyBcJiaoFenResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcJiaoFenResult) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcJiaoFenResult) GetScore() SdyEnumJScore {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return SdyEnumJScore_SDY_SIXTY
}

// 换底结束，开始游戏 (广播)
type SdyBcStartPlay struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Banker           *uint32      `protobuf:"varint,2,opt,name=banker" json:"banker,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcStartPlay) Reset()                    { *m = SdyBcStartPlay{} }
func (m *SdyBcStartPlay) String() string            { return proto.CompactTextString(m) }
func (*SdyBcStartPlay) ProtoMessage()               {}
func (*SdyBcStartPlay) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *SdyBcStartPlay) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcStartPlay) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

// 出牌
type SdyReqOutCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	OutCards         *int32       `protobuf:"varint,2,opt,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqOutCards) Reset()                    { *m = SdyReqOutCards{} }
func (m *SdyReqOutCards) String() string            { return proto.CompactTextString(m) }
func (*SdyReqOutCards) ProtoMessage()               {}
func (*SdyReqOutCards) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *SdyReqOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqOutCards) GetOutCards() int32 {
	if m != nil && m.OutCards != nil {
		return *m.OutCards
	}
	return 0
}

type SdyAckOutCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	RemainPokers     *int32       `protobuf:"varint,3,opt,name=remainPokers" json:"remainPokers,omitempty"`
	OutCards         *int32       `protobuf:"varint,4,opt,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckOutCards) Reset()                    { *m = SdyAckOutCards{} }
func (m *SdyAckOutCards) String() string            { return proto.CompactTextString(m) }
func (*SdyAckOutCards) ProtoMessage()               {}
func (*SdyAckOutCards) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *SdyAckOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckOutCards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyAckOutCards) GetRemainPokers() int32 {
	if m != nil && m.RemainPokers != nil {
		return *m.RemainPokers
	}
	return 0
}

func (m *SdyAckOutCards) GetOutCards() int32 {
	if m != nil && m.OutCards != nil {
		return *m.OutCards
	}
	return 0
}

// 轮到谁操作
type SdyBcOverTurn struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Time             *int32               `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	PreOutPai        *int32               `protobuf:"varint,4,opt,name=preOutPai" json:"preOutPai,omitempty"`
	PlayerInfos      []*SdyBasePlayerInfo `protobuf:"bytes,5,rep,name=playerInfos" json:"playerInfos,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcOverTurn) Reset()                    { *m = SdyBcOverTurn{} }
func (m *SdyBcOverTurn) String() string            { return proto.CompactTextString(m) }
func (*SdyBcOverTurn) ProtoMessage()               {}
func (*SdyBcOverTurn) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

func (m *SdyBcOverTurn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcOverTurn) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcOverTurn) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *SdyBcOverTurn) GetPreOutPai() int32 {
	if m != nil && m.PreOutPai != nil {
		return *m.PreOutPai
	}
	return 0
}

func (m *SdyBcOverTurn) GetPlayerInfos() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

// 广播上一个玩家出了什么牌的消息
type SdyBcWhatPai struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PaiId            *int32       `protobuf:"varint,2,opt,name=paiId" json:"paiId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcWhatPai) Reset()                    { *m = SdyBcWhatPai{} }
func (m *SdyBcWhatPai) String() string            { return proto.CompactTextString(m) }
func (*SdyBcWhatPai) ProtoMessage()               {}
func (*SdyBcWhatPai) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{11} }

func (m *SdyBcWhatPai) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcWhatPai) GetPaiId() int32 {
	if m != nil && m.PaiId != nil {
		return *m.PaiId
	}
	return 0
}

// 一圈打完后发送是否闲家得分的广播，以及得的多少分
type SdyBcScorePai struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ScorePai         []int32      `protobuf:"varint,2,rep,name=scorePai" json:"scorePai,omitempty"`
	IsPoPai          *bool        `protobuf:"varint,3,opt,name=isPoPai" json:"isPoPai,omitempty"`
	IsShangChe       *bool        `protobuf:"varint,4,opt,name=isShangChe" json:"isShangChe,omitempty"`
	IsKouDi          *bool        `protobuf:"varint,5,opt,name=isKouDi" json:"isKouDi,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcScorePai) Reset()                    { *m = SdyBcScorePai{} }
func (m *SdyBcScorePai) String() string            { return proto.CompactTextString(m) }
func (*SdyBcScorePai) ProtoMessage()               {}
func (*SdyBcScorePai) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{12} }

func (m *SdyBcScorePai) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcScorePai) GetScorePai() []int32 {
	if m != nil {
		return m.ScorePai
	}
	return nil
}

func (m *SdyBcScorePai) GetIsPoPai() bool {
	if m != nil && m.IsPoPai != nil {
		return *m.IsPoPai
	}
	return false
}

func (m *SdyBcScorePai) GetIsShangChe() bool {
	if m != nil && m.IsShangChe != nil {
		return *m.IsShangChe
	}
	return false
}

func (m *SdyBcScorePai) GetIsKouDi() bool {
	if m != nil && m.IsKouDi != nil {
		return *m.IsKouDi
	}
	return false
}

// 游戏信息(广播)
type SdyBcGameInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	SenderUserId     *uint32              `protobuf:"varint,2,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *bool                `protobuf:"varint,3,opt,name=isReconnect" json:"isReconnect,omitempty"`
	PlayerInfo       []*SdyBasePlayerInfo `protobuf:"bytes,4,rep,name=playerInfo" json:"playerInfo,omitempty"`
	DeskInfo         *SdyBaseDeskInfo     `protobuf:"bytes,5,opt,name=deskInfo" json:"deskInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcGameInfo) Reset()                    { *m = SdyBcGameInfo{} }
func (m *SdyBcGameInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBcGameInfo) ProtoMessage()               {}
func (*SdyBcGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{13} }

func (m *SdyBcGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *SdyBcGameInfo) GetIsReconnect() bool {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return false
}

func (m *SdyBcGameInfo) GetPlayerInfo() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *SdyBcGameInfo) GetDeskInfo() *SdyBaseDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

// 定主， 广播消息，通知其他玩家谁在定主，包含定主的人
type SdyBcDingZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	FootPokers       []int32      `protobuf:"varint,3,rep,name=footPokers" json:"footPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcDingZhu) Reset()                    { *m = SdyBcDingZhu{} }
func (m *SdyBcDingZhu) String() string            { return proto.CompactTextString(m) }
func (*SdyBcDingZhu) ProtoMessage()               {}
func (*SdyBcDingZhu) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{14} }

func (m *SdyBcDingZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcDingZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcDingZhu) GetFootPokers() []int32 {
	if m != nil {
		return m.FootPokers
	}
	return nil
}

// 选好花色后请求定主
type SdyReqDingZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Flower           *int32       `protobuf:"varint,3,opt,name=flower" json:"flower,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqDingZhu) Reset()                    { *m = SdyReqDingZhu{} }
func (m *SdyReqDingZhu) String() string            { return proto.CompactTextString(m) }
func (*SdyReqDingZhu) ProtoMessage()               {}
func (*SdyReqDingZhu) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{15} }

func (m *SdyReqDingZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqDingZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReqDingZhu) GetFlower() int32 {
	if m != nil && m.Flower != nil {
		return *m.Flower
	}
	return 0
}

type SdyBcDingZhuResult struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Flower           *int32       `protobuf:"varint,2,opt,name=flower" json:"flower,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcDingZhuResult) Reset()                    { *m = SdyBcDingZhuResult{} }
func (m *SdyBcDingZhuResult) String() string            { return proto.CompactTextString(m) }
func (*SdyBcDingZhuResult) ProtoMessage()               {}
func (*SdyBcDingZhuResult) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{16} }

func (m *SdyBcDingZhuResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcDingZhuResult) GetFlower() int32 {
	if m != nil && m.Flower != nil {
		return *m.Flower
	}
	return 0
}

// 定主的回复
type SdyAckDingZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckDingZhu) Reset()                    { *m = SdyAckDingZhu{} }
func (m *SdyAckDingZhu) String() string            { return proto.CompactTextString(m) }
func (*SdyAckDingZhu) ProtoMessage()               {}
func (*SdyAckDingZhu) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{17} }

func (m *SdyAckDingZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckDingZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 广播换底，定主完毕后广播换底，
type SdyBcHuanDi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcHuanDi) Reset()                    { *m = SdyBcHuanDi{} }
func (m *SdyBcHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyBcHuanDi) ProtoMessage()               {}
func (*SdyBcHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{18} }

func (m *SdyBcHuanDi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcHuanDi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type SdyReqHuanDi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	OutCards         []int32      `protobuf:"varint,3,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqHuanDi) Reset()                    { *m = SdyReqHuanDi{} }
func (m *SdyReqHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyReqHuanDi) ProtoMessage()               {}
func (*SdyReqHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{19} }

func (m *SdyReqHuanDi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqHuanDi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReqHuanDi) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

type SdyAckHuanDi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsHuanDiOver     *bool        `protobuf:"varint,3,opt,name=isHuanDiOver" json:"isHuanDiOver,omitempty"`
	OutCards         []int32      `protobuf:"varint,4,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckHuanDi) Reset()                    { *m = SdyAckHuanDi{} }
func (m *SdyAckHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyAckHuanDi) ProtoMessage()               {}
func (*SdyAckHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{20} }

func (m *SdyAckHuanDi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckHuanDi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyAckHuanDi) GetIsHuanDiOver() bool {
	if m != nil && m.IsHuanDiOver != nil {
		return *m.IsHuanDiOver
	}
	return false
}

func (m *SdyAckHuanDi) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

func init() {
	proto.RegisterType((*SdyBcOpening)(nil), "yjprotogo.sdy_bc_opening")
	proto.RegisterType((*SdyDealCards)(nil), "yjprotogo.sdy_dealCards")
	proto.RegisterType((*SdyBcPlayerPokers)(nil), "yjprotogo.sdy_bc_playerPokers")
	proto.RegisterType((*SdyBcJiaoFen)(nil), "yjprotogo.sdy_bc_jiaoFen")
	proto.RegisterType((*SdyReqJiaoFen)(nil), "yjprotogo.sdy_req_jiaoFen")
	proto.RegisterType((*SdyAckJiaoFen)(nil), "yjprotogo.sdy_ack_jiaoFen")
	proto.RegisterType((*SdyBcJiaoFenResult)(nil), "yjprotogo.sdy_bc_jiaoFenResult")
	proto.RegisterType((*SdyBcStartPlay)(nil), "yjprotogo.sdy_bc_startPlay")
	proto.RegisterType((*SdyReqOutCards)(nil), "yjprotogo.sdy_req_outCards")
	proto.RegisterType((*SdyAckOutCards)(nil), "yjprotogo.sdy_ack_outCards")
	proto.RegisterType((*SdyBcOverTurn)(nil), "yjprotogo.sdy_bc_overTurn")
	proto.RegisterType((*SdyBcWhatPai)(nil), "yjprotogo.sdy_bc_whatPai")
	proto.RegisterType((*SdyBcScorePai)(nil), "yjprotogo.sdy_bc_scorePai")
	proto.RegisterType((*SdyBcGameInfo)(nil), "yjprotogo.sdy_bc_gameInfo")
	proto.RegisterType((*SdyBcDingZhu)(nil), "yjprotogo.sdy_bc_dingZhu")
	proto.RegisterType((*SdyReqDingZhu)(nil), "yjprotogo.sdy_req_dingZhu")
	proto.RegisterType((*SdyBcDingZhuResult)(nil), "yjprotogo.sdy_bc_dingZhuResult")
	proto.RegisterType((*SdyAckDingZhu)(nil), "yjprotogo.sdy_ack_dingZhu")
	proto.RegisterType((*SdyBcHuanDi)(nil), "yjprotogo.sdy_bc_huanDi")
	proto.RegisterType((*SdyReqHuanDi)(nil), "yjprotogo.sdy_req_huanDi")
	proto.RegisterType((*SdyAckHuanDi)(nil), "yjprotogo.sdy_ack_huanDi")
	proto.RegisterEnum("yjprotogo.SdyEnumJScore", SdyEnumJScore_name, SdyEnumJScore_value)
}

var fileDescriptor11 = []byte{
	// 860 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4d, 0x6e, 0x1b, 0x37,
	0x14, 0xee, 0x68, 0x24, 0x57, 0x7e, 0x92, 0x62, 0x95, 0x0e, 0x8c, 0x81, 0x10, 0x18, 0x02, 0xd1,
	0x85, 0xd0, 0x85, 0x51, 0x78, 0x95, 0x55, 0x51, 0xc4, 0x76, 0x60, 0xb5, 0x80, 0x3d, 0x18, 0xbb,
	0x6e, 0x9d, 0x45, 0x04, 0x7a, 0x86, 0x96, 0x68, 0x49, 0xa4, 0x42, 0xce, 0x24, 0xd1, 0xb2, 0xcb,
	0xa2, 0x9b, 0x2e, 0x8b, 0xde, 0xa2, 0x37, 0xe8, 0x05, 0x7a, 0x99, 0x9e, 0xa0, 0x20, 0x87, 0x9a,
	0x1f, 0xc1, 0x6e, 0x21, 0x61, 0x80, 0xac, 0x34, 0xef, 0xe3, 0xfb, 0xf9, 0xde, 0x0f, 0x1f, 0x05,
	0xcf, 0x54, 0xb4, 0x1c, 0x2d, 0x66, 0x64, 0x79, 0xb4, 0x90, 0x22, 0x16, 0x68, 0x77, 0xf9, 0x60,
	0x3e, 0xc6, 0xa2, 0xb7, 0x1f, 0x8a, 0xf9, 0x5c, 0xf0, 0x51, 0x38, 0x63, 0x94, 0xc7, 0xe9, 0x79,
	0xcf, 0xe8, 0xdf, 0x11, 0x45, 0x53, 0x19, 0xff, 0xe3, 0xa4, 0x2e, 0xee, 0xc2, 0x91, 0x58, 0x50,
	0xce, 0xf8, 0x18, 0x1d, 0xc1, 0xce, 0x84, 0x92, 0x88, 0x4a, 0xcf, 0xe9, 0x3b, 0x83, 0xd6, 0xf1,
	0xc1, 0x51, 0xe6, 0xf3, 0xc8, 0xd7, 0xbf, 0xe7, 0xe6, 0x34, 0xb0, 0x5a, 0xe8, 0x00, 0x76, 0x12,
	0x45, 0xe5, 0x30, 0xf2, 0x6a, 0x7d, 0x67, 0xd0, 0x09, 0xac, 0xa4, 0xf1, 0x98, 0x85, 0x53, 0x1a,
	0x7b, 0x6e, 0xdf, 0x19, 0xb8, 0x81, 0x95, 0x50, 0x0f, 0x9a, 0x5a, 0xe3, 0x44, 0x30, 0xee, 0xd5,
	0xcd, 0x49, 0x26, 0x23, 0x0c, 0x6d, 0xfd, 0x1d, 0x08, 0x31, 0x3f, 0x21, 0x32, 0xf2, 0x1a, 0xe6,
	0xbc, 0x84, 0xa1, 0x6f, 0xa1, 0xa5, 0x13, 0xa6, 0x72, 0xc8, 0xef, 0x85, 0xf2, 0x76, 0xfa, 0xee,
	0xa0, 0x75, 0x7c, 0x58, 0x20, 0xb9, 0x4a, 0x71, 0x94, 0xab, 0x05, 0x45, 0x13, 0x1c, 0x42, 0x47,
	0xeb, 0x44, 0x94, 0xcc, 0xb4, 0x47, 0xb5, 0x71, 0xca, 0x18, 0xda, 0xa9, 0x3f, 0x5f, 0x4c, 0xa9,
	0x54, 0x5e, 0xad, 0xef, 0x0e, 0x1a, 0x41, 0x09, 0xc3, 0xbf, 0x38, 0xb0, 0x6f, 0x2b, 0x5b, 0xc4,
	0x37, 0x8e, 0xf5, 0xea, 0x91, 0x58, 0x4f, 0xe4, 0xab, 0xcb, 0xe4, 0x13, 0x36, 0x8c, 0xd4, 0x1a,
	0x97, 0xbf, 0xf2, 0x2e, 0x3f, 0x30, 0x22, 0x5e, 0x53, 0xbe, 0x31, 0x8d, 0x2f, 0xa1, 0x13, 0x26,
	0x52, 0x52, 0x1e, 0xff, 0x50, 0x6c, 0x76, 0x19, 0x44, 0x87, 0x00, 0x9c, 0x7e, 0x5c, 0xa9, 0xb8,
	0x46, 0xa5, 0x80, 0xa0, 0x97, 0xb0, 0x1b, 0x12, 0xfe, 0xdd, 0x55, 0x28, 0x24, 0xf5, 0xea, 0x7d,
	0x77, 0xf0, 0xec, 0xb8, 0xb7, 0x96, 0x09, 0xe5, 0xc9, 0x7c, 0xf4, 0x60, 0x34, 0x82, 0x5c, 0x19,
	0xff, 0xea, 0xc0, 0x9e, 0x3e, 0x96, 0xf4, 0xdd, 0xd6, 0x39, 0x3c, 0x35, 0xa9, 0x5f, 0x43, 0x43,
	0x19, 0x46, 0x9a, 0xf0, 0x7f, 0x33, 0x4a, 0x15, 0xf1, 0x6d, 0x4a, 0x86, 0x84, 0xd3, 0xaa, 0xc9,
	0xe0, 0xdf, 0x1c, 0x78, 0x5e, 0xee, 0x55, 0x40, 0x55, 0x32, 0x8b, 0x3f, 0x61, 0xb6, 0x6f, 0xa0,
	0x6b, 0x19, 0xa9, 0x98, 0xc8, 0xd8, 0x9f, 0x91, 0xe5, 0x36, 0x6c, 0xee, 0x08, 0x9f, 0x52, 0xb9,
	0x62, 0x93, 0x4a, 0xf8, 0x6d, 0xea, 0x5b, 0xb7, 0x55, 0x24, 0xf1, 0x76, 0xd7, 0xb1, 0x07, 0xcd,
	0x95, 0xad, 0xf1, 0xde, 0x08, 0x32, 0x19, 0xff, 0xe1, 0xa4, 0x01, 0x74, 0xab, 0xb6, 0x0e, 0xf0,
	0x54, 0x29, 0x31, 0xb4, 0x25, 0x9d, 0x13, 0xc6, 0xed, 0xdd, 0x74, 0x4d, 0xf0, 0x12, 0x56, 0x22,
	0x57, 0x5f, 0x23, 0xf7, 0xb7, 0x1d, 0x6a, 0xbd, 0x7d, 0xdf, 0x53, 0x79, 0x9d, 0xc8, 0xea, 0x86,
	0x1a, 0x41, 0x3d, 0x66, 0x73, 0x6a, 0x39, 0x99, 0x6f, 0xf4, 0x02, 0x76, 0x17, 0x92, 0x5e, 0x26,
	0xb1, 0x4f, 0x98, 0x25, 0x93, 0x03, 0xeb, 0x8b, 0xb5, 0xb1, 0xf9, 0x62, 0xbd, 0xc9, 0xd6, 0xcc,
	0x87, 0x09, 0x31, 0x3e, 0x37, 0xcd, 0xe6, 0x39, 0x34, 0x16, 0x7a, 0x83, 0xd9, 0x3e, 0xa6, 0x02,
	0xfe, 0x33, 0xaf, 0x93, 0x99, 0xc8, 0x6d, 0x3c, 0xf7, 0xa0, 0xb9, 0xb2, 0xb5, 0xfb, 0x3a, 0x93,
	0x91, 0x07, 0x9f, 0x33, 0xe5, 0x0b, 0x7d, 0xa4, 0xcb, 0xd5, 0x0c, 0x56, 0xa2, 0x5e, 0x68, 0x4c,
	0x5d, 0x4d, 0x08, 0x1f, 0x9f, 0x4c, 0xa8, 0x29, 0x59, 0x33, 0x28, 0x20, 0xa9, 0xe5, 0xf7, 0x22,
	0x39, 0x65, 0xe6, 0xad, 0x32, 0x96, 0x46, 0xc4, 0x3f, 0xd7, 0x32, 0xce, 0x63, 0x32, 0xa7, 0xba,
	0x40, 0xdb, 0xbc, 0x33, 0x8a, 0xf2, 0x88, 0xca, 0xd2, 0xce, 0x2d, 0x61, 0xa8, 0x0f, 0x2d, 0xa6,
	0x02, 0x1a, 0x0a, 0xce, 0x69, 0x18, 0x5b, 0xfe, 0x45, 0x08, 0x7d, 0x03, 0x90, 0x37, 0xc9, 0x6c,
	0xdd, 0xff, 0x6f, 0x6b, 0xc1, 0x02, 0xbd, 0x84, 0x66, 0x44, 0xd5, 0xd4, 0x58, 0x37, 0x0c, 0xef,
	0x17, 0x8f, 0x59, 0xaf, 0x74, 0x82, 0x4c, 0x1b, 0x7f, 0xcc, 0xe6, 0x21, 0x62, 0x7c, 0xfc, 0x66,
	0x92, 0x54, 0x36, 0xdd, 0x87, 0x00, 0xf7, 0x42, 0xc4, 0xd9, 0xbd, 0xd3, 0xfd, 0x2c, 0x20, 0xf8,
	0x5d, 0xfe, 0x5a, 0x54, 0x1d, 0xfa, 0x00, 0x76, 0xee, 0x67, 0xe2, 0x03, 0x95, 0xf6, 0x6a, 0x59,
	0x09, 0xbf, 0xcd, 0xf6, 0xb6, 0x8d, 0xb8, 0xfd, 0xde, 0xb6, 0xfe, 0x6b, 0x25, 0xff, 0x85, 0x37,
	0xa7, 0xe2, 0x94, 0xf0, 0x8f, 0xe9, 0x1f, 0xa2, 0xbb, 0x70, 0x34, 0x49, 0x08, 0x3f, 0x65, 0x95,
	0x39, 0x8e, 0xd3, 0x01, 0xd0, 0x6d, 0xa8, 0xd6, 0x73, 0x69, 0xad, 0xa6, 0xed, 0xcf, 0xd7, 0xea,
	0xef, 0xf6, 0xef, 0x8e, 0x2e, 0x55, 0xc5, 0x61, 0x31, 0xb4, 0x99, 0x3a, 0x37, 0x3e, 0x2f, 0xdf,
	0xdb, 0x11, 0x68, 0x06, 0x25, 0x6c, 0x6d, 0xe3, 0x97, 0xa8, 0x7d, 0x35, 0x4e, 0x9b, 0x58, 0x78,
	0x64, 0x51, 0x07, 0x76, 0xaf, 0x4e, 0x6f, 0x47, 0x57, 0xc3, 0x9f, 0xae, 0x6f, 0xbb, 0x0e, 0xfa,
	0x02, 0x3a, 0x99, 0xf8, 0x7a, 0x78, 0x73, 0xd6, 0xad, 0xa1, 0x3d, 0x68, 0x19, 0xe8, 0xec, 0xe6,
	0xec, 0xe2, 0xfa, 0xb6, 0xeb, 0xa2, 0x7d, 0xd8, 0x2b, 0x00, 0x46, 0xab, 0x8e, 0xda, 0xd0, 0xd4,
	0xe0, 0xc5, 0xe5, 0xc5, 0x59, 0xb7, 0xf1, 0xaa, 0x76, 0xee, 0xfa, 0x9f, 0xf9, 0xce, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x92, 0x0a, 0xae, 0x54, 0x21, 0x0c, 0x00, 0x00,
}
