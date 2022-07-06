/*
 * @Author: ww
 * @Date: 2022-07-07 06:53:07
 * @Description:
 * @FilePath: /danmuplay/danmu/send.go
 */
package danmu

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/uerax/danmuplay/cfg"
)

func Send(msg string) {
	// 懒得写直接curl转换
	roomid, err := cfg.Config.GetValue("roomid")
	if err != nil {
		ulog.Error(err)
		return
	}
	data := fmt.Sprintf("------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"bubble\"\r\n\r\n0\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"msg\"\r\n\r\n%s\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"color\"\r\n\r\n65532\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"mode\"\r\n\r\n1\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"fontsize\"\r\n\r\n25\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"rnd\"\r\n\r\n1657146441\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"roomid\"\r\n\r\n%d\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"csrf\"\r\n\r\n066a925463717b2ab4e7344619e0a515\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU\r\nContent-Disposition: form-data; name=\"csrf_token\"\r\n\r\n066a925463717b2ab4e7344619e0a515\r\n------WebKitFormBoundaryLP53IeC9HDisGqhU--\r\n", msg, roomid)
	body := strings.NewReader(data)
	req, err := http.NewRequest("POST", "https://api.live.bilibili.com/msg/send", body)
	if err != nil {
		ulog.Error(err)
		return
	}
	req.Header.Set("Authority", "api.live.bilibili.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundaryLP53IeC9HDisGqhU")
	cookie, err := cfg.Config.GetValue("cookie")
	if err != nil {
		ulog.Error(err)
		return
	}
	req.Header.Set("Cookie", cookie.(string))
	req.Header.Set("Origin", "https://live.bilibili.com")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://live.bilibili.com/746504?broadcast_type=0&is_room_feed=1&spm_id_from=333.999.0.0")
	req.Header.Set("Sec-Ch-Ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ulog.Error(err)
	}
	defer resp.Body.Close()
}
