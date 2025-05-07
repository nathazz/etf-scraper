package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"scraper-go/src/scraper"
	"scraper-go/src/utils"
	"time"
)

type EtfRequest struct {
	Isins []string `json:"isins"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetEtf(c *gin.Context) {
	isin := c.Param("isin")

	if isin == "" {
		utils.RespondError(c, http.StatusBadRequest, "ISIN is required")
		return
	}

	if !utils.IsValidISIN(isin) {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ISIN format")
		return
	}

	results := scraper.EtfScraper([]string{isin})
	c.JSON(http.StatusOK, results)
}

func GetMoreEtfs(c *gin.Context) {
	var req EtfRequest

	if err := c.BindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input. Expected JSON with 'isins' array")
		return
	}

	if len(req.Isins) == 0 {
		utils.RespondError(c, http.StatusBadRequest, "At least one ISIN is required")
		return
	}

	if len(req.Isins) > 10 {
		utils.RespondError(c, http.StatusBadRequest, "You can only request up to 10 ISINs")
		return
	}

	for _, isin := range req.Isins {
		if !utils.IsValidISIN(isin) {
			utils.RespondError(c, http.StatusBadRequest, fmt.Sprintf("Invalid ISIN format: %s", isin))
			return
		}
	}

	results := scraper.EtfScraper(req.Isins)
	c.JSON(http.StatusOK, results)
}

func GenerateEtfsPdfs(c *gin.Context) {
	var req EtfRequest

	if err := c.BindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input. Expected JSON with 'isins' array")
		return
	}

	if len(req.Isins) == 0 {
		utils.RespondError(c, http.StatusBadRequest, "At least one ISIN is required")
		return
	}

	if len(req.Isins) > 10 {
		utils.RespondError(c, http.StatusBadRequest, "You can only request up to 10 ISINs")
		return
	}

	for _, isin := range req.Isins {
		if !utils.IsValidISIN(isin) {
			utils.RespondError(c, http.StatusBadRequest, fmt.Sprintf("Invalid ISIN format: %s", isin))
			return
		}
	}

	etfs := scraper.EtfScraper(req.Isins)
	pdfBytes, err := utils.SaveToPDF(etfs)

	if err != nil {
		log.Printf("Error generating PDF: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Error generating PDF")
		return
	}

	filename := fmt.Sprintf("etf-report-%s.pdf", time.Now().Format("20060102"))
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
