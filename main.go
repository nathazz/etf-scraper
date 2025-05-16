package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"scraper-go/src/middleware"
	"scraper-go/src/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r := gin.Default()
	r.Use(middleware.RateLimitByIP())

	routes.SetupRouter(r)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Error to init server:", err)
	}
}
