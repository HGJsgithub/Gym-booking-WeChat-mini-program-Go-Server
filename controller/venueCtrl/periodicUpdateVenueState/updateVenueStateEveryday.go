package periodicUpdateVenueState

import (
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"fmt"
	"github.com/robfig/cron/v3"
)

func UpdateVenueStateEveryday(c *cron.Cron, venueList ...string) {
	_, err := c.AddFunc("@daily", func() {
		var vs venueModel.VenueState
		for _, venue := range venueList {
			vs.UpdateVenueStateEveryday(venue)
		}
	})
	if err != nil {
		fmt.Println("每天定期修改场地状态出错：", err)
	}

}
