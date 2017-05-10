// Code generated by protoc-gen-go.
// source: common_mj_play.proto
// DO NOT EDIT!

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of CardInfo from common_mj.proto

// Ignoring public import of RoomTypeInfo from common_mj.proto

// Ignoring public import of ComposeCard from common_mj.proto

// Ignoring public import of PlayerCard from common_mj.proto

// Ignoring public import of ErrorCode from common_mj.proto

// Ignoring public import of mj_enum_color from common_mj.proto

// Ignoring public import of mj_enum_gangType from common_mj.proto

// Ignoring public import of mj_enum_huType from common_mj.proto

// Ignoring public import of mj_enum_composeCardType from common_mj.proto

// Ignoring public import of mj_enum_paiType from common_mj.proto

// Ignoring public import of mj_enum_userGameStatus from common_mj.proto

// Ignoring public import of mj_enum_deskGameStatus from common_mj.proto

// Ignoring public import of MJRoomType from common_mj.proto

type GameReqBuxiazi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GameReqBuxiazi) Reset()                    { *m = GameReqBuxiazi{} }
func (m *GameReqBuxiazi) String() string            { return proto.CompactTextString(m) }
func (*GameReqBuxiazi) ProtoMessage()               {}
func (*GameReqBuxiazi) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *GameReqBuxiazi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GameReqBuxiazi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type GameAckBuxiazi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	BuCards          []*CardInfo  `protobuf:"bytes,2,rep,name=buCards" json:"buCards,omitempty"`
	XiaCouont        *int32       `protobuf:"varint,3,opt,name=xiaCouont" json:"xiaCouont,omitempty"`
	MoCards          []*CardInfo  `protobuf:"bytes,4,rep,name=moCards" json:"moCards,omitempty"`
	UserId           *uint32      `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	IsZhuoXiaBu      *bool        `protobuf:"varint,6,opt,name=IsZhuoXiaBu" json:"IsZhuoXiaBu,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GameAckBuxiazi) Reset()                    { *m = GameAckBuxiazi{} }
func (m *GameAckBuxiazi) String() string            { return proto.CompactTextString(m) }
func (*GameAckBuxiazi) ProtoMessage()               {}
func (*GameAckBuxiazi) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *GameAckBuxiazi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GameAckBuxiazi) GetBuCards() []*CardInfo {
	if m != nil {
		return m.BuCards
	}
	return nil
}

func (m *GameAckBuxiazi) GetXiaCouont() int32 {
	if m != nil && m.XiaCouont != nil {
		return *m.XiaCouont
	}
	return 0
}

func (m *GameAckBuxiazi) GetMoCards() []*CardInfo {
	if m != nil {
		return m.MoCards
	}
	return nil
}

func (m *GameAckBuxiazi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *GameAckBuxiazi) GetIsZhuoXiaBu() bool {
	if m != nil && m.IsZhuoXiaBu != nil {
		return *m.IsZhuoXiaBu
	}
	return false
}

func init() {
	proto.RegisterType((*GameReqBuxiazi)(nil), "ddproto.game_req_buxiazi")
	proto.RegisterType((*GameAckBuxiazi)(nil), "ddproto.game_ack_buxiazi")
}

var fileDescriptor5 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0xcf, 0xcd, 0x8a, 0x2f, 0xc8, 0x49, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0xa4, 0x84, 0xa1, 0xd2, 0xc9, 0x39, 0x99, 0xa9, 0x79, 0x25,
	0x10, 0x59, 0x29, 0x7e, 0xb8, 0x1e, 0x88, 0x80, 0x52, 0x04, 0x97, 0x40, 0x7a, 0x62, 0x6e, 0x6a,
	0x7c, 0x51, 0x6a, 0x61, 0x7c, 0x52, 0x69, 0x45, 0x66, 0x62, 0x55, 0xa6, 0x90, 0x0e, 0x17, 0x5b,
	0x46, 0x6a, 0x62, 0x4a, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e, 0xd4,
	0x4c, 0xbd, 0x00, 0x10, 0xe9, 0x01, 0x96, 0x0b, 0x82, 0xaa, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x2d,
	0x4e, 0x2d, 0xf2, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0xf2, 0x94, 0x3e, 0x31,
	0x42, 0x8d, 0x4e, 0x4c, 0xce, 0x26, 0xd3, 0x68, 0x6d, 0x2e, 0xf6, 0xa4, 0x52, 0xe7, 0xc4, 0xa2,
	0x94, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x41, 0xb8, 0x72, 0x90, 0xa8, 0x67, 0x5e,
	0x5a, 0x7e, 0x10, 0x4c, 0x85, 0x90, 0x0c, 0x17, 0x67, 0x45, 0x66, 0xa2, 0x73, 0x7e, 0x69, 0x7e,
	0x5e, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x42, 0x00, 0x64, 0x54, 0x6e, 0x3e, 0xc4,
	0x28, 0x16, 0x9c, 0x46, 0x41, 0x55, 0x20, 0x79, 0x89, 0x15, 0xd9, 0x4b, 0x42, 0x0a, 0x5c, 0xdc,
	0x9e, 0xc5, 0x51, 0x19, 0xa5, 0xf9, 0x11, 0x99, 0x89, 0x4e, 0xa5, 0x12, 0x6c, 0x0a, 0x8c, 0x1a,
	0x1c, 0x41, 0xc8, 0x42, 0x01, 0x0c, 0x01, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x57,
	0xab, 0x77, 0x98, 0x01, 0x00, 0x00,
}
