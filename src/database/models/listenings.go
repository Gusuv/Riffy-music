package models

import "time"

type Listenings struct {
	Id         int64 `gorm:"primaryKey"`
	Track_id   int64 `gorm:"REFERENCES tracks(id)"`
	Users_id   int64 `gorm:"REFERENCES users(id)"`
	List_count int64
	List_at    time.Time `gorm:"DEFAULT NOW()"`
}
