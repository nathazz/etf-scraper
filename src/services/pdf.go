package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"scraper-go/src/model"
	"scraper-go/src/scraper"
	"scraper-go/src/utils"
	"time"
)

type PdfService struct{}

func NewpPdfService() *PdfService {
	return &PdfService{}
}

func (s *PdfService) CreatePdf(c *gin.Context, req model.EtfRequest) {
	if !utils.ValidateIsins(c, req.Isins) {
		return
	}

	etfs := scraper.EtfScraper(req.Isins)

	if utils.ValidateEtfInfos(c, etfs) {
		return
	}

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
