// Code generated by protoc-gen-go.
// source: common_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	Mid                *string `protobuf:"bytes,1,opt,name=mid" json:"mid,omitempty"`
	Pwd                *string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Coin               *int64  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                 *uint32 `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName           *string `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	LoginTurntable     *bool   `protobuf:"varint,6,opt,name=loginTurntable" json:"loginTurntable,omitempty"`
	LoginTurntableTime *string `protobuf:"bytes,7,opt,name=loginTurntableTime" json:"loginTurntableTime,omitempty"`
	Scores             *int32  `protobuf:"varint,8,opt,name=scores" json:"scores,omitempty"`
	LastSignTime       *string `protobuf:"bytes,9,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignCount          *int32  `protobuf:"varint,10,opt,name=signCount" json:"signCount,omitempty"`
	Diamond            *int64  `protobuf:"varint,11,opt,name=Diamond" json:"Diamond,omitempty"`
	OpenId             *string `protobuf:"bytes,12,opt,name=openId" json:"openId,omitempty"`
	HeadUrl            *string `protobuf:"bytes,13,opt,name=headUrl" json:"headUrl,omitempty"`
	City               *string `protobuf:"bytes,14,opt,name=city" json:"city,omitempty"`
	Sex                *int32  `protobuf:"varint,15,opt,name=sex" json:"sex,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *User) GetMid() string {
	if m != nil && m.Mid != nil {
		return *m.Mid
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *User) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *User) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *User) GetLoginTurntable() bool {
	if m != nil && m.LoginTurntable != nil {
		return *m.LoginTurntable
	}
	return false
}

func (m *User) GetLoginTurntableTime() string {
	if m != nil && m.LoginTurntableTime != nil {
		return *m.LoginTurntableTime
	}
	return ""
}

func (m *User) GetScores() int32 {
	if m != nil && m.Scores != nil {
		return *m.Scores
	}
	return 0
}

func (m *User) GetLastSignTime() string {
	if m != nil && m.LastSignTime != nil {
		return *m.LastSignTime
	}
	return ""
}

func (m *User) GetSignCount() int32 {
	if m != nil && m.SignCount != nil {
		return *m.SignCount
	}
	return 0
}

func (m *User) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *User) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *User) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

type TNotice struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NoticeType       *int32   `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string  `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string  `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string  `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Fileds           []string `protobuf:"bytes,6,rep,name=fileds" json:"fileds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TNotice) Reset()                    { *m = TNotice{} }
func (m *TNotice) String() string            { return proto.CompactTextString(m) }
func (*TNotice) ProtoMessage()               {}
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *TNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *TNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *TNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *TNotice) GetFileds() []string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

// 用户的seesion 主要是游戏的
type GameSession struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	RoomId           *int32  `protobuf:"varint,2,opt,name=RoomId" json:"RoomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,3,opt,name=DeskId" json:"DeskId,omitempty"`
	GameStatus       *int32  `protobuf:"varint,4,opt,name=GameStatus" json:"GameStatus,omitempty"`
	GameId           *int32  `protobuf:"varint,5,opt,name=GameId" json:"GameId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,6,opt,name=GameNumber" json:"GameNumber,omitempty"`
	GameCustomStatus *int32  `protobuf:"varint,7,opt,name=GameCustomStatus" json:"GameCustomStatus,omitempty"`
	IsBreak          *bool   `protobuf:"varint,8,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,9,opt,name=isLeave" json:"isLeave,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameSession) Reset()                    { *m = GameSession{} }
func (m *GameSession) String() string            { return proto.CompactTextString(m) }
func (*GameSession) ProtoMessage()               {}
func (*GameSession) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GameSession) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *GameSession) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *GameSession) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *GameSession) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *GameSession) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameSession) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *GameSession) GetGameCustomStatus() int32 {
	if m != nil && m.GameCustomStatus != nil {
		return *m.GameCustomStatus
	}
	return 0
}

func (m *GameSession) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *GameSession) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

// 服务器游戏玩家通过的协议
type CommonSrvGameUser struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	IsBreak          *bool   `protobuf:"varint,3,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,4,opt,name=isLeave" json:"isLeave,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommonSrvGameUser) Reset()                    { *m = CommonSrvGameUser{} }
func (m *CommonSrvGameUser) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvGameUser) ProtoMessage()               {}
func (*CommonSrvGameUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CommonSrvGameUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CommonSrvGameUser) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *CommonSrvGameUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *CommonSrvGameUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "ddproto.User")
	proto.RegisterType((*TNotice)(nil), "ddproto.TNotice")
	proto.RegisterType((*GameSession)(nil), "ddproto.GameSession")
	proto.RegisterType((*CommonSrvGameUser)(nil), "ddproto.common_srv_game_user")
}

var fileDescriptor1 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x52, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0xb2, 0xd9, 0xee, 0x66, 0xf2, 0xd3, 0xe2, 0x16, 0x64, 0x71, 0x42, 0x39, 0xf5, 0xc4,
	0x43, 0x90, 0x4a, 0x28, 0x12, 0xe4, 0x40, 0xd3, 0x0b, 0x97, 0x68, 0xbb, 0x1e, 0x82, 0x95, 0xb5,
	0x1d, 0xd9, 0xde, 0x40, 0x5e, 0x80, 0xf7, 0xe1, 0x0d, 0x19, 0x4f, 0x76, 0xab, 0xd0, 0x9b, 0xbf,
	0xcf, 0xdf, 0xcc, 0x7c, 0xf3, 0x03, 0xb7, 0xb5, 0x33, 0xc6, 0xd9, 0x6d, 0x40, 0x7f, 0x44, 0xff,
	0xf1, 0xe0, 0x5d, 0x74, 0xa2, 0x50, 0x8a, 0x1f, 0x8b, 0x3f, 0x43, 0x18, 0x3d, 0xd1, 0x97, 0x98,
	0x40, 0x66, 0xb4, 0x92, 0x83, 0x0f, 0x83, 0xfb, 0x71, 0x02, 0x87, 0x5f, 0x4a, 0x0e, 0x19, 0x4c,
	0x61, 0x54, 0x3b, 0x6d, 0x65, 0x46, 0x28, 0x13, 0x00, 0x43, 0x92, 0x8d, 0xe8, 0x3d, 0x13, 0x37,
	0x50, 0x5a, 0x5d, 0xef, 0xd7, 0x95, 0x41, 0x99, 0xb3, 0xf6, 0x1d, 0xcc, 0x1b, 0xb7, 0xd3, 0x76,
	0xd3, 0x7a, 0x1b, 0xab, 0xe7, 0x06, 0xe5, 0x15, 0xf1, 0xa5, 0x78, 0x0f, 0xe2, 0x7f, 0x7e, 0xa3,
	0x29, 0xa6, 0xe0, 0x98, 0x39, 0x5c, 0x85, 0xda, 0x79, 0x0c, 0xb2, 0x24, 0x9c, 0x8b, 0x3b, 0x98,
	0x36, 0x55, 0x88, 0x8f, 0x7a, 0x67, 0x59, 0x35, 0x66, 0xd5, 0x1b, 0x18, 0x07, 0x62, 0x96, 0xae,
	0xb5, 0x51, 0x02, 0x0b, 0xaf, 0xa1, 0x78, 0xd0, 0x15, 0xf5, 0xa6, 0xe4, 0x84, 0xbd, 0x51, 0x26,
	0x77, 0x40, 0xbb, 0x52, 0x72, 0xca, 0x31, 0x24, 0xf8, 0x89, 0x95, 0x7a, 0xf2, 0x8d, 0x9c, 0xbd,
	0xb4, 0xa2, 0xe3, 0x49, 0xce, 0xfb, 0x2e, 0x03, 0xfe, 0x96, 0xd7, 0x29, 0xd9, 0xe2, 0x04, 0xc5,
	0x66, 0xed, 0xa2, 0xae, 0xb1, 0x6b, 0x71, 0xc0, 0x35, 0x08, 0x58, 0x66, 0x37, 0xa7, 0x03, 0xf2,
	0x40, 0x72, 0x71, 0x0b, 0x93, 0x8e, 0xd3, 0x91, 0x3a, 0xcc, 0x38, 0xd9, 0x5b, 0x98, 0x9d, 0xc9,
	0xa5, 0xb3, 0x11, 0xc9, 0xe3, 0x88, 0xe9, 0x97, 0xf8, 0xaf, 0x68, 0x5c, 0x37, 0x24, 0xb2, 0xf9,
	0x43, 0x37, 0xa8, 0x02, 0x0d, 0x27, 0xbb, 0x1f, 0x2f, 0xfe, 0x0e, 0x60, 0xf2, 0x99, 0x66, 0xf8,
	0x88, 0x21, 0x68, 0x67, 0xd3, 0x7f, 0x5a, 0xc9, 0xea, 0xec, 0x61, 0x96, 0xf0, 0x37, 0xe7, 0xcc,
	0x4a, 0x75, 0xf5, 0x09, 0x3f, 0x60, 0xd8, 0x13, 0xce, 0x7a, 0x8f, 0x1c, 0x1e, 0xab, 0xd8, 0x06,
	0xae, 0xcb, 0x9a, 0xc4, 0x91, 0x26, 0xbf, 0xd4, 0xac, 0x5b, 0xf3, 0x8c, 0x9e, 0x97, 0x92, 0x0b,
	0x09, 0x37, 0x89, 0x5b, 0xb6, 0x21, 0x3a, 0xd3, 0x45, 0x17, 0xfd, 0x64, 0x75, 0xf8, 0xe4, 0xb1,
	0xda, 0xf3, 0x4e, 0xca, 0x33, 0xf1, 0x05, 0xab, 0xe3, 0x79, 0x1d, 0xe5, 0xe2, 0x3b, 0xdc, 0xf5,
	0x77, 0xe5, 0x8f, 0xdb, 0x1d, 0xa5, 0xd9, 0xb6, 0xe9, 0x8c, 0xa8, 0x6e, 0x7b, 0xe9, 0xfd, 0xf2,
	0x44, 0x86, 0xfd, 0x52, 0xfa, 0xdc, 0xd9, 0xeb, 0xdc, 0xc9, 0x7b, 0xf9, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x4c, 0xfe, 0x37, 0x39, 0xb2, 0x02, 0x00, 0x00,
}
