package routes

import (
	"github.com/gin-gonic/gin"
	"scraper-go/src/controllers"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", controllers.Health)
	r.GET("/etf/:isin", controllers.GetEtf)
	r.POST("/etf/compare-etfs", controllers.CompareEtf)
	r.POST("/etf", controllers.GetMoreEtfs)
	r.POST("/etf/generate-pdf", controllers.EtfPdfGenerator)
}
