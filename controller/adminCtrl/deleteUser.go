package adminCtrl

//func DeleteUser(c *gin.Context) {
//	db := database.ConnectToMySQL()
//	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
//	user := model.User{
//		ID: id,
//	}
//	db.Delete(&user)
//	c.JSON(http.StatusOK, gin.H{
//		"msg": "成功删除！",
//	})
//	defer func(db *gorm.DB) {
//		err := db.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(db)
//
//}
