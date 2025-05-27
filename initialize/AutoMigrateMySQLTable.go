package initialize

import (
	"Gym_booking_WeChat_mini_program/model"
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"gorm.io/gorm"
)

func AutoMigrate(mysql *gorm.DB) error {
	err := mysql.AutoMigrate(&model.Announcement{}, &venueModel.VenueState{}, &model.User{}, &model.Order{})
	if err != nil {
		return err
	}
	return nil
}
