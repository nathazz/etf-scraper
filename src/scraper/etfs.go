package scraper

import (
	"github.com/gocolly/colly"
	"log"
	"scraper-go/src/model"
	"scraper-go/src/utils"
	"strings"
)

func EtfScraper(isins []string) []model.EtfInfo {
	etfInfos := make([]model.EtfInfo, 0, len(isins))

	c := colly.NewCollector(
		colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"),
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9,en;q=0.8")
		log.Println("Visiting:", r.URL.String())

		if r.Ctx.GetAny("etfInfo") == nil {
			r.Abort()
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("h1.page-title", func(e *colly.HTMLElement) {
		info := e.Request.Ctx.GetAny("etfInfo").(*model.EtfInfo)
		info.Title = e.Text
	})

	c.OnHTML("p.mt-2", func(e *colly.HTMLElement) {
		info := e.Request.Ctx.GetAny("etfInfo").(*model.EtfInfo)
		info.Description = e.Text
	})

	c.OnHTML("div.descfloat p.desc", func(e *colly.HTMLElement) {
		info := e.Request.Ctx.GetAny("etfInfo").(*model.EtfInfo)
		selection := e.DOM
		nodes := selection.Children().Nodes

		if len(nodes) == 3 {
			description := strings.TrimSpace(selection.Find("span.desctitle").Text())
			value := selection.FindNodes(nodes[2]).Text()

			switch description {
			case "Replication":
				info.Replication = value
			case "TER":
				info.TotalExpenseRatio = value
			case "TD":
				info.TrackingDifference = value
			case "Earnings":
				info.Earnings = value
			case "Fund size":
				info.FundSize = value
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		info := r.Ctx.GetAny("etfInfo").(*model.EtfInfo)
		etfInfos = append(etfInfos, *info)
	})

	for _, isin := range isins {
		info := &model.EtfInfo{
			Isin: isin,
		}

		ctx := colly.NewContext()
		ctx.Put("etfInfo", info)

		err := c.Request("GET", utils.ScrapeUrls(isin), nil, ctx, nil)
		if err != nil {
			log.Println("Failed to visit:", isin, err)
		}
	}

	c.Wait()

	return etfInfos
}
