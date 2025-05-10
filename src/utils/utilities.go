package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func ScrapeUrls(isin string) string {
	return "https://www.trackingdifferences.com/ETF/ISIN/" + isin
}

func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func ValidateIsins(c *gin.Context, isins []string) bool {
	if len(isins) == 0 {
		RespondError(c, http.StatusBadRequest, "At least one ISIN is required")
		return false
	}

	if len(isins) > 10 {
		RespondError(c, http.StatusBadRequest, "You can only request up to 10 ISINs")
		return false
	}

	for _, isin := range isins {
		if !isValidISIN(isin) {
			RespondError(c, http.StatusBadRequest, fmt.Sprintf("Invalid ISIN format: %s", isin))
			return false
		}
	}

	return true
}

func isValidISIN(isin string) bool {
	matched, _ := regexp.MatchString(`^[A-Z]{2}[A-Z0-9]{10}$`, isin)
	return matched
}
