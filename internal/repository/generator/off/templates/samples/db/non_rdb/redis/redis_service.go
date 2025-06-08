package redis

import (
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Addr string
	Conn *redis.Client
}

func InitRedisConnection(addr, username, password string, db int) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       db,
	})

	return &RedisService{
		Addr: addr,
		Conn: rdb,
	}
}
