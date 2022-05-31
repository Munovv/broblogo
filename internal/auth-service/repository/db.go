package repository

import (
	"context"
	mongoCfg "github.com/Munovv/broblogo/auth-service/auth-service/config/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InitDb(cfg *mongoCfg.MongoConfig) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		log.Fatalf("an error occurred while init connection to mongoDB: %s", err.Error())

		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client.Database(cfg.Name), nil
}
