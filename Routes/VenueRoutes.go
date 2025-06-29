package Routes

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/VenueCtrl/ChangeVenueState"
	"Gym_booking_WeChat_mini_program/Controller/VenueCtrl/GetVenueState"
	"Gym_booking_WeChat_mini_program/Middleware"
	"github.com/gin-gonic/gin"
)

func InitVenueRoutes(r *gin.Engine) {
	readMySQL := Database.ConnectToMySQL()
	writeMySQL := Database.ConnectToMySQL()
	redis := Database.ConnectToRedis()
	venueRoutes := r.Group("/venue")
	{
		state := venueRoutes.Group("/state")
		{
			state.GET("/table", Middleware.SetMySQL(readMySQL, "readMySQL"), GetVenueState.GetStateTable)
			state.POST("/change", Middleware.SetMySQL(writeMySQL, "writeMySQL"), Middleware.SetRedis(redis), ChangeVenueState.ChangeVenueState)
		}
	}
}
