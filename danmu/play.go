/*
 * @Author: UerAx
 * @Date: 2022-07-08 14:35:55
 * @FilePath: /danmuplay/danmu/play.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package danmu

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/uerax/danmuplay/cfg"
	"github.com/uerax/danmuplay/game"
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

	ulog.Infof("[%s]: %s", msg.Info.([]interface{})[2].([]interface{})[1], msg.Info.([]interface{})[1])

	m := msg.Info.([]interface{})[1].(string)
	m = strings.TrimSpace(m)
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
	case "运势":
		getFate(uid, name)
	case "能力":
		getSuperpower(uid, name)
	}

	// redis.HGetAll(([]interface{})[2].([]interface{})[0].(string))
}

func SCMsgHandler(sc *model.SuperChatInfo) {
	ulog.Infof("[%s] %d元: %s %d", sc.Data.UserInfo.Uname, sc.Data.Price, sc.Data.Message, sc.Data.ID)
	uid := strconv.Itoa(sc.Data.UID)
	incr := sc.Data.Price * 10
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
		_, err = redis.HSet(uid, "name", sc.Data.UserInfo.Uname)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "point", incr)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "checkin", "")
		if err != nil {
			ulog.Error(err)
		}
		return
	}
	_, err = redis.Hincrby(uid, "point", int64(incr))
	if err != nil {
		ulog.Error(err)
	}
}

func GuardHandler(ci *model.CrewInfo) {
	ulog.Infof("[%s] 开通 %s * %d%s %d元", ci.Data.Username, ci.Data.RoleName, ci.Data.Num, ci.Data.Unit, ci.Data.Price/1000)
	uid := strconv.Itoa(ci.Data.UID)
	incr := ci.Data.Price/100
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
		_, err = redis.HSet(uid, "name", ci.Data.Username)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "point", incr)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "checkin", "")
		if err != nil {
			ulog.Error(err)
		}
		return
	}
	_, err = redis.Hincrby(uid, "point", int64(incr))
	if err != nil {
		ulog.Error(err)
	}
}

func GiftHandler(gf *model.GiftInfo) {
	ulog.Infof("[%s] 赠送 %d个 %s %.1f元", gf.Data.Uname, gf.Data.Num, gf.Data.GiftName, float64(gf.Data.Price)/1000*float64(gf.Data.Num))
	uid := strconv.Itoa(gf.Data.UID)
	incr := gf.Data.Price/100
	if incr == 0 {
		incr = gf.Data.Num
	}
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
		_, err = redis.HSet(uid, "name", gf.Data.Uname)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "point", incr)
		if err != nil {
			ulog.Error(err)
			return
		}
		_, err = redis.HSet(uid, "checkin", "")
		if err != nil {
			ulog.Error(err)
		}
		return
	}
	_, err = redis.Hincrby(uid, "point", int64(incr))
	if err != nil {
		ulog.Error(err)
	}
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
		Send(fmt.Sprintf("[%s] 签到成功", name))
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
	Send(fmt.Sprintf("[%s] 签到成功", name))
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

func getFate(uid, name string) {

	msg := game.GetFate(uid, name)
	ulog.Info(msg)
	Send(msg)
}

func getSuperpower(uid, name string) {
	msg := game.GetSuperpower(uid, name)
	ulog.Info(msg)
	Send(msg)
}
