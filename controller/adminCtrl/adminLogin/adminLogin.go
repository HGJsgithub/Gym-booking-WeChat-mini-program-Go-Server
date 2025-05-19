package adminLogin

//func AdminVerification(c *gin.Context) {
//	db := database.ConnectToMySQL()
//	account := c.PostForm("account")
//	password := c.PostForm("password")
//	isAccRight := CheckAccount(db, account)
//	isPwdRight := CheckAdminPwd(db, password)
//
//	if isAccRight && isPwdRight {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 0,
//			"msg":  "账号密码正确",
//		})
//	}
//
//	if isAccRight == false {
//		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
//			"code": 1,
//			"msg":  "账号错误！",
//		})
//	} else if isPwdRight == false {
//		c.JSON(http.StatusUnauthorized, gin.H{
//			"code": 2,
//			"msg":  "密码错误！",
//		})
//	}
//	defer func(db *gorm.DB) {
//		err := db.Close()
//		if err != nil {
//			fmt.Println("断开数据库连接出错：", err)
//		}
//	}(db)
//
//}
