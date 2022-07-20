/*
 * @Author: UerAx
 * @Date: 2022-07-20 00:47:13
 * @FilePath: /danmuplay/danmu/luckdraw.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package danmu

import (
	"strconv"

	"github.com/uerax/danmuplay/game"
)

func startLuckDraw(uid, name, msg string) error {
	i, err := strconv.Atoi(msg)
	if err != nil {
		return err
	}
	game.NewLuckDraw(i)
	
	return nil
}

func luckDraw(uid, name string) {
	game.JoinLuckDraw(uid)
}

func endLuckDraw() {
	// todo 后期可加上对uid的私信通知
	_, err := game.EndLuckDraw()
	if err != nil {
		ulog.Error(err)
	}
	
}
