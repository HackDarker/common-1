// Code generated by protoc-gen-go.
// source: common_server_msg.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// 游戏结束后的统计数据消息
type CommonSrvMsgGameCount struct {
	UserId           *uint32               `protobuf:"varint,1,opt,name=user_id" json:"user_id,omitempty"`
	GameId           *CommonEnumGame       `protobuf:"varint,2,opt,name=game_id,enum=ddproto.CommonEnumGame" json:"game_id,omitempty"`
	GameNumber       *int32                `protobuf:"varint,3,opt,name=game_number" json:"game_number,omitempty"`
	RoomType         *COMMON_ENUM_ROOMTYPE `protobuf:"varint,4,opt,name=room_type,enum=ddproto.COMMON_ENUM_ROOMTYPE" json:"room_type,omitempty"`
	RoomLevel        *int32                `protobuf:"varint,5,opt,name=room_level" json:"room_level,omitempty"`
	Bill             *float64              `protobuf:"fixed64,6,opt,name=bill" json:"bill,omitempty"`
	IsWine           *bool                 `protobuf:"varint,7,opt,name=is_wine" json:"is_wine,omitempty"`
	Data             *string               `protobuf:"bytes,8,opt,name=data" json:"data,omitempty"`
	StartTime        *int64                `protobuf:"varint,9,opt,name=start_time" json:"start_time,omitempty"`
	EndTime          *int64                `protobuf:"varint,10,opt,name=end_time" json:"end_time,omitempty"`
	CoinFee          *int32                `protobuf:"varint,11,opt,name=coin_fee" json:"coin_fee,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *CommonSrvMsgGameCount) Reset()                    { *m = CommonSrvMsgGameCount{} }
func (m *CommonSrvMsgGameCount) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvMsgGameCount) ProtoMessage()               {}
func (*CommonSrvMsgGameCount) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *CommonSrvMsgGameCount) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetGameId() CommonEnumGame {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return CommonEnumGame_GID_SRC
}

func (m *CommonSrvMsgGameCount) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetRoomType() COMMON_ENUM_ROOMTYPE {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return COMMON_ENUM_ROOMTYPE_DESK_FRIEND
}

func (m *CommonSrvMsgGameCount) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetBill() float64 {
	if m != nil && m.Bill != nil {
		return *m.Bill
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetIsWine() bool {
	if m != nil && m.IsWine != nil {
		return *m.IsWine
	}
	return false
}

func (m *CommonSrvMsgGameCount) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *CommonSrvMsgGameCount) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetCoinFee() int32 {
	if m != nil && m.CoinFee != nil {
		return *m.CoinFee
	}
	return 0
}

func init() {
	proto.RegisterType((*CommonSrvMsgGameCount)(nil), "ddproto.common_srv_msg_game_count")
}

var fileDescriptor8 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8e, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0x49, 0x1f, 0x24, 0x99, 0xf2, 0x34, 0x0b, 0x5c, 0x24, 0x24, 0xc4, 0x0a, 0xb1, 0x88,
	0x10, 0xbf, 0x80, 0xba, 0x4c, 0x53, 0x21, 0x58, 0xb0, 0x1a, 0xa5, 0xf1, 0x50, 0x59, 0x8a, 0xed,
	0xca, 0x71, 0x82, 0xf8, 0x13, 0x3e, 0x17, 0xdb, 0x6d, 0x11, 0xbb, 0xcc, 0xb9, 0xf7, 0x9e, 0x18,
	0xae, 0x1b, 0xa3, 0x94, 0xd1, 0xd8, 0x91, 0x1d, 0xc8, 0xa2, 0xea, 0x36, 0xc5, 0xd6, 0x1a, 0x67,
	0x58, 0x2a, 0x44, 0xfc, 0xb8, 0xb9, 0xdc, 0x37, 0x48, 0xf7, 0x6a, 0x97, 0xdd, 0xff, 0x8c, 0x60,
	0x7e, 0xd8, 0xd9, 0x21, 0x8c, 0x70, 0x53, 0x2b, 0xc2, 0xc6, 0xf4, 0xda, 0xb1, 0x73, 0x48, 0x7b,
	0xaf, 0x43, 0x29, 0x78, 0x72, 0x97, 0x3c, 0x9c, 0xb2, 0x47, 0x48, 0x63, 0xec, 0xc1, 0xc8, 0x83,
	0xb3, 0xe7, 0x79, 0xb1, 0x97, 0x17, 0xff, 0xdc, 0x51, 0xc1, 0xae, 0x60, 0x16, 0xbb, 0x1e, 0xac,
	0xc9, 0xf2, 0xb1, 0xef, 0x4f, 0xd9, 0x13, 0xe4, 0xd6, 0x18, 0x85, 0xee, 0x7b, 0x4b, 0x7c, 0x12,
	0x15, 0xb7, 0x7f, 0x8a, 0x97, 0xaa, 0x2c, 0xab, 0x25, 0x2e, 0x96, 0xef, 0x25, 0xbe, 0x56, 0x55,
	0xf9, 0xf6, 0xb1, 0x5a, 0x30, 0x06, 0x10, 0x17, 0x2d, 0x0d, 0xd4, 0xf2, 0x69, 0xb4, 0x9c, 0xc0,
	0x64, 0x2d, 0xdb, 0x96, 0x1f, 0xfb, 0x2b, 0x09, 0xaf, 0x94, 0x1d, 0x7e, 0x49, 0x4d, 0x3c, 0xf5,
	0x20, 0x0b, 0xb1, 0xa8, 0x5d, 0xcd, 0x33, 0x7f, 0xe5, 0x41, 0xd0, 0xb9, 0xda, 0x3a, 0x74, 0x52,
	0x11, 0xcf, 0x3d, 0x1b, 0xb3, 0x0b, 0xc8, 0x48, 0x8b, 0x1d, 0x81, 0x03, 0x69, 0x8c, 0xd4, 0xf8,
	0x49, 0xc4, 0x67, 0xe1, 0x27, 0xab, 0xa3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x42, 0x2b,
	0x60, 0x52, 0x01, 0x00, 0x00,
}
