package VenueModel

import (
	"Gym_booking_WeChat_mini_program/Config/Database"
	"Gym_booking_WeChat_mini_program/Model"
	"gorm.io/gorm"
	"unsafe"
)

type VenueState struct {
	VenueType string `gorm:"primaryKey;type:char(16)" json:"venueType"`
	Date      string `gorm:"primaryKey;type:char(8)" json:"date"`
	ID        int    `gorm:"primaryKey;autoIncrement:false" json:"id"`
	T9        bool   `json:"T9"`
	T10       bool   `json:"T10"`
	T11       bool   `json:"T11"`
	T12       bool   `json:"T12"`
	T13       bool   `json:"T13"`
	T14       bool   `json:"T14"`
	T15       bool   `json:"T15"`
	T16       bool   `json:"T16"`
	T17       bool   `json:"T17"`
	T18       bool   `json:"T18"`
	T19       bool   `json:"T19"`
	T20       bool   `json:"T20"`
	T21       bool   `json:"T21"`
}

func (V *VenueState) VenueStateStructToSlice(structure []VenueState) [][]byte {
	var rawState [][]byte
	var byteSlice []byte
	Len := unsafe.Sizeof(VenueState{})
	for i := range structure {
		sliceMock := &Model.SliceMock{
			Addr: uintptr(unsafe.Pointer(&structure[i])),
			Cap:  int(Len),
			Len:  int(Len),
		}
		byteSlice = *(*[]byte)(unsafe.Pointer(sliceMock))
		rawState = append(rawState, byteSlice)
	}
	return rawState
}

func (V *VenueState) UpdateVenueStateEveryday(venueType string) {
	db := Database.ConnectToMySQL()
	var vs []VenueState
	db.Delete(V, "date = ?", "today")
	db.Where("venue_type = ?", venueType).Find(&vs)
	venueNum := len(vs)
	db.Model(V).Where("venue_type = ? AND date = ?", venueType, "tomorrow").Update("date", "today")
	for i := 0; i < venueNum; i++ {
		newVS := VenueState{VenueType: venueType, ID: i + 1, Date: "tomorrow"}
		db.Save(&newVS)
	}
}

func (V *VenueState) UpdateVenueStateEveryHour(timeField string, mysql *gorm.DB) {
	mysql.Model(&VenueState{}).Where("date = ?", "today").Update(timeField, true)
}
