package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scraper-go/src/model"
	"scraper-go/src/scraper"
	"scraper-go/src/utils"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetEtf(c *gin.Context) {
	isin := c.Param("isin")

	if !utils.ValidateIsins(c, []string{isin}) {
		return
	}

	results := scraper.EtfScraper([]string{isin})
	c.JSON(http.StatusOK, results)
}

func GetMoreEtfs(c *gin.Context) {
	var req model.EtfRequest

	if err := c.BindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input. Expected JSON with 'isins' array")
		return
	}

	if !utils.ValidateIsins(c, req.Isins) {
		return
	}

	results := scraper.EtfScraper(req.Isins)
	c.JSON(http.StatusOK, results)
}

func CompareEtf(c *gin.Context) {
	var req model.EtfRequest

	if err := c.BindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input. Expected JSON with 'isins' array")
		return
	}

	if len(req.Isins) < 2 {
		utils.RespondError(c, http.StatusBadRequest, "For comparison it is necessary to have at least two 'isins'")
		return
	}

	if !utils.ValidateIsins(c, req.Isins) {
		return
	}

	etfs := scraper.EtfScraper(req.Isins)

	results := utils.CompareEtf(etfs, true)
	c.JSON(http.StatusOK, results)

}
