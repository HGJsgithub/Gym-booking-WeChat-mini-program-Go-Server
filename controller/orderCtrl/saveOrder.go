package orderCtrl

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"Gym_booking_WeChat_mini_program/controller/orderCtrl/overdueOrder"
	"Gym_booking_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"time"
)

func SaveOrder(c *gin.Context) {
	mysql, err := getDBFromContext.GetGormDBFromContext(c, "writeMySQL")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var order model.Order
	err = c.ShouldBindJSON(&order)
	if err != nil {
		log.Println("绑定order到结构体出错！", err)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	order.CreatedAt = time.Now()
	order.ExpireAt = time.Now().Add(30 * time.Second)
	rabbitMQ, exist := c.Get("rabbitMQ")
	rbmq := rabbitMQ.(*amqp.Connection)
	if !exist {
		log.Println("未连接到RabbitMQ!")
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "未连接到RabbitMQ!!",
		})
	}
	err = overdueOrder.DepositOrderIntoMQ(order.ID, rbmq)
	if err != nil {
		log.Println("将订单消息存入RabbitMQ出错！", err)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	mysql.Save(&order)
	c.Status(http.StatusOK)
	redis, err := getDBFromContext.GetRedisFromContext(c)
	if err != nil {
		log.Println("从context获取redis失败！", err)
		log.Fatal(err)
	}
	go overdueOrder.ListenUnpaidOrderQueue(rbmq, mysql, redis)
}
