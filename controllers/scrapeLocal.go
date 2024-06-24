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
)

func ScrapeLocalData() {
	// DateBase setup
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
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

		fmt.Println("=================")
		fmt.Println("Gracz:", UserName)
		fmt.Println("=================")

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
		fmt.Println("--------------------------------------")
		fmt.Println("Tytuł:", Title)
		fmt.Println("ID gry:", GameID)
		fmt.Println("GameURL:", GameURL)
		fmt.Println("USER:", ActualUser)
		fmt.Println("--------------------------------------")

		var user models.User
		DB.Where("name = ?", ActualUser).First(&user)

		game := models.Game{
			Title:   Title,
			GameID:  GameID,
			GameURL: GameURL,
		}

		err := DB.Create(&game).Error
		if err != nil {
			println("not able to write User name to DataBase")
		}

		if user.ID == 0 {
			fmt.Println(" there is no such user as:", ActualUser, "in data base")
		}

	})

	FileList, err := os.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range FileList {
		// fmt.Println(file.Name())

		c.Visit("file://" + dir + "/data/" + file.Name())
		fmt.Println("file://" + dir + "/data/" + file.Name())
	}

}
