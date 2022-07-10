/*
 * @Author: UerAx
 * @Date: 2022-07-08 16:21:39
 * @FilePath: /danmuplay/redis/redis.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/uerax/danmuplay/cfg"
)

var rdb *redis.Client
var ctx = context.Background()

func Init() {
	client()
}

func client() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.GetStringWithDefault("localhost:6379", "redis", "url"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Exists(key string) (bool, error) {
	i, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return i == 1, nil
}

func HSet(key string, val ...interface{}) (int64, error) {
	return rdb.HSet(ctx, key, val).Result()
}

func HGetAll(key string) (map[string]string, error) {
	return rdb.HGetAll(ctx, key).Result()
}

func HGet(key string, field string) (string, error) {
	return rdb.HGet(ctx, key, field).Result()
}

func Hincrby(key string, field string, incr int64) (int64, error) {
	return rdb.HIncrBy(ctx, key, field, incr).Result()
}

func Check() (string, error) {
	return rdb.Ping(ctx).Result()
}
