package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"steamFriendsGames/models"

	"github.com/gocolly/colly"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ScrapeLocalData() {
	// DateBase setup
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("scraper faild to connect to data.db")
	}
	// location of Steam Downloaded DATA
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(t)

	var ActualUser string

	c.OnHTML("a.persona_name_text_content", func(h *colly.HTMLElement) {
		var UserName string
		UserName = h.Text
		UserName = strings.TrimSpace(UserName)

		fmt.Println("Checkin local data file of:", UserName)

		user := models.User{
			Name: UserName,
		}
		err := DB.Create(&user).Error
		if err != nil {
			println("not able to write User name to DataBase")
		}

		ActualUser = UserName

	})

	c.OnHTML("._22awlPiAoaZjQMqxJhp-KP", func(e *colly.HTMLElement) {
		Title := e.Text
		GameURL := e.Attr("href")
		GameID := strings.ReplaceAll(GameURL, "https://store.steampowered.com/app/", "")
		GameID = strings.TrimSpace(GameID)

		var user models.User
		DB.Where("name = ?", ActualUser).First(&user)

		game := models.Game{
			Title:   Title,
			GameID:  GameID,
			GameURL: GameURL,
		}

		err := DB.Create(&game).Error
		if err == nil {
			DB.Model(&user).Association("Games").Append(&game)
			fmt.Println(game.Title, "- save in database")
		} else {
			var actualGame models.Game
			gameScaned := DB.Where("game_id = ?", GameID).First(&actualGame)
			DB.Model(&user).Association("Games").Append(&gameScaned)
			fmt.Println(game.Title, "- Game already in database")

		}

	})

	FileList, err := os.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range FileList {
		// fmt.Println(file.Name())

		c.Visit("file://" + dir + "/data/" + file.Name())
	}

}
