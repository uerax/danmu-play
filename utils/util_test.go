/*
 * @Author: UerAx
 * @Date: 2022-07-07 23:54:13
 * @FilePath: \danmu-play\utils\util_test.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */

package utils

import "testing"

func TestCookieFetchCsrf(t *testing.T) {
	cookie := `LIVE_BUVID=AUTO2915927237938989; buvid3=212C9FB4-FAF1-46D4-ACFE-A71838139B8F143106infoc; rpdid=|(um|k~|RJYk0J'ulmm|kJuJ|; _ga=GA1.2.712988129.1614841082; video_page_version=v_old_home; CURRENT_BLACKGAP=0; blackside_state=0; fingerprint_s=bfa9684d727dca783aeb37c99dd223c1; DedeUserID=211336; DedeUserID__ckMd5=11a37516312489c6; b_ut=5; buvid4=318031D9-0F28-5520-3483-5E65B98E05A365328-022012119-+V62Zou/9yraW1CqBHK4WQ%3D%3D; i-wanna-go-back=-1; buvid_fp_plain=undefined; buvid_fp=cab4f8874a979dfd92cce715957d31c6; nostalgia_conf=-1; hit-dyn-v2=1; is-2022-channel=1; _uuid=16DEDCE6-4118-8E24-81E5-91FF2A67F9E535003infoc; fingerprint3=d1a134f25cc6479ad01da88ac5bd819f; fingerprint=4cbde8634d08d45ea02e4bb90a4eba23; SESSDATA=c0570311%2C1672579662%2C36388%2A71; bili_jct=867b55d3680d3c6bcf27aa99aa428cae; Hm_lvt_8a6e55dbd2870f0f5bc9194cddf32a02=1656603785,1656944819,1657012892,1657111906; CURRENT_QUALITY=80; sid=5fxb6ffr; CURRENT_FNVAL=4048; bsource=search_baidu; b_lsid=E682FC98_181D90D268E; bp_video_offset_211336=680179019989123200; _dfcaptcha=db8df8d0e8ce230e2596311c21474397; innersign=0; b_timer=%7B%22ffp%22%3A%7B%22333.1007.fp.risk_212C9FB4%22%3A%22181D9590BC9%22%2C%22333.788.fp.risk_212C9FB4%22%3A%22181D83E0F08%22%2C%22444.8.fp.risk_212C9FB4%22%3A%22181D94E5115%22%2C%22333.976.fp.risk_212C9FB4%22%3A%22181D94D6A7D%22%2C%22333.880.fp.risk_212C9FB4%22%3A%22181D6D8AE52%22%2C%22666.25.fp.risk_212C9FB4%22%3A%22181D9504AA5%22%2C%22333.6.fp.risk_212C9FB4%22%3A%22181D84DA810%22%2C%22333.999.fp.risk_212C9FB4%22%3A%22181D95912C2%22%7D%7D; PVID=12`
	type args struct {
		cookie string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{cookie}, "867b55d3680d3c6bcf27aa99aa428cae"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CookieFetchCsrf(tt.args.cookie); got != tt.want {
				t.Errorf("CookieFetchCsrf() = %v, want %v", got, tt.want)
			}
		})
	}

}
