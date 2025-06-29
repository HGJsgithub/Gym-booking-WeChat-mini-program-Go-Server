package Initialize

import (
	"Gym_booking_WeChat_mini_program/Model"
	"Gym_booking_WeChat_mini_program/Model/VenueModel"
	"gorm.io/gorm"
)

func AutoMigrate(mysql *gorm.DB) error {
	err := mysql.AutoMigrate(&Model.Announcement{}, &VenueModel.VenueState{}, &Model.User{}, &Model.Order{})
	if err != nil {
		return err
	}
	return nil
}
