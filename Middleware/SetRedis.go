package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetRedis(redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		if redis == nil {
			fmt.Println("RedisMid未连接到Redis！")
			return
		}
		c.Set("redis", redis)
	}
}
