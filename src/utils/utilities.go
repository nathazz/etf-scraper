package utils

func ScrapeUrls(isin string) string {
	return "https://www.trackingdifferences.com/ETF/ISIN/" + isin
}
