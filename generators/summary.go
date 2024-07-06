package generators

import (
	"fmt"
	"html/template"
	"os"

	"steamFriendsGames/models"

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

	var GameDATA []models.Game

	// SQL promt is incorrect,
	err = DB.Preload("Users").Find(&GameDATA).Where("is_cooperative = ?", 1).Error
	//returning wrong data
	if err != nil {
		fmt.Println("error fetching data from database:", err)
		return
	}

	t, err := template.ParseFiles("templates/coop_game_list.html")
	if err != nil {
		fmt.Println("can NOT pare template, error:", err)
	}

	file, err := os.Create("coop_games.html")
	if err != nil {
		fmt.Println("can NOT create file", file.Name())
	}
	defer file.Close()

	err = t.Execute(file, GameDATA)
	if err != nil {
		fmt.Println("error executing tempalte:", err)
	}

	fmt.Println(GameDATA)

}
