package controller

import (
	"net/http"

	"github.com/brijeshshah13/url-shortener/services/frontend/model"
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

	// TODO: to uncomment below code once ready for use
	// urlOptions := struct {
	// 	is_active bool
	// 	is_used   bool
	// }{
	// 	is_active: true,
	// 	is_used:   false,
	// }

	// urlData := struct {
	// 	original_url    string
	// 	expiration_date int64
	// 	is_used         bool
	// }{
	// 	original_url:    url.OriginalURL,
	// 	expiration_date: time.Now().UTC().UnixNano() / 1e6,
	// 	is_used:         true,
	// }

	ctx.String(http.StatusOK, url.OriginalURL)
}
