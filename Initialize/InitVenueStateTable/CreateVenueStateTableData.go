package InitVenueStateTable

import (
	"Gym_booking_WeChat_mini_program/Model/VenueModel"
	"fmt"
	"gorm.io/gorm"
)

func CreateVenueStateTableData(db *gorm.DB, venueType string, venueNum int) {
	for i := 0; i < venueNum; i++ {
		todayVenueState := VenueModel.VenueState{VenueType: venueType, ID: i + 1, Date: "today"}
		tomorrowVenueState := VenueModel.VenueState{VenueType: venueType, ID: i + 1, Date: "tomorrow"}
		var count int64
		db.First(&todayVenueState).Count(&count)
		if count == 0 {
			db.Create(&VenueModel.VenueState{
				VenueType: venueType,
				ID:        i + 1,
				Date:      "today",
			})
		} else {
			fmt.Printf("今天的%s %d号场地已经存在！\n", venueType, i+1)
		}
		count = 0
		db.First(&tomorrowVenueState).Count(&count)
		if count == 0 {
			db.Create(&VenueModel.VenueState{
				VenueType: venueType,
				ID:        i + 1,
				Date:      "tomorrow",
			})
		} else {
			fmt.Printf("明天的%s %d号场地已经存在！\n", venueType, i+1)
		}
	}
}
