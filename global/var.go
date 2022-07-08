/*
 * @Author: ww
 * @Date: 2022-07-02 02:42:36
 * @Description:
 * @FilePath: \danmu-play\global\var.go
 */
package global

import "github.com/uerax/ulog"

var Log *ulog.Ulog

func init() {
	Log = ulog.New()
}

func Out(out ...string) {
	if len(out[0]) != 0 {
		Log.OutFile(out[0])
	}
}
