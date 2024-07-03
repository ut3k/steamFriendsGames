package controllers

import (
	"fmt"
	"steamFriendsGames/models"

	"github.com/gocolly/colly"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CheckIfGameIsCoop() {
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Function CheckIfGameIsCoop can NOT connect to data.db")
	}

	c := colly.NewCollector()

	for _, id := range MultiUserGameList {
		var game models.Game
		result := DB.First(&game, id)
		if result.Error != nil {
			fmt.Println("Game not found in data base")
		}
		fmt.Println("==================================")
		fmt.Println("Visiting : ", game.GameURL)
		c.Visit(game.GameURL)
		fmt.Println("==================================")

		// for some reason this creates duplicates maybe due to double use of the same icon for diferent types of COOP
		c.OnHTML("div.icon", func(h *colly.HTMLElement) {
			attrData := h.ChildAttr("img.category_icon", "src")
			wantedAttr := "https://store.akamai.steamstatic.com/public/images/v6/ico/ico_coop.png"
			if attrData == wantedAttr {
				fmt.Println("-----------------------")
				fmt.Println(attrData)
				fmt.Println("id GRY:", game.ID)
				fmt.Println("CODE GRY:", game.GameCODE)
				fmt.Println("State of:", game.Title, "should be changed to COOP")
				fmt.Println("-----------------------")
			}
		})
		c.Wait()

	}
	// var games []models.Game
	// err = DB.Preload("usergames").Where("user_id > ?", 1).Find(&games).Error
	//
	//	if err != nil {
	//		panic(err)
	//	} else {
	//
	//		fmt.Println(err)
	//	}
}
