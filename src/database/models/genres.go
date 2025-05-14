package models

type Genres struct {
	Id         int    `gorm:"primary_key"`
	Genre_name string `gorm:"NOT NULL"`
}
