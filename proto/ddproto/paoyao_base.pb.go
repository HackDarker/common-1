// Code generated by protoc-gen-go. DO NOT EDIT.
// source: paoyao_base.proto

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

type PaoyaoEnumProtoid int32

const (
	// //////////////////////////////////
	PaoyaoEnumProtoid_PAOYAO_PID_HEARTBEAT          PaoyaoEnumProtoid = 0
	PaoyaoEnumProtoid_PAOYAO_PID_CREATE_DESK_REQ    PaoyaoEnumProtoid = 1
	PaoyaoEnumProtoid_PAOYAO_PID_ENTER_DESK_REQ     PaoyaoEnumProtoid = 2
	PaoyaoEnumProtoid_PAOYAO_PID_ENTER_DESK_ACK     PaoyaoEnumProtoid = 3
	PaoyaoEnumProtoid_PAOYAO_PID_ENTER_DESK_BC      PaoyaoEnumProtoid = 4
	PaoyaoEnumProtoid_PAOYAO_PID_READY_REQ          PaoyaoEnumProtoid = 5
	PaoyaoEnumProtoid_PAOYAO_PID_READY_BC           PaoyaoEnumProtoid = 6
	PaoyaoEnumProtoid_PAOYAO_PID_FAPAI_OT           PaoyaoEnumProtoid = 7
	PaoyaoEnumProtoid_PAOYAO_PID_DOUBLE_REQ         PaoyaoEnumProtoid = 8
	PaoyaoEnumProtoid_PAOYAO_PID_DOUBLE_BC          PaoyaoEnumProtoid = 9
	PaoyaoEnumProtoid_PAOYAP_PID_CHUPAI_OT          PaoyaoEnumProtoid = 10
	PaoyaoEnumProtoid_PAOYAO_PID_CHUPAI_REQ         PaoyaoEnumProtoid = 11
	PaoyaoEnumProtoid_PAOYAO_PID_CHUPAI_BC          PaoyaoEnumProtoid = 12
	PaoyaoEnumProtoid_PAOYAO_PID_PASS_REQ           PaoyaoEnumProtoid = 13
	PaoyaoEnumProtoid_PAOYAO_PID_PASS_BC            PaoyaoEnumProtoid = 14
	PaoyaoEnumProtoid_PAOYAO_PID_GAME_END_ONE_BC    PaoyaoEnumProtoid = 15
	PaoyaoEnumProtoid_PAOYAO_PID_GAME_END_ALL_BC    PaoyaoEnumProtoid = 16
	PaoyaoEnumProtoid_PAOYAO_PID_APPLY_DISSOLVE_REQ PaoyaoEnumProtoid = 17
	PaoyaoEnumProtoid_PAOYAO_PID_APPLY_DISSOLVE_ACK PaoyaoEnumProtoid = 18
	PaoyaoEnumProtoid_PAOYAO_PID_DISSOLVE_BACK_REQ  PaoyaoEnumProtoid = 19
	PaoyaoEnumProtoid_PAOYAO_PID_DISSOLVE_BACK_ACK  PaoyaoEnumProtoid = 20
	PaoyaoEnumProtoid_PAOYAO_PID_OWNER_DISSOLVE_REQ PaoyaoEnumProtoid = 21
	PaoyaoEnumProtoid_PAOYAO_PID_OWNER_DISSOLVE_ACK PaoyaoEnumProtoid = 22
	PaoyaoEnumProtoid_PAOYAO_PID_SEND_MESSAGE_REQ   PaoyaoEnumProtoid = 23
	PaoyaoEnumProtoid_PAOYAO_PID_MESSAGE_BC         PaoyaoEnumProtoid = 24
	PaoyaoEnumProtoid_PAOYAO_PID_LEAVE_DESK_REQ     PaoyaoEnumProtoid = 25
	PaoyaoEnumProtoid_PAOYAO_PID_LEAVE_DESK_ACK     PaoyaoEnumProtoid = 26
	PaoyaoEnumProtoid_PAOYAO_OFFLINE_BC             PaoyaoEnumProtoid = 27
	PaoyaoEnumProtoid_PAOYAO_COIN_ROOM_LIST_REQ     PaoyaoEnumProtoid = 28
	PaoyaoEnumProtoid_PAOYAO_COIN_ROOM_LIST_ACK     PaoyaoEnumProtoid = 29
	PaoyaoEnumProtoid_PAOYAO_PID_SITE_DOWN_REQ      PaoyaoEnumProtoid = 30
	PaoyaoEnumProtoid_PAOYAO_PID_SITE_UP_REQ        PaoyaoEnumProtoid = 31
	PaoyaoEnumProtoid_PAOYAO_PID_PRE_END_OT         PaoyaoEnumProtoid = 32
	PaoyaoEnumProtoid_PAOYAO_PID_PRE_END_REQ        PaoyaoEnumProtoid = 33
	PaoyaoEnumProtoid_PAOYAO_PID_PRE_END_ACK        PaoyaoEnumProtoid = 34
)

var PaoyaoEnumProtoid_name = map[int32]string{
	0:  "PAOYAO_PID_HEARTBEAT",
	1:  "PAOYAO_PID_CREATE_DESK_REQ",
	2:  "PAOYAO_PID_ENTER_DESK_REQ",
	3:  "PAOYAO_PID_ENTER_DESK_ACK",
	4:  "PAOYAO_PID_ENTER_DESK_BC",
	5:  "PAOYAO_PID_READY_REQ",
	6:  "PAOYAO_PID_READY_BC",
	7:  "PAOYAO_PID_FAPAI_OT",
	8:  "PAOYAO_PID_DOUBLE_REQ",
	9:  "PAOYAO_PID_DOUBLE_BC",
	10: "PAOYAP_PID_CHUPAI_OT",
	11: "PAOYAO_PID_CHUPAI_REQ",
	12: "PAOYAO_PID_CHUPAI_BC",
	13: "PAOYAO_PID_PASS_REQ",
	14: "PAOYAO_PID_PASS_BC",
	15: "PAOYAO_PID_GAME_END_ONE_BC",
	16: "PAOYAO_PID_GAME_END_ALL_BC",
	17: "PAOYAO_PID_APPLY_DISSOLVE_REQ",
	18: "PAOYAO_PID_APPLY_DISSOLVE_ACK",
	19: "PAOYAO_PID_DISSOLVE_BACK_REQ",
	20: "PAOYAO_PID_DISSOLVE_BACK_ACK",
	21: "PAOYAO_PID_OWNER_DISSOLVE_REQ",
	22: "PAOYAO_PID_OWNER_DISSOLVE_ACK",
	23: "PAOYAO_PID_SEND_MESSAGE_REQ",
	24: "PAOYAO_PID_MESSAGE_BC",
	25: "PAOYAO_PID_LEAVE_DESK_REQ",
	26: "PAOYAO_PID_LEAVE_DESK_ACK",
	27: "PAOYAO_OFFLINE_BC",
	28: "PAOYAO_COIN_ROOM_LIST_REQ",
	29: "PAOYAO_COIN_ROOM_LIST_ACK",
	30: "PAOYAO_PID_SITE_DOWN_REQ",
	31: "PAOYAO_PID_SITE_UP_REQ",
	32: "PAOYAO_PID_PRE_END_OT",
	33: "PAOYAO_PID_PRE_END_REQ",
	34: "PAOYAO_PID_PRE_END_ACK",
}
var PaoyaoEnumProtoid_value = map[string]int32{
	"PAOYAO_PID_HEARTBEAT":          0,
	"PAOYAO_PID_CREATE_DESK_REQ":    1,
	"PAOYAO_PID_ENTER_DESK_REQ":     2,
	"PAOYAO_PID_ENTER_DESK_ACK":     3,
	"PAOYAO_PID_ENTER_DESK_BC":      4,
	"PAOYAO_PID_READY_REQ":          5,
	"PAOYAO_PID_READY_BC":           6,
	"PAOYAO_PID_FAPAI_OT":           7,
	"PAOYAO_PID_DOUBLE_REQ":         8,
	"PAOYAO_PID_DOUBLE_BC":          9,
	"PAOYAP_PID_CHUPAI_OT":          10,
	"PAOYAO_PID_CHUPAI_REQ":         11,
	"PAOYAO_PID_CHUPAI_BC":          12,
	"PAOYAO_PID_PASS_REQ":           13,
	"PAOYAO_PID_PASS_BC":            14,
	"PAOYAO_PID_GAME_END_ONE_BC":    15,
	"PAOYAO_PID_GAME_END_ALL_BC":    16,
	"PAOYAO_PID_APPLY_DISSOLVE_REQ": 17,
	"PAOYAO_PID_APPLY_DISSOLVE_ACK": 18,
	"PAOYAO_PID_DISSOLVE_BACK_REQ":  19,
	"PAOYAO_PID_DISSOLVE_BACK_ACK":  20,
	"PAOYAO_PID_OWNER_DISSOLVE_REQ": 21,
	"PAOYAO_PID_OWNER_DISSOLVE_ACK": 22,
	"PAOYAO_PID_SEND_MESSAGE_REQ":   23,
	"PAOYAO_PID_MESSAGE_BC":         24,
	"PAOYAO_PID_LEAVE_DESK_REQ":     25,
	"PAOYAO_PID_LEAVE_DESK_ACK":     26,
	"PAOYAO_OFFLINE_BC":             27,
	"PAOYAO_COIN_ROOM_LIST_REQ":     28,
	"PAOYAO_COIN_ROOM_LIST_ACK":     29,
	"PAOYAO_PID_SITE_DOWN_REQ":      30,
	"PAOYAO_PID_SITE_UP_REQ":        31,
	"PAOYAO_PID_PRE_END_OT":         32,
	"PAOYAO_PID_PRE_END_REQ":        33,
	"PAOYAO_PID_PRE_END_ACK":        34,
}

func (x PaoyaoEnumProtoid) Enum() *PaoyaoEnumProtoid {
	p := new(PaoyaoEnumProtoid)
	*p = x
	return p
}
func (x PaoyaoEnumProtoid) String() string {
	return proto.EnumName(PaoyaoEnumProtoid_name, int32(x))
}
func (x *PaoyaoEnumProtoid) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PaoyaoEnumProtoid_value, data, "PaoyaoEnumProtoid")
	if err != nil {
		return err
	}
	*x = PaoyaoEnumProtoid(value)
	return nil
}
func (PaoyaoEnumProtoid) EnumDescriptor() ([]byte, []int) { return fileDescriptor36, []int{0} }

// =================================公共================================
// 牌型
type PaoyaoEnumPokerType int32

const (
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_OTHER        PaoyaoEnumPokerType = 0
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_DAN_PAI      PaoyaoEnumPokerType = 1
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_DUI_PAI      PaoyaoEnumPokerType = 2
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_DAN_SHUN     PaoyaoEnumPokerType = 3
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_SHUANG_SHUN  PaoyaoEnumPokerType = 4
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_3_LU         PaoyaoEnumPokerType = 5
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_2_XIAO_WANG  PaoyaoEnumPokerType = 6
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_4_LU         PaoyaoEnumPokerType = 7
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_DA_XIAO_WANG PaoyaoEnumPokerType = 8
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_XIAO_YAO     PaoyaoEnumPokerType = 9
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_5_LU         PaoyaoEnumPokerType = 10
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_2_DA_WANG    PaoyaoEnumPokerType = 11
	PaoyaoEnumPokerType_PAOYAO_POKER_2XIAO_1DA_WANG    PaoyaoEnumPokerType = 12
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_6_LU         PaoyaoEnumPokerType = 13
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_ZHONG_YAO    PaoyaoEnumPokerType = 14
	PaoyaoEnumPokerType_PAOYAO_POKER_2DA_1XIAO_WANG    PaoyaoEnumPokerType = 15
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_7_LU         PaoyaoEnumPokerType = 16
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_8_LU         PaoyaoEnumPokerType = 17
	PaoyaoEnumPokerType_PAOYAO_POKER_4_WANG            PaoyaoEnumPokerType = 18
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_LAO_YAO      PaoyaoEnumPokerType = 19
	PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_LAOLAO_YAO   PaoyaoEnumPokerType = 20
)

var PaoyaoEnumPokerType_name = map[int32]string{
	0:  "PAOYAO_POKER_TYPE_OTHER",
	1:  "PAOYAO_POKER_TYPE_DAN_PAI",
	2:  "PAOYAO_POKER_TYPE_DUI_PAI",
	3:  "PAOYAO_POKER_TYPE_DAN_SHUN",
	4:  "PAOYAO_POKER_TYPE_SHUANG_SHUN",
	5:  "PAOYAO_POKER_TYPE_3_LU",
	6:  "PAOYAO_POKER_TYPE_2_XIAO_WANG",
	7:  "PAOYAO_POKER_TYPE_4_LU",
	8:  "PAOYAO_POKER_TYPE_DA_XIAO_WANG",
	9:  "PAOYAO_POKER_TYPE_XIAO_YAO",
	10: "PAOYAO_POKER_TYPE_5_LU",
	11: "PAOYAO_POKER_TYPE_2_DA_WANG",
	12: "PAOYAO_POKER_2XIAO_1DA_WANG",
	13: "PAOYAO_POKER_TYPE_6_LU",
	14: "PAOYAO_POKER_TYPE_ZHONG_YAO",
	15: "PAOYAO_POKER_2DA_1XIAO_WANG",
	16: "PAOYAO_POKER_TYPE_7_LU",
	17: "PAOYAO_POKER_TYPE_8_LU",
	18: "PAOYAO_POKER_4_WANG",
	19: "PAOYAO_POKER_TYPE_LAO_YAO",
	20: "PAOYAO_POKER_TYPE_LAOLAO_YAO",
}
var PaoyaoEnumPokerType_value = map[string]int32{
	"PAOYAO_POKER_TYPE_OTHER":        0,
	"PAOYAO_POKER_TYPE_DAN_PAI":      1,
	"PAOYAO_POKER_TYPE_DUI_PAI":      2,
	"PAOYAO_POKER_TYPE_DAN_SHUN":     3,
	"PAOYAO_POKER_TYPE_SHUANG_SHUN":  4,
	"PAOYAO_POKER_TYPE_3_LU":         5,
	"PAOYAO_POKER_TYPE_2_XIAO_WANG":  6,
	"PAOYAO_POKER_TYPE_4_LU":         7,
	"PAOYAO_POKER_TYPE_DA_XIAO_WANG": 8,
	"PAOYAO_POKER_TYPE_XIAO_YAO":     9,
	"PAOYAO_POKER_TYPE_5_LU":         10,
	"PAOYAO_POKER_TYPE_2_DA_WANG":    11,
	"PAOYAO_POKER_2XIAO_1DA_WANG":    12,
	"PAOYAO_POKER_TYPE_6_LU":         13,
	"PAOYAO_POKER_TYPE_ZHONG_YAO":    14,
	"PAOYAO_POKER_2DA_1XIAO_WANG":    15,
	"PAOYAO_POKER_TYPE_7_LU":         16,
	"PAOYAO_POKER_TYPE_8_LU":         17,
	"PAOYAO_POKER_4_WANG":            18,
	"PAOYAO_POKER_TYPE_LAO_YAO":      19,
	"PAOYAO_POKER_TYPE_LAOLAO_YAO":   20,
}

func (x PaoyaoEnumPokerType) Enum() *PaoyaoEnumPokerType {
	p := new(PaoyaoEnumPokerType)
	*p = x
	return p
}
func (x PaoyaoEnumPokerType) String() string {
	return proto.EnumName(PaoyaoEnumPokerType_name, int32(x))
}
func (x *PaoyaoEnumPokerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PaoyaoEnumPokerType_value, data, "PaoyaoEnumPokerType")
	if err != nil {
		return err
	}
	*x = PaoyaoEnumPokerType(value)
	return nil
}
func (PaoyaoEnumPokerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor36, []int{1} }

// 雪的类型
type PaoyaoSnowType int32

const (
	PaoyaoSnowType_PAOYAO_SNOW_WU_XUE   PaoyaoSnowType = 1
	PaoyaoSnowType_PAOYAO_SNOW_XIAO_XUE PaoyaoSnowType = 2
	PaoyaoSnowType_PAOYAO_SNOW_DA_XUE   PaoyaoSnowType = 3
)

var PaoyaoSnowType_name = map[int32]string{
	1: "PAOYAO_SNOW_WU_XUE",
	2: "PAOYAO_SNOW_XIAO_XUE",
	3: "PAOYAO_SNOW_DA_XUE",
}
var PaoyaoSnowType_value = map[string]int32{
	"PAOYAO_SNOW_WU_XUE":   1,
	"PAOYAO_SNOW_XIAO_XUE": 2,
	"PAOYAO_SNOW_DA_XUE":   3,
}

func (x PaoyaoSnowType) Enum() *PaoyaoSnowType {
	p := new(PaoyaoSnowType)
	*p = x
	return p
}
func (x PaoyaoSnowType) String() string {
	return proto.EnumName(PaoyaoSnowType_name, int32(x))
}
func (x *PaoyaoSnowType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PaoyaoSnowType_value, data, "PaoyaoSnowType")
	if err != nil {
		return err
	}
	*x = PaoyaoSnowType(value)
	return nil
}
func (PaoyaoSnowType) EnumDescriptor() ([]byte, []int) { return fileDescriptor36, []int{2} }

// 房间状态
type PaoyaoEnumDeskStatus int32

const (
	PaoyaoEnumDeskStatus_PAOYAO_DESK_STATUS_WAIT_READY PaoyaoEnumDeskStatus = 1
	PaoyaoEnumDeskStatus_PAOYAO_DESK_STATUS_ON_GAMMING PaoyaoEnumDeskStatus = 2
)

var PaoyaoEnumDeskStatus_name = map[int32]string{
	1: "PAOYAO_DESK_STATUS_WAIT_READY",
	2: "PAOYAO_DESK_STATUS_ON_GAMMING",
}
var PaoyaoEnumDeskStatus_value = map[string]int32{
	"PAOYAO_DESK_STATUS_WAIT_READY": 1,
	"PAOYAO_DESK_STATUS_ON_GAMMING": 2,
}

func (x PaoyaoEnumDeskStatus) Enum() *PaoyaoEnumDeskStatus {
	p := new(PaoyaoEnumDeskStatus)
	*p = x
	return p
}
func (x PaoyaoEnumDeskStatus) String() string {
	return proto.EnumName(PaoyaoEnumDeskStatus_name, int32(x))
}
func (x *PaoyaoEnumDeskStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PaoyaoEnumDeskStatus_value, data, "PaoyaoEnumDeskStatus")
	if err != nil {
		return err
	}
	*x = PaoyaoEnumDeskStatus(value)
	return nil
}
func (PaoyaoEnumDeskStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor36, []int{3} }

// 打出去的牌
type PaoyaoClientPoker struct {
	Pais             []int32              `protobuf:"varint,2,rep,name=pais" json:"pais,omitempty"`
	PokerType        *PaoyaoEnumPokerType `protobuf:"varint,3,opt,name=pokerType,enum=ddproto.PaoyaoEnumPokerType" json:"pokerType,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PaoyaoClientPoker) Reset()                    { *m = PaoyaoClientPoker{} }
func (m *PaoyaoClientPoker) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoClientPoker) ProtoMessage()               {}
func (*PaoyaoClientPoker) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{0} }

func (m *PaoyaoClientPoker) GetPais() []int32 {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *PaoyaoClientPoker) GetPokerType() PaoyaoEnumPokerType {
	if m != nil && m.PokerType != nil {
		return *m.PokerType
	}
	return PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_OTHER
}

// 对局账单信息
type PaoyaoUserBill struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Score            *int32  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	CountWin         *int32  `protobuf:"varint,3,opt,name=count_win,json=countWin" json:"count_win,omitempty"`
	CountLost        *int32  `protobuf:"varint,4,opt,name=count_lost,json=countLost" json:"count_lost,omitempty"`
	KangQiNum        *int32  `protobuf:"varint,5,opt,name=kangQiNum" json:"kangQiNum,omitempty"`
	DaXueNum         *int32  `protobuf:"varint,6,opt,name=daXueNum" json:"daXueNum,omitempty"`
	XiaoXueNum       *int32  `protobuf:"varint,7,opt,name=xiaoXueNum" json:"xiaoXueNum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PaoyaoUserBill) Reset()                    { *m = PaoyaoUserBill{} }
func (m *PaoyaoUserBill) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoUserBill) ProtoMessage()               {}
func (*PaoyaoUserBill) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{1} }

func (m *PaoyaoUserBill) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PaoyaoUserBill) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *PaoyaoUserBill) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *PaoyaoUserBill) GetCountLost() int32 {
	if m != nil && m.CountLost != nil {
		return *m.CountLost
	}
	return 0
}

func (m *PaoyaoUserBill) GetKangQiNum() int32 {
	if m != nil && m.KangQiNum != nil {
		return *m.KangQiNum
	}
	return 0
}

func (m *PaoyaoUserBill) GetDaXueNum() int32 {
	if m != nil && m.DaXueNum != nil {
		return *m.DaXueNum
	}
	return 0
}

func (m *PaoyaoUserBill) GetXiaoXueNum() int32 {
	if m != nil && m.XiaoXueNum != nil {
		return *m.XiaoXueNum
	}
	return 0
}

// desk 配置选项
type PaoyaoDeskOption struct {
	// 金币场
	GammerNum        *int32 `protobuf:"varint,1,opt,name=gammerNum" json:"gammerNum,omitempty"`
	BoardsCout       *int32 `protobuf:"varint,2,opt,name=boardsCout" json:"boardsCout,omitempty"`
	AllyType         *int32 `protobuf:"varint,3,opt,name=allyType" json:"allyType,omitempty"`
	DisplayYao       *int32 `protobuf:"varint,4,opt,name=displayYao" json:"displayYao,omitempty"`
	AllowJiaBei      *bool  `protobuf:"varint,5,opt,name=allowJiaBei" json:"allowJiaBei,omitempty"`
	HasAnimation     *bool  `protobuf:"varint,8,opt,name=hasAnimation" json:"hasAnimation,omitempty"`
	BaseChip         *int32 `protobuf:"varint,9,opt,name=baseChip" json:"baseChip,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PaoyaoDeskOption) Reset()                    { *m = PaoyaoDeskOption{} }
func (m *PaoyaoDeskOption) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoDeskOption) ProtoMessage()               {}
func (*PaoyaoDeskOption) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{2} }

func (m *PaoyaoDeskOption) GetGammerNum() int32 {
	if m != nil && m.GammerNum != nil {
		return *m.GammerNum
	}
	return 0
}

func (m *PaoyaoDeskOption) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *PaoyaoDeskOption) GetAllyType() int32 {
	if m != nil && m.AllyType != nil {
		return *m.AllyType
	}
	return 0
}

func (m *PaoyaoDeskOption) GetDisplayYao() int32 {
	if m != nil && m.DisplayYao != nil {
		return *m.DisplayYao
	}
	return 0
}

func (m *PaoyaoDeskOption) GetAllowJiaBei() bool {
	if m != nil && m.AllowJiaBei != nil {
		return *m.AllowJiaBei
	}
	return false
}

func (m *PaoyaoDeskOption) GetHasAnimation() bool {
	if m != nil && m.HasAnimation != nil {
		return *m.HasAnimation
	}
	return false
}

func (m *PaoyaoDeskOption) GetBaseChip() int32 {
	if m != nil && m.BaseChip != nil {
		return *m.BaseChip
	}
	return 0
}

// room 的信息
type PaoyaoSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	BaseChip         *int32  `protobuf:"varint,5,opt,name=baseChip" json:"baseChip,omitempty"`
	EnterCoin        *int32  `protobuf:"varint,6,opt,name=enterCoin" json:"enterCoin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PaoyaoSrvRoom) Reset()                    { *m = PaoyaoSrvRoom{} }
func (m *PaoyaoSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoSrvRoom) ProtoMessage()               {}
func (*PaoyaoSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{3} }

func (m *PaoyaoSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *PaoyaoSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *PaoyaoSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *PaoyaoSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func (m *PaoyaoSrvRoom) GetBaseChip() int32 {
	if m != nil && m.BaseChip != nil {
		return *m.BaseChip
	}
	return 0
}

func (m *PaoyaoSrvRoom) GetEnterCoin() int32 {
	if m != nil && m.EnterCoin != nil {
		return *m.EnterCoin
	}
	return 0
}

func init() {
	proto.RegisterType((*PaoyaoClientPoker)(nil), "ddproto.paoyao_client_poker")
	proto.RegisterType((*PaoyaoUserBill)(nil), "ddproto.paoyao_user_bill")
	proto.RegisterType((*PaoyaoDeskOption)(nil), "ddproto.paoyao_desk_option")
	proto.RegisterType((*PaoyaoSrvRoom)(nil), "ddproto.paoyao_srv_room")
	proto.RegisterEnum("ddproto.PaoyaoEnumProtoid", PaoyaoEnumProtoid_name, PaoyaoEnumProtoid_value)
	proto.RegisterEnum("ddproto.PaoyaoEnumPokerType", PaoyaoEnumPokerType_name, PaoyaoEnumPokerType_value)
	proto.RegisterEnum("ddproto.PaoyaoSnowType", PaoyaoSnowType_name, PaoyaoSnowType_value)
	proto.RegisterEnum("ddproto.PaoyaoEnumDeskStatus", PaoyaoEnumDeskStatus_name, PaoyaoEnumDeskStatus_value)
}

func init() { proto.RegisterFile("paoyao_base.proto", fileDescriptor36) }

var fileDescriptor36 = []byte{
	// 1082 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0xdb, 0x72, 0xe2, 0x38,
	0x10, 0x86, 0x03, 0x09, 0x09, 0x28, 0x27, 0x45, 0x39, 0x39, 0xc7, 0x61, 0xb8, 0x4a, 0xe5, 0x22,
	0x55, 0x93, 0x9d, 0x3d, 0x5c, 0xec, 0x8d, 0x30, 0x0e, 0x78, 0xe3, 0xd8, 0x1e, 0x1f, 0x86, 0xb0,
	0xb5, 0x55, 0x2a, 0x27, 0xb8, 0x32, 0xae, 0x18, 0x8b, 0xc2, 0x66, 0xb2, 0x79, 0x93, 0xbd, 0xde,
	0xd7, 0xd8, 0xf7, 0xd8, 0xe7, 0xd8, 0x47, 0xd8, 0x92, 0x6c, 0x40, 0x38, 0x26, 0x57, 0x48, 0xfd,
	0x75, 0xff, 0xea, 0x56, 0x37, 0x02, 0xb0, 0x33, 0xf4, 0xe8, 0xab, 0x47, 0xc9, 0x83, 0x17, 0xfb,
	0x57, 0xc3, 0x11, 0x4d, 0x28, 0x5a, 0xeb, 0xf7, 0xf9, 0xe2, 0x78, 0xf7, 0x91, 0x0e, 0x06, 0x34,
	0x22, 0x8f, 0x61, 0xe0, 0x47, 0x49, 0x4a, 0x1b, 0x4f, 0x60, 0x37, 0x0b, 0x49, 0xcd, 0x64, 0x48,
	0x9f, 0xfd, 0x11, 0x42, 0x60, 0x65, 0xe8, 0x05, 0xb1, 0x54, 0xae, 0x2f, 0x5f, 0x54, 0x2c, 0xbe,
	0x46, 0xbf, 0x82, 0x1a, 0x87, 0xce, 0xeb, 0xd0, 0x97, 0x96, 0xeb, 0xa5, 0x8b, 0xad, 0xeb, 0xf3,
	0xab, 0x4c, 0xfc, 0x2a, 0x13, 0xf1, 0xa3, 0xf1, 0x80, 0x4c, 0xbd, 0xac, 0x59, 0x40, 0xe3, 0xdf,
	0x12, 0x80, 0x99, 0xd3, 0x38, 0xf6, 0x47, 0xe4, 0x21, 0x08, 0x43, 0x74, 0x00, 0x56, 0xd9, 0x46,
	0xed, 0x4b, 0xa5, 0x7a, 0xe9, 0x62, 0xd3, 0xca, 0x76, 0x68, 0x0f, 0x54, 0xe2, 0x47, 0x3a, 0xf2,
	0xa5, 0x72, 0xbd, 0x74, 0x51, 0xb1, 0xd2, 0x0d, 0x3a, 0x01, 0xb5, 0x47, 0x3a, 0x8e, 0x12, 0xf2,
	0x12, 0x44, 0x3c, 0x81, 0x8a, 0x55, 0xe5, 0x86, 0x6e, 0x10, 0xa1, 0x33, 0x00, 0x52, 0x18, 0xd2,
	0x38, 0x91, 0x56, 0x38, 0x4d, 0xdd, 0x35, 0x1a, 0x27, 0xe8, 0x14, 0xd4, 0x9e, 0xbd, 0xe8, 0xe9,
	0x4b, 0xa0, 0x8f, 0x07, 0x52, 0x25, 0xa5, 0x53, 0x03, 0x3a, 0x06, 0xd5, 0xbe, 0x77, 0x3f, 0xf6,
	0x19, 0x5c, 0x4d, 0x85, 0x27, 0x7b, 0x74, 0x0e, 0xc0, 0x9f, 0x81, 0x47, 0x33, 0xba, 0xc6, 0xa9,
	0x60, 0x69, 0xfc, 0x57, 0x02, 0x28, 0x2b, 0xac, 0xef, 0xc7, 0xcf, 0x84, 0x0e, 0x93, 0x80, 0x46,
	0xec, 0xc0, 0x27, 0x6f, 0x30, 0xf0, 0x47, 0x2c, 0xaa, 0x94, 0x1e, 0x38, 0x35, 0x30, 0xd1, 0x07,
	0xea, 0x8d, 0xfa, 0xb1, 0x4c, 0xc7, 0x49, 0x56, 0xa5, 0x60, 0x61, 0x09, 0x79, 0x61, 0xf8, 0x3a,
	0xbd, 0xea, 0x8a, 0x35, 0xdd, 0xb3, 0xd8, 0x7e, 0x10, 0x0f, 0x43, 0xef, 0xb5, 0xe7, 0xd1, 0xac,
	0x52, 0xc1, 0x82, 0xea, 0x60, 0xdd, 0x0b, 0x43, 0xfa, 0xf2, 0x5b, 0xe0, 0x35, 0xfd, 0x80, 0x17,
	0x5b, 0xb5, 0x44, 0x13, 0x6a, 0x80, 0x8d, 0x6f, 0x5e, 0x8c, 0xa3, 0x60, 0xe0, 0xb1, 0x5c, 0xa5,
	0x2a, 0x77, 0x99, 0xb3, 0xb1, 0x0c, 0xd8, 0x10, 0xc9, 0xdf, 0x82, 0xa1, 0x54, 0x4b, 0x33, 0x98,
	0xec, 0x1b, 0xff, 0x94, 0xc0, 0x76, 0x56, 0x72, 0x3c, 0xfa, 0x4e, 0x46, 0x94, 0x0e, 0x58, 0x2b,
	0xd9, 0x67, 0xd6, 0xca, 0x8a, 0x95, 0xed, 0x98, 0x0e, 0x5b, 0xf1, 0x4a, 0xd2, 0x3a, 0xa7, 0x7b,
	0x76, 0x47, 0x6c, 0xad, 0xf9, 0xdf, 0xfd, 0x30, 0x2b, 0x73, 0x66, 0x98, 0x50, 0x27, 0x48, 0x42,
	0x9f, 0x97, 0x59, 0xb3, 0x66, 0x86, 0xb9, 0xfc, 0x2a, 0xf3, 0xf9, 0xb1, 0x48, 0x3f, 0x4a, 0xfc,
	0x91, 0x4c, 0x83, 0x28, 0xeb, 0xe7, 0xcc, 0x70, 0xf9, 0x57, 0x75, 0x3a, 0xf3, 0xe9, 0xb8, 0xb2,
	0x01, 0x0e, 0xfa, 0x48, 0x02, 0x7b, 0x26, 0x36, 0x7a, 0xd8, 0x20, 0xa6, 0xda, 0x22, 0x1d, 0x05,
	0x5b, 0x4e, 0x53, 0xc1, 0x0e, 0x5c, 0x42, 0xe7, 0xe0, 0x58, 0x20, 0xb2, 0xa5, 0x60, 0x47, 0x21,
	0x2d, 0xc5, 0xbe, 0x25, 0x96, 0xf2, 0x05, 0x96, 0xd0, 0x19, 0x38, 0x12, 0xb8, 0xa2, 0x3b, 0x8a,
	0x35, 0xc3, 0xe5, 0xc5, 0x18, 0xcb, 0xb7, 0x70, 0x19, 0x9d, 0x02, 0xa9, 0x18, 0x37, 0x65, 0xb8,
	0x92, 0xcb, 0xca, 0x52, 0x70, 0xab, 0xc7, 0x65, 0x2b, 0xe8, 0x10, 0xec, 0xbe, 0x21, 0x4d, 0x19,
	0xae, 0xe6, 0xc0, 0x0d, 0x36, 0xb1, 0x4a, 0x0c, 0x07, 0xae, 0xa1, 0x23, 0xb0, 0x2f, 0x80, 0x96,
	0xe1, 0x36, 0x35, 0x85, 0x8b, 0x55, 0x73, 0xc7, 0x64, 0xa8, 0x29, 0xc3, 0xda, 0x94, 0x98, 0x69,
	0xf1, 0x1d, 0x37, 0x93, 0x03, 0x39, 0xb9, 0x8c, 0x30, 0xb9, 0xf5, 0x9c, 0x5c, 0x86, 0x9a, 0x32,
	0xdc, 0xc8, 0x25, 0x67, 0x62, 0xdb, 0xe6, 0x21, 0x9b, 0xe8, 0x00, 0xa0, 0x3c, 0x68, 0xca, 0x70,
	0x2b, 0x77, 0xf9, 0x6d, 0x7c, 0xa7, 0x10, 0x45, 0x6f, 0x11, 0x43, 0xe7, 0xf9, 0x6d, 0x2f, 0xe2,
	0x58, 0xd3, 0x18, 0x87, 0xe8, 0x23, 0x38, 0x13, 0x38, 0x36, 0x4d, 0xad, 0x47, 0x5a, 0xaa, 0x6d,
	0x1b, 0xda, 0xd7, 0xb4, 0xf8, 0x9d, 0xf7, 0x5d, 0x58, 0x93, 0x10, 0xaa, 0x83, 0x53, 0xf1, 0x7e,
	0x26, 0xb0, 0x89, 0xe5, 0xb4, 0xcb, 0xbb, 0xef, 0x7a, 0x30, 0x8d, 0xbd, 0xdc, 0x31, 0x46, 0x57,
	0x67, 0x8d, 0x16, 0x33, 0xd9, 0x7f, 0xdf, 0x85, 0xa9, 0x1c, 0xa0, 0x0f, 0xe0, 0x44, 0x70, 0xb1,
	0x59, 0xad, 0x77, 0x8a, 0x6d, 0xe3, 0x76, 0xaa, 0x71, 0x98, 0x6b, 0xcb, 0x84, 0x35, 0x65, 0x28,
	0xe5, 0x26, 0x51, 0x53, 0xf0, 0x57, 0x61, 0x8e, 0x8f, 0x16, 0x63, 0x76, 0xf2, 0x31, 0xda, 0x07,
	0x3b, 0x19, 0x36, 0x6e, 0x6e, 0x34, 0x35, 0x6d, 0xc0, 0x89, 0x10, 0x25, 0x1b, 0xaa, 0x4e, 0x2c,
	0xc3, 0xb8, 0x23, 0x9a, 0x6a, 0x3b, 0x5c, 0xf4, 0x74, 0x31, 0x66, 0xa2, 0x67, 0xb9, 0xe9, 0xb7,
	0x55, 0xf6, 0xcd, 0x32, 0xba, 0x3a, 0x0f, 0x3e, 0x47, 0xc7, 0xe0, 0x20, 0x4f, 0x5d, 0x93, 0xb3,
	0x0f, 0xb9, 0x3a, 0x4d, 0x2b, 0x9b, 0x0b, 0x07, 0xd6, 0x73, 0x61, 0x13, 0xc4, 0xc2, 0x3e, 0x2e,
	0x60, 0x2c, 0x99, 0xc6, 0xe5, 0xdf, 0x15, 0xb0, 0x5f, 0xf8, 0x4b, 0x86, 0x4e, 0xc0, 0xe1, 0x24,
	0xca, 0xb8, 0x55, 0x2c, 0xe2, 0xf4, 0x4c, 0x85, 0x18, 0x4e, 0x47, 0xb1, 0xe0, 0x92, 0x78, 0x6f,
	0x33, 0xd8, 0xc2, 0x3a, 0x31, 0xb1, 0x3a, 0xff, 0x3c, 0x08, 0xd8, 0x55, 0x39, 0x2e, 0x8b, 0x03,
	0x3c, 0x1f, 0x6d, 0x77, 0x5c, 0x1d, 0x2e, 0x8b, 0x33, 0x31, 0xe3, 0x76, 0xc7, 0xc5, 0x7a, 0x3b,
	0x75, 0x59, 0x11, 0x6b, 0x9a, 0xb9, 0xfc, 0x40, 0x34, 0x17, 0x56, 0x8a, 0xc3, 0xaf, 0xc9, 0xbd,
	0x8a, 0x0d, 0xd2, 0xc5, 0x7a, 0x1b, 0xae, 0x16, 0x87, 0x7f, 0x66, 0xe1, 0x6b, 0xa8, 0x01, 0xce,
	0x8b, 0xb2, 0x13, 0xe2, 0xab, 0xc5, 0x15, 0x70, 0x87, 0x1e, 0x36, 0x60, 0xad, 0x58, 0xff, 0x47,
	0xa6, 0x0f, 0xc4, 0x71, 0x16, 0xd3, 0x6b, 0xe1, 0x54, 0x7c, 0xfd, 0x8d, 0xc3, 0x35, 0x17, 0xfe,
	0x34, 0x71, 0xd8, 0x28, 0x56, 0xff, 0x89, 0xa9, 0x6f, 0x16, 0xab, 0xff, 0xde, 0x31, 0xf4, 0x36,
	0x4f, 0x6d, 0xeb, 0xad, 0x7a, 0x0b, 0x93, 0x4f, 0xb3, 0xda, 0xb6, 0x8b, 0xd5, 0x7f, 0x66, 0xea,
	0xb0, 0x98, 0xfd, 0xc2, 0xd8, 0x8e, 0xf8, 0xce, 0x71, 0xf6, 0x39, 0x15, 0x44, 0xc5, 0xd3, 0xa0,
	0x65, 0x77, 0x35, 0xf7, 0x8c, 0xcc, 0xe1, 0x89, 0xc7, 0xde, 0xe5, 0x1f, 0xd3, 0x3f, 0x52, 0x71,
	0x44, 0x5f, 0x48, 0xc2, 0xc6, 0x73, 0xf6, 0x78, 0xda, 0xba, 0xd1, 0x25, 0x5d, 0x97, 0xdc, 0xbb,
	0x0a, 0x2c, 0x09, 0xef, 0x30, 0xb7, 0xf3, 0xc2, 0x18, 0x29, 0xe7, 0x23, 0x58, 0x47, 0x5d, 0x05,
	0x2e, 0x5f, 0x12, 0x70, 0x28, 0x7e, 0x03, 0xf8, 0x5f, 0x9a, 0x38, 0xf1, 0x92, 0x71, 0x2c, 0x4c,
	0x12, 0x7f, 0x14, 0x6c, 0x07, 0x3b, 0xae, 0x4d, 0xba, 0x58, 0x75, 0xd2, 0x5f, 0x1f, 0x58, 0x5a,
	0xe0, 0x62, 0xe8, 0xec, 0x5d, 0xbe, 0x53, 0xf5, 0x36, 0x2c, 0x9b, 0x4b, 0xff, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x76, 0xe9, 0xd4, 0x9e, 0xa5, 0x0a, 0x00, 0x00,
}
