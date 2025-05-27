package routes

import (
	"Gym_booking_WeChat_mini_program/config/database"
	"Gym_booking_WeChat_mini_program/config/messageQueue"
	"Gym_booking_WeChat_mini_program/controller/orderCtrl"
	"Gym_booking_WeChat_mini_program/middleware"
	"github.com/gin-gonic/gin"
)

func InitOrderRoute(r *gin.Engine) {
	readMySQL := database.ConnectToMySQL()
	writeMySQL := database.ConnectToMySQL()
	rbmq := messageQueue.ConnectToRabbitMQ()
	redis := database.ConnectToRedis()

	orderRoute := r.Group("/order")
	orderRoute.Use(userAuth())
	//获取订单数据
	orderRoute.GET("/list", middleware.SetMySQL(readMySQL, "readMySQL"), orderCtrl.GetOrderList)
	{
		changeOrderRoute := orderRoute.Group("/changeOrder").Use(middleware.SetMySQL(writeMySQL, "writeMySQL"))
		{
			//保存订单数据
			changeOrderRoute.POST("/saveOrder", middleware.SetRabbitMQ(rbmq), middleware.SetRedis(redis), orderCtrl.SaveOrder)
			//改变订单状态
			changeOrderRoute.POST("/changeOrderState", orderCtrl.ChangeOrderState)
			//改变订单的取消状态
			changeOrderRoute.POST("/changeCancelFlag", orderCtrl.ChangeCancelFlag)
			//删除订单
			changeOrderRoute.POST("/delete", orderCtrl.DeleteOrder)
		}
	}
}
