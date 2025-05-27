package database

import (
	"Gym_booking_WeChat_mini_program/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client {
	var conf config.Config
	conf.LoadConfig()
	Redis := conf.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: Redis.Address,
		DB:   Redis.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("连接Redis出错！", err)
		return nil
	}
	return rdb
}
