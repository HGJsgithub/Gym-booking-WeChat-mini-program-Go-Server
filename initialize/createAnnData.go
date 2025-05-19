package initialize

import (
	"Gym_booking_WeChat_mini_program/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

func CreateAnnData(db *gorm.DB) {
	annList := []model.Announcement{
		{
			ID:      1,
			Title:   "公告一",
			Content: "这是第一条公告",
		},
		{
			ID:      2,
			Title:   "公告二",
			Content: "这是第二条公告",
		},
		{
			ID:      3,
			Title:   "公告三",
			Content: "这是第三条公告",
		},
	}
	for i := 0; i < len(annList); i++ {
		ann := annList[i]
		if db.First(&ann).RecordNotFound() == true {
			fmt.Println(i+1, "没有这条公告！")
			db.Create(&model.Announcement{
				ID:      i + 1,
				Title:   annList[i].Title,
				Content: annList[i].Content,
			})
		} else {
			fmt.Println(i+1, "公告已经存在！")
			fmt.Println()
		}
	}
}
