package main

import (
	"log"
	"net/http"

	"github.com/brijeshshah13/url-shortener/client/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	// creating a connection to the gRPC server
	// conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	// check if any error occured while connecting to the gRPC server
	// if err != nil {
	// 	panic(err)
	// }

	// client := proto.NewUrlShortenerClient(conn)

	g := gin.Default()

	urlGroup := g.Group("api/v1/url-shortener")
	{
		urlGroup.POST("/create", controller.Create)
	}

	// status check
	g.GET("/status", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ok")
	})

	// start the server on port 9090
	if err := g.Run(":9090"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
