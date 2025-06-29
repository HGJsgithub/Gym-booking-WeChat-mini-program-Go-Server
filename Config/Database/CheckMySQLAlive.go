package Database

import (
	"gorm.io/gorm"
)

func CheckMySQLAlive(mysql *gorm.DB) bool {
	db, err := mysql.DB()
	if err != nil {
		return false
	}
	err = db.Ping()
	if err != nil {
		return false
	}
	return true
}
