package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	collector := colly.NewCollector()

	collector.OnHTML("file", func(e *colly.HTMLElement) {
		Title := e.DOM.HasClass("_22awlPiAoaZjQMqxJhp-KP")
		fmt.Println("--------------------------------------")
		fmt.Println("Tytuł:", Title)
		fmt.Println("--------------------------------------")
	})

	files, err := os.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := "./data" + file.Name()
		err := collector.OnFile(filePath, e.Find)
	}

}
