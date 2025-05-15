package cache

import (
	"github.com/patrickmn/go-cache"
	"scraper-go/src/model"
	"time"
)

var EtfCache = cache.New(5*time.Minute, 10*time.Minute)

func GetEtfsWithCache(isins []string, scraperFunc func([]string) []model.EtfInfo) []model.EtfInfo {
	var cachedResults []model.EtfInfo
	var toScrape []string

	for _, isin := range isins {
		if data, found := EtfCache.Get(isin); found {
			cachedResults = append(cachedResults, data.([]model.EtfInfo)...)
		} else {
			toScrape = append(toScrape, isin)
		}
	}

	var scrapedResults []model.EtfInfo

	if len(toScrape) > 0 {
		scrapedResults = scraperFunc(toScrape)

		for _, etf := range scrapedResults {
			EtfCache.Set(etf.Isin, []model.EtfInfo{etf}, cache.DefaultExpiration)
		}
	}

	return append(cachedResults, scrapedResults...)
}
