// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_server.proto

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
	RoomType         *int32  `protobuf:"varint,10,opt,name=roomType" json:"roomType,omitempty"`
	RoomPassword     *string `protobuf:"bytes,11,opt,name=roomPassword" json:"roomPassword,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,12,opt,name=RoomLevel" json:"RoomLevel,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameSession) Reset()                    { *m = GameSession{} }
func (m *GameSession) String() string            { return proto.CompactTextString(m) }
func (*GameSession) ProtoMessage()               {}
func (*GameSession) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

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

func (m *GameSession) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *GameSession) GetRoomPassword() string {
	if m != nil && m.RoomPassword != nil {
		return *m.RoomPassword
	}
	return ""
}

func (m *GameSession) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

// 服务器游戏玩家通过的协议
type CommonSrvGameUser struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	Coin             *int64  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	RoomId           *int32  `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,5,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,6,opt,name=gameNumber" json:"gameNumber,omitempty"`
	IsBreak          *bool   `protobuf:"varint,7,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,8,opt,name=isLeave" json:"isLeave,omitempty"`
	Status           *int32  `protobuf:"varint,9,opt,name=status" json:"status,omitempty"`
	WaitTime         *int64  `protobuf:"varint,10,opt,name=waitTime" json:"waitTime,omitempty"`
	RoomType         *int32  `protobuf:"varint,11,opt,name=roomType" json:"roomType,omitempty"`
	RoomPassword     *string `protobuf:"bytes,12,opt,name=roomPassword" json:"roomPassword,omitempty"`
	IsRobot          *bool   `protobuf:"varint,13,opt,name=isRobot" json:"isRobot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommonSrvGameUser) Reset()                    { *m = CommonSrvGameUser{} }
func (m *CommonSrvGameUser) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvGameUser) ProtoMessage()               {}
func (*CommonSrvGameUser) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

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

func (m *CommonSrvGameUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *CommonSrvGameUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *CommonSrvGameUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *CommonSrvGameUser) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
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

func (m *CommonSrvGameUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *CommonSrvGameUser) GetWaitTime() int64 {
	if m != nil && m.WaitTime != nil {
		return *m.WaitTime
	}
	return 0
}

func (m *CommonSrvGameUser) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *CommonSrvGameUser) GetRoomPassword() string {
	if m != nil && m.RoomPassword != nil {
		return *m.RoomPassword
	}
	return ""
}

func (m *CommonSrvGameUser) GetIsRobot() bool {
	if m != nil && m.IsRobot != nil {
		return *m.IsRobot
	}
	return false
}

// 通用的desk 的协议
type CommonSrvGameDesk struct {
	DeskId           *int32  `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,2,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32  `protobuf:"varint,3,opt,name=roomId" json:"roomId,omitempty"`
	Round            *int32  `protobuf:"varint,4,opt,name=round" json:"round,omitempty"`
	Password         *string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	Owner            *uint32 `protobuf:"varint,6,opt,name=owner" json:"owner,omitempty"`
	Banker           *uint32 `protobuf:"varint,7,opt,name=banker" json:"banker,omitempty"`
	CreateFee        *int64  `protobuf:"varint,8,opt,name=createFee" json:"createFee,omitempty"`
	Status           *int32  `protobuf:"varint,9,opt,name=status" json:"status,omitempty"`
	BoardsCout       *int32  `protobuf:"varint,10,opt,name=boardsCout" json:"boardsCout,omitempty"`
	BaseValue        *int64  `protobuf:"varint,11,opt,name=baseValue" json:"baseValue,omitempty"`
	BeginTime        *string `protobuf:"bytes,12,opt,name=beginTime" json:"beginTime,omitempty"`
	EndTime          *string `protobuf:"bytes,13,opt,name=endTime" json:"endTime,omitempty"`
	NInitActionTime  *int32  `protobuf:"varint,14,opt,name=nInitActionTime" json:"nInitActionTime,omitempty"`
	UserCountLimit   *int32  `protobuf:"varint,15,opt,name=userCountLimit" json:"userCountLimit,omitempty"`
	IsDaikai         *bool   `protobuf:"varint,16,opt,name=isDaikai" json:"isDaikai,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommonSrvGameDesk) Reset()                    { *m = CommonSrvGameDesk{} }
func (m *CommonSrvGameDesk) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvGameDesk) ProtoMessage()               {}
func (*CommonSrvGameDesk) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *CommonSrvGameDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *CommonSrvGameDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *CommonSrvGameDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *CommonSrvGameDesk) GetRound() int32 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *CommonSrvGameDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CommonSrvGameDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *CommonSrvGameDesk) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

func (m *CommonSrvGameDesk) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *CommonSrvGameDesk) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *CommonSrvGameDesk) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *CommonSrvGameDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *CommonSrvGameDesk) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *CommonSrvGameDesk) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

func (m *CommonSrvGameDesk) GetNInitActionTime() int32 {
	if m != nil && m.NInitActionTime != nil {
		return *m.NInitActionTime
	}
	return 0
}

func (m *CommonSrvGameDesk) GetUserCountLimit() int32 {
	if m != nil && m.UserCountLimit != nil {
		return *m.UserCountLimit
	}
	return 0
}

func (m *CommonSrvGameDesk) GetIsDaikai() bool {
	if m != nil && m.IsDaikai != nil {
		return *m.IsDaikai
	}
	return false
}

// 房间代开，desk列表redis结构
type RedisDeskByAgent struct {
	Data             []*CommonDeskByAgent `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *RedisDeskByAgent) Reset()                    { *m = RedisDeskByAgent{} }
func (m *RedisDeskByAgent) String() string            { return proto.CompactTextString(m) }
func (*RedisDeskByAgent) ProtoMessage()               {}
func (*RedisDeskByAgent) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *RedisDeskByAgent) GetData() []*CommonDeskByAgent {
	if m != nil {
		return m.Data
	}
	return nil
}

// 玩家游戏账单数据 存储某一游戏某玩法下玩家的账单和所处的模式 redis结构
type RedisUserGameBill struct {
	Data             []*UserGameBill `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	IsWinMode        *bool           `protobuf:"varint,2,opt,name=isWinMode" json:"isWinMode,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RedisUserGameBill) Reset()                    { *m = RedisUserGameBill{} }
func (m *RedisUserGameBill) String() string            { return proto.CompactTextString(m) }
func (*RedisUserGameBill) ProtoMessage()               {}
func (*RedisUserGameBill) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *RedisUserGameBill) GetData() []*UserGameBill {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RedisUserGameBill) GetIsWinMode() bool {
	if m != nil && m.IsWinMode != nil {
		return *m.IsWinMode
	}
	return false
}

type UserGameBill struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	DeskId           *int32  `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	GameNumber       *int32  `protobuf:"varint,5,opt,name=gameNumber" json:"gameNumber,omitempty"`
	WinAmount        *int64  `protobuf:"varint,6,opt,name=winAmount" json:"winAmount,omitempty"`
	EndTime          *int64  `protobuf:"varint,7,opt,name=endTime" json:"endTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserGameBill) Reset()                    { *m = UserGameBill{} }
func (m *UserGameBill) String() string            { return proto.CompactTextString(m) }
func (*UserGameBill) ProtoMessage()               {}
func (*UserGameBill) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *UserGameBill) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *UserGameBill) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *UserGameBill) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *UserGameBill) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *UserGameBill) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *UserGameBill) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

func (m *UserGameBill) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func init() {
	proto.RegisterType((*GameSession)(nil), "ddproto.GameSession")
	proto.RegisterType((*CommonSrvGameUser)(nil), "ddproto.common_srv_game_user")
	proto.RegisterType((*CommonSrvGameDesk)(nil), "ddproto.common_srv_game_desk")
	proto.RegisterType((*RedisDeskByAgent)(nil), "ddproto.redis_desk_by_agent")
	proto.RegisterType((*RedisUserGameBill)(nil), "ddproto.redis_user_game_bill")
	proto.RegisterType((*UserGameBill)(nil), "ddproto.user_game_bill")
}

func init() { proto.RegisterFile("common_server.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x6a, 0x1b, 0x3b,
	0x10, 0xc6, 0xcf, 0x7a, 0xed, 0xd8, 0x96, 0xe3, 0x24, 0x28, 0x21, 0x47, 0x84, 0x9c, 0x60, 0x7c,
	0x71, 0x30, 0x2d, 0x04, 0xda, 0x37, 0x48, 0x13, 0xda, 0x1a, 0xd2, 0x10, 0xd4, 0xb4, 0xbd, 0x34,
	0x5a, 0x4b, 0x18, 0x61, 0xaf, 0x14, 0x24, 0xad, 0x4d, 0xde, 0xb2, 0x8f, 0xd0, 0xfb, 0xde, 0xf4,
	0x11, 0x8a, 0x34, 0xfb, 0xd7, 0x6e, 0xd2, 0xbb, 0xfd, 0xbe, 0xd1, 0x8c, 0xa4, 0xdf, 0x8c, 0x16,
	0x1d, 0xcf, 0x75, 0x9a, 0x6a, 0x35, 0xb3, 0xc2, 0xac, 0x85, 0xb9, 0x7c, 0x34, 0xda, 0x69, 0xdc,
	0xe5, 0x3c, 0x7c, 0x9c, 0x15, 0xd1, 0xf9, 0x4a, 0x0a, 0xe5, 0x20, 0x3a, 0xfe, 0xd9, 0x42, 0x83,
	0x0f, 0x2c, 0x15, 0x9f, 0x85, 0xb5, 0x52, 0x2b, 0x7c, 0x8a, 0xf6, 0xbe, 0x58, 0x61, 0xa6, 0x9c,
	0x44, 0xa3, 0x68, 0x32, 0xa4, 0xb9, 0xf2, 0x3e, 0xd5, 0x3a, 0x9d, 0x72, 0xd2, 0x1a, 0x45, 0x93,
	0x0e, 0xcd, 0x95, 0xf7, 0x6f, 0x84, 0x5d, 0x4e, 0x39, 0x89, 0xc1, 0x07, 0x85, 0x2f, 0x10, 0x0a,
	0x65, 0x1d, 0x73, 0x99, 0x25, 0xed, 0x10, 0xab, 0x39, 0x3e, 0xcf, 0xab, 0x29, 0x27, 0x1d, 0xc8,
	0x03, 0x55, 0xe4, 0xdd, 0x65, 0x69, 0x22, 0x0c, 0xd9, 0xab, 0xf2, 0xc0, 0xc1, 0xaf, 0xd0, 0x91,
	0x57, 0xd7, 0x99, 0x75, 0x3a, 0xcd, 0xab, 0x77, 0xc3, 0xaa, 0x1d, 0x1f, 0x13, 0xd4, 0x95, 0xf6,
	0x9d, 0x11, 0x6c, 0x49, 0x7a, 0xa3, 0x68, 0xd2, 0xa3, 0x85, 0x84, 0xc8, 0xad, 0x60, 0x6b, 0x41,
	0xfa, 0x45, 0x24, 0x48, 0x7c, 0x86, 0x7a, 0x46, 0xeb, 0xf4, 0xe1, 0xe9, 0x51, 0x10, 0x14, 0xea,
	0x96, 0x1a, 0x8f, 0xd1, 0xbe, 0xff, 0xbe, 0x67, 0xd6, 0x6e, 0xb4, 0xe1, 0x64, 0x30, 0x8a, 0x26,
	0x7d, 0xda, 0xf0, 0xf0, 0x39, 0xea, 0x7b, 0x32, 0xb7, 0x62, 0x2d, 0x56, 0x64, 0x3f, 0x14, 0xa8,
	0x8c, 0xf1, 0xaf, 0x16, 0x3a, 0x29, 0x7a, 0x64, 0xd6, 0xb3, 0x05, 0x4b, 0xc5, 0x2c, 0xb3, 0xc2,
	0x78, 0x1c, 0x59, 0x03, 0x3b, 0x28, 0x7f, 0x1c, 0x25, 0xe7, 0xcb, 0x3b, 0x96, 0x8a, 0x00, 0xbe,
	0x4f, 0x4b, 0x8d, 0x31, 0x6a, 0xcf, 0xb5, 0x54, 0x01, 0x7c, 0x4c, 0xc3, 0xb7, 0xaf, 0x63, 0xa0,
	0x4d, 0x80, 0x3c, 0x57, 0xde, 0xe7, 0xd0, 0xa6, 0x1c, 0x37, 0x2f, 0xdb, 0xb4, 0xd8, 0xc1, 0x5d,
	0x39, 0x75, 0x84, 0xdd, 0x67, 0x11, 0xf6, 0x9a, 0x08, 0x4f, 0xd1, 0x9e, 0x85, 0xc6, 0xf4, 0x61,
	0x2f, 0x50, 0xfe, 0x2e, 0x1b, 0x26, 0xdd, 0x83, 0x4c, 0x01, 0x6d, 0x4c, 0x4b, 0xdd, 0xc0, 0x3e,
	0xf8, 0x0b, 0xf6, 0xfd, 0x3f, 0x60, 0x0f, 0xa7, 0xa1, 0x3a, 0xd1, 0x8e, 0x0c, 0x8b, 0xd3, 0x04,
	0x39, 0xfe, 0x11, 0xef, 0x22, 0xf7, 0x97, 0xaf, 0x21, 0x89, 0x5e, 0x40, 0xd2, 0xda, 0x41, 0x52,
	0x21, 0x8e, 0x1b, 0x88, 0x4f, 0x50, 0xc7, 0xe8, 0x4c, 0x15, 0xe4, 0x41, 0xf8, 0x8b, 0x3d, 0x16,
	0x07, 0xef, 0x40, 0x03, 0x0b, 0xed, 0x33, 0xf4, 0x46, 0xe5, 0xdc, 0x87, 0x14, 0x84, 0xaf, 0x9f,
	0x30, 0xb5, 0x14, 0x26, 0x10, 0x1f, 0xd2, 0x5c, 0xf9, 0xc9, 0x9a, 0x1b, 0xc1, 0x9c, 0x78, 0x2f,
	0x00, 0x79, 0x4c, 0x2b, 0xe3, 0x59, 0xe8, 0x17, 0x08, 0x25, 0x9a, 0x19, 0x6e, 0xaf, 0x75, 0xe6,
	0xf2, 0x89, 0xae, 0x39, 0xbe, 0x6a, 0xc2, 0xac, 0xf8, 0xca, 0x56, 0x19, 0x90, 0x8f, 0x69, 0x65,
	0x84, 0xa8, 0x58, 0x48, 0x15, 0x7a, 0x06, 0xdc, 0x2b, 0xc3, 0x43, 0x17, 0x8a, 0x87, 0xd8, 0x30,
	0xc4, 0x0a, 0x89, 0x27, 0xe8, 0x50, 0x4d, 0x95, 0x74, 0x57, 0x73, 0x27, 0x35, 0x64, 0x1f, 0x84,
	0xad, 0xb7, 0x6d, 0xfc, 0x3f, 0x3a, 0xf0, 0xa3, 0x7e, 0xad, 0x33, 0xe5, 0x6e, 0x65, 0x2a, 0x1d,
	0x39, 0x0c, 0x0b, 0xb7, 0x5c, 0xcf, 0x51, 0xda, 0x1b, 0x26, 0x97, 0x4c, 0x92, 0xa3, 0xd0, 0xe1,
	0x52, 0x8f, 0x3f, 0xa2, 0x63, 0x23, 0xb8, 0xb4, 0xa1, 0xaf, 0xb3, 0xe4, 0x69, 0xc6, 0x16, 0x42,
	0x39, 0xfc, 0x06, 0xb5, 0x39, 0x73, 0x8c, 0x44, 0xa3, 0x78, 0x32, 0x78, 0xfb, 0xdf, 0x65, 0xfe,
	0x1f, 0xbc, 0xcc, 0xa7, 0xa1, 0xb1, 0x98, 0x86, 0xa5, 0x63, 0x86, 0x4e, 0xa0, 0x92, 0xdf, 0x1d,
	0x66, 0x25, 0x91, 0xab, 0x15, 0x7e, 0xdd, 0x28, 0xf5, 0x6f, 0x59, 0xaa, 0xb9, 0x0c, 0x8a, 0x78,
	0x68, 0xd2, 0x7e, 0x93, 0xea, 0x93, 0xe6, 0xf0, 0x68, 0x7b, 0xb4, 0x32, 0xc6, 0xdf, 0x23, 0xb8,
	0x71, 0xad, 0xfa, 0x0b, 0x8f, 0xbf, 0x7c, 0x14, 0xad, 0xad, 0x47, 0x51, 0x4d, 0x6f, 0xdc, 0x98,
	0xde, 0xfa, 0xbc, 0xb5, 0xb7, 0xe6, 0xad, 0x39, 0xd9, 0x9d, 0x9d, 0xc9, 0x3e, 0x47, 0xfd, 0x8d,
	0x54, 0x57, 0xa9, 0xa7, 0x1e, 0x66, 0x32, 0xa6, 0x95, 0x51, 0xef, 0x76, 0x37, 0xc4, 0x0a, 0x79,
	0xff, 0xcf, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x3c, 0xb8, 0x22, 0x79, 0x06, 0x00, 0x00,
}
