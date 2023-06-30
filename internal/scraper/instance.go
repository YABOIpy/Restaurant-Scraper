package scraper

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	HttpClient "Scraper/internal/client"
	"Scraper/internal/utils"
	"github.com/Danny-Dasilva/fhttp"
)

var (
	util utils.Util
)

func (s *Scrape) Configuration() ([]Scrape, error) {
	Conf, err := util.LoadConfig("config.json")
	if err != nil {
		log.Println("Failed To Load Config")
		time.Sleep(3 * time.Second)
		return nil, err
	}

	var (
		ClientErr error
		Instances []Scrape
		Client    *http.Client
		proxy     string
		cfg       = Conf
	)

	sites, _ := util.ReadFile("sites.txt")

	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(len(sites))

	fmt.Print("\033[s")

	for i := 0; i < len(sites); i++ {
		go func(i int) {
			defer wg.Done()

			url := strings.Split(sites[i], "https://")[1]
			cookie := util.Cookie(url)

			fmt.Print("\033[u\033[K")
			fmt.Printf(utils.ChaceLoading, cookie[:20], url)

			if strings.Count(cfg.Mode.Network.Proxy, "") > 2 {
				proxy = "http://" + cfg.Mode.Network.Proxy
			}
			Client, ClientErr = HttpClient.NewClient(HttpClient.Browser{
				JA3:       cfg.Mode.Network.Ja3,
				UserAgent: cfg.Mode.Network.Agent,
				Cookies:   nil,
			},
				cfg.Mode.Network.TimeOut,
				cfg.Mode.Network.Redirect,
				cfg.Mode.Network.Agent,
				proxy,
			)
			if ClientErr != nil {
				mutex.Lock()
				err = ClientErr
				mutex.Unlock()
				return
			}

			mutex.Lock()
			Instances = append(Instances, Scrape{
				Client:  Client,
				SiteUrl: url,
				Cookie:  cookie,
			})
			mutex.Unlock()

		}(i)
	}
	wg.Wait()

	return Instances, err
}
