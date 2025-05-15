package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"scraper-go/src/model"
)

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

func ValidateEtfInfos(c *gin.Context, etfs []model.EtfInfo) bool {
	for _, etf := range etfs {
		if isInvalidEtfInfo(&etf) {
			RespondError(c, http.StatusNotFound, "One or more ISINs returned invalid or empty data. Make sure your exemptions are correct")
			return true
		}
	}
	return false
}

func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func isValidISIN(isin string) bool {
	matched, _ := regexp.MatchString(`^[A-Z]{2}[A-Z0-9]{10}$`, isin)
	return matched
}

func isInvalidEtfInfo(etf *model.EtfInfo) bool {
	return etf.Title == "" ||
		etf.Description == "" ||
		etf.Replication == "" ||
		etf.TotalExpenseRatio == "" ||
		etf.TrackingDifference == "" ||
		etf.Earnings == "" ||
		etf.FundSize == ""
}
