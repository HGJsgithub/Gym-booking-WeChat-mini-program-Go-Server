package initialize

import (
	"Gym_booking_WeChat_mini_program/model"
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(mysql *gorm.DB) {
	mysql.AutoMigrate(&model.Announcement{}, &venueModel.VenueState{}, &model.User{}, &model.Order{})
}
