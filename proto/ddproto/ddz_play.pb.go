// Code generated by protoc-gen-go.
// source: ddz_play.proto
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

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

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

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of ddz_base_roomTypeInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_commonRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_timerInfo from ddz_base.proto

// Ignoring public import of ddz_base_deskInfo from ddz_base.proto

// Ignoring public import of ddz_enum_protoId from ddz_base.proto

// Ignoring public import of ddz_enum_errorCode from ddz_base.proto

// Ignoring public import of ddz_enum_paiType from ddz_base.proto

// Ignoring public import of ddz_enum_actType from ddz_base.proto

// Ignoring public import of ddz_enum_gameStatus from ddz_base.proto

// Ignoring public import of ddz_enum_playerStatus from ddz_base.proto

// Ignoring public import of ddz_enum_roomType from ddz_base.proto

// Ignoring public import of ddz_enum_enterType from ddz_base.proto

// Ignoring public import of ddz_enum_coinRoomLevel from ddz_base.proto

// Ignoring public import of ddz_enum_deskGameStatus from ddz_base.proto

type DdzEnumJdScore int32

const (
	DdzEnumJdScore_DDZ_ONE   DdzEnumJdScore = 1
	DdzEnumJdScore_DDZ_TWO   DdzEnumJdScore = 2
	DdzEnumJdScore_DDZ_THREE DdzEnumJdScore = 3
)

var DdzEnumJdScore_name = map[int32]string{
	1: "DDZ_ONE",
	2: "DDZ_TWO",
	3: "DDZ_THREE",
}
var DdzEnumJdScore_value = map[string]int32{
	"DDZ_ONE":   1,
	"DDZ_TWO":   2,
	"DDZ_THREE": 3,
}

func (x DdzEnumJdScore) Enum() *DdzEnumJdScore {
	p := new(DdzEnumJdScore)
	*p = x
	return p
}
func (x DdzEnumJdScore) String() string {
	return proto.EnumName(DdzEnumJdScore_name, int32(x))
}
func (x *DdzEnumJdScore) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DdzEnumJdScore_value, data, "DdzEnumJdScore")
	if err != nil {
		return err
	}
	*x = DdzEnumJdScore(value)
	return nil
}
func (DdzEnumJdScore) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

type DdzEnumDoubleType int32

const (
	DdzEnumDoubleType_DDZ_JIA   DdzEnumDoubleType = 1
	DdzEnumDoubleType_DDZ_BUJIA DdzEnumDoubleType = 2
)

var DdzEnumDoubleType_name = map[int32]string{
	1: "DDZ_JIA",
	2: "DDZ_BUJIA",
}
var DdzEnumDoubleType_value = map[string]int32{
	"DDZ_JIA":   1,
	"DDZ_BUJIA": 2,
}

func (x DdzEnumDoubleType) Enum() *DdzEnumDoubleType {
	p := new(DdzEnumDoubleType)
	*p = x
	return p
}
func (x DdzEnumDoubleType) String() string {
	return proto.EnumName(DdzEnumDoubleType_name, int32(x))
}
func (x *DdzEnumDoubleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DdzEnumDoubleType_value, data, "DdzEnumDoubleType")
	if err != nil {
		return err
	}
	*x = DdzEnumDoubleType(value)
	return nil
}
func (DdzEnumDoubleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

type DdzEnumLaDaoType int32

const (
	DdzEnumLaDaoType_DDZ_LA      DdzEnumLaDaoType = 1
	DdzEnumLaDaoType_DDZ_DAO     DdzEnumLaDaoType = 2
	DdzEnumLaDaoType_DDZ_BULADAO DdzEnumLaDaoType = 3
)

var DdzEnumLaDaoType_name = map[int32]string{
	1: "DDZ_LA",
	2: "DDZ_DAO",
	3: "DDZ_BULADAO",
}
var DdzEnumLaDaoType_value = map[string]int32{
	"DDZ_LA":      1,
	"DDZ_DAO":     2,
	"DDZ_BULADAO": 3,
}

func (x DdzEnumLaDaoType) Enum() *DdzEnumLaDaoType {
	p := new(DdzEnumLaDaoType)
	*p = x
	return p
}
func (x DdzEnumLaDaoType) String() string {
	return proto.EnumName(DdzEnumLaDaoType_name, int32(x))
}
func (x *DdzEnumLaDaoType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DdzEnumLaDaoType_value, data, "DdzEnumLaDaoType")
	if err != nil {
		return err
	}
	*x = DdzEnumLaDaoType(value)
	return nil
}
func (DdzEnumLaDaoType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

// 开局（接收服务端消息）
type DdzBcOpening struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzBcOpening) Reset()                    { *m = DdzBcOpening{} }
func (m *DdzBcOpening) String() string            { return proto.CompactTextString(m) }
func (*DdzBcOpening) ProtoMessage()               {}
func (*DdzBcOpening) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *DdzBcOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 发牌
type DdzBcDealCards struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerPokers     []*ClientBasePoker `protobuf:"bytes,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzBcDealCards) Reset()                    { *m = DdzBcDealCards{} }
func (m *DdzBcDealCards) String() string            { return proto.CompactTextString(m) }
func (*DdzBcDealCards) ProtoMessage()               {}
func (*DdzBcDealCards) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *DdzBcDealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBcDealCards) GetPlayerPokers() []*ClientBasePoker {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

// 明牌
type DdzReqShowHandPokers struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Rate             *int32       `protobuf:"varint,3,opt,name=rate" json:"rate,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqShowHandPokers) Reset()                    { *m = DdzReqShowHandPokers{} }
func (m *DdzReqShowHandPokers) String() string            { return proto.CompactTextString(m) }
func (*DdzReqShowHandPokers) ProtoMessage()               {}
func (*DdzReqShowHandPokers) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *DdzReqShowHandPokers) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqShowHandPokers) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzReqShowHandPokers) GetRate() int32 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

// 明牌ack
type DdzAckShowHandPokers struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	PlayerPokers     []*ClientBasePoker `protobuf:"bytes,3,rep,name=playerPokers" json:"playerPokers,omitempty"`
	PlayerRate       *int32             `protobuf:"varint,4,opt,name=playerRate" json:"playerRate,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzAckShowHandPokers) Reset()                    { *m = DdzAckShowHandPokers{} }
func (m *DdzAckShowHandPokers) String() string            { return proto.CompactTextString(m) }
func (*DdzAckShowHandPokers) ProtoMessage()               {}
func (*DdzAckShowHandPokers) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *DdzAckShowHandPokers) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckShowHandPokers) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckShowHandPokers) GetPlayerPokers() []*ClientBasePoker {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

func (m *DdzAckShowHandPokers) GetPlayerRate() int32 {
	if m != nil && m.PlayerRate != nil {
		return *m.PlayerRate
	}
	return 0
}

// 叫地主
type DdzReqJiaoDiZhu struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32         `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Jiao             *bool           `protobuf:"varint,3,opt,name=jiao" json:"jiao,omitempty"`
	Score            *DdzEnumJdScore `protobuf:"varint,4,opt,name=score,enum=ddproto.DdzEnumJdScore" json:"score,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *DdzReqJiaoDiZhu) Reset()                    { *m = DdzReqJiaoDiZhu{} }
func (m *DdzReqJiaoDiZhu) String() string            { return proto.CompactTextString(m) }
func (*DdzReqJiaoDiZhu) ProtoMessage()               {}
func (*DdzReqJiaoDiZhu) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *DdzReqJiaoDiZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqJiaoDiZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzReqJiaoDiZhu) GetJiao() bool {
	if m != nil && m.Jiao != nil {
		return *m.Jiao
	}
	return false
}

func (m *DdzReqJiaoDiZhu) GetScore() DdzEnumJdScore {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return DdzEnumJdScore_DDZ_ONE
}

// 叫地主回复
type DdzAckJiaoDiZhu struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32         `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Jiao             *bool           `protobuf:"varint,3,opt,name=jiao" json:"jiao,omitempty"`
	Score            *DdzEnumJdScore `protobuf:"varint,4,opt,name=score,enum=ddproto.DdzEnumJdScore" json:"score,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *DdzAckJiaoDiZhu) Reset()                    { *m = DdzAckJiaoDiZhu{} }
func (m *DdzAckJiaoDiZhu) String() string            { return proto.CompactTextString(m) }
func (*DdzAckJiaoDiZhu) ProtoMessage()               {}
func (*DdzAckJiaoDiZhu) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *DdzAckJiaoDiZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckJiaoDiZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckJiaoDiZhu) GetJiao() bool {
	if m != nil && m.Jiao != nil {
		return *m.Jiao
	}
	return false
}

func (m *DdzAckJiaoDiZhu) GetScore() DdzEnumJdScore {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return DdzEnumJdScore_DDZ_ONE
}

// 抢地主
type DdzReqRobDiZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Rob              *bool        `protobuf:"varint,3,opt,name=rob" json:"rob,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqRobDiZhu) Reset()                    { *m = DdzReqRobDiZhu{} }
func (m *DdzReqRobDiZhu) String() string            { return proto.CompactTextString(m) }
func (*DdzReqRobDiZhu) ProtoMessage()               {}
func (*DdzReqRobDiZhu) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *DdzReqRobDiZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqRobDiZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzReqRobDiZhu) GetRob() bool {
	if m != nil && m.Rob != nil {
		return *m.Rob
	}
	return false
}

type DdzAckRobDiZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Rob              *bool        `protobuf:"varint,3,opt,name=rob" json:"rob,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzAckRobDiZhu) Reset()                    { *m = DdzAckRobDiZhu{} }
func (m *DdzAckRobDiZhu) String() string            { return proto.CompactTextString(m) }
func (*DdzAckRobDiZhu) ProtoMessage()               {}
func (*DdzAckRobDiZhu) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

func (m *DdzAckRobDiZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckRobDiZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckRobDiZhu) GetRob() bool {
	if m != nil && m.Rob != nil {
		return *m.Rob
	}
	return false
}

// 加倍
type DdzReqDouble struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Double           *DdzEnumDoubleType `protobuf:"varint,3,opt,name=double,enum=ddproto.DdzEnumDoubleType" json:"double,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzReqDouble) Reset()                    { *m = DdzReqDouble{} }
func (m *DdzReqDouble) String() string            { return proto.CompactTextString(m) }
func (*DdzReqDouble) ProtoMessage()               {}
func (*DdzReqDouble) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *DdzReqDouble) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqDouble) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzReqDouble) GetDouble() DdzEnumDoubleType {
	if m != nil && m.Double != nil {
		return *m.Double
	}
	return DdzEnumDoubleType_DDZ_JIA
}

type DdzAckDouble struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Double           *DdzEnumDoubleType `protobuf:"varint,3,opt,name=double,enum=ddproto.DdzEnumDoubleType" json:"double,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzAckDouble) Reset()                    { *m = DdzAckDouble{} }
func (m *DdzAckDouble) String() string            { return proto.CompactTextString(m) }
func (*DdzAckDouble) ProtoMessage()               {}
func (*DdzAckDouble) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{9} }

func (m *DdzAckDouble) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckDouble) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckDouble) GetDouble() DdzEnumDoubleType {
	if m != nil && m.Double != nil {
		return *m.Double
	}
	return DdzEnumDoubleType_DDZ_JIA
}

// 叫地主结束，开始游戏 (广播)
type DdzBcStartPlay struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	FootPokers       []*ClientBasePoker `protobuf:"bytes,2,rep,name=footPokers" json:"footPokers,omitempty"`
	FootRate         *int32             `protobuf:"varint,3,opt,name=footRate" json:"footRate,omitempty"`
	Dizhu            *uint32            `protobuf:"varint,4,opt,name=dizhu" json:"dizhu,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzBcStartPlay) Reset()                    { *m = DdzBcStartPlay{} }
func (m *DdzBcStartPlay) String() string            { return proto.CompactTextString(m) }
func (*DdzBcStartPlay) ProtoMessage()               {}
func (*DdzBcStartPlay) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{10} }

func (m *DdzBcStartPlay) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBcStartPlay) GetFootPokers() []*ClientBasePoker {
	if m != nil {
		return m.FootPokers
	}
	return nil
}

func (m *DdzBcStartPlay) GetFootRate() int32 {
	if m != nil && m.FootRate != nil {
		return *m.FootRate
	}
	return 0
}

func (m *DdzBcStartPlay) GetDizhu() uint32 {
	if m != nil && m.Dizhu != nil {
		return *m.Dizhu
	}
	return 0
}

// 闷抓
type DdzReqMenuZhua struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqMenuZhua) Reset()                    { *m = DdzReqMenuZhua{} }
func (m *DdzReqMenuZhua) String() string            { return proto.CompactTextString(m) }
func (*DdzReqMenuZhua) ProtoMessage()               {}
func (*DdzReqMenuZhua) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{11} }

func (m *DdzReqMenuZhua) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqMenuZhua) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type DdzAckMenuZhua struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	HandPokers       []*ClientBasePoker `protobuf:"bytes,3,rep,name=handPokers" json:"handPokers,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzAckMenuZhua) Reset()                    { *m = DdzAckMenuZhua{} }
func (m *DdzAckMenuZhua) String() string            { return proto.CompactTextString(m) }
func (*DdzAckMenuZhua) ProtoMessage()               {}
func (*DdzAckMenuZhua) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{12} }

func (m *DdzAckMenuZhua) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckMenuZhua) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckMenuZhua) GetHandPokers() []*ClientBasePoker {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

// 看牌
type DdzReqSeeCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqSeeCards) Reset()                    { *m = DdzReqSeeCards{} }
func (m *DdzReqSeeCards) String() string            { return proto.CompactTextString(m) }
func (*DdzReqSeeCards) ProtoMessage()               {}
func (*DdzReqSeeCards) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{13} }

func (m *DdzReqSeeCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqSeeCards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type DdzAckSeeCards struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	HandPokers       []*ClientBasePoker `protobuf:"bytes,3,rep,name=handPokers" json:"handPokers,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzAckSeeCards) Reset()                    { *m = DdzAckSeeCards{} }
func (m *DdzAckSeeCards) String() string            { return proto.CompactTextString(m) }
func (*DdzAckSeeCards) ProtoMessage()               {}
func (*DdzAckSeeCards) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{14} }

func (m *DdzAckSeeCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckSeeCards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckSeeCards) GetHandPokers() []*ClientBasePoker {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

// 拉/倒
type DdzReqPull struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32           `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Act              *DdzEnumLaDaoType `protobuf:"varint,3,opt,name=act,enum=ddproto.DdzEnumLaDaoType" json:"act,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *DdzReqPull) Reset()                    { *m = DdzReqPull{} }
func (m *DdzReqPull) String() string            { return proto.CompactTextString(m) }
func (*DdzReqPull) ProtoMessage()               {}
func (*DdzReqPull) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{15} }

func (m *DdzReqPull) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqPull) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzReqPull) GetAct() DdzEnumLaDaoType {
	if m != nil && m.Act != nil {
		return *m.Act
	}
	return DdzEnumLaDaoType_DDZ_LA
}

type DdzAckPull struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32           `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Act              *DdzEnumLaDaoType `protobuf:"varint,3,opt,name=act,enum=ddproto.DdzEnumLaDaoType" json:"act,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *DdzAckPull) Reset()                    { *m = DdzAckPull{} }
func (m *DdzAckPull) String() string            { return proto.CompactTextString(m) }
func (*DdzAckPull) ProtoMessage()               {}
func (*DdzAckPull) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{16} }

func (m *DdzAckPull) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckPull) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckPull) GetAct() DdzEnumLaDaoType {
	if m != nil && m.Act != nil {
		return *m.Act
	}
	return DdzEnumLaDaoType_DDZ_LA
}

// 出牌
type DdzReqOutCards struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	OutCards         []*ClientBasePoker `protobuf:"bytes,2,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzReqOutCards) Reset()                    { *m = DdzReqOutCards{} }
func (m *DdzReqOutCards) String() string            { return proto.CompactTextString(m) }
func (*DdzReqOutCards) ProtoMessage()               {}
func (*DdzReqOutCards) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{17} }

func (m *DdzReqOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqOutCards) GetOutCards() []*ClientBasePoker {
	if m != nil {
		return m.OutCards
	}
	return nil
}

type DdzAckOutCards struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	CardType         *DdzEnumPaiType    `protobuf:"varint,3,opt,name=cardType,enum=ddproto.DdzEnumPaiType" json:"cardType,omitempty"`
	RemainPokers     *int32             `protobuf:"varint,4,opt,name=remainPokers" json:"remainPokers,omitempty"`
	OutCards         []*ClientBasePoker `protobuf:"bytes,5,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DdzAckOutCards) Reset()                    { *m = DdzAckOutCards{} }
func (m *DdzAckOutCards) String() string            { return proto.CompactTextString(m) }
func (*DdzAckOutCards) ProtoMessage()               {}
func (*DdzAckOutCards) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{18} }

func (m *DdzAckOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckOutCards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckOutCards) GetCardType() DdzEnumPaiType {
	if m != nil && m.CardType != nil {
		return *m.CardType
	}
	return DdzEnumPaiType_ERRORCARD
}

func (m *DdzAckOutCards) GetRemainPokers() int32 {
	if m != nil && m.RemainPokers != nil {
		return *m.RemainPokers
	}
	return 0
}

func (m *DdzAckOutCards) GetOutCards() []*ClientBasePoker {
	if m != nil {
		return m.OutCards
	}
	return nil
}

// 过牌
type DdzReqActGuo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqActGuo) Reset()                    { *m = DdzReqActGuo{} }
func (m *DdzReqActGuo) String() string            { return proto.CompactTextString(m) }
func (*DdzReqActGuo) ProtoMessage()               {}
func (*DdzReqActGuo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{19} }

func (m *DdzReqActGuo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqActGuo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

//
type DdzAckGuoAck struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzAckGuoAck) Reset()                    { *m = DdzAckGuoAck{} }
func (m *DdzAckGuoAck) String() string            { return proto.CompactTextString(m) }
func (*DdzAckGuoAck) ProtoMessage()               {}
func (*DdzAckGuoAck) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{20} }

func (m *DdzAckGuoAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckGuoAck) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 轮到谁操作
type DdzBcOverTurn struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Time             *int32               `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	ActType          *DdzEnumActType      `protobuf:"varint,4,opt,name=actType,enum=ddproto.DdzEnumActType" json:"actType,omitempty"`
	CanDouble        *bool                `protobuf:"varint,5,opt,name=canDouble" json:"canDouble,omitempty"`
	PullOrPush       *int32               `protobuf:"varint,6,opt,name=pullOrPush" json:"pullOrPush,omitempty"`
	CanOutCards      *bool                `protobuf:"varint,7,opt,name=canOutCards" json:"canOutCards,omitempty"`
	PlayerInfos      []*DdzBasePlayerInfo `protobuf:"bytes,8,rep,name=playerInfos" json:"playerInfos,omitempty"`
	JiaoScore        *DdzEnumJdScore      `protobuf:"varint,9,opt,name=jiaoScore,enum=ddproto.DdzEnumJdScore" json:"jiaoScore,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzBcOverTurn) Reset()                    { *m = DdzBcOverTurn{} }
func (m *DdzBcOverTurn) String() string            { return proto.CompactTextString(m) }
func (*DdzBcOverTurn) ProtoMessage()               {}
func (*DdzBcOverTurn) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{21} }

func (m *DdzBcOverTurn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBcOverTurn) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzBcOverTurn) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *DdzBcOverTurn) GetActType() DdzEnumActType {
	if m != nil && m.ActType != nil {
		return *m.ActType
	}
	return DdzEnumActType_T_NORMAL_ACT
}

func (m *DdzBcOverTurn) GetCanDouble() bool {
	if m != nil && m.CanDouble != nil {
		return *m.CanDouble
	}
	return false
}

func (m *DdzBcOverTurn) GetPullOrPush() int32 {
	if m != nil && m.PullOrPush != nil {
		return *m.PullOrPush
	}
	return 0
}

func (m *DdzBcOverTurn) GetCanOutCards() bool {
	if m != nil && m.CanOutCards != nil {
		return *m.CanOutCards
	}
	return false
}

func (m *DdzBcOverTurn) GetPlayerInfos() []*DdzBasePlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

func (m *DdzBcOverTurn) GetJiaoScore() DdzEnumJdScore {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return DdzEnumJdScore_DDZ_ONE
}

// 游戏信息(广播)
type DdzBcGameInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       []*DdzBasePlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	DdzDeskInfo      *DdzBaseDeskInfo     `protobuf:"bytes,3,opt,name=ddzDeskInfo" json:"ddzDeskInfo,omitempty"`
	SenderUserId     *uint32              `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *int32               `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzBcGameInfo) Reset()                    { *m = DdzBcGameInfo{} }
func (m *DdzBcGameInfo) String() string            { return proto.CompactTextString(m) }
func (*DdzBcGameInfo) ProtoMessage()               {}
func (*DdzBcGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{22} }

func (m *DdzBcGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBcGameInfo) GetPlayerInfo() []*DdzBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *DdzBcGameInfo) GetDdzDeskInfo() *DdzBaseDeskInfo {
	if m != nil {
		return m.DdzDeskInfo
	}
	return nil
}

func (m *DdzBcGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *DdzBcGameInfo) GetIsReconnect() int32 {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return 0
}

func init() {
	proto.RegisterType((*DdzBcOpening)(nil), "ddproto.ddz_bc_opening")
	proto.RegisterType((*DdzBcDealCards)(nil), "ddproto.ddz_bc_dealCards")
	proto.RegisterType((*DdzReqShowHandPokers)(nil), "ddproto.ddz_req_showHandPokers")
	proto.RegisterType((*DdzAckShowHandPokers)(nil), "ddproto.ddz_ack_showHandPokers")
	proto.RegisterType((*DdzReqJiaoDiZhu)(nil), "ddproto.ddz_req_jiaoDiZhu")
	proto.RegisterType((*DdzAckJiaoDiZhu)(nil), "ddproto.ddz_ack_jiaoDiZhu")
	proto.RegisterType((*DdzReqRobDiZhu)(nil), "ddproto.ddz_req_robDiZhu")
	proto.RegisterType((*DdzAckRobDiZhu)(nil), "ddproto.ddz_ack_robDiZhu")
	proto.RegisterType((*DdzReqDouble)(nil), "ddproto.ddz_req_double")
	proto.RegisterType((*DdzAckDouble)(nil), "ddproto.ddz_ack_double")
	proto.RegisterType((*DdzBcStartPlay)(nil), "ddproto.ddz_bc_startPlay")
	proto.RegisterType((*DdzReqMenuZhua)(nil), "ddproto.ddz_req_menuZhua")
	proto.RegisterType((*DdzAckMenuZhua)(nil), "ddproto.ddz_ack_menuZhua")
	proto.RegisterType((*DdzReqSeeCards)(nil), "ddproto.ddz_req_seeCards")
	proto.RegisterType((*DdzAckSeeCards)(nil), "ddproto.ddz_ack_seeCards")
	proto.RegisterType((*DdzReqPull)(nil), "ddproto.ddz_req_pull")
	proto.RegisterType((*DdzAckPull)(nil), "ddproto.ddz_ack_pull")
	proto.RegisterType((*DdzReqOutCards)(nil), "ddproto.ddz_req_outCards")
	proto.RegisterType((*DdzAckOutCards)(nil), "ddproto.ddz_ack_outCards")
	proto.RegisterType((*DdzReqActGuo)(nil), "ddproto.ddz_req_actGuo")
	proto.RegisterType((*DdzAckGuoAck)(nil), "ddproto.ddz_ack_guoAck")
	proto.RegisterType((*DdzBcOverTurn)(nil), "ddproto.ddz_bc_overTurn")
	proto.RegisterType((*DdzBcGameInfo)(nil), "ddproto.ddz_bc_gameInfo")
	proto.RegisterEnum("ddproto.DdzEnumJdScore", DdzEnumJdScore_name, DdzEnumJdScore_value)
	proto.RegisterEnum("ddproto.DdzEnumDoubleType", DdzEnumDoubleType_name, DdzEnumDoubleType_value)
	proto.RegisterEnum("ddproto.DdzEnumLaDaoType", DdzEnumLaDaoType_name, DdzEnumLaDaoType_value)
}

var fileDescriptor13 = []byte{
	// 811 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0x4d, 0x4f, 0x1a, 0x5d,
	0x14, 0x7e, 0x87, 0x91, 0xaf, 0x03, 0x22, 0x8e, 0xe6, 0x0d, 0x2f, 0x6f, 0x17, 0x86, 0x74, 0x41,
	0xa8, 0xa1, 0xd5, 0x45, 0x37, 0x4d, 0x9a, 0x60, 0xd1, 0x62, 0x63, 0x0a, 0xa1, 0x9a, 0x26, 0xa6,
	0x09, 0xb9, 0xcc, 0x5c, 0x65, 0x04, 0xee, 0xa5, 0x77, 0x66, 0xda, 0xea, 0xae, 0xcb, 0xee, 0xbb,
	0xee, 0xff, 0xe8, 0x0f, 0xe8, 0xff, 0xea, 0xfd, 0x98, 0x0f, 0x15, 0x2b, 0x30, 0x1a, 0xbb, 0x21,
	0xdc, 0x99, 0x73, 0x9e, 0xe7, 0x3c, 0xe7, 0x9c, 0x7b, 0xce, 0x40, 0xc1, 0xb2, 0x2e, 0x7a, 0x93,
	0x11, 0x3a, 0xaf, 0x4f, 0x18, 0x75, 0xa9, 0x91, 0xb6, 0x2c, 0xf9, 0xa7, 0xbc, 0x66, 0xd2, 0xf1,
	0x98, 0x92, 0x9e, 0x39, 0xb2, 0x31, 0x71, 0xd5, 0xdb, 0xb2, 0xb4, 0xee, 0x23, 0x07, 0xab, 0x73,
	0xe5, 0xb9, 0xf2, 0xef, 0x9b, 0x3d, 0x3a, 0xc1, 0xc4, 0x26, 0xa7, 0xc6, 0x63, 0x48, 0x0d, 0x30,
	0xb2, 0x30, 0x2b, 0x69, 0x1b, 0x5a, 0x35, 0xb7, 0xbd, 0x5e, 0xf7, 0x01, 0xeb, 0x1d, 0xf1, 0xdb,
	0x92, 0xef, 0x2a, 0x67, 0x50, 0xf4, 0xfd, 0x2c, 0x8c, 0x46, 0xaf, 0x10, 0xb3, 0x9c, 0xf9, 0x3c,
	0x8d, 0x67, 0x90, 0x17, 0xd1, 0x62, 0xd6, 0xa1, 0x43, 0xcc, 0x9c, 0x52, 0x62, 0x43, 0xe7, 0xb6,
	0xe5, 0xd0, 0x56, 0x85, 0x2b, 0x63, 0xec, 0x4d, 0x84, 0x49, 0xe5, 0x03, 0xfc, 0x2b, 0xb8, 0x18,
	0xfe, 0xd8, 0x73, 0x06, 0xf4, 0x73, 0x0b, 0x11, 0x4b, 0xf9, 0xce, 0xc9, 0x58, 0x80, 0x94, 0xe7,
	0x60, 0xb6, 0x6f, 0x71, 0x2e, 0xad, 0xba, 0x6c, 0xe4, 0x61, 0x89, 0x21, 0x17, 0x97, 0x74, 0x7e,
	0x4a, 0x56, 0xbe, 0x6b, 0x0a, 0x1e, 0x99, 0xc3, 0xfb, 0x81, 0xbf, 0x2e, 0x50, 0x9f, 0x25, 0xd0,
	0x30, 0x00, 0x94, 0x47, 0x57, 0x84, 0xb5, 0x24, 0xc3, 0xfa, 0xaa, 0xc1, 0x6a, 0xa0, 0xfa, 0xcc,
	0x46, 0xb4, 0x69, 0x1f, 0x0f, 0xbc, 0xf8, 0x82, 0x05, 0x84, 0x14, 0x9c, 0x31, 0xaa, 0x90, 0x74,
	0x4c, 0xca, 0x14, 0x51, 0x61, 0xfb, 0xbf, 0x10, 0x42, 0xd0, 0x61, 0xe2, 0x8d, 0x7b, 0x67, 0xd6,
	0x3b, 0x61, 0x10, 0xc6, 0x20, 0x52, 0xf3, 0xb7, 0x62, 0x38, 0x52, 0x8d, 0x26, 0xd2, 0xc0, 0x68,
	0xff, 0x2e, 0x11, 0xe4, 0x40, 0xe7, 0x08, 0x2a, 0x80, 0x00, 0x56, 0x28, 0xbb, 0x4f, 0x58, 0x57,
	0x5d, 0x27, 0x11, 0xad, 0x45, 0xbd, 0xfe, 0x08, 0xc7, 0x04, 0xdd, 0x84, 0x94, 0xf2, 0x97, 0xb8,
	0x85, 0xed, 0x47, 0xd3, 0x09, 0x52, 0xef, 0x0f, 0xcf, 0x27, 0x38, 0x60, 0x15, 0x62, 0x1e, 0x90,
	0xf5, 0x9b, 0x16, 0xce, 0x00, 0xc7, 0x45, 0xcc, 0xed, 0xf0, 0x16, 0x9e, 0x93, 0xb8, 0x0e, 0x70,
	0x42, 0xa9, 0x3b, 0xef, 0x04, 0x30, 0x8a, 0x90, 0x11, 0xf6, 0xdd, 0xf0, 0xd6, 0x1a, 0xcb, 0x90,
	0xb4, 0xec, 0x8b, 0x81, 0x27, 0x1b, 0x68, 0xb9, 0xd2, 0x8a, 0xba, 0x64, 0xcc, 0xe3, 0xe4, 0xd5,
	0x44, 0xf1, 0x72, 0x50, 0xf9, 0x12, 0x35, 0xc6, 0xdd, 0x90, 0x84, 0xc8, 0x41, 0x38, 0x4b, 0x66,
	0x4f, 0x81, 0xcb, 0x1a, 0x1c, 0x8c, 0x17, 0x19, 0xa9, 0xb7, 0x68, 0xb8, 0x1b, 0xd2, 0xc2, 0x1a,
	0x08, 0xe4, 0x03, 0x0d, 0x13, 0x6f, 0x34, 0x8a, 0xc9, 0x5a, 0x05, 0x1d, 0x99, 0xae, 0xdf, 0x84,
	0xff, 0x4f, 0x37, 0xe1, 0x08, 0x35, 0x11, 0x95, 0x3d, 0xe8, 0xf3, 0x09, 0xa5, 0x0f, 0xc2, 0x77,
	0x12, 0xd5, 0x88, 0x7a, 0xee, 0x22, 0x99, 0xdd, 0x84, 0x4c, 0xe0, 0x31, 0xc7, 0xca, 0xfb, 0xa9,
	0x45, 0x25, 0x5c, 0x90, 0xe8, 0xba, 0xb8, 0x27, 0x90, 0x31, 0xb9, 0xbb, 0x08, 0xdf, 0x57, 0x78,
	0xc3, 0xb4, 0x9d, 0x20, 0x5b, 0x18, 0x18, 0xeb, 0x90, 0x67, 0x78, 0x8c, 0x6c, 0xe2, 0x57, 0x5c,
	0xee, 0xa2, 0x2b, 0xb1, 0x27, 0x67, 0xc6, 0xbe, 0x17, 0xcd, 0x40, 0x9e, 0xd5, 0xd7, 0x1e, 0x8d,
	0xd9, 0xc5, 0x7b, 0xd1, 0x54, 0x3b, 0xf5, 0x68, 0xc3, 0x1c, 0xc6, 0xc4, 0xf9, 0x91, 0x80, 0x95,
	0xe0, 0x1b, 0xe7, 0x13, 0x66, 0x87, 0x1e, 0x23, 0xf1, 0x77, 0x98, 0x6b, 0x8f, 0x83, 0x11, 0x54,
	0x83, 0x34, 0xd7, 0x27, 0xf3, 0xfa, 0xc7, 0x2d, 0xe6, 0x1b, 0x18, 0xab, 0x90, 0x35, 0x11, 0x69,
	0xaa, 0xe1, 0x9a, 0x94, 0x2b, 0x50, 0x2c, 0x7d, 0xde, 0xb2, 0x6d, 0xd6, 0xf1, 0x9c, 0x41, 0x29,
	0x25, 0x21, 0xd7, 0x20, 0xc7, 0xcd, 0xda, 0x41, 0xae, 0xd3, 0xd2, 0x70, 0x0b, 0x72, 0xea, 0xeb,
	0x60, 0x9f, 0x9c, 0x50, 0xa7, 0x94, 0x91, 0x05, 0xb8, 0x3a, 0x9a, 0x55, 0xf6, 0x43, 0x23, 0x5e,
	0xb0, 0xac, 0x58, 0xb6, 0x72, 0x83, 0x96, 0xb2, 0xb3, 0x56, 0xec, 0x2f, 0x2d, 0x4c, 0xd0, 0x29,
	0x1a, 0x63, 0x89, 0x30, 0xef, 0xb7, 0x1c, 0x44, 0xac, 0x7e, 0x5b, 0xdf, 0x1e, 0xd9, 0x53, 0xc8,
	0xf1, 0xc7, 0x4d, 0xec, 0x0c, 0xa5, 0x8b, 0x2e, 0xc1, 0xcb, 0xd3, 0x2e, 0x96, 0x6f, 0x21, 0x3a,
	0xd2, 0xc1, 0x84, 0x93, 0x1d, 0xa9, 0x4a, 0xc8, 0x79, 0x2f, 0x12, 0x65, 0x3b, 0x5d, 0x6c, 0x52,
	0x42, 0x30, 0xbf, 0xb9, 0x22, 0xa3, 0xc9, 0xda, 0x0b, 0x75, 0x67, 0x2e, 0x6b, 0xe3, 0xdb, 0x39,
	0xdd, 0x6c, 0x1e, 0xf7, 0xda, 0x6f, 0x77, 0x8b, 0x5a, 0x70, 0x38, 0x7c, 0xdf, 0x2e, 0x26, 0xf8,
	0x06, 0xc9, 0xca, 0x43, 0xab, 0xbb, 0xbb, 0x5b, 0xd4, 0x6b, 0x5b, 0xb0, 0x76, 0xc3, 0x92, 0x0b,
	0x5c, 0xde, 0xec, 0x37, 0xb8, 0xbf, 0xef, 0xb2, 0x73, 0x24, 0x8e, 0x89, 0xda, 0x4b, 0x30, 0xa6,
	0x47, 0x04, 0x2f, 0x6b, 0x4a, 0x18, 0x1d, 0x34, 0x22, 0xc2, 0x66, 0x43, 0x10, 0xae, 0x40, 0x4e,
	0x79, 0x1f, 0x34, 0xc4, 0x03, 0x7d, 0x27, 0xd1, 0xd2, 0x3b, 0xff, 0x74, 0xb4, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xdd, 0x24, 0x74, 0xfd, 0xc2, 0x0b, 0x00, 0x00,
}
