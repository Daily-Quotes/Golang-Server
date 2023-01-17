package routes

import (
	"Quotes_Server/controllers"
	"github.com/gin-gonic/gin"
)

func QuoteRoute(router *gin.Engine) {
	// ---- GET

	// Get all Quotes - admin
	// Get all quotes that are not approved - admin

	// Get quotes stored in used (LIMIT 20) - app

	// ---- POST

	// Add Quote - Admin
	router.POST("/quote", controllers.CreateQuote())

	// Create Quote request - Web

	// ---- UPDATE

	// Update Quote info - Admin

	// Like a Quote - App

	// ---- DELETE

	// Delete Quote - Admin
}
