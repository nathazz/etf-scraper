package routes

import (
	"github.com/gin-gonic/gin"
	"scraper-go/src/controllers"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/etfs", controllers.GetEtfs)
	r.GET("/etfs/:id")
	r.GET("/etf/generate-pdf", controllers.GetEtfsPdf)
}
