package orderCtrl

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"Gym_booking_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DeleteOrder(c *gin.Context) {
	mysql, err := getDBFromContext.GetGormDBFromContext(c, "writeMySQL")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	order := model.Order{ID: id}
	mysql.Delete(&order)
	c.Status(http.StatusOK)
}
