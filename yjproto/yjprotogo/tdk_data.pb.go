// Code generated by protoc-gen-go.
// source: tdk_data.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TdkDeskUserData struct {
	Userid           *uint32 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Headurl          *string `protobuf:"bytes,3,opt,name=headurl" json:"headurl,omitempty"`
	Score            *int32  `protobuf:"varint,4,opt,name=score" json:"score,omitempty"`
	Already          *bool   `protobuf:"varint,5,opt,name=already" json:"already,omitempty"`
	Fold             *bool   `protobuf:"varint,6,opt,name=fold" json:"fold,omitempty"`
	Isowner          *bool   `protobuf:"varint,7,opt,name=isowner" json:"isowner,omitempty"`
	Pos              *int32  `protobuf:"varint,8,opt,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TdkDeskUserData) Reset()                    { *m = TdkDeskUserData{} }
func (m *TdkDeskUserData) String() string            { return proto.CompactTextString(m) }
func (*TdkDeskUserData) ProtoMessage()               {}
func (*TdkDeskUserData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *TdkDeskUserData) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *TdkDeskUserData) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TdkDeskUserData) GetHeadurl() string {
	if m != nil && m.Headurl != nil {
		return *m.Headurl
	}
	return ""
}

func (m *TdkDeskUserData) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *TdkDeskUserData) GetAlready() bool {
	if m != nil && m.Already != nil {
		return *m.Already
	}
	return false
}

func (m *TdkDeskUserData) GetFold() bool {
	if m != nil && m.Fold != nil {
		return *m.Fold
	}
	return false
}

func (m *TdkDeskUserData) GetIsowner() bool {
	if m != nil && m.Isowner != nil {
		return *m.Isowner
	}
	return false
}

func (m *TdkDeskUserData) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

type TdkUserPokerData struct {
	Userid           *uint32  `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Pokerlist        []uint32 `protobuf:"varint,2,rep,name=pokerlist" json:"pokerlist,omitempty"`
	Score            *uint32  `protobuf:"varint,3,opt,name=score" json:"score,omitempty"`
	Borrow           *bool    `protobuf:"varint,4,opt,name=borrow" json:"borrow,omitempty"`
	Pubscore         *uint32  `protobuf:"varint,5,opt,name=pubscore" json:"pubscore,omitempty"`
	Fold             *bool    `protobuf:"varint,6,opt,name=fold" json:"fold,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TdkUserPokerData) Reset()                    { *m = TdkUserPokerData{} }
func (m *TdkUserPokerData) String() string            { return proto.CompactTextString(m) }
func (*TdkUserPokerData) ProtoMessage()               {}
func (*TdkUserPokerData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *TdkUserPokerData) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *TdkUserPokerData) GetPokerlist() []uint32 {
	if m != nil {
		return m.Pokerlist
	}
	return nil
}

func (m *TdkUserPokerData) GetScore() uint32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *TdkUserPokerData) GetBorrow() bool {
	if m != nil && m.Borrow != nil {
		return *m.Borrow
	}
	return false
}

func (m *TdkUserPokerData) GetPubscore() uint32 {
	if m != nil && m.Pubscore != nil {
		return *m.Pubscore
	}
	return 0
}

func (m *TdkUserPokerData) GetFold() bool {
	if m != nil && m.Fold != nil {
		return *m.Fold
	}
	return false
}

type TdkDisUserData struct {
	Userid           *uint32 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Winnum           *uint32 `protobuf:"varint,2,opt,name=winnum" json:"winnum,omitempty"`
	Languonum        *uint32 `protobuf:"varint,3,opt,name=languonum" json:"languonum,omitempty"`
	Baozinum         *uint32 `protobuf:"varint,4,opt,name=baozinum" json:"baozinum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TdkDisUserData) Reset()                    { *m = TdkDisUserData{} }
func (m *TdkDisUserData) String() string            { return proto.CompactTextString(m) }
func (*TdkDisUserData) ProtoMessage()               {}
func (*TdkDisUserData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *TdkDisUserData) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *TdkDisUserData) GetWinnum() uint32 {
	if m != nil && m.Winnum != nil {
		return *m.Winnum
	}
	return 0
}

func (m *TdkDisUserData) GetLanguonum() uint32 {
	if m != nil && m.Languonum != nil {
		return *m.Languonum
	}
	return 0
}

func (m *TdkDisUserData) GetBaozinum() uint32 {
	if m != nil && m.Baozinum != nil {
		return *m.Baozinum
	}
	return 0
}

type TdkDeskConfig struct {
	Roundcount       *uint32 `protobuf:"varint,1,opt,name=roundcount" json:"roundcount,omitempty"`
	Playercount      *uint32 `protobuf:"varint,2,opt,name=playercount" json:"playercount,omitempty"`
	Allinrate        *uint32 `protobuf:"varint,3,opt,name=allinrate" json:"allinrate,omitempty"`
	Basecoin         *uint32 `protobuf:"varint,4,opt,name=basecoin" json:"basecoin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TdkDeskConfig) Reset()                    { *m = TdkDeskConfig{} }
func (m *TdkDeskConfig) String() string            { return proto.CompactTextString(m) }
func (*TdkDeskConfig) ProtoMessage()               {}
func (*TdkDeskConfig) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *TdkDeskConfig) GetRoundcount() uint32 {
	if m != nil && m.Roundcount != nil {
		return *m.Roundcount
	}
	return 0
}

func (m *TdkDeskConfig) GetPlayercount() uint32 {
	if m != nil && m.Playercount != nil {
		return *m.Playercount
	}
	return 0
}

func (m *TdkDeskConfig) GetAllinrate() uint32 {
	if m != nil && m.Allinrate != nil {
		return *m.Allinrate
	}
	return 0
}

func (m *TdkDeskConfig) GetBasecoin() uint32 {
	if m != nil && m.Basecoin != nil {
		return *m.Basecoin
	}
	return 0
}

type TdkZhanJiData struct {
	Userid           *uint32 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Winnum           *uint32 `protobuf:"varint,2,opt,name=winnum" json:"winnum,omitempty"`
	Languonum        *uint32 `protobuf:"varint,3,opt,name=languonum" json:"languonum,omitempty"`
	Baozinum         *uint32 `protobuf:"varint,4,opt,name=baozinum" json:"baozinum,omitempty"`
	Score            *int32  `protobuf:"varint,5,opt,name=score" json:"score,omitempty"`
	Isowner          *bool   `protobuf:"varint,6,opt,name=isowner" json:"isowner,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TdkZhanJiData) Reset()                    { *m = TdkZhanJiData{} }
func (m *TdkZhanJiData) String() string            { return proto.CompactTextString(m) }
func (*TdkZhanJiData) ProtoMessage()               {}
func (*TdkZhanJiData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *TdkZhanJiData) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *TdkZhanJiData) GetWinnum() uint32 {
	if m != nil && m.Winnum != nil {
		return *m.Winnum
	}
	return 0
}

func (m *TdkZhanJiData) GetLanguonum() uint32 {
	if m != nil && m.Languonum != nil {
		return *m.Languonum
	}
	return 0
}

func (m *TdkZhanJiData) GetBaozinum() uint32 {
	if m != nil && m.Baozinum != nil {
		return *m.Baozinum
	}
	return 0
}

func (m *TdkZhanJiData) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *TdkZhanJiData) GetIsowner() bool {
	if m != nil && m.Isowner != nil {
		return *m.Isowner
	}
	return false
}

func init() {
	proto.RegisterType((*TdkDeskUserData)(nil), "yjprotogo.TdkDeskUserData")
	proto.RegisterType((*TdkUserPokerData)(nil), "yjprotogo.TdkUserPokerData")
	proto.RegisterType((*TdkDisUserData)(nil), "yjprotogo.TdkDisUserData")
	proto.RegisterType((*TdkDeskConfig)(nil), "yjprotogo.TdkDeskConfig")
	proto.RegisterType((*TdkZhanJiData)(nil), "yjprotogo.TdkZhanJiData")
}

var fileDescriptor14 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x41, 0xae, 0xd3, 0x30,
	0x10, 0x86, 0xe5, 0x97, 0x26, 0x2f, 0x1d, 0x94, 0xc7, 0x93, 0x85, 0x2a, 0x0b, 0x21, 0x14, 0x65,
	0x95, 0x15, 0x97, 0xa0, 0x2b, 0x56, 0x28, 0x2a, 0x1b, 0x36, 0xc8, 0xad, 0xdd, 0xd6, 0xc4, 0xf5,
	0x44, 0x76, 0xa2, 0xaa, 0xbd, 0x02, 0xf7, 0x40, 0xe2, 0x24, 0x5c, 0x0b, 0xd9, 0x4e, 0x48, 0x41,
	0x48, 0xdd, 0xb1, 0x9b, 0xff, 0xf7, 0x8c, 0xfd, 0xcd, 0x78, 0xe0, 0xa9, 0x17, 0xed, 0x17, 0xc1,
	0x7b, 0xfe, 0xae, 0xb3, 0xd8, 0x23, 0x5d, 0x5e, 0xbe, 0x86, 0xe0, 0x80, 0xd5, 0x4f, 0x02, 0x2f,
	0x37, 0xa2, 0x5d, 0x4b, 0xd7, 0x7e, 0x72, 0xd2, 0xae, 0x79, 0xcf, 0xe9, 0x0a, 0xb2, 0xc1, 0x49,
	0xab, 0x04, 0x23, 0x25, 0xa9, 0x8b, 0x66, 0x54, 0x94, 0xc2, 0xc2, 0xf0, 0x93, 0x64, 0x0f, 0x25,
	0xa9, 0x97, 0x4d, 0x88, 0x29, 0x83, 0xc7, 0xa3, 0xe4, 0x62, 0xb0, 0x9a, 0x25, 0xc1, 0x9e, 0x24,
	0x7d, 0x05, 0xa9, 0xdb, 0xa1, 0x95, 0x6c, 0x51, 0x92, 0x3a, 0x6d, 0xa2, 0xf0, 0xf9, 0x5c, 0x5b,
	0xc9, 0xc5, 0x85, 0xa5, 0x25, 0xa9, 0xf3, 0x66, 0x92, 0xfe, 0xf6, 0x3d, 0x6a, 0xc1, 0xb2, 0x60,
	0x87, 0xd8, 0x67, 0x2b, 0x87, 0x67, 0x23, 0x2d, 0x7b, 0x8c, 0xd9, 0xa3, 0xa4, 0xcf, 0x90, 0x74,
	0xe8, 0x58, 0x1e, 0xee, 0xf6, 0x61, 0xf5, 0x9d, 0xc0, 0xf3, 0x46, 0x84, 0x2e, 0x3e, 0x62, 0x7b,
	0xa7, 0x95, 0x37, 0xb0, 0xec, 0x7c, 0x92, 0x56, 0xae, 0x67, 0x0f, 0x65, 0x52, 0x17, 0xcd, 0x6c,
	0xcc, 0xe8, 0x49, 0x28, 0x1a, 0xd1, 0x57, 0x90, 0x6d, 0xd1, 0x5a, 0x3c, 0x87, 0x8e, 0xf2, 0x66,
	0x54, 0xf4, 0x35, 0xe4, 0xdd, 0xb0, 0x8d, 0x05, 0x69, 0x28, 0xf8, 0xad, 0xff, 0xd5, 0x54, 0x75,
	0x85, 0x27, 0x3f, 0x71, 0xe5, 0xee, 0x0e, 0x7c, 0x05, 0xd9, 0x59, 0x19, 0x33, 0x9c, 0xc2, 0xc8,
	0x8b, 0x66, 0x54, 0x9e, 0x5e, 0x73, 0x73, 0x18, 0xd0, 0x1f, 0x45, 0xc6, 0xd9, 0xf0, 0x3c, 0x5b,
	0x8e, 0x57, 0xe5, 0x0f, 0x17, 0x91, 0x67, 0xd2, 0xd5, 0x37, 0x02, 0xc5, 0xf8, 0xdd, 0xef, 0xd1,
	0xec, 0xd5, 0x81, 0xbe, 0x05, 0xb0, 0x38, 0x18, 0xb1, 0xc3, 0xc1, 0xf4, 0xe3, 0xfb, 0x37, 0x0e,
	0x2d, 0xe1, 0x45, 0xa7, 0xf9, 0x45, 0xda, 0x98, 0x10, 0x41, 0x6e, 0x2d, 0x4f, 0xc3, 0xb5, 0x56,
	0xc6, 0xf2, 0x7e, 0x9a, 0xd8, 0x6c, 0x44, 0x1a, 0x27, 0x77, 0xa8, 0xcc, 0x4c, 0x13, 0x75, 0xf5,
	0x23, 0xd2, 0x7c, 0x3e, 0x72, 0xf3, 0x41, 0xfd, 0xdf, 0x49, 0xcc, 0x7f, 0x9c, 0xfe, 0xb5, 0x9e,
	0xd3, 0xc2, 0x65, 0x7f, 0x2c, 0xdc, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x0b, 0xf4, 0xb4,
	0x44, 0x03, 0x00, 0x00,
}