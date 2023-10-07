package common

import (
	"sample-tabungan2/config"

	"github.com/gomodule/redigo/redis"
)

func NewRedisPool(cfg config.Config) *redis.Pool {
	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.RedisPort)
		},
	}

	return redisPool
}
