package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/brijeshshah13/url-shortener/internal/dialer"
	tracer "github.com/brijeshshah13/url-shortener/internal/trace"
	"github.com/brijeshshah13/url-shortener/models/dbs"
	frontend "github.com/brijeshshah13/url-shortener/services/frontend"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type server interface {
	Run(int) error
}

func main() {
	var (
		port          = flag.Int("port", 8080, "The service port")
		jaegeraddr    = flag.String("jaeger", "jaeger:6831", "Jaeger address")
		shorteneraddr = flag.String("shorteneraddr", "shortener:8080", "Shortener service addr")
	)
	flag.Parse()

	flush := tracer.InitTracer(os.Args[1], *jaegeraddr)
	defer flush()

	tracer := otel.Tracer(os.Args[1])

	var srv server

	switch os.Args[1] {
	case "shortener":
	case "frontend":
		srv = frontend.NewFrontend(
			tracer,
			initGRPCConn(*shorteneraddr, tracer),
			initDBConn(dbs.DBNames["main"]),
		)
	default:
		log.Fatalf("unknown command %s", os.Args[1])
	}

	srv.Run(*port)
}

func initGRPCConn(addr string, tracer trace.Tracer) *grpc.ClientConn {
	conn, err := dialer.Dial(addr, dialer.WithTracer(tracer))
	if err != nil {
		panic(fmt.Sprintf("ERROR: dial error: %v", err))
	}
	return conn
}

func initDBConn(dbName string) *mongo.Database {
	conn, err := dbs.ConnectDB(dbName)
	if err != nil {
		panic(fmt.Sprintf("ERROR: db conn error: %v", err))
	}
	return conn
}
