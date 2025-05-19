package orderCtrl

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"Gym_booking_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetOrderList(c *gin.Context) {
	mysql, err := getDBFromContext.GetGormDBFromContext(c, "readMySQL")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	userID, _ := strconv.ParseInt(c.Query("userID"), 10, 64)
	state := c.Query("state")
	var orderList []model.Order
	mysql.Where(map[string]interface{}{"user_id": userID, "state": state}).Find(&orderList)
	if len(orderList) > 0 && orderList[0].State == "待支付" {
		remainingTime := int(orderList[0].ExpireAt.Sub(time.Now()).Seconds())
		remainingMinutes := remainingTime / 60
		remainingSeconds := remainingTime % 60
		c.JSON(http.StatusOK, gin.H{
			"orderList":        orderList,
			"remainingMinutes": remainingMinutes,
			"remainingSeconds": remainingSeconds,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"orderList": orderList,
		})
	}
}
