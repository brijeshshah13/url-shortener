package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brijeshshah13/url-shortener/internal/proto/shortener"
	"github.com/brijeshshah13/url-shortener/services/frontend/controller"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// NewFrontend returns a new server
func NewFrontend(t opentracing.Tracer, shortenerconn *grpc.ClientConn) *Frontend {
	return &Frontend{
		shortenerClient: shortener.NewShortenerClient(shortenerconn),
		tracer:          t,
	}
}

// Frontend implements frontend service
type Frontend struct {
	shortenerClient shortener.ShortenerClient
	tracer          opentracing.Tracer
}

// Run the server
func (s *Frontend) Run(port int) error {

	router := gin.Default()

	router.Use(ginhttp.Middleware(s.tracer))

	// status check
	router.GET("/status", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ok")
	})

	urlGroup := router.Group("api/v1/url-shortener")
	{
		urlGroup.POST("/create", controller.Create)
	}

	// start the server on port 9090
	err := router.Run(fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	return err

}

// func (s *Frontend) statusHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, "Please specify inDate/outDate params", http.StatusBadRequest)
// 	return
// }
