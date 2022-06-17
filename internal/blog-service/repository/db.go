package repository

import (
	"context"
	"github.com/Munovv/broblogo/internal/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type database struct {
	db *mongo.Database
}

func (db *database) GetAbstractDatabase() *mongo.Database {
	return db.db
}

func NewDatabase(cfg *config.Mongo) (*database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Connection))
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

	return &database{
		db: client.Database(cfg.Database),
	}, nil
}
