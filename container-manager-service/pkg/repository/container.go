package repository

import (
	"context"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContainerMongo struct {
	mongo *mongo.Database
}

func NewContainerRepository(db *mongo.Database) *ContainerMongo {
	return &ContainerMongo{
		mongo: db,
	}
}

func (r *ContainerMongo) Create(ctx context.Context, cont *model.Container) (string, error) {
	// TODO: создание записей с теми контейнерами, которых выбрал пользователь.

	return "", nil
}

func (r *ContainerMongo) Get(ctx context.Context, guid string) (*model.Container, error) {
	// TODO: получение контейнера для пользователя

	return &model.Container{}, nil
}

func (r *ContainerMongo) Remove(ctx context.Context, guid string) error {
	// TODO: удаление контейнера

	return nil
}
