package client

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func Set(key string, value interface{}) (interface{}, error) {
	return g.Redis().Do("set", key, value)
}

func Get(key string) (interface{}, error) {
	return toString(g.Redis().Do("get", key))
}

func Del(key string) (interface{}, error) {
	return g.Redis().Do("del", key)
}

func HSet(key string, field interface{}, value interface{}) (interface{}, error) {
	return g.Redis().Do("hSet", key, field, value)
}

func HGet(key string, field interface{}) (interface{}, error) {
	return toString(g.Redis().Do("hGet", key, field))
}

func HGetAll(key string) (interface{}, error) {
	return toStrings(g.Redis().Do("hGetAll", key))
}

func HDel(key string, field interface{}) (interface{}, error) {
	return g.Redis().Do("hDel", key, field)
}

func LPush(key string, value interface{}) (interface{}, error) {
	return g.Redis().Do("lPush", key, value)
}

func RPush(key string, value interface{}) (interface{}, error) {
	return g.Redis().Do("rPush", key, value)
}

func LPop(key string) (string, error) {
	return toString(g.Redis().Do("lPop", key))
}

func RPop(key string) (string, error) {
	return toString(g.Redis().Do("rPop", key))
}

func toString(data interface{}, err error) (string, error) {
	return gconv.String(data), err
}

func toStrings(data interface{}, err error) (interface{}, error) {
	return gconv.Strings(data), err
}
