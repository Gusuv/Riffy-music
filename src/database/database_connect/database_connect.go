package database_connect

import (
	Configuration "Riffy-music/src/config"
	"Riffy-music/src/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func DBConnect() *gorm.DB {
	Configuration.Init()

	db_name := os.Getenv("DATABASE_DSN")

	db, err := gorm.Open(postgres.Open(db_name), &gorm.Config{})
	if err != nil {
		log.Panic(err, "Нет подключения к БД")
	}
	err = db.AutoMigrate(&models.Users{})

	if err != nil {
		log.Panic("Проблема с миграцией")

	}

	return db

}
