package controllers

import (
	"Quotes_Server/configs"
	"Quotes_Server/models"
	"Quotes_Server/responses"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var quoteCollection *mongo.Collection = configs.GetCollection(configs.DB, "Quotes")
var validate = validator.New()

func CreateQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var quote models.Quote
		defer cancel()

		// Validate the request body
		if err := c.BindJSON(&quote); err != nil {
			c.JSON(http.StatusBadRequest, responses.QuoteResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		// Use the validator library to validate required fields
		if validationErr := validate.Struct(&quote); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.QuoteResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Data:    map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newQuote := models.Quote{
			Id:       primitive.NewObjectID(),
			Quote:    quote.Quote,
			Author:   quote.Author,
			Category: quote.Category,
			Tags:     quote.Tags,
			Likes:    0,
		}

		// Check if the quote already exists
		if isQuoteDuplicated(newQuote) {
			c.JSON(http.StatusBadRequest, responses.QuoteResponse{
				Status:  http.StatusBadRequest,
				Message: "Quote already exists",
				Data:    map[string]interface{}{"data": newQuote}})
			return
		}

		// Insert the quote into the database
		res, err := quoteCollection.InsertOne(ctx, newQuote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.QuoteResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		// Return the quote
		c.JSON(http.StatusCreated, responses.QuoteResponse{
			Status:  http.StatusCreated,
			Message: "Quote created successfully",
			Data:    map[string]interface{}{"data": res}})
	}
}

func isQuoteDuplicated(quote models.Quote) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the quote already exists
	res := quoteCollection.FindOne(ctx, bson.M{"quote": quote.Quote})
	if res.Err() != nil {
		return false
	}
	return true
}
