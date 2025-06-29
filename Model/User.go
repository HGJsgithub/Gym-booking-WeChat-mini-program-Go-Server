package Model

import (
	"time"
)

type User struct {
	ID        int64  `json:"id"`
	Nickname  string `gorm:"size:11" json:"nickname"`
	Phone     string `gorm:"type:char(11)" json:"phone"`
	Password  string `gorm:"size:16" json:"password"`
	AvatarSRC string `gorm:"size:255" json:"avatarSRC"`
	CreatedAt time.Time
}
