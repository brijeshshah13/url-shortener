package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	services "github.com/brijeshshah13/url-shortener"
	"github.com/brijeshshah13/url-shortener/internal/dialer"
	"github.com/brijeshshah13/url-shortener/internal/trace"
	"github.com/opentracing/opentracing-go"
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

	tracer, err := trace.New("search", *jaegeraddr)
	if err != nil {
		log.Fatalf("trace new error: %v", err)
	}

	var srv server

	switch os.Args[1] {
	case "shortener":
	case "frontend":
		srv = services.NewFrontend(
			tracer,
			initGRPCConn(*shorteneraddr, tracer),
		)
	default:
		log.Fatalf("unknown command %s", os.Args[1])
	}

	srv.Run(*port)
}

func initGRPCConn(addr string, tracer opentracing.Tracer) *grpc.ClientConn {
	conn, err := dialer.Dial(addr, dialer.WithTracer(tracer))
	if err != nil {
		panic(fmt.Sprintf("ERROR: dial error: %v", err))
	}
	return conn
}
