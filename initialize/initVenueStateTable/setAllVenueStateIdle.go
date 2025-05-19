package initVenueStateTable

import (
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"github.com/jinzhu/gorm"
)

func SetAllVenueStateIdle(db *gorm.DB, venueList ...string) {
	for _, venue := range venueList {
		var vs []venueModel.VenueState
		db.Where("venue_type = ?", venue).Find(&vs)
		venueNum := len(vs) / 2
		for i := 1; i <= venueNum; i++ {
			todayVS := venueModel.VenueState{VenueType: venue, ID: i, Date: "today"}
			tomorrowVS := venueModel.VenueState{VenueType: venue, ID: i, Date: "tomorrow"}
			db.Save(&todayVS)
			db.Save(&tomorrowVS)
		}
	}
}
