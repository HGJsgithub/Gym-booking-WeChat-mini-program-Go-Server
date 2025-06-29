package Routes

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/OrderCtrl/ChangeCancelFlag"
	"Gym_booking_WeChat_mini_program/Controller/OrderCtrl/ChangeOrderState"
	"Gym_booking_WeChat_mini_program/Controller/OrderCtrl/DeleteOrder"
	"Gym_booking_WeChat_mini_program/Controller/OrderCtrl/GetOrderList"
	"Gym_booking_WeChat_mini_program/Controller/OrderCtrl/SaveOrder"
	"Gym_booking_WeChat_mini_program/Middleware"
	"github.com/gin-gonic/gin"
)

func InitOrderRoute(r *gin.Engine) {
	readMySQL := Database.ConnectToMySQL()
	writeMySQL := Database.ConnectToMySQL()
	redis := Database.ConnectToRedis()

	orderRoute := r.Group("/order")
	orderRoute.Use(userAuth())
	//获取订单数据
	orderRoute.GET("/list", Middleware.SetMySQL(readMySQL, "readMySQL"), GetOrderList.GetOrderList)
	{
		changeOrderRoute := orderRoute.Group("/changeOrder").Use(Middleware.SetMySQL(writeMySQL, "writeMySQL"))
		{
			//保存订单数据
			changeOrderRoute.POST("/saveOrder", Middleware.SetRedis(redis), SaveOrder.SaveOrder)
			//改变订单状态
			changeOrderRoute.POST("/changeOrderState", ChangeOrderState.ChangeOrderState)
			//改变订单的取消状态
			changeOrderRoute.POST("/changeCancelFlag", ChangeCancelFlag.ChangeCancelFlag)
			//删除订单
			changeOrderRoute.POST("/delete", DeleteOrder.DeleteOrder)
		}
	}
}
