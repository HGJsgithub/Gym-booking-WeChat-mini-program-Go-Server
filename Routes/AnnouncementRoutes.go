package Routes

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/AnnouncementCtrl"
	"Gym_booking_WeChat_mini_program/Middleware"
	"github.com/gin-gonic/gin"
)

func InitAnnRoute(r *gin.Engine) {
	redis := Database.ConnectToRedis()
	annRoute := r.Group("/announcement")
	{
		//获取公告
		annRoute.Use(Middleware.SetRedis(redis)).GET("/get", AnnouncementCtrl.GetAnnouncement)
	}
}
