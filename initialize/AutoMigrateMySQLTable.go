package initialize

import (
	"Gym_booking_WeChat_mini_program/model"
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(mysql *gorm.DB) {
	mysql.AutoMigrate(&model.Announcement{})
	mysql.AutoMigrate(&venueModel.VenueState{})
	mysql.AutoMigrate(&model.User{})
	//mysql.AutoMigrate(&model.Admin{})
	mysql.AutoMigrate(&model.Order{})
}
