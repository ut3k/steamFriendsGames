package controllers

import (
	"fmt"

	"steamFriendsGames/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MultiUserGameList []int64

func CheckIfGameHasManyUsers() {
	DB, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Function CheckIfGameHasManyUsers can not connect to data.db")
	}

	var games []models.Game

	DB.Raw(`
		SELECT games.id, games.name 
		FROM games
		JOIN usergames ON usergames.game_id = games.id
		GROUP BY games.id
		HAVING COUNT(usergames.user_id) > 1
	`).Scan(&games)

	for _, game := range games {
		fmt.Println(game.Title)
		MultiUserGameList = append(MultiUserGameList, int64(game.ID))

	}
	fmt.Println(MultiUserGameList)

}
