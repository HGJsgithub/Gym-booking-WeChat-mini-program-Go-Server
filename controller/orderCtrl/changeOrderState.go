package orderCtrl

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func ChangeOrderState(c *gin.Context) {
	mysql, err := getDBFromContext.GetGormDBFromContext(c, "writeMySQL")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)

	state := c.PostForm("state")
	if mysql.Table("orders").Where("id = ?", id).RecordNotFound() == false {
		mysql.Table("orders").Where("id = ?", id).Update("state", state)
		//成功找到订单并修改相应的订单状态
		c.Status(http.StatusOK)
	} else {
		//没有找到相应的订单
		c.Status(http.StatusNotFound)
	}
}
