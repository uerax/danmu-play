/*
 * @Author: ww
 * @Date: 2022-07-07 06:53:07
 * @Description:
 * @FilePath: \danmu-play\danmu\send.go
 */
package danmu

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/uerax/danmuplay/cfg"
	"github.com/uerax/danmuplay/utils"
)

func Send(msg string) {
	// 直接curl转换懒得改了
	roomid, err := cfg.Config.GetValue("roomid")
	if err != nil {
		ulog.Error(err)
		return
	}
	cookie, err := cfg.Config.GetValue("cookie")
	if err != nil {
		ulog.Error(err)
		return
	}
	csrf := utils.CookieFetchCsrf(cookie.(string))

	data := `color=16777215&fontsize=25&mode=1&msg=` + msg + `&rnd=` + fmt.Sprint(time.Now().Unix()) + `&roomid=` + fmt.Sprint(roomid) + `&bubble=0&csrf_token=` + csrf + `&csrf=` + csrf

	req, err := http.NewRequest("POST", "https://api.live.bilibili.com/msg/send", strings.NewReader(data))
	if err != nil {
		ulog.Error(err)
		return
	}
	req.Header.Set("Host", "api.live.bilibili.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	req.Header.Set("Cookie", cookie.(string))
	req.Header.Set("Origin", "https://live.bilibili.com")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", fmt.Sprintf("https://live.bilibili.com/%d", roomid))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ulog.Error(err)
	}
	defer resp.Body.Close()
}
