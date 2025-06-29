package ChangeVenueState

import (
	"encoding/json"
	"fmt"
)

type UpdateInfo struct {
	VenueType string   `json:"venueType"`
	Date      string   `json:"date"`
	ID        int      `json:"id"`
	TimeSlot  []string `json:"timeSlot"`
	State     bool     `json:"state"`
}

type UpdateRequest struct {
	UpdateInfo []UpdateInfo `json:"updateInfo"`
	OrderID    int64        `json:"orderID"`
}

func (u *UpdateRequest) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *UpdateRequest) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

// BuildUpdateStatement state是目标状态
func BuildUpdateStatement(info *UpdateInfo, state bool) (string, []interface{}, map[string]interface{}) {
	// 构建 WHERE 条件：所有目标字段必须为 false
	whereCondition := "venue_type = ? AND date = ? AND id = ?"
	args := []interface{}{info.VenueType, info.Date, info.ID}

	// 构建更新字段：将目标字段设为 true
	updateMap := make(map[string]interface{})
	for _, timeSlot := range info.TimeSlot {
		whereCondition += fmt.Sprintf(" AND %s = ?", timeSlot)
		args = append(args, !state)
		updateMap[timeSlot] = state
	}
	return whereCondition, args, updateMap
}
