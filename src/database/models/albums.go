package models

import "time"

type Albums struct {
	Id           int64  `gorm:"primary_key;auto_increment"`
	Title        string `gorm:"NOT NULL"`
	Artist_id    int64  `gorm:"REFERENCES users(id) ON DELETE CASCADE"`
	Cover_url    string
	Release_date time.Time `gorm:"DEFAULT:NOW()"`
}
