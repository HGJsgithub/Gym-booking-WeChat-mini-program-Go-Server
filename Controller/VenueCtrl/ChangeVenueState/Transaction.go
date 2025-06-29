package ChangeVenueState

import (
	"fmt"
	"gorm.io/gorm"
)

func Transaction(db *gorm.DB, info []UpdateInfo, state bool) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, updateReq := range info {
			whereCondition, args, updateMap := BuildUpdateStatement(&updateReq, state)
			res := tx.Table("venue_states").Where(whereCondition, args...).Updates(updateMap)
			if res.Error != nil {
				//修改场地状态出错
				return fmt.Errorf("update error")
			}
			if res.RowsAffected == 0 {
				return fmt.Errorf("场地 %d 已被预约", updateReq.ID)
			}
		}
		return nil
	})
}
