package repository

import (
	"context"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewMongoDb(cfg *config.Mongo) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		log.Fatalf("an error occurred while init connect to mongo: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	return client.Database(cfg.Name)
}
