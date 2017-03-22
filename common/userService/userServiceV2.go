package userService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/sessionService"
)

type UREDIS uint32

func (s UREDIS) GetMainSession() *ddproto.GameSession {
	return sessionService.GetSessionAuto(uint32(s))
}

func (s UREDIS) GetNickName() string {
	return ""
}