/*
 * @Author: UerAx
 * @Date: 2022-07-20 00:47:13
 * @FilePath: \danmu-play\danmu\luckdraw.go
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
	return game.Luckdraw.NewTask(i)

}

func luckDraw(uid, name string) {
	if err := game.Luckdraw.Join(uid); err != nil {
		ulog.Error(err)
	}
}
