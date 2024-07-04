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
		fmt.Println(id)
		c.Visit(game.GameURL)
		fmt.Println("==================================")

		// for some reason this creates duplicates maybe due to double use of the same icon for diferent types of COOP
		c.OnHTML("div.icon", func(h *colly.HTMLElement) {
			attrData := h.ChildAttr("img.category_icon", "src")
			wantedAttr := "https://store.akamai.steamstatic.com/public/images/v6/ico/ico_coop.png"
			if attrData == wantedAttr && int64(game.ID) == id {
				fmt.Println("-----------------------")
				fmt.Println("Game:", game.Title, "- COOP")
				fmt.Println("game id:", game.ID, "id ze zmiennej:", id)
				fmt.Println("-----------------------")
				game.IsCooperative = true
				c.OnHTML("img.game_header_image_full", func(iurl *colly.HTMLElement) {
					game.MainIMG = iurl.Attr("src")
				})

				result.Save(&game.MainIMG)
				result.Save(&game.IsCooperative)

			}
		})

		c.Wait()
	}
}
