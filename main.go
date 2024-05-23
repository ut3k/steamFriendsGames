package main

import (
	"fmt"
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
	c.Wait()

	c.OnHTML("div", func(e *colly.HTMLElement) {
		Title := e.Text
		fmt.Println("--------------------------------------")
		fmt.Println("Tytuł:", Title)
		fmt.Println("--------------------------------------")
	})

}
