/*
 * @Author: ww
 * @Date: 2022-07-19 02:38:00
 * @Description:
 * @FilePath: /danmuplay/game/luckdraw.go
 */
package game

import (
	"errors"
	"math/rand"
	"time"

	"github.com/uerax/danmuplay/redis"
)

var Luckdraw = NewLuckDraw()

var errPreNotEnd = errors.New("有抽奖未结束")
var errNotTask = errors.New("没有抽奖活动")
var errNotParticipant = errors.New("没有参与者或不存在任务")

var redisPre = "luckdraw:"
var redisBase = ":base"
var redisList = ":list"

type LuckDraw struct {
	Start      int64  // 开始时间
	End        int64  // 结束时间
	InProgress bool   // 抽奖中
	Id         string // 抽奖活动id
	Name       string // 抽奖描述
	Cnt        int    // 参与人数
	// luckdog 抽奖结束后生成
}

func NewLuckDraw() *LuckDraw {
	return &LuckDraw{}
}

func (t *LuckDraw) ForceNewTask(expire int) {
	t.Stop()
	t.NewTask(expire)
}

func (t *LuckDraw) NewTask(expire int) error {

	if t.InProgress && t.isValid() {
		return errPreNotEnd
	}

	now := time.Now()
	t.Start = now.Unix()
	t.End = now.Add(time.Minute * time.Duration(expire)).Unix()
	t.InProgress = true
	t.Id = now.Format("20060102150405")
	t.Cnt = 0
	redis.HSetStruct(redisPre+t.Id+redisBase, t)
	return nil
}

func (t *LuckDraw) Stop() {
	if t.InProgress {
		t.InProgress = false
	}
	
}

func (t *LuckDraw) isValid() bool {
	return time.Now().Unix() < t.End
}

func (t *LuckDraw) Open() (string, error) {

	if t.Cnt == 0 {
		return "", errNotParticipant
	}

	rand.Seed(t.Start)
	if !redis.HExists(redisPre+t.Id+redisBase, "luckydog") {
		luckdog := rand.Intn(t.Cnt)
		redis.HSet(redisPre+t.Id+redisBase, "luckydog", luckdog)
		return redis.LIndex(redisPre+t.Id+redisList, int64(luckdog))
	}
	return redis.HGet(redisPre+t.Id+redisBase, "luckydog")
}

func (t *LuckDraw) Join(uid string) error {
	if !t.InProgress {
		return errNotTask
	}
	//todo
	redis.RPush(redisPre+t.Id+redisList, uid)
	t.Cnt++
	return nil
}
