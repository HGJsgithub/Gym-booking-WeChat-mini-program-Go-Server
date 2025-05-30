package routes

import (
	"Gym_booking_WeChat_mini_program/config/database"
	"Gym_booking_WeChat_mini_program/controller/userCtrl/avatar"
	"Gym_booking_WeChat_mini_program/controller/userCtrl/changeNickname"
	"Gym_booking_WeChat_mini_program/controller/userCtrl/changePassword"
	"Gym_booking_WeChat_mini_program/controller/userCtrl/registration"
	"Gym_booking_WeChat_mini_program/controller/userCtrl/userLogin"
	"Gym_booking_WeChat_mini_program/middleware"
	"Gym_booking_WeChat_mini_program/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func userAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			fmt.Println("token是空的！")
			c.AbortWithStatusJSON(http.StatusAccepted, gin.H{"error": "Invalid token"})
			return
		}
		userClaims, err := utils.ValidateTokenWithCustomClaims(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusAccepted, gin.H{"error": "Invalid token"})
			fmt.Println("解析JWT出错！", err)
			return
		}
		c.Set("user", userClaims)
		c.Next()
	}
}

func InitUserRoute(r *gin.Engine) {
	mysql := database.ConnectToMySQL()
	userRoute := r.Group("/user").Use(middleware.SetMySQL(mysql, "mysql"))
	{
		//用户注册
		userRoute.POST("/registration", registration.Registration)
		//用户登录
		userRoute.POST("/login", userLogin.UserLogin)
		//获取头像
		userRoute.Use(userAuth()).Static("/avatar", "./avatar")
		//修改头像
		userRoute.POST("/avatar", avatar.SaveAvatar)
		//修改用户昵称
		userRoute.POST("/changeNickname", changeNickname.ChangeNickname)
		//修改用户密码
		userRoute.POST("/changePassword", changePassword.ChangePassword)
	}
}
