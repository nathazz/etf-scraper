package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scraper-go/src/model"
	"scraper-go/src/services"
	"scraper-go/src/utils"
)

var etfService = services.NewEtfService()

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetEtf(c *gin.Context) {
	isin := c.Param("isin")

	etfService.GetEtfByIsin(c, isin)
}

func GetMoreEtfs(c *gin.Context) {
	var req model.EtfRequest

	if err := c.BindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input. Expected JSON with 'isins' array")
		return
	}

	etfService.GetMultipleEtfs(c, req)
}

func CompareEtf(c *gin.Context) {
	var req model.EtfRequest

	if err := c.BindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input. Expected JSON with 'isins' array")
		return
	}

	etfService.ComparingEtfs(c, req)
}
