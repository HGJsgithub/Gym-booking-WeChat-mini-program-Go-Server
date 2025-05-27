package main

import (
	"Gym_booking_WeChat_mini_program/config/database"
	"Gym_booking_WeChat_mini_program/controller/venueCtrl/periodicUpdateVenueState"
	"Gym_booking_WeChat_mini_program/initialize"
	"Gym_booking_WeChat_mini_program/initialize/initVenueStateTable"
	"Gym_booking_WeChat_mini_program/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	db := database.ConnectToMySQL()
	redis := database.ConnectToRedis()
	//创建数据库表并赋值
	err := initialize.AutoMigrate(db)
	if err != nil {
		fmt.Println("AutoMigrate err:", err)
		return
	}
	initialize.CreateAnnData(db)
	initVenueStateTable.CreateVenueStateTableData(db, "badminton", 4)
	initVenueStateTable.CreateVenueStateTableData(db, "tableTennis", 3)
	initVenueStateTable.CreateVenueStateTableData(db, "tennis", 2)

	//把所有场地状态变成空闲
	venueList := []string{"badminton", "tableTennis", "tennis"}
	initVenueStateTable.SetAllVenueStateIdle(db, venueList...)

	initVenueStateTable.UpdateAllTimeVenueStateAtOnce(db)

	database.SetMySQLDataToRedis(db, redis)

	r := gin.Default()
	//初始化所有路由
	routes.AllRouteInit(r)

	c := cron.New()

	db = database.ConnectToMySQL()
	//每小时更新场地状态
	periodicUpdateVenueState.UpdateVenueStateEveryHour(c, db)

	//每天更新场地状态
	periodicUpdateVenueState.UpdateVenueStateEveryday(c, venueList...)

	c.Start()

	err = r.Run(":8080")
	fmt.Println("体育馆预约微信小程序服务端开始运行~(～￣▽￣)～")
	if err != nil {
		fmt.Println("r.Run出错：", err)
		return
	}
}
