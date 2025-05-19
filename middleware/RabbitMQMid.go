package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitMQMid(rbmq *amqp.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		if rbmq == nil {
			fmt.Println("RabbitMQMid未连接到RabbitMQ！")
			return
		}
		c.Set("rabbitMQ", rbmq)
	}
}
