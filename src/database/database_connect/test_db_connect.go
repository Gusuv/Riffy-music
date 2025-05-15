package Dbconnect

import (
	"Riffy-music/src/database/models"
	"log"
)

func TestConnection() {

	db := DBConnect()
	var (
		username string = "Test2"
		password string = "Test2"
		email    string = "test2@test.com"
		avatar   string = "https://avatars2.githubusercontent.com/u/999"
		birthday int64  = 2000
		country  string = "China"
	)
	err := db.Create(&models.Users{Username: username, Password_hash: password, Email: email, Avatar_url: avatar, Birth_year: birthday, Country: country})

	if err != nil {
		log.Panic(err, "Данные не были добавлены")
		return

	}

}
