package InitVenueStateTable

import (
	"Gym_booking_WeChat_mini_program/Model/VenueModel"
	"gorm.io/gorm"
)

func SetAllVenueStateIdle(db *gorm.DB, venueList ...string) {
	for _, venue := range venueList {
		var vs []VenueModel.VenueState
		db.Where("venue_type = ?", venue).Find(&vs)
		venueNum := len(vs) / 2
		for i := 1; i <= venueNum; i++ {
			todayVS := VenueModel.VenueState{VenueType: venue, ID: i, Date: "today"}
			tomorrowVS := VenueModel.VenueState{VenueType: venue, ID: i, Date: "tomorrow"}
			db.Save(&todayVS)
			db.Save(&tomorrowVS)
		}
	}
}
