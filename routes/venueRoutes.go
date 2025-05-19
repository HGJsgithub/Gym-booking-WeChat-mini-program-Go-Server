package routes

import (
	"Gym_booking_WeChat_mini_program/config/database"
	"Gym_booking_WeChat_mini_program/controller/venueCtrl/changeVenueState"
	"Gym_booking_WeChat_mini_program/controller/venueCtrl/getVenueState"
	"Gym_booking_WeChat_mini_program/middleware"
	"github.com/gin-gonic/gin"
)

func InitVenueRoutes(r *gin.Engine) {
	readMySQL := database.ConnectToMySQL()
	writeMySQL := database.ConnectToMySQL()
	redis := database.ConnectToRedis()
	venueRoutes := r.Group("/venue")
	{
		state := venueRoutes.Group("/state")
		{
			state.GET("/table", middleware.SetMySQL(readMySQL, "readMySQL"), getVenueState.GetStateTable)
			state.POST("/change", middleware.SetMySQL(writeMySQL, "writeMySQL"), middleware.SetRedis(redis), changeVenueState.ChangeVenueState)
		}
	}
}
