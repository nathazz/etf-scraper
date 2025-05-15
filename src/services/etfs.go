package services

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	cache2 "scraper-go/src/cache"
	"scraper-go/src/model"
	"scraper-go/src/scraper"
	"scraper-go/src/utils"
)

type EtfService struct{}

func NewEtfService() *EtfService {
	return &EtfService{}
}

func (s *EtfService) GetEtfByIsin(c *gin.Context, isin string) {
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

func (s *EtfService) GetMultipleEtfs(c *gin.Context, req model.EtfRequest) {

	if len(req.Isins) < 2 {
		utils.RespondError(c, http.StatusBadRequest, "The route is for multiple isins, not for one")
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

func (s *EtfService) ComparingEtfs(c *gin.Context, req model.EtfRequest) {
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
