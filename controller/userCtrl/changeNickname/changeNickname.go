package changeNickname

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"Gym_booking_WeChat_mini_program/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ChangeNickname(c *gin.Context) {
	mysql, err := getDBFromContext.GetGormDBFromContext(c, "mysql")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.String(http.StatusBadRequest, "不存在该用户!")
		return
	}
	id := user.(*utils.UserClaims).ID
	nickname := c.PostForm("nickname")
	mysql.Table("users").Where("id = ?", id).Update("nickname", nickname)
	c.Status(http.StatusOK)
}
