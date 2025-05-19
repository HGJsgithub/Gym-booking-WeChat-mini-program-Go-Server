package annCtrl

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"Gym_booking_WeChat_mini_program/model"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAnn(c *gin.Context) {
	rdb, err := getDBFromContext.GetRedisFromContext(c)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
	}
	var annList []model.Announcement
	err = rdb.ZRange(context.Background(), "ann", 0, -1).ScanSlice(&annList)
	if err != nil {
		log.Println("从Redis读取公告出错!", err)
		c.String(http.StatusAccepted, err.Error())
		return
	}
	c.JSON(http.StatusOK, annList)
}
