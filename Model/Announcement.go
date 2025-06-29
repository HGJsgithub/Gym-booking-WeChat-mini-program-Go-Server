package Model

import "encoding/json"

type Announcement struct {
	ID      int    `json:"id"`
	Title   string `gorm:"size:255" json:"title"`
	Content string `gorm:"size:4096" json:"content"`
}

func (a *Announcement) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Announcement) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}
