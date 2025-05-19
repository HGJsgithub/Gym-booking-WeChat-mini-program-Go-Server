package periodicUpdateVenueState

import (
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron/v3"
	"strconv"
	"time"
)

func UpdateVenueStateEveryHour(c *cron.Cron, mysql *gorm.DB) {
	_, err := c.AddFunc("@hourly", func() {
		nowHour := time.Now().Hour()
		nowHour--
		if nowHour >= 9 && nowHour <= 21 {
			timeField := "t" + strconv.Itoa(nowHour)
			var vs venueModel.VenueState
			vs.UpdateVenueStateEveryHour(timeField, mysql)
		}
	})
	if err != nil {
		fmt.Println("每个小时修改场地状态出错：", err)
	}
}
