package Database

import (
	"Gym_booking_WeChat_mini_program/Model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetMySQLDataToRedis(db *gorm.DB, rdb *redis.Client) {
	if db == nil {
		fmt.Println("SetMySQLDataToRedis未连接到MySQL！")
		return
	}
	if rdb == nil {
		fmt.Println("SetMySQLDataToRedis未连接到Redis！")
		return
	}
	var annList []Model.Announcement
	db.Find(&annList)
	for _, ann := range annList {
		annMarshal, _ := json.Marshal(ann)
		err := rdb.ZAdd(context.Background(), "ann", redis.Z{
			Score:  float64(ann.ID),
			Member: annMarshal,
		}).Err()
		if err != nil {
			fmt.Println("Redis ZAdd Error", err.Error())
		}
	}
}
