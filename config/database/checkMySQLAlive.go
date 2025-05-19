package database

import (
	"github.com/jinzhu/gorm"
)

func CheckMySQLAlive(mysql *gorm.DB) bool {
	return mysql.DB().Ping() == nil
}
