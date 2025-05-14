package database_connect

import (
	"Riffy-music/src/database/models"
	"fmt"
)

func Test_connection() {

	db := DBConnect()
	var (
		username string
		password string
		email    string
		avatar   string
		birthday int64
		country  string
	)

	fmt.Scan(&username, &password, &email, &avatar, &birthday, &country)

	err := db.Create(&models.Users{Username: username, Password_hash: password, Email: email, Avatar_url: avatar, Birth_year: birthday, Country: country})

	if err != nil {
		fmt.Errorf("Проблема с подключением к бд")
		return

	}

}
