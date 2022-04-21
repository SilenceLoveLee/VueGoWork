package gredis

import "gin/pkg/setting"

var RedisConn *Cacher

func Setup() {
	options := Options{
		Addr:        setting.RedisSetting.Host,
		Password:    setting.RedisSetting.Password,
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Expire:      setting.RedisSetting.Expire,
	}
	RedisConn, _ = New(options)
}
