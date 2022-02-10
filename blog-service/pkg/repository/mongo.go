package repository

import (
	"context"
	"github.com/Munovv/broblogo/blog-service/pkg/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewMongoDb(cfg *configs.Mongo) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		log.Fatalf("messages occurred while connection to mongo with message: %s", err.Error())
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
