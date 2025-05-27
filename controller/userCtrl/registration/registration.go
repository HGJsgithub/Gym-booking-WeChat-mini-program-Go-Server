package registration

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"Gym_booking_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Registration(c *gin.Context) {
	mysql, err := getDBFromContext.GetGormDBFromContext(c, "mysql")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var newUser model.User
	err = c.ShouldBindJSON(&newUser)
	if err != nil {
		log.Println("绑定user到结构体出错！", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(newUser)
	var count int64
	mysql.First(&newUser).Count(&count)
	if count == 0 {
		newUser.ID, _ = strconv.ParseInt(newUser.Phone, 10, 64)
		mysql.Create(&newUser)
		c.JSON(http.StatusCreated, gin.H{
			"newUser": newUser,
		})
	} else {
		//说明已经注册
		c.Status(http.StatusForbidden)
	}
}
