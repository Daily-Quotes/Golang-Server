package main

import (
	"Quotes_Server/configs"
	"Quotes_Server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Router
	router := gin.Default()

	// Run DB
	configs.ConnectDB()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.QuoteRoute(router)

	router.Run("localhost:8080")
}
