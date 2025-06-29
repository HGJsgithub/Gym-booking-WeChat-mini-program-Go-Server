package AnnouncementCtrl

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"Gym_booking_WeChat_mini_program/Model"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAnnouncement(c *gin.Context) {
	rdb, err := GetDBFromContext.GetRedisFromContext(c)
	if err != nil {
		rdb = Database.ReconnectToRedisRegularly()
		if rdb == nil {
			log.Println("重连redis失败")
			c.String(http.StatusInternalServerError, "重连redis失败")
			return
		}
	}
	var annList []Model.Announcement
	err = rdb.ZRange(context.Background(), "ann", 0, -1).ScanSlice(&annList)
	if err != nil {
		log.Println("从Redis读取公告出错!", err)
		c.String(http.StatusAccepted, err.Error())
		return
	}
	c.JSON(http.StatusOK, annList)
}
