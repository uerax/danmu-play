/*
 * @Author: ww
 * @Date: 2022-08-22 16:13:45
 * @Description:
 * @FilePath: /danmuplay/game/luckywheel_test.go
 */
package game

import (
	"fmt"
	"testing"
)

func TestNewLuckyWheel(t *testing.T) {
	t1 := &Gift{"", 100, 100, 30.0}
	t2 := &Gift{"", 100, 100, 30.0}
	t3 := &Gift{"", 100, 100, 3.0}
	t4 := &Gift{"", 3, 100, 0.001}
	ts := make([]*Gift, 0)
	ts = append(ts, t1)
	ts = append(ts, t2)
	ts = append(ts, t3)
	ts = append(ts, t4)
	lw := NewLuckyWheel(ts)
	fmt.Println(lw.binarySearch(0, 0, len(ts)))

}
