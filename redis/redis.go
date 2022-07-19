/*
 * @Author: UerAx
 * @Date: 2022-07-08 16:21:39
 * @FilePath: /danmuplay/redis/redis.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package redis

import (
	"context"
	"reflect"

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

func HExists(key, field string) bool {
	b, _ := rdb.HExists(ctx, key,field).Result()
	return b
}

func HSet(key string, val ...interface{}) (int64, error) {
	return rdb.HSet(ctx, key, val).Result()
}

func HSetStruct(key string, val interface{}) (int64, error) {
	return rdb.HSet(ctx, key, structToMap(val)).Result()
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

func LPush(key string, val ...interface{}) (int64, error) {
	return rdb.LPush(ctx, key, val...).Result()
}

func RPush(key string, val ...interface{}) (int64, error) {
	return rdb.RPush(ctx, key, val...).Result()
}

func LIndex(key string, index int64) (string, error) {
	return rdb.LIndex(ctx, key, index).Result()
}

func structToMap(val any) map[string]any {

	if val == nil {
		return nil
	}

	hash := make(map[string]any, 0)

	t := reflect.TypeOf(val) // 获取obj的反射类型

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	v := reflect.Indirect(reflect.ValueOf(val)) // 获取obj的反射值

	for i := 0; i < t.NumField(); i++ {
		// 存在tag，则使用tag作为key
		var key string
		tag := t.Field(i).Tag.Get("json")
		if tag != "" {
			key = tag
		} else {
			key = t.Field(i).Name
		}
		hash[key] = v.Field(i).Interface()
	}

	return hash
}
