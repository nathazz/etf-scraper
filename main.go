package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"scraper-go/src/middleware"
	"scraper-go/src/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := gin.Default()

	r.Use(cors.Default())

	_ = r.SetTrustedProxies(nil)

	r.Use(middleware.RateLimitByIP())
	routes.SetupRouter(r)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Error to init server:", err)
	}
}
