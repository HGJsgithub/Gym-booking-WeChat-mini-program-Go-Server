package initVenueStateTable

import (
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func UpdateAllTimeVenueStateAtOnce(db *gorm.DB) {
	nowHour := time.Now().Hour()
	timeField := "t" + strconv.Itoa(nowHour)
	var vs venueModel.VenueState
	if nowHour > 9 {
		for i := 9; i < nowHour; i++ {
			timeField = "t" + strconv.Itoa(i)
			vs.UpdateVenueStateEveryHour(timeField, db)
		}
	}
}
