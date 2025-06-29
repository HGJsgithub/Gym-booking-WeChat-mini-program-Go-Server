package UserLogin

import (
	"Gym_booking_WeChat_mini_program/Model"
	"gorm.io/gorm"
)

func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user Model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}
