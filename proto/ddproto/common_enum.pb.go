// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_enum.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏的id 端口号
type CommonEnumGame int32

const (
	CommonEnumGame_GID_SRC                    CommonEnumGame = 0
	CommonEnumGame_GID_TEXAS                  CommonEnumGame = 1
	CommonEnumGame_GID_MAHJONG                CommonEnumGame = 2
	CommonEnumGame_GID_DDZ                    CommonEnumGame = 3
	CommonEnumGame_GID_ZJH                    CommonEnumGame = 4
	CommonEnumGame_GID_HALL                   CommonEnumGame = 5
	CommonEnumGame_GID_PDK                    CommonEnumGame = 7
	CommonEnumGame_GID_ZXZ                    CommonEnumGame = 8
	CommonEnumGame_GID_MJBAISHAN              CommonEnumGame = 9
	CommonEnumGame_GID_MJCHANGCHUN            CommonEnumGame = 10
	CommonEnumGame_GID_NIUNIUJINGDIAN         CommonEnumGame = 11
	CommonEnumGame_GID_BANTUOZI               CommonEnumGame = 12
	CommonEnumGame_GID_DOUDIZHUERREN          CommonEnumGame = 13
	CommonEnumGame_GIDDOUDIZHUJINGDIANDONGBEI CommonEnumGame = 14
	CommonEnumGame_GID_TIANDAKENG             CommonEnumGame = 15
	CommonEnumGame_GID_MJ_SONGJIANGHE         CommonEnumGame = 16
	CommonEnumGame_GID_BAIRENNIUNIU           CommonEnumGame = 17
	CommonEnumGame_GID_ZHIZUNWUZHANG          CommonEnumGame = 18
	CommonEnumGame_GID_MJ_YIBIN               CommonEnumGame = 19
	CommonEnumGame_GID_MJ_PENGZHOU            CommonEnumGame = 20
	CommonEnumGame_GID_SANDAYI                CommonEnumGame = 21
	CommonEnumGame_GID_ZHUANZHUAN             CommonEnumGame = 22
	CommonEnumGame_GID_PAOHUZI                CommonEnumGame = 23
	CommonEnumGame_GID_MJ_HAINAN              CommonEnumGame = 24
	CommonEnumGame_GID_PAOYAO                 CommonEnumGame = 25
	CommonEnumGame_GID_CHANGBAI               CommonEnumGame = 26
	CommonEnumGame_GID_ZHADAN                 CommonEnumGame = 27
	CommonEnumGame_GID_MJ_SHENQI              CommonEnumGame = 28
	CommonEnumGame_GID_PHZ_SHAOYANGBOPI       CommonEnumGame = 29
	CommonEnumGame_GID_MJ_CHANGSHA            CommonEnumGame = 30
	CommonEnumGame_GID_MJ_ZIGONG              CommonEnumGame = 31
)

var CommonEnumGame_name = map[int32]string{
	0:  "GID_SRC",
	1:  "GID_TEXAS",
	2:  "GID_MAHJONG",
	3:  "GID_DDZ",
	4:  "GID_ZJH",
	5:  "GID_HALL",
	7:  "GID_PDK",
	8:  "GID_ZXZ",
	9:  "GID_MJBAISHAN",
	10: "GID_MJCHANGCHUN",
	11: "GID_NIUNIUJINGDIAN",
	12: "GID_BANTUOZI",
	13: "GID_DOUDIZHUERREN",
	14: "GIDDOUDIZHUJINGDIANDONGBEI",
	15: "GID_TIANDAKENG",
	16: "GID_MJ_SONGJIANGHE",
	17: "GID_BAIRENNIUNIU",
	18: "GID_ZHIZUNWUZHANG",
	19: "GID_MJ_YIBIN",
	20: "GID_MJ_PENGZHOU",
	21: "GID_SANDAYI",
	22: "GID_ZHUANZHUAN",
	23: "GID_PAOHUZI",
	24: "GID_MJ_HAINAN",
	25: "GID_PAOYAO",
	26: "GID_CHANGBAI",
	27: "GID_ZHADAN",
	28: "GID_MJ_SHENQI",
	29: "GID_PHZ_SHAOYANGBOPI",
	30: "GID_MJ_CHANGSHA",
	31: "GID_MJ_ZIGONG",
}
var CommonEnumGame_value = map[string]int32{
	"GID_SRC":                    0,
	"GID_TEXAS":                  1,
	"GID_MAHJONG":                2,
	"GID_DDZ":                    3,
	"GID_ZJH":                    4,
	"GID_HALL":                   5,
	"GID_PDK":                    7,
	"GID_ZXZ":                    8,
	"GID_MJBAISHAN":              9,
	"GID_MJCHANGCHUN":            10,
	"GID_NIUNIUJINGDIAN":         11,
	"GID_BANTUOZI":               12,
	"GID_DOUDIZHUERREN":          13,
	"GIDDOUDIZHUJINGDIANDONGBEI": 14,
	"GID_TIANDAKENG":             15,
	"GID_MJ_SONGJIANGHE":         16,
	"GID_BAIRENNIUNIU":           17,
	"GID_ZHIZUNWUZHANG":          18,
	"GID_MJ_YIBIN":               19,
	"GID_MJ_PENGZHOU":            20,
	"GID_SANDAYI":                21,
	"GID_ZHUANZHUAN":             22,
	"GID_PAOHUZI":                23,
	"GID_MJ_HAINAN":              24,
	"GID_PAOYAO":                 25,
	"GID_CHANGBAI":               26,
	"GID_ZHADAN":                 27,
	"GID_MJ_SHENQI":              28,
	"GID_PHZ_SHAOYANGBOPI":       29,
	"GID_MJ_CHANGSHA":            30,
	"GID_MJ_ZIGONG":              31,
}

func (x CommonEnumGame) Enum() *CommonEnumGame {
	p := new(CommonEnumGame)
	*p = x
	return p
}
func (x CommonEnumGame) String() string {
	return proto.EnumName(CommonEnumGame_name, int32(x))
}
func (x *CommonEnumGame) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumGame_value, data, "CommonEnumGame")
	if err != nil {
		return err
	}
	*x = CommonEnumGame(value)
	return nil
}
func (CommonEnumGame) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

// 房间的类型
type COMMON_ENUM_ROOMTYPE int32

const (
	COMMON_ENUM_ROOMTYPE_DESK_FRIEND COMMON_ENUM_ROOMTYPE = 1
	COMMON_ENUM_ROOMTYPE_DESK_COIN   COMMON_ENUM_ROOMTYPE = 2
)

var COMMON_ENUM_ROOMTYPE_name = map[int32]string{
	1: "DESK_FRIEND",
	2: "DESK_COIN",
}
var COMMON_ENUM_ROOMTYPE_value = map[string]int32{
	"DESK_FRIEND": 1,
	"DESK_COIN":   2,
}

func (x COMMON_ENUM_ROOMTYPE) Enum() *COMMON_ENUM_ROOMTYPE {
	p := new(COMMON_ENUM_ROOMTYPE)
	*p = x
	return p
}
func (x COMMON_ENUM_ROOMTYPE) String() string {
	return proto.EnumName(COMMON_ENUM_ROOMTYPE_name, int32(x))
}
func (x *COMMON_ENUM_ROOMTYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_ROOMTYPE_value, data, "COMMON_ENUM_ROOMTYPE")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_ROOMTYPE(value)
	return nil
}
func (COMMON_ENUM_ROOMTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

// 游戏的状态
type COMMON_ENUM_GAMESTATUS int32

const (
	COMMON_ENUM_GAMESTATUS_NOGAME COMMON_ENUM_GAMESTATUS = 1
	COMMON_ENUM_GAMESTATUS_GAMING COMMON_ENUM_GAMESTATUS = 2
)

var COMMON_ENUM_GAMESTATUS_name = map[int32]string{
	1: "NOGAME",
	2: "GAMING",
}
var COMMON_ENUM_GAMESTATUS_value = map[string]int32{
	"NOGAME": 1,
	"GAMING": 2,
}

func (x COMMON_ENUM_GAMESTATUS) Enum() *COMMON_ENUM_GAMESTATUS {
	p := new(COMMON_ENUM_GAMESTATUS)
	*p = x
	return p
}
func (x COMMON_ENUM_GAMESTATUS) String() string {
	return proto.EnumName(COMMON_ENUM_GAMESTATUS_name, int32(x))
}
func (x *COMMON_ENUM_GAMESTATUS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_GAMESTATUS_value, data, "COMMON_ENUM_GAMESTATUS")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_GAMESTATUS(value)
	return nil
}
func (COMMON_ENUM_GAMESTATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

// 发布版本
type COMMON_ENUM_RELEASETAG int32

const (
	COMMON_ENUM_RELEASETAG_R_PRO          COMMON_ENUM_RELEASETAG = 1
	COMMON_ENUM_RELEASETAG_R_APPLE_VERIFY COMMON_ENUM_RELEASETAG = 2
)

var COMMON_ENUM_RELEASETAG_name = map[int32]string{
	1: "R_PRO",
	2: "R_APPLE_VERIFY",
}
var COMMON_ENUM_RELEASETAG_value = map[string]int32{
	"R_PRO":          1,
	"R_APPLE_VERIFY": 2,
}

func (x COMMON_ENUM_RELEASETAG) Enum() *COMMON_ENUM_RELEASETAG {
	p := new(COMMON_ENUM_RELEASETAG)
	*p = x
	return p
}
func (x COMMON_ENUM_RELEASETAG) String() string {
	return proto.EnumName(COMMON_ENUM_RELEASETAG_name, int32(x))
}
func (x *COMMON_ENUM_RELEASETAG) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_RELEASETAG_value, data, "COMMON_ENUM_RELEASETAG")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_RELEASETAG(value)
	return nil
}
func (COMMON_ENUM_RELEASETAG) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

// 踢人的踢出类型
type COMMON_ENUM_KICKOUT int32

const (
	COMMON_ENUM_KICKOUT_K_COINUNMATCHED              COMMON_ENUM_KICKOUT = 1
	COMMON_ENUM_KICKOUT_K_SYSTEM                     COMMON_ENUM_KICKOUT = 2
	COMMON_ENUM_KICKOUT_K_OFFLINE                    COMMON_ENUM_KICKOUT = 3
	COMMON_ENUM_KICKOUT_K_GAME_BLOCKED               COMMON_ENUM_KICKOUT = 4
	COMMON_ENUM_KICKOUT_K_TIMEOUT_NOTREADY_LOTTERY   COMMON_ENUM_KICKOUT = 5
	COMMON_ENUM_KICKOUT_K_TIMEOUT_NOTREADY_ENTERDESK COMMON_ENUM_KICKOUT = 6
	COMMON_ENUM_KICKOUT_K_OWNER                      COMMON_ENUM_KICKOUT = 7
)

var COMMON_ENUM_KICKOUT_name = map[int32]string{
	1: "K_COINUNMATCHED",
	2: "K_SYSTEM",
	3: "K_OFFLINE",
	4: "K_GAME_BLOCKED",
	5: "K_TIMEOUT_NOTREADY_LOTTERY",
	6: "K_TIMEOUT_NOTREADY_ENTERDESK",
	7: "K_OWNER",
}
var COMMON_ENUM_KICKOUT_value = map[string]int32{
	"K_COINUNMATCHED":              1,
	"K_SYSTEM":                     2,
	"K_OFFLINE":                    3,
	"K_GAME_BLOCKED":               4,
	"K_TIMEOUT_NOTREADY_LOTTERY":   5,
	"K_TIMEOUT_NOTREADY_ENTERDESK": 6,
	"K_OWNER":                      7,
}

func (x COMMON_ENUM_KICKOUT) Enum() *COMMON_ENUM_KICKOUT {
	p := new(COMMON_ENUM_KICKOUT)
	*p = x
	return p
}
func (x COMMON_ENUM_KICKOUT) String() string {
	return proto.EnumName(COMMON_ENUM_KICKOUT_name, int32(x))
}
func (x *COMMON_ENUM_KICKOUT) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_KICKOUT_value, data, "COMMON_ENUM_KICKOUT")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_KICKOUT(value)
	return nil
}
func (COMMON_ENUM_KICKOUT) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

// 玩家申请解散房间的操作类型
type COMMON_ENUM_APPLYDISSOLVE int32

const (
	COMMON_ENUM_APPLYDISSOLVE_AD_AGREE  COMMON_ENUM_APPLYDISSOLVE = 1
	COMMON_ENUM_APPLYDISSOLVE_AD_REFUSE COMMON_ENUM_APPLYDISSOLVE = 2
	COMMON_ENUM_APPLYDISSOLVE_AD_NOACT  COMMON_ENUM_APPLYDISSOLVE = 3
)

var COMMON_ENUM_APPLYDISSOLVE_name = map[int32]string{
	1: "AD_AGREE",
	2: "AD_REFUSE",
	3: "AD_NOACT",
}
var COMMON_ENUM_APPLYDISSOLVE_value = map[string]int32{
	"AD_AGREE":  1,
	"AD_REFUSE": 2,
	"AD_NOACT":  3,
}

func (x COMMON_ENUM_APPLYDISSOLVE) Enum() *COMMON_ENUM_APPLYDISSOLVE {
	p := new(COMMON_ENUM_APPLYDISSOLVE)
	*p = x
	return p
}
func (x COMMON_ENUM_APPLYDISSOLVE) String() string {
	return proto.EnumName(COMMON_ENUM_APPLYDISSOLVE_name, int32(x))
}
func (x *COMMON_ENUM_APPLYDISSOLVE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_APPLYDISSOLVE_value, data, "COMMON_ENUM_APPLYDISSOLVE")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_APPLYDISSOLVE(value)
	return nil
}
func (COMMON_ENUM_APPLYDISSOLVE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

// 按钮的类型
type BTN_TYPE int32

const (
	BTN_TYPE_BTNTYPE_NEWUSERBUTTON     BTN_TYPE = 1
	BTN_TYPE_BTNTYPE_SHARELINKTIMELINE BTN_TYPE = 2
	BTN_TYPE_REDBAGSHAR                BTN_TYPE = 3
)

var BTN_TYPE_name = map[int32]string{
	1: "BTNTYPE_NEWUSERBUTTON",
	2: "BTNTYPE_SHARELINKTIMELINE",
	3: "REDBAGSHAR",
}
var BTN_TYPE_value = map[string]int32{
	"BTNTYPE_NEWUSERBUTTON":     1,
	"BTNTYPE_SHARELINKTIMELINE": 2,
	"REDBAGSHAR":                3,
}

func (x BTN_TYPE) Enum() *BTN_TYPE {
	p := new(BTN_TYPE)
	*p = x
	return p
}
func (x BTN_TYPE) String() string {
	return proto.EnumName(BTN_TYPE_name, int32(x))
}
func (x *BTN_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BTN_TYPE_value, data, "BTN_TYPE")
	if err != nil {
		return err
	}
	*x = BTN_TYPE(value)
	return nil
}
func (BTN_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

// 进入金币场房间的错误类型
type COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM int32

const (
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_UNDER_MIN       COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = -100
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_OVER_LIMIT      COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = -102
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_UNDER_LIMIT     COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = -103
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_OTHER_LV_DESK_GAMING COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = -104
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_NONE_DESK_AVAILABLE  COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = -105
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_OTHER                COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = -106
)

var COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name = map[int32]string{
	-100: "ERROR_EC_COIN_UNDER_MIN",
	-102: "ERROR_EC_COIN_OVER_LIMIT",
	-103: "ERROR_EC_COIN_UNDER_LIMIT",
	-104: "ERROR_EC_OTHER_LV_DESK_GAMING",
	-105: "ERROR_EC_NONE_DESK_AVAILABLE",
	-106: "ERROR_EC_OTHER",
}
var COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value = map[string]int32{
	"ERROR_EC_COIN_UNDER_MIN":       -100,
	"ERROR_EC_COIN_OVER_LIMIT":      -102,
	"ERROR_EC_COIN_UNDER_LIMIT":     -103,
	"ERROR_EC_OTHER_LV_DESK_GAMING": -104,
	"ERROR_EC_NONE_DESK_AVAILABLE":  -105,
	"ERROR_EC_OTHER":                -106,
}

func (x COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) Enum() *COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM {
	p := new(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM)
	*p = x
	return p
}
func (x COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) String() string {
	return proto.EnumName(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name, int32(x))
}
func (x *COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value, data, "COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM(value)
	return nil
}
func (COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{7}
}

// 房费结算类型
type COMMON_ENUM_ROOMCARD_BILL_TYPE int32

const (
	COMMON_ENUM_ROOMCARD_BILL_TYPE_OWNER_PAY     COMMON_ENUM_ROOMCARD_BILL_TYPE = 1
	COMMON_ENUM_ROOMCARD_BILL_TYPE_AA_PAY        COMMON_ENUM_ROOMCARD_BILL_TYPE = 2
	COMMON_ENUM_ROOMCARD_BILL_TYPE_BIG_WINER_PAY COMMON_ENUM_ROOMCARD_BILL_TYPE = 3
)

var COMMON_ENUM_ROOMCARD_BILL_TYPE_name = map[int32]string{
	1: "OWNER_PAY",
	2: "AA_PAY",
	3: "BIG_WINER_PAY",
}
var COMMON_ENUM_ROOMCARD_BILL_TYPE_value = map[string]int32{
	"OWNER_PAY":     1,
	"AA_PAY":        2,
	"BIG_WINER_PAY": 3,
}

func (x COMMON_ENUM_ROOMCARD_BILL_TYPE) Enum() *COMMON_ENUM_ROOMCARD_BILL_TYPE {
	p := new(COMMON_ENUM_ROOMCARD_BILL_TYPE)
	*p = x
	return p
}
func (x COMMON_ENUM_ROOMCARD_BILL_TYPE) String() string {
	return proto.EnumName(COMMON_ENUM_ROOMCARD_BILL_TYPE_name, int32(x))
}
func (x *COMMON_ENUM_ROOMCARD_BILL_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_ROOMCARD_BILL_TYPE_value, data, "COMMON_ENUM_ROOMCARD_BILL_TYPE")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_ROOMCARD_BILL_TYPE(value)
	return nil
}
func (COMMON_ENUM_ROOMCARD_BILL_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{8}
}

func init() {
	proto.RegisterEnum("ddproto.CommonEnumGame", CommonEnumGame_name, CommonEnumGame_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_ROOMTYPE", COMMON_ENUM_ROOMTYPE_name, COMMON_ENUM_ROOMTYPE_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_GAMESTATUS", COMMON_ENUM_GAMESTATUS_name, COMMON_ENUM_GAMESTATUS_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_RELEASETAG", COMMON_ENUM_RELEASETAG_name, COMMON_ENUM_RELEASETAG_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_KICKOUT", COMMON_ENUM_KICKOUT_name, COMMON_ENUM_KICKOUT_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_APPLYDISSOLVE", COMMON_ENUM_APPLYDISSOLVE_name, COMMON_ENUM_APPLYDISSOLVE_value)
	proto.RegisterEnum("ddproto.BTN_TYPE", BTN_TYPE_name, BTN_TYPE_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM", COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name, COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE", COMMON_ENUM_ROOMCARD_BILL_TYPE_name, COMMON_ENUM_ROOMCARD_BILL_TYPE_value)
}

func init() { proto.RegisterFile("common_enum.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 835 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x72, 0xe4, 0x34,
	0x10, 0x66, 0x26, 0x9b, 0x4d, 0xa2, 0xfc, 0x75, 0x94, 0x9f, 0xcd, 0xec, 0x26, 0x81, 0xc3, 0x42,
	0x15, 0x3e, 0x50, 0x9c, 0xe0, 0xdc, 0xb6, 0x7b, 0x2c, 0x8d, 0xed, 0xd6, 0x20, 0x4b, 0x93, 0xb5,
	0x2f, 0x2a, 0x8a, 0xdd, 0xe2, 0x94, 0x0d, 0x45, 0xc1, 0x73, 0x70, 0xe3, 0xaf, 0x78, 0x02, 0x1e,
	0x89, 0x97, 0x81, 0x92, 0x1d, 0x67, 0x26, 0xd4, 0xce, 0x61, 0x4a, 0xdd, 0xfd, 0xe9, 0xeb, 0xaf,
	0x5b, 0x9f, 0xc5, 0xc9, 0x77, 0xf7, 0x77, 0x77, 0xf7, 0xef, 0xc3, 0xbb, 0xf7, 0x3f, 0xdf, 0x7d,
	0xf1, 0xc3, 0x8f, 0xf7, 0x3f, 0xdd, 0xcb, 0x9d, 0xb7, 0x6f, 0xfb, 0x43, 0xf2, 0xcf, 0x33, 0x01,
	0x1b, 0xe5, 0xf0, 0xfd, 0xb7, 0x77, 0xef, 0xe4, 0xbe, 0xd8, 0x29, 0x74, 0x1e, 0x1a, 0x9b, 0xc1,
	0x47, 0xf2, 0x50, 0xec, 0xc5, 0xc0, 0xd1, 0x1b, 0x6c, 0x60, 0x22, 0x8f, 0xc5, 0x7e, 0x0c, 0x6b,
	0x54, 0x0b, 0xc3, 0x05, 0x4c, 0x47, 0x70, 0x9e, 0x77, 0xb0, 0x35, 0x06, 0xdd, 0x42, 0xc1, 0x33,
	0x79, 0x20, 0x76, 0x63, 0xa0, 0xb0, 0xaa, 0x60, 0x7b, 0x2c, 0x2d, 0xf3, 0x12, 0x76, 0x1e, 0x71,
	0x6f, 0x3a, 0xd8, 0x95, 0x27, 0xe2, 0xb0, 0xa7, 0x5c, 0xa4, 0xa8, 0x1b, 0x85, 0x0c, 0x7b, 0xf2,
	0x54, 0x1c, 0x0f, 0xa9, 0x4c, 0x21, 0x17, 0x99, 0xf2, 0x0c, 0x42, 0x5e, 0x08, 0x19, 0x93, 0xac,
	0x3d, 0x6b, 0xbf, 0xd0, 0x5c, 0xe4, 0x1a, 0x19, 0xf6, 0x25, 0x88, 0x83, 0x98, 0x4f, 0x91, 0x9d,
	0x37, 0x9d, 0x86, 0x03, 0x79, 0x2e, 0x4e, 0x7a, 0x4d, 0xc6, 0xe7, 0xba, 0x53, 0x9e, 0xac, 0x25,
	0x86, 0x43, 0x79, 0x23, 0x5e, 0x16, 0x3a, 0x1f, 0xb3, 0x23, 0x43, 0x6e, 0xb8, 0x48, 0x49, 0xc3,
	0x91, 0x94, 0xe2, 0xa8, 0x1f, 0x35, 0x26, 0xb1, 0x24, 0x2e, 0xe0, 0x78, 0x6c, 0x5a, 0x2f, 0x42,
	0x63, 0xb8, 0x58, 0x68, 0xe4, 0x42, 0x11, 0x80, 0x3c, 0x13, 0x30, 0x34, 0xd5, 0x96, 0x78, 0x90,
	0x04, 0x27, 0x63, 0xe3, 0x4e, 0xe9, 0xce, 0xf3, 0xad, 0xef, 0xa2, 0x7a, 0x90, 0xa3, 0xc2, 0x7a,
	0x11, 0x5a, 0x9d, 0x6a, 0x86, 0xd3, 0xf5, 0x80, 0x61, 0x49, 0x5c, 0x74, 0xca, 0x78, 0x38, 0x1b,
	0x77, 0xdb, 0xc4, 0xf6, 0xad, 0x86, 0xf3, 0x51, 0x50, 0xa7, 0x3c, 0x72, 0xff, 0x07, 0x17, 0x23,
	0x68, 0x89, 0x46, 0xf9, 0x4e, 0xc3, 0x8b, 0xf5, 0xfa, 0x82, 0x42, 0xcd, 0xc8, 0x70, 0x29, 0x8f,
	0x84, 0x78, 0xc0, 0xb4, 0x68, 0x60, 0x36, 0xf6, 0xef, 0x97, 0x99, 0xa2, 0x86, 0x97, 0x23, 0xa2,
	0x53, 0x98, 0x23, 0xc3, 0xab, 0x0d, 0x92, 0x46, 0x11, 0x7f, 0xa3, 0xe1, 0x4a, 0x5e, 0x8a, 0xb3,
	0x9e, 0x44, 0x75, 0xa1, 0x51, 0x91, 0x89, 0x8b, 0xd4, 0x2c, 0x35, 0x5c, 0x6f, 0x88, 0xef, 0x19,
	0x1b, 0x85, 0x70, 0xb3, 0xc1, 0xd0, 0xe9, 0x22, 0x5a, 0xe3, 0xe3, 0xe4, 0x2b, 0x71, 0x96, 0x99,
	0xba, 0x36, 0x1c, 0x88, 0x7d, 0x1d, 0xac, 0x31, 0xb5, 0x6b, 0x97, 0x14, 0x47, 0xc8, 0xa9, 0x29,
	0xc3, 0xdc, 0x6a, 0xe2, 0x1c, 0x26, 0xd1, 0x63, 0x7d, 0x22, 0x33, 0x9a, 0x61, 0x9a, 0x7c, 0x29,
	0x2e, 0x36, 0xef, 0x15, 0x58, 0x53, 0xe3, 0xd0, 0xf9, 0x46, 0x0a, 0xf1, 0x9c, 0x4d, 0x8c, 0x61,
	0x12, 0xcf, 0x05, 0xd6, 0x3a, 0x9a, 0x30, 0xf9, 0xfa, 0xe9, 0x0d, 0x4b, 0x15, 0x61, 0x43, 0x0e,
	0x0b, 0xb9, 0x27, 0xb6, 0x6d, 0x58, 0x5a, 0x03, 0x93, 0xb8, 0x4d, 0x1b, 0x70, 0xb9, 0xac, 0x28,
	0xac, 0xc8, 0xea, 0x79, 0x0b, 0xd3, 0xe4, 0xef, 0x89, 0x38, 0xdd, 0xbc, 0x59, 0xea, 0xac, 0x34,
	0xde, 0xc5, 0x11, 0x07, 0x39, 0x9e, 0x6b, 0x74, 0x99, 0xa2, 0x28, 0xf3, 0x40, 0xec, 0x96, 0xa1,
	0x69, 0x1b, 0x47, 0x35, 0x4c, 0xa3, 0xe8, 0x32, 0x98, 0xf9, 0xbc, 0xd2, 0x4c, 0xb0, 0x15, 0xd9,
	0xcb, 0x5e, 0x6a, 0x48, 0x2b, 0x93, 0x95, 0x94, 0xc3, 0xb3, 0x68, 0xb8, 0x32, 0x38, 0x5d, 0x93,
	0xf1, 0x2e, 0xb0, 0x71, 0x96, 0x30, 0x6f, 0x43, 0x65, 0x9c, 0x23, 0xdb, 0xc2, 0xb6, 0xfc, 0x44,
	0x5c, 0x7d, 0xa0, 0x4e, 0xec, 0xc8, 0xc6, 0x7d, 0xc0, 0xf3, 0xf8, 0xa1, 0x94, 0xc1, 0xdc, 0x32,
	0x59, 0xd8, 0x49, 0xe6, 0x62, 0xb6, 0xa9, 0x35, 0x8e, 0xd2, 0xe6, 0xba, 0x69, 0x4c, 0xb5, 0xa2,
	0x28, 0x0e, 0xf3, 0x80, 0x85, 0x25, 0x1a, 0x36, 0x8a, 0x79, 0xb0, 0x34, 0xf7, 0x0d, 0xc1, 0xf4,
	0xa1, 0xc8, 0x06, 0x33, 0x07, 0x5b, 0x89, 0x13, 0xbb, 0xa9, 0xe3, 0xd0, 0xbf, 0xc5, 0x4c, 0x9c,
	0xa7, 0x8e, 0xe3, 0x31, 0x30, 0xdd, 0xfa, 0x86, 0x6c, 0xea, 0x9d, 0x33, 0x0c, 0x13, 0x79, 0x2d,
	0x66, 0x63, 0xa9, 0x51, 0x68, 0xa9, 0xd2, 0x5c, 0x46, 0xb1, 0xfd, 0xc0, 0xd3, 0x68, 0x21, 0x4b,
	0x79, 0x8a, 0xf1, 0xfd, 0x2d, 0x6c, 0x25, 0xbf, 0x4c, 0xc5, 0xeb, 0x4d, 0x79, 0x64, 0xad, 0xb1,
	0x7d, 0x93, 0x61, 0xa2, 0xb8, 0xcd, 0x68, 0x01, 0xf9, 0x5a, 0xbc, 0x18, 0x6a, 0x94, 0xf5, 0x2b,
	0x0e, 0x9e, 0x73, 0xb2, 0xa1, 0xd6, 0x0c, 0x7f, 0xfd, 0xfb, 0xf0, 0x9b, 0xc8, 0x4f, 0xc5, 0xe5,
	0x53, 0x94, 0x59, 0x91, 0x0d, 0x95, 0xae, 0xb5, 0x83, 0x3f, 0xd7, 0xb0, 0xcf, 0xc4, 0xec, 0x43,
	0x64, 0x03, 0xee, 0x8f, 0x35, 0x2e, 0x11, 0xd7, 0x8f, 0x38, 0xe3, 0x54, 0x84, 0xac, 0x42, 0x6f,
	0xba, 0x07, 0x13, 0xfd, 0xbe, 0xc6, 0x7e, 0x2e, 0xae, 0x1e, 0xb1, 0x6c, 0x98, 0x06, 0x1c, 0xae,
	0x50, 0x57, 0x98, 0x56, 0x04, 0xbf, 0xad, 0xa1, 0xaf, 0xc4, 0xd1, 0x53, 0x5a, 0xf8, 0xf5, 0xb1,
	0x98, 0xb0, 0xb8, 0xf9, 0xbf, 0xff, 0x33, 0xb4, 0x79, 0x48, 0x75, 0x55, 0x0d, 0xdb, 0x3f, 0x14,
	0x7b, 0xfd, 0xe3, 0x86, 0x25, 0xb6, 0x83, 0xa5, 0x11, 0xfb, 0xf3, 0x34, 0x7e, 0x4f, 0xa9, 0x2e,
	0xc2, 0xad, 0x1e, 0xcb, 0x5b, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xd0, 0x2f, 0xb6, 0xc9,
	0x05, 0x00, 0x00,
}
