package scraper

import (
	"fmt"
	"github.com/Danny-Dasilva/fhttp"
	"io"
	"log"
)

func (_ *Scrape) GetHTMLSpareRibExpress(s Scrape, url string) string {
	req, err := http.NewRequest("GET", fmt.Sprintf(
		"https://%s", url),
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	Hd.DefaultHeader(req, map[string]string{
		"Cookie":  "sre_session=" + s.Cookie,
		"Referer": "https://www.spareribexpress.com/",
		"Accept":  "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		//"Accept-Encoding": "gzip, deflate, br",
	})

	resp, err := s.Client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 200, 201, 204:
		return string(body)
	default:
		log.Println("Failed", string(body))
	}
	return ""
}
