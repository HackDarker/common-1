// Code generated by protoc-gen-go. DO NOT EDIT.
// source: niuniu_base.proto

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

type NiuniuEnumProtoid int32

const (
	// //////////////////////////////////
	NiuniuEnumProtoid_NIU_PID_HEARTBEAT          NiuniuEnumProtoid = 0
	NiuniuEnumProtoid_NIU_PID_QUICK_CONN         NiuniuEnumProtoid = 1
	NiuniuEnumProtoid_NIU_PID_QUICK_CONN_ACK     NiuniuEnumProtoid = 2
	NiuniuEnumProtoid_NIU_PID_GAME_LOGIN         NiuniuEnumProtoid = 3
	NiuniuEnumProtoid_NIU_PID_GAME_LOGIN_ACK     NiuniuEnumProtoid = 4
	NiuniuEnumProtoid_NIU_PID_CREATE_DESK_REQ    NiuniuEnumProtoid = 5
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_REQ     NiuniuEnumProtoid = 6
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_ACK     NiuniuEnumProtoid = 7
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_BC      NiuniuEnumProtoid = 8
	NiuniuEnumProtoid_NIU_PID_READY_REQ          NiuniuEnumProtoid = 9
	NiuniuEnumProtoid_NIU_PID_READY_ACK          NiuniuEnumProtoid = 10
	NiuniuEnumProtoid_NIU_PID_READY_BC           NiuniuEnumProtoid = 11
	NiuniuEnumProtoid_NIU_PID_START_GAME_OT      NiuniuEnumProtoid = 12
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_OT     NiuniuEnumProtoid = 13
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_REQ    NiuniuEnumProtoid = 14
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_ACK    NiuniuEnumProtoid = 15
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_BC     NiuniuEnumProtoid = 16
	NiuniuEnumProtoid_NIU_PID_JIABEI_OT          NiuniuEnumProtoid = 17
	NiuniuEnumProtoid_NIU_PID_JIABEI_REQ         NiuniuEnumProtoid = 18
	NiuniuEnumProtoid_NIU_PID_JIABEI_ACK         NiuniuEnumProtoid = 19
	NiuniuEnumProtoid_NIU_PID_JIABEI_BC          NiuniuEnumProtoid = 20
	NiuniuEnumProtoid_NIU_PID_BIPAI_RESULT_BC    NiuniuEnumProtoid = 21
	NiuniuEnumProtoid_NIU_PID_GAME_END_BC        NiuniuEnumProtoid = 22
	NiuniuEnumProtoid_NIU_PID_APPLY_DISSOLVE_REQ NiuniuEnumProtoid = 23
	NiuniuEnumProtoid_NIU_PID_APPLY_DISSOLVE_ACK NiuniuEnumProtoid = 24
	NiuniuEnumProtoid_NIU_PID_DISSOLVE_BACK_REQ  NiuniuEnumProtoid = 25
	NiuniuEnumProtoid_NIU_PID_DISSOLVE_BACK_ACK  NiuniuEnumProtoid = 26
	NiuniuEnumProtoid_NIU_PID_DISSOLVE_DONE_BC   NiuniuEnumProtoid = 27
	NiuniuEnumProtoid_NIU_PID_SEND_MESSAGE_REQ   NiuniuEnumProtoid = 28
	NiuniuEnumProtoid_NIU_PID_MESSAGE_BC         NiuniuEnumProtoid = 29
	NiuniuEnumProtoid_NIU_PID_LEAVE_DESK_REQ     NiuniuEnumProtoid = 30
	NiuniuEnumProtoid_NIU_PID_LEAVE_DESK_ACK     NiuniuEnumProtoid = 31
	NiuniuEnumProtoid_NIU_PID_OWNER_DISSOLVE_REQ NiuniuEnumProtoid = 32
	NiuniuEnumProtoid_NIU_PID_OWNER_DISSOLVE_ACK NiuniuEnumProtoid = 33
	NiuniuEnumProtoid_NIU_OFFLINE_BC             NiuniuEnumProtoid = 34
	NiuniuEnumProtoid_NIU_COIN_ROOM_LIST_REQ     NiuniuEnumProtoid = 35
	NiuniuEnumProtoid_NIU_COIN_ROOM_LIST_ACK     NiuniuEnumProtoid = 36
)

var NiuniuEnumProtoid_name = map[int32]string{
	0:  "NIU_PID_HEARTBEAT",
	1:  "NIU_PID_QUICK_CONN",
	2:  "NIU_PID_QUICK_CONN_ACK",
	3:  "NIU_PID_GAME_LOGIN",
	4:  "NIU_PID_GAME_LOGIN_ACK",
	5:  "NIU_PID_CREATE_DESK_REQ",
	6:  "NIU_PID_ENTER_DESK_REQ",
	7:  "NIU_PID_ENTER_DESK_ACK",
	8:  "NIU_PID_ENTER_DESK_BC",
	9:  "NIU_PID_READY_REQ",
	10: "NIU_PID_READY_ACK",
	11: "NIU_PID_READY_BC",
	12: "NIU_PID_START_GAME_OT",
	13: "NIU_PID_QIANGZHUANG_OT",
	14: "NIU_PID_QIANGZHUANG_REQ",
	15: "NIU_PID_QIANGZHUANG_ACK",
	16: "NIU_PID_QIANGZHUANG_BC",
	17: "NIU_PID_JIABEI_OT",
	18: "NIU_PID_JIABEI_REQ",
	19: "NIU_PID_JIABEI_ACK",
	20: "NIU_PID_JIABEI_BC",
	21: "NIU_PID_BIPAI_RESULT_BC",
	22: "NIU_PID_GAME_END_BC",
	23: "NIU_PID_APPLY_DISSOLVE_REQ",
	24: "NIU_PID_APPLY_DISSOLVE_ACK",
	25: "NIU_PID_DISSOLVE_BACK_REQ",
	26: "NIU_PID_DISSOLVE_BACK_ACK",
	27: "NIU_PID_DISSOLVE_DONE_BC",
	28: "NIU_PID_SEND_MESSAGE_REQ",
	29: "NIU_PID_MESSAGE_BC",
	30: "NIU_PID_LEAVE_DESK_REQ",
	31: "NIU_PID_LEAVE_DESK_ACK",
	32: "NIU_PID_OWNER_DISSOLVE_REQ",
	33: "NIU_PID_OWNER_DISSOLVE_ACK",
	34: "NIU_OFFLINE_BC",
	35: "NIU_COIN_ROOM_LIST_REQ",
	36: "NIU_COIN_ROOM_LIST_ACK",
}
var NiuniuEnumProtoid_value = map[string]int32{
	"NIU_PID_HEARTBEAT":          0,
	"NIU_PID_QUICK_CONN":         1,
	"NIU_PID_QUICK_CONN_ACK":     2,
	"NIU_PID_GAME_LOGIN":         3,
	"NIU_PID_GAME_LOGIN_ACK":     4,
	"NIU_PID_CREATE_DESK_REQ":    5,
	"NIU_PID_ENTER_DESK_REQ":     6,
	"NIU_PID_ENTER_DESK_ACK":     7,
	"NIU_PID_ENTER_DESK_BC":      8,
	"NIU_PID_READY_REQ":          9,
	"NIU_PID_READY_ACK":          10,
	"NIU_PID_READY_BC":           11,
	"NIU_PID_START_GAME_OT":      12,
	"NIU_PID_QIANGZHUANG_OT":     13,
	"NIU_PID_QIANGZHUANG_REQ":    14,
	"NIU_PID_QIANGZHUANG_ACK":    15,
	"NIU_PID_QIANGZHUANG_BC":     16,
	"NIU_PID_JIABEI_OT":          17,
	"NIU_PID_JIABEI_REQ":         18,
	"NIU_PID_JIABEI_ACK":         19,
	"NIU_PID_JIABEI_BC":          20,
	"NIU_PID_BIPAI_RESULT_BC":    21,
	"NIU_PID_GAME_END_BC":        22,
	"NIU_PID_APPLY_DISSOLVE_REQ": 23,
	"NIU_PID_APPLY_DISSOLVE_ACK": 24,
	"NIU_PID_DISSOLVE_BACK_REQ":  25,
	"NIU_PID_DISSOLVE_BACK_ACK":  26,
	"NIU_PID_DISSOLVE_DONE_BC":   27,
	"NIU_PID_SEND_MESSAGE_REQ":   28,
	"NIU_PID_MESSAGE_BC":         29,
	"NIU_PID_LEAVE_DESK_REQ":     30,
	"NIU_PID_LEAVE_DESK_ACK":     31,
	"NIU_PID_OWNER_DISSOLVE_REQ": 32,
	"NIU_PID_OWNER_DISSOLVE_ACK": 33,
	"NIU_OFFLINE_BC":             34,
	"NIU_COIN_ROOM_LIST_REQ":     35,
	"NIU_COIN_ROOM_LIST_ACK":     36,
}

func (x NiuniuEnumProtoid) Enum() *NiuniuEnumProtoid {
	p := new(NiuniuEnumProtoid)
	*p = x
	return p
}
func (x NiuniuEnumProtoid) String() string {
	return proto.EnumName(NiuniuEnumProtoid_name, int32(x))
}
func (x *NiuniuEnumProtoid) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumProtoid_value, data, "NiuniuEnumProtoid")
	if err != nil {
		return err
	}
	*x = NiuniuEnumProtoid(value)
	return nil
}
func (NiuniuEnumProtoid) EnumDescriptor() ([]byte, []int) { return fileDescriptor33, []int{0} }

// =================================公共================================
// 牛牛牌的类型
type NiuniuEnum_PokerType int32

const (
	NiuniuEnum_PokerType_NO_NIU           NiuniuEnum_PokerType = 1
	NiuniuEnum_PokerType_NIU_1            NiuniuEnum_PokerType = 2
	NiuniuEnum_PokerType_NIU_2            NiuniuEnum_PokerType = 3
	NiuniuEnum_PokerType_NIU_3            NiuniuEnum_PokerType = 4
	NiuniuEnum_PokerType_NIU_4            NiuniuEnum_PokerType = 5
	NiuniuEnum_PokerType_NIU_5            NiuniuEnum_PokerType = 6
	NiuniuEnum_PokerType_NIU_6            NiuniuEnum_PokerType = 7
	NiuniuEnum_PokerType_NIU_7            NiuniuEnum_PokerType = 8
	NiuniuEnum_PokerType_NIU_8            NiuniuEnum_PokerType = 9
	NiuniuEnum_PokerType_NIU_9            NiuniuEnum_PokerType = 10
	NiuniuEnum_PokerType_NIU_NIU          NiuniuEnum_PokerType = 11
	NiuniuEnum_PokerType_YIN_NIU          NiuniuEnum_PokerType = 12
	NiuniuEnum_PokerType_JIN_NIU          NiuniuEnum_PokerType = 13
	NiuniuEnum_PokerType_WU_XIAO_NIU      NiuniuEnum_PokerType = 14
	NiuniuEnum_PokerType_NIU_ZHA_DAN      NiuniuEnum_PokerType = 15
	NiuniuEnum_PokerType_NIU_YI_TIAO_LONG NiuniuEnum_PokerType = 16
)

var NiuniuEnum_PokerType_name = map[int32]string{
	1:  "NO_NIU",
	2:  "NIU_1",
	3:  "NIU_2",
	4:  "NIU_3",
	5:  "NIU_4",
	6:  "NIU_5",
	7:  "NIU_6",
	8:  "NIU_7",
	9:  "NIU_8",
	10: "NIU_9",
	11: "NIU_NIU",
	12: "YIN_NIU",
	13: "JIN_NIU",
	14: "WU_XIAO_NIU",
	15: "NIU_ZHA_DAN",
	16: "NIU_YI_TIAO_LONG",
}
var NiuniuEnum_PokerType_value = map[string]int32{
	"NO_NIU":           1,
	"NIU_1":            2,
	"NIU_2":            3,
	"NIU_3":            4,
	"NIU_4":            5,
	"NIU_5":            6,
	"NIU_6":            7,
	"NIU_7":            8,
	"NIU_8":            9,
	"NIU_9":            10,
	"NIU_NIU":          11,
	"YIN_NIU":          12,
	"JIN_NIU":          13,
	"WU_XIAO_NIU":      14,
	"NIU_ZHA_DAN":      15,
	"NIU_YI_TIAO_LONG": 16,
}

func (x NiuniuEnum_PokerType) Enum() *NiuniuEnum_PokerType {
	p := new(NiuniuEnum_PokerType)
	*p = x
	return p
}
func (x NiuniuEnum_PokerType) String() string {
	return proto.EnumName(NiuniuEnum_PokerType_name, int32(x))
}
func (x *NiuniuEnum_PokerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnum_PokerType_value, data, "NiuniuEnum_PokerType")
	if err != nil {
		return err
	}
	*x = NiuniuEnum_PokerType(value)
	return nil
}
func (NiuniuEnum_PokerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor33, []int{1} }

// 房间状态
type NiuniuEnumDeskState int32

const (
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER       NiuniuEnumDeskState = 1
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_READY       NiuniuEnumDeskState = 2
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_START       NiuniuEnumDeskState = 3
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_QIANGZHUANG NiuniuEnumDeskState = 4
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_JIABEI      NiuniuEnumDeskState = 5
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_BIPAI       NiuniuEnumDeskState = 6
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_RESULT      NiuniuEnumDeskState = 7
)

var NiuniuEnumDeskState_name = map[int32]string{
	1: "NIU_DESK_STATUS_WAIT_ENTER",
	2: "NIU_DESK_STATUS_WAIT_READY",
	3: "NIU_DESK_STATUS_WAIT_START",
	4: "NIU_DESK_STATUS_WAIT_QIANGZHUANG",
	5: "NIU_DESK_STATUS_WAIT_JIABEI",
	6: "NIU_DESK_STATUS_WAIT_BIPAI",
	7: "NIU_DESK_STATUS_WAIT_RESULT",
}
var NiuniuEnumDeskState_value = map[string]int32{
	"NIU_DESK_STATUS_WAIT_ENTER":       1,
	"NIU_DESK_STATUS_WAIT_READY":       2,
	"NIU_DESK_STATUS_WAIT_START":       3,
	"NIU_DESK_STATUS_WAIT_QIANGZHUANG": 4,
	"NIU_DESK_STATUS_WAIT_JIABEI":      5,
	"NIU_DESK_STATUS_WAIT_BIPAI":       6,
	"NIU_DESK_STATUS_WAIT_RESULT":      7,
}

func (x NiuniuEnumDeskState) Enum() *NiuniuEnumDeskState {
	p := new(NiuniuEnumDeskState)
	*p = x
	return p
}
func (x NiuniuEnumDeskState) String() string {
	return proto.EnumName(NiuniuEnumDeskState_name, int32(x))
}
func (x *NiuniuEnumDeskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumDeskState_value, data, "NiuniuEnumDeskState")
	if err != nil {
		return err
	}
	*x = NiuniuEnumDeskState(value)
	return nil
}
func (NiuniuEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor33, []int{2} }

// 坐庄规则
type NiuniuEnumBankerRule int32

const (
	NiuniuEnumBankerRule_DING_ZHUANG        NiuniuEnumBankerRule = 1
	NiuniuEnumBankerRule_SUI_JI_ZUO_ZHUANG  NiuniuEnumBankerRule = 2
	NiuniuEnumBankerRule_QIANG_ZHUANG       NiuniuEnumBankerRule = 3
	NiuniuEnumBankerRule_FANGZHU_DINGZHUANG NiuniuEnumBankerRule = 4
)

var NiuniuEnumBankerRule_name = map[int32]string{
	1: "DING_ZHUANG",
	2: "SUI_JI_ZUO_ZHUANG",
	3: "QIANG_ZHUANG",
	4: "FANGZHU_DINGZHUANG",
}
var NiuniuEnumBankerRule_value = map[string]int32{
	"DING_ZHUANG":        1,
	"SUI_JI_ZUO_ZHUANG":  2,
	"QIANG_ZHUANG":       3,
	"FANGZHU_DINGZHUANG": 4,
}

func (x NiuniuEnumBankerRule) Enum() *NiuniuEnumBankerRule {
	p := new(NiuniuEnumBankerRule)
	*p = x
	return p
}
func (x NiuniuEnumBankerRule) String() string {
	return proto.EnumName(NiuniuEnumBankerRule_name, int32(x))
}
func (x *NiuniuEnumBankerRule) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumBankerRule_value, data, "NiuniuEnumBankerRule")
	if err != nil {
		return err
	}
	*x = NiuniuEnumBankerRule(value)
	return nil
}
func (NiuniuEnumBankerRule) EnumDescriptor() ([]byte, []int) { return fileDescriptor33, []int{3} }

// 打出去的牌
type NiuniuClientPoker struct {
	Pais []*ClientBasePoker    `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type *NiuniuEnum_PokerType `protobuf:"varint,3,opt,name=type,enum=ddproto.NiuniuEnum_PokerType" json:"type,omitempty"`
	// 金币场
	SelectedId       []int32 `protobuf:"varint,4,rep,name=selectedId" json:"selectedId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuniuClientPoker) Reset()                    { *m = NiuniuClientPoker{} }
func (m *NiuniuClientPoker) String() string            { return proto.CompactTextString(m) }
func (*NiuniuClientPoker) ProtoMessage()               {}
func (*NiuniuClientPoker) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{0} }

func (m *NiuniuClientPoker) GetPais() []*ClientBasePoker {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *NiuniuClientPoker) GetType() NiuniuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NiuniuEnum_PokerType_NO_NIU
}

func (m *NiuniuClientPoker) GetSelectedId() []int32 {
	if m != nil {
		return m.SelectedId
	}
	return nil
}

// 对局账单信息
type NiuniuUserBill struct {
	Score            *int32  `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	CountHasNiu      *int32  `protobuf:"varint,2,opt,name=count_has_niu,json=countHasNiu" json:"count_has_niu,omitempty"`
	CountNoNiu       *int32  `protobuf:"varint,3,opt,name=count_no_niu,json=countNoNiu" json:"count_no_niu,omitempty"`
	CountWin         *int32  `protobuf:"varint,4,opt,name=count_win,json=countWin" json:"count_win,omitempty"`
	CountLost        *int32  `protobuf:"varint,5,opt,name=count_lost,json=countLost" json:"count_lost,omitempty"`
	UserId           *uint32 `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuniuUserBill) Reset()                    { *m = NiuniuUserBill{} }
func (m *NiuniuUserBill) String() string            { return proto.CompactTextString(m) }
func (*NiuniuUserBill) ProtoMessage()               {}
func (*NiuniuUserBill) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{1} }

func (m *NiuniuUserBill) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *NiuniuUserBill) GetCountHasNiu() int32 {
	if m != nil && m.CountHasNiu != nil {
		return *m.CountHasNiu
	}
	return 0
}

func (m *NiuniuUserBill) GetCountNoNiu() int32 {
	if m != nil && m.CountNoNiu != nil {
		return *m.CountNoNiu
	}
	return 0
}

func (m *NiuniuUserBill) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *NiuniuUserBill) GetCountLost() int32 {
	if m != nil && m.CountLost != nil {
		return *m.CountLost
	}
	return 0
}

func (m *NiuniuUserBill) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// desk 配置选项
type NiuniuDeskOption struct {
	MinUser         *int32                `protobuf:"varint,1,opt,name=minUser" json:"minUser,omitempty"`
	MaxUser         *int32                `protobuf:"varint,2,opt,name=maxUser" json:"maxUser,omitempty"`
	MaxCircle       *int32                `protobuf:"varint,3,opt,name=maxCircle" json:"maxCircle,omitempty"`
	HasFlower       *bool                 `protobuf:"varint,4,opt,name=hasFlower" json:"hasFlower,omitempty"`
	BankRule        *NiuniuEnumBankerRule `protobuf:"varint,5,opt,name=bankRule,enum=ddproto.NiuniuEnumBankerRule" json:"bankRule,omitempty"`
	IsFlowerPlay    *bool                 `protobuf:"varint,6,opt,name=isFlowerPlay" json:"isFlowerPlay,omitempty"`
	IsJiaoFenJiaBei *bool                 `protobuf:"varint,7,opt,name=isJiaoFenJiaBei" json:"isJiaoFenJiaBei,omitempty"`
	HasAnimation    *bool                 `protobuf:"varint,8,opt,name=hasAnimation" json:"hasAnimation,omitempty"`
	// 金币场
	IsCoinRoom       *bool  `protobuf:"varint,9,opt,name=isCoinRoom" json:"isCoinRoom,omitempty"`
	BaseScore        *int32 `protobuf:"varint,10,opt,name=baseScore" json:"baseScore,omitempty"`
	NeedPwd          *bool  `protobuf:"varint,11,opt,name=needPwd" json:"needPwd,omitempty"`
	MinEnterScore    *int32 `protobuf:"varint,12,opt,name=minEnterScore" json:"minEnterScore,omitempty"`
	MaxQzScore       *int32 `protobuf:"varint,13,opt,name=maxQzScore" json:"maxQzScore,omitempty"`
	DenyHalfJoin     *bool  `protobuf:"varint,14,opt,name=denyHalfJoin" json:"denyHalfJoin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NiuniuDeskOption) Reset()                    { *m = NiuniuDeskOption{} }
func (m *NiuniuDeskOption) String() string            { return proto.CompactTextString(m) }
func (*NiuniuDeskOption) ProtoMessage()               {}
func (*NiuniuDeskOption) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{2} }

func (m *NiuniuDeskOption) GetMinUser() int32 {
	if m != nil && m.MinUser != nil {
		return *m.MinUser
	}
	return 0
}

func (m *NiuniuDeskOption) GetMaxUser() int32 {
	if m != nil && m.MaxUser != nil {
		return *m.MaxUser
	}
	return 0
}

func (m *NiuniuDeskOption) GetMaxCircle() int32 {
	if m != nil && m.MaxCircle != nil {
		return *m.MaxCircle
	}
	return 0
}

func (m *NiuniuDeskOption) GetHasFlower() bool {
	if m != nil && m.HasFlower != nil {
		return *m.HasFlower
	}
	return false
}

func (m *NiuniuDeskOption) GetBankRule() NiuniuEnumBankerRule {
	if m != nil && m.BankRule != nil {
		return *m.BankRule
	}
	return NiuniuEnumBankerRule_DING_ZHUANG
}

func (m *NiuniuDeskOption) GetIsFlowerPlay() bool {
	if m != nil && m.IsFlowerPlay != nil {
		return *m.IsFlowerPlay
	}
	return false
}

func (m *NiuniuDeskOption) GetIsJiaoFenJiaBei() bool {
	if m != nil && m.IsJiaoFenJiaBei != nil {
		return *m.IsJiaoFenJiaBei
	}
	return false
}

func (m *NiuniuDeskOption) GetHasAnimation() bool {
	if m != nil && m.HasAnimation != nil {
		return *m.HasAnimation
	}
	return false
}

func (m *NiuniuDeskOption) GetIsCoinRoom() bool {
	if m != nil && m.IsCoinRoom != nil {
		return *m.IsCoinRoom
	}
	return false
}

func (m *NiuniuDeskOption) GetBaseScore() int32 {
	if m != nil && m.BaseScore != nil {
		return *m.BaseScore
	}
	return 0
}

func (m *NiuniuDeskOption) GetNeedPwd() bool {
	if m != nil && m.NeedPwd != nil {
		return *m.NeedPwd
	}
	return false
}

func (m *NiuniuDeskOption) GetMinEnterScore() int32 {
	if m != nil && m.MinEnterScore != nil {
		return *m.MinEnterScore
	}
	return 0
}

func (m *NiuniuDeskOption) GetMaxQzScore() int32 {
	if m != nil && m.MaxQzScore != nil {
		return *m.MaxQzScore
	}
	return 0
}

func (m *NiuniuDeskOption) GetDenyHalfJoin() bool {
	if m != nil && m.DenyHalfJoin != nil {
		return *m.DenyHalfJoin
	}
	return false
}

// room 的信息
type NiuniuSrvRoom struct {
	RoomId           *int32                `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32                `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32                `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string               `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	BaseChip         *int32                `protobuf:"varint,5,opt,name=baseChip" json:"baseChip,omitempty"`
	EnterCoin        *int32                `protobuf:"varint,6,opt,name=enterCoin" json:"enterCoin,omitempty"`
	BankRule         *NiuniuEnumBankerRule `protobuf:"varint,7,opt,name=bankRule,enum=ddproto.NiuniuEnumBankerRule" json:"bankRule,omitempty"`
	MaxDeskGammer    *int32                `protobuf:"varint,8,opt,name=maxDeskGammer" json:"maxDeskGammer,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuSrvRoom) Reset()                    { *m = NiuniuSrvRoom{} }
func (m *NiuniuSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvRoom) ProtoMessage()               {}
func (*NiuniuSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{3} }

func (m *NiuniuSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func (m *NiuniuSrvRoom) GetBaseChip() int32 {
	if m != nil && m.BaseChip != nil {
		return *m.BaseChip
	}
	return 0
}

func (m *NiuniuSrvRoom) GetEnterCoin() int32 {
	if m != nil && m.EnterCoin != nil {
		return *m.EnterCoin
	}
	return 0
}

func (m *NiuniuSrvRoom) GetBankRule() NiuniuEnumBankerRule {
	if m != nil && m.BankRule != nil {
		return *m.BankRule
	}
	return NiuniuEnumBankerRule_DING_ZHUANG
}

func (m *NiuniuSrvRoom) GetMaxDeskGammer() int32 {
	if m != nil && m.MaxDeskGammer != nil {
		return *m.MaxDeskGammer
	}
	return 0
}

func init() {
	proto.RegisterType((*NiuniuClientPoker)(nil), "ddproto.niuniu_client_poker")
	proto.RegisterType((*NiuniuUserBill)(nil), "ddproto.niuniu_user_bill")
	proto.RegisterType((*NiuniuDeskOption)(nil), "ddproto.niuniu_desk_option")
	proto.RegisterType((*NiuniuSrvRoom)(nil), "ddproto.niuniu_srv_room")
	proto.RegisterEnum("ddproto.NiuniuEnumProtoid", NiuniuEnumProtoid_name, NiuniuEnumProtoid_value)
	proto.RegisterEnum("ddproto.NiuniuEnum_PokerType", NiuniuEnum_PokerType_name, NiuniuEnum_PokerType_value)
	proto.RegisterEnum("ddproto.NiuniuEnumDeskState", NiuniuEnumDeskState_name, NiuniuEnumDeskState_value)
	proto.RegisterEnum("ddproto.NiuniuEnumBankerRule", NiuniuEnumBankerRule_name, NiuniuEnumBankerRule_value)
}

func init() { proto.RegisterFile("niuniu_base.proto", fileDescriptor33) }

var fileDescriptor33 = []byte{
	// 1223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x8e, 0xfe, 0xa5, 0x91, 0x7f, 0x36, 0xeb, 0xd8, 0x66, 0xec, 0xc4, 0x51, 0xd5, 0x1c, 0x0c,
	0x1f, 0x0c, 0xd4, 0xfd, 0x07, 0x7a, 0xa1, 0x28, 0x5a, 0xa6, 0xa3, 0x90, 0x0a, 0x49, 0xc5, 0x75,
	0x2e, 0x0b, 0xc6, 0xda, 0xc2, 0x0b, 0x53, 0xa4, 0x40, 0x4a, 0x89, 0xdd, 0x53, 0x5f, 0xa3, 0xa7,
	0xbe, 0x4b, 0xdf, 0xa5, 0xaf, 0xd0, 0x6b, 0x8b, 0x59, 0x92, 0x12, 0xd5, 0x48, 0x01, 0x7a, 0xf2,
	0xcc, 0xf7, 0xcd, 0x7c, 0x3b, 0x3f, 0xbb, 0xb4, 0xe0, 0x71, 0x20, 0x66, 0x81, 0x98, 0xb1, 0xf7,
	0x5e, 0xcc, 0x4f, 0x27, 0x51, 0x38, 0x0d, 0x69, 0x6d, 0x34, 0x92, 0xc6, 0xc1, 0xce, 0x4d, 0x38,
	0x1e, 0x87, 0x01, 0xbb, 0xf1, 0x05, 0x0f, 0xa6, 0x09, 0xdb, 0xfe, 0xbd, 0x00, 0x3b, 0x69, 0x4e,
	0x82, 0xb3, 0x49, 0x78, 0xc7, 0x23, 0x7a, 0x0a, 0xe5, 0x89, 0x27, 0x62, 0xa5, 0xd8, 0x2a, 0x1d,
	0x37, 0xcf, 0x0e, 0x4e, 0x53, 0x91, 0xd3, 0x34, 0x08, 0xf5, 0x93, 0x48, 0x5b, 0xc6, 0xd1, 0x33,
	0x28, 0x4f, 0x1f, 0x26, 0x5c, 0x29, 0xb5, 0x0a, 0xc7, 0x5b, 0x67, 0x47, 0xf3, 0xf8, 0x54, 0x9b,
	0x07, 0xb3, 0x31, 0x1b, 0x60, 0xbc, 0xfb, 0x30, 0xe1, 0xb6, 0x8c, 0xa5, 0x47, 0x00, 0x31, 0xf7,
	0xf9, 0xcd, 0x94, 0x8f, 0x8c, 0x91, 0x52, 0x6e, 0x95, 0x8e, 0x2b, 0x76, 0x0e, 0x69, 0xff, 0x59,
	0x00, 0x92, 0xe6, 0xcf, 0x62, 0x1e, 0xb1, 0xf7, 0xc2, 0xf7, 0xe9, 0x13, 0xa8, 0xc4, 0x37, 0x61,
	0xc4, 0x95, 0x42, 0xab, 0x70, 0x5c, 0xb1, 0x13, 0x87, 0xb6, 0x61, 0xf3, 0x26, 0x9c, 0x05, 0x53,
	0x76, 0xeb, 0xc5, 0x2c, 0x10, 0x33, 0xa5, 0x28, 0xd9, 0xa6, 0x04, 0x2f, 0xbc, 0xd8, 0x14, 0x33,
	0xda, 0x82, 0x8d, 0x24, 0x26, 0x08, 0x65, 0x48, 0x49, 0x86, 0x80, 0xc4, 0xcc, 0x10, 0x23, 0x0e,
	0xa1, 0x91, 0x44, 0x7c, 0x14, 0x81, 0x52, 0x96, 0x74, 0x5d, 0x02, 0x57, 0x22, 0xa0, 0xcf, 0x21,
	0x09, 0x65, 0x7e, 0x18, 0x4f, 0x95, 0x8a, 0x64, 0x93, 0xf0, 0x7e, 0x18, 0x4f, 0xe9, 0x1e, 0x54,
	0xb1, 0x48, 0x63, 0xa4, 0x54, 0x5b, 0x85, 0xe3, 0x4d, 0x3b, 0xf5, 0xda, 0x7f, 0x97, 0x80, 0xa6,
	0x4d, 0x8c, 0x78, 0x7c, 0xc7, 0xc2, 0xc9, 0x54, 0x84, 0x01, 0x55, 0xa0, 0x36, 0x16, 0xc1, 0x30,
	0xe6, 0x51, 0xda, 0x48, 0xe6, 0x4a, 0xc6, 0xbb, 0x97, 0x4c, 0x31, 0x65, 0x12, 0x97, 0x3e, 0x83,
	0xc6, 0xd8, 0xbb, 0xd7, 0x44, 0x74, 0xe3, 0xf3, 0xb4, 0xfa, 0x05, 0x80, 0xec, 0xad, 0x17, 0x9f,
	0xfb, 0xe1, 0x47, 0x1e, 0xc9, 0xe2, 0xeb, 0xf6, 0x02, 0xa0, 0x3f, 0x41, 0xfd, 0xbd, 0x17, 0xdc,
	0xd9, 0x33, 0x9f, 0xcb, 0xda, 0xb7, 0xce, 0x5a, 0x2b, 0x77, 0x84, 0x41, 0x3c, 0x62, 0xd1, 0xcc,
	0xe7, 0xf6, 0x3c, 0x83, 0xb6, 0x61, 0x43, 0xa4, 0x4a, 0x03, 0xdf, 0x7b, 0x90, 0x2d, 0xd6, 0xed,
	0x25, 0x8c, 0x1e, 0xc3, 0xb6, 0x88, 0x2f, 0x85, 0x17, 0x9e, 0xf3, 0xe0, 0x52, 0x78, 0x1d, 0x2e,
	0x94, 0x9a, 0x0c, 0xfb, 0x2f, 0x8c, 0x6a, 0xb7, 0x5e, 0xac, 0x06, 0x62, 0xec, 0xe1, 0x2c, 0x94,
	0x7a, 0xa2, 0x96, 0xc7, 0xf0, 0x6e, 0x88, 0x58, 0x0b, 0x45, 0x60, 0x87, 0xe1, 0x58, 0x69, 0xc8,
	0x88, 0x1c, 0x82, 0xdd, 0xe2, 0x1d, 0x74, 0xe4, 0x55, 0x80, 0x64, 0x16, 0x73, 0x00, 0x67, 0x18,
	0x70, 0x3e, 0x1a, 0x7c, 0x1c, 0x29, 0x4d, 0x99, 0x9a, 0xb9, 0xf4, 0x25, 0x6c, 0x8e, 0x45, 0xa0,
	0x07, 0x53, 0x1e, 0x25, 0xb9, 0x1b, 0x32, 0x77, 0x19, 0xc4, 0xd3, 0xc7, 0xde, 0xfd, 0x9b, 0x5f,
	0x93, 0x90, 0xcd, 0xe4, 0xa2, 0x2c, 0x10, 0xec, 0x60, 0xc4, 0x83, 0x87, 0x0b, 0xcf, 0xff, 0xe5,
	0x32, 0x14, 0x81, 0xb2, 0x95, 0x74, 0x90, 0xc7, 0xda, 0x7f, 0x14, 0x61, 0x3b, 0x9d, 0x6c, 0x1c,
	0x7d, 0x60, 0x11, 0x56, 0xbd, 0x07, 0x55, 0xfc, 0x6b, 0x8c, 0xd2, 0xa5, 0xa7, 0x1e, 0x3d, 0x80,
	0x3a, 0x5a, 0xf8, 0x36, 0xd2, 0xa5, 0xcf, 0x7d, 0xec, 0x14, 0xed, 0x3e, 0xff, 0xc0, 0xfd, 0x6c,
	0xeb, 0x73, 0x20, 0x63, 0x5d, 0x31, 0xf5, 0xb9, 0xdc, 0x7a, 0xc3, 0x5e, 0x00, 0xa8, 0x8b, 0x43,
	0xd1, 0x6e, 0xc5, 0x24, 0xbd, 0xb1, 0x73, 0x1f, 0x33, 0x39, 0x76, 0x8c, 0x23, 0x95, 0x0b, 0xad,
	0xd8, 0x0b, 0x60, 0xe9, 0xbe, 0xd4, 0xfe, 0xf7, 0x7d, 0xc1, 0x29, 0x7b, 0xf7, 0x5d, 0x1e, 0xdf,
	0xf5, 0xbc, 0xf1, 0x98, 0x47, 0x72, 0xc5, 0x38, 0xe5, 0x3c, 0x78, 0xf2, 0x4f, 0x6d, 0xfe, 0xed,
	0x91, 0x5a, 0x52, 0x5d, 0x8c, 0xe8, 0x2e, 0x3c, 0x36, 0x8d, 0x21, 0x1b, 0x18, 0x5d, 0x76, 0xa1,
	0xab, 0xb6, 0xdb, 0xd1, 0x55, 0x97, 0x3c, 0xa2, 0x7b, 0x40, 0x33, 0xf8, 0xcd, 0xd0, 0xd0, 0x5e,
	0x31, 0xcd, 0x32, 0x4d, 0x52, 0xa0, 0x07, 0xb0, 0xf7, 0x29, 0xce, 0x54, 0xed, 0x15, 0x29, 0xe6,
	0x73, 0x7a, 0xea, 0x6b, 0x9d, 0xf5, 0xad, 0x9e, 0x61, 0x92, 0x52, 0x3e, 0x67, 0x81, 0xcb, 0x9c,
	0x32, 0x3d, 0x84, 0xfd, 0x8c, 0xd3, 0x6c, 0x5d, 0x75, 0x75, 0xd6, 0xd5, 0x9d, 0x57, 0xcc, 0xd6,
	0xdf, 0x90, 0x4a, 0x3e, 0x51, 0x37, 0x5d, 0xdd, 0x5e, 0x70, 0xd5, 0x35, 0x1c, 0x8a, 0xd6, 0xe8,
	0x53, 0xd8, 0x5d, 0xc1, 0x75, 0x34, 0x52, 0xcf, 0xb7, 0x6b, 0xeb, 0x6a, 0xf7, 0x5a, 0xaa, 0x35,
	0x3e, 0x85, 0x51, 0x08, 0xe8, 0x13, 0x20, 0xcb, 0x70, 0x47, 0x23, 0xcd, 0xbc, 0xbc, 0xe3, 0xaa,
	0xb6, 0x9b, 0x74, 0x65, 0xb9, 0x64, 0x63, 0x69, 0x3c, 0x86, 0x6a, 0xf6, 0xde, 0x5d, 0x0c, 0x55,
	0xb3, 0x87, 0xdc, 0x66, 0xbe, 0xd5, 0x3c, 0x87, 0x05, 0x6c, 0xad, 0x23, 0xb1, 0x8c, 0xed, 0x75,
	0xaa, 0x1d, 0x8d, 0x90, 0x7c, 0xe5, 0x97, 0x86, 0xda, 0xd1, 0x0d, 0x3c, 0xec, 0x71, 0x7e, 0x17,
	0x29, 0x8c, 0xe7, 0xd0, 0x15, 0x38, 0x1e, 0xb1, 0xb3, 0x42, 0xa6, 0xa3, 0x91, 0x27, 0xf9, 0xb2,
	0x3a, 0xc6, 0x40, 0x45, 0x15, 0x67, 0xd8, 0x77, 0x91, 0xdc, 0xa5, 0xfb, 0xb0, 0xb3, 0xb4, 0x57,
	0xdd, 0xec, 0x22, 0xb1, 0x47, 0x8f, 0xe0, 0x20, 0x23, 0xd4, 0xc1, 0xa0, 0x7f, 0xcd, 0xba, 0x86,
	0xe3, 0x58, 0xfd, 0xb7, 0xba, 0x2c, 0x62, 0xff, 0x33, 0x3c, 0x16, 0xa3, 0xd0, 0xe7, 0xf0, 0x34,
	0xe3, 0xe7, 0x4c, 0x47, 0xd5, 0x92, 0xd5, 0x3f, 0x5d, 0x4f, 0x63, 0xf6, 0x01, 0x7d, 0x06, 0xca,
	0x27, 0x74, 0xd7, 0x32, 0x75, 0xac, 0xed, 0x30, 0xcf, 0x3a, 0x58, 0xf0, 0x6b, 0xdd, 0x71, 0xd4,
	0x5e, 0x52, 0xd9, 0xb3, 0xfc, 0x78, 0x32, 0xa2, 0xa3, 0x91, 0xe7, 0xf9, 0x0d, 0xf4, 0x75, 0xf5,
	0x6d, 0xee, 0x96, 0x1e, 0xad, 0xe1, 0xb0, 0x96, 0x17, 0xf9, 0x4e, 0xad, 0x2b, 0x13, 0x6f, 0x62,
	0x7e, 0x12, 0xad, 0xcf, 0xf0, 0x98, 0xff, 0x05, 0xa5, 0xb0, 0x85, 0xbc, 0x75, 0x7e, 0xde, 0x37,
	0x92, 0x0e, 0xda, 0xd9, 0x79, 0x9a, 0x65, 0x98, 0xcc, 0xb6, 0xac, 0xd7, 0xac, 0x6f, 0x38, 0xae,
	0xd4, 0xfb, 0x72, 0x0d, 0x87, 0x5a, 0x2f, 0x4f, 0xfe, 0x2a, 0xc0, 0xee, 0xca, 0x5f, 0x08, 0x14,
	0xa0, 0x6a, 0x5a, 0xcc, 0x34, 0x86, 0xa4, 0x40, 0x1b, 0x50, 0x41, 0x85, 0xaf, 0x48, 0x31, 0x33,
	0xcf, 0x48, 0x29, 0x33, 0xbf, 0x26, 0xe5, 0xcc, 0xfc, 0x86, 0x54, 0x32, 0xf3, 0x5b, 0x52, 0xcd,
	0xcc, 0xef, 0x48, 0x2d, 0x33, 0xbf, 0x27, 0xf5, 0xcc, 0xfc, 0x81, 0x34, 0x32, 0xf3, 0x47, 0x02,
	0xb4, 0x09, 0x35, 0x34, 0xf1, 0xbc, 0x26, 0x3a, 0xd7, 0x86, 0x29, 0x9d, 0x0d, 0x74, 0x2e, 0x53,
	0x67, 0x93, 0x6e, 0x43, 0xf3, 0x6a, 0xc8, 0x7e, 0x36, 0xd4, 0xa4, 0xb4, 0x2d, 0x04, 0x30, 0xef,
	0xdd, 0x85, 0xca, 0xba, 0xaa, 0x49, 0xb6, 0xb3, 0xe7, 0x79, 0x6d, 0x30, 0x17, 0xa3, 0xfa, 0x96,
	0xd9, 0x23, 0xe4, 0xe4, 0xb7, 0x22, 0xec, 0xe5, 0xfb, 0x94, 0xbf, 0x04, 0xe2, 0xa9, 0x37, 0xe5,
	0xd9, 0xb8, 0xe5, 0x82, 0x1c, 0x57, 0x75, 0x87, 0x0e, 0xbb, 0x52, 0x0d, 0x37, 0xf9, 0x4a, 0x90,
	0xc2, 0x5a, 0x5e, 0x3e, 0x7e, 0x52, 0x5c, 0xcb, 0xcb, 0xcf, 0x00, 0x29, 0xd1, 0x97, 0xd0, 0x5a,
	0xc9, 0xe7, 0x5e, 0x2d, 0x29, 0xd3, 0x17, 0x70, 0xb8, 0x32, 0x2a, 0x79, 0x78, 0xa4, 0xb2, 0xf6,
	0x18, 0xf9, 0x04, 0x49, 0x75, 0xad, 0x40, 0xf2, 0x38, 0x49, 0xed, 0xe4, 0x0e, 0xf6, 0xd7, 0xfc,
	0xdf, 0xc0, 0x21, 0x76, 0x0d, 0xb3, 0xc7, 0xd2, 0x6a, 0x0a, 0xf8, 0xf2, 0x9d, 0xa1, 0xc1, 0x2e,
	0x0d, 0xf6, 0x6e, 0x68, 0x65, 0x70, 0x91, 0x12, 0xd8, 0x90, 0x55, 0x67, 0x48, 0x09, 0xdf, 0xc6,
	0x79, 0xd2, 0x06, 0x43, 0x85, 0xac, 0x9d, 0xc1, 0xa3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd1,
	0xd3, 0x5c, 0xdd, 0x09, 0x0b, 0x00, 0x00,
}
