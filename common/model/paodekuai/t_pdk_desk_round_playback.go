package paodekuai

import (
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
)

//回放可以通过这个表来查询
type T_pdk_desk_round_playback struct {
	DeskId     int32
	GameNumber int32
	PlayBackData []*ddproto.PdkPlaybackSnapshot
}

func GetPdkPlayBack(gamenumber int32) []*ddproto.PdkPlaybackSnapshot {
	ret := &T_pdk_desk_round_playback{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_PDK_DESK_ROUND_PLAYBACK).Find(bson.M{"gamenumber": gamenumber}).One(ret)
	})
	if ret.DeskId > 0 {
		return ret.PlayBackData
	} else {
		return nil
	}
}
