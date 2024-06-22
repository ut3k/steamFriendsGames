package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(t)

	// pages := []string{}

	c.Visit("file://" + dir + "/data/view-source_https___steamcommunity.com_id_aquiceeo_games__tab=all.html")
	println("file://" + dir + "/data/view-source_https___steamcommunity.com_id_aquiceeo_games__tab=all.html")
	c.OnHTML("a.persona_name_text_content", func(h *colly.HTMLElement) {
		var UserName string
		UserName = h.Text
		UserName = strings.TrimSpace(UserName)

		fmt.Println("=================")
		fmt.Println("Gracz:", UserName)
		fmt.Println("=================")

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
		fmt.Println("--------------------------------------")
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
