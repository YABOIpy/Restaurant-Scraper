package scraper

import http "github.com/Danny-Dasilva/fhttp"

type Scrape struct {
	Client  *http.Client
	SiteUrl string
	Cookie  string
}

type Header struct {
}

var (
	Hd = Header{}
)
