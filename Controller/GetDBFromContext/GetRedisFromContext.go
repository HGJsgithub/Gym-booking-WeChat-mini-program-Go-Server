package GetDBFromContext

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetRedisFromContext(c *gin.Context) (*redis.Client, error) {
	Redis, exist := c.Get("redis")
	if !exist {
		return nil, fmt.Errorf("未获取Redis连接")
	}
	rdb := Redis.(*redis.Client)
	if !Database.CheckRedisAlive(rdb) {
		return nil, fmt.Errorf("redis连接已经断开")
	}
	return rdb, nil
}
