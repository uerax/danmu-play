/*
 * @Author: UerAx
 * @Date: 2022-07-08 15:18:34
 * @FilePath: \danmu-play\danmuplay.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package main

import (
	"flag"

	"github.com/uerax/danmuplay/cfg"
	"github.com/uerax/danmuplay/danmu"
	"github.com/uerax/danmuplay/global"
	"github.com/uerax/danmuplay/redis"
)

var path string
var out string

func init() {
	flag.StringVar(&path, "c", "./etc", "请输入配置文件路径或者文件夹 usage: -c /etc")
	flag.StringVar(&out, "o", "", "请输入日志输出位置 usage: -o /logs")
	flag.Parse()
}

func main() {

	if out != "" {
		global.Out(out)
	}

	err := cfg.Init(path)

	if err != nil {
		global.Log.Panic(err)
	}

	redis.Init()

	danmu.Init()

}
