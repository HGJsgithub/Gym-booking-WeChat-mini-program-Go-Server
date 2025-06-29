package Database

import (
	"Gym_booking_WeChat_mini_program/Config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMySQL() *gorm.DB {
	var conf Config.Config
	conf.LoadConfig()
	mysqlConf := conf.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DBName, mysqlConf.Charset, mysqlConf.ParseTime, mysqlConf.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm连接mysql失败！", err)
		return nil
	}
	sqlDB, _ := db.DB()
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(1000) // 空闲连接数
	return db
}
