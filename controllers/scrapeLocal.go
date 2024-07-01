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
		println(ActualUser)

	})

	c.OnHTML("._22awlPiAoaZjQMqxJhp-KP", func(e *colly.HTMLElement) {
		Title := e.Text
		GameURL := e.Attr("href")
		GameCODE := strings.ReplaceAll(GameURL, "https://store.steampowered.com/app/", "")
		GameCODE = strings.TrimSpace(GameCODE)

		var user models.User
		DB.Where("name = ?", ActualUser).First(&user)

		game := models.Game{
			Title:    Title,
			GameCODE: GameCODE,
			GameURL:  GameURL,
		}

		err := DB.Create(&game).Error
		if err == nil {
			DB.Model(&user).Association("Games").Append(&game)
			fmt.Println(game.Title, "- save in database")
			fmt.Println(ActualUser)
		} else {
			var gameScaned models.Game
			DB.Where("game_code = ?", GameCODE).First(&gameScaned)
			DB.Model(&user).Association("Games").Append(&gameScaned)
			fmt.Println("===========")
			fmt.Println("game scanned:", gameScaned)
			fmt.Println("Actual user:", ActualUser)
			fmt.Println("GameCODE user:", GameCODE)
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
