package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scraper-go/src/scraper"
	"scraper-go/src/utils"
	"time"
)

func GetEtfs(c *gin.Context) {
	etfs := scraper.EtfScraper()

	c.JSON(http.StatusOK, etfs)
}

func GetEtfsPdf(c *gin.Context) {
	etfs := scraper.EtfScraper()
	pdfBytes, err := utils.SaveToPDF(etfs)

	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao gerar PDF")
		return
	}

	filename := fmt.Sprintf("etf-report-%s.pdf", time.Now().Format("20060102"))
	c.Header("Content-Disposition", "attachment; filename="+filename)

	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
