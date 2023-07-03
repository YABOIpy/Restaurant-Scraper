package parsers

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Express struct {
	Title    string
	Subtitle string
	Price    string
}

func SpareRibExpress(content string) (menus []Express) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".menu-dish--large").Each(func(i int, s *goquery.Selection) {
		data := Express{
			Title:    s.Find("h3.menu-dish--large__title").Text(),
			Subtitle: s.Find("p.menu-dish--large__sub").Text(),
			Price:    s.Find("p.menu-dish--large__price").Text(),
		}
		menus = append(menus, data)
	})

	return menus
}
