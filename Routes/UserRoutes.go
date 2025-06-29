package Routes

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Controller/UserCtrl/Avatar"
	"Gym_booking_WeChat_mini_program/Controller/UserCtrl/ChangeNickname"
	"Gym_booking_WeChat_mini_program/Controller/UserCtrl/ChangePassword"
	"Gym_booking_WeChat_mini_program/Controller/UserCtrl/Registration"
	"Gym_booking_WeChat_mini_program/Controller/UserCtrl/UserLogin"
	"Gym_booking_WeChat_mini_program/Middleware"
	"Gym_booking_WeChat_mini_program/Utils"
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
		userClaims, err := Utils.ValidateTokenWithCustomClaims(token)
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
	mysql := Database.ConnectToMySQL()
	userRoute := r.Group("/user").Use(Middleware.SetMySQL(mysql, "mysql"))
	{
		//用户注册
		userRoute.POST("/registration", Registration.Registration)
		//用户登录
		userRoute.POST("/login", UserLogin.UserLogin)
		//获取头像
		userRoute.Use(userAuth()).Static("/avatar", "./avatar")
		//修改头像
		userRoute.POST("/avatar", Avatar.SaveAvatar)
		//修改用户昵称
		userRoute.POST("/changeNickname", ChangeNickname.ChangeNickname)
		//修改用户密码
		userRoute.POST("/changePassword", ChangePassword.ChangePassword)
	}
}
