package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scraper-go/src/scraper"
	"scraper-go/src/utils"
	"time"
)

type EtfRequest struct {
	Isins []string `json:"isins"`
}

func GetEtf(c *gin.Context) {
	isin := c.Param("isin")

	if isin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ISIN required"})
		return
	}

	results := scraper.EtfScraper([]string{isin})
	c.JSON(http.StatusOK, results)
}

func GetMoreEtfs(c *gin.Context) {
	var req EtfRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if len(req.Isins) > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can only request up to 10 ISINs"})
		return
	}

	results := scraper.EtfScraper(req.Isins)
	c.JSON(http.StatusOK, results)
}

func GenerateEtfsPdfs(c *gin.Context) {
	var req EtfRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

	if len(req.Isins) > 10 {
		c.JSON(http.StatusOK, gin.H{"error": "You can only request up to 10 ISINs"})
		return
	}

	etfs := scraper.EtfScraper(req.Isins)
	pdfBytes, err := utils.SaveToPDF(etfs)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error to create PDF")
		return
	}

	filename := fmt.Sprintf("etf-report-%s.pdf", time.Now().Format("20060102"))
	c.Header("Content-Disposition", "attachment; filename="+filename)

	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
