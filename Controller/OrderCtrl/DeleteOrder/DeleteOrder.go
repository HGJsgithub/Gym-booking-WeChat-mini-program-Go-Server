package DeleteOrder

import (
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"Gym_booking_WeChat_mini_program/Model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DeleteOrder(c *gin.Context) {
	mysql, err := GetDBFromContext.GetGormDBFromContext(c, "writeMySQL")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	order := Model.Order{ID: id}
	mysql.Delete(&order)
	c.Status(http.StatusOK)
}
