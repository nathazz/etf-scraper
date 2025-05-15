package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	cache2 "scraper-go/src/cache"
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

	if data, found := cache2.EtfCache.Get(isin); found {
		c.JSON(http.StatusOK, data)
		return
	}

	results := scraper.EtfScraper([]string{isin})

	if utils.ValidateEtfInfos(c, results) {
		return
	}

	cache2.EtfCache.Set(isin, results, cache.DefaultExpiration)

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

	allResults := cache2.GetEtfsWithCache(req.Isins, scraper.EtfScraper)

	if utils.ValidateEtfInfos(c, allResults) {
		return
	}

	c.JSON(http.StatusOK, allResults)
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

	allResults := cache2.GetEtfsWithCache(req.Isins, scraper.EtfScraper)

	if utils.ValidateEtfInfos(c, allResults) {
		return
	}

	results := utils.CompareEtf(allResults, true)
	c.JSON(http.StatusOK, results)
}
