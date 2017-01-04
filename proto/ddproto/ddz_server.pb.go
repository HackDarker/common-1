// Code generated by protoc-gen-go.
// source: ddz_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of ddz_base_roomTypeInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_commonRateInfo from ddz_base.proto

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

// 打出去的牌
type DdzSrvOutPokerPais struct {
	KeyValue         *int32               `protobuf:"varint,1,opt,name=keyValue" json:"keyValue,omitempty"`
	PokerPais        []*CommonSrvPokerPai `protobuf:"bytes,2,rep,name=pokerPais" json:"pokerPais,omitempty"`
	Type             *int32               `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	IsBomb           *bool                `protobuf:"varint,4,opt,name=isBomb" json:"isBomb,omitempty"`
	CountDuizi       *int32               `protobuf:"varint,5,opt,name=countDuizi" json:"countDuizi,omitempty"`
	CountSanzhang    *int32               `protobuf:"varint,6,opt,name=countSanzhang" json:"countSanzhang,omitempty"`
	CountSizhang     *int32               `protobuf:"varint,7,opt,name=countSizhang" json:"countSizhang,omitempty"`
	CountYizhang     *int32               `protobuf:"varint,8,opt,name=countYizhang" json:"countYizhang,omitempty"`
	UserId           *uint32              `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzSrvOutPokerPais) Reset()                    { *m = DdzSrvOutPokerPais{} }
func (m *DdzSrvOutPokerPais) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvOutPokerPais) ProtoMessage()               {}
func (*DdzSrvOutPokerPais) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *DdzSrvOutPokerPais) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetPokerPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.PokerPais
	}
	return nil
}

func (m *DdzSrvOutPokerPais) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetIsBomb() bool {
	if m != nil && m.IsBomb != nil {
		return *m.IsBomb
	}
	return false
}

func (m *DdzSrvOutPokerPais) GetCountDuizi() int32 {
	if m != nil && m.CountDuizi != nil {
		return *m.CountDuizi
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountSanzhang() int32 {
	if m != nil && m.CountSanzhang != nil {
		return *m.CountSanzhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountSizhang() int32 {
	if m != nil && m.CountSizhang != nil {
		return *m.CountSizhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountYizhang() int32 {
	if m != nil && m.CountYizhang != nil {
		return *m.CountYizhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 对 desk的统计
type DdzSrvDeskTongJi struct {
	Bombs            []*DdzSrvOutPokerPais `protobuf:"bytes,1,rep,name=bombs" json:"bombs,omitempty"`
	CountQiangDiZhu  *int32                `protobuf:"varint,2,opt,name=countQiangDiZhu" json:"countQiangDiZhu,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvDeskTongJi) Reset()                    { *m = DdzSrvDeskTongJi{} }
func (m *DdzSrvDeskTongJi) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDeskTongJi) ProtoMessage()               {}
func (*DdzSrvDeskTongJi) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *DdzSrvDeskTongJi) GetBombs() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.Bombs
	}
	return nil
}

func (m *DdzSrvDeskTongJi) GetCountQiangDiZhu() int32 {
	if m != nil && m.CountQiangDiZhu != nil {
		return *m.CountQiangDiZhu
	}
	return 0
}

// desk
type DdzSrvDesk struct {
	DeskId           *int32                 `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	Key              *string                `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	UserCountLimit   *int32                 `protobuf:"varint,3,opt,name=userCountLimit" json:"userCountLimit,omitempty"`
	AllPokerPai      []*CommonSrvPokerPai   `protobuf:"bytes,4,rep,name=allPokerPai" json:"allPokerPai,omitempty"`
	DiPokerPai       []*CommonSrvPokerPai   `protobuf:"bytes,5,rep,name=diPokerPai" json:"diPokerPai,omitempty"`
	OutPai           *DdzSrvOutPokerPais    `protobuf:"bytes,6,opt,name=outPai" json:"outPai,omitempty"`
	Owner            *uint32                `protobuf:"varint,7,opt,name=owner" json:"owner,omitempty"`
	DizhuPaiUser     *uint32                `protobuf:"varint,8,opt,name=dizhuPaiUser" json:"dizhuPaiUser,omitempty"`
	DiZhuUserId      *uint32                `protobuf:"varint,9,opt,name=diZhuUserId" json:"diZhuUserId,omitempty"`
	ActiveUserId     *uint32                `protobuf:"varint,10,opt,name=activeUserId" json:"activeUserId,omitempty"`
	Tongji           *DdzSrvDeskTongJi      `protobuf:"bytes,11,opt,name=tongji" json:"tongji,omitempty"`
	BaseValue        *int64                 `protobuf:"varint,12,opt,name=baseValue" json:"baseValue,omitempty"`
	QingDizhuValue   *int64                 `protobuf:"varint,13,opt,name=qingDizhuValue" json:"qingDizhuValue,omitempty"`
	WinValue         *int64                 `protobuf:"varint,14,opt,name=winValue" json:"winValue,omitempty"`
	DdzType          *int32                 `protobuf:"varint,15,opt,name=ddzType" json:"ddzType,omitempty"`
	RoomType         *int32                 `protobuf:"varint,16,opt,name=RoomType" json:"RoomType,omitempty"`
	BoardsCount      *int32                 `protobuf:"varint,17,opt,name=BoardsCount" json:"BoardsCount,omitempty"`
	CapMax           *int64                 `protobuf:"varint,18,opt,name=CapMax" json:"CapMax,omitempty"`
	IsJiaoFen        *bool                  `protobuf:"varint,19,opt,name=IsJiaoFen" json:"IsJiaoFen,omitempty"`
	RoomId           *int32                 `protobuf:"varint,20,opt,name=roomId" json:"roomId,omitempty"`
	CommonRateInfo   *DdzBaseCommonRateInfo `protobuf:"bytes,21,opt,name=commonRateInfo" json:"commonRateInfo,omitempty"`
	PlayRate         *int32                 `protobuf:"varint,22,opt,name=playRate" json:"playRate,omitempty"`
	GameStatus       *int32                 `protobuf:"varint,23,opt,name=GameStatus" json:"GameStatus,omitempty"`
	InitRoomCoin     *int64                 `protobuf:"varint,24,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32                 `protobuf:"varint,25,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32                 `protobuf:"varint,26,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	PlayerNum        *int32                 `protobuf:"varint,27,opt,name=playerNum" json:"playerNum,omitempty"`
	TimeOutSecond    *int32                 `protobuf:"varint,28,opt,name=timeOutSecond" json:"timeOutSecond,omitempty"`
	UserMinCoin      *int64                 `protobuf:"varint,29,opt,name=userMinCoin" json:"userMinCoin,omitempty"`
	UserMaxCoin      *int64                 `protobuf:"varint,30,opt,name=userMaxCoin" json:"userMaxCoin,omitempty"`
	CoinTicket       *int64                 `protobuf:"varint,31,opt,name=coinTicket" json:"coinTicket,omitempty"`
	CoinRoomLv       *DdzEnumCoinRoomLevel  `protobuf:"varint,32,opt,name=coinRoomLv,enum=ddproto.DdzEnumCoinRoomLevel" json:"coinRoomLv,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *DdzSrvDesk) Reset()                    { *m = DdzSrvDesk{} }
func (m *DdzSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDesk) ProtoMessage()               {}
func (*DdzSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *DdzSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzSrvDesk) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *DdzSrvDesk) GetUserCountLimit() int32 {
	if m != nil && m.UserCountLimit != nil {
		return *m.UserCountLimit
	}
	return 0
}

func (m *DdzSrvDesk) GetAllPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.AllPokerPai
	}
	return nil
}

func (m *DdzSrvDesk) GetDiPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.DiPokerPai
	}
	return nil
}

func (m *DdzSrvDesk) GetOutPai() *DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *DdzSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *DdzSrvDesk) GetDizhuPaiUser() uint32 {
	if m != nil && m.DizhuPaiUser != nil {
		return *m.DizhuPaiUser
	}
	return 0
}

func (m *DdzSrvDesk) GetDiZhuUserId() uint32 {
	if m != nil && m.DiZhuUserId != nil {
		return *m.DiZhuUserId
	}
	return 0
}

func (m *DdzSrvDesk) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *DdzSrvDesk) GetTongji() *DdzSrvDeskTongJi {
	if m != nil {
		return m.Tongji
	}
	return nil
}

func (m *DdzSrvDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *DdzSrvDesk) GetQingDizhuValue() int64 {
	if m != nil && m.QingDizhuValue != nil {
		return *m.QingDizhuValue
	}
	return 0
}

func (m *DdzSrvDesk) GetWinValue() int64 {
	if m != nil && m.WinValue != nil {
		return *m.WinValue
	}
	return 0
}

func (m *DdzSrvDesk) GetDdzType() int32 {
	if m != nil && m.DdzType != nil {
		return *m.DdzType
	}
	return 0
}

func (m *DdzSrvDesk) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *DdzSrvDesk) GetBoardsCount() int32 {
	if m != nil && m.BoardsCount != nil {
		return *m.BoardsCount
	}
	return 0
}

func (m *DdzSrvDesk) GetCapMax() int64 {
	if m != nil && m.CapMax != nil {
		return *m.CapMax
	}
	return 0
}

func (m *DdzSrvDesk) GetIsJiaoFen() bool {
	if m != nil && m.IsJiaoFen != nil {
		return *m.IsJiaoFen
	}
	return false
}

func (m *DdzSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *DdzSrvDesk) GetCommonRateInfo() *DdzBaseCommonRateInfo {
	if m != nil {
		return m.CommonRateInfo
	}
	return nil
}

func (m *DdzSrvDesk) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvDesk) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *DdzSrvDesk) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *DdzSrvDesk) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *DdzSrvDesk) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *DdzSrvDesk) GetTimeOutSecond() int32 {
	if m != nil && m.TimeOutSecond != nil {
		return *m.TimeOutSecond
	}
	return 0
}

func (m *DdzSrvDesk) GetUserMinCoin() int64 {
	if m != nil && m.UserMinCoin != nil {
		return *m.UserMinCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetUserMaxCoin() int64 {
	if m != nil && m.UserMaxCoin != nil {
		return *m.UserMaxCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetCoinTicket() int64 {
	if m != nil && m.CoinTicket != nil {
		return *m.CoinTicket
	}
	return 0
}

func (m *DdzSrvDesk) GetCoinRoomLv() DdzEnumCoinRoomLevel {
	if m != nil && m.CoinRoomLv != nil {
		return *m.CoinRoomLv
	}
	return DdzEnumCoinRoomLevel_LV1
}

// 游戏数据
type DdzSrvGameData struct {
	HandPokers       []*CommonSrvPokerPai  `protobuf:"bytes,1,rep,name=handPokers" json:"handPokers,omitempty"`
	OutPaiList       []*DdzSrvOutPokerPais `protobuf:"bytes,2,rep,name=outPaiList" json:"outPaiList,omitempty"`
	OutPai           *DdzSrvOutPokerPais   `protobuf:"bytes,3,opt,name=outPai" json:"outPai,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvGameData) Reset()                    { *m = DdzSrvGameData{} }
func (m *DdzSrvGameData) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvGameData) ProtoMessage()               {}
func (*DdzSrvGameData) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *DdzSrvGameData) GetHandPokers() []*CommonSrvPokerPai {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *DdzSrvGameData) GetOutPaiList() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPaiList
	}
	return nil
}

func (m *DdzSrvGameData) GetOutPai() *DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

type DdzSrvBillBean struct {
	// 斗地主的账单
	Coin             *int64  `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	LoseUser         *uint32 `protobuf:"varint,2,opt,name=loseUser" json:"loseUser,omitempty"`
	WinUser          *uint32 `protobuf:"varint,3,opt,name=winUser" json:"winUser,omitempty"`
	Desc             *string `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DdzSrvBillBean) Reset()                    { *m = DdzSrvBillBean{} }
func (m *DdzSrvBillBean) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBillBean) ProtoMessage()               {}
func (*DdzSrvBillBean) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *DdzSrvBillBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvBillBean) GetLoseUser() uint32 {
	if m != nil && m.LoseUser != nil {
		return *m.LoseUser
	}
	return 0
}

func (m *DdzSrvBillBean) GetWinUser() uint32 {
	if m != nil && m.WinUser != nil {
		return *m.WinUser
	}
	return 0
}

func (m *DdzSrvBillBean) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

type DdzSrvBill struct {
	// 斗地主的账单
	WinCoin          *int64            `protobuf:"varint,1,opt,name=winCoin" json:"winCoin,omitempty"`
	BillBean         []*DdzSrvBillBean `protobuf:"bytes,2,rep,name=billBean" json:"billBean,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *DdzSrvBill) Reset()                    { *m = DdzSrvBill{} }
func (m *DdzSrvBill) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBill) ProtoMessage()               {}
func (*DdzSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *DdzSrvBill) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzSrvBill) GetBillBean() []*DdzSrvBillBean {
	if m != nil {
		return m.BillBean
	}
	return nil
}

type DdzSrvUserStatisticsRound struct {
	// 玩家每一局的统计信息
	Round            *int32                `protobuf:"varint,1,opt,name=round" json:"round,omitempty"`
	WinCoin          *int64                `protobuf:"varint,2,opt,name=winCoin" json:"winCoin,omitempty"`
	CountBomb        *int32                `protobuf:"varint,3,opt,name=countBomb" json:"countBomb,omitempty"`
	BomBean          []*DdzSrvOutPokerPais `protobuf:"bytes,4,rep,name=bomBean" json:"bomBean,omitempty"`
	PlayRate         *int32                `protobuf:"varint,5,opt,name=playRate" json:"playRate,omitempty"`
	IsSpring         *bool                 `protobuf:"varint,6,opt,name=isSpring" json:"isSpring,omitempty"`
	IsDizhu          *bool                 `protobuf:"varint,7,opt,name=isDizhu" json:"isDizhu,omitempty"`
	IsWin            *bool                 `protobuf:"varint,8,opt,name=isWin" json:"isWin,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUserStatisticsRound) Reset()                    { *m = DdzSrvUserStatisticsRound{} }
func (m *DdzSrvUserStatisticsRound) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatisticsRound) ProtoMessage()               {}
func (*DdzSrvUserStatisticsRound) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *DdzSrvUserStatisticsRound) GetRound() int32 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetBomBean() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.BomBean
	}
	return nil
}

func (m *DdzSrvUserStatisticsRound) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

func (m *DdzSrvUserStatisticsRound) GetIsDizhu() bool {
	if m != nil && m.IsDizhu != nil {
		return *m.IsDizhu
	}
	return false
}

func (m *DdzSrvUserStatisticsRound) GetIsWin() bool {
	if m != nil && m.IsWin != nil {
		return *m.IsWin
	}
	return false
}

type DdzSrvUserStatistics struct {
	// 玩家统计信息
	RoundBean        []*DdzSrvUserStatisticsRound `protobuf:"bytes,1,rep,name=roundBean" json:"roundBean,omitempty"`
	CountWin         *int32                       `protobuf:"varint,2,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32                       `protobuf:"varint,3,opt,name=countLose" json:"countLose,omitempty"`
	CountSpring      *int32                       `protobuf:"varint,4,opt,name=countSpring" json:"countSpring,omitempty"`
	CountDizhu       *int32                       `protobuf:"varint,5,opt,name=countDizhu" json:"countDizhu,omitempty"`
	CountBomb        *int32                       `protobuf:"varint,6,opt,name=countBomb" json:"countBomb,omitempty"`
	MaxWinCoin       *int64                       `protobuf:"varint,7,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *DdzSrvUserStatistics) Reset()                    { *m = DdzSrvUserStatistics{} }
func (m *DdzSrvUserStatistics) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatistics) ProtoMessage()               {}
func (*DdzSrvUserStatistics) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

func (m *DdzSrvUserStatistics) GetRoundBean() []*DdzSrvUserStatisticsRound {
	if m != nil {
		return m.RoundBean
	}
	return nil
}

func (m *DdzSrvUserStatistics) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountSpring() int32 {
	if m != nil && m.CountSpring != nil {
		return *m.CountSpring
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountDizhu() int32 {
	if m != nil && m.CountDizhu != nil {
		return *m.CountDizhu
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetMaxWinCoin() int64 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

// user
type DdzSrvUser struct {
	UserId           *uint32               `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	GameData         *DdzSrvGameData       `protobuf:"bytes,2,opt,name=gameData" json:"gameData,omitempty"`
	Status           *int32                `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	IsBreak          *bool                 `protobuf:"varint,4,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool                 `protobuf:"varint,5,opt,name=isLeave" json:"isLeave,omitempty"`
	StatusHLQiang    *int32                `protobuf:"varint,7,opt,name=statusHLQiang" json:"statusHLQiang,omitempty"`
	StatusHLJiao     *int32                `protobuf:"varint,8,opt,name=statusHLJiao" json:"statusHLJiao,omitempty"`
	StatusHLJiaBei   *int32                `protobuf:"varint,9,opt,name=statusHLJiaBei" json:"statusHLJiaBei,omitempty"`
	StatusSCMenZhua  *int32                `protobuf:"varint,10,opt,name=statusSCMenZhua" json:"statusSCMenZhua,omitempty"`
	StatusSCZhuaPai  *int32                `protobuf:"varint,11,opt,name=statusSCZhuaPai" json:"statusSCZhuaPai,omitempty"`
	StatusSCLaDao    *int32                `protobuf:"varint,12,opt,name=statusSCLaDao" json:"statusSCLaDao,omitempty"`
	StatusJDJiao     *int32                `protobuf:"varint,13,opt,name=statusJDJiao" json:"statusJDJiao,omitempty"`
	IsShowPokers     *bool                 `protobuf:"varint,15,opt,name=isShowPokers" json:"isShowPokers,omitempty"`
	Bill             *DdzSrvBill           `protobuf:"bytes,16,opt,name=bill" json:"bill,omitempty"`
	DeskId           *int32                `protobuf:"varint,17,opt,name=deskId" json:"deskId,omitempty"`
	RoomId           *int32                `protobuf:"varint,18,opt,name=roomId" json:"roomId,omitempty"`
	Coin             *int64                `protobuf:"varint,19,opt,name=coin" json:"coin,omitempty"`
	Statistics       *DdzSrvUserStatistics `protobuf:"bytes,20,opt,name=statistics" json:"statistics,omitempty"`
	PlayRate         *int32                `protobuf:"varint,21,opt,name=playRate" json:"playRate,omitempty"`
	JiaoScore        *int32                `protobuf:"varint,22,opt,name=jiaoScore" json:"jiaoScore,omitempty"`
	TimeOutCount     *int32                `protobuf:"varint,23,opt,name=timeOutCount" json:"timeOutCount,omitempty"`
	IsAgent          *bool                 `protobuf:"varint,24,opt,name=isAgent" json:"isAgent,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUser) Reset()                    { *m = DdzSrvUser{} }
func (m *DdzSrvUser) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUser) ProtoMessage()               {}
func (*DdzSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *DdzSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzSrvUser) GetGameData() *DdzSrvGameData {
	if m != nil {
		return m.GameData
	}
	return nil
}

func (m *DdzSrvUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *DdzSrvUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *DdzSrvUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *DdzSrvUser) GetStatusHLQiang() int32 {
	if m != nil && m.StatusHLQiang != nil {
		return *m.StatusHLQiang
	}
	return 0
}

func (m *DdzSrvUser) GetStatusHLJiao() int32 {
	if m != nil && m.StatusHLJiao != nil {
		return *m.StatusHLJiao
	}
	return 0
}

func (m *DdzSrvUser) GetStatusHLJiaBei() int32 {
	if m != nil && m.StatusHLJiaBei != nil {
		return *m.StatusHLJiaBei
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCMenZhua() int32 {
	if m != nil && m.StatusSCMenZhua != nil {
		return *m.StatusSCMenZhua
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCZhuaPai() int32 {
	if m != nil && m.StatusSCZhuaPai != nil {
		return *m.StatusSCZhuaPai
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCLaDao() int32 {
	if m != nil && m.StatusSCLaDao != nil {
		return *m.StatusSCLaDao
	}
	return 0
}

func (m *DdzSrvUser) GetStatusJDJiao() int32 {
	if m != nil && m.StatusJDJiao != nil {
		return *m.StatusJDJiao
	}
	return 0
}

func (m *DdzSrvUser) GetIsShowPokers() bool {
	if m != nil && m.IsShowPokers != nil {
		return *m.IsShowPokers
	}
	return false
}

func (m *DdzSrvUser) GetBill() *DdzSrvBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *DdzSrvUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzSrvUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *DdzSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvUser) GetStatistics() *DdzSrvUserStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *DdzSrvUser) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUser) GetJiaoScore() int32 {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return 0
}

func (m *DdzSrvUser) GetTimeOutCount() int32 {
	if m != nil && m.TimeOutCount != nil {
		return *m.TimeOutCount
	}
	return 0
}

func (m *DdzSrvUser) GetIsAgent() bool {
	if m != nil && m.IsAgent != nil {
		return *m.IsAgent
	}
	return false
}

// room
type DdzSrvRoom struct {
	RoomId           *int32 `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DdzSrvRoom) Reset()                    { *m = DdzSrvRoom{} }
func (m *DdzSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvRoom) ProtoMessage()               {}
func (*DdzSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{9} }

func (m *DdzSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

// 备份专用...
type DdzSrvBak struct {
	Desk             *DdzSrvDesk   `protobuf:"bytes,1,opt,name=desk" json:"desk,omitempty"`
	Users            []*DdzSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DdzSrvBak) Reset()                    { *m = DdzSrvBak{} }
func (m *DdzSrvBak) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBak) ProtoMessage()               {}
func (*DdzSrvBak) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{10} }

func (m *DdzSrvBak) GetDesk() *DdzSrvDesk {
	if m != nil {
		return m.Desk
	}
	return nil
}

func (m *DdzSrvBak) GetUsers() []*DdzSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*DdzSrvOutPokerPais)(nil), "ddproto.ddz_srv_outPokerPais")
	proto.RegisterType((*DdzSrvDeskTongJi)(nil), "ddproto.ddz_srv_deskTongJi")
	proto.RegisterType((*DdzSrvDesk)(nil), "ddproto.ddz_srv_desk")
	proto.RegisterType((*DdzSrvGameData)(nil), "ddproto.ddz_srv_gameData")
	proto.RegisterType((*DdzSrvBillBean)(nil), "ddproto.ddz_srv_billBean")
	proto.RegisterType((*DdzSrvBill)(nil), "ddproto.ddz_srv_bill")
	proto.RegisterType((*DdzSrvUserStatisticsRound)(nil), "ddproto.ddz_srv_userStatisticsRound")
	proto.RegisterType((*DdzSrvUserStatistics)(nil), "ddproto.ddz_srv_userStatistics")
	proto.RegisterType((*DdzSrvUser)(nil), "ddproto.ddz_srv_user")
	proto.RegisterType((*DdzSrvRoom)(nil), "ddproto.ddz_srv_room")
	proto.RegisterType((*DdzSrvBak)(nil), "ddproto.ddz_srv_bak")
}

var fileDescriptor13 = []byte{
	// 1132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xdb, 0x52, 0x23, 0x37,
	0x13, 0xfe, 0xbd, 0xf6, 0x18, 0x5b, 0xc6, 0x86, 0x1d, 0x4e, 0x02, 0xf6, 0x40, 0xcd, 0xbf, 0x17,
	0x5b, 0xb5, 0x09, 0xc9, 0x92, 0x8b, 0xe4, 0x36, 0x40, 0x25, 0x81, 0x32, 0x89, 0x17, 0xb3, 0x21,
	0x87, 0x0b, 0x97, 0xf0, 0x28, 0xa0, 0xb5, 0x67, 0xe4, 0xcc, 0xc1, 0x1c, 0x5e, 0x21, 0xef, 0x91,
	0xc7, 0xc8, 0x55, 0x2a, 0x8f, 0x91, 0x67, 0x49, 0x77, 0x4b, 0xb2, 0xc7, 0xc1, 0x15, 0xb8, 0x82,
	0xf9, 0xd4, 0x52, 0xf7, 0xd7, 0xfa, 0xfa, 0x93, 0xd9, 0x72, 0x18, 0xde, 0xf5, 0x52, 0x99, 0x8c,
	0x65, 0xb2, 0x3b, 0x4a, 0x74, 0xa6, 0xfd, 0x85, 0x30, 0xa4, 0x7f, 0xb6, 0x36, 0xfb, 0x3a, 0x8a,
	0x74, 0x6c, 0x57, 0x7b, 0x23, 0x3d, 0x70, 0x31, 0x5b, 0x2d, 0xdc, 0x75, 0x21, 0x52, 0x69, 0xbe,
	0x83, 0xbf, 0x4b, 0x6c, 0x95, 0x0e, 0x4a, 0xc6, 0x3d, 0x9d, 0x67, 0x1d, 0x0c, 0xed, 0x08, 0x95,
	0xfa, 0xcb, 0xac, 0x36, 0x90, 0xb7, 0xdf, 0x8b, 0x61, 0x2e, 0x79, 0x69, 0xa7, 0xf4, 0xda, 0xf3,
	0x3f, 0x61, 0xf5, 0x91, 0x5b, 0xe6, 0x4f, 0x76, 0xca, 0xaf, 0x1b, 0x7b, 0xcf, 0x76, 0x6d, 0xca,
	0x5d, 0x97, 0x11, 0x8e, 0x71, 0x41, 0xfe, 0x22, 0xab, 0x64, 0xb7, 0x23, 0xc9, 0xcb, 0xb4, 0xbd,
	0xc5, 0xaa, 0x2a, 0xdd, 0xd7, 0xd1, 0x05, 0xaf, 0xc0, 0x77, 0xcd, 0xf7, 0x19, 0xeb, 0xeb, 0x3c,
	0xce, 0x0e, 0x73, 0x75, 0xa7, 0xb8, 0x47, 0x31, 0x6b, 0xac, 0x49, 0x58, 0x57, 0xc4, 0x77, 0x57,
	0x22, 0xbe, 0xe4, 0x55, 0x82, 0x57, 0xd9, 0xa2, 0x81, 0x95, 0x41, 0x17, 0x66, 0xd0, 0x1f, 0x2d,
	0x5a, 0x73, 0x69, 0x72, 0xe0, 0x7d, 0x14, 0xf2, 0x3a, 0x7c, 0x37, 0x83, 0x9f, 0x99, 0xef, 0xf8,
	0x85, 0x32, 0x1d, 0x9c, 0xe9, 0xf8, 0xf2, 0x58, 0xf9, 0x1f, 0x31, 0xef, 0x02, 0x4a, 0x49, 0x81,
	0x1a, 0xf2, 0x78, 0x3e, 0xe1, 0x31, 0xb7, 0x17, 0x1b, 0x6c, 0x89, 0x32, 0xbd, 0x53, 0x90, 0xe7,
	0x50, 0xfd, 0x74, 0x95, 0x03, 0x7f, 0x48, 0x16, 0xfc, 0x51, 0x65, 0x8b, 0xc5, 0xd3, 0x31, 0x3b,
	0xfe, 0x85, 0xec, 0xa6, 0x67, 0x0d, 0x56, 0x86, 0x2e, 0x52, 0x74, 0xdd, 0x5f, 0x67, 0x2d, 0x2c,
	0xed, 0x00, 0x8f, 0x6a, 0xab, 0x48, 0x65, 0xb6, 0x33, 0x6f, 0x59, 0x43, 0x0c, 0x87, 0x2e, 0x1d,
	0xb4, 0xe7, 0xe1, 0xd6, 0x7e, 0xca, 0x58, 0xa8, 0x26, 0x3b, 0xbc, 0x47, 0xec, 0xf8, 0x98, 0x55,
	0x91, 0x13, 0x44, 0x63, 0x4f, 0x1f, 0xa4, 0xdc, 0x64, 0x9e, 0xbe, 0x8e, 0x65, 0x42, 0xbd, 0x6e,
	0x62, 0xaf, 0x43, 0x68, 0x73, 0x0e, 0x6b, 0xef, 0x81, 0x02, 0xf5, 0xba, 0xe9, 0xaf, 0xb0, 0x46,
	0x88, 0xdd, 0x78, 0x5f, 0x68, 0x38, 0x86, 0x8a, 0x7e, 0xa6, 0xc6, 0xd2, 0xa2, 0x8c, 0xd0, 0x37,
	0xac, 0x9a, 0x41, 0xeb, 0x3f, 0x28, 0xde, 0xa0, 0xf4, 0xdb, 0xf7, 0xd2, 0x17, 0x6e, 0xe7, 0x29,
	0xab, 0xa3, 0x44, 0x8d, 0xf8, 0x16, 0x21, 0xbe, 0x8c, 0xbd, 0xfb, 0x55, 0x61, 0xf3, 0xa1, 0x08,
	0x83, 0x37, 0x09, 0x07, 0x99, 0x5e, 0xab, 0xd8, 0x20, 0x2d, 0x42, 0x96, 0x18, 0xcc, 0xc1, 0xdd,
	0x19, 0x0a, 0x6f, 0x89, 0xda, 0x0b, 0x21, 0xa7, 0x5a, 0x47, 0x84, 0x2c, 0x13, 0x02, 0x75, 0xef,
	0x6b, 0x91, 0x84, 0x29, 0x5d, 0x05, 0x7f, 0xea, 0x84, 0x73, 0x20, 0x46, 0x27, 0xe2, 0x86, 0xfb,
	0x74, 0x0e, 0x14, 0x71, 0x94, 0x1e, 0x2b, 0xa1, 0xbf, 0x92, 0x31, 0x5f, 0x21, 0xc9, 0x42, 0x48,
	0x02, 0x27, 0x01, 0xa9, 0x55, 0xda, 0xf2, 0x05, 0x6b, 0x99, 0x56, 0x9f, 0x8a, 0x4c, 0x1e, 0xc5,
	0xbf, 0x68, 0xbe, 0x46, 0xe4, 0x76, 0x66, 0xc8, 0x21, 0x95, 0xde, 0x6c, 0x1c, 0xd6, 0x34, 0x1a,
	0x8a, 0x5b, 0xfc, 0xe6, 0xeb, 0x74, 0x16, 0x8c, 0xc3, 0xd7, 0x22, 0x92, 0xdd, 0x4c, 0x64, 0x79,
	0xca, 0x37, 0x9c, 0xc2, 0x55, 0xac, 0x32, 0xac, 0xfe, 0x40, 0xab, 0x98, 0x73, 0x2a, 0x0c, 0x87,
	0x24, 0x4f, 0x92, 0x0e, 0xec, 0x37, 0xf5, 0x6f, 0x52, 0x30, 0x74, 0x28, 0xd3, 0x99, 0x18, 0x4e,
	0xf1, 0x2d, 0xc2, 0x81, 0x07, 0xa6, 0x92, 0xc9, 0xb7, 0x79, 0xc4, 0xb7, 0xdd, 0x98, 0x65, 0x2a,
	0x92, 0xdf, 0xe5, 0x59, 0x57, 0xf6, 0x75, 0x1c, 0xf2, 0x67, 0xae, 0x2d, 0xa8, 0xcf, 0x13, 0x15,
	0x53, 0xb6, 0xe7, 0x94, 0xcd, 0x81, 0xe2, 0x86, 0xc0, 0x17, 0x04, 0xd2, 0xec, 0xaa, 0xf8, 0x4c,
	0xf5, 0x07, 0x32, 0xe3, 0x2f, 0x09, 0xfb, 0xcc, 0x60, 0x58, 0x6c, 0x7b, 0xcc, 0x77, 0x00, 0x6b,
	0xed, 0xbd, 0x9c, 0x69, 0x84, 0x8c, 0xf3, 0xa8, 0x37, 0x89, 0x91, 0x63, 0x39, 0x0c, 0x7e, 0x2f,
	0x59, 0x1f, 0x03, 0x01, 0x5c, 0x02, 0xfd, 0x43, 0x91, 0x09, 0x14, 0x37, 0x0c, 0x74, 0x48, 0x62,
	0x74, 0x13, 0xfa, 0xdf, 0xe2, 0x7e, 0xcb, 0x98, 0x11, 0x77, 0x5b, 0xa5, 0x99, 0xf5, 0xa6, 0x07,
	0x04, 0x3e, 0x9d, 0x87, 0xf2, 0x23, 0xe6, 0x21, 0x78, 0x37, 0xad, 0xf3, 0x42, 0x0d, 0x87, 0xfb,
	0x52, 0xc4, 0xe8, 0x6f, 0xc8, 0x86, 0x46, 0x9d, 0x94, 0x38, 0xd4, 0x29, 0xa9, 0x9e, 0xe6, 0xbd,
	0x89, 0x4a, 0x04, 0x6d, 0x12, 0x50, 0x26, 0x00, 0x36, 0x80, 0xca, 0xfb, 0x64, 0x80, 0xf5, 0xa0,
	0x3d, 0xf5, 0x0e, 0x3c, 0xd2, 0x86, 0x1f, 0x4c, 0x4f, 0x7c, 0xc3, 0x6a, 0x2e, 0x97, 0xe5, 0xb4,
	0x79, 0xaf, 0x48, 0x17, 0x10, 0xfc, 0x59, 0x62, 0xdb, 0x0e, 0xc4, 0x0b, 0x43, 0x21, 0x41, 0x33,
	0x54, 0x3f, 0x3d, 0x05, 0x31, 0x84, 0x38, 0xd0, 0x09, 0xfe, 0x63, 0x8d, 0xa9, 0x90, 0xec, 0x89,
	0x93, 0x3b, 0x79, 0x1c, 0x39, 0xb4, 0xf1, 0xa5, 0x5d, 0xb6, 0x00, 0x26, 0x49, 0xe9, 0x2b, 0x8f,
	0x69, 0x69, 0x51, 0xd4, 0x9e, 0x1b, 0x3d, 0x95, 0x76, 0x47, 0x89, 0xb2, 0x56, 0x5e, 0xc3, 0xbc,
	0x2a, 0xa5, 0x29, 0x26, 0x67, 0xa9, 0x61, 0x5d, 0x2a, 0x3d, 0x87, 0x32, 0xd0, 0x52, 0x6a, 0xc1,
	0x5f, 0x25, 0xb6, 0x3e, 0x9f, 0x86, 0xff, 0x39, 0xab, 0x13, 0x03, 0x2a, 0xc8, 0xa8, 0xe2, 0xd5,
	0xbd, 0x82, 0xe6, 0x51, 0x87, 0x2a, 0x88, 0xda, 0xb9, 0x25, 0xeb, 0x4d, 0xc8, 0xb6, 0xe1, 0xc2,
	0x2c, 0x59, 0xd0, 0xb9, 0x79, 0x63, 0x4c, 0xb5, 0x15, 0x37, 0x94, 0xe6, 0x8d, 0xa2, 0x82, 0xbd,
	0x99, 0xbd, 0xd4, 0xa8, 0xaa, 0x0b, 0x8b, 0xc4, 0xcd, 0xb9, 0xed, 0x27, 0xf2, 0x2a, 0x07, 0xbf,
	0x55, 0xa6, 0xd7, 0x8b, 0x45, 0x15, 0x1e, 0xa6, 0x92, 0x75, 0xc4, 0x9a, 0x53, 0x3c, 0x55, 0x35,
	0xef, 0x76, 0x27, 0x23, 0x01, 0x9b, 0x53, 0xe3, 0x0c, 0x65, 0x77, 0x7d, 0xf0, 0x98, 0x26, 0x52,
	0x0c, 0xec, 0x6b, 0x4a, 0x40, 0x5b, 0x8a, 0xb1, 0x69, 0x7d, 0x0d, 0x67, 0xdc, 0xec, 0xf8, 0xa6,
	0x4d, 0xcf, 0xd6, 0xf4, 0xd1, 0x74, 0x30, 0x7a, 0x9b, 0x7d, 0x34, 0xc1, 0x3b, 0x0a, 0xe8, 0xbe,
	0x54, 0xe4, 0xe5, 0x1e, 0x3e, 0x7c, 0x06, 0xef, 0x1e, 0x9c, 0xc8, 0x18, 0x9c, 0x5e, 0x90, 0x9d,
	0xcf, 0x2c, 0x20, 0x8a, 0x63, 0xd4, 0x70, 0xd6, 0xe2, 0x16, 0xda, 0xe2, 0x10, 0x12, 0x2c, 0xce,
	0xa6, 0x3d, 0x3e, 0xa4, 0xb4, 0xcd, 0x89, 0xbf, 0xa5, 0xdd, 0x2b, 0x7d, 0x6d, 0x47, 0x7d, 0x89,
	0x2a, 0xff, 0x3f, 0xab, 0xa0, 0xaa, 0xc9, 0xab, 0x1b, 0x7b, 0x6b, 0x73, 0x25, 0x5f, 0x78, 0x68,
	0x27, 0xee, 0x6d, 0xad, 0xd9, 0xa7, 0x6f, 0x37, 0x9b, 0x2b, 0xce, 0x9b, 0xd2, 0x89, 0x28, 0xc8,
	0xbc, 0x1b, 0xff, 0xf2, 0xa6, 0x39, 0x7a, 0x2b, 0xca, 0x79, 0xcd, 0x5d, 0xfd, 0x07, 0xa8, 0xbe,
	0xdb, 0xd7, 0x89, 0xb3, 0x6d, 0xa0, 0x60, 0xad, 0xd4, 0x78, 0xee, 0xc6, 0xf4, 0x7a, 0xbe, 0xbc,
	0x94, 0x00, 0x70, 0x92, 0xf5, 0x8b, 0xa9, 0x18, 0xb0, 0xcc, 0x42, 0xb9, 0x34, 0x8e, 0xc1, 0x0f,
	0xf0, 0x92, 0x3a, 0x7a, 0x62, 0x80, 0x2d, 0x40, 0x76, 0xb4, 0x38, 0xaf, 0x05, 0xf4, 0x5b, 0xe3,
	0x15, 0xf3, 0xb0, 0x62, 0xf7, 0x5b, 0x6c, 0x6d, 0x2e, 0x9f, 0xce, 0xff, 0x3a, 0xa5, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x64, 0x07, 0x54, 0x84, 0x2b, 0x0a, 0x00, 0x00,
}
