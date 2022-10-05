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
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	return &RedisService{
		Addr: addr,
		Conn: rdb,
	}
}
