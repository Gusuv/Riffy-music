package models

import "time"

type Users struct {
	Id              int64  `gorm:"primary_key;auto_increment"`
	Username        string `gorm:"not null;unique"`
	Password_hash   string `gorm:"not null"`
	Email           string `gorm:"not null;unique"`
	Avatar_url      string
	Birth_year      int64     `gorm:"not null"`
	Country         string    `gorm:"not null"`
	registered_date time.Time `gorm:"default:now()"`
}
