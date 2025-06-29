package Registration

import (
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"Gym_booking_WeChat_mini_program/Model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Registration(c *gin.Context) {
	mysql, err := GetDBFromContext.GetGormDBFromContext(c, "mysql")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var newUser Model.User
	err = c.ShouldBindJSON(&newUser)
	if err != nil {
		log.Println("绑定user到结构体出错！", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(newUser)
	var count int64
	mysql.Where("phone", newUser.Phone).First(&newUser).Count(&count)
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
