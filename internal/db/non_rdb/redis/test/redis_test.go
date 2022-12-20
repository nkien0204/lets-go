package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/nkien0204/lets-go/internal/db/non_rdb/redis"
)

func TestRedis(t *testing.T) {
	conn := redis.InitRedisConnection("localhost:55000", "default", "redispw", 0)
	ctx := context.Background()
	err := conn.Conn.Set(ctx, "key123", "value1231234", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := conn.Conn.Get(ctx, "key123").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
