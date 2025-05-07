package main

import (
	"github.com/gin-gonic/gin"
	"scraper-go/src/routes"
)

func main() {

	r := gin.Default()
	routes.SetupRouter(r)

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
