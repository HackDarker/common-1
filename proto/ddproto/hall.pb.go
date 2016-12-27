// Code generated by protoc-gen-go.
// source: hall.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of hall_lottery_item from hall_data.proto

// Ignoring public import of hall_sign_lottery_info from hall_data.proto

// Ignoring public import of hall_item_event from hall_data.proto

// Ignoring public import of hall_item_email from hall_data.proto

// Ignoring public import of hall_item_bag from hall_data.proto

// Ignoring public import of hall_item_task from hall_data.proto

// Ignoring public import of hall_userData from hall_data.proto

// Ignoring public import of hall_item_rank from hall_data.proto

// Ignoring public import of CoinZone from hall_data.proto

// Ignoring public import of DiamondZone from hall_data.proto

// Ignoring public import of ExchangeZone from hall_data.proto

// Ignoring public import of BuyZone from hall_data.proto

// Ignoring public import of GoodsItem from hall_data.proto

// Ignoring public import of hall_enum_protoId from hall_data.proto

// Ignoring public import of hall_enum_event from hall_data.proto

// Ignoring public import of hall_enum_Reward from hall_data.proto

// Ignoring public import of hall_enum_emailType from hall_data.proto

// Ignoring public import of hall_enum_bagItemType from hall_data.proto

// Ignoring public import of hall_enum_bagItemRank from hall_data.proto

// Ignoring public import of hall_enum_taskType from hall_data.proto

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 道具类型
type HallEnumPropsType int32

const (
	HallEnumPropsType_TYPE_LABA    HallEnumPropsType = 1
	HallEnumPropsType_TYPE_JIPAIQI HallEnumPropsType = 2
	HallEnumPropsType_TYPE_FANGKA  HallEnumPropsType = 3
)

var HallEnumPropsType_name = map[int32]string{
	1: "TYPE_LABA",
	2: "TYPE_JIPAIQI",
	3: "TYPE_FANGKA",
}
var HallEnumPropsType_value = map[string]int32{
	"TYPE_LABA":    1,
	"TYPE_JIPAIQI": 2,
	"TYPE_FANGKA":  3,
}

func (x HallEnumPropsType) Enum() *HallEnumPropsType {
	p := new(HallEnumPropsType)
	*p = x
	return p
}
func (x HallEnumPropsType) String() string {
	return proto.EnumName(HallEnumPropsType_name, int32(x))
}
func (x *HallEnumPropsType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumPropsType_value, data, "HallEnumPropsType")
	if err != nil {
		return err
	}
	*x = HallEnumPropsType(value)
	return nil
}
func (HallEnumPropsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

// ======================商城商品列表======================
// 商品列表页面类型
type HallEnumGoodsType int32

const (
	HallEnumGoodsType_TYPE_COIN     HallEnumGoodsType = 1
	HallEnumGoodsType_TYPE_DIAMOND  HallEnumGoodsType = 2
	HallEnumGoodsType_TYPE_EXCHANGE HallEnumGoodsType = 3
	HallEnumGoodsType_TYPE_BUY      HallEnumGoodsType = 4
	HallEnumGoodsType_TYPE_VIP      HallEnumGoodsType = 5
)

var HallEnumGoodsType_name = map[int32]string{
	1: "TYPE_COIN",
	2: "TYPE_DIAMOND",
	3: "TYPE_EXCHANGE",
	4: "TYPE_BUY",
	5: "TYPE_VIP",
}
var HallEnumGoodsType_value = map[string]int32{
	"TYPE_COIN":     1,
	"TYPE_DIAMOND":  2,
	"TYPE_EXCHANGE": 3,
	"TYPE_BUY":      4,
	"TYPE_VIP":      5,
}

func (x HallEnumGoodsType) Enum() *HallEnumGoodsType {
	p := new(HallEnumGoodsType)
	*p = x
	return p
}
func (x HallEnumGoodsType) String() string {
	return proto.EnumName(HallEnumGoodsType_name, int32(x))
}
func (x *HallEnumGoodsType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumGoodsType_value, data, "HallEnumGoodsType")
	if err != nil {
		return err
	}
	*x = HallEnumGoodsType(value)
	return nil
}
func (HallEnumGoodsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

// 货币与商品类型
type HallEnumTradeType int32

const (
	HallEnumTradeType_TRADE_COIN    HallEnumTradeType = 1
	HallEnumTradeType_TRADE_DIAMOND HallEnumTradeType = 2
	HallEnumTradeType_TRADE_BONUS   HallEnumTradeType = 3
	HallEnumTradeType_TRADE_CODE    HallEnumTradeType = 4
	HallEnumTradeType_TRADE_RMB     HallEnumTradeType = 5
	HallEnumTradeType_TRADE_PROPS   HallEnumTradeType = 6
	HallEnumTradeType_TRADE_VIP     HallEnumTradeType = 7
)

var HallEnumTradeType_name = map[int32]string{
	1: "TRADE_COIN",
	2: "TRADE_DIAMOND",
	3: "TRADE_BONUS",
	4: "TRADE_CODE",
	5: "TRADE_RMB",
	6: "TRADE_PROPS",
	7: "TRADE_VIP",
}
var HallEnumTradeType_value = map[string]int32{
	"TRADE_COIN":    1,
	"TRADE_DIAMOND": 2,
	"TRADE_BONUS":   3,
	"TRADE_CODE":    4,
	"TRADE_RMB":     5,
	"TRADE_PROPS":   6,
	"TRADE_VIP":     7,
}

func (x HallEnumTradeType) Enum() *HallEnumTradeType {
	p := new(HallEnumTradeType)
	*p = x
	return p
}
func (x HallEnumTradeType) String() string {
	return proto.EnumName(HallEnumTradeType_name, int32(x))
}
func (x *HallEnumTradeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumTradeType_value, data, "HallEnumTradeType")
	if err != nil {
		return err
	}
	*x = HallEnumTradeType(value)
	return nil
}
func (HallEnumTradeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

// 活动列表
type HallReqEvent struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqEvent) Reset()                    { *m = HallReqEvent{} }
func (m *HallReqEvent) String() string            { return proto.CompactTextString(m) }
func (*HallReqEvent) ProtoMessage()               {}
func (*HallReqEvent) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *HallReqEvent) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 活动列表ACK
type HallAckEvent struct {
	Header           *ProtoHeader     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	EventList        []*HallItemEvent `protobuf:"bytes,2,rep,name=eventList" json:"eventList,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallAckEvent) Reset()                    { *m = HallAckEvent{} }
func (m *HallAckEvent) String() string            { return proto.CompactTextString(m) }
func (*HallAckEvent) ProtoMessage()               {}
func (*HallAckEvent) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *HallAckEvent) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckEvent) GetEventList() []*HallItemEvent {
	if m != nil {
		return m.EventList
	}
	return nil
}

// 邮件
type HallReqMail struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqMail) Reset()                    { *m = HallReqMail{} }
func (m *HallReqMail) String() string            { return proto.CompactTextString(m) }
func (*HallReqMail) ProtoMessage()               {}
func (*HallReqMail) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *HallReqMail) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckMail struct {
	Header           *ProtoHeader     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	EmailList        []*HallItemEmail `protobuf:"bytes,2,rep,name=emailList" json:"emailList,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallAckMail) Reset()                    { *m = HallAckMail{} }
func (m *HallAckMail) String() string            { return proto.CompactTextString(m) }
func (*HallAckMail) ProtoMessage()               {}
func (*HallAckMail) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *HallAckMail) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckMail) GetEmailList() []*HallItemEmail {
	if m != nil {
		return m.EmailList
	}
	return nil
}

// 任务
type HallReqTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqTask) Reset()                    { *m = HallReqTask{} }
func (m *HallReqTask) String() string            { return proto.CompactTextString(m) }
func (*HallReqTask) ProtoMessage()               {}
func (*HallReqTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *HallReqTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckTask struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	TaskList         []*HallItemTask `protobuf:"bytes,2,rep,name=taskList" json:"taskList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HallAckTask) Reset()                    { *m = HallAckTask{} }
func (m *HallAckTask) String() string            { return proto.CompactTextString(m) }
func (*HallAckTask) ProtoMessage()               {}
func (*HallAckTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *HallAckTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckTask) GetTaskList() []*HallItemTask {
	if m != nil {
		return m.TaskList
	}
	return nil
}

// 领取任务
type HallReqGetTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqGetTask) Reset()                    { *m = HallReqGetTask{} }
func (m *HallReqGetTask) String() string            { return proto.CompactTextString(m) }
func (*HallReqGetTask) ProtoMessage()               {}
func (*HallReqGetTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *HallReqGetTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckGetTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	FinishedNum      *int32       `protobuf:"varint,2,opt,name=finishedNum" json:"finishedNum,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallAckGetTask) Reset()                    { *m = HallAckGetTask{} }
func (m *HallAckGetTask) String() string            { return proto.CompactTextString(m) }
func (*HallAckGetTask) ProtoMessage()               {}
func (*HallAckGetTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

func (m *HallAckGetTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckGetTask) GetFinishedNum() int32 {
	if m != nil && m.FinishedNum != nil {
		return *m.FinishedNum
	}
	return 0
}

// ==========================背包道具列表============================
// 背包道具列表
type HallReqBagItems struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqBagItems) Reset()                    { *m = HallReqBagItems{} }
func (m *HallReqBagItems) String() string            { return proto.CompactTextString(m) }
func (*HallReqBagItems) ProtoMessage()               {}
func (*HallReqBagItems) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *HallReqBagItems) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 道具列表ACK
type HallAckBagItems struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Items            []*HallBagItem `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *HallAckBagItems) Reset()                    { *m = HallAckBagItems{} }
func (m *HallAckBagItems) String() string            { return proto.CompactTextString(m) }
func (*HallAckBagItems) ProtoMessage()               {}
func (*HallAckBagItems) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{9} }

func (m *HallAckBagItems) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckBagItems) GetItems() []*HallBagItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// 单个道具
type HallBagItem struct {
	Type             *HallEnumPropsType `protobuf:"varint,1,opt,name=type,enum=ddproto.HallEnumPropsType" json:"type,omitempty"`
	Amount           *int32             `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallBagItem) Reset()                    { *m = HallBagItem{} }
func (m *HallBagItem) String() string            { return proto.CompactTextString(m) }
func (*HallBagItem) ProtoMessage()               {}
func (*HallBagItem) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{10} }

func (m *HallBagItem) GetType() HallEnumPropsType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumPropsType_TYPE_LABA
}

func (m *HallBagItem) GetAmount() int32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

// 个人信息
type HallReqUserData struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqUserData) Reset()                    { *m = HallReqUserData{} }
func (m *HallReqUserData) String() string            { return proto.CompactTextString(m) }
func (*HallReqUserData) ProtoMessage()               {}
func (*HallReqUserData) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{11} }

func (m *HallReqUserData) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckUserData struct {
	Header           *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	User             *HallUserData `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *HallAckUserData) Reset()                    { *m = HallAckUserData{} }
func (m *HallAckUserData) String() string            { return proto.CompactTextString(m) }
func (*HallAckUserData) ProtoMessage()               {}
func (*HallAckUserData) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{12} }

func (m *HallAckUserData) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckUserData) GetUser() *HallUserData {
	if m != nil {
		return m.User
	}
	return nil
}

// 商品列表Req
type HallReqGoodsList struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	GoodsType        *HallEnumGoodsType `protobuf:"varint,2,opt,name=goods_type,enum=ddproto.HallEnumGoodsType" json:"goods_type,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallReqGoodsList) Reset()                    { *m = HallReqGoodsList{} }
func (m *HallReqGoodsList) String() string            { return proto.CompactTextString(m) }
func (*HallReqGoodsList) ProtoMessage()               {}
func (*HallReqGoodsList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{13} }

func (m *HallReqGoodsList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallReqGoodsList) GetGoodsType() HallEnumGoodsType {
	if m != nil && m.GoodsType != nil {
		return *m.GoodsType
	}
	return HallEnumGoodsType_TYPE_COIN
}

// 商品列表Ack
type HallAckGoodsList struct {
	Header           *ProtoHeader        `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	GoodsType        *HallEnumGoodsType  `protobuf:"varint,2,opt,name=goods_type,enum=ddproto.HallEnumGoodsType" json:"goods_type,omitempty"`
	Items            []*HallGoodsItemMsg `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *HallAckGoodsList) Reset()                    { *m = HallAckGoodsList{} }
func (m *HallAckGoodsList) String() string            { return proto.CompactTextString(m) }
func (*HallAckGoodsList) ProtoMessage()               {}
func (*HallAckGoodsList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{14} }

func (m *HallAckGoodsList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckGoodsList) GetGoodsType() HallEnumGoodsType {
	if m != nil && m.GoodsType != nil {
		return *m.GoodsType
	}
	return HallEnumGoodsType_TYPE_COIN
}

func (m *HallAckGoodsList) GetItems() []*HallGoodsItemMsg {
	if m != nil {
		return m.Items
	}
	return nil
}

// 单个商品信息
type HallGoodsItemMsg struct {
	GoodsId          *int32             `protobuf:"varint,1,opt,name=goods_id" json:"goods_id,omitempty"`
	Name             *string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PriceType        *HallEnumTradeType `protobuf:"varint,3,opt,name=price_type,enum=ddproto.HallEnumTradeType" json:"price_type,omitempty"`
	Price            *int32             `protobuf:"varint,4,opt,name=price" json:"price,omitempty"`
	ValueType        *HallEnumTradeType `protobuf:"varint,5,opt,name=value_type,enum=ddproto.HallEnumTradeType" json:"value_type,omitempty"`
	Value            *int32             `protobuf:"varint,6,opt,name=value" json:"value,omitempty"`
	Discount         *string            `protobuf:"bytes,7,opt,name=discount" json:"discount,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallGoodsItemMsg) Reset()                    { *m = HallGoodsItemMsg{} }
func (m *HallGoodsItemMsg) String() string            { return proto.CompactTextString(m) }
func (*HallGoodsItemMsg) ProtoMessage()               {}
func (*HallGoodsItemMsg) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{15} }

func (m *HallGoodsItemMsg) GetGoodsId() int32 {
	if m != nil && m.GoodsId != nil {
		return *m.GoodsId
	}
	return 0
}

func (m *HallGoodsItemMsg) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *HallGoodsItemMsg) GetPriceType() HallEnumTradeType {
	if m != nil && m.PriceType != nil {
		return *m.PriceType
	}
	return HallEnumTradeType_TRADE_COIN
}

func (m *HallGoodsItemMsg) GetPrice() int32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *HallGoodsItemMsg) GetValueType() HallEnumTradeType {
	if m != nil && m.ValueType != nil {
		return *m.ValueType
	}
	return HallEnumTradeType_TRADE_COIN
}

func (m *HallGoodsItemMsg) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *HallGoodsItemMsg) GetDiscount() string {
	if m != nil && m.Discount != nil {
		return *m.Discount
	}
	return ""
}

// 排行
type HallReqRank struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqRank) Reset()                    { *m = HallReqRank{} }
func (m *HallReqRank) String() string            { return proto.CompactTextString(m) }
func (*HallReqRank) ProtoMessage()               {}
func (*HallReqRank) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{16} }

func (m *HallReqRank) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckRank struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RankList         []*HallItemRank `protobuf:"bytes,2,rep,name=rankList" json:"rankList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HallAckRank) Reset()                    { *m = HallAckRank{} }
func (m *HallAckRank) String() string            { return proto.CompactTextString(m) }
func (*HallAckRank) ProtoMessage()               {}
func (*HallAckRank) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{17} }

func (m *HallAckRank) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckRank) GetRankList() []*HallItemRank {
	if m != nil {
		return m.RankList
	}
	return nil
}

// 转盘签到奖励物品信息
type HallAckDsLotteryItems struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DialItems        []*HallLotteryItem   `protobuf:"bytes,2,rep,name=dialItems" json:"dialItems,omitempty"`
	SignInfo         *HallSignLotteryInfo `protobuf:"bytes,3,opt,name=signInfo" json:"signInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *HallAckDsLotteryItems) Reset()                    { *m = HallAckDsLotteryItems{} }
func (m *HallAckDsLotteryItems) String() string            { return proto.CompactTextString(m) }
func (*HallAckDsLotteryItems) ProtoMessage()               {}
func (*HallAckDsLotteryItems) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{18} }

func (m *HallAckDsLotteryItems) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckDsLotteryItems) GetDialItems() []*HallLotteryItem {
	if m != nil {
		return m.DialItems
	}
	return nil
}

func (m *HallAckDsLotteryItems) GetSignInfo() *HallSignLotteryInfo {
	if m != nil {
		return m.SignInfo
	}
	return nil
}

// 签到动作
type HallReqSignLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqSignLottery) Reset()                    { *m = HallReqSignLottery{} }
func (m *HallReqSignLottery) String() string            { return proto.CompactTextString(m) }
func (*HallReqSignLottery) ProtoMessage()               {}
func (*HallReqSignLottery) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{19} }

func (m *HallReqSignLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckSignLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	LotteryId        *int32       `protobuf:"varint,2,opt,name=lotteryId" json:"lotteryId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallAckSignLottery) Reset()                    { *m = HallAckSignLottery{} }
func (m *HallAckSignLottery) String() string            { return proto.CompactTextString(m) }
func (*HallAckSignLottery) ProtoMessage()               {}
func (*HallAckSignLottery) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{20} }

func (m *HallAckSignLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckSignLottery) GetLotteryId() int32 {
	if m != nil && m.LotteryId != nil {
		return *m.LotteryId
	}
	return 0
}

// 抽奖动作
type HallReqDrawLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqDrawLottery) Reset()                    { *m = HallReqDrawLottery{} }
func (m *HallReqDrawLottery) String() string            { return proto.CompactTextString(m) }
func (*HallReqDrawLottery) ProtoMessage()               {}
func (*HallReqDrawLottery) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{21} }

func (m *HallReqDrawLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckDrawLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	LotteryId        *int32       `protobuf:"varint,2,opt,name=lotteryId" json:"lotteryId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallAckDrawLottery) Reset()                    { *m = HallAckDrawLottery{} }
func (m *HallAckDrawLottery) String() string            { return proto.CompactTextString(m) }
func (*HallAckDrawLottery) ProtoMessage()               {}
func (*HallAckDrawLottery) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{22} }

func (m *HallAckDrawLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckDrawLottery) GetLotteryId() int32 {
	if m != nil && m.LotteryId != nil {
		return *m.LotteryId
	}
	return 0
}

func init() {
	proto.RegisterType((*HallReqEvent)(nil), "ddproto.hall_req_event")
	proto.RegisterType((*HallAckEvent)(nil), "ddproto.hall_ack_event")
	proto.RegisterType((*HallReqMail)(nil), "ddproto.hall_req_mail")
	proto.RegisterType((*HallAckMail)(nil), "ddproto.hall_ack_mail")
	proto.RegisterType((*HallReqTask)(nil), "ddproto.hall_req_task")
	proto.RegisterType((*HallAckTask)(nil), "ddproto.hall_ack_task")
	proto.RegisterType((*HallReqGetTask)(nil), "ddproto.hall_req_getTask")
	proto.RegisterType((*HallAckGetTask)(nil), "ddproto.hall_ack_getTask")
	proto.RegisterType((*HallReqBagItems)(nil), "ddproto.hall_req_bag_items")
	proto.RegisterType((*HallAckBagItems)(nil), "ddproto.hall_ack_bag_items")
	proto.RegisterType((*HallBagItem)(nil), "ddproto.hall_bag_item")
	proto.RegisterType((*HallReqUserData)(nil), "ddproto.hall_req_userData")
	proto.RegisterType((*HallAckUserData)(nil), "ddproto.hall_ack_userData")
	proto.RegisterType((*HallReqGoodsList)(nil), "ddproto.hall_req_goods_list")
	proto.RegisterType((*HallAckGoodsList)(nil), "ddproto.hall_ack_goods_list")
	proto.RegisterType((*HallGoodsItemMsg)(nil), "ddproto.hall_goods_item_msg")
	proto.RegisterType((*HallReqRank)(nil), "ddproto.hall_req_rank")
	proto.RegisterType((*HallAckRank)(nil), "ddproto.hall_ack_rank")
	proto.RegisterType((*HallAckDsLotteryItems)(nil), "ddproto.hall_ack_ds_lottery_items")
	proto.RegisterType((*HallReqSignLottery)(nil), "ddproto.hall_req_sign_lottery")
	proto.RegisterType((*HallAckSignLottery)(nil), "ddproto.hall_ack_sign_lottery")
	proto.RegisterType((*HallReqDrawLottery)(nil), "ddproto.hall_req_draw_lottery")
	proto.RegisterType((*HallAckDrawLottery)(nil), "ddproto.hall_ack_draw_lottery")
	proto.RegisterEnum("ddproto.HallEnumPropsType", HallEnumPropsType_name, HallEnumPropsType_value)
	proto.RegisterEnum("ddproto.HallEnumGoodsType", HallEnumGoodsType_name, HallEnumGoodsType_value)
	proto.RegisterEnum("ddproto.HallEnumTradeType", HallEnumTradeType_name, HallEnumTradeType_value)
}

var fileDescriptor13 = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0xdb, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0x7f, 0x37, 0x87, 0x26, 0x3b, 0x4d, 0xea, 0xba, 0xfd, 0x21, 0x54, 0x20, 0x2a, 0x0b,
	0xa4, 0x52, 0x44, 0x25, 0x2a, 0x81, 0x00, 0x89, 0x0b, 0xa7, 0x09, 0xd4, 0xd0, 0x26, 0xa6, 0x07,
	0x44, 0xaf, 0xcc, 0x34, 0x9e, 0xa6, 0x56, 0x7d, 0x08, 0xb6, 0x53, 0xd4, 0x47, 0xe0, 0x19, 0xb8,
	0xe4, 0x71, 0x78, 0x29, 0x66, 0xb6, 0x27, 0x8e, 0x2d, 0x5a, 0xd5, 0x06, 0x71, 0x13, 0xc5, 0x7b,
	0xd6, 0x5a, 0xfb, 0x9b, 0x23, 0xc0, 0x19, 0x71, 0x9c, 0xcd, 0x71, 0xe0, 0x47, 0xbe, 0x32, 0x6f,
	0x59, 0xf8, 0x67, 0x75, 0x91, 0x17, 0x4d, 0x8b, 0x44, 0x24, 0x1e, 0x59, 0x5d, 0x1e, 0xfa, 0xae,
	0xeb, 0x7b, 0xe6, 0xd0, 0xb1, 0xa9, 0x17, 0xc5, 0x45, 0xf5, 0x39, 0xb4, 0x50, 0x17, 0xd0, 0x2f,
	0x26, 0xbd, 0x60, 0x75, 0xe5, 0x01, 0x54, 0xcf, 0x28, 0xb1, 0x68, 0xd0, 0x96, 0xd6, 0xa4, 0xf5,
	0xc6, 0xd6, 0xca, 0xa6, 0x48, 0xdc, 0x34, 0xf8, 0xef, 0x0e, 0x8e, 0xa9, 0x43, 0xe1, 0x23, 0xc3,
	0xf3, 0x22, 0x3e, 0xe5, 0x31, 0xd4, 0x51, 0xbe, 0x6b, 0x87, 0x51, 0x7b, 0x6e, 0xad, 0xc4, 0x84,
	0xed, 0x44, 0x88, 0x89, 0x76, 0x44, 0xdd, 0x38, 0x52, 0x7d, 0x06, 0xcd, 0x04, 0xce, 0x25, 0xb6,
	0x93, 0x93, 0xed, 0x44, 0xd8, 0x38, 0x5b, 0x7e, 0x1b, 0xa2, 0x71, 0xf9, 0x4d, 0x68, 0x5c, 0x93,
	0x41, 0x8b, 0x48, 0x78, 0x9e, 0x13, 0xed, 0x73, 0x0a, 0x2d, 0xbf, 0x4d, 0x79, 0x04, 0x35, 0xae,
	0x4e, 0x91, 0xdd, 0xbe, 0x82, 0x8c, 0x4b, 0xd4, 0x17, 0x20, 0x27, 0x60, 0x23, 0x1a, 0x1d, 0xe6,
	0x67, 0xdb, 0x13, 0x4e, 0xce, 0x56, 0xc8, 0xa9, 0x2c, 0x43, 0xe3, 0xd4, 0xf6, 0xec, 0xf0, 0x8c,
	0x5a, 0xfd, 0x89, 0xcb, 0x08, 0xa5, 0xf5, 0x8a, 0xfa, 0x0a, 0x94, 0x04, 0xe4, 0x84, 0x8c, 0x10,
	0x31, 0xcc, 0x89, 0x42, 0x84, 0x97, 0xa3, 0x14, 0xf4, 0x2a, 0x0f, 0xa1, 0x82, 0x72, 0xb1, 0x50,
	0xb7, 0xb2, 0x0b, 0x35, 0x4d, 0x53, 0x77, 0xc5, 0x4e, 0x4c, 0x0b, 0x6c, 0xfb, 0xcb, 0xd1, 0xe5,
	0x98, 0x62, 0x76, 0x6b, 0xeb, 0x5e, 0xd6, 0x46, 0xbd, 0x89, 0x6b, 0xb2, 0xef, 0x71, 0x68, 0x72,
	0x91, 0xd2, 0x82, 0x2a, 0x71, 0xfd, 0x89, 0x17, 0x89, 0xc9, 0xbe, 0x84, 0xa5, 0x64, 0xb2, 0x93,
	0x90, 0x06, 0x5d, 0x76, 0xed, 0x72, 0xce, 0xd5, 0x14, 0x56, 0x3e, 0xd7, 0x62, 0x56, 0xa6, 0x2a,
	0x73, 0x07, 0x32, 0xfc, 0x36, 0xd3, 0x69, 0x96, 0xea, 0xc1, 0xf2, 0xec, 0x44, 0xf8, 0xbe, 0x15,
	0x9a, 0x0e, 0x3b, 0x47, 0x39, 0x5b, 0x3c, 0x05, 0x88, 0x3d, 0xb8, 0x36, 0x73, 0xd7, 0xae, 0xcd,
	0x4c, 0xa4, 0x7e, 0x97, 0x44, 0x43, 0x3c, 0x48, 0xff, 0xbe, 0x21, 0xdb, 0x39, 0xb1, 0xe3, 0x25,
	0xdc, 0xf1, 0xbb, 0x59, 0x75, 0x2c, 0xc4, 0x0b, 0xe2, 0x86, 0x23, 0xf5, 0xe7, 0x94, 0x2e, 0x5b,
	0x57, 0x64, 0xa8, 0x89, 0x8a, 0x85, 0x7c, 0x15, 0x65, 0x01, 0xca, 0x1e, 0x71, 0x63, 0x86, 0x3a,
	0xe7, 0x1a, 0x07, 0xf6, 0x90, 0xc6, 0x5c, 0xa5, 0x6b, 0xb9, 0xa2, 0x80, 0x4d, 0x23, 0xe6, 0x6a,
	0x42, 0x05, 0x2d, 0xed, 0x32, 0xe6, 0xb1, 0x84, 0x0b, 0xe2, 0x4c, 0x44, 0x42, 0x25, 0x67, 0x02,
	0x5a, 0xda, 0x55, 0x4c, 0x60, 0x8c, 0x96, 0x1d, 0x0e, 0xf1, 0xdc, 0xcd, 0x73, 0xaa, 0xcc, 0x33,
	0x14, 0x10, 0xef, 0x4f, 0x9e, 0xa1, 0xfc, 0x36, 0xfe, 0x0c, 0x71, 0xf5, 0x0d, 0xcf, 0x10, 0x97,
	0xa8, 0x3f, 0x24, 0xb8, 0x93, 0xb4, 0xe0, 0x27, 0xc0, 0x8f, 0x22, 0x1a, 0x5c, 0x16, 0xba, 0xc9,
	0x4f, 0xa0, 0x6e, 0xd9, 0xc4, 0xd1, 0x53, 0xb7, 0x79, 0x35, 0xdb, 0x2f, 0x9d, 0xca, 0xd6, 0xb7,
	0x16, 0xda, 0x23, 0x4f, 0xf7, 0x4e, 0x7d, 0xdc, 0x9f, 0xc6, 0xd6, 0xfd, 0xac, 0x9a, 0x8f, 0xce,
	0x2c, 0x4c, 0xa6, 0xbe, 0x86, 0xff, 0x93, 0xe5, 0x4b, 0x8f, 0xe6, 0x5c, 0x46, 0x43, 0xd8, 0xf9,
	0x1c, 0x8b, 0xdb, 0x95, 0x25, 0xa8, 0x0b, 0x83, 0x6e, 0x89, 0x77, 0x24, 0x0d, 0x64, 0x05, 0xe4,
	0xeb, 0x5f, 0x00, 0x15, 0xb7, 0x5f, 0x01, 0xb4, 0xb1, 0x03, 0x2b, 0x57, 0x3e, 0x80, 0x4d, 0xa8,
	0x1f, 0x1e, 0x1b, 0x3d, 0x73, 0x57, 0xeb, 0x68, 0xb2, 0xc4, 0x4e, 0xe6, 0x02, 0x7e, 0xbe, 0xd3,
	0x0d, 0x4d, 0xff, 0xa0, 0xcb, 0x73, 0xca, 0x22, 0x34, 0xb0, 0xf2, 0x46, 0xeb, 0xbf, 0x7d, 0xaf,
	0xc9, 0xa5, 0x8d, 0xd3, 0x74, 0x52, 0xea, 0xf6, 0x4e, 0x93, 0xb6, 0x07, 0x7a, 0x3f, 0x95, 0xd4,
	0xd5, 0xb5, 0xbd, 0x41, 0xbf, 0xcb, 0x92, 0x96, 0xa0, 0x89, 0x95, 0xde, 0xa7, 0xed, 0x1d, 0x16,
	0xd6, 0x93, 0x4b, 0xec, 0x6a, 0xd6, 0xb0, 0xd4, 0x39, 0x3a, 0x96, 0xcb, 0xc9, 0xd7, 0x47, 0xdd,
	0x90, 0x2b, 0x1b, 0xdf, 0xa4, 0x74, 0xa3, 0xd4, 0x65, 0x6a, 0x01, 0x1c, 0xee, 0x6b, 0xdd, 0xa4,
	0x13, 0xcf, 0xc5, 0xef, 0x59, 0x2b, 0x0e, 0x8d, 0xa5, 0xce, 0xa0, 0x7f, 0x74, 0xc0, 0x1a, 0xa5,
	0x3c, 0xdd, 0x1e, 0x6b, 0xc5, 0x61, 0xf1, 0x7b, 0x7f, 0xaf, 0x23, 0x57, 0x66, 0x7a, 0x63, 0x7f,
	0x60, 0x1c, 0xc8, 0xd5, 0xd9, 0x38, 0x67, 0x99, 0x37, 0xfe, 0x33, 0xa4, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x92, 0x28, 0xb1, 0x58, 0x9d, 0x09, 0x00, 0x00,
}
