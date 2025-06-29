package ChangeOrderState

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"Gym_booking_WeChat_mini_program/Controller/VenueCtrl/ChangeVenueState"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strconv"
)

func ChangeOrderState(c *gin.Context) {
	mysql, err := GetDBFromContext.GetGormDBFromContext(c, "writeMySQL")
	if err != nil {
		mysql = Database.ReconnectToMySQLRegularly()
		if mysql == nil {
			log.Println("重连MySQL失败")
			c.String(http.StatusInternalServerError, "重连MySQL失败")
			return
		}
	}

	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)

	state := c.PostForm("state")
	var count int64
	mysql.Table("orders").Where("id = ?", id).Count(&count)
	if count > 0 {
		//说明用户提前取消了订单，需要把场地重新设为可预约
		if state == "已取消" {
			rdb := Database.ConnectToRedis()
			defer func(rdb *redis.Client) {
				_ = rdb.Close()
			}(rdb)
			if rdb == nil {
				log.Println("连接redis失败！")
				c.String(http.StatusInternalServerError, "连接redis失败！")
				return
			}
			tmpID := strconv.FormatInt(id, 10)
			var updateReq ChangeVenueState.UpdateRequest
			//从redis获取要修改场地状态的信息
			err = rdb.Get(context.Background(), tmpID).Scan(&updateReq)
			if err != nil {
				log.Println("从redis获取场地状态更改条件失败！", err)
				c.String(http.StatusInternalServerError, "从redis获取场地状态更改条件失败！")
				return
			}
			updateInfo := updateReq.UpdateInfo
			//事务修改
			err = ChangeVenueState.Transaction(mysql, updateInfo, false)
			if err != nil {
				log.Println("更改场地状态失败！", err)
				c.String(http.StatusInternalServerError, "更改场地状态失败！")
				return
			}
			//删除对应订单的场地状态修改信息
			rdb.Del(context.Background(), tmpID)
		}
		mysql.Table("orders").Where("id = ?", id).Update("state", state)
		//成功找到订单并修改相应的订单状态
		c.Status(http.StatusOK)
	} else {
		//没有找到相应的订单
		c.Status(http.StatusNotFound)
	}
}
