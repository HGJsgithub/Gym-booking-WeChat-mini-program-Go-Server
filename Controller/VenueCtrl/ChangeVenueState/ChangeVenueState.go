package ChangeVenueState

import (
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strconv"
)

func ChangeVenueState(c *gin.Context) {
	mysql, err := GetDBFromContext.GetGormDBFromContext(c, "writeMySQL")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusAccepted, gin.H{
			"serverErr": true,
			"err":       err.Error(),
		})
		return
	}
	var rdb *redis.Client
	rdb, err = GetDBFromContext.GetRedisFromContext(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusAccepted, gin.H{
			"serverErr": true,
			"err":       err.Error(),
		})
		return
	}

	var updateReq UpdateRequest
	err = c.ShouldBindJSON(&updateReq)
	if err != nil {
		log.Println("绑定updateRequest结构体出错!", err)
		c.JSON(http.StatusAccepted, gin.H{
			"serverErr": true,
			"err":       "绑定updateRequest结构体出错!",
		})
		return
	}
	log.Println("UpdateReq:", updateReq)

	ctx := context.Background()
	updateReqMarshal, _ := updateReq.MarshalBinary()
	rdb.Set(ctx, strconv.FormatInt(updateReq.OrderID, 10), updateReqMarshal, 0)

	err = Transaction(mysql, updateReq.UpdateInfo, true)
	if err != nil {
		log.Println("err:", err)
		if err.Error() == "update error" {
			c.JSON(http.StatusAccepted, gin.H{
				"serverErr": true,
				"err":       err.Error(),
			})
		} else {
			c.JSON(http.StatusAccepted, gin.H{
				"serverErr": false,
				"msg":       err.Error(),
			})
		}
	} else {
		c.Status(http.StatusOK)
	}
}
