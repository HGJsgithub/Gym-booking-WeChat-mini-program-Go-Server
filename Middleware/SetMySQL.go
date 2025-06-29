package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetMySQL(mysql *gorm.DB, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if mysql == nil {
			fmt.Println("MySQLMid未获取MySQL连接！")
			return
		}
		c.Set(key, mysql)
	}
}
