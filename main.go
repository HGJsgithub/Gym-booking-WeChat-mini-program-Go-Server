package main

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/VenueCtrl/PeriodicUpdateVenueState"
	"Gym_booking_WeChat_mini_program/Initialize"
	"Gym_booking_WeChat_mini_program/Initialize/InitVenueStateTable"
	"Gym_booking_WeChat_mini_program/Routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	db := Database.ReconnectToMySQLRegularly()
	redis := Database.ReconnectToRedisRegularly()
	//创建数据库表并赋值
	err := Initialize.AutoMigrate(db)
	if err != nil {
		fmt.Println("AutoMigrate err:", err)
		return
	}
	Initialize.CreateAnnData(db)
	InitVenueStateTable.CreateVenueStateTableData(db, "badminton", 4)
	InitVenueStateTable.CreateVenueStateTableData(db, "tableTennis", 3)
	InitVenueStateTable.CreateVenueStateTableData(db, "tennis", 2)

	//把所有场地状态变成空闲
	venueList := []string{"badminton", "tableTennis", "tennis"}
	InitVenueStateTable.SetAllVenueStateIdle(db, venueList...)
	InitVenueStateTable.UpdateAllTimeVenueStateAtOnce(db)

	Database.SetMySQLDataToRedis(db, redis)

	// 设置为 release 模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//初始化所有路由
	Routes.AllRouteInit(r)

	c := cron.New()

	db = Database.ConnectToMySQL()
	//每小时更新场地状态
	PeriodicUpdateVenueState.UpdateVenueStateEveryHour(c, db)

	//每天更新场地状态
	PeriodicUpdateVenueState.UpdateVenueStateEveryday(c, venueList...)

	c.Start()

	err = r.Run(":8080")
	fmt.Println("体育馆预约微信小程序服务端开始运行~(～￣▽￣)～")
	if err != nil {
		panic("r.Run err:" + err.Error())
	}
}
