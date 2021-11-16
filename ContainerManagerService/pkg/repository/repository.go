package repository

import (
	"context"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container interface {
	Create(ctx context.Context, cont *model.Container) (string, error)
	Get(ctx context.Context, guid string) (*model.Container, error)
	Remove(ctx context.Context, guid string) error
}

type Repository struct {
	Container
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Container: NewContainerRepository(db),
	}
}
