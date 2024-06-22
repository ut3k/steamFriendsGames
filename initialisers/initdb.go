package initialisers

import (
	"fmt"

	"steamFriendsGames/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDataBase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err == nil {
		fmt.Println("connection ESTABLISHED")
	}

	if err != nil {
		fmt.Println("Failed to connect to DataBase")

	}

}

func SyncDB() {
	DB.AutoMigrate(&models.Game{})
	DB.AutoMigrate(&models.User{})
}
