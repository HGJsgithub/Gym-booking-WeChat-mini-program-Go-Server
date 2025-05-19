package main

import (
	"Gym_booking_WeChat_mini_program/config/database"
	"Gym_booking_WeChat_mini_program/controller/venueCtrl/periodicUpdateVenueState"
	"Gym_booking_WeChat_mini_program/initialize"
	"Gym_booking_WeChat_mini_program/initialize/initVenueStateTable"
	"Gym_booking_WeChat_mini_program/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robfig/cron/v3"
)

func main() {
	mysql := database.ConnectToMySQL()
	redis := database.ConnectToRedis()
	//创建数据库表并赋值
	initialize.AutoMigrate(mysql)
	initialize.CreateAnnData(mysql)
	initVenueStateTable.CreateVenueStateTableData(mysql, "badminton", 4)
	initVenueStateTable.CreateVenueStateTableData(mysql, "tableTennis", 3)
	initVenueStateTable.CreateVenueStateTableData(mysql, "tennis", 2)

	//把所有场地状态变成空闲
	venueList := []string{"badminton", "tableTennis", "tennis"}
	initVenueStateTable.SetAllVenueStateIdle(mysql, venueList...)

	initVenueStateTable.UpdateAllTimeVenueStateAtOnce(mysql)

	database.SetMySQLDataToRedis(mysql, redis)

	r := gin.Default()
	//初始化所有路由
	routes.AllRouteInit(r)

	c := cron.New()

	mysql = database.ConnectToMySQL()
	//每小时更新场地状态
	periodicUpdateVenueState.UpdateVenueStateEveryHour(c, mysql)

	//每天更新场地状态
	periodicUpdateVenueState.UpdateVenueStateEveryday(c, venueList...)

	c.Start()

	fmt.Println("体育馆预约微信小程序服务端开始运行~(～￣▽￣)～")

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("r.Run出错：", err)
		return
	}
}
