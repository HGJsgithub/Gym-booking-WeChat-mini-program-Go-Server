package PeriodicUpdateVenueState

import (
	"Gym_booking_WeChat_mini_program/Model/VenueModel"
	"fmt"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func UpdateVenueStateEveryHour(c *cron.Cron, db *gorm.DB) {
	_, err := c.AddFunc("@hourly", func() {
		nowHour := time.Now().Hour()
		nowHour--
		if nowHour >= 9 && nowHour <= 21 {
			timeField := "t" + strconv.Itoa(nowHour)
			var vs VenueModel.VenueState
			vs.UpdateVenueStateEveryHour(timeField, db)
		}
	})
	if err != nil {
		fmt.Println("每个小时修改场地状态出错：", err)
	}
}
