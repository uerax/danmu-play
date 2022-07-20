/*
 * @Author: ww
 * @Date: 2022-07-19 02:38:00
 * @Description:
 * @FilePath: /danmuplay/game/luckdraw.go
 */
package game

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/uerax/danmuplay/redis"
)

var Luckdraw = &LuckDrawTask{}

var redisPre = "luckdraw:"
var redisBase = ":base"
var redisList = ":list"

type LuckDrawTask struct {
	id string // id
	luckdraw *LuckDraw
}
type LuckDraw struct {
	Start      int64  `json:"start,omitempty"`        // 开始时间
	End        int64  `json:"end,omitempty"`          // 结束时间
	InProgress bool   `json:"inprogress,omitempty"` // 抽奖中
	Id         string `json:"id,omitempty"`         // 抽奖活动id
	Name       string `json:"name,omitempty"`       // 抽奖描述
	Cnt        int    `json:"cnt,omitempty"`        // 参与人数
	// luckdog 抽奖结束后生成
}

func NewLuckDraw(expire int) {
	ld := newLuckDraw(expire)
	Luckdraw.id = ld.Id
	Luckdraw.luckdraw = ld
	redis.HSetStruct(redisPre+ld.Id+redisBase,ld)
	go func(key string, end int64, expire int) {
		t := time.NewTicker(time.Second * time.Duration(expire))
		defer t.Stop()
		for range t.C {
			redis.HSet(redisPre+ld.Id+redisBase, "inprogress", false)
		}
	}(ld.Id, ld.End, expire)
}

func newLuckDraw(expire int) *LuckDraw {
	new := new(LuckDraw)
	new.NewTask(expire)
	return new
}

func (t *LuckDraw) NewTask(expire int) {
	now := time.Now()
	t.Start = now.Unix()
	t.End = now.Add(time.Minute * time.Duration(expire)).Unix()
	t.InProgress = true
	t.Id = now.Format("20060102150405")
	t.Cnt = 0	
}

func (t *LuckDraw) Stop() {
	redis.HSet(redisPre+t.Id+redisBase, "inprogress", false)
}

func (t *LuckDraw) Open() (string, error) {

	rand.Seed(t.Start)
	if !redis.HExists(redisPre+t.Id+redisBase, "luckydog") {
		count, err := redis.HGet(redisPre+t.Id+redisBase, "cnt")
		if err != nil {
			return "", err
		}
		cnt, err := strconv.Atoi(count)
		if err != nil {
			return "", err
		}
		luckdog := rand.Intn(cnt)
		uid, err := redis.LIndex(redisPre+t.Id+redisList, int64(luckdog))
		if err != nil {
			return "", err
		}
		redis.HSet(redisPre+t.Id+redisBase, "luckydog", uid)
		return uid, nil
	}
	return redis.HGet(redisPre+t.Id+redisBase, "luckydog")
}

func (t *LuckDraw) Join(uid string) {

	//todo
	redis.RPush(redisPre+t.Id+redisList, uid)
	redis.Hincrby(redisPre + t.Id + redisBase, "cnt", 1)
}

func JoinLuckDraw(uid string) {
	Luckdraw.luckdraw.Join(uid)
}

func EndLuckDraw() (string, error) {
	Luckdraw.luckdraw.Stop()
	return Luckdraw.luckdraw.Open()
}