package userLogin

import (
	"Gym_booking_WeChat_mini_program/model"
	"gorm.io/gorm"
)

func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}
