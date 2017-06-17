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

// 三打一换桌message
type SdyReqChangeDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsChangeDesk     *bool        `protobuf:"varint,2,opt,name=isChangeDesk" json:"isChangeDesk,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqChangeDesk) Reset()                    { *m = SdyReqChangeDesk{} }
func (m *SdyReqChangeDesk) String() string            { return proto.CompactTextString(m) }
func (*SdyReqChangeDesk) ProtoMessage()               {}
func (*SdyReqChangeDesk) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{21} }

func (m *SdyReqChangeDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqChangeDesk) GetIsChangeDesk() bool {
	if m != nil && m.IsChangeDesk != nil {
		return *m.IsChangeDesk
	}
	return false
}

type SdyAckChangeDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsChangeOk       *bool        `protobuf:"varint,2,opt,name=isChangeOk" json:"isChangeOk,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckChangeDesk) Reset()                    { *m = SdyAckChangeDesk{} }
func (m *SdyAckChangeDesk) String() string            { return proto.CompactTextString(m) }
func (*SdyAckChangeDesk) ProtoMessage()               {}
func (*SdyAckChangeDesk) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{22} }

func (m *SdyAckChangeDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckChangeDesk) GetIsChangeOk() bool {
	if m != nil && m.IsChangeOk != nil {
		return *m.IsChangeOk
	}
	return false
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
	proto.RegisterType((*SdyReqChangeDesk)(nil), "yjprotogo.sdy_req_changeDesk")
	proto.RegisterType((*SdyAckChangeDesk)(nil), "yjprotogo.sdy_ack_changeDesk")
	proto.RegisterEnum("yjprotogo.SdyEnumJScore", SdyEnumJScore_name, SdyEnumJScore_value)
}

var fileDescriptor11 = []byte{
<<<<<<< HEAD
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0x5f, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0xa5, 0x8e, 0xfb, 0x73, 0x36, 0x4d, 0x1b, 0xdc, 0x82, 0xa2, 0x0a, 0xa1, 0xca, 0x0f,
	0xa8, 0x20, 0x91, 0x87, 0xc2, 0x09, 0x68, 0x53, 0x9a, 0x02, 0xa9, 0xd5, 0x94, 0x8a, 0x56, 0x48,
	0xd1, 0xd6, 0xde, 0x26, 0xdb, 0xc4, 0xbb, 0x61, 0xd7, 0x6e, 0xc9, 0x23, 0x57, 0xe0, 0x02, 0x9c,
	0x82, 0x3b, 0x70, 0x2c, 0x66, 0x77, 0x9d, 0xbf, 0x2a, 0x10, 0x27, 0x12, 0x4f, 0xf1, 0xce, 0xce,
	0x7e, 0x33, 0xf3, 0xcd, 0xcc, 0x17, 0xb4, 0x2e, 0xc3, 0x41, 0xab, 0xdf, 0xc3, 0x83, 0x6a, 0x5f,
	0xf0, 0x98, 0xbb, 0x85, 0xc1, 0x8d, 0xfe, 0x68, 0xf3, 0xed, 0xcd, 0x80, 0x47, 0x11, 0x67, 0xad,
	0xa0, 0x47, 0x09, 0x8b, 0xcd, 0xfd, 0xb6, 0xf6, 0xbf, 0xc2, 0x92, 0x98, 0xb3, 0xf7, 0x23, 0x67,
	0x20, 0xae, 0x82, 0x16, 0xef, 0x13, 0x46, 0x59, 0xdb, 0x7d, 0x8a, 0x56, 0x3b, 0x04, 0x87, 0x44,
	0x54, 0x72, 0x3b, 0xb9, 0xdd, 0xe2, 0xde, 0xa3, 0xea, 0x08, 0xb3, 0xea, 0xab, 0xdf, 0x23, 0x7d,
	0xeb, 0xae, 0xa3, 0xd5, 0x44, 0x12, 0x51, 0x0f, 0x2b, 0x2b, 0xe0, 0x57, 0x52, 0xe7, 0x98, 0x06,
	0x5d, 0x12, 0x57, 0x2c, 0x38, 0x5b, 0x6e, 0x19, 0x39, 0xea, 0x7e, 0x9f, 0x53, 0x56, 0xc9, 0x6b,
	0xcb, 0x16, 0x5a, 0x53, 0x96, 0x53, 0xce, 0xa3, 0x7d, 0x2c, 0xc2, 0x8a, 0xad, 0xad, 0x2f, 0x51,
	0x51, 0x15, 0x00, 0x48, 0xec, 0x9a, 0xcb, 0xca, 0xea, 0x8e, 0x05, 0x41, 0x9f, 0x4c, 0x04, 0x1d,
	0xa6, 0xdc, 0x1a, 0xbb, 0x79, 0xef, 0x51, 0x49, 0x99, 0x43, 0x82, 0x7b, 0x0a, 0x4a, 0xce, 0x9d,
	0x35, 0xe4, 0x60, 0x60, 0x7c, 0xde, 0x25, 0x42, 0x42, 0xee, 0xd6, 0xae, 0xed, 0x49, 0xb4, 0x99,
	0xb2, 0x30, 0x79, 0x39, 0x37, 0xe8, 0xab, 0x7b, 0x40, 0x7f, 0x53, 0x83, 0x22, 0xc0, 0xc7, 0xb4,
	0x1e, 0x4a, 0xef, 0xdb, 0x98, 0xfb, 0x1b, 0x8a, 0xf9, 0x21, 0x61, 0x73, 0x07, 0x7c, 0x88, 0x4a,
	0x41, 0x22, 0x04, 0xf4, 0xf5, 0xc3, 0x64, 0x0b, 0x5c, 0x84, 0x18, 0xf9, 0x32, 0xb4, 0x59, 0xda,
	0xf6, 0x02, 0x15, 0x02, 0xcc, 0x8e, 0x9b, 0x01, 0x17, 0x04, 0xfa, 0x60, 0xed, 0xae, 0xef, 0x6d,
	0xcf, 0x24, 0x46, 0x58, 0x12, 0xb5, 0x6e, 0xb4, 0x87, 0x17, 0xa3, 0x0d, 0x65, 0x12, 0xe4, 0x73,
	0xe6, 0xa4, 0x66, 0x07, 0xe2, 0x19, 0xb2, 0xa5, 0x8e, 0xaa, 0x12, 0xf9, 0x73, 0xd4, 0xba, 0x89,
	0x8a, 0x83, 0xee, 0xb2, 0x51, 0xbd, 0x01, 0xda, 0x9a, 0x26, 0xf5, 0x94, 0xc8, 0xa4, 0x17, 0xff,
	0x8b, 0x2a, 0x8e, 0x51, 0x39, 0x0d, 0x2d, 0x63, 0x2c, 0x62, 0x1f, 0x46, 0x22, 0x4b, 0xd8, 0x2b,
	0xcc, 0x60, 0x7a, 0xd2, 0x32, 0xde, 0x19, 0x2c, 0xd5, 0x07, 0x9e, 0xc4, 0xd9, 0x66, 0x1c, 0x36,
	0x6f, 0xf8, 0x46, 0xa3, 0xd9, 0x9e, 0x30, 0x68, 0x8a, 0xdf, 0xcc, 0x68, 0xb3, 0x84, 0xc0, 0x06,
	0x09, 0x12, 0x61, 0xca, 0xd2, 0x61, 0x57, 0xbc, 0xd8, 0x53, 0x31, 0xf3, 0x3a, 0xe6, 0xf7, 0x9c,
	0x69, 0xaa, 0x92, 0x96, 0x5b, 0x22, 0xce, 0x12, 0xb1, 0xf8, 0x28, 0xad, 0xa1, 0x7c, 0x4c, 0x23,
	0x92, 0xc6, 0x7a, 0x80, 0x0a, 0x7d, 0x41, 0x4e, 0x92, 0x18, 0x16, 0xc9, 0x04, 0x9b, 0x15, 0x11,
	0x7b, 0x2e, 0x11, 0x79, 0x33, 0xda, 0xbf, 0xbb, 0x0e, 0x56, 0x60, 0x73, 0xe7, 0x57, 0x42, 0x76,
	0x5f, 0x2d, 0x71, 0x4a, 0xef, 0xd7, 0x71, 0xa9, 0x7a, 0x56, 0xb2, 0x40, 0x01, 0x71, 0xc3, 0x37,
	0x46, 0x8c, 0xdc, 0x0d, 0xf4, 0x3f, 0x95, 0x3e, 0x57, 0x06, 0x55, 0xaf, 0xa3, 0xd6, 0x9a, 0xca,
	0x66, 0x07, 0xb3, 0xf6, 0x7e, 0x87, 0xe8, 0x82, 0x1d, 0xe3, 0xf4, 0x96, 0x27, 0x07, 0x54, 0xcb,
	0xa8, 0xe3, 0xfd, 0x1c, 0xe7, 0xd0, 0xc6, 0x11, 0x51, 0x05, 0x66, 0x11, 0x45, 0x49, 0x18, 0x7c,
	0x4d, 0xa9, 0xc9, 0x26, 0x2a, 0x52, 0x79, 0x4a, 0x02, 0xce, 0x18, 0x09, 0xe2, 0x34, 0x97, 0x3d,
	0x84, 0xc6, 0x0c, 0x6a, 0x3d, 0xf9, 0x2b, 0xcf, 0x6e, 0x15, 0x39, 0x21, 0x91, 0x5d, 0xfd, 0xc2,
	0xd6, 0x89, 0x3c, 0xbe, 0xef, 0xc5, 0xd0, 0xc7, 0xfb, 0x34, 0xea, 0x4b, 0x08, 0x7f, 0x48, 0x97,
	0x9d, 0x64, 0xe1, 0xb9, 0x01, 0xe6, 0xae, 0x39, 0x8f, 0x47, 0x93, 0xaa, 0xb4, 0xfe, 0x62, 0xac,
	0x70, 0xcb, 0xc2, 0xc3, 0xf9, 0xba, 0xc7, 0xef, 0xe0, 0x9d, 0x1e, 0x4c, 0xaf, 0x31, 0xd2, 0x9e,
	0x14, 0x39, 0xbb, 0xf6, 0xa4, 0x78, 0x66, 0xae, 0x26, 0x64, 0x71, 0xc9, 0x54, 0x61, 0xd6, 0x4b,
	0x69, 0x6a, 0x9d, 0x04, 0xb3, 0x03, 0xba, 0x30, 0xd0, 0xa5, 0x69, 0x8e, 0xa2, 0x6f, 0x39, 0xa4,
	0x29, 0xc9, 0x30, 0xad, 0xe9, 0x1b, 0x6c, 0x55, 0xef, 0x92, 0xd8, 0x30, 0xd1, 0x54, 0x1e, 0x69,
	0x8c, 0x93, 0xdb, 0xb4, 0x3f, 0xce, 0x8c, 0x48, 0x41, 0xc4, 0xe7, 0x6d, 0xc3, 0xf0, 0x84, 0x8a,
	0xc3, 0x6e, 0x17, 0x9a, 0x07, 0x17, 0xad, 0x66, 0xfd, 0xe3, 0xd9, 0x45, 0x39, 0x07, 0x62, 0x53,
	0x1a, 0x1d, 0x0f, 0xeb, 0xe7, 0xb5, 0xf2, 0x0a, 0xec, 0x5e, 0x51, 0x9b, 0x6a, 0xe7, 0xb5, 0x06,
	0xf8, 0x58, 0xb0, 0x29, 0x1b, 0x13, 0x06, 0xed, 0x95, 0x07, 0xcd, 0x72, 0x94, 0xb1, 0x71, 0xd2,
	0xa8, 0x95, 0xed, 0xd7, 0x2b, 0x47, 0x96, 0xff, 0x9f, 0x9f, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x79, 0xdb, 0x2d, 0xb4, 0xb1, 0x09, 0x00, 0x00,
=======
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x66, 0x3c, 0x76, 0x70, 0xca, 0xf6, 0xae, 0xe9, 0xac, 0xa2, 0x91, 0xb5, 0x8a, 0xac, 0x16,
	0x07, 0x8b, 0x43, 0x84, 0x72, 0xda, 0x13, 0x42, 0x9b, 0x64, 0x15, 0x83, 0x94, 0x58, 0x93, 0x10,
	0xc8, 0x1e, 0xd6, 0xea, 0xcc, 0x74, 0xec, 0x8e, 0xed, 0x6e, 0x6f, 0xf7, 0xcc, 0xee, 0xe6, 0xc8,
	0x11, 0x71, 0xe1, 0x88, 0x78, 0x0b, 0xde, 0x80, 0x17, 0xe0, 0x65, 0x78, 0x02, 0xd4, 0x3f, 0xf3,
	0x67, 0x6d, 0x40, 0x36, 0x96, 0x38, 0xd9, 0xf5, 0x75, 0x75, 0xd5, 0x57, 0x3f, 0xf3, 0xcd, 0xc0,
	0x13, 0x15, 0x3f, 0x8c, 0x97, 0x73, 0xf2, 0x70, 0xb8, 0x94, 0x22, 0x11, 0x68, 0xf7, 0xe1, 0xde,
	0xfc, 0x99, 0x88, 0xde, 0x5e, 0x24, 0x16, 0x0b, 0xc1, 0xc7, 0xd1, 0x9c, 0x51, 0x9e, 0xd8, 0xf3,
	0x9e, 0xf1, 0xbf, 0x25, 0x8a, 0x5a, 0x1b, 0xff, 0xe5, 0xd9, 0x10, 0xb7, 0xd1, 0x58, 0x2c, 0x29,
	0x67, 0x7c, 0x82, 0x0e, 0x61, 0x67, 0x4a, 0x49, 0x4c, 0x65, 0xe0, 0xf5, 0xbd, 0x41, 0xeb, 0x68,
	0xff, 0x30, 0x8f, 0x79, 0x38, 0xd2, 0xbf, 0x67, 0xe6, 0x34, 0x74, 0x5e, 0x68, 0x1f, 0x76, 0x52,
	0x45, 0xe5, 0x30, 0x0e, 0x6a, 0x7d, 0x6f, 0xd0, 0x09, 0x9d, 0xa5, 0xf1, 0x84, 0x45, 0x33, 0x9a,
	0x04, 0x7e, 0xdf, 0x1b, 0xf8, 0xa1, 0xb3, 0x50, 0x0f, 0x9a, 0xda, 0xe3, 0x58, 0x30, 0x1e, 0xd4,
	0xcd, 0x49, 0x6e, 0x23, 0x0c, 0x6d, 0xfd, 0x3f, 0x14, 0x62, 0x71, 0x4c, 0x64, 0x1c, 0x34, 0xcc,
	0x79, 0x05, 0x43, 0x5f, 0x43, 0x4b, 0x17, 0x4c, 0xe5, 0x90, 0xdf, 0x09, 0x15, 0xec, 0xf4, 0xfd,
	0x41, 0xeb, 0xe8, 0xa0, 0x44, 0x32, 0x2b, 0x71, 0x5c, 0xb8, 0x85, 0xe5, 0x2b, 0x38, 0x82, 0x8e,
	0xf6, 0x89, 0x29, 0x99, 0xeb, 0x88, 0x6a, 0xed, 0x92, 0x31, 0xb4, 0x6d, 0xbc, 0x91, 0x98, 0x51,
	0xa9, 0x82, 0x5a, 0xdf, 0x1f, 0x34, 0xc2, 0x0a, 0x86, 0x7f, 0xf2, 0x60, 0xcf, 0x75, 0xb6, 0x8c,
	0xaf, 0x9d, 0xeb, 0xe5, 0x47, 0x72, 0x3d, 0x52, 0xaf, 0x6e, 0xd3, 0x88, 0xb0, 0x61, 0xac, 0x56,
	0xb8, 0xfc, 0x51, 0x4c, 0xf9, 0x9e, 0x11, 0xf1, 0x8a, 0xf2, 0xb5, 0x69, 0x7c, 0x0e, 0x9d, 0x28,
	0x95, 0x92, 0xf2, 0xe4, 0xbb, 0xf2, 0xb0, 0xab, 0x20, 0x3a, 0x00, 0xe0, 0xf4, 0x43, 0xe6, 0xe2,
	0x1b, 0x97, 0x12, 0x82, 0x5e, 0xc0, 0x6e, 0x44, 0xf8, 0x37, 0x97, 0x91, 0x90, 0x34, 0xa8, 0xf7,
	0xfd, 0xc1, 0x93, 0xa3, 0xde, 0x4a, 0x25, 0x94, 0xa7, 0x8b, 0xf1, 0xbd, 0xf1, 0x08, 0x0b, 0x67,
	0xfc, 0xb3, 0x07, 0x4f, 0xf5, 0xb1, 0xa4, 0x6f, 0x37, 0xae, 0xe1, 0xb1, 0x4d, 0xfd, 0x12, 0x1a,
	0xca, 0x30, 0xd2, 0x84, 0xff, 0x99, 0x91, 0x75, 0xc4, 0x37, 0x96, 0x0c, 0x89, 0x66, 0xdb, 0x26,
	0x83, 0x7f, 0xf1, 0xe0, 0x59, 0x75, 0x56, 0x21, 0x55, 0xe9, 0x3c, 0xf9, 0x1f, 0xab, 0x7d, 0x0d,
	0x5d, 0xc7, 0x48, 0x25, 0x44, 0x26, 0xa3, 0x39, 0x79, 0xd8, 0x84, 0xcd, 0x2d, 0xe1, 0x33, 0x2a,
	0x33, 0x36, 0xd6, 0xc2, 0x6f, 0x6c, 0x6c, 0x3d, 0x56, 0x91, 0x26, 0x9b, 0x3d, 0x8e, 0x3d, 0x68,
	0x66, 0x77, 0x4d, 0xf4, 0x46, 0x98, 0xdb, 0xf8, 0x37, 0xcf, 0x26, 0xd0, 0xa3, 0xda, 0x38, 0xc1,
	0x63, 0xad, 0xc4, 0xd0, 0x96, 0x74, 0x41, 0x18, 0x77, 0xcf, 0xa6, 0x6f, 0x92, 0x57, 0xb0, 0x0a,
	0xb9, 0xfa, 0x0a, 0xb9, 0x3f, 0xdd, 0x52, 0x6b, 0xf5, 0x7d, 0x47, 0xe5, 0x55, 0x2a, 0xb7, 0xb7,
	0xd4, 0x08, 0xea, 0x09, 0x5b, 0x50, 0xc7, 0xc9, 0xfc, 0x47, 0xcf, 0x61, 0x77, 0x29, 0xe9, 0x45,
	0x9a, 0x8c, 0x08, 0x73, 0x64, 0x0a, 0x60, 0x55, 0x58, 0x1b, 0xeb, 0x0b, 0xeb, 0x75, 0x2e, 0x33,
	0xef, 0xa7, 0xc4, 0xc4, 0x5c, 0xb7, 0x9a, 0x67, 0xd0, 0x58, 0x6a, 0x05, 0x73, 0x73, 0xb4, 0x06,
	0xfe, 0xbd, 0xe8, 0x93, 0xd9, 0xc8, 0x4d, 0x22, 0xf7, 0xa0, 0x99, 0xdd, 0x75, 0x7a, 0x9d, 0xdb,
	0x28, 0x80, 0x4f, 0x99, 0x1a, 0x09, 0x7d, 0xa4, 0xdb, 0xd5, 0x0c, 0x33, 0x53, 0x0b, 0x1a, 0x53,
	0x97, 0x53, 0xc2, 0x27, 0xc7, 0x53, 0x6a, 0x5a, 0xd6, 0x0c, 0x4b, 0x88, 0xbd, 0xf9, 0xad, 0x48,
	0x4f, 0x98, 0x79, 0x57, 0x99, 0x9b, 0xc6, 0xc4, 0x3f, 0xd6, 0x72, 0xce, 0x13, 0xb2, 0xa0, 0xba,
	0x41, 0x9b, 0xbc, 0x67, 0x14, 0xe5, 0x31, 0x95, 0x15, 0xcd, 0xad, 0x60, 0xa8, 0x0f, 0x2d, 0xa6,
	0x42, 0x1a, 0x09, 0xce, 0x69, 0x94, 0x38, 0xfe, 0x65, 0x08, 0x7d, 0x05, 0x50, 0x0c, 0xc9, 0xa8,
	0xee, 0xbf, 0x8f, 0xb5, 0x74, 0x03, 0xbd, 0x80, 0x66, 0x4c, 0xd5, 0xcc, 0xdc, 0x6e, 0x18, 0xde,
	0xcf, 0x3f, 0x76, 0x3b, 0xf3, 0x09, 0x73, 0x6f, 0xfc, 0x21, 0xdf, 0x87, 0x98, 0xf1, 0xc9, 0xeb,
	0x69, 0xba, 0xb5, 0xed, 0x3e, 0x00, 0xb8, 0x13, 0x22, 0xc9, 0x9f, 0x3b, 0x3d, 0xcf, 0x12, 0x82,
	0xdf, 0x16, 0x6f, 0x8b, 0x6d, 0xa7, 0xde, 0x87, 0x9d, 0xbb, 0xb9, 0x78, 0x4f, 0xa5, 0x7b, 0xb4,
	0x9c, 0x85, 0xdf, 0xe4, 0xba, 0xed, 0x32, 0x6e, 0xae, 0xdb, 0x2e, 0x7e, 0xad, 0x12, 0xbf, 0xf4,
	0xce, 0xd9, 0x72, 0x49, 0xf8, 0x7b, 0xfb, 0x41, 0x74, 0x1b, 0x8d, 0xa7, 0x29, 0xe1, 0x27, 0x6c,
	0x6b, 0x81, 0x13, 0xbb, 0x00, 0x7a, 0x0c, 0xdb, 0x8d, 0x5c, 0x91, 0x55, 0x3b, 0xfe, 0x42, 0x56,
	0x7f, 0x75, 0x9f, 0x3b, 0xba, 0x55, 0x5b, 0x4e, 0x8b, 0xa1, 0xcd, 0xd4, 0x99, 0x89, 0x79, 0xf1,
	0xce, 0xad, 0x40, 0x33, 0xac, 0x60, 0x2b, 0x8a, 0x5f, 0xa5, 0x36, 0x05, 0x94, 0x35, 0x24, 0xd2,
	0x1a, 0x42, 0x4f, 0xa8, 0x9a, 0x6d, 0xa2, 0x0b, 0x4c, 0x1d, 0xe7, 0xf7, 0x0d, 0x47, 0xc3, 0xa2,
	0xc0, 0x70, 0x6c, 0x33, 0xe9, 0x1e, 0xfc, 0x87, 0x4c, 0x46, 0xff, 0x6c, 0xd4, 0x8b, 0x2c, 0x4f,
	0x09, 0xf9, 0x62, 0x62, 0x97, 0xb2, 0xf4, 0xd1, 0x80, 0x3a, 0xb0, 0x7b, 0x79, 0x72, 0x33, 0xbe,
	0x1c, 0xfe, 0x70, 0x75, 0xd3, 0xf5, 0xd0, 0x67, 0xd0, 0xc9, 0xcd, 0x57, 0xc3, 0xeb, 0xd3, 0x6e,
	0x0d, 0x3d, 0x85, 0x96, 0x81, 0x4e, 0xaf, 0x4f, 0xcf, 0xaf, 0x6e, 0xba, 0x3e, 0xda, 0x83, 0xa7,
	0x25, 0xc0, 0x78, 0xd5, 0x51, 0x1b, 0x9a, 0x1a, 0x3c, 0xbf, 0x38, 0x3f, 0xed, 0x36, 0x5e, 0xd6,
	0xce, 0xfc, 0xd1, 0x27, 0x23, 0xef, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x26, 0x31, 0xae,
	0xf1, 0x0c, 0x00, 0x00,
>>>>>>> 1905a8d1dfb58c418a3623ced4808a60bfecf519
}
