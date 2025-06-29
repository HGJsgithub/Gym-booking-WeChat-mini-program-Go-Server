package UserLogin

import (
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"Gym_booking_WeChat_mini_program/Model"
	"Gym_booking_WeChat_mini_program/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserLogin(c *gin.Context) {
	mysql, err := GetDBFromContext.GetGormDBFromContext(c, "mysql")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	phone := c.PostForm("phone")
	password := c.PostForm("password")
	log.Println("phone:", phone, "password:", password)
	var user Model.User
	userExist, pwdRight := SearchUser(mysql, phone, password, &user)
	if userExist == true && pwdRight == true {
		token, err := Utils.GenerateToken(user.ID)
		if err != nil {
			log.Println("生成token出错！", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"userInfo": user,
			"token":    token,
		})
		return
	}
	if userExist == true && pwdRight == false {
		//密码错误
		c.Status(http.StatusUnauthorized)
		return
	}
	//用户未注册
	c.Status(http.StatusMethodNotAllowed)
}
