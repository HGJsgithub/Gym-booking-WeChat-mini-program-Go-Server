package Avatar

import (
	"Gym_booking_WeChat_mini_program/Controller/GetDBFromContext"
	"Gym_booking_WeChat_mini_program/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SaveAvatar(c *gin.Context) {
	mysql, err := GetDBFromContext.GetGormDBFromContext(c, "mysql")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		log.Println("未获取头像！", err)
		c.String(http.StatusBadRequest, "未获取头像！")
		return
	}
	user, exists := c.Get("user")
	if !exists {
		c.String(http.StatusBadRequest, "不存在该用户!")
		return
	}
	id := user.(*Utils.UserClaims).ID
	strID := strconv.FormatInt(id, 10)
	avatarType := "png"
	avatarName := strID + "." + avatarType
	savePath := "./avatar/" + avatarName
	err = c.SaveUploadedFile(avatar, savePath)
	if err != nil {
		log.Println("保存头像出错！", err)
		c.String(http.StatusInternalServerError, "保存头像出错")
		return
	}
	avatarSRC := "https://credible-halibut-sound.ngrok-free.app/user/avatar/" + avatarName
	mysql.Table("users").Where("id = ?", id).Update("avatar_src", avatarSRC)
	c.String(http.StatusOK, avatarSRC)
}
