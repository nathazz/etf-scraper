package utils

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

func ScrapeUrls(isin string) string {
	return "https://www.trackingdifferences.com/ETF/ISIN/" + isin
}

func IsValidISIN(isin string) bool {
	matched, _ := regexp.MatchString(`^[A-Z]{2}[A-Z0-9]{10}$`, isin)
	return matched
}

func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
