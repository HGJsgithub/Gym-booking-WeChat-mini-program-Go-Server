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
			})
		} else {
			fmt.Printf("今天的%s %d号场地已经存在！\n", venueType, i+1)
		}
		if db.First(&tomorrowVenueState).RecordNotFound() == true {
			db.Create(&venueModel.VenueState{
				VenueType: venueType,
				ID:        i + 1,
				Date:      "tomorrow",
			})
		} else {
			fmt.Printf("明天的%s %d号场地已经存在！\n", venueType, i+1)
		}
	}
}
