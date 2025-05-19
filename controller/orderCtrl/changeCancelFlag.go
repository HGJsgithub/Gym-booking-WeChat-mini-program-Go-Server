package orderCtrl

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// ChangeCancelFlag 修改订单是否可撤销的标记
func ChangeCancelFlag(c *gin.Context) {
	mysql, err := getDBFromContext.GetGormDBFromContext(c, "writeMySQL")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if mysql.Table("orders").Where("id = ?", id).RecordNotFound() == false {
		mysql.Table("orders").Where("id=?", id).Update("cancel_flag", true)
		//成功找到订单并修改相应的订单的可取消标记
		c.Status(http.StatusOK)
	} else {
		//没有找到相应的订单
		c.Status(http.StatusNotFound)
	}
}
