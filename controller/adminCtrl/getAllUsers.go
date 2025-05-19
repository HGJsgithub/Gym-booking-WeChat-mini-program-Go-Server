package adminCtrl

//func GetAllUsers(c *gin.Context) {
//	db := database.ConnectToMySQL()
//	var userList []model.User
//	db.Find(&userList)
//	c.JSON(http.StatusOK, userList)
//	defer func(db *gorm.DB) {
//		err := db.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(db)
//
//}
