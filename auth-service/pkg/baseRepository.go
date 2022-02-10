package pkg

import (
	"context"
	mongoConfig "github.com/Munovv/broblogo/auth-service/config/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InitDb(cfg *mongoConfig.MongoConfig) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		log.Fatalf("an error occurred while init connection to mongoDB: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(cfg.Name)
}
