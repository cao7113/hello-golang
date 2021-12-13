package db

import (
	"context"
	"github.com/cao7113/hellogolang/config"
	"github.com/go-redis/redis/v8"
)

// https://github.com/go-redis/redis
var rCtx = context.Background()
var RedisConn *redis.Client

func init() {
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisAddr,
		Password: config.Config.RedisPasswd,
		DB:       0, // use default DB
	})
}
