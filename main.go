package main

import (
	"Quotes_Server/configs"
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

	router.Run("localhost:8080")
}
