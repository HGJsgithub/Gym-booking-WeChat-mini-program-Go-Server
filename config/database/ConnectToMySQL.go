package database

import (
	"Gym_booking_WeChat_mini_program/config"
	"fmt"
	"github.com/jinzhu/gorm"
)

func ConnectToMySQL() *gorm.DB {
	var conf config.Config
	conf.LoadConfig()
	mysql := conf.MySQL
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		mysql.User, mysql.Password, mysql.Host, mysql.Port, mysql.DBName, mysql.Charset, mysql.ParseTime, mysql.Loc)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		fmt.Println("gorm连接mysql出错！", err)
		return nil
	}
	sqlDB := db.DB()
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(1000) // 空闲连接数
	return db
}
