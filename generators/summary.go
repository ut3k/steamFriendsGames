package generators

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GenerateHTMLlist() {
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Function CheckIfGameIsCoop can NOT connect to data.db")
	}

}
