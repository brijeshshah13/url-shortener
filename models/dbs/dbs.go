package dbs

import (
	"context"
	"log"
	"time"

	"github.com/brijeshshah13/url-shortener/config/environments"
	"github.com/brijeshshah13/url-shortener/models/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBNames = map[string]string{
		"main": "urlshortener",
	}
)

var allowedDBNames = make(map[string]struct{})

var collectionRegistry = map[string]string{
	utils.CollectionNames["url"]: DBNames["main"],
}

type dbRegistryConfig struct {
	connected bool
	dbConfig  environments.DBConfig
}

var dbRegistry = map[string]dbRegistryConfig{
	DBNames["main"]: {
		connected: false,
		dbConfig:  environments.Mongo["main"],
	},
}

func init() {
	// initialize list of allowed db names
	for _, v := range DBNames {
		allowedDBNames[v] = struct{}{}
	}
}

func ConnectDB(dbName string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbRegistry[dbName].dbConfig.URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	conn := client.Database(dbName)
	return conn, nil
}
