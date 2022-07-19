/*
 * @Author: ww
 * @Date: 2022-07-19 02:38:00
 * @Description:
 * @FilePath: \danmu-play\game\luckdraw_test.go
 */
package game

import (
	"reflect"
	"testing"
)

func TestNewLuckDraw(t *testing.T) {
	tests := []struct {
		name string
		want *LuckDraw
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLuckDraw(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLuckDraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLuckDraw_ForceNewTask(t *testing.T) {
	type args struct {
		expire int
	}
	tests := []struct {
		name string
		tr   *LuckDraw
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.ForceNewTask(tt.args.expire)
		})
	}
}

func TestLuckDraw_NewTask(t *testing.T) {
	type args struct {
		expire int
	}
	tests := []struct {
		name    string
		tr      *LuckDraw
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case11", NewLuckDraw(), args{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.NewTask(tt.args.expire); (err != nil) != tt.wantErr {
				t.Errorf("LuckDraw.NewTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLuckDraw_Stop(t *testing.T) {
	tests := []struct {
		name string
		tr   *LuckDraw
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Stop()
		})
	}
}

func TestLuckDraw_isValid(t *testing.T) {
	tests := []struct {
		name string
		tr   *LuckDraw
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.isValid(); got != tt.want {
				t.Errorf("LuckDraw.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLuckDraw_Open(t *testing.T) {
	tests := []struct {
		name    string
		tr      *LuckDraw
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.Open()
			if (err != nil) != tt.wantErr {
				t.Errorf("LuckDraw.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LuckDraw.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLuckDraw_Join(t *testing.T) {
	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		tr      *LuckDraw
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Join(tt.args.uid); (err != nil) != tt.wantErr {
				t.Errorf("LuckDraw.Join() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
