/*
 * @Author: UerAx
 * @Date: 2022-07-07 23:54:13
 * @FilePath: \danmu-play\utils\util.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package utils

import "strings"

func CookieFetchCsrf(cookie string) string {
	slice := strings.Split(cookie, ";")
	for _, s := range slice {
		s = strings.TrimSpace(s)
		tit := strings.Index(s, "=")
		if tit != -1 && "bili_jct" == s[:tit] {
			return s[tit+1:]
		}
	}

	return ""
}
