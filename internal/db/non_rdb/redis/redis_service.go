package redis

import (
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Addr string
	Conn *redis.Client
}

func InitRedisConnection(addr, password string) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &RedisService{
		Addr: addr,
		Conn: rdb,
	}
}
