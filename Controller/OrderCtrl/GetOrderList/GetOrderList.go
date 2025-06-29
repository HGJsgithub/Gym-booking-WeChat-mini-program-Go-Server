package GetOrderList

import (
	"Gym_booking_WeChat_mini_program/ConstDefine"
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"Gym_booking_WeChat_mini_program/Model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetOrderList(c *gin.Context) {
	mysql, err := GetDBFromContext.GetGormDBFromContext(c, "readMySQL")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	userID, _ := strconv.ParseInt(c.Query("userID"), 10, 64)
	state := c.Query("state")
	var orderList []Model.Order
	mysql.Where(map[string]interface{}{"user_id": userID, "state": state}).Find(&orderList)
	if state == "待使用" {
		now := time.Now()
		for i, order := range orderList {
			if order.FinishedAt.Before(now) {
				mysql.Table("orders").Where("id = ?", order.ID).Update("state", "已完成")
				orderList = append(orderList[:i], orderList[i+1:]...)
			}
		}
	}
	if len(orderList) > 0 && orderList[0].State == ConstDefine.Pending {
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
