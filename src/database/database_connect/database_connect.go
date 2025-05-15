package Dbconnect

import (
	Config "Riffy-music/src/config"
	"Riffy-music/src/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DBConnect() *gorm.DB {
	Config.Init()

	db, err := gorm.Open(postgres.Open(Config.DatabaseDns()), &gorm.Config{})
	if err != nil {
		log.Panic(err, "Нет подключения к БД")
		return nil
	}
	err = db.AutoMigrate(&models.Users{})

	if err != nil {
		log.Panic("Проблема с миграцией")

	}

	return db

}
