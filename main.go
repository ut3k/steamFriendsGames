package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	c.OnHTML("._22awlPiAoaZjQMqxJhp-KP", func(e *colly.HTMLElement) {
		Title := e.Text
		fmt.Println("--------------------------------------")
		fmt.Println("Tytuł:", Title)
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
