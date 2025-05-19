package routes

import (
	"Gym_booking_WeChat_mini_program/config/database"
	"Gym_booking_WeChat_mini_program/controller/annCtrl"
	"Gym_booking_WeChat_mini_program/middleware"
	"github.com/gin-gonic/gin"
)

func InitAnnRoute(r *gin.Engine) {
	redis := database.ConnectToRedis()
	annRoute := r.Group("/announcement")
	{
		//获取公告
		annRoute.Use(middleware.RedisMid(redis)).GET("/get", annCtrl.GetAnn)
		//删除公告
		annRoute.POST("/delete", annCtrl.DeleteAnn)
	}
}
