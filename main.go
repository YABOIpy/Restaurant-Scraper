package main

import (
	"Scraper/internal/scraper"
	"Scraper/internal/scraper/parsers"
	"Scraper/internal/task"
	"Scraper/internal/utils"
	"fmt"
	"log"
	"time"
)

var (
	s = scraper.Scrape{}
)

func initialize() []scraper.Scrape {
	fmt.Println(utils.Initializing)
	data, err := s.Configuration()
	if err != nil {
		log.Println(err)
		time.Sleep(3 * time.Second)
		return nil
	}
	return data
}

func GetSpareRibExpress(c scraper.Scrape) {
	Webdata := s.GetHTMLSpareRibExpress(c, c.SiteUrl)
	data := parsers.SpareRibExpress(Webdata)
	for _, d := range data {
		fmt.Printf("Titel: %s\n Subtitel: %s | Prijs: %s", d.Title, d.Subtitle, d.Price)
		//kan dit schrijven naar een csv file of database
	}

}

func main() {
	in := initialize()
	task.StartTask(in, func(c scraper.Scrape) {
		GetSpareRibExpress(c)
	})
}
