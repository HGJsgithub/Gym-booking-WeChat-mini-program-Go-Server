package Database

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func CheckRedisAlive(redis *redis.Client) bool {
	pong, err := redis.Ping(context.Background()).Result()
	return err == nil && pong == "PONG"
}
