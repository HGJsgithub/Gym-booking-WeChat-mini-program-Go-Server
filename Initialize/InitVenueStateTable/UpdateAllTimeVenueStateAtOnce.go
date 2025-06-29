package InitVenueStateTable

import (
	"Gym_booking_WeChat_mini_program/Model/VenueModel"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func UpdateAllTimeVenueStateAtOnce(db *gorm.DB) {
	nowHour := time.Now().Hour()
	timeField := "t" + strconv.Itoa(nowHour)
	var vs VenueModel.VenueState
	if nowHour > 9 {
		for i := 9; i < nowHour; i++ {
			timeField = "t" + strconv.Itoa(i)
			vs.UpdateVenueStateEveryHour(timeField, db)
		}
	}
}
