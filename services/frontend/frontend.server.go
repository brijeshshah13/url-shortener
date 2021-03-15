package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brijeshshah13/url-shortener/internal/proto/shortener"
	"github.com/brijeshshah13/url-shortener/services/frontend/controller"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

// NewFrontend returns a new server
func NewFrontend(t trace.Tracer, shortenerconn *grpc.ClientConn) *Frontend {
	return &Frontend{
		shortenerClient: shortener.NewShortenerClient(shortenerconn),
		tracer:          t,
	}
}

// Frontend implements frontend service
type Frontend struct {
	shortenerClient shortener.ShortenerClient
	tracer          trace.Tracer
}

// Run the server
func (s *Frontend) Run(port int) error {

	router := gin.Default()

	router.Use(otelgin.Middleware("frontend"))

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
