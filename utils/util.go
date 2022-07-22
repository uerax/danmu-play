/*
 * @Author: UerAx
 * @Date: 2022-07-07 23:54:13
 * @FilePath: /danmuplay/utils/util.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package utils

import (
	"math/rand"
	"strings"
	"time"
)

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

func GenDevId() string {
	b := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
	s := []byte("xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx")
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(s); i++ {
		if '-' == s[i] || '4' == s[i] {
			continue
		}
		randomInt := rand.Intn(16);
		if 'x' == s[i] {
			s[i] = b[randomInt];
		} else {
			s[i] = b[3 & randomInt | 8];
		}
	}
	return string(s);
}