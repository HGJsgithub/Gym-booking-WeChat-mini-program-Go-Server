package Database

import (
	"Gym_booking_WeChat_mini_program/ConstDefine"
	"gorm.io/gorm"
	"log"
	"time"
)

func ReconnectToMySQLRegularly() *gorm.DB {
	var db *gorm.DB
	for i := 0; i < ConstDefine.Repetitions; i++ {
		db = ConnectToMySQL()
		if db == nil {
			log.Println("重连MySQL失败")
			time.Sleep(ConstDefine.IntervalTime)
		} else {
			log.Println("重连MySQL成功")
			break
		}
	}
	return db
}
