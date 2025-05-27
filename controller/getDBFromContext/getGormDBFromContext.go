package getDBFromContext

import (
	"Gym_booking_WeChat_mini_program/config/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGormDBFromContext(c *gin.Context, key string) (*gorm.DB, error) {
	db, exist := c.Get(key)
	if !exist {
		return nil, fmt.Errorf("mysql连接不存在")
	}
	mysql := db.(*gorm.DB)
	alive := database.CheckMySQLAlive(mysql)
	if !alive {
		return nil, fmt.Errorf("mysql连接已经关闭")
	}
	return mysql, nil
}
