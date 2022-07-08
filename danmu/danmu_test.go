/*
 * @Author: ww
 * @Date: 2022-07-02 02:25:38
 * @Description:
 * @FilePath: \danmu-play\danmu\danmu_test.go
 */
package danmu

import (
	"fmt"
	"os"
	"testing"

	"github.com/uerax/danmuplay/cfg"
	"github.com/uerax/danmuplay/model"
)

func TestLog(t *testing.T) {
	ulog.Info(os.Getwd())
	cfg.Init("../etc")
	Send("test")
}

func TestNewBiliRoom(t *testing.T) {
	cfg.Init("../etc")
	tmp := NewBiliRoom("746504")
	tmp.MsgHandler = func(mi *model.MessageInfo) {
		Send(fmt.Sprintf("[弹幕] %s: 签到成功", mi.Info.([]interface{})[2].([]interface{})[1]))
	}
	go tmp.Start()
	tmp.DanmuHandler()
}
