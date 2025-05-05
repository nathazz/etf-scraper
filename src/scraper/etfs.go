package scraper

import (
	"github.com/gocolly/colly"
	"scraper-go/src/model"
	"scraper-go/src/utils"
	"strings"
)

func EtfScraper() []model.EtfInfo {
	isins := []string{"LU1737652310", "LU0274210672", `IE00B14X4M10`}
	etfInfo := model.EtfInfo{}
	etfInfos := make([]model.EtfInfo, 0, len(isins))

	c := colly.NewCollector(colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"))

	c.OnHTML("h1.page-title", func(e *colly.HTMLElement) {
		etfInfo.Title = e.Text
	})

	c.OnHTML("p.mt-2", func(e *colly.HTMLElement) {
		etfInfo.Description = e.Text
	})

	c.OnHTML("div.descfloat p.desc", func(e *colly.HTMLElement) {
		selection := e.DOM
		nodes := selection.Children().Nodes

		if len(nodes) == 3 {
			description := strings.TrimSpace(selection.Find("span.desctitle").Text())
			value := selection.FindNodes(nodes[2]).Text()

			switch description {
			case "Replication":
				etfInfo.Replication = value
			case "TER":
				etfInfo.TotalExpenseRatio = value
			case "TD":
				etfInfo.TrackingDifference = value
			case "Earnings":
				etfInfo.Earnings = value
			case "Fund size":
				etfInfo.FundSize = value
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		etfInfos = append(etfInfos, etfInfo)

		//err := utils.SaveToPDF(etfInfos, "etfinfo.pdf")

		//if err != nil {
		//	return
		//}

		etfInfo = model.EtfInfo{}
	})

	for _, isin := range isins {
		_ = c.Visit(utils.ScrapeUrls(isin))
	}

	return etfInfos
}
