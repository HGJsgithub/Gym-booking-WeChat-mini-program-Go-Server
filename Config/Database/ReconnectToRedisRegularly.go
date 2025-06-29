package Database

import (
	"Gym_booking_WeChat_mini_program/ConstDefine"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func ReconnectToRedisRegularly() *redis.Client {
	var rdb *redis.Client
	for i := 0; i < ConstDefine.Repetitions; i++ {
		rdb = ConnectToRedis()
		if rdb == nil {
			log.Println("重连Redis失败")
			time.Sleep(ConstDefine.IntervalTime)
		} else {
			log.Println("重连Redis成功")
			break
		}
	}
	return rdb
}
