package task

import (
	"Scraper/internal/scraper"
	"sync"
)

func StartTask(in []scraper.Scrape, Task func(c scraper.Scrape)) {
	var wg sync.WaitGroup
	wg.Add(len(in))

	for i := 0; i < len(in); i++ {
		c := in[i]
		go func(i int, c scraper.Scrape) {
			defer wg.Done()
			Task(c)
		}(i, c)
	}
	wg.Wait()
}
