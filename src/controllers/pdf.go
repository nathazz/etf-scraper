package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scraper-go/src/model"
	"scraper-go/src/services"
	"scraper-go/src/utils"
)

var pdfService = services.NewpPdfService()

func EtfPdfGenerator(c *gin.Context) {
	var req model.EtfRequest

	if err := c.BindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input. Expected JSON with 'isins' array")
		return
	}

	pdfService.CreatePdf(c, req)
}
