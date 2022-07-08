/*
 * @Author: UerAx
 * @Date: 2022-07-08 14:35:55
 * @FilePath: \danmu-play\danmu\play.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package danmu

import (
	"fmt"
	"strconv"
	"time"

	"github.com/uerax/danmuplay/cfg"
	"github.com/uerax/danmuplay/model"
	"github.com/uerax/danmuplay/redis"
)

func Init() {
	dm := NewBiliRoom(cfg.GetStringWithDefault("746504", "roomid"))
	dm.SCMsgHandler = SCMsgHandler
	dm.MsgHandler = MsgHandler
	dm.GuardHandler = GuardHandler
	dm.GiftHandler = GiftHandler
	go dm.Start()
	dm.DanmuHandler()
}

func MsgHandler(msg *model.MessageInfo) {

	// ulog.Info(msg)

	m := msg.Info.([]interface{})[1].(string)

	if m[0] != '#' {
		return
	}

	m = m[1:]

	uid := strconv.Itoa(int(msg.Info.([]interface{})[2].([]interface{})[0].(float64)))

	name := fmt.Sprint(msg.Info.([]interface{})[2].([]interface{})[1])

	switch m {
	case "签到":
		checkIn(uid, name)
	case "积分":
		getPoint(uid, name)
	}

	// redis.HGetAll(([]interface{})[2].([]interface{})[0].(string))
}

func SCMsgHandler(sc *model.SuperChatInfo) {
}

func GuardHandler(gd *model.CrewInfo) {
}

func GiftHandler(gf *model.GiftInfo) {
}

func checkIn(uid, name string) {
	now := time.Now().Format("20060102")
	exists, err := redis.Exists(uid)
	if err != nil {
		ulog.Error(err)
		return
	}
	if !exists {
		_, err = redis.HSet(uid, "uid", uid)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "name", name)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "point", 1)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "checkin", now)
		if err != nil {
			ulog.Error(err)
		}
		return
	}

	lastCheckIn, err := redis.HGet(uid, "checkin")
	if err != nil {
		ulog.Error(err)
		return
	}
	if lastCheckIn != now {
		_, err = redis.Hincrby(uid, "point", 1)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "checkin", now)
		if err != nil {
			ulog.Error(err)
			return
		}
	}
}

func getPoint(uid, name string) {
	pit, err := redis.HGet(uid, "point")
	if err != nil {
		ulog.Error(err)
		return
	}

	msg := fmt.Sprintf("[%s] 当前积分：%s", name, pit)

	Send(msg)

}
