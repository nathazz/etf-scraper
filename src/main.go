package main

import (
	"github.com/gin-gonic/gin"
	"scraper-go/src/routes"
)

func main() {

	r := gin.Default()
	err := r.Run(":8080")

	routes.SetupRouter(r)

	if err != nil {
		return
	}
}
