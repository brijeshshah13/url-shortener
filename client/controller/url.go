package controller

import (
	"net/http"

	"github.com/brijeshshah13/url-shortener/client/model"
	"github.com/gin-gonic/gin"
)

/**
 * Contains controller for gRPC client which is facing the FE
 */

// Create encoded short URL
func Create(ctx *gin.Context) {
	// fetch URL model object to be expected in the request body
	url := model.URL{}
	// bind request body to the model
	err := ctx.BindJSON(&url)
	// check if err is not nil
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}
	ctx.String(http.StatusOK, url.OriginalURL)
}
