package SaveOrder

import (
	"Gym_booking_WeChat_mini_program/ConstDefine"
	"Gym_booking_WeChat_mini_program/Controller/VenueCtrl/ChangeVenueState"
	"Gym_booking_WeChat_mini_program/Model"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func timeoutHandle(rdb *redis.Client, ctx context.Context, db *gorm.DB, orderID int64, cancel context.CancelFunc) {
	defer cancel()
	select {
	case <-ctx.Done():
		unpaidOrder := Model.Order{}
		// 查询订单状态
		db.Table("orders").Where("id = ?", orderID).First(&unpaidOrder)
		tmpID := strconv.FormatInt(unpaidOrder.ID, 10)
		if unpaidOrder.State == ConstDefine.Pending {
			var updateReq ChangeVenueState.UpdateRequest
			//从redis获取要修改场地状态的信息
			err := rdb.Get(context.Background(), tmpID).Scan(&updateReq)
			if err != nil {
				log.Println("从redis获取场地状态更改条件失败！", err)
				return
			}
			updateInfo := updateReq.UpdateInfo
			//事务修改
			err = ChangeVenueState.Transaction(db, updateInfo, false)
			if err != nil {
				log.Println("更改场地状态失败！", err)
				return
			}
			//将订单改为已取消
			db.Model(&unpaidOrder).Where("id = ?", orderID).Update("state", "已取消")
		}
		//删除对应订单的场地状态修改信息
		rdb.Del(context.Background(), tmpID)
		return
	}
}
