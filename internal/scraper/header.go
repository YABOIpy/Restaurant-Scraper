package scraper

import (
	"github.com/Danny-Dasilva/fhttp"
)

func (Hd *Header) DefaultHeader(req *http.Request, header map[string]string) {
	for x, o := range map[string]string{
		"Accept-Language":           "en-US,en;q=0.9,nl;q=0.8",
		"Sec-Ch-Ua":                 `Not.A/Brand";v="8", "Chromium";v="114", "Vivaldi";v="6.1"`,
		"Sec-Ch-Ua-Mobile":          "?0",
		"Sec-Ch-Ua-Platform":        `"Windows"`,
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	} {
		req.Header.Set(x, o)
	}
	for x, o := range header {
		req.Header.Set(x, o)
	}
}
