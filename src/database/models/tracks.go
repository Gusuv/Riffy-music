package models

import "time"

type Tracks struct {
	Id               int64  `gorm:"primary_key;auto_increment"`
	Title            string `gorm:"NOT NULL"`
	Artist_id        int64  `gorm:"REFERENCES users(id) ON DELETE CASCADE"`
	File_url         string `gorm:"NOT NULL"`
	Cover_url        string
	Duration         int64     `gorm:"NOT NULL"`
	Publication_date time.Time `gorm:"DEFAULT:NOW()"`
	Genre_id         int64     `gorm:"REFERENCES genres(id)"`
	Album_id         int64     `gorm:"REFERENCES albums(id)"`
}
