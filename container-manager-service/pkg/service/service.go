package service

import (
	"context"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/model"
	"github.com/Munovv/broblogo/ContainerManagerService/pkg/repository"
)

type Container interface {
	Create(ctx context.Context, cont *model.Container) (string, error)
	Get(ctx context.Context, guid string) (*model.Container, error)
	Delete(ctx context.Context, guid string) error
}

type Service struct {
	Container
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Container: NewContainerService(repo.Container),
	}
}
