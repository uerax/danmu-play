/*
 * @Author: ww
 * @Date: 2022-07-09 04:07:55
 * @Description:
 * @FilePath: /danmuplay/game/fate_test.go
 */
package game

import "testing"

func TestGetFate(t *testing.T) {
	type args struct {
		uid  string
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case", args{"111111","test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFate(tt.args.uid, tt.args.name)
			print(got)
		})
	}
}
