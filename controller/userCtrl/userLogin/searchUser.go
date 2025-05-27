package userLogin

import (
	"Gym_booking_WeChat_mini_program/model"
	"gorm.io/gorm"
)

func SearchUser(db *gorm.DB, phone string, password string, user *model.User) (exist, right bool) {
	if IsPhoneExist(db, phone) {
		db.Where("phone = ?", phone).First(user)
		if CheckPassword(password, user.Password) {
			return true, true
		} else {
			return true, false
		}
	} else {
		return false, false
	}
}
