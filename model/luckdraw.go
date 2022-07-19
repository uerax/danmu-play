/*
 * @Author: ww
 * @Date: 2022-07-19 02:39:02
 * @Description:
 * @FilePath: /danmuplay/model/luckdraw.go
 */
package model

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"sync"
	"time"
)

var errPreNotEnd = errors.New("有抽奖未结束")
var errNotTask = errors.New("没有抽奖活动")
var errNotParticipant = errors.New("没有参与者或不存在任务")

type LuckDraw struct {
	start int64	//开始时间
	end int64 // 结束时间
	inProgress bool // 抽奖中
	Id string // 抽奖活动id
	Participant []string	// 参与人
	Cnt int	// 参与人数
	rw	sync.RWMutex
}

func NewLuckDraw() *LuckDraw {
	return &LuckDraw{
		rw: sync.RWMutex{},
	}
}

func (t *LuckDraw) NewTask(expire int) error {
	if t.inProgress {
		return errPreNotEnd
	}
	t.rw.RLock()
	defer t.rw.Unlock()

	now := time.Now()
	t.start = now.Unix()
	t.end = now.Add(time.Minute * time.Duration(expire)).Unix()
	t.inProgress = true
	t.Id = base64.StdEncoding.EncodeToString([]byte(now.Format("2006-01-02 15:04:05")))
	t.Participant = make([]string, 0)
	t.Cnt = 0
	return nil
}

func (t *LuckDraw) Stop() {
	if t.inProgress {
		t.inProgress = false
	}
}

func (t *LuckDraw) Open() (string, error) {

	if t.Cnt == 0 || len(t.Participant) == 0 {
		return "", errNotParticipant
	}

	rand.Seed(t.start)
	return t.Participant[rand.Intn(t.Cnt)], nil
}

func (t *LuckDraw) Join(uid string) error {
	if !t.inProgress {
		return errNotTask
	}
	t.rw.RLock()
	defer t.rw.Unlock()
	if t.Participant == nil {
		t.Participant = make([]string, 0)
	}
	t.Participant = append(t.Participant, uid)
	t.Cnt++
	return nil
}

