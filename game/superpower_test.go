/*
 * @Author: ww
 * @Date: 2022-07-10 03:54:32
 * @Description:
 * @FilePath: /danmuplay/game/superpower_test.go
 */
package game

import "testing"

func TestGetSuperpower(t *testing.T) {
	s,y := GetSuperpower("1111", "test"), GetSuperpower("111111","etsss")
	t.Error(s, y)
}
