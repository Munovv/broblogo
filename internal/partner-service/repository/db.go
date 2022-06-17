package repository

import (
	"context"
	"github.com/Munovv/broblogo/internal/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoDatabase(cfg config.Mongo) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Connection))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(cfg.Database), nil
}
