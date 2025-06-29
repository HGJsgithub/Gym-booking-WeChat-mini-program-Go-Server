package SaveOrder

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"Gym_booking_WeChat_mini_program/Model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func createFinishedTime(date string, hour int) time.Time {
	tmp := strings.Split(date, "-")
	year, _ := strconv.Atoi(tmp[0])
	month, _ := strconv.Atoi(tmp[1])
	day, _ := strconv.Atoi(tmp[2])
	return time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.Local)
}

func SaveOrder(c *gin.Context) {
	mysql, err := GetDBFromContext.GetGormDBFromContext(c, "writeMySQL")
	if err != nil {
		mysql = Database.ReconnectToMySQLRegularly()
		if mysql == nil {
			log.Println("重连MySQL失败")
			c.String(http.StatusInternalServerError, "重连MySQL失败")
			return
		}
	}

	var order Model.Order
	err = c.ShouldBindJSON(&order)
	if err != nil {
		log.Println("绑定order到结构体出错！", err)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	order.CreatedAt = time.Now()                      //设置订单创建时间
	order.ExpireAt = time.Now().Add(30 * time.Second) //设置订单未支付自动取消的时长
	order.FinishedAt = createFinishedTime(order.UseDate, order.FinishTime)
	mysql.Save(&order)
	//超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var rdb *redis.Client
	rdb, err = GetDBFromContext.GetRedisFromContext(c)
	if err != nil {
		log.Println("从context获取redis失败！", err)
		c.String(http.StatusInternalServerError, err.Error())
		cancel()
		return
	}
	c.Status(http.StatusOK)
	go timeoutHandle(rdb, ctx, mysql, order.ID, cancel)
}
