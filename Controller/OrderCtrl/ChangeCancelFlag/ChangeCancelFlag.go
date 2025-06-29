package ChangeCancelFlag

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// ChangeCancelFlag 修改订单是否可撤销的标记
func ChangeCancelFlag(c *gin.Context) {
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
	var count int64
	mysql.Table("orders").Where("id = ?", id).Count(&count)
	if count > 0 {
		mysql.Table("orders").Where("id=?", id).Update("cancel_flag", true)
		//成功找到订单并修改相应的订单的可取消标记
		c.Status(http.StatusOK)
	} else {
		//没有找到相应的订单
		c.Status(http.StatusNotFound)
	}
}
