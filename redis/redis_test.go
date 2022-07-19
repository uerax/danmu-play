/*
 * @Author: UerAx
 * @Date: 2022-07-08 16:21:39
 * @FilePath: \danmu-play\redis\redis_test.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */

package redis

import (
	"testing"

	"github.com/uerax/danmuplay/cfg"
)

func TestClient(t *testing.T) {
	cfg.Init("../etc")
	Init()
	// HSet(uid, "uid", uid)
	// HSet(uid, "Name", name)
	// HSet(uid, "Point", 1)
	// HSet(uid, "Checkin", 1)
}
