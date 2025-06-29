package Model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type VenueOrderDetail struct {
	ID    int      `json:"id"`
	Time  []string `json:"time"`
	Price int      `json:"price"`
}
type Order struct {
	ID          int64            `json:"id"`
	UserID      int64            `json:"userID"`
	State       string           `gorm:"type:varchar(8)" json:"state"`
	VenueType   string           `gorm:"type:varchar(10)" json:"venueType"`
	Count       int              `gorm:"type:tinyint" json:"count"`
	BookingDate string           `gorm:"type:varchar(10)" json:"bookingDate"`
	UseDate     string           `gorm:"type:varchar(10)" json:"useDate"`
	FinishTime  int              `gorm:"type:tinyint" json:"finishTime"`
	Venue1      VenueOrderDetail `gorm:"type:json" json:"venue1"`
	Venue2      VenueOrderDetail `gorm:"type:json" json:"venue2"`
	CancelFlag  bool             `json:"cancelFlag"`
	CreatedAt   time.Time        `json:"createdAt"`
	ExpireAt    time.Time        `json:"expireAt"`
	FinishedAt  time.Time        `json:"finishedAt"`
}

func (v *VenueOrderDetail) Scan(value interface{}) error {
	byteValue, _ := value.([]byte) //类型断言，断定为[]byte类型，我们在value方法中也是转换为[]byte类型输入到数据库中的
	var receiver VenueOrderDetail
	err := json.Unmarshal(byteValue, &receiver) //反序列化，将[]byte类型转化为我们需要的结构体
	if err != nil {
		return err
	}
	//fmt.Println(receiver)
	*v = receiver //将其内容传输给venue
	return nil
}

// Value 存入数据库，将json转换为数据库可接受类型数据，实现driver.Valuer接口
func (v VenueOrderDetail) Value() (driver.Value, error) {
	return json.Marshal(v) //由结构体转换为json类型数据，返回[]byte

}
