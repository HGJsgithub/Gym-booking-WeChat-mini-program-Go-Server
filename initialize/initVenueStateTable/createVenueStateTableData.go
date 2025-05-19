package initVenueStateTable

import (
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"fmt"
	"github.com/jinzhu/gorm"
)

func CreateVenueStateTableData(db *gorm.DB, venueType string, venueNum int) {
	for i := 0; i < venueNum; i++ {
		todayVenueState := venueModel.VenueState{VenueType: venueType, ID: i + 1, Date: "today"}
		tomorrowVenueState := venueModel.VenueState{VenueType: venueType, ID: i + 1, Date: "tomorrow"}
		if db.First(&todayVenueState).RecordNotFound() == true {
			db.Create(&venueModel.VenueState{
				VenueType: venueType,
				ID:        i + 1,
				Date:      "today",
				T9:        false,
				T10:       false,
				T11:       false,
				T12:       false,
				T13:       false,
				T14:       false,
				T15:       false,
				T16:       false,
				T17:       false,
				T18:       false,
				T19:       false,
				T20:       false,
				T21:       false,
			})
		} else {
			fmt.Printf("今天的%s %d号场地已经存在！\n", venueType, i+1)
			fmt.Println()
		}
		if db.First(&tomorrowVenueState).RecordNotFound() == true {
			db.Create(&venueModel.VenueState{
				VenueType: venueType,
				ID:        i + 1,
				Date:      "tomorrow",
				T9:        false,
				T10:       false,
				T11:       false,
				T12:       false,
				T13:       false,
				T14:       false,
				T15:       false,
				T16:       false,
				T17:       false,
				T18:       false,
				T19:       false,
				T20:       false,
				T21:       false,
			})
		} else {
			fmt.Printf("明天的%s %d号场地已经存在！\n", venueType, i+1)
			fmt.Println()
		}
	}
}
